package gdal

/*
#include "cpl_port_preamble.h"
*/
import "C"

/**
 * \file cpl_port.h
 *
 * Core portability definitions for CPL.
 */

// Int32 type
type GInt32 C.GInt32

// Unsigned int32 type
type GUInt32 C.GUInt32

// Int16 type
type GInt16 C.GInt16

// Unsigned int16 type
type GUInt16 C.GUInt16

// Unsigned byte type
type GByte C.GByte

// Signed int8 type
type GInt8 C.GInt8

// Type for boolean values (alias to int)
type GBool C.GBool

// Large signed integer type (generally 64-bit integer type).
// Use GInt64 when exactly 64 bit is needed.
type GIntBig C.GIntBig

// Large unsigned integer type (generally 64-bit unsigned integer type).
// Use GUInt64 when exactly 64 bit is needed.
type GUIntBig C.GUIntBig

// Minimum GIntBig value
var GIntBigMin = GIntBig(C._GINTBIG_MIN)

// Maximum GIntBig value
var GIntBigMax = GIntBig(C._GINTBIG_MAX)

// Maximum GUIntBig value
var GUIntBigMax = GUIntBig(C._GUINTBIG_MAX)

const CPLHasGInt64 = C.CPL_HAS_GINT64

// Signed 64 bit integer type
type GInt64 C.GInt64

// Unsigned 64 bit integer type
type GUInt64 C.GUInt64

// Minimum GInt64 value
var GInt64Min = GIntBig(C._GINT64_MIN)

// Maximum GInt64 value
var GInt64Max = GIntBig(C._GINT64_MAX)

// Maximum GUInt64 value
var GUInt64Max = GUIntBig(C._GUINT64_MAX)

// Integer type large enough to hold the difference between 2 addresses
type GPtrDiff_t C.GPtrDiff_t

// Note: GUIntptr_t / CPL_IS_ALIGNED are guarded behind GDAL_COMPILATION and are
// not part of the public ABI here, so they are skipped.

// Printf formatting for GIntBig
var CPLFrmtGIB = C.GoString(C._CPL_FRMT_GIB)

// Printf formatting for GUIntBig
var CPLFrmtGUIB = C.GoString(C._CPL_FRMT_GUIB)

// MIN(a, b), MAX(a, b): deferred — value-producing macros to be wrapped in a later pass.

// ABS(x): deferred — value-producing macro to be wrapped in a later pass.

// PI definition
var MPi = float64(C._M_PI)

// CPLIsEqual(x, y): deferred — value-producing macro to be wrapped in a later pass.

// STRCASECMP(a, b), STRNCASECMP(a, b, n): deferred — value-producing macros to be wrapped in a later pass.

// EQUALN(a, b, n), EQUAL(a, b): deferred — value-producing macros to be wrapped in a later pass.

// STARTS_WITH(a, b), STARTS_WITH_CI(a, b): deferred — value-producing macros to be wrapped in a later pass.

// CPLIsNan(x), CPLIsInf(x), CPLIsFinite(x): deferred — value-producing macros to be wrapped in a later pass.

const CPLIsLSB = C.CPL_IS_LSB

// CPL_SWAP16(x), CPL_SWAP32(x), CPL_SWAP64(x): deferred — value-producing macros to be wrapped in a later pass.

// CPL_SWAP16PTR/CPL_SWAP32PTR/CPL_SWAP64PTR/CPL_SWAPDOUBLE, CPL_MSBWORD*/CPL_LSBWORD*,
// CPL_MSBPTR*/CPL_LSBPTR*, CPL_LSBINT16PTR/CPL_LSBINT32PTR, CPL_LSBSINT*PTR/CPL_LSBUINT*PTR:
// deferred — value-producing macros to be wrapped in a later pass.

const False = C.FALSE

const True = C.TRUE

// Type of a constant null-terminated list of nul terminated strings.
// Seen as char** from C (the definition cgo compiles against).
//
// This is the primary string-list handle throughout the binding: option lists
// and list-returning functions are typed CSLConstList rather than []string. Its
// methods (see cpl_string.go) wrap the CSL* API — build a list by chaining
// AddString from NullCSLConstList, read it back with Count/GetField. Ownership
// is per function: an owned list must be released with Destroy, a borrowed one
// (e.g. a view into GDAL-internal state) must not be. See NullCSLConstList for
// the nil sentinel.
type CSLConstList struct {
	cValue C.CSLConstList
}
