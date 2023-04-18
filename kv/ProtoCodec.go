package kv

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
	"google.golang.org/protobuf/proto"
)

type (
	protoToByteCodec[V proto.Message] struct {
		lang.BaseObject
		emptyMessage func() V
	}
)

func ProtoCodec[V proto.Message](emptyMessage func() V) Codec[[]byte, V] {
	return &protoToByteCodec[V]{
		emptyMessage: emptyMessage,
	}
}

func ProtoDeCodec[V proto.Message](emptyMessage func() V) Codec[V, []byte] {
	return InvertCodec(ProtoCodec(emptyMessage))
}

func (p *protoToByteCodec[V]) encode(v2 V) []byte {
	bytes, err := proto.Marshal(v2)
	if err == nil {
		panic(exceptions.Package(err))
	}

	return bytes
}

func (p *protoToByteCodec[V]) decode(v1 []byte) V {
	message := p.emptyMessage()
	if err := proto.Unmarshal(v1, message); err != nil {
		panic(exceptions.Package(err))
	}
	return message
}
