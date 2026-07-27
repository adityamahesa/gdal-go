package gdal

/*
#include "util_preamble.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// cFile aliases a C FILE pointer. cFOpen produces one and the FILE*-based C
// covers accept it directly; aliasing just keeps those signatures reading as
// cFile rather than *C.FILE.
type cFile = *C.FILE

// cFOpen opens filename with the given mode (e.g. "w") and returns a cFile
// handle together with a close func. The caller must call close when done.
func cFOpen(filename, mode string) (fp cFile, closeFn func(), err error) {
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
