package kv

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	prefixCodec struct {
		lang.BaseObject
		prefix string
	}
)

func PrefixCodec(prefix string) Codec[string, string] {
	return &prefixCodec{prefix: prefix}
}

func (p *prefixCodec) Encode(v2 string) string {
	return p.prefix + v2
}

func (p *prefixCodec) Decode(v1 string) string {
	panic(exceptions.NewIllegalAccessException("unsupported method", nil))
}
