package models

// Komisi represents the commission calculation for each marketing.
type Komisi struct {
	MarketingId   int     `json:"marketing_id"`
	Bulan         string  `json:"bulan"`
	Omzet         float64 `json:"omzet"`
	KomisiPersen  float64 `json:"komisi_persen"`
	KomisiNominal float64 `json:"komisi_nominal"`
}
