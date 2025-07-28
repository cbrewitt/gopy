package gopy

import (
	"encoding/binary"
	"math"
	"unsafe"
)

var isLittleEndian bool

func init() {
	var i uint16 = 1
	isLittleEndian = *(*byte)(unsafe.Pointer(&i)) == 1
}

func int32ToBytes1D(src []int32, dest []byte) {
	if len(src) == 0 {
		return
	}
	if isLittleEndian {
		copy(dest, unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4))
	} else {
		i := 0
		for _, val := range src {
			binary.LittleEndian.PutUint32(dest[i:i+4], uint32(val))
			i += 4
		}
	}
}

func int32ToBytes2D(s [][]int32, dest []byte) {
	if len(s) == 0 || len(s[0]) == 1 {
		return
	}
	l := len(s[0]) * 4
	i := 0
	for _, row := range s {
		int32ToBytes1D(row, dest[i:i+l])
		i += l
	}
}

func int32ToBytes3D(s [][][]int32, dest []byte) {
	if len(s) == 0 || len(s[0]) == 0 || len(s[0][0]) == 0 {
		return
	}
	l := len(s[0]) * len(s[0][0]) * 4
	i := 0
	for _, dim := range s {
		int32ToBytes2D(dim, dest[i:i+l])
		i += l
	}
}

func int16ToBytes1D(src []int16, dest []byte) {
	if len(src) == 0 {
		return
	}
	if isLittleEndian {
		copy(dest, unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*2))
	} else {
		i := 0
		for _, val := range src {
			binary.LittleEndian.PutUint16(dest[i:i+2], uint16(val))
			i += 2
		}
	}
}

func int16ToBytes2D(s [][]int16, dest []byte) {
	if len(s) == 0 || len(s[0]) == 1 {
		return
	}
	l := len(s[0]) * 2
	i := 0
	for _, row := range s {
		int16ToBytes1D(row, dest[i:i+l])
		i += l
	}
}

func int16ToBytes3D(s [][][]int16, dest []byte) {
	if len(s) == 0 || len(s[0]) == 0 || len(s[0][0]) == 0 {
		return
	}
	l := len(s[0]) * len(s[0][0]) * 2
	i := 0
	for _, dim := range s {
		int16ToBytes2D(dim, dest[i:i+l])
		i += l
	}
}

// Int64 versions
func int64ToBytes1D(src []int64, dest []byte) {
	if len(src) == 0 {
		return
	}
	if isLittleEndian {
		copy(dest, unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*8))
	} else {
		i := 0
		for _, val := range src {
			binary.LittleEndian.PutUint64(dest[i:i+8], uint64(val))
			i += 8
		}
	}
}

func int64ToBytes2D(s [][]int64, dest []byte) {
	if len(s) == 0 || len(s[0]) == 1 {
		return
	}
	l := len(s[0]) * 8
	i := 0
	for _, row := range s {
		int64ToBytes1D(row, dest[i:i+l])
		i += l
	}
}

func int64ToBytes3D(s [][][]int64, dest []byte) {
	if len(s) == 0 || len(s[0]) == 0 || len(s[0][0]) == 0 {
		return
	}
	l := len(s[0]) * len(s[0][0]) * 8
	i := 0
	for _, dim := range s {
		int64ToBytes2D(dim, dest[i:i+l])
		i += l
	}
}

// Float32 versions
func float32ToBytes1D(src []float32, dest []byte) {
	if len(src) == 0 {
		return
	}
	if isLittleEndian {
		copy(dest, unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*4))
	} else {
		i := 0
		for _, val := range src {
			binary.LittleEndian.PutUint32(dest[i:i+4], math.Float32bits(val))
			i += 4
		}
	}
}

func float32ToBytes2D(s [][]float32, dest []byte) {
	if len(s) == 0 || len(s[0]) == 1 {
		return
	}
	l := len(s[0]) * 4
	i := 0
	for _, row := range s {
		float32ToBytes1D(row, dest[i:i+l])
		i += l
	}
}

func float32ToBytes3D(s [][][]float32, dest []byte) {
	if len(s) == 0 || len(s[0]) == 0 || len(s[0][0]) == 0 {
		return
	}
	l := len(s[0]) * len(s[0][0]) * 4
	i := 0
	for _, dim := range s {
		float32ToBytes2D(dim, dest[i:i+l])
		i += l
	}
}

// Float64 versions
func float64ToBytes1D(src []float64, dest []byte) {
	if len(src) == 0 {
		return
	}
	if isLittleEndian {
		copy(dest, unsafe.Slice((*byte)(unsafe.Pointer(&src[0])), len(src)*8))
	} else {
		i := 0
		for _, val := range src {
			binary.LittleEndian.PutUint64(dest[i:i+8], math.Float64bits(val))
			i += 8
		}
	}
}

func float64ToBytes2D(s [][]float64, dest []byte) {
	if len(s) == 0 || len(s[0]) == 1 {
		return
	}
	l := len(s[0]) * 8
	i := 0
	for _, row := range s {
		float64ToBytes1D(row, dest[i:i+l])
		i += l
	}
}

func float64ToBytes3D(s [][][]float64, dest []byte) {
	if len(s) == 0 || len(s[0]) == 0 || len(s[0][0]) == 0 {
		return
	}
	l := len(s[0]) * len(s[0][0]) * 8
	i := 0
	for _, dim := range s {
		float64ToBytes2D(dim, dest[i:i+l])
		i += l
	}
}

func bytesToInt161D(src []byte, result []int16) {
	if len(result) == 0 {
		return
	}
	if isLittleEndian {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&result[0])), len(result)*2), src[:len(result)*2])
	} else {
		for i := 0; i < len(result); i++ {
			result[i] = int16(binary.LittleEndian.Uint16(src[i*2 : i*2+2]))
		}
	}
}

func bytesToInt321D(src []byte, result []int32) {
	if len(result) == 0 {
		return
	}
	if isLittleEndian {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&result[0])), len(result)*4), src[:len(result)*4])
	} else {
		for i := 0; i < len(result); i++ {
			result[i] = int32(binary.LittleEndian.Uint32(src[i*4 : i*4+4]))
		}
	}
}

func bytesToInt641D(src []byte, result []int64) {
	if len(result) == 0 {
		return
	}
	if isLittleEndian {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&result[0])), len(result)*8), src[:len(result)*8])
	} else {
		for i := 0; i < len(result); i++ {
			result[i] = int64(binary.LittleEndian.Uint64(src[i*8 : i*8+8]))
		}
	}
}

func bytesToFloat321D(src []byte, result []float32) {
	if len(result) == 0 {
		return
	}
	if isLittleEndian {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&result[0])), len(result)*4), src[:len(result)*4])
	} else {
		for i := 0; i < len(result); i++ {
			result[i] = math.Float32frombits(binary.LittleEndian.Uint32(src[i*4 : i*4+4]))
		}
	}
}

func bytesToFloat641D(src []byte, result []float64) {
	if len(result) == 0 {
		return
	}
	if isLittleEndian {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(&result[0])), len(result)*8), src[:len(result)*8])
	} else {
		for i := 0; i < len(result); i++ {
			result[i] = math.Float64frombits(binary.LittleEndian.Uint64(src[i*8 : i*8+8]))
		}
	}
}

// ----  UNSAFE VERSIONS WIP ----

func int32ToBytes1DUnsafe(s []int32) []byte {
	if len(s) == 0 {
		return []byte{}
	}
	sliceHeader := unsafe.SliceData(s)
	return *(*[]byte)(unsafe.Pointer(sliceHeader))
}
