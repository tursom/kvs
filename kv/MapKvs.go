package kv

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	mapKvs struct {
		lang.BaseObject
		m map[string][]byte
	}
)

func MapKvs() Store[string, []byte] {
	return &mapKvs{
		m: make(map[string][]byte),
	}
}

func (m *mapKvs) Put(key string, value []byte) exceptions.Exception {
	m.m[key] = value
	return nil
}

func (m *mapKvs) Get(key string) ([]byte, exceptions.Exception) {
	return m.m[key], nil
}

func (m *mapKvs) Delete(key string) exceptions.Exception {
	delete(m.m, key)
	return nil
}
