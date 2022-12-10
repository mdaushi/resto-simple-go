package responses

type notFound struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ResponsesNotFound(msg string) notFound {
	return notFound{
		Code:    404,
		Status:  false,
		Message: msg,
	}
}
