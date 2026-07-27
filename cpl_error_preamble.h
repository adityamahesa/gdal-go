#ifndef GDAL_GO_CPL_ERROR_PREAMBLE_H
#define GDAL_GO_CPL_ERROR_PREAMBLE_H

#include "cpl_error.h"

static CPLErrorHandler cplLoggingErrorHandlerFn()         { return CPLLoggingErrorHandler; }
static CPLErrorHandler cplDefaultErrorHandlerFn()         { return CPLDefaultErrorHandler; }
static CPLErrorHandler cplQuietErrorHandlerFn()           { return CPLQuietErrorHandler; }
static CPLErrorHandler cplQuietWarningsErrorHandlerFn()   { return CPLQuietWarningsErrorHandler; }

const CPLErr _VALIDATE_POINTER_ERR = VALIDATE_POINTER_ERR;

#endif /* GDAL_GO_CPL_ERROR_PREAMBLE_H */
