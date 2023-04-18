package kv

import "github.com/tursom/GoCollections/lang"

type (
	stringToByteCodec struct {
		lang.BaseObject
	}
)

var (
	StringToByteCodec Codec[[]byte, string] = &stringToByteCodec{}
	ByteToStringCodec                       = InvertCodec[string, []byte](&stringToByteCodec{})
)

func (s *stringToByteCodec) encode(v2 string) []byte {
	return []byte(v2)
}

func (s *stringToByteCodec) decode(v1 []byte) string {
	if len(v1) == 0 {
		return ""
	}

	return string(v1)
}
