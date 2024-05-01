package MemEngine

import (
	"syscall"
	"unsafe"
)

type Module struct {
	Name                 string
	BaseAddress          uintptr
	ModuleSize, ModuleId uint32
}
type Process struct {
	ProcessId uint32
	Handle    uintptr
	Modules   []Module
}

type ModuleEntry struct {
	Size           uint32
	Th32ModuleId   uint32
	Th32ProcessId  uint32
	GlblcntUsage   uint32
	ProccntUsage   uint32
	ModBaseAddress *byte
	ModBaseSize    uint32
	HModule        uintptr
	SzModule       [255]uint8
	SzExePath      [260]uint8
}

const ProcessAllAccess = 0x000F0000 | 0x00100000 | 0xFFFF

func (p *Process) Close() {
	err := syscall.CloseHandle(syscall.Handle(p.Handle))
	if err != nil {
		panic(err)
	}
}

func (p *Process) GetModule(Name string) *Module {
	for _, mod := range p.Modules {
		if mod.Name == Name {
			return &mod
		}
	}
	return nil
}

func findProcess(Name string) *Process {
	handle, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	defer func(handle syscall.Handle) {
		err := syscall.CloseHandle(handle)
		if err != nil {
			panic(err)
		}
	}(handle)
	procEntry := syscall.ProcessEntry32{}
	procEntry.Size = uint32(unsafe.Sizeof(procEntry))

	err = syscall.Process32First(handle, &procEntry)
	if err != nil {
		panic(err)
	}

	for {
		if syscall.UTF16ToString(procEntry.ExeFile[:]) == Name {
			handleModule, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPMODULE, procEntry.ProcessID)
			if err != nil {
				panic(err)
			}

			var Modules []Module

			moduleEntry := ModuleEntry{}
			moduleEntry.Size = uint32(unsafe.Sizeof(moduleEntry))

			success, _, message := module32First.Call(uintptr(handleModule), uintptr(unsafe.Pointer(&moduleEntry)))
			if success == 0 {
				panic(message)
			}

			for {
				Length := 0
				for idx, v := range moduleEntry.SzModule {
					if v == 0 {
						Length = idx
						break
					}
				}
				Modules = append(Modules, Module{
					Name:        string(moduleEntry.SzModule[:Length]),
					BaseAddress: uintptr(unsafe.Pointer(moduleEntry.ModBaseAddress)),
					ModuleId:    moduleEntry.Th32ModuleId,
					ModuleSize:  moduleEntry.ModBaseSize,
				})
				moduleEntry.SzModule = [255]uint8{}
				success, _, message = module32Next.Call(uintptr(handleModule), uintptr(unsafe.Pointer(&moduleEntry)))
				if success == 0 {
					break
				}
			}
			err = syscall.CloseHandle(handleModule)
			if err != nil {
				panic(err)
			}
			OpenHandle, err := syscall.OpenProcess(ProcessAllAccess, false, procEntry.ProcessID)
			if err != nil {
				panic(err)
			}
			return &Process{
				ProcessId: procEntry.ProcessID,
				Handle:    uintptr(OpenHandle),
				Modules:   Modules,
			}
		}
		err = syscall.Process32Next(handle, &procEntry)
		if err != nil {
			return nil
		}
	}

}
