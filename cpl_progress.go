package gdal

/*
#include "cpl_progress.h"

static GDALProgressFunc gdalDummyProgressFn()   { return GDALDummyProgress; }
static GDALProgressFunc gdalTermProgressFn()    { return GDALTermProgress; }
static GDALProgressFunc gdalScaledProgressFn()  { return GDALScaledProgress; }
*/
import "C"
import "unsafe"

type GDALProgressFunc struct {
	cValue C.GDALProgressFunc
}

var GDALDummyProgress = GDALProgressFunc{cValue: C.gdalDummyProgressFn()}
var GDALTermProgress = GDALProgressFunc{cValue: C.gdalTermProgressFn()}
var GDALScaledProgress = GDALProgressFunc{cValue: C.gdalScaledProgressFn()}

func gdalCreateScaledProgress(dfMin, dfMax float64, pfnProgress GDALProgressFunc, pData unsafe.Pointer) (result unsafe.Pointer) {
	result = C.GDALCreateScaledProgress(C.double(dfMin), C.double(dfMax), pfnProgress.cValue, pData)
	return
}

func GDALCreateScaledProgress(dfMin, dfMax float64, pfnProgress GDALProgressFunc, pData unsafe.Pointer) (result unsafe.Pointer) {
	result = gdalCreateScaledProgress(dfMin, dfMax, pfnProgress, pData)
	return
}

func gdalDestroyScaledProgress(pData unsafe.Pointer) {
	C.GDALDestroyScaledProgress(pData)
}

func GDALDestroyScaledProgress(pData unsafe.Pointer) {
	gdalDestroyScaledProgress(pData)
}
