package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "regexp"
    "strconv"
    "strings"

    "golang.org/x/text/encoding/unicode"
    "golang.org/x/text/transform"
)

// ErrSkipLine is returned when a line should be ignored (empty or malformed)
var ErrSkipLine = errors.New("skip")

// numericRE matches numeric values including negative numbers, decimals, and comma-separated numbers
// Examples: "123", "-45.67", "1,234.56", "+0.123"
var numericRE = regexp.MustCompile(`[-+]?\d[\d,]*\.?\d*`)

// ParseFile is the main entry point for parsing DEXA scanner files
// It handles UTF-16 LE BOM encoding and automatically detects the file format
// Returns: DXAType (format detected), interface{} (slice of records), error
func ParseFile(r io.Reader) (DXAType, interface{}, error) {
    // Create a UTF-16 Little Endian decoder that expects a BOM (Byte Order Mark)
    // DEXA scanner files use UTF-16 LE encoding with BOM
    utf16bom := unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM)
    reader := transform.NewReader(r, utf16bom.NewDecoder())
    scanner := bufio.NewScanner(reader)

    lineNum := 0
    header := ""

    // Find the first non-empty line, which contains the header
    // The header determines the file type (BodyComp, TotalBody, or CoreScan)
    for scanner.Scan() {
        lineNum++
        raw := strings.TrimSpace(scanner.Text())
        if raw == "" {
            continue // Skip empty lines
        }
        header = raw
        break // Found header, exit loop
    }

    // Validate that we found a header
    if header == "" {
        return DXATypeUnknown, nil, fmt.Errorf("empty file")
    }

    // Detect which DEXA format this file contains based on header content
    t := detectDXAType(header)

    // Initialize slices to hold records for each possible file type
    bodycomp := []BodyFatRecord{}   // For body composition (fat mass/percent)
    totalbody := []TotalBodyRecord{} // For total body BMD measurements
    corescan := []CoreScanRecord{}   // For visceral adipose tissue (VAT) scans

    // Parse all remaining data lines
    for scanner.Scan() {
        lineNum++
        raw := strings.TrimSpace(scanner.Text())
        if raw == "" {
            continue // Skip empty lines
        }

        // Parse the data line according to detected file type
        rec, err := parseDataLine(t, raw)
        if err != nil {
            if errors.Is(err, ErrSkipLine) {
                continue // Skip lines that are intentionally ignored
            }
            // Return error with line number for debugging
            return DXATypeUnknown, nil, fmt.Errorf("line %d: %w", lineNum, err)
        }

        // Append parsed record to the appropriate slice based on detected type
        switch t {
        case DXATypeBodyComp:
            bodycomp = append(bodycomp, rec.(BodyFatRecord))
        case DXATypeTotalBody:
            totalbody = append(totalbody, rec.(TotalBodyRecord))
        case DXATypeCoreScan:
            corescan = append(corescan, rec.(CoreScanRecord))
        }
    }

    // Check for any scanner errors (I/O issues, etc.)
    if err := scanner.Err(); err != nil {
        return DXATypeUnknown, nil, err
    }

    // Return the appropriate slice based on detected file type
    switch t {
    case DXATypeBodyComp:
        return t, bodycomp, nil
    case DXATypeTotalBody:
        return t, totalbody, nil
    case DXATypeCoreScan:
        return t, corescan, nil
    }

    return DXATypeUnknown, nil, fmt.Errorf("unrecognized file type")
}

// detectDXAType examines the header row to determine which DEXA format the file contains
// Detection is based on distinctive column names unique to each format:
// - BodyComp: contains "arms fat mass" 
// - TotalBody: contains "head bmd" (bone mineral density)
// - CoreScan: contains "vat mass" (visceral adipose tissue)
func detectDXAType(header string) DXAType {
    h := strings.ToLower(header)

    switch {
    case strings.Contains(h, "arms fat mass"):
        return DXATypeBodyComp
    case strings.Contains(h, "head bmd"):
        return DXATypeTotalBody
    case strings.Contains(h, "vat mass"):
        return DXATypeCoreScan
    }
    return DXATypeUnknown
}

// parseDataLine parses a single tab-delimited data row based on the detected file type
// All DEXA formats share the first 4 columns: ID1, ID2, ID3, Date
// Remaining columns vary by format and contain measurement data
func parseDataLine(t DXAType, line string) (interface{}, error) {
    // Split by tab character (DEXA files are tab-delimited)
    fields := strings.Split(line, "\t")
    
    // Require at least 4 fields (the common ID/date columns)
    if len(fields) < 4 {
        return nil, ErrSkipLine
    }

    // Extract the common fields present in all formats
    id1 := strings.TrimSpace(fields[0])  // Patient/Subject ID
    id2 := strings.TrimSpace(fields[1])  // Secondary ID
    id3 := strings.TrimSpace(fields[2])  // Tertiary ID
    date := strings.TrimSpace(fields[3]) // Scan date

    switch t {

    // BODY COMPOSITION FORMAT
    // Contains fat mass and fat percentage for body regions (arms, legs, trunk, etc.)
    // Data structure: first half is mass measurements, second half is percentages
    // Each measurement has 4 values: Total, Left, Right, Delta
    case DXATypeBodyComp:
        // Extract all numeric values from remaining columns
        nums := extractNumbers(fields[4:])
        if len(nums) == 0 {
            return nil, ErrSkipLine
        }

        rec := BodyFatRecord{ID1: id1, ID2: id2, ID3: id3, Date: date}
        
        // Split numbers in half: first half = mass, second half = percent
        rec.Mass = groupMeasurements(nums[:len(nums)/2])
        rec.Percent = groupMeasurements(nums[len(nums)/2:])
        return rec, nil

    // TOTAL BODY FORMAT
    // Contains bone mineral density (BMD) and other body composition values
    // Simple flat array of measurements
    case DXATypeTotalBody:
        nums := extractNumbers(fields[4:])
        rec := TotalBodyRecord{ID1: id1, ID2: id2, ID3: id3, Date: date, Values: nums}
        return rec, nil

    // CORE SCAN FORMAT (VAT - Visceral Adipose Tissue)
    // Contains exactly 2 measurements: VAT mass (lbs) and VAT volume (in³)
    case DXATypeCoreScan:
        nums := extractNumbers(fields[4:])
        if len(nums) < 2 {
            return nil, fmt.Errorf("corescan row missing numeric fields")
        }
        rec := CoreScanRecord{
            ID1:       id1,
            ID2:       id2,
            ID3:       id3,
            Date:      date,
            VATMass:   nums[0], // Visceral adipose tissue mass in pounds
            VATVolume: nums[1], // Visceral adipose tissue volume in cubic inches
        }
        return rec, nil
    }

    return nil, ErrSkipLine
}

// extractNumbers extracts all numeric values from an array of field strings
// Uses regex to find numbers that may include:
// - Negative/positive signs
// - Comma thousands separators (e.g., "1,234.56")
// - Decimal points
// Returns a flat array of float64 values
func extractNumbers(fields []string) []float64 {
    vals := []float64{}

    for _, f := range fields {
        // Find all numeric patterns in this field
        matches := numericRE.FindAllString(f, -1)
        for _, m := range matches {
            // Remove comma separators before parsing
            m = strings.ReplaceAll(m, ",", "")
            if m == "" {
                continue
            }
            // Convert string to float64
            if v, err := strconv.ParseFloat(m, 64); err == nil {
                vals = append(vals, v)
            }
        }
    }

    return vals
}

// groupMeasurements converts a flat array of numbers into structured Measurement blocks
// Each measurement consists of 4 consecutive values:
// - Total: total measurement for both sides
// - Left: left side measurement
// - Right: right side measurement  
// - Delta: difference between left and right
// This pattern repeats for each body region (arms, legs, trunk, etc.)
func groupMeasurements(nums []float64) []Measurement {
    out := []Measurement{}
    
    // Process numbers in groups of 4
    for len(nums) >= 4 {
        out = append(out, Measurement{
            Total: nums[0],
            Left:  nums[1],
            Right: nums[2],
            Delta: nums[3],
        })
        nums = nums[4:] // Advance to next group
    }
    
    return out
}
