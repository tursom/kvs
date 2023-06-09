package sqlite

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tursom/GoCollections/exceptions"

	"gitea.tursom.cn/tursom/kvs/kv"
)

func Test_sqliteKVS(t *testing.T) {
	db := exceptions.Exec2r1(sql.Open, "sqlite3", ":memory:")
	s, exception := New(db, "kv")
	if exception != nil {
		t.Fatal(exception)
	}

	skvs := kv.CodecStore(s, kv.StringToByteCodec, kv.StringToByteCodec)

	if err := skvs.Put("hello", "world!"); err != nil {
		t.Fatal(err)
	}

	value, exception := skvs.Get("hello")
	if exception != nil || value != "world!" {
		t.Fatal(value, exception)
	}

	if exception = skvs.Delete("hello"); exception != nil {
		t.Fatal(value, exception)
	}

	value, exception = skvs.Get("hello")
	if exception != nil || value != "" {
		t.Fatal(value, exception)
	}
}
