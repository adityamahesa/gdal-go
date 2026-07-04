package gdal

/*
#include "gdal_fwd.h"
*/
import "C"

type GDALMajorObject struct {
	cValue C.GDALMajorObjectH
}

type GDALDataset struct {
	cValue C.GDALDatasetH
}

type GDALRasterBand struct {
	cValue C.GDALRasterBandH
}

type GDALComputedRasterBand struct {
	cValue C.GDALComputedRasterBandH
}

type GDALDriver struct {
	cValue C.GDALDriverH
}

type GDALColorTable struct {
	cValue C.GDALColorTableH
}

type GDALRasterAttributeTable struct {
	cValue C.GDALRasterAttributeTableH
}

type GDALAsyncReader struct {
	cValue C.GDALAsyncReaderH
}

type GDALRelationship struct {
	cValue C.GDALRelationshipH
}

type OGRGeometry struct {
	cValue C.OGRGeometryH
}

type OGRGeomTransformer struct {
	cValue C.OGRGeomTransformerH
}

type OGRGeomCoordinatePrecision struct {
	cValue C.OGRGeomCoordinatePrecisionH
}

type OGRwkbExportOptions struct {
	cValue *C.OGRwkbExportOptions
}

type OGRPreparedGeometry struct {
	cValue C.OGRPreparedGeometryH
}

type OGRFieldDefn struct {
	cValue C.OGRFieldDefnH
}

type OGRFeatureDefn struct {
	cValue C.OGRFeatureDefnH
}

type OGRGeomFieldDefn struct {
	cValue C.OGRGeomFieldDefnH
}

type OGRFieldDomain struct {
	cValue C.OGRFieldDomainH
}

type OGRFeature struct {
	cValue C.OGRFeatureH
}

type OGRLayer struct {
	cValue C.OGRLayerH
}

type OGRDataSource struct {
	cValue C.OGRDataSourceH
}

type OGRSFDriver struct {
	cValue C.OGRSFDriverH
}

type OGRStyleMgr struct {
	cValue C.OGRStyleMgrH
}

type OGRStyleTool struct {
	cValue C.OGRStyleToolH
}

type OGRStyleTable struct {
	cValue C.OGRStyleTableH
}

type OGRSpatialReference struct {
	cValue C.OGRSpatialReferenceH
}

type OGRCoordinateTransformation struct {
	cValue C.OGRCoordinateTransformationH
}

type OGRCoordinateTransformationOptions struct {
	cValue C.OGRCoordinateTransformationOptionsH
}

type GNMNetwork struct {
	cValue C.GNMNetworkH
}

type GNMGenericNetwork struct {
	cValue C.GNMGenericNetworkH
}
