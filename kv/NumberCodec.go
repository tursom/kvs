package kv

import (
	"encoding/binary"
	"unsafe"

	"github.com/tursom/GoCollections/lang"
)

type (
	int8ToByteCodec struct {
		lang.BaseObject
	}
	int16ToByteCodec struct {
		lang.BaseObject
	}
	int32ToByteCodec struct {
		lang.BaseObject
	}
	int64ToByteCodec struct {
		lang.BaseObject
	}
	uint8ToByteCodec struct {
		lang.BaseObject
	}
	uint16ToByteCodec struct {
		lang.BaseObject
	}
	uint32ToByteCodec struct {
		lang.BaseObject
	}
	uint64ToByteCodec struct {
		lang.BaseObject
	}
	float32ToByteCodec struct {
		lang.BaseObject
	}
	float64ToByteCodec struct {
		lang.BaseObject
	}
	complex64ToByteCodec struct {
		lang.BaseObject
	}
	complex128ToByteCodec struct {
		lang.BaseObject
	}
)

var (
	Int8ToByteCodec       Codec[[]byte, int8]       = &int8ToByteCodec{}
	Int16ToByteCodec      Codec[[]byte, int16]      = &int16ToByteCodec{}
	Int32ToByteCodec      Codec[[]byte, int32]      = &int32ToByteCodec{}
	Int64ToByteCodec      Codec[[]byte, int64]      = &int64ToByteCodec{}
	Uint8ToByteCodec      Codec[[]byte, uint8]      = &uint8ToByteCodec{}
	Uint16ToByteCodec     Codec[[]byte, uint16]     = &uint16ToByteCodec{}
	Uint32ToByteCodec     Codec[[]byte, uint32]     = &uint32ToByteCodec{}
	Uint64ToByteCodec     Codec[[]byte, uint64]     = &uint64ToByteCodec{}
	Float32ToByteCodec    Codec[[]byte, float32]    = &float32ToByteCodec{}
	Float64ToByteCodec    Codec[[]byte, float64]    = &float64ToByteCodec{}
	Complex64ToByteCodec  Codec[[]byte, complex64]  = &complex64ToByteCodec{}
	Complex128ToByteCodec Codec[[]byte, complex128] = &complex128ToByteCodec{}
	ByteToInt8Codec                                 = InvertCodec[int8, []byte](&int8ToByteCodec{})
	ByteToInt16Codec                                = InvertCodec[int16, []byte](&int16ToByteCodec{})
	ByteToInt32Codec                                = InvertCodec[int32, []byte](&int32ToByteCodec{})
	ByteToInt64Codec                                = InvertCodec[int64, []byte](&int64ToByteCodec{})
	ByteToUint8Codec                                = InvertCodec[uint8, []byte](&uint8ToByteCodec{})
	ByteToUint16Codec                               = InvertCodec[uint16, []byte](&uint16ToByteCodec{})
	ByteToUint32Codec                               = InvertCodec[uint32, []byte](&uint32ToByteCodec{})
	ByteToUint64Codec                               = InvertCodec[uint64, []byte](&uint64ToByteCodec{})
	ByteToFloat32Codec                              = InvertCodec[float32, []byte](&float32ToByteCodec{})
	ByteToFloat64Codec                              = InvertCodec[float64, []byte](&float64ToByteCodec{})
	ByteToComplex64Codec                            = InvertCodec[complex64, []byte](&complex64ToByteCodec{})
	ByteToComplex128Codec                           = InvertCodec[complex128, []byte](&complex128ToByteCodec{})
)

func (u *int8ToByteCodec) encode(v2 int8) []byte {
	return []byte{byte(v2)}
}

func (u *int8ToByteCodec) decode(v1 []byte) int8 {
	if len(v1) == 0 {
		return 0
	}

	return int8(v1[0])
}

func (u *int16ToByteCodec) encode(v2 int16) []byte {
	return binary.BigEndian.AppendUint16(nil, uint16(v2))
}

func (u *int16ToByteCodec) decode(v1 []byte) int16 {
	if len(v1) == 0 {
		return 0
	}

	return int16(binary.BigEndian.Uint16(v1))
}

func (u *int32ToByteCodec) encode(v2 int32) []byte {
	return binary.BigEndian.AppendUint32(nil, uint32(v2))
}

func (u *int32ToByteCodec) decode(v1 []byte) int32 {
	if len(v1) == 0 {
		return 0
	}

	return int32(binary.BigEndian.Uint32(v1))
}

func (u *int64ToByteCodec) encode(v2 int64) []byte {
	return binary.BigEndian.AppendUint64(nil, uint64(v2))
}

func (u *int64ToByteCodec) decode(v1 []byte) int64 {
	if len(v1) == 0 {
		return 0
	}

	return int64(binary.BigEndian.Uint64(v1))
}

func (u *uint8ToByteCodec) encode(v2 uint8) []byte {
	return []byte{v2}
}

func (u *uint8ToByteCodec) decode(v1 []byte) uint8 {
	if len(v1) == 0 {
		return 0
	}

	return v1[0]
}

func (u *uint16ToByteCodec) encode(v2 uint16) []byte {
	return binary.BigEndian.AppendUint16(nil, v2)
}

func (u *uint16ToByteCodec) decode(v1 []byte) uint16 {
	if len(v1) == 0 {
		return 0
	}

	return binary.BigEndian.Uint16(v1)
}

func (u *uint32ToByteCodec) encode(v2 uint32) []byte {
	return binary.BigEndian.AppendUint32(nil, v2)
}

func (u *uint32ToByteCodec) decode(v1 []byte) uint32 {
	if len(v1) == 0 {
		return 0
	}

	return binary.BigEndian.Uint32(v1)
}

func (u *uint64ToByteCodec) encode(v2 uint64) []byte {
	return binary.BigEndian.AppendUint64(nil, v2)
}

func (u *uint64ToByteCodec) decode(v1 []byte) uint64 {
	if len(v1) == 0 {
		return 0
	}

	return binary.BigEndian.Uint64(v1)
}

func (u *float32ToByteCodec) encode(v2 float32) []byte {
	return binary.BigEndian.AppendUint32(nil, *(*uint32)(unsafe.Pointer(&v2)))
}

func (u *float32ToByteCodec) decode(v1 []byte) float32 {
	if len(v1) == 0 {
		return 0
	}

	u2 := binary.BigEndian.Uint32(v1)
	return *(*float32)(unsafe.Pointer(&u2))
}

func (u *float64ToByteCodec) encode(v2 float64) []byte {
	return binary.BigEndian.AppendUint64(nil, *(*uint64)(unsafe.Pointer(&v2)))
}

func (u *float64ToByteCodec) decode(v1 []byte) float64 {
	if len(v1) == 0 {
		return 0
	}

	u2 := binary.BigEndian.Uint64(v1)
	return *(*float64)(unsafe.Pointer(&u2))
}

func (u *complex64ToByteCodec) encode(v2 complex64) []byte {
	return binary.BigEndian.AppendUint64(nil, *(*uint64)(unsafe.Pointer(&v2)))
}

func (u *complex64ToByteCodec) decode(v1 []byte) complex64 {
	if len(v1) == 0 {
		return 0
	}

	u2 := binary.BigEndian.Uint64(v1)
	return *(*complex64)(unsafe.Pointer(&u2))
}

func (u *complex128ToByteCodec) encode(v2 complex128) []byte {
	r := real(v2)
	i := imag(v2)

	bytes := make([]byte, 16)

	binary.BigEndian.PutUint64(bytes[0:], *(*uint64)(unsafe.Pointer(&r)))
	binary.BigEndian.PutUint64(bytes[8:], *(*uint64)(unsafe.Pointer(&i)))

	return bytes
}

func (u *complex128ToByteCodec) decode(v1 []byte) complex128 {
	if len(v1) == 0 {
		return 0
	}

	r := binary.BigEndian.Uint64(v1[0:])
	i := binary.BigEndian.Uint64(v1[8:])

	return complex(
		*(*float64)(unsafe.Pointer(&r)),
		*(*float64)(unsafe.Pointer(&i)),
	)
}
