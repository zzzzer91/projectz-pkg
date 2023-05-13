package proto

import (
	context "context"
	"errors"

	"github.com/zzzzer91/zlog"
)

type Json struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (p *Json) reset() {
	p.Code = SuccessCodeDefault
	p.Msg = ""
	p.Data = nil
}

func NewJson(code uint32, msg string, data interface{}) *Json {
	p := jsonPool.Get().(*Json)
	p.Code = code
	p.Msg = msg
	p.Data = data
	return p
}

func NewJsonSuccess(msg string, data interface{}) *Json {
	return NewJson(SuccessCodeDefault, msg, data)
}

func NewJsonError(code uint32, msg string) *Json {
	return NewJson(code, msg, nil)
}

func NewDefaultJsonError() *Json {
	return NewJsonError(ErrorCodeServerDefault, ErrorMsgServerDefault)
}

func ToJsonError(ctx context.Context, err error) *Json {
	var resp *Json
	codeErr := new(CodeErr)
	if errors.As(err, &codeErr) {
		code, msg, cause := codeErr.Code(), codeErr.Msg(), codeErr.Cause()
		resp = NewJsonError(code, msg)
		// 只打印 ErrorCodeServerDefault 级别的错误
		// 没塞 cause，服务端就不打印 err，只返回给客户端错误
		if code >= ErrorCodeServerDefault && cause != nil {
			zlog.Ctx(ctx).WithError(cause).Errorf("code error, code: %d, msg: %s", code, msg)
		}
	} else {
		resp = NewDefaultJsonError()
		zlog.Ctx(ctx).WithError(err).Error("internal error")
	}
	return resp
}
