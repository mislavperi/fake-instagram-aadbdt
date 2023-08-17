package models

import (
	"time"
)

type Log struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	User      User      `json:"user"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
}
