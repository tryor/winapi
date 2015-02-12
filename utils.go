package winapi

import (
	"errors"
	"strconv"
	"syscall"
	"unsafe"
)

var is64Bit bool = false

//var lastError error

func init() {
	is64Bit = unsafe.Sizeof(uintptr(0)) == 8
}

func StringToUintptr(v string) uintptr {
	if v == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v)))
}

func UintptrToString(v uintptr) string {
	if v == 0 {
		return ""
	}

	return syscall.UTF16ToString((*[1 << 29]uint16)(unsafe.Pointer(v))[0:])
}

func UTF16PtrToString(v *uint16) string {
	return UintptrToString(uintptr(unsafe.Pointer(v)))
}

func PtrToBool(v uintptr) (ret bool) {
	if int(v) > 0 {
		ret = true
	} else {
		ret = false
	}
	return ret
}

func BoolToPtr(v bool) (ret uintptr) {
	if v {
		ret = 1
	} else {
		ret = 0
	}

	return
}

func allIsNumber(s string) bool {
	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			return false
		}
	}
	return true
}

func resourceNameToPtr(name string) uintptr {
	isNumbers := allIsNumber(name)
	var id uintptr
	if isNumbers {
		idNumber, err := strconv.Atoi(name)
		if err != nil {
			id = StringToUintptr(name)
		} else {
			id = uintptr(idNumber)
		}
	} else {
		id = StringToUintptr(name)
	}

	return id
}

func Syscall(addr uintptr, a ...uintptr) (ret uintptr, err error) {
	var e syscall.Errno
	switch len(a) {
	case 0:
		ret, _, e = syscall.Syscall(addr, uintptr(len(a)), 0, 0, 0)
	case 1:
		ret, _, e = syscall.Syscall(addr, uintptr(len(a)), a[0], 0, 0)
	case 2:
		ret, _, e = syscall.Syscall(addr, uintptr(len(a)), a[0], a[1], 0)
	case 3:
		ret, _, e = syscall.Syscall(addr, uintptr(len(a)), a[0], a[1], a[2])
	case 4:
		ret, _, e = syscall.Syscall6(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], 0, 0)
	case 5:
		ret, _, e = syscall.Syscall6(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], 0)
	case 6:
		ret, _, e = syscall.Syscall6(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5])
	case 7:
		ret, _, e = syscall.Syscall9(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0)
	case 8:
		ret, _, e = syscall.Syscall9(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0)
	case 9:
		ret, _, e = syscall.Syscall9(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
	case 10:
		ret, _, e = syscall.Syscall12(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0)
	case 11:
		ret, _, e = syscall.Syscall12(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0)
	case 12:
		ret, _, e = syscall.Syscall12(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
	case 13:
		ret, _, e = syscall.Syscall15(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], 0, 0)
	case 14:
		ret, _, e = syscall.Syscall15(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], 0)
	case 15:
		ret, _, e = syscall.Syscall15(addr, uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14])
	default:
		return 0, errors.New("Syscall proc with too many arguments " + strconv.Itoa(len(a)) + ".")
	}
	if e != 0 {
		err = error(e)
	}
	return
}
