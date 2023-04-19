package collection

import (
	"math/rand"
	"testing"

	"gitea.tursom.cn/tursom/kvs/kv"
)

func Test_listNodeCodec(t *testing.T) {
	codec := ListNodeCodec(kv.Int32ToByteCodec)

	prev := uint32(rand.Int31())
	next := uint32(rand.Int31())
	value := rand.Int31()

	encode := codec.Encode(&ListNode[int32]{
		prev:  prev,
		next:  next,
		value: value,
	})
	decode := codec.Decode(encode)

	if decode.prev != prev || decode.next != next || decode.value != value {
		t.Fatal(decode)
	}
}
