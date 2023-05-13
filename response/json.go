package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzzzer91/projectz-pkg/proto"
)

func JsonError(c *gin.Context, err error) {
	var resp *proto.Json
	codeErr := new(proto.CodeErr)
	if errors.As(err, &codeErr) {
		code, msg, cause := codeErr.Code(), codeErr.Msg(), codeErr.Cause()
		resp = proto.NewJsonError(code, msg)
		if code >= proto.ErrorCodeServerDefault && cause != nil {
			_ = c.Error(cause)
		}
	} else {
		resp = proto.NewDefaultJsonError()
		_ = c.Error(err)
	}
	c.Data(http.StatusOK, contentTypeJson, proto.MarshalJson(resp))
}

func JsonSuccess(c *gin.Context, msg string, data interface{}) {
	resp := proto.NewJsonSuccess(msg, data)
	c.Data(http.StatusOK, contentTypeJson, proto.MarshalJson(resp))
}
