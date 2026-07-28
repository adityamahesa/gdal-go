#ifndef GDAL_GO_CPL_VSI_PREAMBLE_H
#define GDAL_GO_CPL_VSI_PREAMBLE_H

// On glibc, GDAL selects "struct stat64" for VSIStatBufL when large-file support
// is detected, but that struct is only fully defined when _LARGEFILE64_SOURCE is
// set before any system header is included. Without it the typedef is incomplete
// and its members (st_size/st_mode) are unaddressable. Must precede all includes.
#ifndef _LARGEFILE64_SOURCE
#define _LARGEFILE64_SOURCE 1
#endif

#include "cpl_vsi.h"

// All-ones offset bound read through an accessor so cgo evaluates it at runtime
// instead of as an untyped -1 constant that would overflow on conversion
// (see cpl_port.go for the same rationale on GUINTBIG_MAX).
static GUIntBig _VSI_L_OFFSET_MAX(void) { return VSI_L_OFFSET_MAX; }

// VSIStatBufL wraps a "struct stat", whose members (st_size/st_mode) are not
// directly addressable from cgo on all platforms (glibc lays them out in a way
// the Go side can't name). Read them in C, where s->field is always valid.
static long long _vsi_statbuf_size(const VSIStatBufL *s) { return (long long)s->st_size; }
static int _vsi_statbuf_mode(const VSIStatBufL *s) { return (int)s->st_mode; }

#endif /* GDAL_GO_CPL_VSI_PREAMBLE_H */
