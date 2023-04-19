package kv

import (
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"
)

type (
	Store[K, V any] interface {
		lang.Object
		Put(key K, value V) exceptions.Exception
		Get(key K) (V, exceptions.Exception)
		Delete(key K) exceptions.Exception
	}
)
