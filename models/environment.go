package models

type Environment struct {
	NetworkUserAgent string  `json:"networkUserAgent"`
	HostUserAgent    string  `json:"hostUserAgent"`
	BenchmarkIndex   float64 `json:"benchmarkIndex"`
}
