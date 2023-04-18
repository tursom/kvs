package kv

import "github.com/tursom/GoCollections/lang"

type (
	boolCodec struct {
		lang.BaseObject
	}
)

var (
	TrueBytes  = []byte{1}
	FalseBytes = []byte{0}

	BoolToByteCodec Codec[[]byte, bool] = &boolCodec{}
)

func (b *boolCodec) encode(v2 bool) []byte {
	if v2 {
		return TrueBytes
	} else {
		return FalseBytes
	}
}

func (b *boolCodec) decode(v1 []byte) bool {
	if len(v1) == 0 {
		return false
	}

	return v1[0] != 0
}
