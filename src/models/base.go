package models

import "time"

type Base struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreatedTimestamp Convert to timestamp
func (b *Base) CreatedTimestamp() int64 {
	return b.CreatedAt.Unix()
}
