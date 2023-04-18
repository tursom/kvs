package kv

import (
	"bytes"
	"io"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

// 不推荐使用

type (
	arrayCodec[V any] struct {
		lang.BaseObject
		codec Codec[io.Reader, V]
	}
)

func ArrayCodec[V any](codec Codec[io.Reader, V]) Codec[[]byte, []V] {
	return &arrayCodec[V]{codec: codec}
}

func (a *arrayCodec[V]) encode(v2 []V) []byte {
	var bs []byte
	for _, v := range v2 {
		encode := a.codec.encode(v)
		all, err := io.ReadAll(encode)
		if err != nil {
			panic(exceptions.Package(err))
		}

		bs = append(bs, all...)
	}

	return bs
}

func (a *arrayCodec[V]) decode(v1 []byte) []V {
	if len(v1) == 0 {
		return []V{}
	}

	reader := bytes.NewReader(v1)

	var values []V

	for func() bool {
		defer recover()

		v := a.codec.decode(reader)
		values = append(values, v)

		return reader.Len() > 0
	}() {
	}

	return values
}
