package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "strings"
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

// BODY COMP — Body Composition with friendly column names
func writeCSVBodyComp(w io.Writer, rows []BodyFatRecord) error {
    writer := csv.NewWriter(w)

    // Define mass measurement labels in the order they appear
    massLabels := []string{
        "Arms_Bone_Mass",
        "Legs_Bone_Mass",
        "Trunk_Bone_Mass",
        "Android_Bone_Mass",
        "Gynoid_Bone_Mass",
        "Total_Bone_Mass",
        "TBLH_Bone_Mass",
        "Arms_Fat_Mass",
        "Legs_Fat_Mass",
        "Trunk_Fat_Mass",
        "Android_Fat_Mass",
        "Gynoid_Fat_Mass",
        "Total_Fat_Mass",
        "TBLH_Fat_Mass",
        "Arms_Lean_Mass",
        "Legs_Lean_Mass",
        "Trunk_Lean_Mass",
        "Android_Lean_Mass",
        "Gynoid_Lean_Mass",
        "Total_Lean_Mass",
        "TBLH_Lean_Mass",
        "Arms_Tissue_Mass",
        "Legs_Tissue_Mass",
        "Trunk_Tissue_Mass",
        "Android_Tissue_Mass",
        "Gynoid_Tissue_Mass",
        "Total_Tissue_Mass",
        "TBLH_Tissue_Mass",
        "Arms_Fat_Free_Mass",
        "Legs_Fat_Free_Mass",
        "Trunk_Fat_Free_Mass",
        "Android_Fat_Free_Mass",
        "Gynoid_Fat_Free_Mass",
        "Total_Fat_Free_Mass",
        "TBLH_Fat_Free_Mass",
        "Arms_Total_Mass",
        "Legs_Total_Mass",
        "Trunk_Total_Mass",
        "Android_Total_Mass",
        "Gynoid_Total_Mass",
        "Total_Total_Mass",
        "TBLH_Total_Mass",
    }

    // Define percentage measurement labels in the order they appear
    percentLabels := []string{
        "Arms_Region_Percent_Fat",
        "Legs_Region_Percent_Fat",
        "Trunk_Region_Percent_Fat",
        "Android_Region_Percent_Fat",
        "Gynoid_Region_Percent_Fat",
        "Total_Region_Percent_Fat",
        "TBLH_Region_Percent_Fat",
        "Arms_Tissue_Percent_Fat",
        "Legs_Tissue_Percent_Fat",
        "Trunk_Tissue_Percent_Fat",
        "Android_Tissue_Percent_Fat",
        "Gynoid_Tissue_Percent_Fat",
        "Total_Tissue_Percent_Fat",
        "TBLH_Tissue_Percent_Fat",
    }

    // Start with base identifier columns
    header := []string{"Last_Name", "First_Name", "Patient_ID", "Measure_Date"}

    // Calculate the maximum number of measurement blocks across all records
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

    // Generate headers for mass measurements with friendly names
    for i := 0; i < maxMass; i++ {
        label := ""
        if i < len(massLabels) {
            label = massLabels[i]
        } else {
            label = fmt.Sprintf("Mass_%d", i)
        }
        header = append(header,
            label+"_Total",
            label+"_Left",
            label+"_Right",
            label+"_Delta",
        )
    }

    // Generate headers for percentage measurements with friendly names
    for i := 0; i < maxPct; i++ {
        label := ""
        if i < len(percentLabels) {
            label = percentLabels[i]
        } else {
            label = fmt.Sprintf("Percent_%d", i)
        }
        header = append(header,
            label+"_Total",
            label+"_Left",
            label+"_Right",
            label+"_Delta",
        )
    }

    // Write the header row
    writer.Write(header)

    // Write data rows
    for _, r := range rows {
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

// TOTAL BODY — BMD measurements with friendly column names
func writeCSVTotalBody(w io.Writer, rows []TotalBodyRecord) error {
    writer := csv.NewWriter(w)

    // Define all Total Body measurement labels in exact order
    totalBodyLabels := []string{
        // BMD measurements (0-16)
        "Head_BMD",
        "Arms_BMD",
        "Legs_BMD",
        "Trunk_BMD",
        "Ribs_BMD",
        "Pelvis_BMD",
        "Spine_BMD",
        "Arm_Left_BMD",
        "Leg_Left_BMD",
        "Arm_Right_BMD",
        "Leg_Right_BMD",
        "Total_BMD",
        "TBLH_BMD",
        "Trunk_Left_BMD",
        "Total_Left_BMD",
        "Trunk_Right_BMD",
        "Total_Right_BMD",
        // BMC measurements (17-33)
        "Head_BMC",
        "Arms_BMC",
        "Legs_BMC",
        "Trunk_BMC",
        "Ribs_BMC",
        "Pelvis_BMC",
        "Spine_BMC",
        "Arm_Left_BMC",
        "Leg_Left_BMC",
        "Arm_Right_BMC",
        "Leg_Right_BMC",
        "Total_BMC",
        "TBLH_BMC",
        "Trunk_Left_BMC",
        "Total_Left_BMC",
        "Trunk_Right_BMC",
        "Total_Right_BMC",
        // Area measurements (34-50)
        "Head_Area",
        "Arms_Area",
        "Legs_Area",
        "Trunk_Area",
        "Ribs_Area",
        "Pelvis_Area",
        "Spine_Area",
        "Arm_Left_Area",
        "Leg_Left_Area",
        "Arm_Right_Area",
        "Leg_Right_Area",
        "Total_Area",
        "TBLH_Area",
        "Trunk_Left_Area",
        "Total_Left_Area",
        "Trunk_Right_Area",
        "Total_Right_Area",
        // T-Scores (51-67)
        "Head_T_Score",
        "Arms_T_Score",
        "Legs_T_Score",
        "Trunk_T_Score",
        "Ribs_T_Score",
        "Pelvis_T_Score",
        "Spine_T_Score",
        "Arm_Left_T_Score",
        "Leg_Left_T_Score",
        "Arm_Right_T_Score",
        "Leg_Right_T_Score",
        "Total_T_Score",
        "TBLH_T_Score",
        "Trunk_Left_T_Score",
        "Total_Left_T_Score",
        "Trunk_Right_T_Score",
        "Total_Right_T_Score",
        // Z-Scores (68-84)
        "Head_Z_Score",
        "Arms_Z_Score",
        "Legs_Z_Score",
        "Trunk_Z_Score",
        "Ribs_Z_Score",
        "Pelvis_Z_Score",
        "Spine_Z_Score",
        "Arm_Left_Z_Score",
        "Leg_Left_Z_Score",
        "Arm_Right_Z_Score",
        "Leg_Right_Z_Score",
        "Total_Z_Score",
        "TBLH_Z_Score",
        "Trunk_Left_Z_Score",
        "Total_Left_Z_Score",
        "Trunk_Right_Z_Score",
        "Total_Right_Z_Score",
        // Average Height (85-101)
        "Head_Average_Height",
        "Arms_Average_Height",
        "Legs_Average_Height",
        "Trunk_Average_Height",
        "Ribs_Average_Height",
        "Pelvis_Average_Height",
        "Spine_Average_Height",
        "Arm_Left_Average_Height",
        "Leg_Left_Average_Height",
        "Arm_Right_Average_Height",
        "Leg_Right_Average_Height",
        "Total_Average_Height",
        "TBLH_Average_Height",
        "Trunk_Left_Average_Height",
        "Total_Left_Average_Height",
        "Trunk_Right_Average_Height",
        "Total_Right_Average_Height",
        // Average Width (102-118)
        "Head_Average_Width",
        "Arms_Average_Width",
        "Legs_Average_Width",
        "Trunk_Average_Width",
        "Ribs_Average_Width",
        "Pelvis_Average_Width",
        "Spine_Average_Width",
        "Arm_Left_Average_Width",
        "Leg_Left_Average_Width",
        "Arm_Right_Average_Width",
        "Leg_Right_Average_Width",
        "Total_Average_Width",
        "TBLH_Average_Width",
        "Trunk_Left_Average_Width",
        "Total_Left_Average_Width",
        "Trunk_Right_Average_Width",
        "Total_Right_Average_Width",
    }

    // Build header with base columns
    header := []string{"Last_Name", "First_Name", "Patient_ID", "Measure_Date"}

    // Add friendly column names for each measurement
    for i := range rows[0].Values {
        if i < len(totalBodyLabels) {
            header = append(header, totalBodyLabels[i])
        } else {
            // Fallback for any additional values beyond our label list
            header = append(header, fmt.Sprintf("Value_%d", i))
        }
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

// CORE SCAN — VAT measurements (already has friendly names)
func writeCSVCoreScan(w io.Writer, rows []CoreScanRecord) error {
    writer := csv.NewWriter(w)

    // Write header with friendly column names
    writer.Write([]string{
        "Last_Name",
        "First_Name",
        "Patient_ID",
        "Measure_Date",
        "VAT_Mass_lbs",
        "VAT_Volume_in3",
    })

    // Write each record
    for _, r := range rows {
        row := []string{
            r.ID1,
            r.ID2,
            r.ID3,
            r.Date,
            fmt.Sprintf("%f", r.VATMass),
            fmt.Sprintf("%f", r.VATVolume),
        }
        writer.Write(row)
    }

    writer.Flush()
    return writer.Error()
}

// sanitizeColumnName converts header text to friendly underscore-separated names
// Example: "Arms Fat Mass" -> "Arms_Fat_Mass"
//          "Region %Fat" -> "Region_Percent_Fat"
func sanitizeColumnName(name string) string {
    // Replace common symbols
    name = strings.ReplaceAll(name, "%", "Percent_")
    name = strings.ReplaceAll(name, " ", "_")
    name = strings.ReplaceAll(name, "-", "_")
    name = strings.ReplaceAll(name, "/", "_")
    name = strings.ReplaceAll(name, "(", "")
    name = strings.ReplaceAll(name, ")", "")
    
    // Remove any double underscores
    for strings.Contains(name, "__") {
        name = strings.ReplaceAll(name, "__", "_")
    }
    
    // Trim trailing underscores
    name = strings.Trim(name, "_")
    
    return name
}
