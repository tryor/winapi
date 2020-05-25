package winapi

import (
	"errors"
	"log"
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

	procGetForegroundWindow = moduser32.NewProc("GetForegroundWindow")
	procGetActiveWindow     = moduser32.NewProc("GetActiveWindow")

	procTrackPopupMenu    = moduser32.NewProc("TrackPopupMenu")
	procSetActiveWindow   = moduser32.NewProc("SetActiveWindow")
	procAttachThreadInput = moduser32.NewProc("AttachThreadInput")

	procSetTimer  = moduser32.NewProc("SetTimer")
	procKillTimer = moduser32.NewProc("KillTimer")

	procLoadBitmapW = moduser32.NewProc("LoadBitmapW")

	procGetWindowRect  = moduser32.NewProc("GetWindowRect")
	procGetMessageTime = moduser32.NewProc("GetMessageTime")
	procPeekMessageW   = moduser32.NewProc("PeekMessageW")

	procSetParent = moduser32.NewProc("SetParent")

	procCallWindowProcW = moduser32.NewProc("CallWindowProcW")
	procScreenToClient  = moduser32.NewProc("ScreenToClient")

	procSetWindowPos     = moduser32.NewProc("SetWindowPos")
	procBringWindowToTop = moduser32.NewProc("BringWindowToTop")

	procAdjustWindowRect = moduser32.NewProc("AdjustWindowRect")

	procMoveWindow = moduser32.NewProc("MoveWindow")

	procEnableWindow = moduser32.NewProc("EnableWindow")

	procSetWindowTextW = moduser32.NewProc("SetWindowTextW")
	procGetWindowTextW = moduser32.NewProc("GetWindowTextW")

	procSetClassLong = moduser32.NewProc("SetClassLongW")

	procSetLayeredWindowAttributes = moduser32.NewProc("SetLayeredWindowAttributes")

	procGetDesktopWindow = moduser32.NewProc("GetDesktopWindow")

	procCreatePopupMenu = moduser32.NewProc("CreatePopupMenu")
	procAppendMenu      = moduser32.NewProc("AppendMenuW")

	procLoadImage = moduser32.NewProc("LoadImageW")

	procGetClientRect = moduser32.NewProc("GetClientRect")

	procClientToScreen = moduser32.NewProc("ClientToScreen")

	procGetCursorPos = moduser32.NewProc("GetCursorPos")

	procSetForegroundWindow = moduser32.NewProc("SetForegroundWindow")

	procMessageBoxTimeoutW = moduser32.NewProc("MessageBoxTimeoutW")

	procGetWindowThreadProcessId = moduser32.NewProc("GetWindowThreadProcessId")

	procRegisterHotKey   = moduser32.NewProc("RegisterHotKey")
	procUnregisterHotKey = moduser32.NewProc("UnregisterHotKey")

	procGetKeyState = moduser32.NewProc("GetKeyState")

	procSetCursorPos = moduser32.NewProc("SetCursorPos")

	procMouseEvent = moduser32.NewProc("mouse_event")
	procKeyBDEvent = moduser32.NewProc("keybd_event")
)

/*
　　VOID keybd_event(

　　BYTE bVk, // virtual-key code

　　BYTE bScan, // hardware scan code

　　DWORD dwFlags, // flags specifying various function options

　　DWORD dwExtraInfo // additional data associated with keystroke

　　);

*/

func SimulateKeyBDEvent(bVk BYTE, dwFlags DWORD) bool {
	r, _, _ := procKeyBDEvent.Call(uintptr(bVk), 0, uintptr(dwFlags), 0)
	return r != 0
}

// VOID mouse_event(
//   DWORD     dwFlags,     // motion and click options
//   DWORD     dx,          // horizontal position or change
//   DWORD     dy,          // vertical position or change
//   DWORD     dwData,      // wheel movement
//   ULONG_PTR dwExtraInfo  // application-defined information
// );

func SimulateMouseEvent(flags MouseEventFlags, dx, dy INT) bool {
	r, _, _ := procMouseEvent.Call(uintptr(flags), uintptr(dx), uintptr(dy), 0, 0)
	return r != 0
}

func SetCursorPos(x, y INT) bool {
	r, _, _ := procSetCursorPos.Call(uintptr(x), uintptr(y))
	return r != 0
}

func GetKeyState(nVirtKey INT) (state SHORT) {
	r, _, _ := procGetKeyState.Call(uintptr(nVirtKey))
	state = SHORT(r)
	return
}

//__in_opt HWND hWnd,
//__in int id,
//__in UINT fsModifiers,
//__in UINT vk

func RegisterHotKey(hwnd HWND, id INT, fsModifiers UINT, vk UINT) (err error) {
	r, _, e1 := procRegisterHotKey.Call(uintptr(hwnd), uintptr(id), uintptr(fsModifiers), uintptr(vk))
	if r == 0 {
		if e1 == nil {
			err = syscall.EINVAL
		} else {
			err = error(e1)
		}
	}
	return
}

func UnregisterHotKey(hwnd HWND, id INT) (err error) {
	r, _, e1 := procUnregisterHotKey.Call(uintptr(hwnd), uintptr(id))
	if r == 0 {
		if e1 == nil {
			err = syscall.EINVAL
		} else {
			err = error(e1)
		}
	}
	return
}

func GetWindowThreadProcessId(hwnd HWND) (id, pid uintptr) {
	id, _, _ = procGetWindowThreadProcessId.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&pid)))
	return
}

func SetForegroundWindow(hWnd HWND) bool {
	r, _, _ := procSetForegroundWindow.Call(uintptr(hWnd))
	return r != 0
}

func GetCursorPos() (p POINT, err error) {
	r, e1 := Syscall(procGetCursorPos.Addr(), uintptr(unsafe.Pointer(&p)))
	if r == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}

	}
	return
}

//BOOL GetClientRect(
//HWND hWnd, // 窗口句柄
//LPRECT lpRect // 客户区坐标
//);

func GetClientRect(hWnd HWND) (rect RECT, err error) {
	r, e1 := Syscall(procGetClientRect.Addr(), uintptr(hWnd), uintptr(unsafe.Pointer(&rect)))
	if r == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}

	}
	return
}

//    _In_opt_ HINSTANCE hInst,
//    _In_ LPCWSTR name,
//    _In_ UINT type,
//    _In_ int cx,
//    _In_ int cy,
//    _In_ UINT fuLoad);

func LoadImageFromFile(name string, typ IMAGE_TYPE) (h HWND, err error) {
	return LoadImageW(0, StringToUintptr(name), typ, 0, 0, LR_LOADFROMFILE)
}

func LoadImageW(hInst HINSTANCE, name uintptr, typ IMAGE_TYPE, cx, cy int, fuLoad UINT) (h HWND, err error) {
	r, e1 := Syscall(procLoadImage.Addr(), uintptr(hInst), name, uintptr(typ), uintptr(cx), uintptr(cy), uintptr(fuLoad))
	if r == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}

	} else {
		h = HWND(r)
	}
	return
}

func AppendMenuString(h HWND, wFlags, wIDNewItem int, lpNewItem string) (mh HWND, err error) {
	wFlags = wFlags | MF_STRING

	return AppendMenu(h, wFlags, wIDNewItem, StringToUintptr(lpNewItem))
}

func AppendMenu(h HWND, wFlags, wIDNewItem int, lpNewItem uintptr) (mh HWND, err error) {
	r, e1 := Syscall(procAppendMenu.Addr(), uintptr(h), uintptr(wFlags), uintptr(wIDNewItem), uintptr(lpNewItem))
	if r == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}

	}
	h = HWND(r)
	return
}

func CreatePopupMenu() (HWND, error) {
	r, err := Syscall(procCreatePopupMenu.Addr())
	if r == 0 {
		if err != nil {
			return 0, err
		} else {
			return 0, errors.New("CreatePopupMenu error")
		}

	}
	return HWND(r), err
}

func GetDesktopWindow() (HWND, error) {
	r, err := Syscall(procGetDesktopWindow.Addr())
	return HWND(r), err
}

func SetLayeredWindowAttributes(hWnd HWND, crKey COLORREF, bAlpha BYTE, dwFlags DWORD) (bool, error) {
	ret, err := Syscall(procSetLayeredWindowAttributes.Addr(), uintptr(hWnd), uintptr(crKey), uintptr(bAlpha), uintptr(dwFlags))
	return PtrToBool(ret), err
}

func SetClassLong(hWnd HWND, nIndex int, dwNewLong HCURSOR) (bool, error) {
	ret, err := Syscall(procSetClassLong.Addr(), uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return PtrToBool(ret), err
}

func GetWindowText(hWnd HWND) (text string, err error) {
	buff := make([]uint16, 255+1)
	_, err = Syscall(procGetWindowTextW.Addr(), uintptr(hWnd), uintptr(unsafe.Pointer(&buff[0])), 255)
	return syscall.UTF16ToString(buff), err
}

func SetWindowTextW(hWnd HWND, text string) (bool, error) {
	ret, err := Syscall(procSetWindowTextW.Addr(), uintptr(hWnd), StringToUintptr(text))
	return PtrToBool(ret), err
}

func EnableWindow(hWnd HWND, b bool) (bool, error) {
	ret, err := Syscall(procEnableWindow.Addr(), uintptr(hWnd), BoolToPtr(b))
	return PtrToBool(ret), err
}

func MoveWindow(hWnd HWND, x, y, w, h int32) (bool, error) {
	ret, err := Syscall(procMoveWindow.Addr(), uintptr(hWnd), uintptr(x), uintptr(y), uintptr(w), uintptr(h), BoolToPtr(false))
	return PtrToBool(ret), err
}

//func SetWindowRect(hWnd HWND, x, y, w, h int32) (bool, error) {

//}

func AdjustWindowRect(lpRectUpdate LPCRECT, style uintptr, bMenu bool) (bool, error) {
	ret, err := Syscall(procAdjustWindowRect.Addr(), uintptr(unsafe.Pointer(lpRectUpdate)), style, BoolToPtr(bMenu))
	return PtrToBool(ret), err
}

func BringWindowToTop(hWnd HWND) bool {
	r, _, _ := procBringWindowToTop.Call(uintptr(hWnd))
	return r != 0
}

func SetWindowPos(hWnd, hWndInsertAfter HWND, x, y, cx, cy int32, uFlags uint32) bool {
	r, _, _ := procSetWindowPos.Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return r != 0
}

func ScreenToClient(hWnd HWND, lpPoint *POINT) bool {
	r, _, _ := procScreenToClient.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

func ClientToScreen(hWnd HWND, lpPoint *POINT) bool {
	r, _, _ := procClientToScreen.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpPoint)))
	return r != 0
}

func CallWindowProcW(lpPrevWndFunc uintptr, hWnd HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := procCallWindowProcW.Call(lpPrevWndFunc, uintptr(hWnd), uintptr(Msg), wParam, lParam)
	return r
}

func SetParent(hWndChild, hWndNewParent HWND) (h HWND, err error) {
	r0, _, e1 := procSetParent.Call(uintptr(hWndChild), uintptr(hWndNewParent))
	if r0 == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	h = HWND(r0)
	return
}

func PeekMessage(msg *Msg, hWnd HWND, wMSGfilterMin, wMsgFilterMax, wRemoveMsg UINT) bool {
	r0, _, _ := procPeekMessageW.Call(uintptr(unsafe.Pointer(msg)), uintptr(hWnd), uintptr(wMSGfilterMin), uintptr(wMsgFilterMax), uintptr(wRemoveMsg))
	return r0 != 0
}
func GetMessageTime() DWORD {
	r0, _, _ := procGetMessageTime.Call()
	return DWORD(r0)
}

func GetWindowRect(h HWND) (rect RECT, err error) {
	r0, _, e1 := procGetWindowRect.Call(uintptr(h), uintptr(unsafe.Pointer(&rect)))
	if r0 == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LoadBitmapS(instance HINSTANCE, name string) (icon HBITMAP, err error) {
	return LoadBitmapW(instance, resourceNameToPtr(name))
}

func LoadBitmap(instance HINSTANCE, name *uint16) (icon HBITMAP, err error) {
	return LoadBitmapW(instance, uintptr(unsafe.Pointer(name)))
}

func LoadBitmapW(instance HINSTANCE, name uintptr) (icon HBITMAP, err error) {
	r0, _, e1 := syscall.Syscall(procLoadBitmapW.Addr(), 2, uintptr(instance), name, 0)
	icon = HBITMAP(r0)
	if icon == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

//SetTimer(
//  hWnd: HWND;               {与定时器相关联的窗口句柄}
//  nIDEvent: UINT;           {指定一个非 0 的定时器标识符}
//  uElapse: UINT;            {指定间隔时间, 单位是毫秒}
//  lpTimerFunc: TFNTimerProc {每到时间后, 要调用的函数的指针}
//): UINT;                    {返回定时器标识符; 失败返回 0}

//TimerProc(
//  hWnd: HWND;    {与定时器相关联的窗口句柄}
//  uMsg: UINT;    {WM_TIMER 消息}
//  idEvent: UINT; {定时器的标识符}
//  Time: DWORD    {以世界时间公约格式(UTC)指定的系统时间}
//);

//移除定时器函数的声明:
//KillTimer(
//  hWnd: HWND;    {与定时器相关联的窗口句柄}
//  uIDEvent: UINT {定时器标识符}
//): BOOL;

type TFNTimerProc func(hWnd uintptr, uMsg UINT, idEvent TimerEventID, Time DWORD)

func SetTimer(hWnd uintptr, nIDEvent TimerEventID, uElapse UINT, lpTimerFunc TFNTimerProc) UINT {
	var lpTimerFunc_ uintptr
	if lpTimerFunc != nil {
		lpTimerFunc_ = syscall.NewCallback(lpTimerFunc)
	}
	r0, _, _ := procSetTimer.Call(hWnd, uintptr(nIDEvent), uintptr(uElapse), lpTimerFunc_)
	return UINT(r0)
}

func KillTimer(hWnd uintptr, nIDEvent TimerEventID) bool {
	r0, _, _ := procKillTimer.Call(hWnd, uintptr(nIDEvent))
	return r0 != 0
}

//idAttach As Long, ByVal idAttachTo As Long, ByVal fAttach
func AttachThreadInput(idAttach uintptr, idAttachTo uintptr, fAttach bool) (bool, error) {
	var fAttach_ uintptr = 0
	if fAttach {
		fAttach_ = 1
	}
	r0, _, err := procAttachThreadInput.Call(idAttach, uintptr(idAttachTo), fAttach_)

	if r0 == 0 {
		log.Println("AttachThreadInput.err:", err)
	}

	return r0 != 0, err
}

func SetActiveWindow(hWnd HWND) bool {
	r0, _, _ := procSetActiveWindow.Call(uintptr(hWnd))
	return r0 != 0
}

func TrackPopupMenu(m HMENU, uFlags UINT, x, y int, hWnd HWND) (bool, *RECT) {
	rect := &RECT{}
	r0, _, _ := procTrackPopupMenu.Call(uintptr(m), uintptr(uFlags), uintptr(x), uintptr(y), uintptr(0), uintptr(hWnd), uintptr(unsafe.Pointer(rect)))
	return r0 != 0, rect
}

func GetActiveWindow() HWND {
	r0, _, _ := procGetActiveWindow.Call()
	return HWND(r0)
}

func GetForegroundWindow() HWND {
	r0, _, _ := procGetForegroundWindow.Call()
	return HWND(r0)
}

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
	classnamePtr := uintptr(0)
	if classname != "" {
		classnamePtr = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(classname)))
	}

	r0, _, e1 := syscall.Syscall12(procCreateWindowExW.Addr(), 12, uintptr(exstyle), classnamePtr, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowname))), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(wndparent), uintptr(menu), uintptr(instance), uintptr(param))
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

func MessageBoxTimeoutW(parent HWND, text, title string, boxType BoxType, wLanguageId uint16, dwMilliseconds uint32) (int, error) {
	ret, err := Syscall(procMessageBoxTimeoutW.Addr(), uintptr(parent),
		StringToUintptr(text), StringToUintptr(title), uintptr(boxType), uintptr(wLanguageId), uintptr(dwMilliseconds))
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
	var lpRectUpdatePtr = uintptr(0)
	if lpRectUpdate != nil {
		lpRectUpdatePtr = uintptr(unsafe.Pointer(lpRectUpdate))
	}
	ret, err := Syscall(procRedrawWindow.Addr(), uintptr(hwnd), lpRectUpdatePtr, uintptr(prgnUpdate), uintptr(flags))
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
