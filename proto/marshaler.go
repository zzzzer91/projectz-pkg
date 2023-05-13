package proto

import (
	"github.com/bytedance/sonic"
)

func MarshalJson(v *Json) []byte {
	d, _ := sonic.Marshal(v)
	v.reset()
	jsonPool.Put(v)
	return d
}

func MarshalProtobuf(v *Protobuf) []byte {
	buf := make([]byte, v.Size())
	v.FastWrite(buf)
	v.Reset()
	protobufPool.Put(v)
	return buf
}
