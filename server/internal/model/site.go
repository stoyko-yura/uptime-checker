package model

import (
	"time"

	"gorm.io/gorm"
)

type Site struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	URL         string         `gorm:"not null" json:"url"`
	Name        string         `json:"name"`
	IntervalSec int            `gorm:"default:60" json:"interval_sec"`
	LastStatus  int            `json:"last_status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Check struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	SiteID         uint      `json:"site_id"`
	StatusCode     int       `json:"status_code"`
	ResponseTimeMs int64     `json:"response_time_ms"`
	CheckedAt      time.Time `json:"checked_at"`
}
