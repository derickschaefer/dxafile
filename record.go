package main

// DXAType represents the three different DEXA scanner file formats
// Each format contains different types of measurements
type DXAType int

const (
    DXATypeUnknown   DXAType = iota // Unrecognized or invalid format
    DXATypeBodyComp                  // Body Composition: fat mass & percentage by body region
    DXATypeTotalBody                 // Total Body: BMD and body composition measurements
    DXATypeCoreScan                  // Core Scan: visceral adipose tissue (VAT) measurements
)

// Measurement represents a symmetric body measurement with left/right comparison
// Used in Body Composition format for body regions (arms, legs, trunk, etc.)
type Measurement struct {
    Total float64 `json:"total"` // Combined measurement for both sides
    Left  float64 `json:"left"`  // Left side measurement
    Right float64 `json:"right"` // Right side measurement
    Delta float64 `json:"delta"` // Difference between left and right (asymmetry indicator)
}

// BodyFatRecord represents a Body Composition scan
// Contains fat mass and fat percentage for multiple body regions
// Each region has measurements for total, left, right, and delta values
// Example regions: arms, legs, trunk, android, gynoid, total body
type BodyFatRecord struct {
    ID1     string        `json:"id1"`               // Primary patient/subject identifier
    ID2     string        `json:"id2"`               // Secondary identifier
    ID3     string        `json:"id3"`               // Tertiary identifier
    Date    string        `json:"date"`              // Scan date
    Mass    []Measurement `json:"mass,omitempty"`    // Fat mass measurements by region (in grams or kg)
    Percent []Measurement `json:"percent,omitempty"` // Fat percentage measurements by region
}

// TotalBodyRecord represents a Total Body scan
// Contains bone mineral density (BMD) and comprehensive body composition values
// The exact meaning of each value depends on the column order in the source file
// Common measurements include: head BMD, arms BMD, legs BMD, trunk BMD, total BMD,
// tissue percentages, lean mass, fat mass, etc.
type TotalBodyRecord struct {
    ID1    string    `json:"id1"`    // Primary patient/subject identifier
    ID2    string    `json:"id2"`    // Secondary identifier
    ID3    string    `json:"id3"`    // Tertiary identifier
    Date   string    `json:"date"`   // Scan date
    Values []float64 `json:"values"` // Array of measurements (BMD, mass, percentages, etc.)
}

// CoreScanRecord represents a Core Scan (VAT measurement)
// Measures visceral adipose tissue - the abdominal fat that surrounds internal organs
// VAT is a key health indicator associated with metabolic syndrome and cardiovascular risk
type CoreScanRecord struct {
    ID1       string  `json:"id1"`            // Primary patient/subject identifier
    ID2       string  `json:"id2"`            // Secondary identifier
    ID3       string  `json:"id3"`            // Tertiary identifier
    Date      string  `json:"date"`           // Scan date
    VATMass   float64 `json:"vat_mass_lbs"`   // Visceral adipose tissue mass in pounds
    VATVolume float64 `json:"vat_volume_in3"` // Visceral adipose tissue volume in cubic inches
}
