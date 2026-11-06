package gdal

/*
#include "gdal_fwd_basic_preamble.h"
*/
import "C"

import "unsafe"

func (ds GDALDataset) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(ds.cValue))}
	return
}

func (b GDALRasterBand) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(b.cValue))}
	return
}

func (b GDALComputedRasterBand) RasterBand() (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALRasterBandH(unsafe.Pointer(b.cValue))}
	return
}

func (b GDALComputedRasterBand) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(b.cValue))}
	return
}

func (d GDALDriver) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(d.cValue))}
	return
}

func (l OGRLayer) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(l.cValue))}
	return
}

func (ds OGRDataSource) Dataset() (result GDALDataset) {
	result = GDALDataset{cValue: C.GDALDatasetH(unsafe.Pointer(ds.cValue))}
	return
}

func (ds OGRDataSource) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(ds.cValue))}
	return
}

func (dr OGRSFDriver) Driver() (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALDriverH(unsafe.Pointer(dr.cValue))}
	return
}

func (dr OGRSFDriver) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(dr.cValue))}
	return
}

func (n GNMNetwork) Dataset() (result GDALDataset) {
	result = GDALDataset{cValue: C.GDALDatasetH(unsafe.Pointer(n.cValue))}
	return
}

func (n GNMNetwork) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(n.cValue))}
	return
}

func (n GNMGenericNetwork) Network() (result GNMNetwork) {
	result = GNMNetwork{cValue: C.GNMNetworkH(unsafe.Pointer(n.cValue))}
	return
}

func (n GNMGenericNetwork) Dataset() (result GDALDataset) {
	result = GDALDataset{cValue: C.GDALDatasetH(unsafe.Pointer(n.cValue))}
	return
}

func (n GNMGenericNetwork) MajorObject() (result GDALMajorObject) {
	result = GDALMajorObject{cValue: C.GDALMajorObjectH(unsafe.Pointer(n.cValue))}
	return
}
