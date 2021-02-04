package common

type JsonResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func NewJsonResult(errNo int, errMsg string, data interface{}, status bool) *JsonResult {
	return &JsonResult{
		ErrorCode: errNo,
		Message:   errMsg,
		Data:      data,
		Success:   status,
	}
}

func JsonData(data interface{}) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

func JsonError(codeErr *CodeErrMsg) *JsonResult {
	return &JsonResult{
		ErrorCode: codeErr.Code,
		Message:   codeErr.Msg,
		Data:      codeErr.Data,
		Success:   false,
	}
}

func JsonErrorMsg(message string) *JsonResult {
	return &JsonResult{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

type RespBuilder struct {
	Data map[string]interface{}
}

func NewEmptyRespBuild() *RespBuilder {
	return &RespBuilder{Data: map[string]interface{}{}}
}

func (r *RespBuilder) Put(key string, val interface{}) *RespBuilder {
	r.Data[key] = val
	return r
}

func (r *RespBuilder) JsonResult() *JsonResult {
	return JsonData(r.Data)
}

type CodeErrMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewError(code int,msg string) *CodeErrMsg {
	return &CodeErrMsg{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func NewErrorMsg(msg string) *CodeErrMsg {
	return &CodeErrMsg{
		Code: 0,
		Msg:  msg,
		Data: nil,
	}
}

func (e *CodeErrMsg) Error() string {
	return e.Msg
}