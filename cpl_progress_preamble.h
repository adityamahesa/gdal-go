#ifndef GDAL_GO_CPL_PROGRESS_PREAMBLE_H
#define GDAL_GO_CPL_PROGRESS_PREAMBLE_H

#include "cpl_progress.h"

static GDALProgressFunc gdalDummyProgressFn()   { return GDALDummyProgress; }
static GDALProgressFunc gdalTermProgressFn()    { return GDALTermProgress; }
static GDALProgressFunc gdalScaledProgressFn()  { return GDALScaledProgress; }

#endif /* GDAL_GO_CPL_PROGRESS_PREAMBLE_H */
