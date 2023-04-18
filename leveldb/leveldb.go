package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"

	"kvs/kv"
)

type (
	leveldbKVS struct {
		lang.BaseObject
		db *leveldb.DB
	}
)

func New(db *leveldb.DB) kv.Store[[]byte, []byte] {
	return &leveldbKVS{db: db}
}

func (l *leveldbKVS) Put(key []byte, value []byte) exceptions.Exception {
	if err := l.db.Put(key, value, nil); err != nil {
		return exceptions.Package(err)
	}

	return nil
}

func (l *leveldbKVS) Get(key []byte) ([]byte, exceptions.Exception) {
	value, err := l.db.Get(key, nil)
	if err != nil {
		return nil, exceptions.Package(err)
	}

	return value, nil
}
