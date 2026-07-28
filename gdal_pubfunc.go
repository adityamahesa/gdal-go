package gdal

import "unsafe"

// Deprecated: use SizeBits or SizeBytes instead.
func (dt GDALDataType) Size() (result int) {
	result = gdalGetDataTypeSize(dt)
	return
}

func (dt GDALDataType) SizeBits() (result int) {
	result = gdalGetDataTypeSizeBits(dt)
	return
}

func (dt GDALDataType) SizeBytes() (result int) {
	result = gdalGetDataTypeSizeBytes(dt)
	return
}

func (dt GDALDataType) IsComplex() (result bool) {
	result = gdalDataTypeIsComplex(dt)
	return
}

func (dt GDALDataType) IsInteger() (result bool) {
	result = gdalDataTypeIsInteger(dt)
	return
}

func (dt GDALDataType) IsFloating() (result bool) {
	result = gdalDataTypeIsFloating(dt)
	return
}

func (dt GDALDataType) IsSigned() (result bool) {
	result = gdalDataTypeIsSigned(dt)
	return
}

func (dt GDALDataType) GetName() (result string) {
	result = gdalGetDataTypeName(dt)
	return
}

func GDALGetDataTypeByName(name string) (result GDALDataType) {
	result = gdalGetDataTypeByName(name)
	return
}

func (dt GDALDataType) Union(other GDALDataType) (result GDALDataType) {
	result = gdalDataTypeUnion(dt, other)
	return
}

func (dt GDALDataType) UnionWithValue(value float64, complex int) (result GDALDataType) {
	result = gdalDataTypeUnionWithValue(dt, value, complex)
	return
}

func GDALFindDataType(bits, signed, floating, complex int) (result GDALDataType) {
	result = gdalFindDataType(bits, signed, floating, complex)
	return
}

func GDALFindDataTypeForValue(value float64, complex int) (result GDALDataType) {
	result = gdalFindDataTypeForValue(value, complex)
	return
}

func (dt GDALDataType) AdjustValue(value float64) (result float64, clamped, rounded int) {
	result = gdalAdjustValueToDataType(dt, value, &clamped, &rounded)
	return
}

func (dt GDALDataType) IsValueExactAs(value float64) (result bool) {
	result = gdalIsValueExactAs(value, dt)
	return
}

func (dt GDALDataType) IsValueInRangeOf(value float64) (result bool) {
	result = gdalIsValueInRangeOf(value, dt)
	return
}

func (dt GDALDataType) GetNonComplex() (result GDALDataType) {
	result = gdalGetNonComplexDataType(dt)
	return
}

func (dt GDALDataType) IsConversionLossyTo(to GDALDataType) (result bool) {
	result = gdalDataTypeIsConversionLossy(dt, to)
	return
}

func (s GDALAsyncStatusType) GetName() (result string) {
	result = gdalGetAsyncStatusTypeName(s)
	return
}

func GDALGetAsyncStatusTypeByName(name string) (result GDALAsyncStatusType) {
	result = gdalGetAsyncStatusTypeByName(name)
	return
}

func (ci GDALColorInterp) GetName() (result string) {
	result = gdalGetColorInterpretationName(ci)
	return
}

func GDALGetColorInterpretationByName(name string) (result GDALColorInterp) {
	result = gdalGetColorInterpretationByName(name)
	return
}

func (pi GDALPaletteInterp) GetName() (result string) {
	result = gdalGetPaletteInterpretationName(pi)
	return
}

func GDALAllRegister() {
	gdalAllRegister()
}

func GDALRegisterPlugins() {
	gdalRegisterPlugins()
}

func GDALRegisterPlugin(name string) (err error) {
	err = cplErr(gdalRegisterPlugin(name))
	return
}

func (d GDALDriver) Create(name string, xSize, ySize, bands int, dataType GDALDataType, options CSLConstList) (result GDALDataset, err error) {
	result = gdalCreate(d, name, xSize, ySize, bands, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (d GDALDriver) CreateCopy(name string, src GDALDataset, strict int, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result GDALDataset, err error) {
	result = gdalCreateCopy(d, name, src, strict, options, progress, progressData)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALIdentifyDriver(filename string, fileList CSLConstList) (result GDALDriver, err error) {
	result = gdalIdentifyDriver(filename, fileList)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALIdentifyDriverEx(filename string, identifyFlags uint, allowedDrivers, fileList CSLConstList) (result GDALDriver, err error) {
	result = gdalIdentifyDriverEx(filename, identifyFlags, allowedDrivers, fileList)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALOpen(filename string, access GDALAccess) (result GDALDataset, err error) {
	result = gdalOpen(filename, access)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALOpenShared(filename string, access GDALAccess) (result GDALDataset, err error) {
	result = gdalOpenShared(filename, access)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALOpenEx(filename string, openFlags GDALOpenFlag, allowedDrivers, openOptions, siblingFiles CSLConstList) (result GDALDataset, err error) {
	result = gdalOpenEx(filename, openFlags, allowedDrivers, openOptions, siblingFiles)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALDumpOpenDatasets(filename string) (result int, err error) {
	return gdalDumpOpenDatasets(filename)
}

func GDALGetDriverByName(name string) (result GDALDriver, err error) {
	result = gdalGetDriverByName(name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALGetDriverCount() (result int) {
	result = gdalGetDriverCount()
	return
}

func GDALGetDriver(index int) (result GDALDriver, err error) {
	result = gdalGetDriver(index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALCreateDriver() (result GDALDriver, err error) {
	result = gdalCreateDriver()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (d GDALDriver) Destroy() {
	gdalDestroyDriver(d)
}

func (d GDALDriver) Register() (result int) {
	result = gdalRegisterDriver(d)
	return
}

func (d GDALDriver) Deregister() {
	gdalDeregisterDriver(d)
}

func GDALDestroyDriverManager() {
	gdalDestroyDriverManager()
}

func GDALDestroy() {
	gdalDestroy()
}

func (d GDALDriver) DeleteDataset(filename string) (err error) {
	err = cplErr(gdalDeleteDataset(d, filename))
	return
}

func (d GDALDriver) RenameDataset(newName, oldName string) (err error) {
	err = cplErr(gdalRenameDataset(d, newName, oldName))
	return
}

func (d GDALDriver) CopyDatasetFiles(newName, oldName string) (err error) {
	err = cplErr(gdalCopyDatasetFiles(d, newName, oldName))
	return
}

func (d GDALDriver) ValidateCreationOptions(options CSLConstList) (result bool) {
	result = gdalValidateCreationOptions(d, options)
	return
}

func GDALGetOutputDriversForDatasetName(destFilename string, flagRasterVector int, singleMatch, emitWarning bool) (result CSLConstList) {
	result = gdalGetOutputDriversForDatasetName(destFilename, flagRasterVector, singleMatch, emitWarning)
	return
}

func (d GDALDriver) HasOpenOption(openOptionName string) (result bool) {
	result = gdalDriverHasOpenOption(d, openOptionName)
	return
}

func (d GDALDriver) GetShortName() (result string) {
	result = gdalGetDriverShortName(d)
	return
}

func (d GDALDriver) GetLongName() (result string) {
	result = gdalGetDriverLongName(d)
	return
}

func (d GDALDriver) GetHelpTopic() (result string) {
	result = gdalGetDriverHelpTopic(d)
	return
}

func (d GDALDriver) GetCreationOptionList() (result string) {
	result = gdalGetDriverCreationOptionList(d)
	return
}

// GDALInitGCPs allocates and initializes a list of count Ground Control Points.
func GDALInitGCPs(count int) (result GDALGCPs) {
	result = make(GDALGCPs, count)
	gdalInitGCPs(count, result)
	return
}

func (g GDALGCPs) Deinit() {
	gdalDeinitGCPs(len(g), g)
}

func (g GDALGCPs) Duplicate() (result GDALGCPs) {
	result = gdalDuplicateGCPs(len(g), g)
	return
}

func (g GDALGCPs) ToGeoTransform(approxOK int) (geoTransform [6]float64, ok bool) {
	ok = gdalGCPsToGeoTransform(len(g), g, &geoTransform, approxOK) != 0
	return
}

func GDALInvGeoTransform(geoTransform [6]float64) (result [6]float64, ok bool) {
	ok = gdalInvGeoTransform(geoTransform, &result) != 0
	return
}

func GDALApplyGeoTransform(geoTransform [6]float64, pixel, line float64) (geoX, geoY float64) {
	gdalApplyGeoTransform(geoTransform, pixel, line, &geoX, &geoY)
	return
}

func GDALComposeGeoTransforms(a, b [6]float64) (result [6]float64) {
	gdalComposeGeoTransforms(a, b, &result)
	return
}

func (g GDALGCPs) ToHomography() (homography [9]float64, ok bool) {
	ok = gdalGCPsToHomography(len(g), g, &homography) != 0
	return
}

func GDALInvHomography(homography [9]float64) (result [9]float64, ok bool) {
	ok = gdalInvHomography(homography, &result) != 0
	return
}

func GDALApplyHomography(homography [9]float64, x, y float64) (outX, outY float64, ok bool) {
	ok = gdalApplyHomography(homography, x, y, &outX, &outY) != 0
	return
}

func GDALComposeHomographies(a, b [9]float64) (result [9]float64) {
	gdalComposeHomographies(a, b, &result)
	return
}

func (o GDALMajorObject) GetMetadataDomainList() (result CSLConstList) {
	result = gdalGetMetadataDomainList(o)
	return
}

func (o GDALMajorObject) GetMetadata(domain string) (result CSLConstList) {
	result = gdalGetMetadata(o, domain)
	return
}

func (o GDALMajorObject) SetMetadata(metadata CSLConstList, domain string) (err error) {
	err = cplErr(gdalSetMetadata(o, metadata, domain))
	return
}

func (o GDALMajorObject) GetMetadataItem(name, domain string) (result string) {
	result = gdalGetMetadataItem(o, name, domain)
	return
}

func (o GDALMajorObject) SetMetadataItem(name, value, domain string) (err error) {
	err = cplErr(gdalSetMetadataItem(o, name, value, domain))
	return
}

func (o GDALMajorObject) GetDescription() (result string) {
	result = gdalGetDescription(o)
	return
}

func (o GDALMajorObject) SetDescription(description string) {
	gdalSetDescription(o, description)
}

func (ds GDALDataset) GetDriver() (result GDALDriver) {
	result = gdalGetDatasetDriver(ds)
	return
}

func (ds GDALDataset) GetFileList() (result CSLConstList) {
	result = gdalGetFileList(ds)
	return
}

func (ds GDALDataset) MarkSuppressOnClose() {
	gdalDatasetMarkSuppressOnClose(ds)
}

func (ds GDALDataset) Close() (err error) {
	err = cplErr(gdalClose(ds))
	return
}

func (ds GDALDataset) RunCloseWithoutDestroying() (err error) {
	err = cplErr(gdalDatasetRunCloseWithoutDestroying(ds))
	return
}

func (ds GDALDataset) GetRasterXSize() (result int) {
	result = gdalGetRasterXSize(ds)
	return
}

func (ds GDALDataset) GetRasterYSize() (result int) {
	result = gdalGetRasterYSize(ds)
	return
}

func (ds GDALDataset) GetRasterCount() (result int) {
	result = gdalGetRasterCount(ds)
	return
}

func (ds GDALDataset) GetRasterBand(band int) (result GDALRasterBand, err error) {
	result = gdalGetRasterBand(ds, band)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) IsThreadSafe(scopeFlags int, options CSLConstList) (result bool) {
	result = gdalDatasetIsThreadSafe(ds, scopeFlags, options)
	return
}

func (ds GDALDataset) GetThreadSafeDataset(scopeFlags int, options CSLConstList) (result GDALDataset, err error) {
	result = gdalGetThreadSafeDataset(ds, scopeFlags, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) AddBand(dataType GDALDataType, options CSLConstList) (err error) {
	err = cplErr(gdalAddBand(ds, dataType, options))
	return
}

func (ds GDALDataset) BeginAsyncReader(xOff, yOff, xSize, ySize int, buf []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandMap []int, pixelSpace, lineSpace, bandSpace int, options CSLConstList) (result GDALAsyncReader, err error) {
	result = gdalBeginAsyncReader(ds, xOff, yOff, xSize, ySize, cBytes(buf), bufXSize, bufYSize, bufType, bandCount, bandMap, pixelSpace, lineSpace, bandSpace, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) EndAsyncReader(reader GDALAsyncReader) {
	gdalEndAsyncReader(ds, reader)
}

func (ds GDALDataset) RasterIO(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int) (err error) {
	err = cplErr(gdalDatasetRasterIO(ds, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, bandCount, bandList, pixelSpace, lineSpace, bandSpace))
	return
}

func (ds GDALDataset) RasterIOEx(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int64, extraArg GDALRasterIOExtraArg) (err error) {
	err = cplErr(gdalDatasetRasterIOEx(ds, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, bandCount, bandList, pixelSpace, lineSpace, bandSpace, extraArg))
	return
}

func (ds GDALDataset) AdviseRead(xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, options CSLConstList) (err error) {
	err = cplErr(gdalDatasetAdviseRead(ds, xOff, yOff, xSize, ySize, bufXSize, bufYSize, bufType, bandCount, bandList, options))
	return
}

func (ds GDALDataset) GetCompressionFormats(xOff, yOff, xSize, ySize, bandCount int, bandList []int) (result CSLConstList) {
	result = gdalDatasetGetCompressionFormats(ds, xOff, yOff, xSize, ySize, bandCount, bandList)
	return
}

func (ds GDALDataset) ReadCompressedData(format string, xOff, yOff, xSize, ySize, bandCount int, bandList []int) (buffer []byte, detailedFormat string, err error) {
	var cBuffer unsafe.Pointer
	var size int
	err = cplErr(gdalDatasetReadCompressedData(ds, format, xOff, yOff, xSize, ySize, bandCount, bandList, &cBuffer, &size, &detailedFormat))
	if err != nil {
		return
	}
	if cBuffer != nil {
		buffer = goBytes(cBuffer, size)
		vsiFree(cBuffer)
	}
	return
}

func (ds GDALDataset) GetProjectionRef() (result string) {
	result = gdalGetProjectionRef(ds)
	return
}

func (ds GDALDataset) GetSpatialRef() (result OGRSpatialReference) {
	result = gdalGetSpatialRef(ds)
	return
}

func (ds GDALDataset) SetProjection(projection string) (err error) {
	err = cplErr(gdalSetProjection(ds, projection))
	return
}

func (ds GDALDataset) SetSpatialRef(srs OGRSpatialReference) (err error) {
	err = cplErr(gdalSetSpatialRef(ds, srs))
	return
}

func (ds GDALDataset) GetGeoTransform() (geoTransform [6]float64, err error) {
	err = cplErr(gdalGetGeoTransform(ds, &geoTransform))
	return
}

func (ds GDALDataset) SetGeoTransform(geoTransform [6]float64) (err error) {
	err = cplErr(gdalSetGeoTransform(ds, geoTransform))
	return
}

func (ds GDALDataset) GetExtent(crs OGRSpatialReference) (envelope OGREnvelope, err error) {
	envelope = InitOGREnvelope()
	err = cplErr(gdalGetExtent(ds, envelope, crs))
	return
}

func (ds GDALDataset) GetExtentWGS84LongLat() (envelope OGREnvelope, err error) {
	envelope = InitOGREnvelope()
	err = cplErr(gdalGetExtentWGS84LongLat(ds, envelope))
	return
}

func (ds GDALDataset) GeolocationToPixelLine(geolocX, geolocY float64, srs OGRSpatialReference, transformerOptions CSLConstList) (pixel, line float64, err error) {
	err = cplErr(gdalDatasetGeolocationToPixelLine(ds, geolocX, geolocY, srs, &pixel, &line, transformerOptions))
	return
}

func (ds GDALDataset) GetGCPCount() (result int) {
	result = gdalGetGCPCount(ds)
	return
}

func (ds GDALDataset) GetGCPProjection() (result string) {
	result = gdalGetGCPProjection(ds)
	return
}

func (ds GDALDataset) GetGCPSpatialRef() (result OGRSpatialReference) {
	result = gdalGetGCPSpatialRef(ds)
	return
}

func (ds GDALDataset) GetGCPs() (result GDALGCPs) {
	result = gdalGetGCPs(ds)
	return
}

func (ds GDALDataset) SetGCPs(gcps GDALGCPs, projection string) (err error) {
	err = cplErr(gdalSetGCPs(ds, len(gcps), gcps, projection))
	return
}

func (ds GDALDataset) SetGCPs2(gcps GDALGCPs, srs OGRSpatialReference) (err error) {
	err = cplErr(gdalSetGCPs2(ds, len(gcps), gcps, srs))
	return
}

func (ds GDALDataset) GetInternalHandle(request string) unsafe.Pointer {
	return gdalGetInternalHandle(ds, request)
}

func (ds GDALDataset) Reference() (result int) {
	result = gdalReferenceDataset(ds)
	return
}

func (ds GDALDataset) Dereference() (result int) {
	result = gdalDereferenceDataset(ds)
	return
}

func (ds GDALDataset) Release() (result int) {
	result = gdalReleaseDataset(ds)
	return
}

func (ds GDALDataset) BuildOverviews(resampling string, overviewList, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalBuildOverviews(ds, resampling, len(overviewList), overviewList, len(bandList), bandList, progress, progressData))
	return
}

func (ds GDALDataset) BuildOverviewsEx(resampling string, overviewList, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer, options CSLConstList) (err error) {
	err = cplErr(gdalBuildOverviewsEx(ds, resampling, len(overviewList), overviewList, len(bandList), bandList, progress, progressData, options))
	return
}

func GDALGetOpenDatasets() (result []GDALDataset) {
	var datasets GDALDatasets
	var count int
	gdalGetOpenDatasets(&datasets, &count)
	if datasets.cValue == nil || count == 0 {
		return
	}
	src := unsafe.Slice(datasets.cValue, count)
	result = make([]GDALDataset, count)
	for i := range result {
		result[i] = GDALDataset{cValue: src[i]}
	}
	return
}

func (ds GDALDataset) GetAccess() (result int) {
	result = gdalGetAccess(ds)
	return
}

func (ds GDALDataset) FlushCache() (err error) {
	err = cplErr(gdalFlushCache(ds))
	return
}

func (ds GDALDataset) DropCache() (err error) {
	err = cplErr(gdalDropCache(ds))
	return
}

func (ds GDALDataset) CreateMaskBand(flags int) (err error) {
	err = cplErr(gdalCreateDatasetMaskBand(ds, flags))
	return
}

func (src GDALDataset) CopyWholeRaster(dst GDALDataset, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalDatasetCopyWholeRaster(src, dst, options, progress, progressData))
	return
}

func (src GDALRasterBand) CopyWholeRaster(dst GDALRasterBand, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalRasterBandCopyWholeRaster(src, dst, options, progress, progressData))
	return
}

func (src GDALRasterBand) RegenerateOverviews(overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalRegenerateOverviews(src, len(overviewBands), overviewBands, resampling, progress, progressData))
	return
}

func (src GDALRasterBand) RegenerateOverviewsEx(overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer, options CSLConstList) (err error) {
	err = cplErr(gdalRegenerateOverviewsEx(src, len(overviewBands), overviewBands, resampling, progress, progressData, options))
	return
}

func (ds GDALDataset) GetLayerCount() (result int) {
	result = gdalDatasetGetLayerCount(ds)
	return
}

func (ds GDALDataset) GetLayer(index int) (result OGRLayer, err error) {
	result = gdalDatasetGetLayer(ds, index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (l OGRLayer) GetDataset() (result GDALDataset, err error) {
	result = ogrLGetDataset(l)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) GetLayerByName(name string) (result OGRLayer, err error) {
	result = gdalDatasetGetLayerByName(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) IsLayerPrivate(index int) (result bool) {
	result = gdalDatasetIsLayerPrivate(ds, index)
	return
}

func (ds GDALDataset) DeleteLayer(index int) (err error) {
	err = ogrError(gdalDatasetDeleteLayer(ds, index))
	return
}

func (ds GDALDataset) CreateLayer(name string, srs OGRSpatialReference, geomType OGRwkbGeometryType, options CSLConstList) (result OGRLayer, err error) {
	result = gdalDatasetCreateLayer(ds, name, srs, geomType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) CreateLayerFromGeomFieldDefn(name string, geomFieldDefn OGRGeomFieldDefn, options CSLConstList) (result OGRLayer, err error) {
	result = gdalDatasetCreateLayerFromGeomFieldDefn(ds, name, geomFieldDefn, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) CopyLayer(srcLayer OGRLayer, newName string, options CSLConstList) (result OGRLayer, err error) {
	result = gdalDatasetCopyLayer(ds, srcLayer, newName, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) ResetReading() {
	gdalDatasetResetReading(ds)
}

func (ds GDALDataset) GetNextFeature(progress GDALProgressFunc, progressData unsafe.Pointer) (feature OGRFeature, belongingLayer OGRLayer, progressPct float64, err error) {
	feature = gdalDatasetGetNextFeature(ds, &belongingLayer, &progressPct, progress, progressData)
	if feature.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) TestCapability(capability string) (result bool) {
	result = gdalDatasetTestCapability(ds, capability)
	return
}

func (ds GDALDataset) ExecuteSQL(statement string, spatialFilter OGRGeometry, dialect string) (result OGRLayer, err error) {
	result = gdalDatasetExecuteSQL(ds, statement, spatialFilter, dialect)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) AbortSQL() (err error) {
	err = ogrError(gdalDatasetAbortSQL(ds))
	return
}

func (ds GDALDataset) ReleaseResultSet(layer OGRLayer) {
	gdalDatasetReleaseResultSet(ds, layer)
}

func (ds GDALDataset) GetStyleTable() (result OGRStyleTable, err error) {
	result = gdalDatasetGetStyleTable(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) SetStyleTableDirectly(styleTable OGRStyleTable) {
	gdalDatasetSetStyleTableDirectly(ds, styleTable)
}

func (ds GDALDataset) SetStyleTable(styleTable OGRStyleTable) {
	gdalDatasetSetStyleTable(ds, styleTable)
}

func (ds GDALDataset) StartTransaction(force int) (err error) {
	err = ogrError(gdalDatasetStartTransaction(ds, force))
	return
}

func (ds GDALDataset) CommitTransaction() (err error) {
	err = ogrError(gdalDatasetCommitTransaction(ds))
	return
}

func (ds GDALDataset) RollbackTransaction() (err error) {
	err = ogrError(gdalDatasetRollbackTransaction(ds))
	return
}

func (ds GDALDataset) ClearStatistics() {
	gdalDatasetClearStatistics(ds)
}

func (ds GDALDataset) AsMDArray(options CSLConstList) (result GDALMDArray, err error) {
	result = gdalDatasetAsMDArray(ds, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) GetFieldDomainNames(options CSLConstList) (result CSLConstList) {
	result = gdalDatasetGetFieldDomainNames(ds, options)
	return
}

func (ds GDALDataset) GetFieldDomain(name string) (result OGRFieldDomain, err error) {
	result = gdalDatasetGetFieldDomain(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) AddFieldDomain(fieldDomain OGRFieldDomain) (ok bool, failureReason string) {
	ok = gdalDatasetAddFieldDomain(ds, fieldDomain, &failureReason)
	return
}

func (ds GDALDataset) DeleteFieldDomain(name string) (ok bool, failureReason string) {
	ok = gdalDatasetDeleteFieldDomain(ds, name, &failureReason)
	return
}

func (ds GDALDataset) UpdateFieldDomain(fieldDomain OGRFieldDomain) (ok bool, failureReason string) {
	ok = gdalDatasetUpdateFieldDomain(ds, fieldDomain, &failureReason)
	return
}

func (ds GDALDataset) GetRelationshipNames(options CSLConstList) (result CSLConstList) {
	result = gdalDatasetGetRelationshipNames(ds, options)
	return
}

func (ds GDALDataset) GetRelationship(name string) (result GDALRelationship, err error) {
	result = gdalDatasetGetRelationship(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) AddRelationship(relationship GDALRelationship) (ok bool, failureReason string) {
	ok = gdalDatasetAddRelationship(ds, relationship, &failureReason)
	return
}

func (ds GDALDataset) DeleteRelationship(name string) (ok bool, failureReason string) {
	ok = gdalDatasetDeleteRelationship(ds, name, &failureReason)
	return
}

func (ds GDALDataset) UpdateRelationship(relationship GDALRelationship) (ok bool, failureReason string) {
	ok = gdalDatasetUpdateRelationship(ds, relationship, &failureReason)
	return
}

func GDALGetSubdatasetInfo(fileName string) (result GDALSubdatasetInfo, err error) {
	result = gdalGetSubdatasetInfo(fileName)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (i GDALSubdatasetInfo) GetPathComponent() (result string) {
	result = gdalSubdatasetInfoGetPathComponent(i)
	return
}

func (i GDALSubdatasetInfo) GetSubdatasetComponent() (result string) {
	result = gdalSubdatasetInfoGetSubdatasetComponent(i)
	return
}

func (i GDALSubdatasetInfo) ModifyPathComponent(newPath string) (result string) {
	result = gdalSubdatasetInfoModifyPathComponent(i, newPath)
	return
}

func (i GDALSubdatasetInfo) Destroy() {
	gdalDestroySubdatasetInfo(i)
}

func (b GDALRasterBand) GetRasterDataType() (result GDALDataType) {
	result = gdalGetRasterDataType(b)
	return
}

func (b GDALRasterBand) GetBlockSize() (xSize, ySize int) {
	gdalGetBlockSize(b, &xSize, &ySize)
	return
}

func (b GDALRasterBand) GetActualBlockSize(xBlockOff, yBlockOff int) (xValid, yValid int, err error) {
	err = cplErr(gdalGetActualBlockSize(b, xBlockOff, yBlockOff, &xValid, &yValid))
	return
}

func (b GDALRasterBand) AdviseRead(xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, options CSLConstList) (err error) {
	err = cplErr(gdalRasterAdviseRead(b, xOff, yOff, xSize, ySize, bufXSize, bufYSize, bufType, options))
	return
}

func (b GDALRasterBand) RasterIO(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int) (err error) {
	err = cplErr(gdalRasterIO(b, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, pixelSpace, lineSpace))
	return
}

func (b GDALRasterBand) RasterIOEx(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int64, extraArg GDALRasterIOExtraArg) (err error) {
	err = cplErr(gdalRasterIOEx(b, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, pixelSpace, lineSpace, extraArg))
	return
}

func (b GDALRasterBand) ReadBlock(xBlockOff, yBlockOff int, buffer []byte) (err error) {
	err = cplErr(gdalReadBlock(b, xBlockOff, yBlockOff, cBytes(buffer)))
	return
}

func (b GDALRasterBand) WriteBlock(xBlockOff, yBlockOff int, buffer []byte) (err error) {
	err = cplErr(gdalWriteBlock(b, xBlockOff, yBlockOff, cBytes(buffer)))
	return
}

func (b GDALRasterBand) GetXSize() (result int) {
	result = gdalGetRasterBandXSize(b)
	return
}

func (b GDALRasterBand) GetYSize() (result int) {
	result = gdalGetRasterBandYSize(b)
	return
}

func (b GDALRasterBand) GetAccess() (result GDALAccess) {
	result = gdalGetRasterAccess(b)
	return
}

func (b GDALRasterBand) GetBandNumber() (result int) {
	result = gdalGetBandNumber(b)
	return
}

func (b GDALRasterBand) GetDataset() (result GDALDataset, err error) {
	result = gdalGetBandDataset(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) GetColorInterpretation() (result GDALColorInterp) {
	result = gdalGetRasterColorInterpretation(b)
	return
}

func (b GDALRasterBand) SetColorInterpretation(colorInterp GDALColorInterp) (err error) {
	err = cplErr(gdalSetRasterColorInterpretation(b, colorInterp))
	return
}

func (b GDALRasterBand) GetColorTable() (result GDALColorTable, err error) {
	result = gdalGetRasterColorTable(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) SetColorTable(colorTable GDALColorTable) (err error) {
	err = cplErr(gdalSetRasterColorTable(b, colorTable))
	return
}

func (b GDALRasterBand) HasArbitraryOverviews() (result bool) {
	result = gdalHasArbitraryOverviews(b)
	return
}

func (b GDALRasterBand) GetOverviewCount() (result int) {
	result = gdalGetOverviewCount(b)
	return
}

func (b GDALRasterBand) GetOverview(index int) (result GDALRasterBand, err error) {
	result = gdalGetOverview(b, index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) GetNoDataValue() (value float64, ok bool) {
	var success int
	value = gdalGetRasterNoDataValue(b, &success)
	ok = success != 0
	return
}

func (b GDALRasterBand) GetNoDataValueAsInt64() (value int64, ok bool) {
	var success int
	value = gdalGetRasterNoDataValueAsInt64(b, &success)
	ok = success != 0
	return
}

func (b GDALRasterBand) GetNoDataValueAsUInt64() (value uint64, ok bool) {
	var success int
	value = gdalGetRasterNoDataValueAsUInt64(b, &success)
	ok = success != 0
	return
}

func (b GDALRasterBand) SetNoDataValue(value float64) (err error) {
	err = cplErr(gdalSetRasterNoDataValue(b, value))
	return
}

func (b GDALRasterBand) SetNoDataValueAsInt64(value int64) (err error) {
	err = cplErr(gdalSetRasterNoDataValueAsInt64(b, value))
	return
}

func (b GDALRasterBand) SetNoDataValueAsUInt64(value uint64) (err error) {
	err = cplErr(gdalSetRasterNoDataValueAsUInt64(b, value))
	return
}

func (b GDALRasterBand) DeleteNoDataValue() (err error) {
	err = cplErr(gdalDeleteRasterNoDataValue(b))
	return
}

func (b GDALRasterBand) GetCategoryNames() (result CSLConstList) {
	result = gdalGetRasterCategoryNames(b)
	return
}

func (b GDALRasterBand) SetCategoryNames(names CSLConstList) (err error) {
	err = cplErr(gdalSetRasterCategoryNames(b, names))
	return
}

func (b GDALRasterBand) GetMinimum() (value float64, ok bool) {
	var success int
	value = gdalGetRasterMinimum(b, &success)
	ok = success != 0
	return
}

func (b GDALRasterBand) GetMaximum() (value float64, ok bool) {
	var success int
	value = gdalGetRasterMaximum(b, &success)
	ok = success != 0
	return
}

func (b GDALRasterBand) GetStatistics(approxOK, force int) (min, max, mean, stdDev float64, err error) {
	err = cplErr(gdalGetRasterStatistics(b, approxOK, force, &min, &max, &mean, &stdDev))
	return
}

func (b GDALRasterBand) ComputeStatistics(approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max, mean, stdDev float64, err error) {
	err = cplErr(gdalComputeRasterStatistics(b, approxOK, &min, &max, &mean, &stdDev, progress, progressData))
	return
}

func (b GDALRasterBand) SetStatistics(min, max, mean, stdDev float64) (err error) {
	err = cplErr(gdalSetRasterStatistics(b, min, max, mean, stdDev))
	return
}

func (b GDALRasterBand) AsMDArray() (result GDALMDArray, err error) {
	result = gdalRasterBandAsMDArray(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) GetUnitType() (result string) {
	result = gdalGetRasterUnitType(b)
	return
}

func (b GDALRasterBand) SetUnitType(newValue string) (err error) {
	err = cplErr(gdalSetRasterUnitType(b, newValue))
	return
}

func (b GDALRasterBand) GetOffset() (value float64, ok bool) {
	var success int
	value = gdalGetRasterOffset(b, &success)
	ok = success != 0
	return
}

func (b GDALRasterBand) SetOffset(newOffset float64) (err error) {
	err = cplErr(gdalSetRasterOffset(b, newOffset))
	return
}

func (b GDALRasterBand) GetScale() (value float64, ok bool) {
	var success int
	value = gdalGetRasterScale(b, &success)
	ok = success != 0
	return
}

func (b GDALRasterBand) SetScale(newScale float64) (err error) {
	err = cplErr(gdalSetRasterScale(b, newScale))
	return
}

func (b GDALRasterBand) ComputeMinMax(approxOK int) (minMax [2]float64, err error) {
	err = cplErr(gdalComputeRasterMinMax(b, approxOK, &minMax))
	return
}

func (b GDALRasterBand) ComputeMinMaxLocation() (min, max float64, minX, minY, maxX, maxY int, err error) {
	err = cplErr(gdalComputeRasterMinMaxLocation(b, &min, &max, &minX, &minY, &maxX, &maxY))
	return
}

func (b GDALRasterBand) FlushCache() (err error) {
	err = cplErr(gdalFlushRasterCache(b))
	return
}

func (b GDALRasterBand) DropCache() (err error) {
	err = cplErr(gdalDropRasterCache(b))
	return
}

// Deprecated: use GetHistogramEx.
func (b GDALRasterBand) GetHistogram(min, max float64, nBuckets, includeOutOfRange, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (histogram []int, err error) {
	histogram = make([]int, nBuckets)
	err = cplErr(gdalGetRasterHistogram(b, min, max, nBuckets, histogram, includeOutOfRange, approxOK, progress, progressData))
	return
}

func (b GDALRasterBand) GetHistogramEx(min, max float64, nBuckets, includeOutOfRange, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (histogram []uint64, err error) {
	histogram = make([]uint64, nBuckets)
	err = cplErr(gdalGetRasterHistogramEx(b, min, max, nBuckets, histogram, includeOutOfRange, approxOK, progress, progressData))
	return
}

func (b GDALRasterBand) GetDefaultHistogramEx(force int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max float64, buckets int, histogram []uint64, err error) {
	err = cplErr(gdalGetDefaultHistogramEx(b, &min, &max, &buckets, &histogram, force, progress, progressData))
	return
}

func (b GDALRasterBand) SetDefaultHistogramEx(min, max float64, histogram []uint64) (err error) {
	err = cplErr(gdalSetDefaultHistogramEx(b, min, max, len(histogram), histogram))
	return
}

func (b GDALRasterBand) GetRandomSample(samples int) (buffer []float32, count int) {
	buffer = make([]float32, samples)
	count = gdalGetRandomRasterSample(b, samples, buffer)
	return
}

func (b GDALRasterBand) GetSampleOverview(desiredSamples int) (result GDALRasterBand, err error) {
	result = gdalGetRasterSampleOverview(b, desiredSamples)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) GetSampleOverviewEx(desiredSamples uint64) (result GDALRasterBand, err error) {
	result = gdalGetRasterSampleOverviewEx(b, desiredSamples)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) Fill(realValue, imaginaryValue float64) (err error) {
	err = cplErr(gdalFillRaster(b, realValue, imaginaryValue))
	return
}

func (b GDALRasterBand) ComputeBandStats(sampleStep int, progress GDALProgressFunc, progressData unsafe.Pointer) (mean, stdDev float64, err error) {
	err = cplErr(gdalComputeBandStats(b, sampleStep, &mean, &stdDev, progress, progressData))
	return
}

func (b GDALRasterBand) OverviewMagnitudeCorrection(overviews GDALRasterBands, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalOverviewMagnitudeCorrection(b, len(overviews), overviews, progress, progressData))
	return
}

func (b GDALRasterBand) GetDefaultRAT() (result GDALRasterAttributeTable, err error) {
	result = gdalGetDefaultRAT(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) SetDefaultRAT(rat GDALRasterAttributeTable) (err error) {
	err = cplErr(gdalSetDefaultRAT(b, rat))
	return
}

func (b GDALRasterBand) InterpolateAtPoint(pixel, line float64, interpolation GDALRIOResampleAlg) (realValue, imagValue float64, err error) {
	err = cplErr(gdalRasterInterpolateAtPoint(b, pixel, line, interpolation, &realValue, &imagValue))
	return
}

func (b GDALRasterBand) InterpolateAtGeolocation(geolocX, geolocY float64, srs OGRSpatialReference, interpolation GDALRIOResampleAlg, transformerOptions CSLConstList) (realValue, imagValue float64, err error) {
	err = cplErr(gdalRasterInterpolateAtGeolocation(b, geolocX, geolocY, srs, interpolation, &realValue, &imagValue, transformerOptions))
	return
}

func (b GDALRasterBand) GetMaskBand() (result GDALRasterBand, err error) {
	result = gdalGetMaskBand(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) GetMaskFlags() (result int) {
	result = gdalGetMaskFlags(b)
	return
}

func (b GDALRasterBand) CreateMaskBand(flags int) (err error) {
	err = cplErr(gdalCreateMaskBand(b, flags))
	return
}

func (b GDALRasterBand) IsMaskBand() (result bool) {
	result = gdalIsMaskBand(b)
	return
}

func (b GDALRasterBand) GetDataCoverageStatus(xOff, yOff, xSize, ySize, maskFlagStop int) (status int, dataPct float64) {
	status = gdalGetDataCoverageStatus(b, xOff, yOff, xSize, ySize, maskFlagStop, &dataPct)
	return
}

func (b GDALComputedRasterBand) Release() {
	gdalComputedRasterBandRelease(b)
}

func (b GDALRasterBand) UnaryOp(op GDALRasterAlgebraUnaryOperation) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandUnaryOp(b, op)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) BinaryOpBand(op GDALRasterAlgebraBinaryOperation, otherBand GDALRasterBand) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandBinaryOpBand(b, op, otherBand)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) BinaryOpDouble(op GDALRasterAlgebraBinaryOperation, constant float64) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandBinaryOpDouble(b, op, constant)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALRasterBandBinaryOpDoubleToBand(constant float64, op GDALRasterAlgebraBinaryOperation, band GDALRasterBand) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandBinaryOpDoubleToBand(constant, op, band)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) IfThenElse(thenBand, elseBand GDALRasterBand) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandIfThenElse(b, thenBand, elseBand)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) AsDataType(dataType GDALDataType) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandAsDataType(b, dataType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALMaximumOfNBands(bands GDALRasterBands) (result GDALComputedRasterBand, err error) {
	result = gdalMaximumOfNBands(len(bands), bands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) MaxConstant(constant float64) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandMaxConstant(b, constant)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALMinimumOfNBands(bands GDALRasterBands) (result GDALComputedRasterBand, err error) {
	result = gdalMinimumOfNBands(len(bands), bands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (b GDALRasterBand) MinConstant(constant float64) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandMinConstant(b, constant)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALMeanOfNBands(bands GDALRasterBands) (result GDALComputedRasterBand, err error) {
	result = gdalMeanOfNBands(len(bands), bands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (r GDALAsyncReader) GetNextUpdatedRegion(timeout float64) (status GDALAsyncStatusType, xBufOff, yBufOff, xBufSize, yBufSize int) {
	status = gdalARGetNextUpdatedRegion(r, timeout, &xBufOff, &yBufOff, &xBufSize, &yBufSize)
	return
}

func (r GDALAsyncReader) LockBuffer(timeout float64) (result bool) {
	result = gdalARLockBuffer(r, timeout)
	return
}

func (r GDALAsyncReader) UnlockBuffer() {
	gdalARUnlockBuffer(r)
}

func GDALSwapWords(data []byte, wordSize, wordCount, wordSkip int) {
	gdalSwapWords(cBytes(data), wordSize, wordCount, wordSkip)
}

func GDALSwapWordsEx(data []byte, wordSize, wordCount, wordSkip int) {
	gdalSwapWordsEx(cBytes(data), wordSize, wordCount, wordSkip)
}

func GDALCopyWords(src []byte, srcType GDALDataType, srcPixelOffset int, dst []byte, dstType GDALDataType, dstPixelOffset, wordCount int) {
	gdalCopyWords(cBytes(src), srcType, srcPixelOffset, cBytes(dst), dstType, dstPixelOffset, wordCount)
}

func GDALCopyWords64(src []byte, srcType GDALDataType, srcPixelOffset int, dst []byte, dstType GDALDataType, dstPixelOffset int, wordCount int64) {
	gdalCopyWords64(cBytes(src), srcType, srcPixelOffset, cBytes(dst), dstType, dstPixelOffset, wordCount)
}

func GDALCopyBits(src []byte, srcOffset, srcStep int, dst []byte, dstOffset, dstStep, bitCount, stepCount int) {
	gdalCopyBits(cBytes(src), srcOffset, srcStep, cBytes(dst), dstOffset, dstStep, bitCount, stepCount)
}

func GDALTranspose2D(src []byte, srcType GDALDataType, dst []byte, dstType GDALDataType, srcWidth, srcHeight int) {
	gdalTranspose2D(cBytes(src), srcType, cBytes(dst), dstType, srcWidth, srcHeight)
}

func (dt GDALDataType) GetNoDataReplacementValue(value float64) (result float64) {
	result = gdalGetNoDataReplacementValue(dt, value)
	return
}

func GDALLoadWorldFile(filename string) (geoTransform [6]float64, ok bool) {
	ok = gdalLoadWorldFile(filename, &geoTransform) != 0
	return
}

func GDALReadWorldFile(baseFilename, extension string) (geoTransform [6]float64, ok bool) {
	ok = gdalReadWorldFile(baseFilename, extension, &geoTransform) != 0
	return
}

func GDALWriteWorldFile(baseFilename, extension string, geoTransform [6]float64) (ok bool) {
	ok = gdalWriteWorldFile(baseFilename, extension, geoTransform) != 0
	return
}

func GDALDecToDMS(angle float64, axis string, precision int) (result string) {
	result = gdalDecToDMS(angle, axis, precision)
	return
}

func GDALPackedDMSToDec(packed float64) (result float64) {
	result = gdalPackedDMSToDec(packed)
	return
}

func GDALDecToPackedDMS(dec float64) (result float64) {
	result = gdalDecToPackedDMS(dec)
	return
}

func GDALVersionInfo(request string) (result string) {
	result = gdalVersionInfo(request)
	return
}

func GDALCheckVersion(versionMajor, versionMinor int, callingComponentName string) (result bool) {
	result = gdalCheckVersion(versionMajor, versionMinor, callingComponentName)
	return
}

func GDALExtractRPCInfoV2(metadata CSLConstList) (rpcInfo GDALRPCInfoV2, ok bool) {
	ok = gdalExtractRPCInfoV2(metadata, &rpcInfo) != 0
	return
}

func GDALCreateColorTable(paletteInterp GDALPaletteInterp) (result GDALColorTable, err error) {
	result = gdalCreateColorTable(paletteInterp)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ct GDALColorTable) Destroy() {
	gdalDestroyColorTable(ct)
}

func (ct GDALColorTable) Clone() (result GDALColorTable, err error) {
	result = gdalCloneColorTable(ct)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ct GDALColorTable) GetPaletteInterpretation() (result GDALPaletteInterp) {
	result = gdalGetPaletteInterpretation(ct)
	return
}

func (ct GDALColorTable) GetColorEntryCount() (result int) {
	result = gdalGetColorEntryCount(ct)
	return
}

func (ct GDALColorTable) GetColorEntry(index int) (result GDALColorEntry) {
	result = gdalGetColorEntry(ct, index)
	return
}

func (ct GDALColorTable) GetColorEntryAsRGB(index int) (result GDALColorEntry, ok bool) {
	ok = gdalGetColorEntryAsRGB(ct, index, &result) != 0
	return
}

func (ct GDALColorTable) SetColorEntry(index int, entry GDALColorEntry) {
	gdalSetColorEntry(ct, index, entry)
}

func (ct GDALColorTable) CreateColorRamp(startIndex int, startColor GDALColorEntry, endIndex int, endColor GDALColorEntry) {
	gdalCreateColorRamp(ct, startIndex, startColor, endIndex, endColor)
}

func GDALCreateRasterAttributeTable() (result GDALRasterAttributeTable, err error) {
	result = gdalCreateRasterAttributeTable()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (rat GDALRasterAttributeTable) Destroy() {
	gdalDestroyRasterAttributeTable(rat)
}

func (rat GDALRasterAttributeTable) GetColumnCount() (result int) {
	result = gdalRATGetColumnCount(rat)
	return
}

func (rat GDALRasterAttributeTable) GetNameOfCol(col int) (result string) {
	result = gdalRATGetNameOfCol(rat, col)
	return
}

func (rat GDALRasterAttributeTable) GetUsageOfCol(col int) (result GDALRATFieldUsage) {
	result = gdalRATGetUsageOfCol(rat, col)
	return
}

func (rat GDALRasterAttributeTable) GetTypeOfCol(col int) (result GDALRATFieldType) {
	result = gdalRATGetTypeOfCol(rat, col)
	return
}

func (ft GDALRATFieldType) GetName() (result string) {
	result = gdalGetRATFieldTypeName(ft)
	return
}

func (fu GDALRATFieldUsage) GetName() (result string) {
	result = gdalGetRATFieldUsageName(fu)
	return
}

func (rat GDALRasterAttributeTable) GetColOfUsage(usage GDALRATFieldUsage) (result int) {
	result = gdalRATGetColOfUsage(rat, usage)
	return
}

func (rat GDALRasterAttributeTable) GetRowCount() (result int) {
	result = gdalRATGetRowCount(rat)
	return
}

func (rat GDALRasterAttributeTable) GetValueAsString(row, field int) (result string) {
	result = gdalRATGetValueAsString(rat, row, field)
	return
}

func (rat GDALRasterAttributeTable) GetValueAsInt(row, field int) (result int) {
	result = gdalRATGetValueAsInt(rat, row, field)
	return
}

func (rat GDALRasterAttributeTable) GetValueAsDouble(row, field int) (result float64) {
	result = gdalRATGetValueAsDouble(rat, row, field)
	return
}

func (rat GDALRasterAttributeTable) GetValueAsBoolean(row, field int) (result bool) {
	result = gdalRATGetValueAsBoolean(rat, row, field)
	return
}

func (rat GDALRasterAttributeTable) GetValueAsDateTime(row, field int) (dateTime GDALRATDateTime, err error) {
	err = cplErr(gdalRATGetValueAsDateTime(rat, row, field, &dateTime))
	return
}

func (rat GDALRasterAttributeTable) GetValueAsWKBGeometry(row, field int) (result []byte) {
	result = gdalRATGetValueAsWKBGeometry(rat, row, field)
	return
}

func (rat GDALRasterAttributeTable) SetValueAsString(row, field int, value string) {
	gdalRATSetValueAsString(rat, row, field, value)
}

func (rat GDALRasterAttributeTable) SetValueAsInt(row, field, value int) {
	gdalRATSetValueAsInt(rat, row, field, value)
}

func (rat GDALRasterAttributeTable) SetValueAsDouble(row, field int, value float64) {
	gdalRATSetValueAsDouble(rat, row, field, value)
}

func (rat GDALRasterAttributeTable) SetValueAsBoolean(row, field int, value bool) (err error) {
	err = cplErr(gdalRATSetValueAsBoolean(rat, row, field, value))
	return
}

func (rat GDALRasterAttributeTable) SetValueAsDateTime(row, field int, dateTime GDALRATDateTime) (err error) {
	err = cplErr(gdalRATSetValueAsDateTime(rat, row, field, dateTime))
	return
}

func (rat GDALRasterAttributeTable) SetValueAsWKBGeometry(row, field int, wkb []byte) (err error) {
	err = cplErr(gdalRATSetValueAsWKBGeometry(rat, row, field, cBytes(wkb), len(wkb)))
	return
}

func (rat GDALRasterAttributeTable) ChangesAreWrittenToFile() (result bool) {
	result = gdalRATChangesAreWrittenToFile(rat)
	return
}

func (rat GDALRasterAttributeTable) ValuesIOAsDouble(rwFlag GDALRWFlag, field, startRow int, data []float64) (err error) {
	err = cplErr(gdalRATValuesIOAsDouble(rat, rwFlag, field, startRow, len(data), data))
	return
}

func (rat GDALRasterAttributeTable) ValuesIOAsInteger(rwFlag GDALRWFlag, field, startRow int, data []int) (err error) {
	err = cplErr(gdalRATValuesIOAsInteger(rat, rwFlag, field, startRow, len(data), data))
	return
}

func (rat GDALRasterAttributeTable) SetRowCount(count int) {
	gdalRATSetRowCount(rat, count)
}

func (rat GDALRasterAttributeTable) CreateColumn(name string, fieldType GDALRATFieldType, fieldUsage GDALRATFieldUsage) (err error) {
	err = cplErr(gdalRATCreateColumn(rat, name, fieldType, fieldUsage))
	return
}

func (rat GDALRasterAttributeTable) SetLinearBinning(row0Min, binSize float64) (err error) {
	err = cplErr(gdalRATSetLinearBinning(rat, row0Min, binSize))
	return
}

func (rat GDALRasterAttributeTable) GetLinearBinning() (row0Min, binSize float64, ok bool) {
	ok = gdalRATGetLinearBinning(rat, &row0Min, &binSize) != 0
	return
}

func (rat GDALRasterAttributeTable) SetTableType(tableType GDALRATTableType) (err error) {
	err = cplErr(gdalRATSetTableType(rat, tableType))
	return
}

func (rat GDALRasterAttributeTable) GetTableType() (result GDALRATTableType) {
	result = gdalRATGetTableType(rat)
	return
}

func (rat GDALRasterAttributeTable) InitializeFromColorTable(colorTable GDALColorTable) (err error) {
	err = cplErr(gdalRATInitializeFromColorTable(rat, colorTable))
	return
}

func (rat GDALRasterAttributeTable) TranslateToColorTable(entryCount int) (result GDALColorTable, err error) {
	result = gdalRATTranslateToColorTable(rat, entryCount)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (rat GDALRasterAttributeTable) DumpReadable(filename string) (err error) {
	return gdalRATDumpReadable(rat, filename)
}

func (rat GDALRasterAttributeTable) Clone() (result GDALRasterAttributeTable, err error) {
	result = gdalRATClone(rat)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (rat GDALRasterAttributeTable) GetRowOfValue(value float64) (result int) {
	result = gdalRATGetRowOfValue(rat, value)
	return
}

func (rat GDALRasterAttributeTable) RemoveStatistics() {
	gdalRATRemoveStatistics(rat)
}

func GDALRelationshipCreate(name, leftTableName, rightTableName string, cardinality GDALRelationshipCardinality) (result GDALRelationship, err error) {
	result = gdalRelationshipCreate(name, leftTableName, rightTableName, cardinality)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (r GDALRelationship) Destroy() {
	gdalDestroyRelationship(r)
}

func (r GDALRelationship) GetName() (result string) {
	result = gdalRelationshipGetName(r)
	return
}

func (r GDALRelationship) GetCardinality() (result GDALRelationshipCardinality) {
	result = gdalRelationshipGetCardinality(r)
	return
}

func (r GDALRelationship) GetLeftTableName() (result string) {
	result = gdalRelationshipGetLeftTableName(r)
	return
}

func (r GDALRelationship) GetRightTableName() (result string) {
	result = gdalRelationshipGetRightTableName(r)
	return
}

func (r GDALRelationship) GetMappingTableName() (result string) {
	result = gdalRelationshipGetMappingTableName(r)
	return
}

func (r GDALRelationship) SetMappingTableName(name string) {
	gdalRelationshipSetMappingTableName(r, name)
}

func (r GDALRelationship) GetLeftTableFields() (result CSLConstList) {
	result = gdalRelationshipGetLeftTableFields(r)
	return
}

func (r GDALRelationship) GetRightTableFields() (result CSLConstList) {
	result = gdalRelationshipGetRightTableFields(r)
	return
}

func (r GDALRelationship) SetLeftTableFields(fields CSLConstList) {
	gdalRelationshipSetLeftTableFields(r, fields)
}

func (r GDALRelationship) SetRightTableFields(fields CSLConstList) {
	gdalRelationshipSetRightTableFields(r, fields)
}

func (r GDALRelationship) GetLeftMappingTableFields() (result CSLConstList) {
	result = gdalRelationshipGetLeftMappingTableFields(r)
	return
}

func (r GDALRelationship) GetRightMappingTableFields() (result CSLConstList) {
	result = gdalRelationshipGetRightMappingTableFields(r)
	return
}

func (r GDALRelationship) SetLeftMappingTableFields(fields CSLConstList) {
	gdalRelationshipSetLeftMappingTableFields(r, fields)
}

func (r GDALRelationship) SetRightMappingTableFields(fields CSLConstList) {
	gdalRelationshipSetRightMappingTableFields(r, fields)
}

func (r GDALRelationship) GetType() (result GDALRelationshipType) {
	result = gdalRelationshipGetType(r)
	return
}

func (r GDALRelationship) SetType(relationshipType GDALRelationshipType) {
	gdalRelationshipSetType(r, relationshipType)
}

func (r GDALRelationship) GetForwardPathLabel() (result string) {
	result = gdalRelationshipGetForwardPathLabel(r)
	return
}

func (r GDALRelationship) SetForwardPathLabel(label string) {
	gdalRelationshipSetForwardPathLabel(r, label)
}

func (r GDALRelationship) GetBackwardPathLabel() (result string) {
	result = gdalRelationshipGetBackwardPathLabel(r)
	return
}

func (r GDALRelationship) SetBackwardPathLabel(label string) {
	gdalRelationshipSetBackwardPathLabel(r, label)
}

func (r GDALRelationship) GetRelatedTableType() (result string) {
	result = gdalRelationshipGetRelatedTableType(r)
	return
}

func (r GDALRelationship) SetRelatedTableType(relatedTableType string) {
	gdalRelationshipSetRelatedTableType(r, relatedTableType)
}

func GDALSetCacheMax(bytes int) {
	gdalSetCacheMax(bytes)
}

func GDALGetCacheMax() (result int) {
	result = gdalGetCacheMax()
	return
}

func GDALGetCacheUsed() (result int) {
	result = gdalGetCacheUsed()
	return
}

func GDALSetCacheMax64(bytes int64) {
	gdalSetCacheMax64(bytes)
}

func GDALGetCacheMax64() (result int64) {
	result = gdalGetCacheMax64()
	return
}

func GDALGetCacheUsed64() (result int64) {
	result = gdalGetCacheUsed64()
	return
}

func GDALFlushCacheBlock() (result bool) {
	result = gdalFlushCacheBlock()
	return
}

func GDALCreatePansharpenedVRT(xml string, panchroBand GDALRasterBand, inputSpectralBands GDALRasterBands) (result GDALDataset, err error) {
	result = gdalCreatePansharpenedVRT(xml, panchroBand, len(inputSpectralBands), inputSpectralBands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALGetJPEG2000Structure(filename string, options CSLConstList) (result CPLXMLNode, err error) {
	result = gdalGetJPEG2000Structure(filename, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (d GDALDriver) CreateMultiDimensional(name string, rootGroupOptions, options CSLConstList) (result GDALDataset, err error) {
	result = gdalCreateMultiDimensional(d, name, rootGroupOptions, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALExtendedDataTypeCreate(dataType GDALDataType) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreate(dataType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALExtendedDataTypeCreateString(maxStringLength int) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreateString(maxStringLength)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALExtendedDataTypeCreateStringEx(maxStringLength int, subType GDALExtendedDataTypeSubType) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreateStringEx(maxStringLength, subType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALExtendedDataTypeCreateCompound(name string, totalSize int, comps []GDALEDTComponent) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreateCompound(name, totalSize, len(comps), comps)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (edt GDALExtendedDataType) Release() {
	gdalExtendedDataTypeRelease(edt)
}

func (edt GDALExtendedDataType) GetName() (result string) {
	result = gdalExtendedDataTypeGetName(edt)
	return
}

func (edt GDALExtendedDataType) GetClass() (result GDALExtendedDataTypeClass) {
	result = gdalExtendedDataTypeGetClass(edt)
	return
}

func (edt GDALExtendedDataType) GetNumericDataType() (result GDALDataType) {
	result = gdalExtendedDataTypeGetNumericDataType(edt)
	return
}

func (edt GDALExtendedDataType) GetSize() (result int) {
	result = gdalExtendedDataTypeGetSize(edt)
	return
}

func (edt GDALExtendedDataType) GetMaxStringLength() (result int) {
	result = gdalExtendedDataTypeGetMaxStringLength(edt)
	return
}

func (edt GDALExtendedDataType) CanConvertTo(targetEDT GDALExtendedDataType) (result bool) {
	result = gdalExtendedDataTypeCanConvertTo(edt, targetEDT)
	return
}

func (edt GDALExtendedDataType) Equals(other GDALExtendedDataType) (result bool) {
	result = gdalExtendedDataTypeEquals(edt, other)
	return
}

func (edt GDALExtendedDataType) GetSubType() (result GDALExtendedDataTypeSubType) {
	result = gdalExtendedDataTypeGetSubType(edt)
	return
}

func (edt GDALExtendedDataType) GetRAT() (result GDALRasterAttributeTable, err error) {
	result = gdalExtendedDataTypeGetRAT(edt)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func GDALEDTComponentCreate(name string, offset int, dataType GDALExtendedDataType) (result GDALEDTComponent, err error) {
	result = gdalEDTComponentCreate(name, offset, dataType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (comp GDALEDTComponent) Release() {
	gdalEDTComponentRelease(comp)
}

func (comp GDALEDTComponent) GetName() (result string) {
	result = gdalEDTComponentGetName(comp)
	return
}

func (comp GDALEDTComponent) GetOffset() (result int) {
	result = gdalEDTComponentGetOffset(comp)
	return
}

func (comp GDALEDTComponent) GetType() (result GDALExtendedDataType, err error) {
	result = gdalEDTComponentGetType(comp)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds GDALDataset) GetRootGroup() (result GDALGroup, err error) {
	result = gdalDatasetGetRootGroup(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) Release() {
	gdalGroupRelease(g)
}

func (g GDALGroup) GetName() (result string) {
	result = gdalGroupGetName(g)
	return
}

func (g GDALGroup) GetFullName() (result string) {
	result = gdalGroupGetFullName(g)
	return
}

func (g GDALGroup) GetMDArrayNames(options CSLConstList) (result CSLConstList) {
	result = gdalGroupGetMDArrayNames(g, options)
	return
}

func (g GDALGroup) GetMDArrayFullNamesRecursive(groupOptions, arrayOptions CSLConstList) (result CSLConstList) {
	result = gdalGroupGetMDArrayFullNamesRecursive(g, groupOptions, arrayOptions)
	return
}

func (g GDALGroup) OpenMDArray(name string, options CSLConstList) (result GDALMDArray, err error) {
	result = gdalGroupOpenMDArray(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) OpenMDArrayFromFullname(name string, options CSLConstList) (result GDALMDArray, err error) {
	result = gdalGroupOpenMDArrayFromFullname(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) ResolveMDArray(name, startingPoint string, options CSLConstList) (result GDALMDArray, err error) {
	result = gdalGroupResolveMDArray(g, name, startingPoint, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) GetGroupNames(options CSLConstList) (result CSLConstList) {
	result = gdalGroupGetGroupNames(g, options)
	return
}

func (g GDALGroup) OpenGroup(name string, options CSLConstList) (result GDALGroup, err error) {
	result = gdalGroupOpenGroup(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) OpenGroupFromFullname(name string, options CSLConstList) (result GDALGroup, err error) {
	result = gdalGroupOpenGroupFromFullname(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) GetVectorLayerNames(options CSLConstList) (result CSLConstList) {
	result = gdalGroupGetVectorLayerNames(g, options)
	return
}

func (g GDALGroup) OpenVectorLayer(name string, options CSLConstList) (result OGRLayer, err error) {
	result = gdalGroupOpenVectorLayer(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) GetDimensions(options CSLConstList) (result []GDALDimension) {
	result = gdalGroupGetDimensions(g, options)
	return
}

func (g GDALGroup) GetAttribute(name string) (result GDALAttribute, err error) {
	result = gdalGroupGetAttribute(g, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) GetAttributes(options CSLConstList) (result []GDALAttribute) {
	result = gdalGroupGetAttributes(g, options)
	return
}

func (g GDALGroup) GetStructuralInfo() (result CSLConstList) {
	result = gdalGroupGetStructuralInfo(g)
	return
}

func (g GDALGroup) CreateGroup(name string, options CSLConstList) (result GDALGroup, err error) {
	result = gdalGroupCreateGroup(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) DeleteGroup(name string, options CSLConstList) (result bool) {
	result = gdalGroupDeleteGroup(g, name, options)
	return
}

func (g GDALGroup) CreateDimension(name, dimType, direction string, size uint64, options CSLConstList) (result GDALDimension, err error) {
	result = gdalGroupCreateDimension(g, name, dimType, direction, size, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) CreateMDArray(name string, dimensions []GDALDimension, dataType GDALExtendedDataType, options CSLConstList) (result GDALMDArray, err error) {
	result = gdalGroupCreateMDArray(g, name, len(dimensions), dimensions, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) DeleteMDArray(name string, options CSLConstList) (result bool) {
	result = gdalGroupDeleteMDArray(g, name, options)
	return
}

func (g GDALGroup) CreateAttribute(name string, dimensions []uint64, dataType GDALExtendedDataType, options CSLConstList) (result GDALAttribute, err error) {
	result = gdalGroupCreateAttribute(g, name, len(dimensions), dimensions, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) DeleteAttribute(name string, options CSLConstList) (result bool) {
	result = gdalGroupDeleteAttribute(g, name, options)
	return
}

func (g GDALGroup) Rename(newName string) (result bool) {
	result = gdalGroupRename(g, newName)
	return
}

func (g GDALGroup) SubsetDimensionFromSelection(selection string, options CSLConstList) (result GDALGroup, err error) {
	result = gdalGroupSubsetDimensionFromSelection(g, selection, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g GDALGroup) GetDataTypeCount() (result int) {
	result = gdalGroupGetDataTypeCount(g)
	return
}

func (g GDALGroup) GetDataType(index int) (result GDALExtendedDataType, err error) {
	result = gdalGroupGetDataType(g, index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) Release() {
	gdalMDArrayRelease(a)
}

func (a GDALMDArray) GetName() (result string) {
	result = gdalMDArrayGetName(a)
	return
}

func (a GDALMDArray) GetFullName() (result string) {
	result = gdalMDArrayGetFullName(a)
	return
}

func (a GDALMDArray) GetTotalElementsCount() (result uint64) {
	result = gdalMDArrayGetTotalElementsCount(a)
	return
}

func (a GDALMDArray) GetDimensionCount() (result int) {
	result = gdalMDArrayGetDimensionCount(a)
	return
}

func (a GDALMDArray) GetDimensions() (result []GDALDimension) {
	result = gdalMDArrayGetDimensions(a)
	return
}

func (a GDALMDArray) GetDataType() (result GDALExtendedDataType, err error) {
	result = gdalMDArrayGetDataType(a)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) GetAttribute(name string) (result GDALAttribute, err error) {
	result = gdalMDArrayGetAttribute(a, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) GetAttributes(options CSLConstList) (result []GDALAttribute) {
	result = gdalMDArrayGetAttributes(a, options)
	return
}

func (a GDALMDArray) CreateAttribute(name string, dimensions []uint64, dataType GDALExtendedDataType, options CSLConstList) (result GDALAttribute, err error) {
	result = gdalMDArrayCreateAttribute(a, name, len(dimensions), dimensions, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) DeleteAttribute(name string, options CSLConstList) (result bool) {
	result = gdalMDArrayDeleteAttribute(a, name, options)
	return
}

func (a GDALMDArray) Resize(newDimSizes []uint64, options CSLConstList) (result bool) {
	result = gdalMDArrayResize(a, newDimSizes, options)
	return
}

func (a GDALMDArray) GetNoDataValueAsDouble() (value float64, ok bool) {
	var has int
	value = gdalMDArrayGetNoDataValueAsDouble(a, &has)
	ok = has != 0
	return
}

func (a GDALMDArray) GetNoDataValueAsInt64() (value int64, ok bool) {
	var has int
	value = gdalMDArrayGetNoDataValueAsInt64(a, &has)
	ok = has != 0
	return
}

func (a GDALMDArray) GetNoDataValueAsUInt64() (value uint64, ok bool) {
	var has int
	value = gdalMDArrayGetNoDataValueAsUInt64(a, &has)
	ok = has != 0
	return
}

func (a GDALMDArray) SetNoDataValueAsDouble(value float64) (result bool) {
	result = gdalMDArraySetNoDataValueAsDouble(a, value)
	return
}

func (a GDALMDArray) SetNoDataValueAsInt64(value int64) (result bool) {
	result = gdalMDArraySetNoDataValueAsInt64(a, value)
	return
}

func (a GDALMDArray) SetNoDataValueAsUInt64(value uint64) (result bool) {
	result = gdalMDArraySetNoDataValueAsUInt64(a, value)
	return
}

func (a GDALMDArray) SetScale(scale float64) (result bool) {
	result = gdalMDArraySetScale(a, scale)
	return
}

func (a GDALMDArray) SetScaleEx(scale float64, storageType GDALDataType) (result bool) {
	result = gdalMDArraySetScaleEx(a, scale, storageType)
	return
}

func (a GDALMDArray) GetScale() (value float64, ok bool) {
	var has int
	value = gdalMDArrayGetScale(a, &has)
	ok = has != 0
	return
}

func (a GDALMDArray) SetOffset(offset float64) (result bool) {
	result = gdalMDArraySetOffset(a, offset)
	return
}

func (a GDALMDArray) SetOffsetEx(offset float64, storageType GDALDataType) (result bool) {
	result = gdalMDArraySetOffsetEx(a, offset, storageType)
	return
}

func (a GDALMDArray) GetOffset() (value float64, ok bool) {
	var has int
	value = gdalMDArrayGetOffset(a, &has)
	ok = has != 0
	return
}

func (a GDALMDArray) SetUnit(unit string) (result bool) {
	result = gdalMDArraySetUnit(a, unit)
	return
}

func (a GDALMDArray) GetUnit() (result string) {
	result = gdalMDArrayGetUnit(a)
	return
}

func (a GDALMDArray) SetSpatialRef(srs OGRSpatialReference) (result bool) {
	result = gdalMDArraySetSpatialRef(a, srs)
	return
}

func (a GDALMDArray) GetSpatialRef() (result OGRSpatialReference) {
	result = gdalMDArrayGetSpatialRef(a)
	return
}

func (a GDALMDArray) GetStructuralInfo() (result CSLConstList) {
	result = gdalMDArrayGetStructuralInfo(a)
	return
}

func (a GDALMDArray) GetView(viewExpr string) (result GDALMDArray, err error) {
	result = gdalMDArrayGetView(a, viewExpr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) Transpose(mapNewAxisToOldAxis []int) (result GDALMDArray, err error) {
	result = gdalMDArrayTranspose(a, len(mapNewAxisToOldAxis), mapNewAxisToOldAxis)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) GetUnscaled() (result GDALMDArray, err error) {
	result = gdalMDArrayGetUnscaled(a)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) GetMask(options CSLConstList) (result GDALMDArray, err error) {
	result = gdalMDArrayGetMask(a, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) AsClassicDataset(xDim, yDim int) (result GDALDataset, err error) {
	result = gdalMDArrayAsClassicDataset(a, xDim, yDim)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) AsClassicDatasetEx(xDim, yDim int, rootGroup GDALGroup, options CSLConstList) (result GDALDataset, err error) {
	result = gdalMDArrayAsClassicDatasetEx(a, xDim, yDim, rootGroup, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) GetStatistics(dataset GDALDataset, approxOK, force int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max, mean, stdDev float64, validCount uint64, err error) {
	err = cplErr(gdalMDArrayGetStatistics(a, dataset, approxOK, force, &min, &max, &mean, &stdDev, &validCount, progress, progressData))
	return
}

func (a GDALMDArray) ComputeStatistics(dataset GDALDataset, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max, mean, stdDev float64, validCount uint64, ok bool) {
	ok = gdalMDArrayComputeStatistics(a, dataset, approxOK, &min, &max, &mean, &stdDev, &validCount, progress, progressData)
	return
}

func (a GDALMDArray) GetResampled(newDims []GDALDimension, resampleAlg GDALRIOResampleAlg, targetSRS OGRSpatialReference, options CSLConstList) (result GDALMDArray, err error) {
	result = gdalMDArrayGetResampled(a, len(newDims), newDims, resampleAlg, targetSRS, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) GetGridded(gridOptions string, xArray, yArray GDALMDArray, options CSLConstList) (result GDALMDArray, err error) {
	result = gdalMDArrayGetGridded(a, gridOptions, xArray, yArray, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (a GDALMDArray) GetCoordinateVariables() (result []GDALMDArray) {
	result = gdalMDArrayGetCoordinateVariables(a)
	return
}

func (a GDALMDArray) Cache(options CSLConstList) (result bool) {
	result = gdalMDArrayCache(a, options)
	return
}

func (a GDALMDArray) Rename(newName string) (result bool) {
	result = gdalMDArrayRename(a, newName)
	return
}

func GDALCreateRasterAttributeTableFromMDArrays(tableType GDALRATTableType, arrays []GDALMDArray, usages []GDALRATFieldUsage) (result GDALRasterAttributeTable, err error) {
	result = gdalCreateRasterAttributeTableFromMDArrays(tableType, len(arrays), arrays, usages)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (attr GDALAttribute) Release() {
	gdalAttributeRelease(attr)
}

func (attr GDALAttribute) GetName() (result string) {
	result = gdalAttributeGetName(attr)
	return
}

func (attr GDALAttribute) GetFullName() (result string) {
	result = gdalAttributeGetFullName(attr)
	return
}

func (attr GDALAttribute) GetTotalElementsCount() (result uint64) {
	result = gdalAttributeGetTotalElementsCount(attr)
	return
}

func (attr GDALAttribute) GetDimensionCount() (result int) {
	result = gdalAttributeGetDimensionCount(attr)
	return
}

func (attr GDALAttribute) GetDimensionsSize() (result []uint64) {
	result = gdalAttributeGetDimensionsSize(attr)
	return
}

func (attr GDALAttribute) GetDataType() (result GDALExtendedDataType, err error) {
	result = gdalAttributeGetDataType(attr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (attr GDALAttribute) ReadAsRaw() (result []byte) {
	result = gdalAttributeReadAsRaw(attr)
	return
}

func (attr GDALAttribute) ReadAsString() (result string) {
	result = gdalAttributeReadAsString(attr)
	return
}

func (attr GDALAttribute) ReadAsInt() (result int) {
	result = gdalAttributeReadAsInt(attr)
	return
}

func (attr GDALAttribute) ReadAsInt64() (result int64) {
	result = gdalAttributeReadAsInt64(attr)
	return
}

func (attr GDALAttribute) ReadAsDouble() (result float64) {
	result = gdalAttributeReadAsDouble(attr)
	return
}

func (attr GDALAttribute) ReadAsStringArray() (result CSLConstList) {
	result = gdalAttributeReadAsStringArray(attr)
	return
}

func (attr GDALAttribute) ReadAsIntArray() (result []int) {
	result = gdalAttributeReadAsIntArray(attr)
	return
}

func (attr GDALAttribute) ReadAsInt64Array() (result []int64) {
	result = gdalAttributeReadAsInt64Array(attr)
	return
}

func (attr GDALAttribute) ReadAsDoubleArray() (result []float64) {
	result = gdalAttributeReadAsDoubleArray(attr)
	return
}

func (attr GDALAttribute) WriteRaw(data []byte) (result bool) {
	result = gdalAttributeWriteRaw(attr, cBytes(data), len(data))
	return
}

func (attr GDALAttribute) WriteString(value string) (result bool) {
	result = gdalAttributeWriteString(attr, value)
	return
}

func (attr GDALAttribute) WriteStringArray(values CSLConstList) (result bool) {
	result = gdalAttributeWriteStringArray(attr, values)
	return
}

func (attr GDALAttribute) WriteInt(value int) (result bool) {
	result = gdalAttributeWriteInt(attr, value)
	return
}

func (attr GDALAttribute) WriteIntArray(values []int) (result bool) {
	result = gdalAttributeWriteIntArray(attr, values, len(values))
	return
}

func (attr GDALAttribute) WriteInt64(value int64) (result bool) {
	result = gdalAttributeWriteInt64(attr, value)
	return
}

func (attr GDALAttribute) WriteInt64Array(values []int64) (result bool) {
	result = gdalAttributeWriteInt64Array(attr, values, len(values))
	return
}

func (attr GDALAttribute) WriteDouble(value float64) (result bool) {
	result = gdalAttributeWriteDouble(attr, value)
	return
}

func (attr GDALAttribute) WriteDoubleArray(values []float64) (result bool) {
	result = gdalAttributeWriteDoubleArray(attr, values, len(values))
	return
}

func (attr GDALAttribute) Rename(newName string) (result bool) {
	result = gdalAttributeRename(attr, newName)
	return
}

func (dim GDALDimension) Release() {
	gdalDimensionRelease(dim)
}

func (dim GDALDimension) GetName() (result string) {
	result = gdalDimensionGetName(dim)
	return
}

func (dim GDALDimension) GetFullName() (result string) {
	result = gdalDimensionGetFullName(dim)
	return
}

func (dim GDALDimension) GetType() (result string) {
	result = gdalDimensionGetType(dim)
	return
}

func (dim GDALDimension) GetDirection() (result string) {
	result = gdalDimensionGetDirection(dim)
	return
}

func (dim GDALDimension) GetSize() (result uint64) {
	result = gdalDimensionGetSize(dim)
	return
}

func (dim GDALDimension) GetIndexingVariable() (result GDALMDArray, err error) {
	result = gdalDimensionGetIndexingVariable(dim)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (dim GDALDimension) SetIndexingVariable(array GDALMDArray) (result bool) {
	result = gdalDimensionSetIndexingVariable(dim, array)
	return
}

func (dim GDALDimension) Rename(newName string) (result bool) {
	result = gdalDimensionRename(dim, newName)
	return
}
