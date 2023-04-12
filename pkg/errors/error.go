package errors

type Error interface {
	Message() string
	Code() int
	Status() string
}

type ErrorData struct {
	MessageData string `json:"message"`
	CodeData    int    `json:"code"`
	StatusData  string `json:"status"`
}

func (e ErrorData) Message() string {
	return e.MessageData
}

func (e ErrorData) Code() int {
	return e.CodeData
}

func (e ErrorData) Status() string {
	return e.StatusData
}

func NewUnauthenticatedError(message string) Error {
	return &ErrorData{
		MessageData: message,
		CodeData:    401,
		StatusData:  "UNAUTHENTICATED",
	}
}

func NewNotFoundError(message string) Error {
	return &ErrorData{
		MessageData: message,
		CodeData:    404,
		StatusData:  "NOT_FOUND",
	}
}

func NewConflictError(message string) Error {
	return &ErrorData{
		MessageData: message,
		CodeData:    409,
		StatusData:  "CONFLICT",
	}
}

func NewInternalServerError(message string) Error {
	return &ErrorData{
		MessageData: message,
		CodeData:    500,
		StatusData:  "INTERNAL_SERVER_ERROR",
	}
}
