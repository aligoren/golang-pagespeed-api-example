package models

type AuditItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	ScoreDisplayMode string       `json:"ScoreDisplayMode"`
	DisplayValue     string       `json:"displayValue"`
	Score            float64      `json:"score"`
	Details          AuditDetails `json:"details"`
}

type AuditDetails struct {
	Type     string                `json:"type,omitempty"`
	Items    []AuditDetailItems    `json:"items"`
	Headings []AuditDetailHeadings `json:"headings"`
	Nodes    []AuditDetailNodes    `json:"nodes"`
}

type AuditDetailItems struct {
	Score           float64               `json:"score,omitempty"`
	WastedPercent   float64               `json:"wastedPercent,omitempty"`
	TotalBytes      float64               `json:"totalBytes,omitempty"`
	WastedBytes     float64               `json:"wastedBytes,omitempty"`
	WastedWebpBytes float64               `json:"wastedWebpBytes,omitempty"`
	Url             string                `json:"url,omitempty"`
	IsCrossOrigin   bool                  `json:"isCrossOrigin"`
	FromProtocol    bool                  `json:"fromProtocol"`
	Source          AuditDetailItemSource `json:"source"`
	Node            AuditDetailItemNode   `json:"node"`
}

type AuditDetailItemSource struct {
	Column      float64 `json:"column,omitempty"`
	Line        float64 `json:"line,omitempty"`
	Url         string  `json:"url,omitempty"`
	UrlProvider string  `json:"urlProvider,omitempty"`
	Type        string  `json:"type,omitempty"`
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
	Granularity float64 `json:"granularity"`
	Text        string  `json:"text"`
	ItemType    string  `json:"itemType"`
	Key         string  `json:"key"`
	Label       string  `json:"label"`
	ValueType   string  `json:"valueType"`
}

type AuditDetailNodes struct {
	ResourceBytes float64 `json:"resourceBytes"`
	UnusedBytes   float64 `json:"unusedBytes"`
	Name          string  `json:"name"`
}
