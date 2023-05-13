package response

import (
	"errors"
	"net/http"

	"github.com/cloudwego/fastpb"
	"github.com/gin-gonic/gin"
	"github.com/zzzzer91/projectz-pkg/proto"
)

func ProtobufError(c *gin.Context, err error) {
	var resp *proto.Protobuf
	codeErr := new(proto.CodeErr)
	if errors.As(err, &codeErr) {
		code, msg, cause := codeErr.Code(), codeErr.Msg(), codeErr.Cause()
		resp = proto.NewProtobufError(code, msg)
		if code >= proto.ErrorCodeServerDefault && cause != nil {
			_ = c.Error(cause)
		}
	} else {
		resp = proto.NewDefaultProtobufError()
		_ = c.Error(err)
	}
	c.Data(http.StatusOK, contentTypeProtobuf, proto.MarshalProtobuf(resp))
}

func ProtobufSuccess(c *gin.Context, msg string, data fastpb.Writer) {
	var buf []byte
	if data != nil {
		buf = make([]byte, data.Size())
		data.FastWrite(buf)
	}
	RawProtobufSuccess(c, msg, buf)
}

func RawProtobufSuccess(c *gin.Context, msg string, data []byte) {
	resp := proto.NewProtobufSuccess(msg, data)
	c.Data(http.StatusOK, contentTypeProtobuf, proto.MarshalProtobuf(resp))
}
