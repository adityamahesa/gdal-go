package gdal

/*
#include "gdal_fwd.h"
*/
import "C"

type OGRSpatialReferences struct {
	cValue *C.OGRSpatialReferenceH
}

// GDALDatasets wraps a C GDALDatasetH array (its length carries the C
// dataset-count argument).
type GDALDatasets struct {
	cValue *C.GDALDatasetH
}

// GDALRasterBands is a contiguous list of raster bands, matching a C
// GDALRasterBandH array (its length carries the C band-count argument).
type GDALRasterBands []GDALRasterBand

// cPtr returns a pointer to the first C GDALRasterBandH, or nil for an empty
// list. GDALRasterBand is a single-field struct, so the slice is a contiguous
// C array.
func (b GDALRasterBands) cPtr() *C.GDALRasterBandH {
	if len(b) == 0 {
		return nil
	}
	return &b[0].cValue
}
