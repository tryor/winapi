package winapi

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procSetCursor        = moduser32.NewProc("SetCursor")
	procGetKeyboardState = moduser32.NewProc("GetKeyboardState")
	procSetFocus         = moduser32.NewProc("SetFocus")

	procBeginPaint          = moduser32.NewProc("BeginPaint")
	procCreateDialogParamW  = moduser32.NewProc("CreateDialogParamW")
	procCreateWindowExW     = moduser32.NewProc("CreateWindowExW")
	procCreateWindowExA     = moduser32.NewProc("CreateWindowExA")
	procDefWindowProcW      = moduser32.NewProc("DefWindowProcW")
	procDestroyWindow       = moduser32.NewProc("DestroyWindow")
	procDialogBoxParamW     = moduser32.NewProc("DialogBoxParamW")
	procDispatchMessageW    = moduser32.NewProc("DispatchMessageW")
	procEndDialog           = moduser32.NewProc("EndDialog")
	procEndPaint            = moduser32.NewProc("EndPaint")
	procGetDC               = moduser32.NewProc("GetDC")
	procGetDlgItem          = moduser32.NewProc("GetDlgItem")
	procGetMessageW         = moduser32.NewProc("GetMessageW")
	procGetWindowLongW      = moduser32.NewProc("GetWindowLongW")
	procGetWindowLongPtrW   = moduser32.NewProc("GetWindowLongPtrW")
	procLoadCursorW         = moduser32.NewProc("LoadCursorW")
	procLoadIconW           = moduser32.NewProc("LoadIconW")
	procLoadMenuW           = moduser32.NewProc("LoadMenuW")
	procLoadStringW         = moduser32.NewProc("LoadStringW")
	procMessageBoxW         = moduser32.NewProc("MessageBoxW")
	procUnregisterClassW    = moduser32.NewProc("UnregisterClassW")
	procPostMessageW        = moduser32.NewProc("PostMessageW")
	procPostQuitMessage     = moduser32.NewProc("PostQuitMessage")
	procRegisterClassExW    = moduser32.NewProc("RegisterClassExW")
	procRegisterClassW      = moduser32.NewProc("RegisterClassW")
	procReleaseDC           = moduser32.NewProc("ReleaseDC")
	procSendMessageW        = moduser32.NewProc("SendMessageW")
	procSendDlgItemMessageW = moduser32.NewProc("SendDlgItemMessageW")
	procSetMenu             = moduser32.NewProc("SetMenu")
	procSetWindowLongW      = moduser32.NewProc("SetWindowLongW")
	procSetWindowLongPtrW   = moduser32.NewProc("SetWindowLongPtrW")
	procShowWindow          = moduser32.NewProc("ShowWindow")
	procTranslateMessage    = moduser32.NewProc("TranslateMessage")
	procUpdateWindow        = moduser32.NewProc("UpdateWindow")
	procRedrawWindow        = moduser32.NewProc("RedrawWindow")
	procInvalidateRect      = moduser32.NewProc("InvalidateRect")
)

func GetDC(hwnd HWND) (hdc HDC) {
	r0, _ := Syscall(procGetDC.Addr(), uintptr(hwnd))
	hdc = HDC(r0)
	return
}

func ReleaseDC(h HWND, hdc HDC) bool {
	r0, _, _ := procReleaseDC.Call(uintptr(h), uintptr(hdc))
	return r0 == 1
}

func BeginPaint(hwnd HWND, ps *PAINTSTRUCT) (hdc HDC) {
	r0, _ := Syscall(procBeginPaint.Addr(), uintptr(hwnd), uintptr(unsafe.Pointer(ps)))
	hdc = HDC(r0)
	return
}

func EndPaint(hwnd HWND, ps *PAINTSTRUCT) bool {
	r0, _ := Syscall(procEndPaint.Addr(), uintptr(hwnd), uintptr(unsafe.Pointer(ps)))
	return PtrToBool(r0)
}

func RegisterClassExW(wndclass *Wndclassex) (atom uint16, err error) {
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

func RegisterClassW(wndclass *Wndclass) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassW.Addr(), 1, uintptr(unsafe.Pointer(wndclass)), 0, 0)
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

func CreateWindowExW(exstyle uint32, classname string, windowname string, style uint32, x int32, y int32, width int32, height int32, wndparent HWND, menu HMENU, instance HINSTANCE, param uintptr) (hwnd HWND, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExW.Addr(), 12, uintptr(exstyle), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(classname))), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowname))), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(wndparent), uintptr(menu), uintptr(instance), uintptr(param))
	hwnd = HWND(r0)
	if hwnd == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateWindowExA(exstyle uint32, classname string, windowname string, style uint32, x int32, y int32, width int32, height int32, wndparent HWND, menu HMENU, instance HINSTANCE, param uintptr) (hwnd HWND, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExA.Addr(), 12, uintptr(exstyle), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(classname))), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowname))), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(wndparent), uintptr(menu), uintptr(instance), uintptr(param))
	hwnd = HWND(r0)
	if hwnd == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DefWindowProcW(hwnd HWND, msg UINT, wparam WPARAM, lparam LPARAM) (lresult uintptr) {
	r0, _, _ := syscall.Syscall6(procDefWindowProcW.Addr(), 4, uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam), 0, 0)
	lresult = uintptr(r0)
	return
}

func DestroyWindow(hwnd HWND) (err error) {
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

func ShowWindow(hwnd HWND, cmdshow int32) (wasvisible bool) {
	r0, _, _ := syscall.Syscall(procShowWindow.Addr(), 2, uintptr(hwnd), uintptr(cmdshow), 0)
	wasvisible = bool(r0 != 0)
	return
}

func UpdateWindow(hwnd HWND) (err error) {
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

func GetMessage(msg *Msg, hwnd HWND, MsgFilterMin uint32, MsgFilterMax uint32) (ret int32, err error) {
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

func DispatchMessageW(msg *Msg) (ret int32) {
	r0, _, _ := syscall.Syscall(procDispatchMessageW.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	ret = int32(r0)
	return
}

func LoadIconS(instance HINSTANCE, iconname string) (icon HICON, err error) {
	return LoadIconW(instance, resourceNameToPtr(iconname))
}

func LoadIcon(instance HINSTANCE, iconname *uint16) (icon HICON, err error) {
	return LoadIconW(instance, uintptr(unsafe.Pointer(iconname)))
}

func LoadIconW(instance HINSTANCE, iconname uintptr) (icon HICON, err error) {
	r0, _, e1 := syscall.Syscall(procLoadIconW.Addr(), 2, uintptr(instance), iconname, 0)
	icon = HICON(r0)
	if icon == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LoadCursorS(instance HINSTANCE, cursorname string) (cursor HCURSOR, err error) {
	return LoadCursorW(instance, StringToUintptr(cursorname))
}

func LoadCursor(instance HINSTANCE, cursorname *uint16) (cursor HCURSOR, err error) {
	return LoadCursorW(instance, uintptr(unsafe.Pointer(cursorname)))
}

func LoadCursorW(instance HINSTANCE, cursorname uintptr) (cursor HCURSOR, err error) {
	r0, _, e1 := syscall.Syscall(procLoadCursorW.Addr(), 2, uintptr(instance), cursorname, 0)
	cursor = HCURSOR(r0)
	if cursor == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetCursor(cursor HCURSOR) (precursor HCURSOR, err error) {
	r0, _, e1 := syscall.Syscall(procSetCursor.Addr(), 1, uintptr(cursor), 0, 0)
	precursor = HCURSOR(r0)
	if precursor == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SendMessage(hwnd HWND, msg UINT, wparam WPARAM, lparam LPARAM) (lresult uintptr) {
	r0, _, _ := syscall.Syscall6(procSendMessageW.Addr(), 4, uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam), 0, 0)
	lresult = uintptr(r0)
	return
}

func PostMessage(hwnd HWND, msg UINT, wparam WPARAM, lparam LPARAM) (err error) {
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

func SetFocus(hwnd HWND) (err error) {
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

//------------------------------------------------

func DialogBoxParam(instRes HINSTANCE, name string, parent HWND,
	proc uintptr, param uintptr) (int, error) {
	ret, err := Syscall(procDialogBoxParamW.Addr(), uintptr(instRes), resourceNameToPtr(name), uintptr(parent), proc, param)
	return int(ret), err
}

func EndDialog(h HWND, result int) (bool, error) {
	ret, err := Syscall(procEndDialog.Addr(), uintptr(h), uintptr(result))
	return PtrToBool(ret), err
}

func GetDlgItem(h HWND, id int) HWND {
	ret, _ := Syscall(procGetDlgItem.Addr(), uintptr(h), uintptr(id))
	return HWND(ret)
}

func GetWindowLongPtr(h HWND, index int) (ret uintptr, err error) {
	if is64Bit {
		ret, err = Syscall(procGetWindowLongPtrW.Addr(), uintptr(h), uintptr(index))
	} else {
		ret, err = Syscall(procGetWindowLongW.Addr(), uintptr(h), uintptr(index))
	}
	return
}

func LoadMenu(instRes HINSTANCE, name string) (HMENU, error) {
	ret, err := Syscall(procLoadMenuW.Addr(), uintptr(instRes), resourceNameToPtr(name))
	return HMENU(ret), err
}

func LoadString(inst HINSTANCE, id uint) (ret string, err error) {
	var text [4096]uint16
	var r uintptr
	r, err = Syscall(procLoadStringW.Addr(), uintptr(inst), uintptr(id),
		uintptr(unsafe.Pointer(&text[0])), 4096)
	if int(r) <= 0 {
		ret = ""
	} else {
		ret = string(utf16.Decode(text[0:r]))
	}
	return
}

func MessageBoxW(parent HWND, text, title string, boxType BoxType) (int, error) {
	ret, err := Syscall(procMessageBoxW.Addr(), uintptr(parent),
		StringToUintptr(text), StringToUintptr(title), uintptr(boxType))
	return int(ret), err
}

func UnregisterClassW(name string) (bool, error) {
	ret, err := Syscall(procUnregisterClassW.Addr(), StringToUintptr(name), 0)
	return PtrToBool(ret), err
}

func SendDlgItemMessageW(m *Msg, id int) (LRESULT, error) {
	ret, err := Syscall(procSendDlgItemMessageW.Addr(), uintptr(m.Hwnd), uintptr(id),
		uintptr(m.Message), uintptr(m.Wparam), uintptr(m.Lparam))
	return LRESULT(ret), err
}

func SetMenu(hwnd HWND, menu HMENU) (bool, error) {
	ret, err := Syscall(procSetMenu.Addr(), uintptr(hwnd), uintptr(menu))
	return PtrToBool(ret), err
}

func SetWindowLongPtrW(h HWND, index int, value uintptr) (ret uintptr, err error) {
	if is64Bit {
		ret, err = Syscall(procSetWindowLongPtrW.Addr(), uintptr(h), uintptr(index), value)
	} else {
		ret, err = Syscall(procSetWindowLongW.Addr(), uintptr(h), uintptr(index), value)
	}
	return
}

//prgnUpdate *CRgn
func RedrawWindow(hwnd HWND, lpRectUpdate LPCRECT, prgnUpdate HANDLE, flags UINT) (bool, error) {
	ret, err := Syscall(procRedrawWindow.Addr(), uintptr(hwnd), uintptr(unsafe.Pointer(lpRectUpdate)), uintptr(prgnUpdate), uintptr(flags))
	return PtrToBool(ret), err
}

// LPCRECT lpRect, BOOL bErase = TRUE
func InvalidateRect(hwnd HWND, lpRectUpdate LPCRECT, bErase BOOL) error {
	var bErase_ int8
	if bErase {
		bErase_ = 1
	}
	_, err := Syscall(procInvalidateRect.Addr(), uintptr(hwnd), uintptr(unsafe.Pointer(lpRectUpdate)), uintptr(bErase_))
	return err
}
