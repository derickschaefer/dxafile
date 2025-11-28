package main

type DXAType int

const (
    DXATypeUnknown DXAType = iota
    DXATypeBodyComp
    DXATypeTotalBody
    DXATypeCoreScan
)

type Measurement struct {
    Total float64 `json:"total"`
    Left  float64 `json:"left"`
    Right float64 `json:"right"`
    Delta float64 `json:"delta"`
}

type BodyFatRecord struct {
    ID1     string        `json:"id1"`
    ID2     string        `json:"id2"`
    ID3     string        `json:"id3"`
    Date    string        `json:"date"`
    Mass    []Measurement `json:"mass,omitempty"`
    Percent []Measurement `json:"percent,omitempty"`
}

type TotalBodyRecord struct {
    ID1    string    `json:"id1"`
    ID2    string    `json:"id2"`
    ID3    string    `json:"id3"`
    Date   string    `json:"date"`
    Values []float64 `json:"values"`
}

type CoreScanRecord struct {
    ID1       string  `json:"id1"`
    ID2       string  `json:"id2"`
    ID3       string  `json:"id3"`
    Date      string  `json:"date"`
    VATMass   float64 `json:"vat_mass_lbs"`
    VATVolume float64 `json:"vat_volume_in3"`
}
