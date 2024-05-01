package MemEngine

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func float32ToBytes(Value float32) []byte {
	bits := math.Float32bits(Value)
	Bytes := make([]byte, 4)
	for i := range Bytes {
		Bytes[i] = byte(bits >> uint(i*0x8))
	}
	return Bytes
}

func float64ToBytes(Value float64) []byte {
	bits := math.Float64bits(Value)
	Bytes := make([]byte, 8)
	for i := range Bytes {
		Bytes[i] = byte(bits >> uint(i*0x8))
	}
	return Bytes
}

func int32ToBytes(Value int32) []byte {
	Bytes := make([]byte, 4)
	for i := range Bytes {
		Bytes[i] = byte(Value >> uint(i*0x8))
	}
	return Bytes
}

func int64ToBytes(Value int64) []byte {
	Bytes := make([]byte, 8)
	for i := range Bytes {
		Bytes[i] = byte(Value >> uint(i*0x8))
	}
	return Bytes
}

func uint32ToBytes(Value uint32) []byte {
	Bytes := make([]byte, 4)
	for i := range Bytes {
		Bytes[i] = byte(Value >> uint(i*0x8))
	}
	return Bytes
}
func uint64ToBytes(Value uint64) []byte {
	Bytes := make([]byte, 8)
	for i := range Bytes {
		Bytes[i] = byte(Value >> uint(i*0x8))
	}
	return Bytes
}

func bytesToFloat32(Value []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(Value))
}

func bytesToFloat64(Value []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(Value))
}
func bytesToUint32(Value []byte) uint32 {
	return binary.LittleEndian.Uint32(Value)
}
func bytesToUint64(Value []byte) uint64 {
	return binary.LittleEndian.Uint64(Value)
}
func bytesToInt32(Value []byte) int32 {
	return int32(binary.LittleEndian.Uint32(Value))
}
func bytesToInt64(Value []byte) int64 {
	return int64(binary.LittleEndian.Uint64(Value))
}

func StringToAob(Value string) string {
	s := ""
	for _, v := range Value {
		s += fmt.Sprintf("%X ", v)
	}
	return s
}

func AobToArray(Aob string) []int32 {
	var fields = strings.Fields(Aob)
	var bytes []int32
	for _, v := range fields {
		if v != "??" {
			val, err := strconv.ParseInt(v, 16, 32)
			if err != nil {
				panic(err)
			}
			bytes = append(bytes, int32(val))
			continue
		}
		bytes = append(bytes, -1)
	}
	return bytes
}
