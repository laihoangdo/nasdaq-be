package models

type List[T any] struct {
	Pagination *Pagination
	Data       []T
}

type Pagination struct {
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	HasMore    bool  `json:"has_more"`
}
