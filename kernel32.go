package winapi

import (
	"reflect"
	"syscall"
	"unsafe"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetModuleHandleW    = modkernel32.NewProc("GetModuleHandleW")
	procFreeConsole         = modkernel32.NewProc("FreeConsole")
	procGetLastError        = modkernel32.NewProc("GetLastError")
	procGetLocaleInfoW      = modkernel32.NewProc("GetLocaleInfoW")
	procSetSystemPowerState = modkernel32.NewProc("SetSystemPowerState")
	ProcGetFullPathName     = modkernel32.NewProc("GetFullPathNameA")

	procGetCurrentThreadId = modkernel32.NewProc("GetCurrentThreadId")

	procCreateThread = modkernel32.NewProc("CreateThread")

	//procSetUnhandledExceptionFilter = modkernel32.NewProc("SetUnhandledExceptionFilter")

)

/*
HANDLE
WINAPI
CreateThread(
_In_opt_LPSECURITY_ATTRIBUTES lpThreadAttributes,
_In_SIZE_T dwStackSize,
_In_LPTHREAD_START_ROUTINE lpStartAddress,
_In_opt___drv_aliasesMemLPVOID lpParameter,
_In_DWORD dwCreationFlags,
_Out_opt_LPDWORD lpThreadId
);


DWORD WINAPI ThreadProc(LPVOID lpParam){
    return 0
    }

hThread[i]=CreateThread(
        NULL,//default security attributes
        0,//use default stack size
        ThreadProc,//thread function
        pData,//argument to thread function
        0,//use default creation flags
        &dwThreadId[i]);//returns the thread identifier

*/
type ThreadProc func(lpParam uintptr) uintptr

func CreateThread(proc ThreadProc) (h HANDLE, tid uintptr, err error) {
	r0, e1 := Syscall(procCreateThread.Addr(), 0, 0, syscall.NewCallback(proc), 0, 0, uintptr(unsafe.Pointer(&tid)))
	if e1 != nil {
		err = error(e1)
	} else {
		h = HANDLE(r0)
	}
	return
}

func GetCurrentThreadId() (ret uint, err error) {
	r0, e1 := Syscall(procGetCurrentThreadId.Addr())
	if e1 != nil {
		err = error(e1)
	} else {
		ret = uint(r0)
	}
	return
}

func GetFullPathName(lpFileName string) (string, error) {
	nBufferLength := 255
	lpBuffer := make([]byte, nBufferLength)
	var lpFilePart string
	ret, err := Syscall(ProcGetFullPathName.Addr(), StringToUintptr(lpFileName), uintptr(nBufferLength), uintptr(unsafe.Pointer(&lpBuffer[0])), (uintptr)(unsafe.Pointer(&lpFilePart)))
	if err == nil || reflect.ValueOf(err).IsNil() {
		return string(lpBuffer[0:int(ret)]), nil
	} else {
		return "", err
	}
}

func GetModuleHandle(modname string) (handle HINSTANCE, err error) {
	var modname_ uintptr
	if modname == "" {
		modname_ = 0
	} else {
		modname_ = StringToUintptr(modname)
	}
	r0, _, e1 := syscall.Syscall(procGetModuleHandleW.Addr(), 1, modname_, 0, 0)
	handle = HINSTANCE(r0)
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
