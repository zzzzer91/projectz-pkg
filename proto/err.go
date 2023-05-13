package proto

import (
	"fmt"
)

type CodeErr struct {
	code uint32
	msg  string
	err  error
}

func NewCodeErr(code uint32, msg string, err error) *CodeErr {
	return &CodeErr{
		code: code,
		msg:  msg,
		err:  err,
	}
}

func NewClientDefaultCodeErr(msg string, err error) *CodeErr {
	return NewCodeErr(ErrorCodeClientDefault, msg, err)
}

func NewServerDefaultCodeErr(msg string, err error) *CodeErr {
	return NewCodeErr(ErrorCodeServerDefault, msg, err)
}

func (e *CodeErr) Error() string {
	return fmt.Sprintf("Code: [%d], Msg: [%s], Err: [%v]", e.code, e.msg, e.err)
}

func (e *CodeErr) Code() uint32 {
	return e.code
}

func (e *CodeErr) Msg() string {
	return e.msg
}

func (e *CodeErr) Cause() error {
	return e.err
}
