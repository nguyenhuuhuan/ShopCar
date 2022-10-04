package dtos

type Base struct {
	ID       int64 `json:"id"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
