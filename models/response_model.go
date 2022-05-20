package models

type PageSpeedResponseModel struct {
	CaptchaResult           string            `json:"captchaResult"`
	Kind                    string            `json:"kind"`
	ID                      string            `json:"id"`
	LoadingExperience       LoadingExperience `json:"loadingExperience"`
	OriginLoadingExperience LoadingExperience `json:"originLoadingExperience"`
	LighthouseResult        LighthouseResult  `json:"lighthouseResult"`
}

type LighthouseResult struct {
	RequestedUrl      string        `json:"requestedUrl"`
	FinalUrl          string        `json:"finalUrl"`
	LighthouseVersion string        `json:"lighthouseVersion"`
	UserAgent         string        `json:"userAgent"`
	FetchTime         string        `json:"fetchTime"`
	Environment       Environment   `json:"environment"`
	RunWarnings       []RunWarnings `json:"runWarnings"`
}

type RunWarnings struct {
}

type Environment struct {
	NetworkUserAgent string `json:"networkUserAgent"`
	HostUserAgent    string `json:"hostUserAgent"`
	BenchmarkIndex   int    `json:"benchmarkIndex"`
}

type LoadingExperience struct {
	ID              string  `json:"id"`
	Metrics         Metrics `json:"metrics"`
	OverallCategory string  `json:"overall_category"`
	InitialUrl      string  `json:"initial_url"`
	OriginFallback  bool    `json:"origin_fallback"`
}

type Metrics struct {
	CumulativeLayoutShiftScore         MetricData `json:"CUMULATIVE_LAYOUT_SHIFT_SCORE"`
	ExperimentalInteractionToNextPaint MetricData `json:"EXPERIMENTAL_INTERACTION_TO_NEXT_PAINT"`
	ExperimentalTimeToFirstByte        MetricData `json:"EXPERIMENTAL_TIME_TO_FIRST_BYTE"`
	FirstContentfulPaintMs             MetricData `json:"FIRST_CONTENTFUL_PAINT_MS"`
	LargestContentfulPaintMs           MetricData `json:"LARGEST_CONTENTFUL_PAINT_MS"`
}

type MetricData struct {
	Percentile    float64         `json:"percentile"`
	Distributions []Distributions `json:"distributions"`
	Category      string          `json:"category"`
}

type Distributions struct {
	Min        float64 `json:"min"`
	Max        float64 `json:"max"`
	Proportion float64 `json:"proportion"`
}
