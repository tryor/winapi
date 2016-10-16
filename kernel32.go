package winapi

import "syscall"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetModuleHandleW    = modkernel32.NewProc("GetModuleHandleW")
	procFreeConsole         = modkernel32.NewProc("FreeConsole")
	procGetLastError        = modkernel32.NewProc("GetLastError")
	procGetLocaleInfoW      = modkernel32.NewProc("GetLocaleInfoW")
	procSetSystemPowerState = modkernel32.NewProc("SetSystemPowerState")
)

func GetModuleHandle(modname string) (handle HWND, err error) {
	var modname_ uintptr
	if modname == "" {
		modname_ = 0
	} else {
		modname_ = StringToUintptr(modname)
	}
	r0, _, e1 := syscall.Syscall(procGetModuleHandleW.Addr(), 1, modname_, 0, 0)
	handle = HWND(r0)
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
