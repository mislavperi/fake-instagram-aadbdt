package models

type Plan struct {
	PlanName          string `json:"planName"`
	UploadLimitSizeKb uint32 `json:"uploadLimiSizeKb"`
	DailyUploadLimit  uint32 `json:"dailyUploadLimit"`
	Cost              uint32 `json:"cost"`
}
