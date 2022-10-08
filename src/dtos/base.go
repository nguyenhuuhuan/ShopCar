package dtos

import "time"

type Base struct {
	ID                 int64 `json:"id"`
	CreatedAtTimestamp int64 `json:"created_at"`
	UpdatedAtTimestamp int64 `json:"updated_at"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CreatedTimestamp Convert to timestamp
func (b *Base) CreatedAt(CreatedAt time.Time) {
	b.CreatedAtTimestamp = CreatedAt.Unix()
}

func (b *Base) UpdatedAt(UpdatedAt time.Time) {
	b.UpdatedAtTimestamp = UpdatedAt.Unix()
}
