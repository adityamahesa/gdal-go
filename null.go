package gdal

import "unsafe"

// NullPointer is the nil raw pointer sentinel.
var NullPointer unsafe.Pointer = nil

// NullCSLConstList is the nil (empty) string-list sentinel.
var NullCSLConstList CSLConstList = CSLConstList{
	cValue: nil,
}

var NullCPLErrorHandler CPLErrorHandler = CPLErrorHandler{
	cValue: nil,
}

var NullCPLLockFileHandle CPLLockFileHandle = CPLLockFileHandle{
	cValue: nil,
}

var NullCPLXMLNode CPLXMLNode = CPLXMLNode{
	cValue: nil,
}

var NullGDALAsyncReader GDALAsyncReader = GDALAsyncReader{
	cValue: nil,
}

var NullGDALAttribute GDALAttribute = GDALAttribute{
	cValue: nil,
}

var NullGDALColorTable GDALColorTable = GDALColorTable{
	cValue: nil,
}

var NullGDALComputedRasterBand GDALComputedRasterBand = GDALComputedRasterBand{
	cValue: nil,
}

var NullGDALDataset GDALDataset = GDALDataset{
	cValue: nil,
}

var NullGDALDatasets GDALDatasets = GDALDatasets{
	cValue: nil,
}

var NullGDALDimension GDALDimension = GDALDimension{
	cValue: nil,
}

var NullGDALDriver GDALDriver = GDALDriver{
	cValue: nil,
}

var NullGDALEDTComponent GDALEDTComponent = GDALEDTComponent{
	cValue: nil,
}

var NullGDALExtendedDataType GDALExtendedDataType = GDALExtendedDataType{
	cValue: nil,
}

var NullGDALGroup GDALGroup = GDALGroup{
	cValue: nil,
}

var NullGDALMDArray GDALMDArray = GDALMDArray{
	cValue: nil,
}

var NullGDALMajorObject GDALMajorObject = GDALMajorObject{
	cValue: nil,
}

var NullGDALProgressFunc GDALProgressFunc = GDALProgressFunc{
	cValue: nil,
}

var NullGDALRasterAttributeTable GDALRasterAttributeTable = GDALRasterAttributeTable{
	cValue: nil,
}

var NullGDALRasterBand GDALRasterBand = GDALRasterBand{
	cValue: nil,
}

var NullGDALRasterIOExtraArg GDALRasterIOExtraArg = GDALRasterIOExtraArg{
	cValue: nil,
}

var NullGDALRelationship GDALRelationship = GDALRelationship{
	cValue: nil,
}

var NullGDALSubdatasetInfo GDALSubdatasetInfo = GDALSubdatasetInfo{
	cValue: nil,
}

var NullGNMGenericNetwork GNMGenericNetwork = GNMGenericNetwork{
	cValue: nil,
}

var NullGNMNetwork GNMNetwork = GNMNetwork{
	cValue: nil,
}

var NullOGRCodedValue OGRCodedValue = OGRCodedValue{
	cValue: nil,
}

var NullOGRCoordinateTransformation OGRCoordinateTransformation = OGRCoordinateTransformation{
	cValue: nil,
}

var NullOGRCoordinateTransformationOptions OGRCoordinateTransformationOptions = OGRCoordinateTransformationOptions{
	cValue: nil,
}

var NullOGRDataSource OGRDataSource = OGRDataSource{
	cValue: nil,
}

var NullOGREnvelope OGREnvelope = OGREnvelope{
	cValue: nil,
}

var NullOGREnvelope3D OGREnvelope3D = OGREnvelope3D{
	cValue: nil,
}

var NullOGRFeature OGRFeature = OGRFeature{
	cValue: nil,
}

var NullOGRFeatureDefn OGRFeatureDefn = OGRFeatureDefn{
	cValue: nil,
}

var NullOGRField OGRField = OGRField{
	cValue: nil,
}

var NullOGRFieldDefn OGRFieldDefn = OGRFieldDefn{
	cValue: nil,
}

var NullOGRFieldDomain OGRFieldDomain = OGRFieldDomain{
	cValue: nil,
}

var NullOGRGeomCoordinatePrecision OGRGeomCoordinatePrecision = OGRGeomCoordinatePrecision{
	cValue: nil,
}

var NullOGRGeomFieldDefn OGRGeomFieldDefn = OGRGeomFieldDefn{
	cValue: nil,
}

var NullOGRGeomTransformer OGRGeomTransformer = OGRGeomTransformer{
	cValue: nil,
}

var NullOGRGeometry OGRGeometry = OGRGeometry{
	cValue: nil,
}

var NullOGRGeometryTypeCounter OGRGeometryTypeCounter = OGRGeometryTypeCounter{
	cValue: nil,
}

var NullOGRLayer OGRLayer = OGRLayer{
	cValue: nil,
}

var NullOGRPreparedGeometry OGRPreparedGeometry = OGRPreparedGeometry{
	cValue: nil,
}

var NullOGRSFDriver OGRSFDriver = OGRSFDriver{
	cValue: nil,
}

var NullOGRSpatialReference OGRSpatialReference = OGRSpatialReference{
	cValue: nil,
}

var NullOGRSpatialReferences OGRSpatialReferences = OGRSpatialReferences{
	cValue: nil,
}

var NullOGRStyleMgr OGRStyleMgr = OGRStyleMgr{
	cValue: nil,
}

var NullOGRStyleTable OGRStyleTable = OGRStyleTable{
	cValue: nil,
}

var NullOGRStyleTool OGRStyleTool = OGRStyleTool{
	cValue: nil,
}

var NullOGRwkbExportOptions OGRwkbExportOptions = OGRwkbExportOptions{
	cValue: nil,
}

var NullOSRCRSInfo OSRCRSInfo = OSRCRSInfo{
	cValue: nil,
}

var NullOSRCRSInfos OSRCRSInfos = OSRCRSInfos{
	cValue: nil,
}

var NullOSRCRSListParameters OSRCRSListParameters = OSRCRSListParameters{
	cValue: nil,
}

var NullVSILFile VSILFile = VSILFile{
	cValue: nil,
}

var NullVSIDir VSIDir = VSIDir{
	cValue: nil,
}

var NullVSIDirEntry VSIDirEntry = VSIDirEntry{
	cValue: nil,
}

var NullVSIFilesystemPluginCallbacksStruct VSIFilesystemPluginCallbacksStruct = VSIFilesystemPluginCallbacksStruct{
	cValue: nil,
}
