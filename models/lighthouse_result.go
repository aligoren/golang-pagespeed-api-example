package models

type LighthouseResult struct {
	RequestedUrl      string               `json:"requestedUrl"`
	FinalUrl          string               `json:"finalUrl"`
	LighthouseVersion string               `json:"lighthouseVersion"`
	UserAgent         string               `json:"userAgent"`
	FetchTime         string               `json:"fetchTime"`
	Environment       Environment          `json:"environment"`
	RunWarnings       []RunWarnings        `json:"runWarnings"`
	Audits            map[string]AuditItem `json:"audits"`
}
