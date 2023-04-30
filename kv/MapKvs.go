package kv

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	mapKvs[K comparable, V any] struct {
		lang.BaseObject
		m map[K]V
	}
)

func MapKvs() Store[string, []byte] {
	return MapKvs1[string, []byte]()
}

func MapKvs1[K comparable, V any]() Store[K, V] {
	return &mapKvs[K, V]{
		m: make(map[K]V),
	}
}

func (m *mapKvs[K, V]) Put(key K, value V) exceptions.Exception {
	m.m[key] = value
	return nil
}

func (m *mapKvs[K, V]) Get(key K) (V, exceptions.Exception) {
	return m.m[key], nil
}

func (m *mapKvs[K, V]) Delete(key K) exceptions.Exception {
	delete(m.m, key)
	return nil
}
