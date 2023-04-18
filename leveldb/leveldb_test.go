package leveldb

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"

	"gitea.tursom.cn/tursom/kvs/kv"
)

func Test_leveldbKVS(t *testing.T) {
	db, err := leveldb.OpenFile("test", nil)
	if err != nil {
		t.Fatal(err)
	}

	s := kv.CodecStore(New(db), kv.StringToByteCodec, kv.StringToByteCodec)

	if err := s.Put("hello", "world!"); err != nil {
		t.Fatal(err)
	}

	value, exception := s.Get("hello")
	if exception != nil || value != "world!" {
		t.Fatal(value, exception)
	}
}
