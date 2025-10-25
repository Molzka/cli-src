package api

type APIErrorResponse struct {
	ErrorCode string            `json:"error_code"`
	Message   string            `json:"message"`
	RequestID string            `json:"request_id"`
	Details   map[string]string `json:"details,omitempty"`
}
