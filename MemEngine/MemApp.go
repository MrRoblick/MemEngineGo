package MemEngine

import (
	"unsafe"
)

type MemApp struct {
	*Process
}

type MemoryBasicInformation struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	PartitionId       uint16
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

const memCommit = 0x1000
const pageReadonly = 0x02
const pageReadWrite = 0x04

func (m *MemApp) ReadBytes(Address uintptr, Length uint32) ([]byte, int) {
	var bytesReaded int
	var buffer = make([]byte, Length)
	call, _, _ := readProcessMemory.Call(m.Handle, Address, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)), uintptr(unsafe.Pointer(&bytesReaded)))
	if call == 0 {
		return buffer, 0
	}
	return buffer[:], bytesReaded
}

func (m *MemApp) ReadUint32(Address uintptr) (uint32, int) {
	v, s := m.ReadBytes(Address, 4)
	return BytesToUint32(v), s
}
func (m *MemApp) ReadUint64(Address uintptr) (uint64, int) {
	v, s := m.ReadBytes(Address, 8)
	return BytesToUint64(v), s
}

func (m *MemApp) ReadInt32(Address uintptr) (int32, int) {
	v, s := m.ReadBytes(Address, 4)
	return BytesToInt32(v), s
}
func (m *MemApp) ReadInt64(Address uintptr) (int64, int) {
	v, s := m.ReadBytes(Address, 8)
	return BytesToInt64(v), s
}

func (m *MemApp) ReadByte(Address uintptr) (byte, int) {
	v, s := m.ReadBytes(Address, 1)
	return v[0], s
}
func (m *MemApp) ReadBool(Address uintptr) (bool, int) {
	v, s := m.ReadBytes(Address, 1)
	var b = false
	if v[0] > 0 {
		b = true
	}
	return b, s
}

func (m *MemApp) ReadFloat32(Address uintptr) (float32, int) {
	v, s := m.ReadBytes(Address, 4)
	return BytesToFloat32(v), s
}

func (m *MemApp) ReadFloat64(Address uintptr) (float64, int) {
	v, s := m.ReadBytes(Address, 8)
	return BytesToFloat64(v), s
}

func (m *MemApp) ReadAsciiString(Address uintptr) string {
	str := ""
	for {
		Byte, _ := m.ReadByte(Address)
		Address++
		if Byte == 0 {
			break
		}
		str += string(Byte)
	}
	return str
}

func (m *MemApp) ReadVector3(Address uintptr) (Vector3, int) {
	v, s := m.ReadBytes(Address, 3*4)
	return NewVector3(
		BytesToFloat32(v[0:4]),
		BytesToFloat32(v[4:8]),
		BytesToFloat32(v[8:12]),
	), s
}
func (m *MemApp) ReadVector2(Address uintptr) (Vector2, int) {
	v, s := m.ReadBytes(Address, 2*4)
	return NewVector2(
		BytesToFloat32(v[0:4]),
		BytesToFloat32(v[4:8]),
	), s
}

func (m *MemApp) WriteBytes(Address uintptr, Buffer []byte) bool {
	var bytesWrited int
	call, _, _ := writeProcessMemory.Call(m.Handle, Address, uintptr(unsafe.Pointer(&Buffer[0])), uintptr(len(Buffer)), uintptr(unsafe.Pointer(&bytesWrited)))
	return call != 0 && bytesWrited == len(Buffer)
}
func (m *MemApp) WriteByte(Address uintptr, Value byte) bool {
	return m.WriteBytes(Address, []byte{Value})
}
func (m *MemApp) WriteBool(Address uintptr, Value bool) bool {
	var b byte = 0
	if Value {
		b = 255
	}
	return m.WriteBytes(Address, []byte{b})
}
func (m *MemApp) WriteFloat32(Address uintptr, Value float32) bool {
	return m.WriteBytes(Address, Float32ToBytes(Value))
}
func (m *MemApp) WriteFloat64(Address uintptr, Value float64) bool {
	return m.WriteBytes(Address, Float64ToBytes(Value))
}
func (m *MemApp) WriteUint32(Address uintptr, Value uint32) bool {
	return m.WriteBytes(Address, Uint32ToBytes(Value))
}
func (m *MemApp) WriteUint64(Address uintptr, Value uint64) bool {
	return m.WriteBytes(Address, Uint64ToBytes(Value))
}

func (m *MemApp) WriteInt32(Address uintptr, Value int32) bool {
	return m.WriteBytes(Address, Int32ToBytes(Value))
}
func (m *MemApp) WriteInt64(Address uintptr, Value int64) bool {
	return m.WriteBytes(Address, Int64ToBytes(Value))
}

func (m *MemApp) WriteAsciiString(Address uintptr, Value string) bool {
	return m.WriteBytes(Address, []byte(Value))
}
func (m *MemApp) WriteVector3(Address uintptr, Value Vector3) bool {
	return m.WriteBytes(Address, Value.ConvertToBytes())
}
func (m *MemApp) WriteVector2(Address uintptr, Value Vector2) bool {
	return m.WriteBytes(Address, Value.ConvertToBytes())
}

func (m *MemApp) ScanMemory(Aob string) []uintptr {
	var results []uintptr
	var signature = AobToArray(Aob)

	var currentAddress uintptr
	var stopAddress uintptr = 0x00007fffffffffff

	memBasicInfo := MemoryBasicInformation{}
	size := unsafe.Sizeof(memBasicInfo)

	for {
		s, _, _ := virtualQueryEx.Call(m.Handle, currentAddress, uintptr(unsafe.Pointer(&memBasicInfo)), size)
		if s == 0 || currentAddress > stopAddress {
			break
		}
		if memBasicInfo.State == memCommit && memBasicInfo.Type != pageReadonly {
			buffer, readed := m.ReadBytes(memBasicInfo.BaseAddress, uint32(memBasicInfo.RegionSize))
			for i := 0; i < readed-len(signature); i++ {
				match := true
				for j := 0; j < len(signature); j++ {
					if signature[j] >= 0 && buffer[i+j] != byte(signature[j]) {
						match = false
						break
					}
				}
				if match {
					results = append(results, memBasicInfo.BaseAddress+uintptr(i))
				}
			}
		}
		currentAddress += memBasicInfo.RegionSize
	}
	return results
}
