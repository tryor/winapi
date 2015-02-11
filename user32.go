package winapi

import "unsafe"
import "syscall"

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procGetDC      = moduser32.NewProc("GetDC")
	procBeginPaint = moduser32.NewProc("BeginPaint")
	procEndPaint   = moduser32.NewProc("EndPaint")
	procReleaseDC  = moduser32.NewProc("ReleaseDC")

	procRegisterClassExW = moduser32.NewProc("RegisterClassExW")
	procCreateWindowExW  = moduser32.NewProc("CreateWindowExW")
	procDefWindowProcW   = moduser32.NewProc("DefWindowProcW")
	procDestroyWindow    = moduser32.NewProc("DestroyWindow")
	procPostQuitMessage  = moduser32.NewProc("PostQuitMessage")
	procShowWindow       = moduser32.NewProc("ShowWindow")
	procUpdateWindow     = moduser32.NewProc("UpdateWindow")
	procGetMessageW      = moduser32.NewProc("GetMessageW")
	procTranslateMessage = moduser32.NewProc("TranslateMessage")
	procDispatchMessageW = moduser32.NewProc("DispatchMessageW")
	procLoadIconW        = moduser32.NewProc("LoadIconW")
	procLoadCursorW      = moduser32.NewProc("LoadCursorW")
	procSetCursor        = moduser32.NewProc("SetCursor")
	procSendMessageW     = moduser32.NewProc("SendMessageW")
	procPostMessageW     = moduser32.NewProc("PostMessageW")

	procGetKeyboardState = moduser32.NewProc("GetKeyboardState")
	procSetFocus         = moduser32.NewProc("SetFocus")
)

func GetDC(hwnd HANDLE) (hdc HDC) {
	r0, _, _ := syscall.Syscall(procGetDC.Addr(), 1, uintptr(hwnd), 0, 0)
	hdc = HDC(r0)
	return hdc
}

func ReleaseDC(h HANDLE, hdc HDC) bool {
	var ret uintptr
	ret, _, _ = procReleaseDC.Call(uintptr(h), uintptr(hdc))
	return ret == 1
}

func BeginPaint(hwnd HANDLE, ps *PAINTSTRUCT) (hdc HDC) {
	r0, _, _ := syscall.Syscall(procBeginPaint.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(ps)), 0)
	hdc = HDC(r0)
	return
}

func EndPaint(hwnd HANDLE, ps *PAINTSTRUCT) bool {
	syscall.Syscall(procEndPaint.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(ps)), 0)
	return true
}

func RegisterClassEx(wndclass *Wndclassex) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassExW.Addr(), 1, uintptr(unsafe.Pointer(wndclass)), 0, 0)
	atom = uint16(r0)
	if atom == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateWindowEx(exstyle uint32, classname *uint16, windowname *uint16, style uint32, x int32, y int32, width int32, height int32, wndparent HANDLE, menu HANDLE, instance HANDLE, param uintptr) (hwnd HANDLE, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExW.Addr(), 12, uintptr(exstyle), uintptr(unsafe.Pointer(classname)), uintptr(unsafe.Pointer(windowname)), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(wndparent), uintptr(menu), uintptr(instance), uintptr(param))
	hwnd = HANDLE(r0)
	if hwnd == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DefWindowProc(hwnd HANDLE, msg uint32, wparam uintptr, lparam uintptr) (lresult uintptr) {
	r0, _, _ := syscall.Syscall6(procDefWindowProcW.Addr(), 4, uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam), 0, 0)
	lresult = uintptr(r0)
	return
}

func DestroyWindow(hwnd HANDLE) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	if int(r1) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func PostQuitMessage(exitcode int32) {
	syscall.Syscall(procPostQuitMessage.Addr(), 1, uintptr(exitcode), 0, 0)
	return
}

func ShowWindow(hwnd HANDLE, cmdshow int32) (wasvisible bool) {
	r0, _, _ := syscall.Syscall(procShowWindow.Addr(), 2, uintptr(hwnd), uintptr(cmdshow), 0)
	wasvisible = bool(r0 != 0)
	return
}

func UpdateWindow(hwnd HANDLE) (err error) {
	r1, _, e1 := syscall.Syscall(procUpdateWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	if int(r1) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetMessage(msg *Msg, hwnd HANDLE, MsgFilterMin uint32, MsgFilterMax uint32) (ret int32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetMessageW.Addr(), 4, uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(MsgFilterMin), uintptr(MsgFilterMax), 0, 0)
	ret = int32(r0)
	if ret == -1 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func TranslateMessage(msg *Msg) (done bool) {
	r0, _, _ := syscall.Syscall(procTranslateMessage.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	done = bool(r0 != 0)
	return
}

func DispatchMessage(msg *Msg) (ret int32) {
	r0, _, _ := syscall.Syscall(procDispatchMessageW.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	ret = int32(r0)
	return
}

func LoadIcon(instance HANDLE, iconname *uint16) (icon HANDLE, err error) {
	r0, _, e1 := syscall.Syscall(procLoadIconW.Addr(), 2, uintptr(instance), uintptr(unsafe.Pointer(iconname)), 0)
	icon = HANDLE(r0)
	if icon == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LoadCursor(instance HANDLE, cursorname *uint16) (cursor HANDLE, err error) {
	r0, _, e1 := syscall.Syscall(procLoadCursorW.Addr(), 2, uintptr(instance), uintptr(unsafe.Pointer(cursorname)), 0)
	cursor = HANDLE(r0)
	if cursor == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetCursor(cursor HANDLE) (precursor HANDLE, err error) {
	r0, _, e1 := syscall.Syscall(procSetCursor.Addr(), 1, uintptr(cursor), 0, 0)
	precursor = HANDLE(r0)
	if precursor == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SendMessage(hwnd HANDLE, msg uint32, wparam uintptr, lparam uintptr) (lresult uintptr) {
	r0, _, _ := syscall.Syscall6(procSendMessageW.Addr(), 4, uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam), 0, 0)
	lresult = uintptr(r0)
	return
}

func PostMessage(hwnd HANDLE, msg uint32, wparam uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procPostMessageW.Addr(), 4, uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam), 0, 0)
	if int(r1) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetKeyboardState() (keyState []byte, err error) {
	var keys [256]byte
	r0, _, e1 := syscall.Syscall(procGetKeyboardState.Addr(), 1, uintptr(unsafe.Pointer(&keys)), 0, 0)
	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	keyState = keys[:]
	return
}

func SetFocus(hwnd HANDLE) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFocus.Addr(), 1, uintptr(hwnd), 0, 0)
	if int(r1) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
