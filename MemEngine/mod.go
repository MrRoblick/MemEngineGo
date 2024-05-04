package MemEngine

import (
	"syscall"
)

var kernel = syscall.MustLoadDLL("kernel32.dll")

var module32First = kernel.MustFindProc("Module32First")
var module32Next = kernel.MustFindProc("Module32Next")

var readProcessMemory = kernel.MustFindProc("ReadProcessMemory")
var writeProcessMemory = kernel.MustFindProc("WriteProcessMemory")

var virtualQueryEx = kernel.MustFindProc("VirtualQueryEx")

func NewApplication(ProcessName string) *MemApp {
	return &MemApp{
		findProcess(ProcessName),
	}
}
