package domain

import "time"

type Request struct {
	ID        uint      `json:"id"         gorm:"primary_key"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Timestamp time.Time `json:"timestamp"`
}
