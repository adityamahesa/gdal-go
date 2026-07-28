#ifndef GDAL_GO_CPL_PORT_PREAMBLE_H
#define GDAL_GO_CPL_PORT_PREAMBLE_H

#include "cpl_port.h"

static GIntBig _GINTBIG_MIN(void) { return GINTBIG_MIN; }
static GIntBig _GINTBIG_MAX(void) { return GINTBIG_MAX; }
static GIntBig _GINT64_MIN(void)  { return GINT64_MIN; }
static GIntBig _GINT64_MAX(void)  { return GINT64_MAX; }
// All-ones unsigned bounds: read through an accessor so cgo evaluates them at
// runtime instead of as untyped -1 constants that would overflow on conversion.
static GUIntBig _GUINTBIG_MAX(void) { return GUINTBIG_MAX; }
static GUIntBig _GUINT64_MAX(void)  { return GUINT64_MAX; }
// Read through an accessor so cgo evaluates the full-precision double at runtime;
// as a constant cgo would marshal it through a low-precision decimal form.
static double _M_PI(void)               { return M_PI; }
static const char* _CPL_FRMT_GIB(void)  { return CPL_FRMT_GIB; }
static const char* _CPL_FRMT_GUIB(void) { return CPL_FRMT_GUIB; }

#endif /* GDAL_GO_CPL_PORT_PREAMBLE_H */
