package collection

import (
	"bytes"
	"encoding/binary"

	"github.com/tursom/GoCollections/collections"
	"github.com/tursom/GoCollections/exceptions"
	"github.com/tursom/GoCollections/lang"

	"gitea.tursom.cn/tursom/kvs/kv"
)

type (
	List[T lang.Object] interface {
		collections.MutableList[T]
	}

	ListNode[T any] struct {
		lang.BaseObject
		prev, next uint32
		value      T
	}

	listNodeCodec[T any] struct {
		lang.BaseObject
		codec kv.Codec[[]byte, T]
	}

	listImpl[T lang.Object] struct {
		kvs        kv.Store[uint32, *ListNode[T]]
		head, tail uint32
	}

	listIterator[T any] struct {
		kvs kv.Store[uint32, *ListNode[T]]
	}
)

func ListNodeCodec[T any](codec kv.Codec[[]byte, T]) kv.Codec[[]byte, *ListNode[T]] {
	return &listNodeCodec[T]{
		codec: codec,
	}
}

func NewList[T lang.Object](kvs kv.Store[uint32, *ListNode[T]]) List[T] {
	headPointer, exception := kvs.Get(0)
	if exception != nil {
		panic(exception)
	}

	var head uint32
	if headPointer == nil {
		head = 1
	} else {
		head = headPointer.next
	}

	return &listImpl[T]{
		kvs:  kvs,
		head: head,
	}
}

func (l *listNodeCodec[T]) Encode(v2 *ListNode[T]) []byte {
	if v2 == nil {
		return nil
	}

	buffer := bytes.NewBuffer(nil)

	_ = binary.Write(buffer, binary.BigEndian, v2.prev)
	_ = binary.Write(buffer, binary.BigEndian, v2.next)
	buffer.Write(l.codec.Encode(v2.value))

	return buffer.Bytes()
}

func (l *listNodeCodec[T]) Decode(v1 []byte) *ListNode[T] {
	if len(v1) == 0 {
		return nil
	}

	return &ListNode[T]{
		prev:  binary.BigEndian.Uint32(v1),
		next:  binary.BigEndian.Uint32(v1[4:]),
		value: l.codec.Decode(v1[8:]),
	}
}

func (l *listImpl[T]) Size() int {
	size, err := collections.Size[T](l)
	if err != nil {
		panic(err)
	}
	return size
}

func (l *listImpl[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) Contains(element T) bool {
	return collections.Contains[T](l, element)
}

func (l *listImpl[T]) ContainsAll(c collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) Get(index int) (T, exceptions.Exception) {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) SubList(from, to int) collections.List[T] {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) Add(element T) bool {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) Remove(element T) exceptions.Exception {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) AddAll(c collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) RemoveAll(c collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) RetainAll(c collections.Collection[T]) bool {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) Set(index int, element T) exceptions.Exception {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) AddAtIndex(index int, element T) bool {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) RemoveAt(index int) exceptions.Exception {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) SubMutableList(from, to int) collections.MutableList[T] {
	//TODO implement me
	panic("implement me")
}

func (l *listImpl[T]) Iterator() collections.Iterator[T] {
	return l.MutableListIterator()
}

func (l *listImpl[T]) ListIterator() collections.ListIterator[T] {
	return l.MutableListIterator()
}

func (l *listImpl[T]) MutableIterator() collections.MutableIterator[T] {
	return l.MutableListIterator()
}

func (l *listImpl[T]) MutableListIterator() collections.MutableListIterator[T] {
	//TODO implement me
	panic("implement me")
}
