package winapi

import (
	"unsafe"
)

type IStream interface{}

type RECT struct {
	Left   LONG
	Top    LONG
	Right  LONG
	Bottom LONG
}

var RECT_SIZE = UINT(unsafe.Sizeof(RECT{}))

type RECTL RECT

var RECTL_SIZE = UINT(unsafe.Sizeof(RECTL{}))

type POINT struct {
	X uintptr
	Y uintptr
}
type POINTL POINT

type SIZE struct {
	CX LONG
	CY LONG
}

type SIZEL SIZE

//type HPALETTE__ struct {
//	unused int
//}

//type HENHMETAFILE__ struct {
//	unused int
//}

//type HPALETTE *HPALETTE__
//type HENHMETAFILE *HENHMETAFILE__

//type LOBYTE
func LOBYTE(w DWORD) BYTE { return BYTE((DWORD(w)) & 0xff) }
