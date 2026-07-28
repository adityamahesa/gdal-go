package gdal

import "unsafe"

func GDALCreateScaledProgress(dfMin, dfMax float64, pfnProgress GDALProgressFunc, pData unsafe.Pointer) (result unsafe.Pointer) {
	result = gdalCreateScaledProgress(dfMin, dfMax, pfnProgress, pData)
	return
}

func GDALDestroyScaledProgress(pData unsafe.Pointer) {
	gdalDestroyScaledProgress(pData)
}
