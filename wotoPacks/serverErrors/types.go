package serverErrors

type ErrorType uint8

type EndPointError struct {
	Type    int    `json:"type"`
	Message string `json:"message"`
}
