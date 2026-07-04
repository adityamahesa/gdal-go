package gdal

/*
#include "gdal_fwd.h"
*/
import "C"

type OGRSpatialReferences struct {
	cValue *C.OGRSpatialReferenceH
}
