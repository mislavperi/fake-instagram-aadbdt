package models

type Filter struct {
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	DateRange   *Range    `json:"dateRange"`
	Hashtags    []*string `json:"hashtags"`
	User        *string   `json:"user"`
}

type Range struct {
	Gte *string `json:"gte"`
	Lte *string `json:"lte"`
}
