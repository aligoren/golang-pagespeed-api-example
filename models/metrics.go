package models

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
