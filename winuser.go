package winapi

/*
 * MessageBox() Flags
 */

type BoxType uint

const (
	MB_OK               BoxType = 0x00000000
	MB_OKCANCEL         BoxType = 0x00000001
	MB_ABORTRETRYIGNORE BoxType = 0x00000002
	MB_YESNOCANCEL      BoxType = 0x00000003
	MB_YESNO            BoxType = 0x00000004
	MB_RETRYCANCEL      BoxType = 0x00000005
	//#if(WINVER >= 0x0500)
	MB_CANCELTRYCONTINUE BoxType = 0x00000006
	//#endif /* WINVER >= 0x0500 */

	MB_ICONHAND        BoxType = 0x00000010
	MB_ICONQUESTION    BoxType = 0x00000020
	MB_ICONEXCLAMATION BoxType = 0x00000030
	MB_ICONASTERISK    BoxType = 0x00000040

	//#if(WINVER >= 0x0400)
	MB_USERICON    BoxType = 0x00000080
	MB_ICONWARNING BoxType = MB_ICONEXCLAMATION
	MB_ICONERROR   BoxType = MB_ICONHAND
	//#endif /* WINVER >= 0x0400 */

	MB_ICONINFORMATION BoxType = MB_ICONASTERISK
	MB_ICONSTOP        BoxType = MB_ICONHAND

	MB_DEFBUTTON1 BoxType = 0x00000000
	MB_DEFBUTTON2 BoxType = 0x00000100
	MB_DEFBUTTON3 BoxType = 0x00000200
	//#if(WINVER >= 0x0400)
	MB_DEFBUTTON4 BoxType = 0x00000300
	//#endif /* WINVER >= 0x0400 */

	MB_APPLMODAL   BoxType = 0x00000000
	MB_SYSTEMMODAL BoxType = 0x00001000
	MB_TASKMODAL   BoxType = 0x00002000
	//#if(WINVER >= 0x0400)
	MB_HELP BoxType = 0x00004000 // Help Button
	//#endif /* WINVER >= 0x0400 */

	MB_NOFOCUS              BoxType = 0x00008000
	MB_SETFOREGROUND        BoxType = 0x00010000
	MB_DEFAULT_DESKTOP_ONLY BoxType = 0x00020000

	//#if(WINVER >= 0x0400)
	MB_TOPMOST    BoxType = 0x00040000
	MB_RIGHT      BoxType = 0x00080000
	MB_RTLREADING BoxType = 0x00100000

	//#endif /* WINVER >= 0x0400 */

	//#ifdef _WIN32_WINNT
	//#if (_WIN32_WINNT >= 0x0400)
	MB_SERVICE_NOTIFICATION BoxType = 0x00200000
	//#else
	//	MB_SERVICE_NOTIFICATION BoxType = 0x00040000
	//#endif
	MB_SERVICE_NOTIFICATION_NT3X BoxType = 0x00040000
	//#endif

	MB_TYPEMASK BoxType = 0x0000000F
	MB_ICONMASK BoxType = 0x000000F0
	MB_DEFMASK  BoxType = 0x00000F00
	MB_MODEMASK BoxType = 0x00003000
	MB_MISCMASK BoxType = 0x0000C000
)

type Wndclassex struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HANDLE
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HICON
}

type Wndclass struct {
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HANDLE
	MenuName   *uint16
	ClassName  *uint16
}

//    UINT        style;
//    WNDPROC     lpfnWndProc;
//    int         cbClsExtra;
//    int         cbWndExtra;
//    HINSTANCE   hInstance;
//    HICON       hIcon;
//    HCURSOR     hCursor;
//    HBRUSH      hbrBackground;
//    LPCWSTR     lpszMenuName;
//    LPCWSTR     lpszClassName;

type PAINTSTRUCT struct {
	HDC         HDC
	Erase       int32 // bool
	RcPaint     RECT
	Restore     int32 // bool
	IncUpdate   int32 // bool
	rgbReserved [32]byte
}

type Msg struct {
	Hwnd    HANDLE
	Message uint32
	Wparam  uintptr
	Lparam  uintptr
	Time    uint32
	Pt      POINT
}
