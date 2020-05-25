package winapi

type (
	HANDLE    uintptr
	HWND      HANDLE
	HMENU     HANDLE
	HMODULE   HANDLE
	HINSTANCE HANDLE
	HDC       HANDLE
	HRGN      HANDLE
	HBRUSH    HANDLE
	HICON     HANDLE
	HCURSOR   HANDLE
	HPEN      HANDLE
	HPALETTE  HANDLE
	HBITMAP   HANDLE
	HFONT     HANDLE

	WPARAM   uintptr
	LPARAM   uintptr
	LRESULT  int
	COLORREF uint32
	LANGID   uint16

	PROPID       ULONG
	HMETAFILE    uintptr
	HENHMETAFILE uintptr

	UINT32    uint32
	ULONG_PTR uintptr
	UINT_PTR  uintptr
	DWORD_PTR uintptr

	BOOL     bool
	CHAR     int8
	BYTE     uint8
	SHORT    int16
	INT16    int16
	WORD     uint16
	UINT16   uint16
	INT      int32 //, int
	UINT     uint32
	LONG     int32
	ULONG    uint32
	DWORD    uint32
	LONGLONG int64
	DWORD64  uint64
	FLOAT    float32
	DOUBLE   float64
	WCHAR    uint16
	LPSTR    *byte
	LPCSTR   *byte //syscall.StringBytePtr()
	LPWSTR   *uint16
	LPCWSTR  *uint16 //syscall.StringToUTF16Ptr()
	LPBYTE   *BYTE

	TimerEventID UINT
)

const (
	TRUE  BOOL = true
	FALSE BOOL = false
)

//  Primary language IDs.
//
//  WARNING: These aren't always unique.  Bosnian, Serbian & Croation for example.
//
//  It is recommended that applications test for locale names or actual LCIDs.
//
//  Note that the LANG, SUBLANG construction is not always consistent.
//  The named locale APIs (eg GetLocaleInfoEx) are recommended.
//
const (
	LANG_NEUTRAL             LANGID = 0x00
	LANG_INVARIANT           LANGID = 0x7f
	LANG_AFRIKAANS           LANGID = 0x36
	LANG_ALBANIAN            LANGID = 0x1c
	LANG_ALSATIAN            LANGID = 0x84
	LANG_AMHARIC             LANGID = 0x5e
	LANG_ARABIC              LANGID = 0x01
	LANG_ARMENIAN            LANGID = 0x2b
	LANG_ASSAMESE            LANGID = 0x4d
	LANG_AZERI               LANGID = 0x2c
	LANG_BASHKIR             LANGID = 0x6d
	LANG_BASQUE              LANGID = 0x2d
	LANG_BELARUSIAN          LANGID = 0x23
	LANG_BENGALI             LANGID = 0x45
	LANG_BRETON              LANGID = 0x7e
	LANG_BOSNIAN             LANGID = 0x1a   // Use with SUBLANG_BOSNIAN_* Sublanguage IDs
	LANG_BOSNIAN_NEUTRAL     LANGID = 0x781a // Use with the ConvertDefaultLocale function
	LANG_BULGARIAN           LANGID = 0x02
	LANG_CATALAN             LANGID = 0x03
	LANG_CHINESE             LANGID = 0x04   // Use with SUBLANG_CHINESE_* Sublanguage IDs
	LANG_CHINESE_SIMPLIFIED  LANGID = 0x04   // Use with the ConvertDefaultLocale function
	LANG_CHINESE_TRADITIONAL LANGID = 0x7c04 // Use with the ConvertDefaultLocale function
	LANG_CORSICAN            LANGID = 0x83
	LANG_CROATIAN            LANGID = 0x1a
	LANG_CZECH               LANGID = 0x05
	LANG_DANISH              LANGID = 0x06
	LANG_DARI                LANGID = 0x8c
	LANG_DIVEHI              LANGID = 0x65
	LANG_DUTCH               LANGID = 0x13
	LANG_ENGLISH             LANGID = 0x09
	LANG_ESTONIAN            LANGID = 0x25
	LANG_FAEROESE            LANGID = 0x38
	LANG_FARSI               LANGID = 0x29 // Deprecated: use LANG_PERSIAN instead
	LANG_FILIPINO            LANGID = 0x64
	LANG_FINNISH             LANGID = 0x0b
	LANG_FRENCH              LANGID = 0x0c
	LANG_FRISIAN             LANGID = 0x62
	LANG_GALICIAN            LANGID = 0x56
	LANG_GEORGIAN            LANGID = 0x37
	LANG_GERMAN              LANGID = 0x07
	LANG_GREEK               LANGID = 0x08
	LANG_GREENLANDIC         LANGID = 0x6f
	LANG_GUJARATI            LANGID = 0x47
	LANG_HAUSA               LANGID = 0x68
	LANG_HEBREW              LANGID = 0x0d
	LANG_HINDI               LANGID = 0x39
	LANG_HUNGARIAN           LANGID = 0x0e
	LANG_ICELANDIC           LANGID = 0x0f
	LANG_IGBO                LANGID = 0x70
	LANG_INDONESIAN          LANGID = 0x21
	LANG_INUKTITUT           LANGID = 0x5d
	LANG_IRISH               LANGID = 0x3c // Use with the SUBLANG_IRISH_IRELAND Sublanguage ID
	LANG_ITALIAN             LANGID = 0x10
	LANG_JAPANESE            LANGID = 0x11
	LANG_KANNADA             LANGID = 0x4b
	LANG_KASHMIRI            LANGID = 0x60
	LANG_KAZAK               LANGID = 0x3f
	LANG_KHMER               LANGID = 0x53
	LANG_KICHE               LANGID = 0x86
	LANG_KINYARWANDA         LANGID = 0x87
	LANG_KONKANI             LANGID = 0x57
	LANG_KOREAN              LANGID = 0x12
	LANG_KYRGYZ              LANGID = 0x40
	LANG_LAO                 LANGID = 0x54
	LANG_LATVIAN             LANGID = 0x26
	LANG_LITHUANIAN          LANGID = 0x27
	LANG_LOWER_SORBIAN       LANGID = 0x2e
	LANG_LUXEMBOURGISH       LANGID = 0x6e
	LANG_MACEDONIAN          LANGID = 0x2f // the Former Yugoslav Republic of Macedonia
	LANG_MALAY               LANGID = 0x3e
	LANG_MALAYALAM           LANGID = 0x4c
	LANG_MALTESE             LANGID = 0x3a
	LANG_MANIPURI            LANGID = 0x58
	LANG_MAORI               LANGID = 0x81
	LANG_MAPUDUNGUN          LANGID = 0x7a
	LANG_MARATHI             LANGID = 0x4e
	LANG_MOHAWK              LANGID = 0x7c
	LANG_MONGOLIAN           LANGID = 0x50
	LANG_NEPALI              LANGID = 0x61
	LANG_NORWEGIAN           LANGID = 0x14
	LANG_OCCITAN             LANGID = 0x82
	LANG_ORIYA               LANGID = 0x48
	LANG_PASHTO              LANGID = 0x63
	LANG_PERSIAN             LANGID = 0x29
	LANG_POLISH              LANGID = 0x15
	LANG_PORTUGUESE          LANGID = 0x16
	LANG_PUNJABI             LANGID = 0x46
	LANG_QUECHUA             LANGID = 0x6b
	LANG_ROMANIAN            LANGID = 0x18
	LANG_ROMANSH             LANGID = 0x17
	LANG_RUSSIAN             LANGID = 0x19
	LANG_SAMI                LANGID = 0x3b
	LANG_SANSKRIT            LANGID = 0x4f
	LANG_SCOTTISH_GAELIC     LANGID = 0x91
	LANG_SERBIAN             LANGID = 0x1a   // Use with the SUBLANG_SERBIAN_* Sublanguage IDs
	LANG_SERBIAN_NEUTRAL     LANGID = 0x7c1a // Use with the ConvertDefaultLocale function
	LANG_SINDHI              LANGID = 0x59
	LANG_SINHALESE           LANGID = 0x5b
	LANG_SLOVAK              LANGID = 0x1b
	LANG_SLOVENIAN           LANGID = 0x24
	LANG_SOTHO               LANGID = 0x6c
	LANG_SPANISH             LANGID = 0x0a
	LANG_SWAHILI             LANGID = 0x41
	LANG_SWEDISH             LANGID = 0x1d
	LANG_SYRIAC              LANGID = 0x5a
	LANG_TAJIK               LANGID = 0x28
	LANG_TAMAZIGHT           LANGID = 0x5f
	LANG_TAMIL               LANGID = 0x49
	LANG_TATAR               LANGID = 0x44
	LANG_TELUGU              LANGID = 0x4a
	LANG_THAI                LANGID = 0x1e
	LANG_TIBETAN             LANGID = 0x51
	LANG_TIGRIGNA            LANGID = 0x73
	LANG_TSWANA              LANGID = 0x32
	LANG_TURKISH             LANGID = 0x1f
	LANG_TURKMEN             LANGID = 0x42
	LANG_UIGHUR              LANGID = 0x80
	LANG_UKRAINIAN           LANGID = 0x22
	LANG_UPPER_SORBIAN       LANGID = 0x2e
	LANG_URDU                LANGID = 0x20
	LANG_UZBEK               LANGID = 0x43
	LANG_VIETNAMESE          LANGID = 0x2a
	LANG_WELSH               LANGID = 0x52
	LANG_WOLOF               LANGID = 0x88
	LANG_XHOSA               LANGID = 0x34
	LANG_YAKUT               LANGID = 0x85
	LANG_YI                  LANGID = 0x78
	LANG_YORUBA              LANGID = 0x6a
	LANG_ZULU                LANGID = 0x35
)

/*
 * RedrawWindow() flags
 */
const (
	RDW_INVALIDATE    = 0x0001
	RDW_INTERNALPAINT = 0x0002
	RDW_ERASE         = 0x0004

	RDW_VALIDATE        = 0x0008
	RDW_NOINTERNALPAINT = 0x0010
	RDW_NOERASE         = 0x0020

	RDW_NOCHILDREN  = 0x0040
	RDW_ALLCHILDREN = 0x0080

	RDW_UPDATENOW = 0x0100
	RDW_ERASENOW  = 0x0200

	RDW_FRAME   = 0x0400
	RDW_NOFRAME = 0x0800
)

const (
	LWA_COLORKEY = 0x00000001
	LWA_ALPHA    = 0x00000002
)

const (
	/*
	 * Menu flags for Add/Check/EnableMenuItem()
	 */
	MF_INSERT = 0x00000000
	MF_CHANGE = 0x00000080
	MF_APPEND = 0x00000100
	MF_DELETE = 0x00000200
	MF_REMOVE = 0x00001000

	MF_BYCOMMAND  = 0x00000000
	MF_BYPOSITION = 0x00000400

	MF_SEPARATOR = 0x00000800

	MF_ENABLED  = 0x00000000
	MF_GRAYED   = 0x00000001
	MF_DISABLED = 0x00000002

	MF_UNCHECKED       = 0x00000000
	MF_CHECKED         = 0x00000008
	MF_USECHECKBITMAPS = 0x00000200

	MF_STRING    = 0x00000000
	MF_BITMAP    = 0x00000004
	MF_OWNERDRAW = 0x00000100

	MF_POPUP        = 0x00000010
	MF_MENUBARBREAK = 0x00000020
	MF_MENUBREAK    = 0x00000040

	MF_UNHILITE = 0x00000000
	MF_HILITE   = 0x00000080

	MF_DEFAULT = 0x00001000

	MF_SYSMENU = 0x00002000
	MF_HELP    = 0x00004000
	//MF_RIGHTJUSTIFY     0x00004000
	MF_MOUSESELECT = 0x00008000
)

const (
	/*
	 * Flags for TrackPopupMenu
	 */
	TPM_LEFTBUTTON  = 0x0000
	TPM_RIGHTBUTTON = 0x0002
	TPM_LEFTALIGN   = 0x0000
	TPM_CENTERALIGN = 0x0004
	TPM_RIGHTALIGN  = 0x0008

	TPM_TOPALIGN     = 0x0000
	TPM_VCENTERALIGN = 0x0010
	TPM_BOTTOMALIGN  = 0x0020

	TPM_HORIZONTAL = 0x0000 /* Horz alignment matters more */
	TPM_VERTICAL   = 0x0040 /* Vert alignment matters more */
	TPM_NONOTIFY   = 0x0080 /* Don't send any notification msgs */
	TPM_RETURNCMD  = 0x0100

	TPM_RECURSE         = 0x0001
	TPM_HORPOSANIMATION = 0x0400
	TPM_HORNEGANIMATION = 0x0800
	TPM_VERPOSANIMATION = 0x1000
	TPM_VERNEGANIMATION = 0x2000

	TPM_NOANIMATION = 0x4000

	TPM_LAYOUTRTL = 0x8000

	TPM_WORKAREA = 0x10000
)

type IMAGE_TYPE uint8

const (
	IMAGE_BITMAP IMAGE_TYPE = 0
	IMAGE_ICON              = 1
	IMAGE_CURSOR            = 2
)

const (
	LR_DEFAULTCOLOR     = 0x00000000
	LR_MONOCHROME       = 0x00000001
	LR_COLOR            = 0x00000002
	LR_COPYRETURNORG    = 0x00000004
	LR_COPYDELETEORG    = 0x00000008
	LR_LOADFROMFILE     = 0x00000010
	LR_LOADTRANSPARENT  = 0x00000020
	LR_DEFAULTSIZE      = 0x00000040
	LR_VGACOLOR         = 0x00000080
	LR_LOADMAP3DCOLORS  = 0x00001000
	LR_CREATEDIBSECTION = 0x00002000
	LR_COPYFROMRESOURCE = 0x00004000
	LR_SHARED           = 0x00008000
)

type MouseEventFlags INT

const (
	MOUSEEVENTF_MOVE       MouseEventFlags = 0x0001 /* mouse move */
	MOUSEEVENTF_LEFTDOWN                   = 0x0002 /* left button down */
	MOUSEEVENTF_LEFTUP                     = 0x0004 /* left button up */
	MOUSEEVENTF_RIGHTDOWN                  = 0x0008 /* right button down */
	MOUSEEVENTF_RIGHTUP                    = 0x0010 /* right button up */
	MOUSEEVENTF_MIDDLEDOWN                 = 0x0020 /* middle button down */
	MOUSEEVENTF_MIDDLEUP                   = 0x0040 /* middle button up */
	MOUSEEVENTF_XDOWN                      = 0x0080 /* x button down */
	MOUSEEVENTF_XUP                        = 0x0100 /* x button down */
	MOUSEEVENTF_WHEEL                      = 0x0800 /* wheel button rolled */
	//#if (_WIN32_WINNT >= 0x0600)
	MOUSEEVENTF_HWHEEL = 0x01000 /* hwheel button rolled */
	//#endif
	//#if(WINVER >= 0x0600)
	MOUSEEVENTF_MOVE_NOCOALESCE = 0x2000 /* do not coalesce mouse moves */
	//#endif /* WINVER >= 0x0600 */
	MOUSEEVENTF_VIRTUALDESK = 0x4000 /* map to entire virtual desktop */
	MOUSEEVENTF_ABSOLUTE    = 0x8000 /* absolute move */

)
