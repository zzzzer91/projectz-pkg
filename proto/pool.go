package proto

import "sync"

var (
	jsonPool = sync.Pool{
		New: func() interface{} {
			return new(Json)
		},
	}

	protobufPool = sync.Pool{
		New: func() interface{} {
			return new(Protobuf)
		},
	}
)
