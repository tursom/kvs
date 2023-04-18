package kv

import (
	"fmt"
	"testing"
)

func Test_mapKvs(t *testing.T) {
	kvs := CodecStore(
		MapKvs(),
		ByteToStringCodec,
		ArrayCodec(ChainCodec(StringToByteCodec, LengthFieldCodec)),
	)
	_ = kvs.Put([]byte{1}, []string{"hello", "world!"})
	fmt.Println(kvs.Get([]byte{1}))
}
