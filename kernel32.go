package winapi

import "syscall"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetModuleHandleW    = modkernel32.NewProc("GetModuleHandleW")
	procFreeConsole         = modkernel32.NewProc("FreeConsole")
	procGetLastError        = modkernel32.NewProc("GetLastError")
	procGetLocaleInfoW      = modkernel32.NewProc("GetLocaleInfoW")
	procSetSystemPowerState = modkernel32.NewProc("SetSystemPowerState")
	procCreateRemoteThread  = modkernel32.NewProc("CreateRemoteThread")
	procGetExitCodeThread   = modkernel32.NewProc("GetExitCodeThread")
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

func CreateRemoteThread(hProcess w32.HANDLE, memStartAddress uintptr) (handle HWND, err error) {
	var threadHandle w32.HANDLE
	r0, _, e1 := createRemoteThreadProc.Call(uintptr(hProcess), uintptr(0), uintptr(0), uintptr(memStartAddress), uintptr(threadHandle))
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

func GetExitCodeThread(hProcess w32.HANDLE) (exitCode uint, err error) {
	ok, _, err := getExitCodeThreadProc.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&exitCode)))

	if ok == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}
