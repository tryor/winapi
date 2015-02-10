package winapi

type GUID struct {
	Data1 ULONG
	Data2 WORD
	Data3 WORD
	Data4 [8]BYTE
}

type CLSID GUID

func NewGUID(d1 ULONG, d2, d3 WORD, d40, d41, d42, d43, d44, d45, d46, d47 BYTE) *GUID {
	return &GUID{d1, d2, d3, [8]BYTE{d40, d41, d42, d43, d44, d45, d46, d47}}
}
