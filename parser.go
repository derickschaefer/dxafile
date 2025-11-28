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

var ErrSkipLine = errors.New("skip")
var numericRE = regexp.MustCompile(`[-+]?\d[\d,]*\.?\d*`)

//
// MAIN ENTRY POINT
//
func ParseFile(r io.Reader) (DXAType, interface{}, error) {
    // decode UTF-16 LE
    utf16bom := unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM)
    reader := transform.NewReader(r, utf16bom.NewDecoder())
    scanner := bufio.NewScanner(reader)

    lineNum := 0
    header := ""

    // First non-empty row is the header
    for scanner.Scan() {
        lineNum++
        raw := strings.TrimSpace(scanner.Text())
        if raw == "" {
            continue
        }
        header = raw
        break
    }

    if header == "" {
        return DXATypeUnknown, nil, fmt.Errorf("empty file")
    }

    // Detect file type
    t := detectDXAType(header)

    // Prepare slices
    bodycomp := []BodyFatRecord{}
    totalbody := []TotalBodyRecord{}
    corescan := []CoreScanRecord{}

    // Parse remaining rows
    for scanner.Scan() {
        lineNum++
        raw := strings.TrimSpace(scanner.Text())
        if raw == "" {
            continue
        }

        rec, err := parseDataLine(t, raw)
        if err != nil {
            if errors.Is(err, ErrSkipLine) {
                continue
            }
            return DXATypeUnknown, nil, fmt.Errorf("line %d: %w", lineNum, err)
        }

        switch t {
        case DXATypeBodyComp:
            bodycomp = append(bodycomp, rec.(BodyFatRecord))
        case DXATypeTotalBody:
            totalbody = append(totalbody, rec.(TotalBodyRecord))
        case DXATypeCoreScan:
            corescan = append(corescan, rec.(CoreScanRecord))
        }
    }

    if err := scanner.Err(); err != nil {
        return DXATypeUnknown, nil, err
    }

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

//
// FILE TYPE DETECTION
//
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

//
// PARSE A SINGLE ROW BASED ON FILE TYPE
//
func parseDataLine(t DXAType, line string) (interface{}, error) {
    fields := strings.Split(line, "\t")
    if len(fields) < 4 {
        return nil, ErrSkipLine
    }

    id1 := strings.TrimSpace(fields[0])
    id2 := strings.TrimSpace(fields[1])
    id3 := strings.TrimSpace(fields[2])
    date := strings.TrimSpace(fields[3])

    switch t {

    //
    // BODY COMP
    //
    case DXATypeBodyComp:
        nums := extractNumbers(fields[4:])
        if len(nums) == 0 {
            return nil, ErrSkipLine
        }

        rec := BodyFatRecord{ID1: id1, ID2: id2, ID3: id3, Date: date}
        rec.Mass = groupMeasurements(nums[:len(nums)/2])
        rec.Percent = groupMeasurements(nums[len(nums)/2:])
        return rec, nil

    //
    // TOTAL BODY
    //
    case DXATypeTotalBody:
        nums := extractNumbers(fields[4:])
        rec := TotalBodyRecord{ID1: id1, ID2: id2, ID3: id3, Date: date, Values: nums}
        return rec, nil

    //
    // CORESCAN (VAT)
    //
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
            VATMass:   nums[0],
            VATVolume: nums[1],
        }
        return rec, nil
    }

    return nil, ErrSkipLine
}

//
// EXTRACT NUMBERS FROM A SET OF FIELDS
//
func extractNumbers(fields []string) []float64 {
    vals := []float64{}

    for _, f := range fields {
        matches := numericRE.FindAllString(f, -1)
        for _, m := range matches {
            m = strings.ReplaceAll(m, ",", "")
            if m == "" {
                continue
            }
            if v, err := strconv.ParseFloat(m, 64); err == nil {
                vals = append(vals, v)
            }
        }
    }

    return vals
}

//
// GROUP NUMBERS INTO MEASUREMENT BLOCKS (4 numbers per block)
//
func groupMeasurements(nums []float64) []Measurement {
    out := []Measurement{}
    for len(nums) >= 4 {
        out = append(out, Measurement{
            Total: nums[0],
            Left:  nums[1],
            Right: nums[2],
            Delta: nums[3],
        })
        nums = nums[4:]
    }
    return out
}
