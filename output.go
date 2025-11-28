package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
)

//
// JSON OUTPUT — Works for all DXA types
//
func OutputJSON(w io.Writer, t DXAType, records interface{}) error {
    enc := json.NewEncoder(w)
    enc.SetIndent("", "  ")
    return enc.Encode(records)
}

//
// CSV OUTPUT — Dispatch to specific CSV writers by file type
//
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

//
// ------------------------------
// CSV Writers for Each Format
// ------------------------------
//

// BODY COMP — Very wide dataset (hundreds of measurements)
func writeCSVBodyComp(w io.Writer, rows []BodyFatRecord) error {
    writer := csv.NewWriter(w)

    header := []string{"id1", "id2", "id3", "date"}

    // Identify maximum number of measurement blocks
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

    // Mass headers
    for i := 0; i < maxMass; i++ {
        header = append(header,
            fmt.Sprintf("mass_%d_total", i),
            fmt.Sprintf("mass_%d_left", i),
            fmt.Sprintf("mass_%d_right", i),
            fmt.Sprintf("mass_%d_delta", i),
        )
    }

    // Percent headers
    for i := 0; i < maxPct; i++ {
        header = append(header,
            fmt.Sprintf("pct_%d_total", i),
            fmt.Sprintf("pct_%d_left", i),
            fmt.Sprintf("pct_%d_right", i),
            fmt.Sprintf("pct_%d_delta", i),
        )
    }

    writer.Write(header)

    // Rows
    for _, r := range rows {
        line := []string{r.ID1, r.ID2, r.ID3, r.Date}

        for _, m := range r.Mass {
            line = append(line,
                fmt.Sprintf("%f", m.Total),
                fmt.Sprintf("%f", m.Left),
                fmt.Sprintf("%f", m.Right),
                fmt.Sprintf("%f", m.Delta),
            )
        }

        // pad mass if needed
        for i := len(r.Mass); i < maxMass; i++ {
            line = append(line, "", "", "", "")
        }

        for _, p := range r.Percent {
            line = append(line,
                fmt.Sprintf("%f", p.Total),
                fmt.Sprintf("%f", p.Left),
                fmt.Sprintf("%f", p.Right),
                fmt.Sprintf("%f", p.Delta),
            )
        }

        // pad percent if needed
        for i := len(r.Percent); i < maxPct; i++ {
            line = append(line, "", "", "", "")
        }

        writer.Write(line)
    }

    writer.Flush()
    return writer.Error()
}

// TOTAL BODY — Compact table
func writeCSVTotalBody(w io.Writer, rows []TotalBodyRecord) error {
    writer := csv.NewWriter(w)

    header := []string{
        "id1", "id2", "id3", "date",
    }

    for i := range rows[0].Values {
        header = append(header, fmt.Sprintf("value_%d", i))
    }

    writer.Write(header)

    for _, r := range rows {
        row := []string{r.ID1, r.ID2, r.ID3, r.Date}
        for _, v := range r.Values {
            row = append(row, fmt.Sprintf("%f", v))
        }
        writer.Write(row)
    }

    writer.Flush()
    return writer.Error()
}

// CORE SCAN — Simple 2-column numeric record
func writeCSVCoreScan(w io.Writer, rows []CoreScanRecord) error {
    writer := csv.NewWriter(w)

    writer.Write([]string{
        "id1", "id2", "id3", "date",
        "vat_mass_lbs", "vat_volume_in3",
    })

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
