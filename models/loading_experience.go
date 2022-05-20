package models

type LoadingExperience struct {
	ID              string  `json:"id"`
	Metrics         Metrics `json:"metrics"`
	OverallCategory string  `json:"overall_category"`
	InitialUrl      string  `json:"initial_url"`
	OriginFallback  bool    `json:"origin_fallback"`
}
