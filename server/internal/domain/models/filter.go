package models

type Filter struct {
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	TimeRange   *Range    `json:"timerange"`
	Hashtags    []*string `json:"hashtags"`
	User        *string   `json:"user"`
}

type Range struct {
	Gte *uint64 `json:"gte"`
	Lte *uint64 `json:"lte"`
}
