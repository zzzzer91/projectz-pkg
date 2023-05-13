package proto

import (
	context "context"
	"errors"

	"github.com/zzzzer91/zlog"
)

func NewProtobuf(code uint32, msg string, data []byte) *Protobuf {
	p := protobufPool.Get().(*Protobuf)
	p.Code = code
	p.Msg = msg
	p.Data = data
	return p
}

func NewProtobufSuccess(msg string, data []byte) *Protobuf {
	return NewProtobuf(SuccessCodeDefault, msg, data)
}

func NewProtobufError(code uint32, msg string) *Protobuf {
	return NewProtobuf(code, msg, nil)
}

func NewDefaultProtobufError() *Protobuf {
	return NewProtobufError(ErrorCodeServerDefault, ErrorMsgServerDefault)
}

func ToProtobufError(ctx context.Context, err error) *Protobuf {
	var resp *Protobuf
	codeErr := new(CodeErr)
	if errors.As(err, &codeErr) {
		code, msg, cause := codeErr.Code(), codeErr.Msg(), codeErr.Cause()
		resp = NewProtobufError(code, msg)
		// 只打印 ErrorCodeServerDefault 级别的错误
		// 没塞 cause，服务端就不打印 err，只返回给客户端错误
		if code >= ErrorCodeServerDefault && cause != nil {
			zlog.Ctx(ctx).WithError(cause).Errorf("code error, code: %d, msg: %s", code, msg)
		}
	} else {
		resp = NewDefaultProtobufError()
		zlog.Ctx(ctx).WithError(err).Error("internal error")
	}
	return resp
}
