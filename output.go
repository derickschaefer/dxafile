package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
)

// OutputJSON writes records to JSON format with pretty-printing
// Works universally for all DXA types since it just marshals the records
// The output includes 2-space indentation for readability
func OutputJSON(w io.Writer, t DXAType, records interface{}) error {
    enc := json.NewEncoder(w)
    enc.SetIndent("", "  ") // Pretty-print with 2-space indentation
    return enc.Encode(records)
}

// OutputCSV dispatches to the appropriate CSV writer based on the detected DXA type
// Each format has different column structures, so they need specialized writers
func OutputCSV(w io.Writer, t DXAType, records interface{}) error {
    switch t {
    case DXATypeBodyComp:
        return writeCSVBodyComp(w, records.([]BodyFatRecord))
    case DXATypeTotalBody:
        return writeCSVTotalBody(w, records.([]TotalBodyRecord))
    case DXATypeCoreScan:
        return writeCSVCoreScan(w, records.([]CoreScanRecord))
    }
    return fmt.Errorf("unknown file type")
}

// ------------------------------
// CSV Writers for Each Format
// ------------------------------

// writeCSVBodyComp handles the Body Composition format
// This is the most complex format with hundreds of measurements per record
// Structure:
// - Base columns: id1, id2, id3, date
// - Mass measurements: groups of 4 columns (total, left, right, delta) per body region
// - Percent measurements: groups of 4 columns (total, left, right, delta) per body region
//
// Challenge: Different records may have different numbers of measurements,
// so we find the maximum and pad shorter records with empty strings
func writeCSVBodyComp(w io.Writer, rows []BodyFatRecord) error {
    writer := csv.NewWriter(w)

    // Start with base identifier columns
    header := []string{"id1", "id2", "id3", "date"}

    // Calculate the maximum number of measurement blocks across all records
    // This ensures our CSV has enough columns for the widest record
    maxMass := 0
    maxPct := 0
    for _, r := range rows {
        if len(r.Mass) > maxMass {
            maxMass = len(r.Mass)
        }
        if len(r.Percent) > maxPct {
            maxPct = len(r.Percent)
        }
    }

    // Generate headers for mass measurements
    // Each measurement block gets 4 columns: total, left, right, delta
    for i := 0; i < maxMass; i++ {
        header = append(header,
            fmt.Sprintf("mass_%d_total", i),
            fmt.Sprintf("mass_%d_left", i),
            fmt.Sprintf("mass_%d_right", i),
            fmt.Sprintf("mass_%d_delta", i),
        )
    }

    // Generate headers for percentage measurements
    for i := 0; i < maxPct; i++ {
        header = append(header,
            fmt.Sprintf("pct_%d_total", i),
            fmt.Sprintf("pct_%d_left", i),
            fmt.Sprintf("pct_%d_right", i),
            fmt.Sprintf("pct_%d_delta", i),
        )
    }

    // Write the header row
    writer.Write(header)

    // Write data rows
    for _, r := range rows {
        // Start each row with base identifiers
        line := []string{r.ID1, r.ID2, r.ID3, r.Date}

        // Add all mass measurements
        for _, m := range r.Mass {
            line = append(line,
                fmt.Sprintf("%f", m.Total),
                fmt.Sprintf("%f", m.Left),
                fmt.Sprintf("%f", m.Right),
                fmt.Sprintf("%f", m.Delta),
            )
        }

        // Pad with empty strings if this record has fewer mass measurements than max
        for i := len(r.Mass); i < maxMass; i++ {
            line = append(line, "", "", "", "")
        }

        // Add all percentage measurements
        for _, p := range r.Percent {
            line = append(line,
                fmt.Sprintf("%f", p.Total),
                fmt.Sprintf("%f", p.Left),
                fmt.Sprintf("%f", p.Right),
                fmt.Sprintf("%f", p.Delta),
            )
        }

        // Pad with empty strings if this record has fewer percent measurements than max
        for i := len(r.Percent); i < maxPct; i++ {
            line = append(line, "", "", "", "")
        }

        writer.Write(line)
    }

    writer.Flush()
    return writer.Error()
}

// writeCSVTotalBody handles the Total Body format
// This format contains BMD (bone mineral density) and body composition values
// Structure is simpler: base columns + variable number of numeric measurements
// The exact number of values is consistent across records in the same file
func writeCSVTotalBody(w io.Writer, rows []TotalBodyRecord) error {
    writer := csv.NewWriter(w)

    // Build header with base columns
    header := []string{
        "id1", "id2", "id3", "date",
    }

    // Add a column for each measurement value
    // Assumes the first record has the full set of measurements
    for i := range rows[0].Values {
        header = append(header, fmt.Sprintf("value_%d", i))
    }

    writer.Write(header)

    // Write each record as a row
    for _, r := range rows {
        row := []string{r.ID1, r.ID2, r.ID3, r.Date}
        
        // Append all numeric values
        for _, v := range r.Values {
            row = append(row, fmt.Sprintf("%f", v))
        }
        
        writer.Write(row)
    }

    writer.Flush()
    return writer.Error()
}

// writeCSVCoreScan handles the Core Scan format (VAT measurements)
// This is the simplest format with only 2 measurements:
// - VAT Mass (in pounds)
// - VAT Volume (in cubic inches)
// VAT = Visceral Adipose Tissue (abdominal fat around organs)
func writeCSVCoreScan(w io.Writer, rows []CoreScanRecord) error {
    writer := csv.NewWriter(w)

    // Write header with descriptive column names
    writer.Write([]string{
        "id1", "id2", "id3", "date",
        "vat_mass_lbs",   // Visceral adipose tissue mass in pounds
        "vat_volume_in3", // Visceral adipose tissue volume in cubic inches
    })

    // Write each record
    for _, r := range rows {
        row := []string{
            r.ID1, r.ID2, r.ID3, r.Date,
            fmt.Sprintf("%f", r.VATMass),
            fmt.Sprintf("%f", r.VATVolume),
        }
        writer.Write(row)
    }

    writer.Flush()
    return writer.Error()
}
