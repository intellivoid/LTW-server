package serverErrors

const (
	NoError           ErrorType = iota
	UnknownError      ErrorType = iota
	ServerUnavailable ErrorType = iota + 1
)
