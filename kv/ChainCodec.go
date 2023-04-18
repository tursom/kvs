package kv

import "github.com/tursom/GoCollections/lang"

type (
	chainCodec[V1, V2, V3 any] struct {
		lang.BaseObject
		codec1 Codec[V2, V3]
		codec2 Codec[V1, V2]
	}
)

func ChainCodec[V1, V2, V3 any](codec1 Codec[V2, V3], codec2 Codec[V1, V2]) Codec[V1, V3] {
	return &chainCodec[V1, V2, V3]{
		codec1: codec1,
		codec2: codec2,
	}
}

func (c *chainCodec[V1, V2, V3]) encode(v2 V3) V1 {
	return c.codec2.encode(c.codec1.encode(v2))
}

func (c *chainCodec[V1, V2, V3]) decode(v1 V1) V3 {
	return c.codec1.decode(c.codec2.decode(v1))
}
