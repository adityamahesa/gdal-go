package gdal

/*
#include "util_csl_preamble.h"
*/
import "C"
import "unsafe"

// cslConstList wraps a raw C string list (char**, as returned by CSL and GDAL
// functions) in a CSLConstList handle. It is the single place that builds the
// handle from a C pointer, keeping the C.CSLConstList conversion in one file.
func cslConstList(raw **C.char) CSLConstList {
	return CSLConstList{cValue: C.CSLConstList(unsafe.Pointer(raw))}
}
