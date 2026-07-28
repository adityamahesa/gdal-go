package gdal

import "unsafe"

func OGRGetGEOSVersion() (major, minor, patch int) {
	major, minor, patch = ogrGetGEOSVersion()
	return
}

func OGRGeomCoordinatePrecisionCreate() (result OGRGeomCoordinatePrecision, err error) {
	scope := errScope()
	defer scope()
	result = ogrGeomCoordinatePrecisionCreate()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (p OGRGeomCoordinatePrecision) Destroy() {
	ogrGeomCoordinatePrecisionDestroy(p)
}

func (p OGRGeomCoordinatePrecision) GetXYResolution() (result float64) {
	result = ogrGeomCoordinatePrecisionGetXYResolution(p)
	return
}

func (p OGRGeomCoordinatePrecision) GetZResolution() (result float64) {
	result = ogrGeomCoordinatePrecisionGetZResolution(p)
	return
}

func (p OGRGeomCoordinatePrecision) GetMResolution() (result float64) {
	result = ogrGeomCoordinatePrecisionGetMResolution(p)
	return
}

func (p OGRGeomCoordinatePrecision) GetFormats() (result CSLConstList) {
	result = ogrGeomCoordinatePrecisionGetFormats(p)
	return
}

func (p OGRGeomCoordinatePrecision) GetFormatSpecificOptions(formatName string) (result CSLConstList) {
	result = ogrGeomCoordinatePrecisionGetFormatSpecificOptions(p, formatName)
	return
}

func (p OGRGeomCoordinatePrecision) Set(xyResolution, zResolution, mResolution float64) {
	ogrGeomCoordinatePrecisionSet(p, xyResolution, zResolution, mResolution)
}

func (p OGRGeomCoordinatePrecision) SetFromMeter(sr OGRSpatialReference, xyMeterResolution, zMeterResolution, mResolution float64) {
	ogrGeomCoordinatePrecisionSetFromMeter(p, sr, xyMeterResolution, zMeterResolution, mResolution)
}

func (p OGRGeomCoordinatePrecision) SetFormatSpecificOptions(formatName string, options CSLConstList) {
	ogrGeomCoordinatePrecisionSetFormatSpecificOptions(p, formatName, options)
}

func OGRGCreateFromWkb(data []byte, sr OGRSpatialReference) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGCreateFromWkb(data, sr)
	err = ogrError(status)
	return
}

func OGRGCreateFromWkbEx(data []byte, sr OGRSpatialReference) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGCreateFromWkbEx(data, sr)
	err = ogrError(status)
	return
}

func OGRGCreateFromWkt(wkt string, sr OGRSpatialReference) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGCreateFromWkt(wkt, sr)
	err = ogrError(status)
	return
}

func OGRGCreateFromFgf(data []byte, sr OGRSpatialReference) (result OGRGeometry, bytesConsumed int, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, bytesConsumed, status = ogrGCreateFromFgf(data, sr)
	err = ogrError(status)
	return
}

func OGRGCreateFromEnvelope(minX, maxX, minY, maxY float64, sr OGRSpatialReference) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGCreateFromEnvelope(minX, maxX, minY, maxY, sr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Destroy() {
	ogrGDestroyGeometry(g)
}

func OGRGCreateGeometry(eType OGRwkbGeometryType) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGCreateGeometry(eType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func OGRGApproximateArcAngles(centerX, centerY, z, primaryRadius, secondaryAxis, rotation, startAngle, endAngle, maxAngleStepSizeDegrees float64) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGApproximateArcAngles(centerX, centerY, z, primaryRadius, secondaryAxis, rotation, startAngle, endAngle, maxAngleStepSizeDegrees)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ForceToPolygon() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGForceToPolygon(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ForceToLineString() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGForceToLineString(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ForceToMultiPolygon() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGForceToMultiPolygon(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ForceToMultiPoint() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGForceToMultiPoint(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ForceToMultiLineString() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGForceToMultiLineString(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ForceTo(eTargetType OGRwkbGeometryType, options CSLConstList) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGForceTo(g, eTargetType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) RemoveLowerDimensionSubGeoms() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGRemoveLowerDimensionSubGeoms(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) GetDimension() (result int) {
	result = ogrGGetDimension(g)
	return
}

func (g OGRGeometry) GetCoordinateDimension() (result int) {
	result = ogrGGetCoordinateDimension(g)
	return
}

func (g OGRGeometry) CoordinateDimension() (result int) {
	result = ogrGCoordinateDimension(g)
	return
}

func (g OGRGeometry) SetCoordinateDimension(dimension int) {
	ogrGSetCoordinateDimension(g, dimension)
}

func (g OGRGeometry) Is3D() (result bool) {
	result = ogrGIs3D(g) != 0
	return
}

func (g OGRGeometry) IsMeasured() (result bool) {
	result = ogrGIsMeasured(g) != 0
	return
}

func (g OGRGeometry) Set3D(is3D int) {
	ogrGSet3D(g, is3D)
}

func (g OGRGeometry) SetMeasured(isMeasured int) {
	ogrGSetMeasured(g, isMeasured)
}

func (g OGRGeometry) Clone() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGClone(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) GetEnvelope() (result OGREnvelope) {
	result = ogrGGetEnvelope(g)
	return
}

func (g OGRGeometry) GetEnvelope3D() (result OGREnvelope3D) {
	result = ogrGGetEnvelope3D(g)
	return
}

func (g OGRGeometry) ImportFromWkb(data []byte) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrGImportFromWkb(g, data))
	return
}

func (g OGRGeometry) ExportToWkb(order OGRwkbByteOrder) (result []byte, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGExportToWkb(g, order)
	err = ogrError(status)
	return
}

func (g OGRGeometry) ExportToIsoWkb(order OGRwkbByteOrder) (result []byte, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGExportToIsoWkb(g, order)
	err = ogrError(status)
	return
}

func OGRwkbExportOptionsCreate() (result OGRwkbExportOptions, err error) {
	scope := errScope()
	defer scope()
	result = ogrwkbExportOptionsCreate()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (o OGRwkbExportOptions) Destroy() {
	ogrwkbExportOptionsDestroy(o)
}

func (o OGRwkbExportOptions) SetByteOrder(order OGRwkbByteOrder) {
	ogrwkbExportOptionsSetByteOrder(o, order)
}

func (o OGRwkbExportOptions) SetVariant(variant OGRwkbVariant) {
	ogrwkbExportOptionsSetVariant(o, variant)
}

func (o OGRwkbExportOptions) SetPrecision(precision OGRGeomCoordinatePrecision) {
	ogrwkbExportOptionsSetPrecision(o, precision)
}

func (g OGRGeometry) ExportToWkbEx(opts OGRwkbExportOptions) (result []byte, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGExportToWkbEx(g, opts)
	err = ogrError(status)
	return
}

func (g OGRGeometry) WkbSize() (result int) {
	result = ogrGWkbSize(g)
	return
}

func (g OGRGeometry) WkbSizeEx() (result int) {
	result = ogrGWkbSizeEx(g)
	return
}

func (g OGRGeometry) ImportFromWkt(wkt string) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrGImportFromWkt(g, wkt))
	return
}

func (g OGRGeometry) ExportToWkt() (result string, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGExportToWkt(g)
	err = ogrError(status)
	return
}

func (g OGRGeometry) ExportToIsoWkt() (result string, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrGExportToIsoWkt(g)
	err = ogrError(status)
	return
}

func (g OGRGeometry) GetGeometryType() (result OGRwkbGeometryType) {
	result = ogrGGetGeometryType(g)
	return
}

func (g OGRGeometry) GetGeometryName() (result string) {
	result = ogrGGetGeometryName(g)
	return
}

func (g OGRGeometry) FlattenTo2D() {
	ogrGFlattenTo2D(g)
}

func (g OGRGeometry) CloseRings() {
	ogrGCloseRings(g)
}

func OGRGCreateFromGML(gml string) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGCreateFromGML(gml)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ExportToGML() (result string) {
	result = ogrGExportToGML(g)
	return
}

func (g OGRGeometry) ExportToGMLEx(options CSLConstList) (result string) {
	result = ogrGExportToGMLEx(g, options)
	return
}

func OGRGCreateFromGMLTree(tree CPLXMLNode) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGCreateFromGMLTree(tree)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ExportToGMLTree() (result CPLXMLNode, err error) {
	scope := errScope()
	defer scope()
	result = ogrGExportToGMLTree(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ExportEnvelopeToGMLTree() (result CPLXMLNode, err error) {
	scope := errScope()
	defer scope()
	result = ogrGExportEnvelopeToGMLTree(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ExportToKML(altitudeMode string) (result string) {
	result = ogrGExportToKML(g, altitudeMode)
	return
}

func (g OGRGeometry) ExportToJson() (result string) {
	result = ogrGExportToJson(g)
	return
}

func (g OGRGeometry) ExportToJsonEx(options CSLConstList) (result string) {
	result = ogrGExportToJsonEx(g, options)
	return
}

// /** Create a OGR geometry from a GeoJSON geometry object */
func OGRGCreateGeometryFromJson(json string) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGCreateGeometryFromJson(json)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// /** Create a OGR geometry from a ESRI JSON geometry object */
func OGRGCreateGeometryFromEsriJson(json string) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGCreateGeometryFromEsriJson(json)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) AssignSpatialReference(sr OGRSpatialReference) {
	ogrGAssignSpatialReference(g, sr)
}

func (g OGRGeometry) GetSpatialReference() (result OGRSpatialReference) {
	result = ogrGGetSpatialReference(g)
	return
}

func (g OGRGeometry) Transform(ct OGRCoordinateTransformation) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrGTransform(g, ct))
	return
}

func (g OGRGeometry) TransformTo(sr OGRSpatialReference) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrGTransformTo(g, sr))
	return
}

func OGRGeomTransformerCreate(ct OGRCoordinateTransformation, options CSLConstList) (result OGRGeomTransformer, err error) {
	scope := errScope()
	defer scope()
	result = ogrGeomTransformerCreate(ct, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (t OGRGeomTransformer) Transform(g OGRGeometry) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGeomTransformerTransform(t, g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (t OGRGeomTransformer) Destroy() {
	ogrGeomTransformerDestroy(t)
}

func (g OGRGeometry) Simplify(tolerance float64) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGSimplify(g, tolerance)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) SimplifyPreserveTopology(tolerance float64) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGSimplifyPreserveTopology(g, tolerance)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) DelaunayTriangulation(tolerance float64, onlyEdges int) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGDelaunayTriangulation(g, tolerance, onlyEdges)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ConstrainedDelaunayTriangulation() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGConstrainedDelaunayTriangulation(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Segmentize(maxLength float64) {
	ogrGSegmentize(g, maxLength)
}

func (g OGRGeometry) Intersects(other OGRGeometry) (result bool) {
	result = ogrGIntersects(g, other) != 0
	return
}

func (g OGRGeometry) Equals(other OGRGeometry) (result bool) {
	result = ogrGEquals(g, other) != 0
	return
}

func (g OGRGeometry) Disjoint(other OGRGeometry) (result bool) {
	result = ogrGDisjoint(g, other) != 0
	return
}

func (g OGRGeometry) Touches(other OGRGeometry) (result bool) {
	result = ogrGTouches(g, other) != 0
	return
}

func (g OGRGeometry) Crosses(other OGRGeometry) (result bool) {
	result = ogrGCrosses(g, other) != 0
	return
}

func (g OGRGeometry) Within(other OGRGeometry) (result bool) {
	result = ogrGWithin(g, other) != 0
	return
}

func (g OGRGeometry) Contains(other OGRGeometry) (result bool) {
	result = ogrGContains(g, other) != 0
	return
}

func (g OGRGeometry) Overlaps(other OGRGeometry) (result bool) {
	result = ogrGOverlaps(g, other) != 0
	return
}

func (g OGRGeometry) Boundary() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGBoundary(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ConvexHull() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGConvexHull(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) ConcaveHull(ratio float64, allowHoles bool) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGConcaveHull(g, ratio, allowHoles)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Buffer(dist float64, quadSegs int) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGBuffer(g, dist, quadSegs)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) BufferEx(dist float64, options CSLConstList) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGBufferEx(g, dist, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Intersection(other OGRGeometry) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGIntersection(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Union(other OGRGeometry) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGUnion(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) UnionCascaded() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGUnionCascaded(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) UnaryUnion() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGUnaryUnion(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) PointOnSurface() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGPointOnSurface(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Difference(other OGRGeometry) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGDifference(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) SymDifference(other OGRGeometry) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGSymDifference(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Distance(other OGRGeometry) (result float64) {
	result = ogrGDistance(g, other)
	return
}

func (g OGRGeometry) Distance3D(other OGRGeometry) (result float64) {
	result = ogrGDistance3D(g, other)
	return
}

func (g OGRGeometry) Length() (result float64) {
	result = ogrGLength(g)
	return
}

func (g OGRGeometry) GeodesicLength() (result float64) {
	result = ogrGGeodesicLength(g)
	return
}

func (g OGRGeometry) Area() (result float64) {
	result = ogrGArea(g)
	return
}

func (g OGRGeometry) GeodesicArea() (result float64) {
	result = ogrGGeodesicArea(g)
	return
}

func (g OGRGeometry) IsClockwise() (result bool) {
	result = ogrGIsClockwise(g)
	return
}

func (g OGRGeometry) Centroid() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGCreateGeometry(WkbPoint)
	err = ogrError(OGRErr(ogrGCentroid(g, result)))
	return
}

func (g OGRGeometry) Value(distance float64) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGValue(g, distance)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Empty() {
	ogrGEmpty(g)
}

func (g OGRGeometry) IsEmpty() (result bool) {
	result = ogrGIsEmpty(g) != 0
	return
}

func (g OGRGeometry) IsValid() (result bool) {
	result = ogrGIsValid(g) != 0
	return
}

func (g OGRGeometry) MakeValid() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGMakeValid(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) MakeValidEx(options CSLConstList) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGMakeValidEx(g, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Normalize() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGNormalize(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) IsSimple() (result bool) {
	result = ogrGIsSimple(g) != 0
	return
}

func (g OGRGeometry) IsRing() (result bool) {
	result = ogrGIsRing(g) != 0
	return
}

func (g OGRGeometry) SetPrecision(gridSize float64, flags int) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGSetPrecision(g, gridSize, flags)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) Polygonize() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGPolygonize(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) BuildArea() (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGBuildArea(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) GetPointCount() (result int) {
	result = ogrGGetPointCount(g)
	return
}

func (g OGRGeometry) GetPoints() (x, y, z []float64) {
	x, y, z = ogrGGetPoints(g)
	return
}

func (g OGRGeometry) GetPointsZM() (x, y, z, m []float64) {
	x, y, z, m = ogrGGetPointsZM(g)
	return
}

func (g OGRGeometry) GetX(point int) (result float64) {
	result = ogrGGetX(g, point)
	return
}

func (g OGRGeometry) GetY(point int) (result float64) {
	result = ogrGGetY(g, point)
	return
}

func (g OGRGeometry) GetZ(point int) (result float64) {
	result = ogrGGetZ(g, point)
	return
}

func (g OGRGeometry) GetM(point int) (result float64) {
	result = ogrGGetM(g, point)
	return
}

func (g OGRGeometry) GetPoint(point int) (x, y, z float64) {
	x, y, z = ogrGGetPoint(g, point)
	return
}

func (g OGRGeometry) GetPointZM(point int) (x, y, z, m float64) {
	x, y, z, m = ogrGGetPointZM(g, point)
	return
}

func (g OGRGeometry) SetPointCount(count int) {
	ogrGSetPointCount(g, count)
}

func (g OGRGeometry) SetPoint(point int, x, y, z float64) {
	ogrGSetPoint(g, point, x, y, z)
}

func (g OGRGeometry) SetPoint2D(point int, x, y float64) {
	ogrGSetPoint2D(g, point, x, y)
}

func (g OGRGeometry) SetPointM(point int, x, y, m float64) {
	ogrGSetPointM(g, point, x, y, m)
}

func (g OGRGeometry) SetPointZM(point int, x, y, z, m float64) {
	ogrGSetPointZM(g, point, x, y, z, m)
}

func (g OGRGeometry) AddPoint(x, y, z float64) {
	ogrGAddPoint(g, x, y, z)
}

func (g OGRGeometry) AddPoint2D(x, y float64) {
	ogrGAddPoint2D(g, x, y)
}

func (g OGRGeometry) AddPointM(x, y, m float64) {
	ogrGAddPointM(g, x, y, m)
}

func (g OGRGeometry) AddPointZM(x, y, z, m float64) {
	ogrGAddPointZM(g, x, y, z, m)
}

func (g OGRGeometry) SetPoints(x, y, z []float64) {
	ogrGSetPoints(g, x, y, z)
}

func (g OGRGeometry) SetPointsZM(x, y, z, m []float64) {
	ogrGSetPointsZM(g, x, y, z, m)
}

func (g OGRGeometry) SwapXY() {
	ogrGSwapXY(g)
}

func (g OGRGeometry) GetGeometryCount() (result int) {
	result = ogrGGetGeometryCount(g)
	return
}

func (g OGRGeometry) GetGeometryRef(subGeom int) (result OGRGeometry) {
	result = ogrGGetGeometryRef(g, subGeom)
	return
}

func (g OGRGeometry) AddGeometry(subGeom OGRGeometry) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrGAddGeometry(g, subGeom))
	return
}

func (g OGRGeometry) AddGeometryDirectly(subGeom OGRGeometry) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrGAddGeometryDirectly(g, subGeom))
	return
}

func (g OGRGeometry) RemoveGeometry(subGeom, delete int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrGRemoveGeometry(g, subGeom, delete))
	return
}

func (g OGRGeometry) HasCurveGeometry(lookForNonLinear int) (result bool) {
	result = ogrGHasCurveGeometry(g, lookForNonLinear) != 0
	return
}

func (g OGRGeometry) GetLinearGeometry(maxAngleStepSizeDegrees float64, options CSLConstList) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGGetLinearGeometry(g, maxAngleStepSizeDegrees, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (g OGRGeometry) GetCurveGeometry(options CSLConstList) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrGGetCurveGeometry(g, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func OGRBuildPolygonFromEdges(lines OGRGeometry, bestEffort, autoClose int, tolerance float64) (result OGRGeometry, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrBuildPolygonFromEdges(lines, bestEffort, autoClose, tolerance)
	err = ogrError(status)
	return
}

func OGRSetGenerateDB2V72ByteOrder(generate int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrSetGenerateDB2V72ByteOrder(generate))
	return
}

func OGRGetGenerateDB2V72ByteOrder() (result int) {
	result = ogrGetGenerateDB2V72ByteOrder()
	return
}

func OGRSetNonLinearGeometriesEnabledFlag(flag int) {
	ogrSetNonLinearGeometriesEnabledFlag(flag)
}

func OGRGetNonLinearGeometriesEnabledFlag() (result int) {
	result = ogrGetNonLinearGeometriesEnabledFlag()
	return
}

func OGRHasPreparedGeometrySupport() (result bool) {
	result = ogrHasPreparedGeometrySupport()
	return
}

func OGRCreatePreparedGeometry(g OGRGeometry) (result OGRPreparedGeometry, err error) {
	scope := errScope()
	defer scope()
	result = ogrCreatePreparedGeometry(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (p OGRPreparedGeometry) Destroy() {
	ogrDestroyPreparedGeometry(p)
}

func (p OGRPreparedGeometry) Intersects(other OGRGeometry) (result bool) {
	result = ogrPreparedGeometryIntersects(p, other) != 0
	return
}

func (p OGRPreparedGeometry) Contains(other OGRGeometry) (result bool) {
	result = ogrPreparedGeometryContains(p, other) != 0
	return
}

func OGRFldCreate(name string, eType OGRFieldType) (result OGRFieldDefn, err error) {
	scope := errScope()
	defer scope()
	result = ogrFldCreate(name, eType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (fld OGRFieldDefn) Destroy() {
	ogrFldDestroy(fld)
}

func (fld OGRFieldDefn) SetName(name string) {
	ogrFldSetName(fld, name)
}

func (fld OGRFieldDefn) GetNameRef() (result string) {
	result = ogrFldGetNameRef(fld)
	return
}

func (fld OGRFieldDefn) SetAlternativeName(name string) {
	ogrFldSetAlternativeName(fld, name)
}

func (fld OGRFieldDefn) GetAlternativeNameRef() (result string) {
	result = ogrFldGetAlternativeNameRef(fld)
	return
}

func (fld OGRFieldDefn) GetType() (result OGRFieldType) {
	result = ogrFldGetType(fld)
	return
}

func (fld OGRFieldDefn) SetType(eType OGRFieldType) {
	ogrFldSetType(fld, eType)
}

func (fld OGRFieldDefn) GetSubType() (result OGRFieldSubType) {
	result = ogrFldGetSubType(fld)
	return
}

func (fld OGRFieldDefn) SetSubType(eSubType OGRFieldSubType) {
	ogrFldSetSubType(fld, eSubType)
}

func (fld OGRFieldDefn) GetJustify() (result OGRJustification) {
	result = ogrFldGetJustify(fld)
	return
}

func (fld OGRFieldDefn) SetJustify(eJustify OGRJustification) {
	ogrFldSetJustify(fld, eJustify)
}

func (fld OGRFieldDefn) GetWidth() (result int) {
	result = ogrFldGetWidth(fld)
	return
}

func (fld OGRFieldDefn) SetWidth(width int) {
	ogrFldSetWidth(fld, width)
}

func (fld OGRFieldDefn) GetPrecision() (result int) {
	result = ogrFldGetPrecision(fld)
	return
}

func (fld OGRFieldDefn) SetPrecision(precision int) {
	ogrFldSetPrecision(fld, precision)
}

func (fld OGRFieldDefn) GetTZFlag() (result int) {
	result = ogrFldGetTZFlag(fld)
	return
}

func (fld OGRFieldDefn) SetTZFlag(tzFlag int) {
	ogrFldSetTZFlag(fld, tzFlag)
}

func (fld OGRFieldDefn) Set(name string, eType OGRFieldType, width, precision int, justify OGRJustification) {
	ogrFldSet(fld, name, eType, width, precision, justify)
}

func (fld OGRFieldDefn) IsIgnored() (result bool) {
	result = ogrFldIsIgnored(fld) != 0
	return
}

func (fld OGRFieldDefn) SetIgnored(ignored int) {
	ogrFldSetIgnored(fld, ignored)
}

func (fld OGRFieldDefn) IsNullable() (result bool) {
	result = ogrFldIsNullable(fld) != 0
	return
}

func (fld OGRFieldDefn) SetNullable(nullable int) {
	ogrFldSetNullable(fld, nullable)
}

func (fld OGRFieldDefn) SetGenerated(generated int) {
	ogrFldSetGenerated(fld, generated)
}

func (fld OGRFieldDefn) IsGenerated() (result bool) {
	result = ogrFldIsGenerated(fld) != 0
	return
}

func (fld OGRFieldDefn) IsUnique() (result bool) {
	result = ogrFldIsUnique(fld) != 0
	return
}

func (fld OGRFieldDefn) SetUnique(unique int) {
	ogrFldSetUnique(fld, unique)
}

func (fld OGRFieldDefn) GetDefault() (result string) {
	result = ogrFldGetDefault(fld)
	return
}

func (fld OGRFieldDefn) SetDefault(value string) {
	ogrFldSetDefault(fld, value)
}

func (fld OGRFieldDefn) IsDefaultDriverSpecific() (result bool) {
	result = ogrFldIsDefaultDriverSpecific(fld) != 0
	return
}

func (fld OGRFieldDefn) GetDomainName() (result string) {
	result = ogrFldGetDomainName(fld)
	return
}

func (fld OGRFieldDefn) SetDomainName(name string) {
	ogrFldSetDomainName(fld, name)
}

func (fld OGRFieldDefn) GetComment() (result string) {
	result = ogrFldGetComment(fld)
	return
}

func (fld OGRFieldDefn) SetComment(comment string) {
	ogrFldSetComment(fld, comment)
}

func OGRGetFieldTypeName(eType OGRFieldType) (result string) {
	result = ogrGetFieldTypeName(eType)
	return
}

func OGRGetFieldTypeByName(name string) (result OGRFieldType) {
	result = ogrGetFieldTypeByName(name)
	return
}

func OGRGetFieldSubTypeName(eSubType OGRFieldSubType) (result string) {
	result = ogrGetFieldSubTypeName(eSubType)
	return
}

func OGRGetFieldSubTypeByName(name string) (result OGRFieldSubType) {
	result = ogrGetFieldSubTypeByName(name)
	return
}

func OGRAreTypeSubTypeCompatible(eType OGRFieldType, eSubType OGRFieldSubType) (result bool) {
	result = ogrAreTypeSubTypeCompatible(eType, eSubType) != 0
	return
}

func OGRGFldCreate(name string, eType OGRwkbGeometryType) (result OGRGeomFieldDefn, err error) {
	scope := errScope()
	defer scope()
	result = ogrGFldCreate(name, eType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (gfld OGRGeomFieldDefn) Destroy() {
	ogrGFldDestroy(gfld)
}

func (gfld OGRGeomFieldDefn) SetName(name string) {
	ogrGFldSetName(gfld, name)
}

func (gfld OGRGeomFieldDefn) GetNameRef() (result string) {
	result = ogrGFldGetNameRef(gfld)
	return
}

func (gfld OGRGeomFieldDefn) GetType() (result OGRwkbGeometryType) {
	result = ogrGFldGetType(gfld)
	return
}

func (gfld OGRGeomFieldDefn) SetType(eType OGRwkbGeometryType) {
	ogrGFldSetType(gfld, eType)
}

func (gfld OGRGeomFieldDefn) GetSpatialRef() (result OGRSpatialReference) {
	result = ogrGFldGetSpatialRef(gfld)
	return
}

func (gfld OGRGeomFieldDefn) SetSpatialRef(sr OGRSpatialReference) {
	ogrGFldSetSpatialRef(gfld, sr)
}

func (gfld OGRGeomFieldDefn) IsNullable() (result bool) {
	result = ogrGFldIsNullable(gfld) != 0
	return
}

func (gfld OGRGeomFieldDefn) SetNullable(nullable int) {
	ogrGFldSetNullable(gfld, nullable)
}

func (gfld OGRGeomFieldDefn) IsIgnored() (result bool) {
	result = ogrGFldIsIgnored(gfld) != 0
	return
}

func (gfld OGRGeomFieldDefn) SetIgnored(ignored int) {
	ogrGFldSetIgnored(gfld, ignored)
}

func (gfld OGRGeomFieldDefn) GetCoordinatePrecision() (result OGRGeomCoordinatePrecision) {
	result = ogrGFldGetCoordinatePrecision(gfld)
	return
}

func (gfld OGRGeomFieldDefn) SetCoordinatePrecision(precision OGRGeomCoordinatePrecision) {
	ogrGFldSetCoordinatePrecision(gfld, precision)
}

func OGRFDCreate(name string) (result OGRFeatureDefn, err error) {
	scope := errScope()
	defer scope()
	result = ogrFDCreate(name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (fd OGRFeatureDefn) Destroy() {
	ogrFDDestroy(fd)
}

func (fd OGRFeatureDefn) Release() {
	ogrFDRelease(fd)
}

func (fd OGRFeatureDefn) GetName() (result string) {
	result = ogrFDGetName(fd)
	return
}

func (fd OGRFeatureDefn) GetFieldCount() (result int) {
	result = ogrFDGetFieldCount(fd)
	return
}

func (fd OGRFeatureDefn) GetFieldDefn(field int) (result OGRFieldDefn) {
	result = ogrFDGetFieldDefn(fd, field)
	return
}

func (fd OGRFeatureDefn) GetFieldIndex(name string) (result int) {
	result = ogrFDGetFieldIndex(fd, name)
	return
}

func (fd OGRFeatureDefn) AddFieldDefn(fld OGRFieldDefn) {
	ogrFDAddFieldDefn(fd, fld)
}

func (fd OGRFeatureDefn) DeleteFieldDefn(field int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFDDeleteFieldDefn(fd, field))
	return
}

func (fd OGRFeatureDefn) ReorderFieldDefns(panMap []int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFDReorderFieldDefns(fd, panMap))
	return
}

func (fd OGRFeatureDefn) GetGeomType() (result OGRwkbGeometryType) {
	result = ogrFDGetGeomType(fd)
	return
}

func (fd OGRFeatureDefn) SetGeomType(eType OGRwkbGeometryType) {
	ogrFDSetGeomType(fd, eType)
}

func (fd OGRFeatureDefn) IsGeometryIgnored() (result bool) {
	result = ogrFDIsGeometryIgnored(fd) != 0
	return
}

func (fd OGRFeatureDefn) SetGeometryIgnored(ignored int) {
	ogrFDSetGeometryIgnored(fd, ignored)
}

func (fd OGRFeatureDefn) IsStyleIgnored() (result bool) {
	result = ogrFDIsStyleIgnored(fd) != 0
	return
}

func (fd OGRFeatureDefn) SetStyleIgnored(ignored int) {
	ogrFDSetStyleIgnored(fd, ignored)
}

func (fd OGRFeatureDefn) Reference() (result int) {
	result = ogrFDReference(fd)
	return
}

func (fd OGRFeatureDefn) Dereference() (result int) {
	result = ogrFDDereference(fd)
	return
}

func (fd OGRFeatureDefn) GetReferenceCount() (result int) {
	result = ogrFDGetReferenceCount(fd)
	return
}

func (fd OGRFeatureDefn) GetGeomFieldCount() (result int) {
	result = ogrFDGetGeomFieldCount(fd)
	return
}

func (fd OGRFeatureDefn) GetGeomFieldDefn(geomField int) (result OGRGeomFieldDefn) {
	result = ogrFDGetGeomFieldDefn(fd, geomField)
	return
}

func (fd OGRFeatureDefn) GetGeomFieldIndex(name string) (result int) {
	result = ogrFDGetGeomFieldIndex(fd, name)
	return
}

func (fd OGRFeatureDefn) AddGeomFieldDefn(gfld OGRGeomFieldDefn) {
	ogrFDAddGeomFieldDefn(fd, gfld)
}

func (fd OGRFeatureDefn) DeleteGeomFieldDefn(geomField int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFDDeleteGeomFieldDefn(fd, geomField))
	return
}

func (fd OGRFeatureDefn) IsSame(other OGRFeatureDefn) (result bool) {
	result = ogrFDIsSame(fd, other) != 0
	return
}

func OGRFCreate(fd OGRFeatureDefn) (result OGRFeature, err error) {
	scope := errScope()
	defer scope()
	result = ogrFCreate(fd)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (feat OGRFeature) Destroy() {
	ogrFDestroy(feat)
}

func (feat OGRFeature) GetDefnRef() (result OGRFeatureDefn) {
	result = ogrFGetDefnRef(feat)
	return
}

func (feat OGRFeature) SetGeometryDirectly(geom OGRGeometry) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFSetGeometryDirectly(feat, geom))
	return
}

func (feat OGRFeature) SetGeometry(geom OGRGeometry) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFSetGeometry(feat, geom))
	return
}

func (feat OGRFeature) GetGeometryRef() (result OGRGeometry) {
	result = ogrFGetGeometryRef(feat)
	return
}

func (feat OGRFeature) StealGeometry() (result OGRGeometry) {
	result = ogrFStealGeometry(feat)
	return
}

func (feat OGRFeature) StealGeometryEx(geomField int) (result OGRGeometry) {
	result = ogrFStealGeometryEx(feat, geomField)
	return
}

func (feat OGRFeature) Clone() (result OGRFeature, err error) {
	scope := errScope()
	defer scope()
	result = ogrFClone(feat)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (feat OGRFeature) Equal(other OGRFeature) (result bool) {
	result = ogrFEqual(feat, other) != 0
	return
}

func (feat OGRFeature) GetFieldCount() (result int) {
	result = ogrFGetFieldCount(feat)
	return
}

func (feat OGRFeature) GetFieldDefnRef(field int) (result OGRFieldDefn) {
	result = ogrFGetFieldDefnRef(feat, field)
	return
}

func (feat OGRFeature) GetFieldIndex(name string) (result int) {
	result = ogrFGetFieldIndex(feat, name)
	return
}

func (feat OGRFeature) IsFieldSet(field int) (result bool) {
	result = ogrFIsFieldSet(feat, field) != 0
	return
}

func (feat OGRFeature) UnsetField(field int) {
	ogrFUnsetField(feat, field)
}

func (feat OGRFeature) IsFieldNull(field int) (result bool) {
	result = ogrFIsFieldNull(feat, field) != 0
	return
}

func (feat OGRFeature) IsFieldSetAndNotNull(field int) (result bool) {
	result = ogrFIsFieldSetAndNotNull(feat, field) != 0
	return
}

func (feat OGRFeature) SetFieldNull(field int) {
	ogrFSetFieldNull(feat, field)
}

func (feat OGRFeature) GetRawFieldRef(field int) (result OGRField) {
	result = ogrFGetRawFieldRef(feat, field)
	return
}

func (f OGRField) IsUnset() (result bool) {
	result = ogrRawFieldIsUnset(f) != 0
	return
}

func (f OGRField) IsNull() (result bool) {
	result = ogrRawFieldIsNull(f) != 0
	return
}

func (f OGRField) SetUnset() {
	ogrRawFieldSetUnset(f)
}

func (f OGRField) SetNull() {
	ogrRawFieldSetNull(f)
}

func (feat OGRFeature) GetFieldAsInteger(field int) (result int) {
	result = ogrFGetFieldAsInteger(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsInteger64(field int) (result int64) {
	result = ogrFGetFieldAsInteger64(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsDouble(field int) (result float64) {
	result = ogrFGetFieldAsDouble(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsString(field int) (result string) {
	result = ogrFGetFieldAsString(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsISO8601DateTime(field int, options CSLConstList) (result string) {
	result = ogrFGetFieldAsISO8601DateTime(feat, field, options)
	return
}

func (feat OGRFeature) GetFieldAsIntegerList(field int) (result []int) {
	result = ogrFGetFieldAsIntegerList(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsInteger64List(field int) (result []int64) {
	result = ogrFGetFieldAsInteger64List(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsDoubleList(field int) (result []float64) {
	result = ogrFGetFieldAsDoubleList(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsStringList(field int) (result CSLConstList) {
	result = ogrFGetFieldAsStringList(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsBinary(field int) (result []byte) {
	result = ogrFGetFieldAsBinary(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsDateTime(field int) (year, month, day, hour, minute, second, tzFlag int, ok bool) {
	year, month, day, hour, minute, second, tzFlag, ok = ogrFGetFieldAsDateTime(feat, field)
	return
}

func (feat OGRFeature) GetFieldAsDateTimeEx(field int) (year, month, day, hour, minute int, second float32, tzFlag int, ok bool) {
	year, month, day, hour, minute, second, tzFlag, ok = ogrFGetFieldAsDateTimeEx(feat, field)
	return
}

func (feat OGRFeature) SetFieldInteger(field, value int) {
	ogrFSetFieldInteger(feat, field, value)
}

func (feat OGRFeature) SetFieldInteger64(field int, value int64) {
	ogrFSetFieldInteger64(feat, field, value)
}

func (feat OGRFeature) SetFieldDouble(field int, value float64) {
	ogrFSetFieldDouble(feat, field, value)
}

func (feat OGRFeature) SetFieldString(field int, value string) {
	ogrFSetFieldString(feat, field, value)
}

func (feat OGRFeature) SetFieldIntegerList(field int, values []int) {
	ogrFSetFieldIntegerList(feat, field, values)
}

func (feat OGRFeature) SetFieldInteger64List(field int, values []int64) {
	ogrFSetFieldInteger64List(feat, field, values)
}

func (feat OGRFeature) SetFieldDoubleList(field int, values []float64) {
	ogrFSetFieldDoubleList(feat, field, values)
}

func (feat OGRFeature) SetFieldStringList(field int, values CSLConstList) {
	ogrFSetFieldStringList(feat, field, values)
}

func (feat OGRFeature) SetFieldRaw(field int, value OGRField) {
	ogrFSetFieldRaw(feat, field, value)
}

func (feat OGRFeature) SetFieldBinary(field int, data []byte) {
	ogrFSetFieldBinary(feat, field, data)
}

func (feat OGRFeature) SetFieldDateTime(field, year, month, day, hour, minute, second, tzFlag int) {
	ogrFSetFieldDateTime(feat, field, year, month, day, hour, minute, second, tzFlag)
}

func (feat OGRFeature) SetFieldDateTimeEx(field, year, month, day, hour, minute int, second float32, tzFlag int) {
	ogrFSetFieldDateTimeEx(feat, field, year, month, day, hour, minute, second, tzFlag)
}

func (feat OGRFeature) GetGeomFieldCount() (result int) {
	result = ogrFGetGeomFieldCount(feat)
	return
}

func (feat OGRFeature) GetGeomFieldDefnRef(field int) (result OGRGeomFieldDefn) {
	result = ogrFGetGeomFieldDefnRef(feat, field)
	return
}

func (feat OGRFeature) GetGeomFieldIndex(name string) (result int) {
	result = ogrFGetGeomFieldIndex(feat, name)
	return
}

func (feat OGRFeature) GetGeomFieldRef(field int) (result OGRGeometry) {
	result = ogrFGetGeomFieldRef(feat, field)
	return
}

func (feat OGRFeature) SetGeomFieldDirectly(field int, geom OGRGeometry) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFSetGeomFieldDirectly(feat, field, geom))
	return
}

func (feat OGRFeature) SetGeomField(field int, geom OGRGeometry) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFSetGeomField(feat, field, geom))
	return
}

func (feat OGRFeature) GetFID() (result int64) {
	result = ogrFGetFID(feat)
	return
}

func (feat OGRFeature) SetFID(fid int64) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFSetFID(feat, fid))
	return
}

func (feat OGRFeature) DumpReadableAsString(options CSLConstList) (result string) {
	result = ogrFDumpReadableAsString(feat, options)
	return
}

func (feat OGRFeature) SetFrom(other OGRFeature, forgiving int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFSetFrom(feat, other, forgiving))
	return
}

func (feat OGRFeature) SetFromWithMap(other OGRFeature, forgiving int, panMap []int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrFSetFromWithMap(feat, other, forgiving, panMap))
	return
}

func (feat OGRFeature) GetStyleString() (result string) {
	result = ogrFGetStyleString(feat)
	return
}

func (feat OGRFeature) SetStyleString(style string) {
	ogrFSetStyleString(feat, style)
}

func (feat OGRFeature) SetStyleStringDirectly(style string) {
	ogrFSetStyleStringDirectly(feat, style)
}

// /** Return style table */
func (feat OGRFeature) GetStyleTable() (result OGRStyleTable) {
	result = ogrFGetStyleTable(feat)
	return
}

// /** Set style table and take ownership */
func (feat OGRFeature) SetStyleTableDirectly(styleTable OGRStyleTable) {
	ogrFSetStyleTableDirectly(feat, styleTable)
}

// /** Set style table */
func (feat OGRFeature) SetStyleTable(styleTable OGRStyleTable) {
	ogrFSetStyleTable(feat, styleTable)
}

func (feat OGRFeature) GetNativeData() (result string) {
	result = ogrFGetNativeData(feat)
	return
}

func (feat OGRFeature) SetNativeData(data string) {
	ogrFSetNativeData(feat, data)
}

func (feat OGRFeature) GetNativeMediaType() (result string) {
	result = ogrFGetNativeMediaType(feat)
	return
}

func (feat OGRFeature) SetNativeMediaType(mediaType string) {
	ogrFSetNativeMediaType(feat, mediaType)
}

func (feat OGRFeature) FillUnsetWithDefault(notNullableOnly int, options CSLConstList) {
	ogrFFillUnsetWithDefault(feat, notNullableOnly, options)
}

func (feat OGRFeature) Validate(validateFlags, emitError int) (result bool) {
	result = ogrFValidate(feat, validateFlags, emitError) != 0
	return
}

func (dom OGRFieldDomain) Destroy() {
	ogrFldDomainDestroy(dom)
}

func (dom OGRFieldDomain) GetName() (result string) {
	result = ogrFldDomainGetName(dom)
	return
}

func (dom OGRFieldDomain) GetDescription() (result string) {
	result = ogrFldDomainGetDescription(dom)
	return
}

func (dom OGRFieldDomain) GetDomainType() (result OGRFieldDomainType) {
	result = ogrFldDomainGetDomainType(dom)
	return
}

func (dom OGRFieldDomain) GetFieldType() (result OGRFieldType) {
	result = ogrFldDomainGetFieldType(dom)
	return
}

func (dom OGRFieldDomain) GetFieldSubType() (result OGRFieldSubType) {
	result = ogrFldDomainGetFieldSubType(dom)
	return
}

func (dom OGRFieldDomain) GetSplitPolicy() (result OGRFieldDomainSplitPolicy) {
	result = ogrFldDomainGetSplitPolicy(dom)
	return
}

func (dom OGRFieldDomain) SetSplitPolicy(policy OGRFieldDomainSplitPolicy) {
	ogrFldDomainSetSplitPolicy(dom, policy)
}

func (dom OGRFieldDomain) GetMergePolicy() (result OGRFieldDomainMergePolicy) {
	result = ogrFldDomainGetMergePolicy(dom)
	return
}

func (dom OGRFieldDomain) SetMergePolicy(policy OGRFieldDomainMergePolicy) {
	ogrFldDomainSetMergePolicy(dom, policy)
}

func OGRCodedFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, enumeration OGRCodedValue) (result OGRFieldDomain, err error) {
	scope := errScope()
	defer scope()
	result = ogrCodedFldDomainCreate(name, description, eFieldType, eFieldSubType, enumeration)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (dom OGRFieldDomain) GetEnumeration() (result OGRCodedValue) {
	result = ogrCodedFldDomainGetEnumeration(dom)
	return
}

func OGRRangeFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, min OGRField, minInclusive bool, max OGRField, maxInclusive bool) (result OGRFieldDomain, err error) {
	scope := errScope()
	defer scope()
	result = ogrRangeFldDomainCreate(name, description, eFieldType, eFieldSubType, min, minInclusive, max, maxInclusive)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (dom OGRFieldDomain) GetMin() (result OGRField, inclusive bool) {
	result, inclusive = ogrRangeFldDomainGetMin(dom)
	return
}

func (dom OGRFieldDomain) GetMax() (result OGRField, inclusive bool) {
	result, inclusive = ogrRangeFldDomainGetMax(dom)
	return
}

func OGRGlobFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, glob string) (result OGRFieldDomain, err error) {
	scope := errScope()
	defer scope()
	result = ogrGlobFldDomainCreate(name, description, eFieldType, eFieldSubType, glob)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (dom OGRFieldDomain) GetGlob() (result string) {
	result = ogrGlobFldDomainGetGlob(dom)
	return
}

func (l OGRLayer) GetName() (result string) {
	result = ogrLGetName(l)
	return
}

func (l OGRLayer) GetGeomType() (result OGRwkbGeometryType) {
	result = ogrLGetGeomType(l)
	return
}

func (c OGRGeometryTypeCounter) GeomType() (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(c.cValue.eGeomType)
	return
}

func (c OGRGeometryTypeCounter) Count() (result int64) {
	result = int64(c.cValue.nCount)
	return
}

func (l OGRLayer) GetGeometryTypes(iGeomField, flags int, progress GDALProgressFunc, progressData unsafe.Pointer) (result []OGRGeometryTypeCounter) {
	result = ogrLGetGeometryTypes(l, iGeomField, flags, progress, progressData)
	return
}

func (l OGRLayer) GetSpatialFilter() (result OGRGeometry) {
	result = ogrLGetSpatialFilter(l)
	return
}

func (l OGRLayer) SetSpatialFilter(geom OGRGeometry) {
	ogrLSetSpatialFilter(l, geom)
}

func (l OGRLayer) SetSpatialFilterRect(minX, minY, maxX, maxY float64) {
	ogrLSetSpatialFilterRect(l, minX, minY, maxX, maxY)
}

func (l OGRLayer) SetSpatialFilterEx(iGeomField int, geom OGRGeometry) {
	ogrLSetSpatialFilterEx(l, iGeomField, geom)
}

func (l OGRLayer) SetSpatialFilterRectEx(iGeomField int, minX, minY, maxX, maxY float64) {
	ogrLSetSpatialFilterRectEx(l, iGeomField, minX, minY, maxX, maxY)
}

func (l OGRLayer) SetAttributeFilter(query string) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLSetAttributeFilter(l, query))
	return
}

func (l OGRLayer) ResetReading() {
	ogrLResetReading(l)
}

func (l OGRLayer) GetNextFeature() (result OGRFeature) {
	result = ogrLGetNextFeature(l)
	return
}

func (l OGRLayer) SetNextByIndex(index int64) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLSetNextByIndex(l, index))
	return
}

func (l OGRLayer) GetFeature(fid int64) (result OGRFeature, err error) {
	scope := errScope()
	defer scope()
	result = ogrLGetFeature(l, fid)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (l OGRLayer) SetFeature(feat OGRFeature) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLSetFeature(l, feat))
	return
}

func (l OGRLayer) CreateFeature(feat OGRFeature) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLCreateFeature(l, feat))
	return
}

func (l OGRLayer) DeleteFeature(fid int64) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLDeleteFeature(l, fid))
	return
}

func (l OGRLayer) UpsertFeature(feat OGRFeature) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLUpsertFeature(l, feat))
	return
}

func (l OGRLayer) UpdateFeature(feat OGRFeature, updatedFieldsIdx, updatedGeomFieldsIdx []int, updateStyleString bool) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLUpdateFeature(l, feat, updatedFieldsIdx, updatedGeomFieldsIdx, updateStyleString))
	return
}

func (l OGRLayer) GetLayerDefn() (result OGRFeatureDefn) {
	result = ogrLGetLayerDefn(l)
	return
}

func (l OGRLayer) GetSpatialRef() (result OGRSpatialReference) {
	result = ogrLGetSpatialRef(l)
	return
}

func (l OGRLayer) GetSupportedSRSList(iGeomField int) (result []OGRSpatialReference) {
	result = ogrLGetSupportedSRSList(l, iGeomField)
	return
}

func (l OGRLayer) SetActiveSRS(iGeomField int, sr OGRSpatialReference) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLSetActiveSRS(l, iGeomField, sr))
	return
}

func (l OGRLayer) FindFieldIndex(name string, exactMatch int) (result int) {
	result = ogrLFindFieldIndex(l, name, exactMatch)
	return
}

func (l OGRLayer) GetFeatureCount(force int) (result int64) {
	result = ogrLGetFeatureCount(l, force)
	return
}

func (l OGRLayer) GetExtent(force int) (result OGREnvelope, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrLGetExtent(l, force)
	err = ogrError(status)
	return
}

func (l OGRLayer) GetExtentEx(iGeomField, force int) (result OGREnvelope, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrLGetExtentEx(l, iGeomField, force)
	err = ogrError(status)
	return
}

func (l OGRLayer) GetExtent3D(iGeomField, force int) (result OGREnvelope3D, err error) {
	scope := errScope()
	defer scope()
	var status OGRErr
	result, status = ogrLGetExtent3D(l, iGeomField, force)
	err = ogrError(status)
	return
}

func (l OGRLayer) TestCapability(capability string) (result bool) {
	result = ogrLTestCapability(l, capability) != 0
	return
}

func (l OGRLayer) CreateField(fld OGRFieldDefn, approxOK int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLCreateField(l, fld, approxOK))
	return
}

func (l OGRLayer) CreateGeomField(gfld OGRGeomFieldDefn, force int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLCreateGeomField(l, gfld, force))
	return
}

func (l OGRLayer) DeleteField(field int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLDeleteField(l, field))
	return
}

func (l OGRLayer) ReorderFields(panMap []int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLReorderFields(l, panMap))
	return
}

func (l OGRLayer) ReorderField(oldFieldPos, newFieldPos int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLReorderField(l, oldFieldPos, newFieldPos))
	return
}

func (l OGRLayer) AlterFieldDefn(field int, newFieldDefn OGRFieldDefn, flags int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLAlterFieldDefn(l, field, newFieldDefn, flags))
	return
}

func (l OGRLayer) AlterGeomFieldDefn(field int, newGeomFieldDefn OGRGeomFieldDefn, flags int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLAlterGeomFieldDefn(l, field, newGeomFieldDefn, flags))
	return
}

func (l OGRLayer) StartTransaction() (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLStartTransaction(l))
	return
}

func (l OGRLayer) CommitTransaction() (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLCommitTransaction(l))
	return
}

func (l OGRLayer) RollbackTransaction() (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLRollbackTransaction(l))
	return
}

func (l OGRLayer) Rename(newName string) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLRename(l, newName))
	return
}

func (l OGRLayer) Reference() (result int) {
	result = ogrLReference(l)
	return
}

func (l OGRLayer) Dereference() (result int) {
	result = ogrLDereference(l)
	return
}

func (l OGRLayer) GetRefCount() (result int) {
	result = ogrLGetRefCount(l)
	return
}

func (l OGRLayer) SyncToDisk() (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLSyncToDisk(l))
	return
}

func (l OGRLayer) GetFeaturesRead() (result int64) {
	result = ogrLGetFeaturesRead(l)
	return
}

func (l OGRLayer) GetFIDColumn() (result string) {
	result = ogrLGetFIDColumn(l)
	return
}

func (l OGRLayer) GetGeometryColumn() (result string) {
	result = ogrLGetGeometryColumn(l)
	return
}

// /** Get style table */
func (l OGRLayer) GetStyleTable() (result OGRStyleTable) {
	result = ogrLGetStyleTable(l)
	return
}

// /** Set style table (and take ownership) */
func (l OGRLayer) SetStyleTableDirectly(styleTable OGRStyleTable) {
	ogrLSetStyleTableDirectly(l, styleTable)
}

// /** Set style table */
func (l OGRLayer) SetStyleTable(styleTable OGRStyleTable) {
	ogrLSetStyleTable(l, styleTable)
}

func (l OGRLayer) SetIgnoredFields(fields CSLConstList) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLSetIgnoredFields(l, fields))
	return
}

func OGRLIntersection(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLIntersection(input, method, result, options, progress, progressData))
	return
}

func OGRLUnion(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLUnion(input, method, result, options, progress, progressData))
	return
}

func OGRLSymDifference(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLSymDifference(input, method, result, options, progress, progressData))
	return
}

func OGRLIdentity(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLIdentity(input, method, result, options, progress, progressData))
	return
}

func OGRLUpdate(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLUpdate(input, method, result, options, progress, progressData))
	return
}

func OGRLClip(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLClip(input, method, result, options, progress, progressData))
	return
}

func OGRLErase(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrLErase(input, method, result, options, progress, progressData))
	return
}

func (ds OGRDataSource) Destroy() {
	ogrDSDestroy(ds)
}

func (ds OGRDataSource) GetName() (result string) {
	result = ogrDSGetName(ds)
	return
}

func (ds OGRDataSource) GetLayerCount() (result int) {
	result = ogrDSGetLayerCount(ds)
	return
}

func (ds OGRDataSource) GetLayer(layer int) (result OGRLayer, err error) {
	scope := errScope()
	defer scope()
	result = ogrDSGetLayer(ds, layer)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds OGRDataSource) GetLayerByName(name string) (result OGRLayer, err error) {
	scope := errScope()
	defer scope()
	result = ogrDSGetLayerByName(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds OGRDataSource) DeleteLayer(layer int) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrDSDeleteLayer(ds, layer))
	return
}

func (ds OGRDataSource) GetDriver() (result OGRSFDriver, err error) {
	scope := errScope()
	defer scope()
	result = ogrDSGetDriver(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds OGRDataSource) CreateLayer(name string, sr OGRSpatialReference, geomType OGRwkbGeometryType, options CSLConstList) (result OGRLayer, err error) {
	scope := errScope()
	defer scope()
	result = ogrDSCreateLayer(ds, name, sr, geomType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds OGRDataSource) CopyLayer(srcLayer OGRLayer, newName string, options CSLConstList) (result OGRLayer, err error) {
	scope := errScope()
	defer scope()
	result = ogrDSCopyLayer(ds, srcLayer, newName, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds OGRDataSource) TestCapability(capability string) (result bool) {
	result = ogrDSTestCapability(ds, capability) != 0
	return
}

func (ds OGRDataSource) ExecuteSQL(statement string, spatialFilter OGRGeometry, dialect string) (result OGRLayer, err error) {
	scope := errScope()
	defer scope()
	result = ogrDSExecuteSQL(ds, statement, spatialFilter, dialect)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds OGRDataSource) ReleaseResultSet(layer OGRLayer) {
	ogrDSReleaseResultSet(ds, layer)
}

func (ds OGRDataSource) Reference() (result int) {
	result = ogrDSReference(ds)
	return
}

func (ds OGRDataSource) Dereference() (result int) {
	result = ogrDSDereference(ds)
	return
}

func (ds OGRDataSource) GetRefCount() (result int) {
	result = ogrDSGetRefCount(ds)
	return
}

func (ds OGRDataSource) GetSummaryRefCount() (result int) {
	result = ogrDSGetSummaryRefCount(ds)
	return
}

// /** Flush pending changes to disk. See GDALDataset::FlushCache() */
func (ds OGRDataSource) SyncToDisk() (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrDSSyncToDisk(ds))
	return
}

// /** Get style table */
func (ds OGRDataSource) GetStyleTable() (result OGRStyleTable) {
	result = ogrDSGetStyleTable(ds)
	return
}

// /** Set style table (and take ownership) */
func (ds OGRDataSource) SetStyleTableDirectly(styleTable OGRStyleTable) {
	ogrDSSetStyleTableDirectly(ds, styleTable)
}

// /** Set style table */
func (ds OGRDataSource) SetStyleTable(styleTable OGRStyleTable) {
	ogrDSSetStyleTable(ds, styleTable)
}

func (dr OGRSFDriver) GetName() (result string) {
	result = ogrDrGetName(dr)
	return
}

func (dr OGRSFDriver) Open(name string, update int) (result OGRDataSource, err error) {
	scope := errScope()
	defer scope()
	result = ogrDrOpen(dr, name, update)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (dr OGRSFDriver) TestCapability(capability string) (result bool) {
	result = ogrDrTestCapability(dr, capability) != 0
	return
}

func (dr OGRSFDriver) CreateDataSource(name string, options CSLConstList) (result OGRDataSource, err error) {
	scope := errScope()
	defer scope()
	result = ogrDrCreateDataSource(dr, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (dr OGRSFDriver) CopyDataSource(srcDS OGRDataSource, newName string, options CSLConstList) (result OGRDataSource, err error) {
	scope := errScope()
	defer scope()
	result = ogrDrCopyDataSource(dr, srcDS, newName, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (dr OGRSFDriver) DeleteDataSource(name string) (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrDrDeleteDataSource(dr, name))
	return
}

func OGROpen(name string, update int) (result OGRDataSource, driver OGRSFDriver, err error) {
	scope := errScope()
	defer scope()
	result, driver = ogrOpen(name, update)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func OGROpenShared(name string, update int) (result OGRDataSource, driver OGRSFDriver, err error) {
	scope := errScope()
	defer scope()
	result, driver = ogrOpenShared(name, update)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (ds OGRDataSource) Release() (err error) {
	scope := errScope()
	defer scope()
	err = ogrError(ogrReleaseDataSource(ds))
	return
}

func (dr OGRSFDriver) Register() {
	ogrRegisterDriver(dr)
}

func (dr OGRSFDriver) Deregister() {
	ogrDeregisterDriver(dr)
}

func OGRGetDriverCount() (result int) {
	result = ogrGetDriverCount()
	return
}

func OGRGetDriver(driver int) (result OGRSFDriver, err error) {
	scope := errScope()
	defer scope()
	result = ogrGetDriver(driver)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func OGRGetDriverByName(name string) (result OGRSFDriver, err error) {
	scope := errScope()
	defer scope()
	result = ogrGetDriverByName(name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func OGRGetOpenDSCount() (result int) {
	result = ogrGetOpenDSCount()
	return
}

func OGRGetOpenDS(ds int) (result OGRDataSource, err error) {
	scope := errScope()
	defer scope()
	result = ogrGetOpenDS(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func OGRRegisterAll() {
	ogrRegisterAll()
}

// /** Clean-up all drivers, including raster ones.
//   - See GDALDestroyDriverManager() */
func OGRCleanupAll() {
	ogrCleanupAll()
}

func OGRSMCreate(styleTable OGRStyleTable) (result OGRStyleMgr, err error) {
	scope := errScope()
	defer scope()
	result = ogrSMCreate(styleTable)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (sm OGRStyleMgr) Destroy() {
	ogrSMDestroy(sm)
}

func (sm OGRStyleMgr) InitFromFeature(feat OGRFeature) (result string) {
	result = ogrSMInitFromFeature(sm, feat)
	return
}

func (sm OGRStyleMgr) InitStyleString(styleString string) (result bool) {
	result = ogrSMInitStyleString(sm, styleString) != 0
	return
}

func (sm OGRStyleMgr) GetPartCount(styleString string) (result int) {
	result = ogrSMGetPartCount(sm, styleString)
	return
}

func (sm OGRStyleMgr) GetPart(partId int, styleString string) (result OGRStyleTool) {
	result = ogrSMGetPart(sm, partId, styleString)
	return
}

func (sm OGRStyleMgr) AddPart(st OGRStyleTool) (result bool) {
	result = ogrSMAddPart(sm, st) != 0
	return
}

func (sm OGRStyleMgr) AddStyle(styleName, styleString string) (result bool) {
	result = ogrSMAddStyle(sm, styleName, styleString) != 0
	return
}

func OGRSTCreate(classId OGRSTClassId) (result OGRStyleTool, err error) {
	scope := errScope()
	defer scope()
	result = ogrSTCreate(classId)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (st OGRStyleTool) Destroy() {
	ogrSTDestroy(st)
}

func (st OGRStyleTool) GetType() (result OGRSTClassId) {
	result = ogrSTGetType(st)
	return
}

func (st OGRStyleTool) GetUnit() (result OGRSTUnitId) {
	result = ogrSTGetUnit(st)
	return
}

func (st OGRStyleTool) SetUnit(unit OGRSTUnitId, groundPaperScale float64) {
	ogrSTSetUnit(st, unit, groundPaperScale)
}

func (st OGRStyleTool) GetParamStr(param int) (result string, isNull bool) {
	result, isNull = ogrSTGetParamStr(st, param)
	return
}

func (st OGRStyleTool) GetParamNum(param int) (result int, isNull bool) {
	result, isNull = ogrSTGetParamNum(st, param)
	return
}

func (st OGRStyleTool) GetParamDbl(param int) (result float64, isNull bool) {
	result, isNull = ogrSTGetParamDbl(st, param)
	return
}

func (st OGRStyleTool) SetParamStr(param int, value string) {
	ogrSTSetParamStr(st, param, value)
}

func (st OGRStyleTool) SetParamNum(param, value int) {
	ogrSTSetParamNum(st, param, value)
}

func (st OGRStyleTool) SetParamDbl(param int, value float64) {
	ogrSTSetParamDbl(st, param, value)
}

func (st OGRStyleTool) GetStyleString() (result string) {
	result = ogrSTGetStyleString(st)
	return
}

func (st OGRStyleTool) GetRGBFromString(color string) (red, green, blue, alpha int, ok bool) {
	red, green, blue, alpha, ok = ogrSTGetRGBFromString(st, color)
	return
}

func OGRSTBLCreate() (result OGRStyleTable, err error) {
	scope := errScope()
	defer scope()
	result = ogrSTBLCreate()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (stbl OGRStyleTable) Destroy() {
	ogrSTBLDestroy(stbl)
}

func (stbl OGRStyleTable) AddStyle(name, styleString string) (result bool) {
	result = ogrSTBLAddStyle(stbl, name, styleString) != 0
	return
}

func (stbl OGRStyleTable) SaveStyleTable(filename string) (result bool) {
	result = ogrSTBLSaveStyleTable(stbl, filename) != 0
	return
}

func (stbl OGRStyleTable) LoadStyleTable(filename string) (result bool) {
	result = ogrSTBLLoadStyleTable(stbl, filename) != 0
	return
}

func (stbl OGRStyleTable) Find(name string) (result string) {
	result = ogrSTBLFind(stbl, name)
	return
}

func (stbl OGRStyleTable) ResetStyleStringReading() {
	ogrSTBLResetStyleStringReading(stbl)
}

func (stbl OGRStyleTable) GetNextStyle() (result string) {
	result = ogrSTBLGetNextStyle(stbl)
	return
}

func (stbl OGRStyleTable) GetLastStyleName() (result string) {
	result = ogrSTBLGetLastStyleName(stbl)
	return
}
