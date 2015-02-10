package winapi

import (
	"unsafe"
)

type Wndclassex struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HANDLE
	Icon       HANDLE
	Cursor     HANDLE
	Background HANDLE
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HANDLE
}

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

const (
	GWL_EXSTYLE    = -20 //设定一个新的扩展风格。
	GWL_HINSTANCE  = -6  //设置一个新的应用程序实例句柄。
	GWL_ID         = -12 //设置一个新的窗口标识符。
	GWL_STYLE      = -16 //设定一个新的窗口风格。
	GWL_USERDATA   = -21 //设置与窗口有关的32位值。每个窗口均有一个由创建该窗口的应用程序使用的32位值。
	GWL_WNDPROC    = -4  //为窗口过程设定一个新的地址。
	GWL_HWNDPARENT = -8  //改变子窗口的父窗口,应使用SetParent函数。
)

const (
	// Window styles
	WS_OVERLAPPED   = 0
	WS_POPUP        = 0x80000000
	WS_CHILD        = 0x40000000
	WS_MINIMIZE     = 0x20000000
	WS_VISIBLE      = 0x10000000
	WS_DISABLED     = 0x8000000
	WS_CLIPSIBLINGS = 0x4000000
	WS_CLIPCHILDREN = 0x2000000
	WS_MAXIMIZE     = 0x1000000
	WS_CAPTION      = WS_BORDER | WS_DLGFRAME
	WS_BORDER       = 0x800000
	WS_DLGFRAME     = 0x400000
	WS_VSCROLL      = 0x200000
	WS_HSCROLL      = 0x100000
	WS_SYSMENU      = 0x80000
	WS_THICKFRAME   = 0x40000
	WS_GROUP        = 0x20000
	WS_TABSTOP      = 0x10000
	WS_MINIMIZEBOX  = 0x20000
	WS_MAXIMIZEBOX  = 0x10000
	WS_TILED        = WS_OVERLAPPED
	WS_ICONIC       = WS_MINIMIZE
	WS_SIZEBOX      = WS_THICKFRAME
	// Common Window Styles
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_TILEDWINDOW      = WS_OVERLAPPEDWINDOW
	WS_POPUPWINDOW      = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW      = WS_CHILD
	//扩展样式
	WS_EX_ACCEPTFILES   = 0x00000010 //指明了一个已创建视窗具有拖拽文件功能(指定以该风格创建的窗口接受一个拖拽文件)
	WS_EX_APPWINDOW     = 0x00040000 //强制一个可见的顶级视窗到工具栏上(当窗口可见时，将一个顶层窗口放置到任务条上)
	WS_EX_CLIENTEDGE    = 0x00000200 //使一个视窗具有凹陷边框(指定窗口有一个带阴影的边界)
	WS_EX_COMPOSITED    = 0x02000000 //Windows XP:将一个窗体的所有子窗口使用双缓冲按照从低到高方式绘制出来,参阅remark项.如果这个视窗已经使用经典样式中的下列值CS_OWNDC , CS_CLASSDC,WS_EX_CONTEXTHELP.此参数将不能使用.这个样式的视窗在标题栏上有一个问号,当拥护点击着个问号,鼠标变成一个问号,如果用户然后点击一个子窗口,子窗就会收到一条WM_HELP消息.子窗口将把这个消息传递给他的父进程,这个父进程将用HELP_WM_HELP命令调用WinHelp函数.这个帮助程序常常弹出一个典型的包含其子窗口的帮助的窗口本参数不能和WS_MAXIMIZEBOX ,WS_MINIMIZEBOX一起使用.
	WS_EX_CONTEXTHELP   = 0x00000400 //在窗口的标题条包含一个问号标志。当用户点击了问号时，鼠标光标变为一个问号的指针、如果点击了一个子窗口，则子窗口接收到WM_HELP消息。子窗口应该将这个消息传递给父窗口过程，父窗口再通过HELP_WM_HELP命令调用WinHelp函数。这个Help应用程序显示一个包含子窗口帮助信息的弹出式窗口。WS_EX_CONTEXTHELP不能与WS_MAXIMIZEBOX和WS_MINIMIZEBOX同时使用。
	WS_EX_CONTROLPARENT = 0x00010000 //这个窗体本身包含了参与对话框导航的子窗口.如果使用了这个参数,对话框管理器进入这个窗体的子窗口,当执行导航操作时,比如按住TAB键,方向键.(允许用户使用Tab键在窗口的子窗口间搜索)
	WS_EX_DLGMODALFRAME = 0x00000001 //创建一个具有双边框的窗口,这个窗口可以通过使用WS_CAPTION样式被创建成具有一个标题栏的窗口.（创建一个带双边的窗口；该窗口可以在dwStyle中指定WS_CAPTION风格来创建一个标题栏。）
	WS_EX_LAYERED       = 0x00080000 //Windows 2000/XP:创建一个分层的窗口.注意,这不能用在子窗口上.同样,如果窗口具有CS_OWNDC ,CS_CLASSDC样式,这也不用使用.
	WS_EX_LAYOUTRTL     = 0x00400000 //阿拉伯以及西伯来版本的98/ME,2000/XP创建一个水平起点在右边的窗口.越往左边水平坐标值变大.
	WS_EX_LEFT          = 0x00000000 //创建一个窗口具有一般的左对齐属性.此为默认（窗口具有左对齐属性，这是缺省设置的）
	WS_EX_LEFTSCROLLBAR = 0x00004000 //如果外壳语言是西伯来,阿拉伯,或者其他阅读顺序的语言,竖滚动条将会在客户区的左边.对其他语言,此参数忽略.（如果外壳语言是如Hebrew，Arabic，或其他支持reading order alignment的语言，则标题条（如果存在）则在客户区的左部分。若是其他语言，在该风格被忽略并且不作为错误处理）
	WS_EX_LTRREADING    = 0x00000000 //窗体的文字按照从左到右排列.此为默认.（窗口文本以LEFT到RIGHT（自左向右）属性的顺序显示。这是缺省设置的）
	WS_EX_MDICHILD      = 0x00000040 //创建一个多文档界面的子窗口.（创建一个MDI子窗口）
	WS_EX_NOACTIVATE    = 0x08000000 //Windows 2000/XP:一个使用此参数创建的顶级窗口不会变成前台窗口,当用户点击他时.系统不会将此窗口放到前台,当用户最小化或者关闭这个前台窗口. 要激活这样的窗口,使用SetActiveWindow或者SetForegroundWindow函数此类型的窗口默认不会显示在任务栏上.要强行将这样的窗口显示到任务栏上,使用WS_EX_APPWINDOW参数.
	// WS_EX_NODRAG = //防止窗口被移动
	WS_EX_NOINHERITLAYOUT  = 0x00100000                                            //
	WS_EX_NOPARENTNOTIFY   = 0x00000004                                            //指明一个使用此参数创建的窗口不发送WM_PARENTNOTIFY消息给他的Windows 2000/XP:用此参数创建的窗体不会传递他的窗口布局给他的子窗父窗口当这个窗口被创建或者销毁的时候.（指明以这个风格创建的窗口在被创建和销毁时不向父窗口发送WM_PARENTNOTFY消息）
	WS_EX_OVERLAPPEDWINDOW = (WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE)                 //联合了WS_EX_CLIENTEDGE and WS_EX_WINDOWEDGE styles（WS_EX_CLIENTEDGE和WS_EX_WINDOWEDGE的组合）
	WS_EX_PALETTEWINDOW    = (WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST) //联合了WS_EX_WINDOWEDGE, WS_EX_TOOLWINDOW, and WS_EX_TOPMOST styles.（WS_EX_WINDOWEDGE, WS_EX_TOOLWINDOW和WS_WX_TOPMOST风格的组合WS_EX_RIGHT:窗口具有普通的右对齐属性，这依赖于窗口类。只有在外壳语言是如Hebrew,Arabic或其他支持读顺序对齐（reading order alignment）的语言时该风格才有效，否则，忽略该标志并且不作为错误处理）
	WS_EX_RIGHT            = 0x00001000                                            //窗口具有一般的右对齐属性.这要依靠这个窗口的类.这个样式只有外壳语言是西伯来语,阿拉伯语等其他阅读顺序的语言才有影响,否则此样式别忽略对文字标签或者编辑框使用WS_EX_RIGHT样式跟使用SS_RIGHT 或者 ES_RIGHT影响是一样的.对按钮使用这个样式跟使用BS_RIGHT 和BS_RIGHTBUTTON的影响是一样的
	WS_EX_RIGHTSCROLLBAR   = 0x00000000                                            //竖直滚动条显示在客户区的右边.默认.（垂直滚动条在窗口的右边界。这是缺省设置的）
	WS_EX_RTLREADING       = 0x00002000                                            //如果外壳语言是西伯来语,阿拉伯语等支持排列方式阅读的语言,窗体文字将按照从右到左的阅读顺序.对其他语言,此样式忽略.（如果外壳语言是如Hebrew，Arabic，或其他支持读顺序对齐（reading order alignment）的语言，则窗口文本是一自左向右）RIGHT到LEFT顺序的读出顺序。若是其他语言，在该风格被忽略并且不作为错误处理）
	WS_EX_STATICEDGE       = 0x00020000                                            //创建一个窗口具有三维边框用来表示一个项目不接受用户输入（为不接受用户输入的项创建一个3维边界风格）
	WS_EX_TOOLWINDOW       = 0x00000080                                            //创建一个工具窗口:也就是说,这个窗口被用来做浮动工具条.一个工具窗口具有一个比一般的标题栏短的标题栏,并且系统在标题栏使用小字体.作为工具窗口,它不显示在工具栏上.当用户用ALT+TAB切换时也不出现在对话框中.如果一个工具窗有系统菜单,那么他的图标不会被显示在标题栏上.但是,你可以通过键入ALT+TAB或者右键点击标题栏来显示系统菜单.（创建工具窗口，即窗口是一个游动的工具条。工具窗口的标题条比一般窗口的标题条短，并且窗口标题以小字体显示。工具窗口不在任务栏里显示，当用户按下Alt+Tab键时工具窗口不在对话框里显示。如果工具窗口有一个系统菜单，它的图标也不会显示在标题栏里，但是，可以通过点击鼠标右键或Alt+Space来显示菜单）
	WS_EX_TOPMOST          = 0x00000008                                            //指明用此参数创建的窗口将会放在所有顶级视窗上并且停在最上面.即使这个窗口不是活动的.要添加或者移除他,使用SetWindowPos函数.（指明以该风格创建的窗口应放置在所有非最高层窗口的上面并且停留在其L，即使窗口未被激活。使用函数SetWindowPos来设置和移去这个风格）
	WS_EX_TRANSPARENT      = 0x00000020                                            //用此参数创建的的窗口在他同一线程的窗口被绘制前将不会被绘制.这个窗口透明的显示,因为同一线程的窗口已经绘制出来要脱离这个限制激活透明,使用SetWindowRgn函数.（指定以这个风格创建的窗口在窗口下的同属窗口已重画时，该窗口才可以重画。由于其下的同属窗口已被重画，该窗口是透明的）
	WS_EX_WINDOWEDGE       = 0x00000100                                            //使一个窗口具有凸起的边框.

	//窗体风格
	//WS_EX_TOPMOST = 0x00000008 //总在顶层的窗口
	//WS_EX_ACCEPTFILES      = 0x00000010 //允许窗口进行鼠标拖放操作
	//WS_EX_TOOLWINDOW = 0x00000080 //工具窗口（很窄的标题栏）
	//WS_EX_WINDOWEDGE = 0x00000100 //立体感的边框
	//WS_EX_CLIENTEDGE       = 0x00000200 //客户区立体边框
	//WS_EX_OVERLAPPEDWINDOW = WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE
	//WS_EX_PALETTEWINDOW = WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST

	// Some windows messages
	WM_CREATE  = 1
	WM_DESTROY = 2
	WM_CLOSE   = 16
	WM_COMMAND = 273

	// Some button control styles
	BS_DEFPUSHBUTTON = 1

	// Some color constants
	COLOR_WINDOW  = 5
	COLOR_BTNFACE = 15

	// Default window position
	CW_USEDEFAULT = 0x80000000 - 0x100000000

	// Show window default style
	SW_HIDE            = 0  //{隐藏, 并且任务栏也没有最小化图标}
	SW_SHOWNORMAL      = 1  //{用最近的大小和位置显示, 激活}
	SW_NORMAL          = 1  //{同 SW_SHOWNORMAL}
	SW_SHOWMINIMIZED   = 2  //{最小化, 激活}
	SW_SHOWMAXIMIZED   = 3  //{最大化, 激活}
	SW_MAXIMIZE        = 3  //{同 SW_SHOWMAXIMIZED}
	SW_SHOWNOACTIVATE  = 4  //{用最近的大小和位置显示, 不激活}
	SW_SHOW            = 5  //{同 SW_SHOWNORMAL}
	SW_MINIMIZE        = 6  //{最小化, 不激活}
	SW_SHOWMINNOACTIVE = 7  //{同 SW_MINIMIZE}
	SW_SHOWNA          = 8  //{同 SW_SHOWNOACTIVATE}
	SW_RESTORE         = 9  //{同 SW_SHOWNORMAL}
	SW_SHOWDEFAULT     = 10 //{同 SW_SHOWNORMAL}
	SW_MAX             = 10 //{同 SW_SHOWNORMAL}

)

var (
	// Some globally known cursors
	IDC_ARROW = MakeIntResource(32512)
	IDC_IBEAM = MakeIntResource(32513)
	IDC_WAIT  = MakeIntResource(32514)
	IDC_CROSS = MakeIntResource(32515)

	IDC_SIZENS   = MakeIntResource(32645)
	IDC_SIZEWE   = MakeIntResource(32644)
	IDC_SIZENWSE = MakeIntResource(32642)
	IDC_SIZENESW = MakeIntResource(32643)
	IDC_SIZEALL  = MakeIntResource(32646)

	//// Some globally known cursors
	//IDC_ARROW = MakeIntResource(32512)
	//IDC_IBEAM = MakeIntResource(32513)
	//IDC_WAIT  = MakeIntResource(32514)
	//IDC_CROSS = MakeIntResource(32515)

	// Some globally known icons
	IDI_APPLICATION = MakeIntResource(32512)
	IDI_HAND        = MakeIntResource(32513)
	IDI_QUESTION    = MakeIntResource(32514)
	IDI_EXCLAMATION = MakeIntResource(32515)
	IDI_ASTERISK    = MakeIntResource(32516)
	IDI_WINLOGO     = MakeIntResource(32517)
	IDI_WARNING     = IDI_EXCLAMATION
	IDI_ERROR       = IDI_HAND
	IDI_INFORMATION = IDI_ASTERISK
)

const (
	WM_PAINT         = 0x000F
	WM_MOUSEFIRST    = 0x0200
	WM_MOUSEMOVE     = 0x0200
	WM_LBUTTONDOWN   = 0x0201
	WM_LBUTTONUP     = 0x0202
	WM_LBUTTONDBLCLK = 0x0203
	WM_RBUTTONDOWN   = 0x0204
	WM_RBUTTONUP     = 0x0205
	WM_RBUTTONDBLCLK = 0x0206
	WM_MBUTTONDOWN   = 0x0207
	WM_MBUTTONUP     = 0x0208

	WM_SETCURSOR = 0x0020

	WM_KEYDOWN = 0x0100
	WM_KEYUP   = 0x0101
	WM_CHAR    = 0x0102 //按下某键，并已发出WM_KEYDOWN， WM_KEYUP消息
)

func LOWORD(l uintptr) int {
	return (int(l) & 0xffff)
}

func HIWORD(h uintptr) int {
	return ((int(h) >> 16) & 0xffff)
}

func KeyState(key byte) bool { return (key & 0x80) > 0 }

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}
