package responses

type success struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponsesSuccess(msg string, data interface{}) success {
	return success{
		Code:    200,
		Status:  true,
		Message: msg,
		Data:    data,
	}
}
