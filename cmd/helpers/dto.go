package helpers

type ErrorMessage struct {
	ErrorMessage string `json:"error_message"`
	InvalidField string `json:"invalid_field;omitempty"`
}
