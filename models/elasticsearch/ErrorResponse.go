package elasticsearch

type ErrorResponse struct {
	Info *ErrorInfo `json:"error,omitempty"`
}
