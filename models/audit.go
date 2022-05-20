package models

type AuditItem struct {
	ID               string       `json:"id"`
	Title            string       `json:"title"`
	Description      string       `json:"description"`
	Score            float64      `json:"score"`
	ScoreDisplayMode string       `json:"ScoreDisplayMode"`
	DisplayValue     string       `json:"displayValue"`
	Details          AuditDetails `json:"details"`
}

type AuditDetails struct {
	Items    []AuditDetailItems    `json:"items"`
	Type     string                `json:"type,omitempty"`
	Headings []AuditDetailHeadings `json:"headings"`
}

type AuditDetailItems struct {
	Source        AuditDetailItemSource `json:"source"`
	Node          AuditDetailItemNode   `json:"node"`
	Score         float64               `json:"score,omitempty"`
	Url           string                `json:"url,omitempty"`
	WastedPercent float64               `json:"wastedPercent,omitempty"`
	TotalBytes    float64               `json:"totalBytes,omitempty"`
	WastedBytes   float64               `json:"wastedBytes,omitempty"`
}

type AuditDetailItemSource struct {
	Url         string  `json:"url,omitempty"`
	UrlProvider string  `json:"urlProvider,omitempty"`
	Column      float64 `json:"column,omitempty"`
	Type        string  `json:"type,omitempty"`
	Line        float64 `json:"line,omitempty"`
}

type AuditDetailItemNode struct {
	Node         string                          `json:"node,omitempty"`
	NodeLabel    string                          `json:"nodeLabel,omitempty"`
	Path         string                          `json:"path,omitempty"`
	Snippet      string                          `json:"snippet,omitempty"`
	LhID         string                          `json:"lhId,omitempty"`
	Selector     string                          `json:"selector,omitempty"`
	BoundingRect AuditDetailItemNodeBoundingRect `json:"boundingRect,omitempty"`
}

type AuditDetailItemNodeBoundingRect struct {
	Top    float64 `json:"top"`
	Width  float64 `json:"width"`
	Right  float64 `json:"right"`
	Left   float64 `json:"left"`
	Bottom float64 `json:"bottom"`
	Height float64 `json:"height"`
}

type AuditDetailHeadings struct {
	Text        string  `json:"text"`
	ItemType    string  `json:"itemType"`
	Granularity float64 `json:"granularity"`
	Key         string  `json:"key"`
	Label       string  `json:"label"`
	ValueType   string  `json:"valueType"`
}
