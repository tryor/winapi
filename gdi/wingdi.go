package gdi

import (
	"unsafe"

	. "github.com/tryor/winapi"
)

const (
	BI_RGB       = 0
	BI_BITFIELDS = 3

	DIB_PAL_COLORS = 1
	DIB_RGB_COLORS = 0

	BLACKNESS   = 0x42
	DSTINVERT   = 0x550009
	MERGECOPY   = 0xC000CA
	MERGEPAINT  = 0xBB0226
	NOTSRCCOPY  = 0x330008
	NOTSRCERASE = 0x1100A6
	PATCOPY     = 0xF00021
	PATINVERT   = 0x5A0049
	PATPAINT    = 0xFB0A09
	SRCAND      = 0x8800C6
	SRCCOPY     = 0xCC0020
	SRCERASE    = 0x440328
	SRCINVERT   = 0x660046
	SRCPAINT    = 0xEE0086
	WHITENESS   = 0xFF0062
)

type BITMAP struct {
	Type       int32
	Width      int32
	Height     int32
	WidthBytes int32
	Planes     uint16
	BitsPixel  uint16
	Bits       *byte
}

type BITMAPINFOHEADER struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}

type BITMAPINFO struct {
	Header BITMAPINFOHEADER
	Colors [1]RGBQUAD
}

type RGBQUAD struct {
	Blue     byte
	Green    byte
	Red      byte
	Reserved byte
}

type METAHEADER struct {
	mtType         WORD
	mtHeaderSize   WORD
	mtVersion      WORD
	mtSize         DWORD
	mtNoObjects    WORD
	mtMaxRecord    DWORD
	mtNoParameters DWORD
}

var METAHEADER_SIZE = UINT(unsafe.Sizeof(METAHEADER{}))

func GetRValue(rgb COLORREF) BYTE { return LOBYTE(DWORD(rgb)) }
func GetGValue(rgb COLORREF) BYTE { return LOBYTE(DWORD(rgb) >> 8) }
func GetBValue(rgb COLORREF) BYTE { return LOBYTE(DWORD(rgb) >> 16) }

func RGB(r, g, b BYTE) COLORREF {
	return ((COLORREF)((DWORD(r) | DWORD(WORD(g)<<8)) | ((DWORD(b)) << 16)))
}

const LF_FACESIZE = 32

type LOGFONTA struct {
	LfHeight         LONG
	LfWidth          LONG
	LfEscapement     LONG
	LfOrientation    LONG
	LfWeight         LONG
	LfItalic         BYTE
	LfUnderline      BYTE
	LfStrikeOut      BYTE
	LfCharSet        BYTE
	LfOutPrecision   BYTE
	LfClipPrecision  BYTE
	LfQuality        BYTE
	LfPitchAndFamily BYTE
	LfFaceName       [LF_FACESIZE]CHAR
}

var LOGFONTA_SIZE = UINT(unsafe.Sizeof(LOGFONTA{}))

type LOGFONTW struct {
	lfHeight         LONG
	lfWidth          LONG
	LfEscapement     LONG
	LfOrientation    LONG
	LfWeight         LONG
	LfItalic         BYTE
	LfUnderline      BYTE
	LfStrikeOut      BYTE
	LfCharSet        BYTE
	LfOutPrecision   BYTE
	LfClipPrecision  BYTE
	LfQuality        BYTE
	LfPitchAndFamily BYTE
	LfFaceName       [LF_FACESIZE]WCHAR
}

var LOGFONTW_SIZE = UINT(unsafe.Sizeof(LOGFONTW{}))
