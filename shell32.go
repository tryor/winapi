package winapi

import (
	//	"errors"
	"syscall"

	//	"unicode/utf16"
	"unsafe"
)

var (
	modshell32 = syscall.NewLazyDLL("SHELL32.dll")

	procShell_NotifyIconW = modshell32.NewProc("Shell_NotifyIconW")
)

const (
	NIM_ADD        = 0x00000000
	NIM_MODIFY     = 0x00000001
	NIM_DELETE     = 0x00000002
	NIM_SETFOCUS   = 0x00000003
	NIM_SETVERSION = 0x00000004
)

const (
	NIF_MESSAGE = 0x00000001
	NIF_ICON    = 0x00000002
	NIF_TIP     = 0x00000004
	NIF_STATE   = 0x00000008
	NIF_INFO    = 0x00000010
	NIF_GUID    = 0x00000020

	NIF_REALTIME = 0x00000040
	NIF_SHOWTIP  = 0x00000080
)

type NOTIFYICONDATA struct {
	CbSize           DWORD
	HWnd             HWND
	UID              UINT
	UFlags           UINT
	UCallbackMessage UINT
	HIcon            HICON

	SzTip       [128]uint16 //WCHAR
	DwState     DWORD
	DwStateMask DWORD
	SzInfo      [256]uint16 //WCHAR

	UVersion UINT

	SzInfoTitle [64]uint16 //WCHAR
	DwInfoFlags DWORD

	GuidItem     GUID
	HBalloonIcon HICON
}

/*
   nid.cbSize = sizeof(NOTIFYICONDATA);
   nid.hWnd = hWnd;;
   nid.uID = 111; // ID_TASKBARICON; //发出的消息中的wParam参数
   nid.uFlags = NIF_ICON | NIF_MESSAGE | NIF_TIP;
   nid.uCallbackMessage = 222; // WM_ICONNOTIFY; //点击托盘图标系统发出的消息（即发出的消息中的lParam参数）
   nid.hIcon = LoadIcon(hInstance, MAKEINTRESOURCE(IDI_WINDOWSPROJECT1));
   lstrcpy(nid.szTip, _T("收银吧"));
*/

//func Shell_NotifyIcon(dwMessage uint32, hWnd HWND, uID UINT, uFlags UINT, uCallbackMessage UINT, hIcon HICON, caption string) (ret bool, err error) {
//	var nid NOTIFYICONDATA
//	nid.cbSize = DWORD(unsafe.Sizeof(nid))
//	nid.hWnd = hWnd
//	nid.uID = uID
//	nid.uFlags = uFlags
//	nid.uCallbackMessage = uCallbackMessage
//	nid.hIcon = hIcon
//	szTip := syscall.StringToUTF16(caption)
//	if len(szTip) > 128 {
//		szTip = szTip[0:128]
//	}
//	copy(nid.szTip[:], szTip)
//	return Shell_NotifyIconW(dwMessage, &nid)
//}

func Shell_NotifyIconW(dwMessage uint32, lpData *NOTIFYICONDATA) (ret bool, err error) {
	r, e1 := Syscall(procShell_NotifyIconW.Addr(), uintptr(dwMessage), uintptr(unsafe.Pointer(lpData)))
	if e1 != nil {
		err = e1
	}
	ret = r != 0
	return
}
