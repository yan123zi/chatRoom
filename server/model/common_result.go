package model

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

type RespBuilder struct {
	Data map[string]interface{}
}

func NewEmptyRespBuild() *RespBuilder {
	return &RespBuilder{Data: map[string]interface{}{}}
}

func (r *RespBuilder) Put(key string,val interface{}) *RespBuilder {
	r.Data[key]=val
	return r
}

func (r *RespBuilder) JsonResult() *JsonResult {
	return JsonData(r.Data)
}