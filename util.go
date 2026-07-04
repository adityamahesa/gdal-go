package gdal

/*
#include <stdio.h>
#include "cpl_string.h" // TODO: implement cpl_string.go
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// cFile opens filename with the given mode (e.g. "w") and returns the C FILE*
// together with a close func. The caller must call close when done.
func cFile(filename, mode string) (fp *C.FILE, closeFn func(), err error) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	cMode := C.CString(mode)
	defer C.free(unsafe.Pointer(cMode))
	fp = C.fopen(cName, cMode)
	if fp == nil {
		return nil, func() {}, fmt.Errorf("gdal: could not open %q with mode %q", filename, mode)
	}
	return fp, func() { C.fclose(fp) }, nil
}

// cStrings converts a Go string slice to a NULL-terminated C string array
// (CSLConstList). The returned free func releases the allocated C strings.
func cStrings(list []string) (arr **C.char, free func()) {
	if len(list) == 0 {
		return nil, func() {}
	}
	raw := make([]*C.char, len(list)+1)
	for i, s := range list {
		raw[i] = C.CString(s)
	}
	free = func() {
		for _, p := range raw {
			if p != nil {
				C.free(unsafe.Pointer(p))
			}
		}
	}
	return &raw[0], free
}

// cBytes returns a C void* pointing at the start of a byte slice, or nil for an
// empty slice. The pointer is only valid while the slice is not moved/collected,
// so it must only be passed to C for the duration of a single call.
func cBytes(data []byte) unsafe.Pointer {
	if len(data) == 0 {
		return nil
	}
	return unsafe.Pointer(&data[0])
}

// cInts converts a Go int slice to a C int array, returning a pointer to the
// first element (nil for empty). The backing array stays alive while the
// returned pointer is reachable, so it is safe to pass to C for a single call.
func cInts(list []int) *C.int {
	if len(list) == 0 {
		return nil
	}
	arr := make([]C.int, len(list))
	for i, v := range list {
		arr[i] = C.int(v)
	}
	return &arr[0]
}

// goStrings converts a C string list (CSLConstList) to a Go slice without
// freeing the C list.
func goStrings(list **C.char) (result []string) {
	csl := C.CSLConstList(unsafe.Pointer(list))
	n := int(C.CSLCount(csl))
	for i := 0; i < n; i++ {
		result = append(result, C.GoString(C.CSLGetField(csl, C.int(i))))
	}
	return
}
