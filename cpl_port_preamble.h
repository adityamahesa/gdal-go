#ifndef GDAL_GO_CPL_PORT_PREAMBLE_H
#define GDAL_GO_CPL_PORT_PREAMBLE_H

#include "cpl_port.h"

const GIntBig  _GINTBIG_MIN  = GINTBIG_MIN;
const GIntBig  _GINTBIG_MAX  = GINTBIG_MAX;
const GIntBig  _GINT64_MIN   = GINT64_MIN;
const GIntBig  _GINT64_MAX   = GINT64_MAX;
// All-ones unsigned bounds: kept as (non-const) globals so cgo reads them at
// runtime instead of as untyped -1 constants that would overflow on conversion.
GUIntBig _GUINTBIG_MAX = GUINTBIG_MAX;
GUIntBig _GUINT64_MAX  = GUINT64_MAX;
// Kept as a (non-const) global so cgo reads the full-precision double at runtime;
// as a constant cgo would marshal it through a low-precision decimal form.
double _M_PI = M_PI;
const char* const _CPL_FRMT_GIB  = CPL_FRMT_GIB;
const char* const _CPL_FRMT_GUIB = CPL_FRMT_GUIB;

#endif /* GDAL_GO_CPL_PORT_PREAMBLE_H */
