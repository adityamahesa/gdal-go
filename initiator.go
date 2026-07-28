package gdal

/*
#include "ogr_core_preamble.h"
#include "ogr_api_preamble.h"
#include "ogr_srs_api_preamble.h"
#include "gdal_preamble.h"
*/
import "C"

func InitOGREnvelope() OGREnvelope {
	return OGREnvelope{cValue: new(C.OGREnvelope)}
}

func InitOGREnvelope3D() OGREnvelope3D {
	return OGREnvelope3D{cValue: new(C.OGREnvelope3D)}
}

func InitOGRField() OGRField {
	return OGRField{cValue: new(C.OGRField)}
}

func InitOGRCodedValue() OGRCodedValue {
	return OGRCodedValue{cValue: new(C.OGRCodedValue)}
}

func InitOGRGeometryTypeCounter() OGRGeometryTypeCounter {
	return OGRGeometryTypeCounter{cValue: new(C.OGRGeometryTypeCounter)}
}

func InitOSRCRSInfo() OSRCRSInfo {
	return OSRCRSInfo{cValue: new(C.OSRCRSInfo)}
}

func InitGDALRasterIOExtraArg() GDALRasterIOExtraArg {
	return GDALRasterIOExtraArg{cValue: new(C.GDALRasterIOExtraArg)}
}
