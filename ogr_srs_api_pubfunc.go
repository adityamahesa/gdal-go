package gdal

import "unsafe"

func (e OGRAxisOrientation) Name() (result string) {
	result = osrAxisEnumToName(e)
	return
}

func OSRSetPROJSearchPaths(paths CSLConstList) {
	osrSetPROJSearchPaths(paths)
}

func OSRGetPROJSearchPaths() (result CSLConstList) {
	result = osrGetPROJSearchPaths()
	return
}

func OSRSetPROJAuxDbPaths(paths CSLConstList) {
	osrSetPROJAuxDbPaths(paths)
}

func OSRGetPROJAuxDbPaths() (result CSLConstList) {
	result = osrGetPROJAuxDbPaths()
	return
}

func OSRSetPROJEnableNetwork(enabled int) {
	osrSetPROJEnableNetwork(enabled)
}

func OSRGetPROJEnableNetwork() (result int) {
	result = osrGetPROJEnableNetwork()
	return
}

func OSRGetPROJVersion() (major, minor, patch int) {
	major, minor, patch = osrGetPROJVersion()
	return
}

func OSRNewSpatialReference(wkt string) (result OGRSpatialReference, err error) {
	result = osrNewSpatialReference(wkt)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (sr OGRSpatialReference) CloneGeogCS() (result OGRSpatialReference, err error) {
	result = osrCloneGeogCS(sr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (sr OGRSpatialReference) Clone() (result OGRSpatialReference, err error) {
	result = osrClone(sr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (sr OGRSpatialReference) Destroy() {
	osrDestroySpatialReference(sr)
}

func (sr OGRSpatialReference) Reference() (result int) {
	result = osrReference(sr)
	return
}

func (sr OGRSpatialReference) Dereference() (result int) {
	result = osrDereference(sr)
	return
}

func (sr OGRSpatialReference) Release() {
	osrRelease(sr)
}

func (sr OGRSpatialReference) Validate() (err error) {
	err = ogrError(osrValidate(sr))
	return
}

func (sr OGRSpatialReference) ImportFromEPSG(nCode int) (err error) {
	err = ogrError(osrImportFromEPSG(sr, nCode))
	return
}

func (sr OGRSpatialReference) ImportFromEPSGA(nCode int) (err error) {
	err = ogrError(osrImportFromEPSGA(sr, nCode))
	return
}

func (sr OGRSpatialReference) ImportFromWkt(wkt string) (err error) {
	err = ogrError(osrImportFromWkt(sr, wkt))
	return
}

func (sr OGRSpatialReference) ImportFromProj4(proj4 string) (err error) {
	err = ogrError(osrImportFromProj4(sr, proj4))
	return
}

func (sr OGRSpatialReference) ImportFromESRI(lines CSLConstList) (err error) {
	err = ogrError(osrImportFromESRI(sr, lines))
	return
}

func (sr OGRSpatialReference) ImportFromPCI(proj, units string, arParams []float64) (err error) {
	err = ogrError(osrImportFromPCI(sr, proj, units, arParams))
	return
}

func (sr OGRSpatialReference) ImportFromUSGS(projSys, zone int, arParams []float64, datum int) (err error) {
	err = ogrError(osrImportFromUSGS(sr, projSys, zone, arParams, datum))
	return
}

func (sr OGRSpatialReference) ImportFromXML(xmlString string) (err error) {
	err = ogrError(osrImportFromXML(sr, xmlString))
	return
}

func (sr OGRSpatialReference) ImportFromDict(dictFile, code string) (err error) {
	err = ogrError(osrImportFromDict(sr, dictFile, code))
	return
}

func (sr OGRSpatialReference) ImportFromPanorama(projSys, datum, ellipsoid int, arParams []float64) (err error) {
	err = ogrError(osrImportFromPanorama(sr, projSys, datum, ellipsoid, arParams))
	return
}

func (sr OGRSpatialReference) ImportFromOzi(lines CSLConstList) (err error) {
	err = ogrError(osrImportFromOzi(sr, lines))
	return
}

func (sr OGRSpatialReference) ImportFromMICoordSys(coordSys string) (err error) {
	err = ogrError(osrImportFromMICoordSys(sr, coordSys))
	return
}

func (sr OGRSpatialReference) ImportFromERM(proj, datum, units string) (err error) {
	err = ogrError(osrImportFromERM(sr, proj, datum, units))
	return
}

func (sr OGRSpatialReference) ImportFromUrl(url string) (err error) {
	err = ogrError(osrImportFromUrl(sr, url))
	return
}

func (sr OGRSpatialReference) ImportFromCF1(keyValues CSLConstList, units string) (err error) {
	err = ogrError(osrImportFromCF1(sr, keyValues, units))
	return
}

func (sr OGRSpatialReference) ExportToWkt() (result string, err error) {
	var status OGRErr
	result, status = osrExportToWkt(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToWktEx(options CSLConstList) (result string, err error) {
	var status OGRErr
	result, status = osrExportToWktEx(sr, options)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToPrettyWkt(simplify int) (result string, err error) {
	var status OGRErr
	result, status = osrExportToPrettyWkt(sr, simplify)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToPROJJSON(options CSLConstList) (result string, err error) {
	var status OGRErr
	result, status = osrExportToPROJJSON(sr, options)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToProj4() (result string, err error) {
	var status OGRErr
	result, status = osrExportToProj4(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToPCI() (proj, units string, params []float64, err error) {
	var status OGRErr
	proj, units, params, status = osrExportToPCI(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToUSGS() (projSys, zone int, params []float64, datum int, err error) {
	var status OGRErr
	projSys, zone, params, datum, status = osrExportToUSGS(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToXML(dialect string) (result string, err error) {
	var status OGRErr
	result, status = osrExportToXML(sr, dialect)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToPanorama() (projSys, datum, ellipsoid, zone int, params []float64, err error) {
	var status OGRErr
	projSys, datum, ellipsoid, zone, params, status = osrExportToPanorama(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToMICoordSys() (result string, err error) {
	var status OGRErr
	result, status = osrExportToMICoordSys(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToERM() (proj, datum, units string, err error) {
	var status OGRErr
	proj, datum, units, status = osrExportToERM(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) ExportToCF1(options CSLConstList) (gridMappingName string, keyValues CSLConstList, units string, err error) {
	var status OGRErr
	gridMappingName, keyValues, units, status = osrExportToCF1(sr, options)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) MorphToESRI() (err error) {
	err = ogrError(osrMorphToESRI(sr))
	return
}

func (sr OGRSpatialReference) MorphFromESRI() (err error) {
	err = ogrError(osrMorphFromESRI(sr))
	return
}

func (sr OGRSpatialReference) StripVertical() (err error) {
	err = ogrError(osrStripVertical(sr))
	return
}

func (sr OGRSpatialReference) ConvertToOtherProjection(targetProjection string, options CSLConstList) (result OGRSpatialReference, err error) {
	result = osrConvertToOtherProjection(sr, targetProjection, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (sr OGRSpatialReference) GetName() (result string) {
	result = osrGetName(sr)
	return
}

func (sr OGRSpatialReference) GetCelestialBodyName() (result string) {
	result = osrGetCelestialBodyName(sr)
	return
}

func (sr OGRSpatialReference) SetAttrValue(nodePath, newValue string) (err error) {
	err = ogrError(osrSetAttrValue(sr, nodePath, newValue))
	return
}

func (sr OGRSpatialReference) GetAttrValue(name string, iChild int) (result string) {
	result = osrGetAttrValue(sr, name, iChild)
	return
}

func (sr OGRSpatialReference) SetAngularUnits(name string, inDegrees float64) (err error) {
	err = ogrError(osrSetAngularUnits(sr, name, inDegrees))
	return
}

func (sr OGRSpatialReference) GetAngularUnits() (result float64, name string) {
	result, name = osrGetAngularUnits(sr)
	return
}

func (sr OGRSpatialReference) SetLinearUnits(name string, inMeters float64) (err error) {
	err = ogrError(osrSetLinearUnits(sr, name, inMeters))
	return
}

func (sr OGRSpatialReference) SetTargetLinearUnits(targetKey, name string, inMeters float64) (err error) {
	err = ogrError(osrSetTargetLinearUnits(sr, targetKey, name, inMeters))
	return
}

func (sr OGRSpatialReference) SetLinearUnitsAndUpdateParameters(name string, inMeters float64) (err error) {
	err = ogrError(osrSetLinearUnitsAndUpdateParameters(sr, name, inMeters))
	return
}

func (sr OGRSpatialReference) GetLinearUnits() (result float64, name string) {
	result, name = osrGetLinearUnits(sr)
	return
}

func (sr OGRSpatialReference) GetTargetLinearUnits(targetKey string) (result float64, name string) {
	result, name = osrGetTargetLinearUnits(sr, targetKey)
	return
}

func (sr OGRSpatialReference) GetPrimeMeridian() (result float64, name string) {
	result, name = osrGetPrimeMeridian(sr)
	return
}

func (sr OGRSpatialReference) IsGeographic() (result bool) {
	result = osrIsGeographic(sr)
	return
}

func (sr OGRSpatialReference) IsDerivedGeographic() (result bool) {
	result = osrIsDerivedGeographic(sr)
	return
}

func (sr OGRSpatialReference) IsLocal() (result bool) {
	result = osrIsLocal(sr)
	return
}

func (sr OGRSpatialReference) IsProjected() (result bool) {
	result = osrIsProjected(sr)
	return
}

func (sr OGRSpatialReference) IsDerivedProjected() (result bool) {
	result = osrIsDerivedProjected(sr)
	return
}

func (sr OGRSpatialReference) IsCompound() (result bool) {
	result = osrIsCompound(sr)
	return
}

func (sr OGRSpatialReference) IsGeocentric() (result bool) {
	result = osrIsGeocentric(sr)
	return
}

func (sr OGRSpatialReference) IsVertical() (result bool) {
	result = osrIsVertical(sr)
	return
}

func (sr OGRSpatialReference) IsDynamic() (result bool) {
	result = osrIsDynamic(sr)
	return
}

func (sr OGRSpatialReference) HasPointMotionOperation() (result bool) {
	result = osrHasPointMotionOperation(sr)
	return
}

func (sr OGRSpatialReference) IsSameGeogCS(other OGRSpatialReference) (result bool) {
	result = osrIsSameGeogCS(sr, other)
	return
}

func (sr OGRSpatialReference) IsSameVertCS(other OGRSpatialReference) (result bool) {
	result = osrIsSameVertCS(sr, other)
	return
}

func (sr OGRSpatialReference) IsSame(other OGRSpatialReference) (result bool) {
	result = osrIsSame(sr, other)
	return
}

func (sr OGRSpatialReference) IsSameEx(other OGRSpatialReference, options CSLConstList) (result bool) {
	result = osrIsSameEx(sr, other, options)
	return
}

func (sr OGRSpatialReference) SetCoordinateEpoch(epoch float64) {
	osrSetCoordinateEpoch(sr, epoch)
}

func (sr OGRSpatialReference) GetCoordinateEpoch() (result float64) {
	result = osrGetCoordinateEpoch(sr)
	return
}

func (sr OGRSpatialReference) SetLocalCS(name string) (err error) {
	err = ogrError(osrSetLocalCS(sr, name))
	return
}

func (sr OGRSpatialReference) SetProjCS(name string) (err error) {
	err = ogrError(osrSetProjCS(sr, name))
	return
}

func (sr OGRSpatialReference) SetGeocCS(name string) (err error) {
	err = ogrError(osrSetGeocCS(sr, name))
	return
}

func (sr OGRSpatialReference) SetWellKnownGeogCS(name string) (err error) {
	err = ogrError(osrSetWellKnownGeogCS(sr, name))
	return
}

func (sr OGRSpatialReference) SetFromUserInput(definition string) (err error) {
	err = ogrError(osrSetFromUserInput(sr, definition))
	return
}

func (sr OGRSpatialReference) SetFromUserInputEx(definition string, options CSLConstList) (err error) {
	err = ogrError(osrSetFromUserInputEx(sr, definition, options))
	return
}

func (sr OGRSpatialReference) CopyGeogCSFrom(src OGRSpatialReference) (err error) {
	err = ogrError(osrCopyGeogCSFrom(sr, src))
	return
}

func (sr OGRSpatialReference) SetTOWGS84(dx, dy, dz, ex, ey, ez, ppm float64) (err error) {
	err = ogrError(osrSetTOWGS84(sr, dx, dy, dz, ex, ey, ez, ppm))
	return
}

func (sr OGRSpatialReference) GetTOWGS84() (params []float64, err error) {
	var status OGRErr
	params, status = osrGetTOWGS84(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) AddGuessedTOWGS84() (err error) {
	err = ogrError(osrAddGuessedTOWGS84(sr))
	return
}

func (sr OGRSpatialReference) SetCompoundCS(name string, horizSRS, vertSRS OGRSpatialReference) (err error) {
	err = ogrError(osrSetCompoundCS(sr, name, horizSRS, vertSRS))
	return
}

func (sr OGRSpatialReference) PromoteTo3D(name string) (err error) {
	err = ogrError(osrPromoteTo3D(sr, name))
	return
}

func (sr OGRSpatialReference) DemoteTo2D(name string) (err error) {
	err = ogrError(osrDemoteTo2D(sr, name))
	return
}

func (sr OGRSpatialReference) SetGeogCS(geogName, datumName, ellipsoidName string, semiMajor, invFlattening float64, pmName string, pmOffset float64, units string, convertToRadians float64) (err error) {
	err = ogrError(osrSetGeogCS(sr, geogName, datumName, ellipsoidName, semiMajor, invFlattening, pmName, pmOffset, units, convertToRadians))
	return
}

func (sr OGRSpatialReference) SetVertCS(vertCSName, vertDatumName string, vertDatumType int) (err error) {
	err = ogrError(osrSetVertCS(sr, vertCSName, vertDatumName, vertDatumType))
	return
}

func (sr OGRSpatialReference) GetSemiMajor() (result float64, err error) {
	var status OGRErr
	result, status = osrGetSemiMajor(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) GetSemiMinor() (result float64, err error) {
	var status OGRErr
	result, status = osrGetSemiMinor(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) GetInvFlattening() (result float64, err error) {
	var status OGRErr
	result, status = osrGetInvFlattening(sr)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) SetAuthority(targetKey, authority string, code int) (err error) {
	err = ogrError(osrSetAuthority(sr, targetKey, authority, code))
	return
}

func (sr OGRSpatialReference) GetAuthorityCode(targetKey string) (result string) {
	result = osrGetAuthorityCode(sr, targetKey)
	return
}

func (sr OGRSpatialReference) GetAuthorityName(targetKey string) (result string) {
	result = osrGetAuthorityName(sr, targetKey)
	return
}

func (sr OGRSpatialReference) GetAreaOfUse() (westLon, southLat, eastLon, northLat float64, areaName string, result bool) {
	westLon, southLat, eastLon, northLat, areaName, result = osrGetAreaOfUse(sr)
	return
}

func (sr OGRSpatialReference) SetProjection(projection string) (err error) {
	err = ogrError(osrSetProjection(sr, projection))
	return
}

func (sr OGRSpatialReference) SetProjParm(name string, value float64) (err error) {
	err = ogrError(osrSetProjParm(sr, name, value))
	return
}

func (sr OGRSpatialReference) GetProjParm(name string, dfDefault float64) (result float64, err error) {
	var status OGRErr
	result, status = osrGetProjParm(sr, name, dfDefault)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) SetNormProjParm(name string, value float64) (err error) {
	err = ogrError(osrSetNormProjParm(sr, name, value))
	return
}

func (sr OGRSpatialReference) GetNormProjParm(name string, dfDefault float64) (result float64, err error) {
	var status OGRErr
	result, status = osrGetNormProjParm(sr, name, dfDefault)
	err = ogrError(status)
	return
}

func (sr OGRSpatialReference) SetUTM(zone int, north bool) (err error) {
	bNorth := 0
	if north {
		bNorth = 1
	}
	err = ogrError(osrSetUTM(sr, zone, bNorth))
	return
}

func (sr OGRSpatialReference) GetUTMZone() (zone int, north bool) {
	zone, north = osrGetUTMZone(sr)
	return
}

func (sr OGRSpatialReference) SetStatePlane(zone int, nad83 bool) (err error) {
	bNad83 := 0
	if nad83 {
		bNad83 = 1
	}
	err = ogrError(osrSetStatePlane(sr, zone, bNad83))
	return
}

func (sr OGRSpatialReference) SetStatePlaneWithUnits(zone int, nad83 bool, unitName string, unit float64) (err error) {
	bNad83 := 0
	if nad83 {
		bNad83 = 1
	}
	err = ogrError(osrSetStatePlaneWithUnits(sr, zone, bNad83, unitName, unit))
	return
}

func (sr OGRSpatialReference) AutoIdentifyEPSG() (err error) {
	err = ogrError(osrAutoIdentifyEPSG(sr))
	return
}

func (sr OGRSpatialReference) FindMatches(options CSLConstList) (result []OGRSpatialReference, matchConfidence []int) {
	var count int
	list := osrFindMatches(sr, options, &count, &matchConfidence)
	if list.cValue == nil || count == 0 {
		return
	}
	src := unsafe.Slice(list.cValue, count)
	result = make([]OGRSpatialReference, count)
	for i := range result {
		result[i] = OGRSpatialReference{cValue: src[i]}
	}
	return
}

func (srs OGRSpatialReferences) FreeSRSArray() {
	osrFreeSRSArray(srs)
}

func (sr OGRSpatialReference) EPSGTreatsAsLatLong() (result bool) {
	result = osrEPSGTreatsAsLatLong(sr)
	return
}

func (sr OGRSpatialReference) EPSGTreatsAsNorthingEasting() (result bool) {
	result = osrEPSGTreatsAsNorthingEasting(sr)
	return
}

func (sr OGRSpatialReference) GetAxis(targetKey string, iAxis int) (name string, orientation OGRAxisOrientation) {
	name, orientation = osrGetAxis(sr, targetKey, iAxis)
	return
}

func (sr OGRSpatialReference) GetAxesCount() (result int) {
	result = osrGetAxesCount(sr)
	return
}

func (sr OGRSpatialReference) SetAxes(targetKey, xAxisName string, xOrientation OGRAxisOrientation, yAxisName string, yOrientation OGRAxisOrientation) (err error) {
	err = ogrError(osrSetAxes(sr, targetKey, xAxisName, xOrientation, yAxisName, yOrientation))
	return
}

func (sr OGRSpatialReference) GetAxisMappingStrategy() (result OSRAxisMappingStrategy) {
	result = osrGetAxisMappingStrategy(sr)
	return
}

func (sr OGRSpatialReference) SetAxisMappingStrategy(strategy OSRAxisMappingStrategy) {
	osrSetAxisMappingStrategy(sr, strategy)
}

func (sr OGRSpatialReference) GetDataAxisToSRSAxisMapping() (result []int) {
	result = osrGetDataAxisToSRSAxisMapping(sr)
	return
}

func (sr OGRSpatialReference) SetDataAxisToSRSAxisMapping(mapping []int) (err error) {
	err = ogrError(osrSetDataAxisToSRSAxisMapping(sr, mapping))
	return
}

func (sr OGRSpatialReference) SetACEA(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetACEA(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetAE(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetAE(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetBonne(dfStandardParallel, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetBonne(sr, dfStandardParallel, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetCEA(dfStdP1, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetCEA(sr, dfStdP1, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetCS(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetCS(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetEC(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetEC(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetEckert(nVariation int, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetEckert(sr, nVariation, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetEckertIV(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetEckertIV(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetEckertVI(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetEckertVI(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetEquirectangular(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetEquirectangular(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetEquirectangular2(dfCenterLat, dfCenterLong, dfPseudoStdParallel1, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetEquirectangular2(sr, dfCenterLat, dfCenterLong, dfPseudoStdParallel1, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetGS(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetGS(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetGH(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetGH(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetIGH() (err error) {
	err = ogrError(osrSetIGH(sr))
	return
}

func (sr OGRSpatialReference) SetGEOS(dfCentralMeridian, dfSatelliteHeight, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetGEOS(sr, dfCentralMeridian, dfSatelliteHeight, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetGaussSchreiberTMercator(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetGaussSchreiberTMercator(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetGnomonic(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetGnomonic(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetHOM(dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetHOM(sr, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetHOMAC(dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetHOMAC(sr, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetHOM2PNO(dfCenterLat, dfLat1, dfLong1, dfLat2, dfLong2, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetHOM2PNO(sr, dfCenterLat, dfLat1, dfLong1, dfLat2, dfLong2, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetIWMPolyconic(dfLat1, dfLat2, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetIWMPolyconic(sr, dfLat1, dfLat2, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetKrovak(dfCenterLat, dfCenterLong, dfAzimuth, dfPseudoStdParallelLat, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetKrovak(sr, dfCenterLat, dfCenterLong, dfAzimuth, dfPseudoStdParallelLat, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetLAEA(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetLAEA(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetLCC(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetLCC(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetLCC1SP(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetLCC1SP(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetLCCB(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetLCCB(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetMC(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetMC(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetMercator(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetMercator(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetMercator2SP(dfStdP1, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetMercator2SP(sr, dfStdP1, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetMollweide(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetMollweide(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetNZMG(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetNZMG(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetOS(dfOriginLat, dfCMeridian, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetOS(sr, dfOriginLat, dfCMeridian, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetOrthographic(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetOrthographic(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetPolyconic(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetPolyconic(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetPS(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetPS(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetRobinson(dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetRobinson(sr, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetSinusoidal(dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetSinusoidal(sr, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetStereographic(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetStereographic(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetSOC(dfLatitudeOfOrigin, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetSOC(sr, dfLatitudeOfOrigin, dfCentralMeridian, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetTM(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetTM(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetTMVariant(pszVariantName string, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetTMVariant(sr, pszVariantName, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetTMG(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetTMG(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetTMSO(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetTMSO(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetTPED(dfLat1, dfLong1, dfLat2, dfLong2, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetTPED(sr, dfLat1, dfLong1, dfLat2, dfLong2, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetVDG(dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetVDG(sr, dfCenterLong, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetWagner(nVariation int, dfCenterLat, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetWagner(sr, nVariation, dfCenterLat, dfFalseEasting, dfFalseNorthing))
	return
}

func (sr OGRSpatialReference) SetQSC(dfCenterLat, dfCenterLong float64) (err error) {
	err = ogrError(osrSetQSC(sr, dfCenterLat, dfCenterLong))
	return
}

func (sr OGRSpatialReference) SetSCH(dfPegLat, dfPegLong, dfPegHeading, dfPegHgt float64) (err error) {
	err = ogrError(osrSetSCH(sr, dfPegLat, dfPegLong, dfPegHeading, dfPegHgt))
	return
}

func (sr OGRSpatialReference) SetVerticalPerspective(dfTopoOriginLat, dfTopoOriginLon, dfTopoOriginHeight, dfViewPointHeight, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(osrSetVerticalPerspective(sr, dfTopoOriginLat, dfTopoOriginLon, dfTopoOriginHeight, dfViewPointHeight, dfFalseEasting, dfFalseNorthing))
	return
}

func OSRCalcInvFlattening(semiMajor, semiMinor float64) (result float64) {
	result = osrCalcInvFlattening(semiMajor, semiMinor)
	return
}

func OSRCalcSemiMinorFromInvFlattening(semiMajor, invFlattening float64) (result float64) {
	result = osrCalcSemiMinorFromInvFlattening(semiMajor, invFlattening)
	return
}

func OSRCleanup() {
	osrCleanup()
}

func OSRGetCRSInfoListFromDatabase(authName string, params OSRCRSListParameters) (result []OSRCRSInfo, err error) {
	var count int
	list := osrGetCRSInfoListFromDatabase(authName, params, &count)
	if list.cValue == nil {
		err = lastError()
		return
	}
	src := unsafe.Slice(list.cValue, count)
	result = make([]OSRCRSInfo, count)
	for i := range result {
		result[i] = OSRCRSInfo{cValue: src[i]}
	}
	return
}

func OSRDestroyCRSInfoList(list OSRCRSInfos) {
	osrDestroyCRSInfoList(list)
}

func OSRGetAuthorityListFromDatabase() (result CSLConstList) {
	result = osrGetAuthorityListFromDatabase()
	return
}

func OCTNewCoordinateTransformation(source, target OGRSpatialReference) (result OGRCoordinateTransformation, err error) {
	result = octNewCoordinateTransformation(source, target)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func OCTNewCoordinateTransformationOptions() (result OGRCoordinateTransformationOptions, err error) {
	result = octNewCoordinateTransformationOptions()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (opts OGRCoordinateTransformationOptions) SetOperation(pszCO string, reverseCO bool) (result bool) {
	bReverseCO := 0
	if reverseCO {
		bReverseCO = 1
	}
	result = octCoordinateTransformationOptionsSetOperation(opts, pszCO, bReverseCO)
	return
}

func (opts OGRCoordinateTransformationOptions) SetAreaOfInterest(dfWestLongitudeDeg, dfSouthLatitudeDeg, dfEastLongitudeDeg, dfNorthLatitudeDeg float64) (result bool) {
	result = octCoordinateTransformationOptionsSetAreaOfInterest(opts, dfWestLongitudeDeg, dfSouthLatitudeDeg, dfEastLongitudeDeg, dfNorthLatitudeDeg)
	return
}

func (opts OGRCoordinateTransformationOptions) SetDesiredAccuracy(dfAccuracy float64) (result bool) {
	result = octCoordinateTransformationOptionsSetDesiredAccuracy(opts, dfAccuracy)
	return
}

func (opts OGRCoordinateTransformationOptions) SetBallparkAllowed(allowBallpark bool) (result bool) {
	bAllowBallpark := 0
	if allowBallpark {
		bAllowBallpark = 1
	}
	result = octCoordinateTransformationOptionsSetBallparkAllowed(opts, bAllowBallpark)
	return
}

func (opts OGRCoordinateTransformationOptions) SetOnlyBest(bOnlyBest bool) (result bool) {
	result = octCoordinateTransformationOptionsSetOnlyBest(opts, bOnlyBest)
	return
}

func (opts OGRCoordinateTransformationOptions) Destroy() {
	octDestroyCoordinateTransformationOptions(opts)
}

func OCTNewCoordinateTransformationEx(source, target OGRSpatialReference, options OGRCoordinateTransformationOptions) (result OGRCoordinateTransformation, err error) {
	result = octNewCoordinateTransformationEx(source, target, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ct OGRCoordinateTransformation) Clone() (result OGRCoordinateTransformation, err error) {
	result = octClone(ct)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ct OGRCoordinateTransformation) GetSourceCS() (result OGRSpatialReference) {
	result = octGetSourceCS(ct)
	return
}

func (ct OGRCoordinateTransformation) GetTargetCS() (result OGRSpatialReference) {
	result = octGetTargetCS(ct)
	return
}

func (ct OGRCoordinateTransformation) GetInverse() (result OGRCoordinateTransformation, err error) {
	result = octGetInverse(ct)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ct OGRCoordinateTransformation) Destroy() {
	octDestroyCoordinateTransformation(ct)
}

func (ct OGRCoordinateTransformation) Transform(x, y, z []float64) (result bool) {
	result = octTransform(ct, x, y, z)
	return
}

func (ct OGRCoordinateTransformation) TransformEx(x, y, z []float64) (success []bool, result bool) {
	success, result = octTransformEx(ct, x, y, z)
	return
}

func (ct OGRCoordinateTransformation) Transform4D(x, y, z, t []float64) (success []bool, result bool) {
	success, result = octTransform4D(ct, x, y, z, t)
	return
}

func (ct OGRCoordinateTransformation) Transform4DWithErrorCodes(x, y, z, t []float64) (errorCodes []int, result bool) {
	errorCodes, result = octTransform4DWithErrorCodes(ct, x, y, z, t)
	return
}

func (ct OGRCoordinateTransformation) TransformBounds(xmin, ymin, xmax, ymax float64, densifyPts int) (outXmin, outYmin, outXmax, outYmax float64, result bool) {
	outXmin, outYmin, outXmax, outYmax, result = octTransformBounds(ct, xmin, ymin, xmax, ymax, densifyPts)
	return
}
