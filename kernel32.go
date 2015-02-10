package winapi

import "unsafe"
import "syscall"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetModuleHandleW = modkernel32.NewProc("GetModuleHandleW")
	procFreeConsole      = modkernel32.NewProc("FreeConsole")
)

func GetModuleHandle(modname *uint16) (handle HANDLE, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleHandleW.Addr(), 1, uintptr(unsafe.Pointer(modname)), 0, 0)
	handle = HANDLE(r0)
	if handle == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func FreeConsole() (ret uint64, err error) {
	r0, _, e1 := syscall.Syscall(procFreeConsole.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
	} else {
		ret = uint64(r0)
	}
	return
}
