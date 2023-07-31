package models

type Plan struct {
	PlanName          string
	UploadLimitSizeKb uint32
	DailyUploadLimit  uint32
	Cost              uint32
}
