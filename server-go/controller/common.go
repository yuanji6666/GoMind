package controller

import "github.com/yuanji6666/gopherAI/common/code"

type Response struct {
	StatusCode	code.Code	`json:"status_code"`
	StatusMsg	string		`json:"status_msg,omitempty"`
}

func (r *Response) CodeOf(code code.Code) Response {
	//自动处理空指针
	if r == nil {
		r = new(Response)
	}
	r.StatusCode = code
	r.StatusMsg = code.Msg()
	//两种使用方式，原始值上直接修改同时返回副本
	return *r
}

func (r *Response) Success(){
	r.CodeOf(code.CodeSuccess)
}