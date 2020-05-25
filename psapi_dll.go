package winapi

import (
	"syscall"
	"unsafe"
)

var (
	modpsapi                 = syscall.NewLazyDLL("psapi.dll")
	procGetModuleFileNameExW = modpsapi.NewProc("GetModuleFileNameExW")
)

func GetModuleFileNameExW(h HANDLE) (filename string) {
	buff := make([]uint16, 255+1)
	_, _ = Syscall(procGetModuleFileNameExW.Addr(), uintptr(h), 0, uintptr(unsafe.Pointer(&buff[0])), 255)
	filename = syscall.UTF16ToString(buff)
	return
}
