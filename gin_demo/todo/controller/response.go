package controller

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(code int, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func SuccessResp(msg string, data interface{}) Response {
	return NewResp(100000, msg, data)
}

func SuccessRespWithData(data interface{}) Response {
	return SuccessResp("success", data)
}

func ArgsErrResp(msg string) Response {
	return NewResp(100400, msg, nil)
}
