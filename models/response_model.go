package models

type PageSpeedResponseModel struct {
	CaptchaResult           string            `json:"captchaResult"`
	Kind                    string            `json:"kind"`
	ID                      string            `json:"id"`
	LoadingExperience       LoadingExperience `json:"loadingExperience"`
	OriginLoadingExperience LoadingExperience `json:"originLoadingExperience"`
	LighthouseResult        LighthouseResult  `json:"lighthouseResult"`
}
