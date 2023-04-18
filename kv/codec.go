package kv

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	Codec[V1, V2 any] interface {
		lang.Object
		Encode(v2 V2) V1
		Decode(v1 V1) V2
	}

	codecStore[K1, K2, V1, V2 any] struct {
		lang.BaseObject
		kvs    Store[K1, V1]
		kCodec Codec[K1, K2]
		vCodec Codec[V1, V2]
	}

	kCodecStore[K1, K2, V any] struct {
		lang.BaseObject
		kvs   Store[K1, V]
		codec Codec[K1, K2]
	}

	vCodecStore[K, V1, V2 any] struct {
		lang.BaseObject
		kvs   Store[K, V1]
		codec Codec[V1, V2]
	}

	invertCodec[V1, V2 any] struct {
		lang.BaseObject
		codec Codec[V2, V1]
	}
)

func InvertCodec[V1, V2 any](codec Codec[V2, V1]) Codec[V1, V2] {
	return &invertCodec[V1, V2]{codec: codec}
}

func CodecStore[K1, K2, V1, V2 any](
	kvs Store[K1, V1],
	kCodec Codec[K1, K2],
	vCodec Codec[V1, V2],
) Store[K2, V2] {
	return &codecStore[K1, K2, V1, V2]{
		kvs:    kvs,
		kCodec: kCodec,
		vCodec: vCodec,
	}
}

func KCodecStore[K1, K2, V any](
	kvs Store[K1, V],
	codec Codec[K1, K2],
) Store[K2, V] {
	return &kCodecStore[K1, K2, V]{
		kvs:   kvs,
		codec: codec,
	}
}

func VCodecStore[K, V1, V2 any](
	kvs Store[K, V1],
	codec Codec[V1, V2],
) Store[K, V2] {
	return &vCodecStore[K, V1, V2]{
		kvs:   kvs,
		codec: codec,
	}
}

func (c *codecStore[K1, K2, V1, V2]) Put(key K2, value V2) exceptions.Exception {
	return c.kvs.Put(c.kCodec.Encode(key), c.vCodec.Encode(value))
}

func (c *codecStore[K1, K2, V1, V2]) Get(key K2) (V2, exceptions.Exception) {
	value, exception := c.kvs.Get(c.kCodec.Encode(key))
	if exception != nil {
		return lang.Nil[V2](), exception
	}

	return c.vCodec.Decode(value), nil
}

func (c *kCodecStore[K1, K2, V]) Put(key K2, value V) exceptions.Exception {
	return c.kvs.Put(c.codec.Encode(key), value)
}

func (c *kCodecStore[K1, K2, V]) Get(key K2) (V, exceptions.Exception) {
	return c.kvs.Get(c.codec.Encode(key))
}

func (c *vCodecStore[K, V1, V2]) Put(key K, value V2) exceptions.Exception {
	return c.kvs.Put(key, c.codec.Encode(value))
}

func (c *vCodecStore[K, V1, V2]) Get(key K) (V2, exceptions.Exception) {
	get, exception := c.kvs.Get(key)
	if exception != nil {
		return lang.Nil[V2](), exception
	}

	return c.codec.Decode(get), nil
}

func (i *invertCodec[V1, V2]) Encode(v2 V2) V1 {
	return i.codec.Decode(v2)
}

func (i *invertCodec[V1, V2]) Decode(v1 V1) V2 {
	return i.codec.Encode(v1)
}
