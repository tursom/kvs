package kv

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

// 不推荐使用

type (
	readerCodec[V any] struct {
		lang.BaseObject
		codec Codec[[]byte, V]
	}
	fixedLengthCodec struct {
		lang.BaseObject
		frameLength uint32
	}
	lengthFieldCodec struct {
		lang.BaseObject
	}
)

var (
	LengthFieldCodec Codec[io.Reader, []byte] = &lengthFieldCodec{}
)

func ReaderCodec[V any](codec Codec[[]byte, V]) Codec[io.Reader, V] {
	return &readerCodec[V]{
		codec: codec,
	}
}

func FixedLengthCodec(frameLength uint32) Codec[io.Reader, []byte] {
	return &fixedLengthCodec{frameLength: frameLength}
}

func (r *readerCodec[V]) encode(v2 V) io.Reader {
	return bytes.NewReader(r.codec.Encode(v2))
}

func (r *readerCodec[V]) decode(v1 io.Reader) V {
	all, err := io.ReadAll(v1)
	if err != nil {
		panic(exceptions.Package(err))
	}

	return r.codec.Decode(all)
}

func (f *fixedLengthCodec) encode(v2 []byte) io.Reader {
	return bytes.NewReader(v2)
}

func (f *fixedLengthCodec) decode(v1 io.Reader) []byte {
	bs := make([]byte, f.frameLength)
	n, err := v1.Read(bs)
	if err != nil {
		panic(exceptions.Package(err))
	}

	return bs[0:n]
}

func (l *lengthFieldCodec) encode(v2 []byte) io.Reader {
	buffer := bytes.NewBuffer(nil)

	_ = binary.Write(buffer, binary.BigEndian, uint32(len(v2)))
	buffer.Write(v2)

	return buffer
}

func (l *lengthFieldCodec) decode(v1 io.Reader) []byte {
	var length uint32
	if err := binary.Read(v1, binary.BigEndian, &length); err != nil {
		panic(exceptions.Package(err))
	}

	bs := make([]byte, length)
	n, err := v1.Read(bs)
	if err != nil {
		panic(exceptions.Package(err))
	}

	return bs[0:n]
}
