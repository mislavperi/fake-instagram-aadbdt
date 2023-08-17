package models

type Plan struct {
	ID                int64  `json:"planID"`
	PlanName          string `json:"planName"`
	UploadLimitSizeKb uint32 `json:"uploadLimitSizeKb"`
	DailyUploadLimit  uint32 `json:"dailyUploadLimit"`
	Cost              uint32 `json:"cost"`
}
