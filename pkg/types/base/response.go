package base

type Pagination struct {
	Page           int `json:"page"`
	Limit          int `json:"limit"`
	TotalDocuments int `json:"total_documents"`
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response[T any] struct {
	Data       *T             `json:"data,omitempty"`
	Pagination *Pagination    `json:"pagination,omitempty"`
	Error      *ErrorResponse `json:"error,omitempty"`
}
