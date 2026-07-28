package gdal

/*
#include "ogr_api_preamble.h"
*/
import "C"
import "unsafe"

// CPL_C_START

func ogrGetGEOSVersion() (major, minor, patch int) {
	var cMajor, cMinor, cPatch C.int
	C.OGRGetGEOSVersion(&cMajor, &cMinor, &cPatch)
	major = int(cMajor)
	minor = int(cMinor)
	patch = int(cPatch)
	return
}

// /* -------------------------------------------------------------------- */
// /*      Geometry related functions (ogr_geometry.h)                     */
// /* -------------------------------------------------------------------- */

// struct _CPLXMLNode;

// /* OGRGeomCoordinatePrecisionH */

// /** Value for a unknown coordinate precision. */
const OGRGeomCoordPrecisionUnknown = C.OGR_GEOM_COORD_PRECISION_UNKNOWN

func ogrGeomCoordinatePrecisionCreate() (result OGRGeomCoordinatePrecision) {
	result = OGRGeomCoordinatePrecision{cValue: C.OGRGeomCoordinatePrecisionCreate()}
	return
}

func ogrGeomCoordinatePrecisionDestroy(p OGRGeomCoordinatePrecision) {
	C.OGRGeomCoordinatePrecisionDestroy(p.cValue)
}

func ogrGeomCoordinatePrecisionGetXYResolution(p OGRGeomCoordinatePrecision) (result float64) {
	result = float64(C.OGRGeomCoordinatePrecisionGetXYResolution(p.cValue))
	return
}

func ogrGeomCoordinatePrecisionGetZResolution(p OGRGeomCoordinatePrecision) (result float64) {
	result = float64(C.OGRGeomCoordinatePrecisionGetZResolution(p.cValue))
	return
}

func ogrGeomCoordinatePrecisionGetMResolution(p OGRGeomCoordinatePrecision) (result float64) {
	result = float64(C.OGRGeomCoordinatePrecisionGetMResolution(p.cValue))
	return
}

func ogrGeomCoordinatePrecisionGetFormats(p OGRGeomCoordinatePrecision) (result CSLConstList) {
	raw := C.OGRGeomCoordinatePrecisionGetFormats(p.cValue)
	result = cslConstList(raw)
	return
}

func ogrGeomCoordinatePrecisionGetFormatSpecificOptions(p OGRGeomCoordinatePrecision, formatName string) (result CSLConstList) {
	cs := C.CString(formatName)
	defer C.free(unsafe.Pointer(cs))
	raw := C.OGRGeomCoordinatePrecisionGetFormatSpecificOptions(p.cValue, cs)
	result = cslConstList(raw)
	return
}

func ogrGeomCoordinatePrecisionSet(p OGRGeomCoordinatePrecision, xyResolution, zResolution, mResolution float64) {
	C.OGRGeomCoordinatePrecisionSet(p.cValue, C.double(xyResolution), C.double(zResolution), C.double(mResolution))
}

func ogrGeomCoordinatePrecisionSetFromMeter(p OGRGeomCoordinatePrecision, sr OGRSpatialReference, xyMeterResolution, zMeterResolution, mResolution float64) {
	C.OGRGeomCoordinatePrecisionSetFromMeter(p.cValue, sr.cValue, C.double(xyMeterResolution), C.double(zMeterResolution), C.double(mResolution))
}

func ogrGeomCoordinatePrecisionSetFormatSpecificOptions(p OGRGeomCoordinatePrecision, formatName string, options CSLConstList) {
	csName := C.CString(formatName)
	defer C.free(unsafe.Pointer(csName))
	cOptions := options.cValue
	C.OGRGeomCoordinatePrecisionSetFormatSpecificOptions(p.cValue, csName, cOptions)
}

// /* From base OGRGeometry class */

func ogrGCreateFromWkb(data []byte, sr OGRSpatialReference) (result OGRGeometry, status OGRErr) {
	var hGeom C.OGRGeometryH
	var p unsafe.Pointer
	if len(data) > 0 {
		p = unsafe.Pointer(&data[0])
	}
	status = OGRErr(C.OGR_G_CreateFromWkb(p, sr.cValue, &hGeom, C.int(len(data))))
	result = OGRGeometry{cValue: hGeom}
	return
}

func ogrGCreateFromWkbEx(data []byte, sr OGRSpatialReference) (result OGRGeometry, status OGRErr) {
	var hGeom C.OGRGeometryH
	var p unsafe.Pointer
	if len(data) > 0 {
		p = unsafe.Pointer(&data[0])
	}
	status = OGRErr(C.OGR_G_CreateFromWkbEx(p, sr.cValue, &hGeom, C.size_t(len(data))))
	result = OGRGeometry{cValue: hGeom}
	return
}

func ogrGCreateFromWkt(wkt string, sr OGRSpatialReference) (result OGRGeometry, status OGRErr) {
	cs := C.CString(wkt)
	defer C.free(unsafe.Pointer(cs))
	pcs := cs
	var hGeom C.OGRGeometryH
	status = OGRErr(C.OGR_G_CreateFromWkt(&pcs, sr.cValue, &hGeom))
	result = OGRGeometry{cValue: hGeom}
	return
}

func ogrGCreateFromFgf(data []byte, sr OGRSpatialReference) (result OGRGeometry, bytesConsumed int, status OGRErr) {
	var hGeom C.OGRGeometryH
	var consumed C.int
	var p unsafe.Pointer
	if len(data) > 0 {
		p = unsafe.Pointer(&data[0])
	}
	status = OGRErr(C.OGR_G_CreateFromFgf(p, sr.cValue, &hGeom, C.int(len(data)), &consumed))
	result = OGRGeometry{cValue: hGeom}
	bytesConsumed = int(consumed)
	return
}

func ogrGCreateFromEnvelope(minX, maxX, minY, maxY float64, sr OGRSpatialReference) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_CreateFromEnvelope(C.double(minX), C.double(maxX), C.double(minY), C.double(maxY), sr.cValue)}
	return
}

func ogrGDestroyGeometry(g OGRGeometry) {
	C.OGR_G_DestroyGeometry(g.cValue)
}

func ogrGCreateGeometry(eType OGRwkbGeometryType) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_CreateGeometry(C.OGRwkbGeometryType(eType))}
	return
}

func ogrGApproximateArcAngles(centerX, centerY, z, primaryRadius, secondaryAxis, rotation, startAngle, endAngle, maxAngleStepSizeDegrees float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ApproximateArcAngles(C.double(centerX), C.double(centerY), C.double(z), C.double(primaryRadius), C.double(secondaryAxis), C.double(rotation), C.double(startAngle), C.double(endAngle), C.double(maxAngleStepSizeDegrees))}
	return
}

func ogrGForceToPolygon(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToPolygon(g.cValue)}
	return
}

func ogrGForceToLineString(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToLineString(g.cValue)}
	return
}

func ogrGForceToMultiPolygon(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToMultiPolygon(g.cValue)}
	return
}

func ogrGForceToMultiPoint(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToMultiPoint(g.cValue)}
	return
}

func ogrGForceToMultiLineString(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToMultiLineString(g.cValue)}
	return
}

func ogrGForceTo(g OGRGeometry, eTargetType OGRwkbGeometryType, options CSLConstList) (result OGRGeometry) {
	cOptions := options.cValue
	result = OGRGeometry{cValue: C.OGR_G_ForceTo(g.cValue, C.OGRwkbGeometryType(eTargetType), cOptions)}
	return
}

func ogrGRemoveLowerDimensionSubGeoms(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_RemoveLowerDimensionSubGeoms(g.cValue)}
	return
}

func ogrGGetDimension(g OGRGeometry) (result int) {
	result = int(C.OGR_G_GetDimension(g.cValue))
	return
}

func ogrGGetCoordinateDimension(g OGRGeometry) (result int) {
	result = int(C.OGR_G_GetCoordinateDimension(g.cValue))
	return
}

func ogrGCoordinateDimension(g OGRGeometry) (result int) {
	result = int(C.OGR_G_CoordinateDimension(g.cValue))
	return
}

func ogrGSetCoordinateDimension(g OGRGeometry, dimension int) {
	C.OGR_G_SetCoordinateDimension(g.cValue, C.int(dimension))
}

func ogrGIs3D(g OGRGeometry) (result int) {
	result = int(C.OGR_G_Is3D(g.cValue))
	return
}

func ogrGIsMeasured(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsMeasured(g.cValue))
	return
}

func ogrGSet3D(g OGRGeometry, is3D int) {
	C.OGR_G_Set3D(g.cValue, C.int(is3D))
}

func ogrGSetMeasured(g OGRGeometry, isMeasured int) {
	C.OGR_G_SetMeasured(g.cValue, C.int(isMeasured))
}

func ogrGClone(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Clone(g.cValue)}
	return
}

func ogrGGetEnvelope(g OGRGeometry) (result OGREnvelope) {
	result = OGREnvelope{cValue: new(C.OGREnvelope)}
	C.OGR_G_GetEnvelope(g.cValue, result.cValue)
	return
}

func ogrGGetEnvelope3D(g OGRGeometry) (result OGREnvelope3D) {
	result = OGREnvelope3D{cValue: new(C.OGREnvelope3D)}
	C.OGR_G_GetEnvelope3D(g.cValue, result.cValue)
	return
}

func ogrGImportFromWkb(g OGRGeometry, data []byte) (result OGRErr) {
	var p unsafe.Pointer
	if len(data) > 0 {
		p = unsafe.Pointer(&data[0])
	}
	result = OGRErr(C.OGR_G_ImportFromWkb(g.cValue, p, C.int(len(data))))
	return
}

func ogrGExportToWkb(g OGRGeometry, order OGRwkbByteOrder) (result []byte, status OGRErr) {
	n := int(C.OGR_G_WkbSize(g.cValue))
	if n <= 0 {
		return
	}
	buf := make([]byte, n)
	status = OGRErr(C.OGR_G_ExportToWkb(g.cValue, C.OGRwkbByteOrder(order), (*C.uchar)(unsafe.Pointer(&buf[0]))))
	if status == OGRErrNone {
		result = buf
	}
	return
}

func ogrGExportToIsoWkb(g OGRGeometry, order OGRwkbByteOrder) (result []byte, status OGRErr) {
	n := int(C.OGR_G_WkbSize(g.cValue))
	if n <= 0 {
		return
	}
	buf := make([]byte, n)
	status = OGRErr(C.OGR_G_ExportToIsoWkb(g.cValue, C.OGRwkbByteOrder(order), (*C.uchar)(unsafe.Pointer(&buf[0]))))
	if status == OGRErrNone {
		result = buf
	}
	return
}

func ogrwkbExportOptionsCreate() (result OGRwkbExportOptions) {
	result = OGRwkbExportOptions{cValue: C.OGRwkbExportOptionsCreate()}
	return
}

func ogrwkbExportOptionsDestroy(o OGRwkbExportOptions) {
	C.OGRwkbExportOptionsDestroy(o.cValue)
}

func ogrwkbExportOptionsSetByteOrder(o OGRwkbExportOptions, order OGRwkbByteOrder) {
	C.OGRwkbExportOptionsSetByteOrder(o.cValue, C.OGRwkbByteOrder(order))
}

func ogrwkbExportOptionsSetVariant(o OGRwkbExportOptions, variant OGRwkbVariant) {
	C.OGRwkbExportOptionsSetVariant(o.cValue, C.OGRwkbVariant(variant))
}

func ogrwkbExportOptionsSetPrecision(o OGRwkbExportOptions, precision OGRGeomCoordinatePrecision) {
	C.OGRwkbExportOptionsSetPrecision(o.cValue, precision.cValue)
}

func ogrGExportToWkbEx(g OGRGeometry, opts OGRwkbExportOptions) (result []byte, status OGRErr) {
	n := int(C.OGR_G_WkbSizeEx(g.cValue))
	if n <= 0 {
		return
	}
	buf := make([]byte, n)
	status = OGRErr(C.OGR_G_ExportToWkbEx(g.cValue, (*C.uchar)(unsafe.Pointer(&buf[0])), opts.cValue))
	if status == OGRErrNone {
		result = buf
	}
	return
}

func ogrGWkbSize(g OGRGeometry) (result int) {
	result = int(C.OGR_G_WkbSize(g.cValue))
	return
}

func ogrGWkbSizeEx(g OGRGeometry) (result int) {
	result = int(C.OGR_G_WkbSizeEx(g.cValue))
	return
}

func ogrGImportFromWkt(g OGRGeometry, wkt string) (result OGRErr) {
	cs := C.CString(wkt)
	defer C.free(unsafe.Pointer(cs))
	pcs := cs
	result = OGRErr(C.OGR_G_ImportFromWkt(g.cValue, &pcs))
	return
}

func ogrGExportToWkt(g OGRGeometry) (result string, status OGRErr) {
	var cs *C.char
	status = OGRErr(C.OGR_G_ExportToWkt(g.cValue, &cs))
	if cs != nil {
		result = C.GoString(cs)
		vsiFree(unsafe.Pointer(cs))
	}
	return
}

func ogrGExportToIsoWkt(g OGRGeometry) (result string, status OGRErr) {
	var cs *C.char
	status = OGRErr(C.OGR_G_ExportToIsoWkt(g.cValue, &cs))
	if cs != nil {
		result = C.GoString(cs)
		vsiFree(unsafe.Pointer(cs))
	}
	return
}

func ogrGGetGeometryType(g OGRGeometry) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_G_GetGeometryType(g.cValue))
	return
}

func ogrGGetGeometryName(g OGRGeometry) (result string) {
	result = C.GoString(C.OGR_G_GetGeometryName(g.cValue))
	return
}

// void CPL_DLL OGR_G_DumpReadable(OGRGeometryH, FILE *, const char *);

func ogrGFlattenTo2D(g OGRGeometry) {
	C.OGR_G_FlattenTo2D(g.cValue)
}

func ogrGCloseRings(g OGRGeometry) {
	C.OGR_G_CloseRings(g.cValue)
}

func ogrGCreateFromGML(gml string) (result OGRGeometry) {
	cs := C.CString(gml)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeometry{cValue: C.OGR_G_CreateFromGML(cs)}
	return
}

func ogrGExportToGML(g OGRGeometry) (result string) {
	cs := C.OGR_G_ExportToGML(g.cValue)
	if cs != nil {
		result = C.GoString(cs)
		vsiFree(unsafe.Pointer(cs))
	}
	return
}

func ogrGExportToGMLEx(g OGRGeometry, options CSLConstList) (result string) {
	cOptions := options.cValue
	cs := C.OGR_G_ExportToGMLEx(g.cValue, cOptions)
	if cs != nil {
		result = C.GoString(cs)
		vsiFree(unsafe.Pointer(cs))
	}
	return
}

func ogrGCreateFromGMLTree(tree CPLXMLNode) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_CreateFromGMLTree(tree.cValue)}
	return
}

func ogrGExportToGMLTree(g OGRGeometry) (result CPLXMLNode) {
	result = CPLXMLNode{cValue: C.OGR_G_ExportToGMLTree(g.cValue)}
	return
}

func ogrGExportEnvelopeToGMLTree(g OGRGeometry) (result CPLXMLNode) {
	result = CPLXMLNode{cValue: C.OGR_G_ExportEnvelopeToGMLTree(g.cValue)}
	return
}

func ogrGExportToKML(g OGRGeometry, altitudeMode string) (result string) {
	var cs *C.char
	if altitudeMode != "" {
		cs = C.CString(altitudeMode)
		defer C.free(unsafe.Pointer(cs))
	}
	raw := C.OGR_G_ExportToKML(g.cValue, cs)
	if raw != nil {
		result = C.GoString(raw)
		vsiFree(unsafe.Pointer(raw))
	}
	return
}

func ogrGExportToJson(g OGRGeometry) (result string) {
	cs := C.OGR_G_ExportToJson(g.cValue)
	if cs != nil {
		result = C.GoString(cs)
		vsiFree(unsafe.Pointer(cs))
	}
	return
}

func ogrGExportToJsonEx(g OGRGeometry, options CSLConstList) (result string) {
	cOptions := options.cValue
	cs := C.OGR_G_ExportToJsonEx(g.cValue, cOptions)
	if cs != nil {
		result = C.GoString(cs)
		vsiFree(unsafe.Pointer(cs))
	}
	return
}

// /** Create a OGR geometry from a GeoJSON geometry object */
func ogrGCreateGeometryFromJson(json string) (result OGRGeometry) {
	cs := C.CString(json)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeometry{cValue: C.OGR_G_CreateGeometryFromJson(cs)}
	return
}

// /** Create a OGR geometry from a ESRI JSON geometry object */
func ogrGCreateGeometryFromEsriJson(json string) (result OGRGeometry) {
	cs := C.CString(json)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeometry{cValue: C.OGR_G_CreateGeometryFromEsriJson(cs)}
	return
}

func ogrGAssignSpatialReference(g OGRGeometry, sr OGRSpatialReference) {
	C.OGR_G_AssignSpatialReference(g.cValue, sr.cValue)
}

func ogrGGetSpatialReference(g OGRGeometry) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OGR_G_GetSpatialReference(g.cValue)}
	return
}

func ogrGTransform(g OGRGeometry, ct OGRCoordinateTransformation) (result OGRErr) {
	result = OGRErr(C.OGR_G_Transform(g.cValue, ct.cValue))
	return
}

func ogrGTransformTo(g OGRGeometry, sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OGR_G_TransformTo(g.cValue, sr.cValue))
	return
}

func ogrGeomTransformerCreate(ct OGRCoordinateTransformation, options CSLConstList) (result OGRGeomTransformer) {
	cOptions := options.cValue
	result = OGRGeomTransformer{cValue: C.OGR_GeomTransformer_Create(ct.cValue, cOptions)}
	return
}

func ogrGeomTransformerTransform(t OGRGeomTransformer, g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_GeomTransformer_Transform(t.cValue, g.cValue)}
	return
}

func ogrGeomTransformerDestroy(t OGRGeomTransformer) {
	C.OGR_GeomTransformer_Destroy(t.cValue)
}

func ogrGSimplify(g OGRGeometry, tolerance float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Simplify(g.cValue, C.double(tolerance))}
	return
}

func ogrGSimplifyPreserveTopology(g OGRGeometry, tolerance float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_SimplifyPreserveTopology(g.cValue, C.double(tolerance))}
	return
}

func ogrGDelaunayTriangulation(g OGRGeometry, tolerance float64, onlyEdges int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_DelaunayTriangulation(g.cValue, C.double(tolerance), C.int(onlyEdges))}
	return
}

func ogrGConstrainedDelaunayTriangulation(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ConstrainedDelaunayTriangulation(g.cValue)}
	return
}

func ogrGSegmentize(g OGRGeometry, maxLength float64) {
	C.OGR_G_Segmentize(g.cValue, C.double(maxLength))
}

func ogrGIntersects(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Intersects(g.cValue, other.cValue))
	return
}

func ogrGEquals(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Equals(g.cValue, other.cValue))
	return
}

func ogrGDisjoint(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Disjoint(g.cValue, other.cValue))
	return
}

func ogrGTouches(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Touches(g.cValue, other.cValue))
	return
}

func ogrGCrosses(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Crosses(g.cValue, other.cValue))
	return
}

func ogrGWithin(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Within(g.cValue, other.cValue))
	return
}

func ogrGContains(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Contains(g.cValue, other.cValue))
	return
}

func ogrGOverlaps(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Overlaps(g.cValue, other.cValue))
	return
}

func ogrGBoundary(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Boundary(g.cValue)}
	return
}

func ogrGConvexHull(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ConvexHull(g.cValue)}
	return
}

func ogrGConcaveHull(g OGRGeometry, ratio float64, allowHoles bool) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ConcaveHull(g.cValue, C.double(ratio), C.bool(allowHoles))}
	return
}

func ogrGBuffer(g OGRGeometry, dist float64, quadSegs int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Buffer(g.cValue, C.double(dist), C.int(quadSegs))}
	return
}

func ogrGBufferEx(g OGRGeometry, dist float64, options CSLConstList) (result OGRGeometry) {
	cOptions := options.cValue
	result = OGRGeometry{cValue: C.OGR_G_BufferEx(g.cValue, C.double(dist), cOptions)}
	return
}

func ogrGIntersection(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Intersection(g.cValue, other.cValue)}
	return
}

func ogrGUnion(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Union(g.cValue, other.cValue)}
	return
}

func ogrGUnionCascaded(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_UnionCascaded(g.cValue)}
	return
}

func ogrGUnaryUnion(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_UnaryUnion(g.cValue)}
	return
}

func ogrGPointOnSurface(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_PointOnSurface(g.cValue)}
	return
}

func ogrGDifference(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Difference(g.cValue, other.cValue)}
	return
}

func ogrGSymDifference(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_SymDifference(g.cValue, other.cValue)}
	return
}

func ogrGDistance(g, other OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Distance(g.cValue, other.cValue))
	return
}

func ogrGDistance3D(g, other OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Distance3D(g.cValue, other.cValue))
	return
}

func ogrGLength(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Length(g.cValue))
	return
}

func ogrGGeodesicLength(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_GeodesicLength(g.cValue))
	return
}

func ogrGArea(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Area(g.cValue))
	return
}

func ogrGGeodesicArea(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_GeodesicArea(g.cValue))
	return
}

func ogrGIsClockwise(g OGRGeometry) (result bool) {
	result = bool(C.OGR_G_IsClockwise(g.cValue))
	return
}

func ogrGCentroid(g, centroid OGRGeometry) (result int) {
	result = int(C.OGR_G_Centroid(g.cValue, centroid.cValue))
	return
}

func ogrGValue(g OGRGeometry, distance float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Value(g.cValue, C.double(distance))}
	return
}

func ogrGEmpty(g OGRGeometry) {
	C.OGR_G_Empty(g.cValue)
}

func ogrGIsEmpty(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsEmpty(g.cValue))
	return
}

func ogrGIsValid(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsValid(g.cValue))
	return
}

// /*char    CPL_DLL *OGR_G_IsValidReason( OGRGeometryH );*/

func ogrGMakeValid(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_MakeValid(g.cValue)}
	return
}

func ogrGMakeValidEx(g OGRGeometry, options CSLConstList) (result OGRGeometry) {
	cOptions := options.cValue
	result = OGRGeometry{cValue: C.OGR_G_MakeValidEx(g.cValue, cOptions)}
	return
}

func ogrGNormalize(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Normalize(g.cValue)}
	return
}

func ogrGIsSimple(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsSimple(g.cValue))
	return
}

func ogrGIsRing(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsRing(g.cValue))
	return
}

// /** This option causes OGR_G_SetPrecision()
//   - to not attempt at preserving the topology */
const OGRGeosPrecNoTopo = C.OGR_GEOS_PREC_NO_TOPO

// /** This option causes OGR_G_SetPrecision()
//   - to retain collapsed elements */
const OGRGeosPrecKeepCollapsed = C.OGR_GEOS_PREC_KEEP_COLLAPSED

func ogrGSetPrecision(g OGRGeometry, gridSize float64, flags int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_SetPrecision(g.cValue, C.double(gridSize), C.int(flags))}
	return
}

func ogrGPolygonize(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Polygonize(g.cValue)}
	return
}

func ogrGBuildArea(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_BuildArea(g.cValue)}
	return
}

// /*! @cond Doxygen_Suppress */
// /* backward compatibility (non-standard methods) */
// int CPL_DLL OGR_G_Intersect(OGRGeometryH, OGRGeometryH) CPL_WARN_DEPRECATED("Non standard method. Use OGR_G_Intersects() instead");
// int CPL_DLL OGR_G_Equal(OGRGeometryH, OGRGeometryH) CPL_WARN_DEPRECATED("Non standard method. Use OGR_G_Equals() instead");
// OGRGeometryH CPL_DLL OGR_G_SymmetricDifference(OGRGeometryH, OGRGeometryH) CPL_WARN_DEPRECATED("Non standard method. Use OGR_G_SymDifference() instead");
// double CPL_DLL OGR_G_GetArea(OGRGeometryH) CPL_WARN_DEPRECATED("Non standard method. Use OGR_G_Area() instead");
// OGRGeometryH CPL_DLL OGR_G_GetBoundary(OGRGeometryH) CPL_WARN_DEPRECATED("Non standard method. Use OGR_G_Boundary() instead");
// /*! @endcond */

// /* Methods for getting/setting vertices in points, line strings and rings */

func ogrGGetPointCount(g OGRGeometry) (result int) {
	result = int(C.OGR_G_GetPointCount(g.cValue))
	return
}

func ogrGGetPoints(g OGRGeometry) (x, y, z []float64) {
	n := int(C.OGR_G_GetPointCount(g.cValue))
	if n == 0 {
		return
	}
	x = make([]float64, n)
	y = make([]float64, n)
	z = make([]float64, n)
	stride := C.int(unsafe.Sizeof(C.double(0)))
	C.OGR_G_GetPoints(g.cValue, unsafe.Pointer(&x[0]), stride, unsafe.Pointer(&y[0]), stride, unsafe.Pointer(&z[0]), stride)
	return
}

func ogrGGetPointsZM(g OGRGeometry) (x, y, z, m []float64) {
	n := int(C.OGR_G_GetPointCount(g.cValue))
	if n == 0 {
		return
	}
	x = make([]float64, n)
	y = make([]float64, n)
	z = make([]float64, n)
	m = make([]float64, n)
	stride := C.int(unsafe.Sizeof(C.double(0)))
	C.OGR_G_GetPointsZM(g.cValue, unsafe.Pointer(&x[0]), stride, unsafe.Pointer(&y[0]), stride, unsafe.Pointer(&z[0]), stride, unsafe.Pointer(&m[0]), stride)
	return
}

func ogrGGetX(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetX(g.cValue, C.int(point)))
	return
}

func ogrGGetY(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetY(g.cValue, C.int(point)))
	return
}

func ogrGGetZ(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetZ(g.cValue, C.int(point)))
	return
}

func ogrGGetM(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetM(g.cValue, C.int(point)))
	return
}

func ogrGGetPoint(g OGRGeometry, point int) (x, y, z float64) {
	var cx, cy, cz C.double
	C.OGR_G_GetPoint(g.cValue, C.int(point), &cx, &cy, &cz)
	x, y, z = float64(cx), float64(cy), float64(cz)
	return
}

func ogrGGetPointZM(g OGRGeometry, point int) (x, y, z, m float64) {
	var cx, cy, cz, cm C.double
	C.OGR_G_GetPointZM(g.cValue, C.int(point), &cx, &cy, &cz, &cm)
	x, y, z, m = float64(cx), float64(cy), float64(cz), float64(cm)
	return
}

func ogrGSetPointCount(g OGRGeometry, count int) {
	C.OGR_G_SetPointCount(g.cValue, C.int(count))
}

func ogrGSetPoint(g OGRGeometry, point int, x, y, z float64) {
	C.OGR_G_SetPoint(g.cValue, C.int(point), C.double(x), C.double(y), C.double(z))
}

func ogrGSetPoint2D(g OGRGeometry, point int, x, y float64) {
	C.OGR_G_SetPoint_2D(g.cValue, C.int(point), C.double(x), C.double(y))
}

func ogrGSetPointM(g OGRGeometry, point int, x, y, m float64) {
	C.OGR_G_SetPointM(g.cValue, C.int(point), C.double(x), C.double(y), C.double(m))
}

func ogrGSetPointZM(g OGRGeometry, point int, x, y, z, m float64) {
	C.OGR_G_SetPointZM(g.cValue, C.int(point), C.double(x), C.double(y), C.double(z), C.double(m))
}

func ogrGAddPoint(g OGRGeometry, x, y, z float64) {
	C.OGR_G_AddPoint(g.cValue, C.double(x), C.double(y), C.double(z))
}

func ogrGAddPoint2D(g OGRGeometry, x, y float64) {
	C.OGR_G_AddPoint_2D(g.cValue, C.double(x), C.double(y))
}

func ogrGAddPointM(g OGRGeometry, x, y, m float64) {
	C.OGR_G_AddPointM(g.cValue, C.double(x), C.double(y), C.double(m))
}

func ogrGAddPointZM(g OGRGeometry, x, y, z, m float64) {
	C.OGR_G_AddPointZM(g.cValue, C.double(x), C.double(y), C.double(z), C.double(m))
}

func ogrGSetPoints(g OGRGeometry, x, y, z []float64) {
	stride := C.int(unsafe.Sizeof(C.double(0)))
	var px, py, pz unsafe.Pointer
	if len(x) > 0 {
		px = unsafe.Pointer(&x[0])
	}
	if len(y) > 0 {
		py = unsafe.Pointer(&y[0])
	}
	if len(z) > 0 {
		pz = unsafe.Pointer(&z[0])
	}
	C.OGR_G_SetPoints(g.cValue, C.int(len(x)), px, stride, py, stride, pz, stride)
}

func ogrGSetPointsZM(g OGRGeometry, x, y, z, m []float64) {
	stride := C.int(unsafe.Sizeof(C.double(0)))
	var px, py, pz, pm unsafe.Pointer
	if len(x) > 0 {
		px = unsafe.Pointer(&x[0])
	}
	if len(y) > 0 {
		py = unsafe.Pointer(&y[0])
	}
	if len(z) > 0 {
		pz = unsafe.Pointer(&z[0])
	}
	if len(m) > 0 {
		pm = unsafe.Pointer(&m[0])
	}
	C.OGR_G_SetPointsZM(g.cValue, C.int(len(x)), px, stride, py, stride, pz, stride, pm, stride)
}

func ogrGSwapXY(g OGRGeometry) {
	C.OGR_G_SwapXY(g.cValue)
}

// /* Methods for getting/setting rings and members collections */

func ogrGGetGeometryCount(g OGRGeometry) (result int) {
	result = int(C.OGR_G_GetGeometryCount(g.cValue))
	return
}

func ogrGGetGeometryRef(g OGRGeometry, subGeom int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_GetGeometryRef(g.cValue, C.int(subGeom))}
	return
}

func ogrGAddGeometry(g, subGeom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_G_AddGeometry(g.cValue, subGeom.cValue))
	return
}

func ogrGAddGeometryDirectly(g, subGeom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_G_AddGeometryDirectly(g.cValue, subGeom.cValue))
	return
}

func ogrGRemoveGeometry(g OGRGeometry, subGeom, delete int) (result OGRErr) {
	result = OGRErr(C.OGR_G_RemoveGeometry(g.cValue, C.int(subGeom), C.int(delete)))
	return
}

func ogrGHasCurveGeometry(g OGRGeometry, lookForNonLinear int) (result int) {
	result = int(C.OGR_G_HasCurveGeometry(g.cValue, C.int(lookForNonLinear)))
	return
}

func ogrGGetLinearGeometry(g OGRGeometry, maxAngleStepSizeDegrees float64, options CSLConstList) (result OGRGeometry) {
	cOptions := options.cValue
	result = OGRGeometry{cValue: C.OGR_G_GetLinearGeometry(g.cValue, C.double(maxAngleStepSizeDegrees), cOptions)}
	return
}

func ogrGGetCurveGeometry(g OGRGeometry, options CSLConstList) (result OGRGeometry) {
	cOptions := options.cValue
	result = OGRGeometry{cValue: C.OGR_G_GetCurveGeometry(g.cValue, cOptions)}
	return
}

func ogrBuildPolygonFromEdges(lines OGRGeometry, bestEffort, autoClose int, tolerance float64) (result OGRGeometry, status OGRErr) {
	var e C.OGRErr
	result = OGRGeometry{cValue: C.OGRBuildPolygonFromEdges(lines.cValue, C.int(bestEffort), C.int(autoClose), C.double(tolerance), &e)}
	status = OGRErr(e)
	return
}

// /*! @cond Doxygen_Suppress */

func ogrSetGenerateDB2V72ByteOrder(generate int) (result OGRErr) {
	result = OGRErr(C.OGRSetGenerate_DB2_V72_BYTE_ORDER(C.int(generate)))
	return
}

func ogrGetGenerateDB2V72ByteOrder() (result int) {
	result = int(C.OGRGetGenerate_DB2_V72_BYTE_ORDER())
	return
}

// /*! @endcond */

func ogrSetNonLinearGeometriesEnabledFlag(flag int) {
	C.OGRSetNonLinearGeometriesEnabledFlag(C.int(flag))
}

func ogrGetNonLinearGeometriesEnabledFlag() (result int) {
	result = int(C.OGRGetNonLinearGeometriesEnabledFlag())
	return
}

func ogrHasPreparedGeometrySupport() (result bool) {
	result = C.OGRHasPreparedGeometrySupport() != 0
	return
}

func ogrCreatePreparedGeometry(g OGRGeometry) (result OGRPreparedGeometry) {
	result = OGRPreparedGeometry{cValue: C.OGRCreatePreparedGeometry(g.cValue)}
	return
}

func ogrDestroyPreparedGeometry(p OGRPreparedGeometry) {
	C.OGRDestroyPreparedGeometry(p.cValue)
}

func ogrPreparedGeometryIntersects(p OGRPreparedGeometry, other OGRGeometry) (result int) {
	result = int(C.OGRPreparedGeometryIntersects(p.cValue, other.cValue))
	return
}

func ogrPreparedGeometryContains(p OGRPreparedGeometry, other OGRGeometry) (result int) {
	result = int(C.OGRPreparedGeometryContains(p.cValue, other.cValue))
	return
}

// /* -------------------------------------------------------------------- */
// /*      Feature related (ogr_feature.h)                                 */
// /* -------------------------------------------------------------------- */

// /* OGRFieldDefn */

func ogrFldCreate(name string, eType OGRFieldType) (result OGRFieldDefn) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRFieldDefn{cValue: C.OGR_Fld_Create(cs, C.OGRFieldType(eType))}
	return
}

func ogrFldDestroy(fld OGRFieldDefn) {
	C.OGR_Fld_Destroy(fld.cValue)
}

func ogrFldSetName(fld OGRFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetName(fld.cValue, cs)
}

func ogrFldGetNameRef(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetNameRef(fld.cValue))
	return
}

func ogrFldSetAlternativeName(fld OGRFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetAlternativeName(fld.cValue, cs)
}

func ogrFldGetAlternativeNameRef(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetAlternativeNameRef(fld.cValue))
	return
}

func ogrFldGetType(fld OGRFieldDefn) (result OGRFieldType) {
	result = OGRFieldType(C.OGR_Fld_GetType(fld.cValue))
	return
}

func ogrFldSetType(fld OGRFieldDefn, eType OGRFieldType) {
	C.OGR_Fld_SetType(fld.cValue, C.OGRFieldType(eType))
}

func ogrFldGetSubType(fld OGRFieldDefn) (result OGRFieldSubType) {
	result = OGRFieldSubType(C.OGR_Fld_GetSubType(fld.cValue))
	return
}

func ogrFldSetSubType(fld OGRFieldDefn, eSubType OGRFieldSubType) {
	C.OGR_Fld_SetSubType(fld.cValue, C.OGRFieldSubType(eSubType))
}

func ogrFldGetJustify(fld OGRFieldDefn) (result OGRJustification) {
	result = OGRJustification(C.OGR_Fld_GetJustify(fld.cValue))
	return
}

func ogrFldSetJustify(fld OGRFieldDefn, eJustify OGRJustification) {
	C.OGR_Fld_SetJustify(fld.cValue, C.OGRJustification(eJustify))
}

func ogrFldGetWidth(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_GetWidth(fld.cValue))
	return
}

func ogrFldSetWidth(fld OGRFieldDefn, width int) {
	C.OGR_Fld_SetWidth(fld.cValue, C.int(width))
}

func ogrFldGetPrecision(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_GetPrecision(fld.cValue))
	return
}

func ogrFldSetPrecision(fld OGRFieldDefn, precision int) {
	C.OGR_Fld_SetPrecision(fld.cValue, C.int(precision))
}

func ogrFldGetTZFlag(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_GetTZFlag(fld.cValue))
	return
}

func ogrFldSetTZFlag(fld OGRFieldDefn, tzFlag int) {
	C.OGR_Fld_SetTZFlag(fld.cValue, C.int(tzFlag))
}

func ogrFldSet(fld OGRFieldDefn, name string, eType OGRFieldType, width, precision int, justify OGRJustification) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_Set(fld.cValue, cs, C.OGRFieldType(eType), C.int(width), C.int(precision), C.OGRJustification(justify))
}

func ogrFldIsIgnored(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsIgnored(fld.cValue))
	return
}

func ogrFldSetIgnored(fld OGRFieldDefn, ignored int) {
	C.OGR_Fld_SetIgnored(fld.cValue, C.int(ignored))
}

func ogrFldIsNullable(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsNullable(fld.cValue))
	return
}

func ogrFldSetNullable(fld OGRFieldDefn, nullable int) {
	C.OGR_Fld_SetNullable(fld.cValue, C.int(nullable))
}

func ogrFldSetGenerated(fld OGRFieldDefn, generated int) {
	C.OGR_Fld_SetGenerated(fld.cValue, C.int(generated))
}

func ogrFldIsGenerated(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsGenerated(fld.cValue))
	return
}

func ogrFldIsUnique(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsUnique(fld.cValue))
	return
}

func ogrFldSetUnique(fld OGRFieldDefn, unique int) {
	C.OGR_Fld_SetUnique(fld.cValue, C.int(unique))
}

func ogrFldGetDefault(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetDefault(fld.cValue))
	return
}

func ogrFldSetDefault(fld OGRFieldDefn, value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetDefault(fld.cValue, cs)
}

func ogrFldIsDefaultDriverSpecific(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsDefaultDriverSpecific(fld.cValue))
	return
}

func ogrFldGetDomainName(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetDomainName(fld.cValue))
	return
}

func ogrFldSetDomainName(fld OGRFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetDomainName(fld.cValue, cs)
}

func ogrFldGetComment(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetComment(fld.cValue))
	return
}

func ogrFldSetComment(fld OGRFieldDefn, comment string) {
	cs := C.CString(comment)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetComment(fld.cValue, cs)
}

func ogrGetFieldTypeName(eType OGRFieldType) (result string) {
	result = C.GoString(C.OGR_GetFieldTypeName(C.OGRFieldType(eType)))
	return
}

func ogrGetFieldTypeByName(name string) (result OGRFieldType) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRFieldType(C.OGR_GetFieldTypeByName(cs))
	return
}

func ogrGetFieldSubTypeName(eSubType OGRFieldSubType) (result string) {
	result = C.GoString(C.OGR_GetFieldSubTypeName(C.OGRFieldSubType(eSubType)))
	return
}

func ogrGetFieldSubTypeByName(name string) (result OGRFieldSubType) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRFieldSubType(C.OGR_GetFieldSubTypeByName(cs))
	return
}

func ogrAreTypeSubTypeCompatible(eType OGRFieldType, eSubType OGRFieldSubType) (result int) {
	result = int(C.OGR_AreTypeSubTypeCompatible(C.OGRFieldType(eType), C.OGRFieldSubType(eSubType)))
	return
}

// /* OGRGeomFieldDefnH */

func ogrGFldCreate(name string, eType OGRwkbGeometryType) (result OGRGeomFieldDefn) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeomFieldDefn{cValue: C.OGR_GFld_Create(cs, C.OGRwkbGeometryType(eType))}
	return
}

func ogrGFldDestroy(gfld OGRGeomFieldDefn) {
	C.OGR_GFld_Destroy(gfld.cValue)
}

func ogrGFldSetName(gfld OGRGeomFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_GFld_SetName(gfld.cValue, cs)
}

func ogrGFldGetNameRef(gfld OGRGeomFieldDefn) (result string) {
	result = C.GoString(C.OGR_GFld_GetNameRef(gfld.cValue))
	return
}

func ogrGFldGetType(gfld OGRGeomFieldDefn) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GFld_GetType(gfld.cValue))
	return
}

func ogrGFldSetType(gfld OGRGeomFieldDefn, eType OGRwkbGeometryType) {
	C.OGR_GFld_SetType(gfld.cValue, C.OGRwkbGeometryType(eType))
}

func ogrGFldGetSpatialRef(gfld OGRGeomFieldDefn) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OGR_GFld_GetSpatialRef(gfld.cValue)}
	return
}

func ogrGFldSetSpatialRef(gfld OGRGeomFieldDefn, sr OGRSpatialReference) {
	C.OGR_GFld_SetSpatialRef(gfld.cValue, sr.cValue)
}

func ogrGFldIsNullable(gfld OGRGeomFieldDefn) (result int) {
	result = int(C.OGR_GFld_IsNullable(gfld.cValue))
	return
}

func ogrGFldSetNullable(gfld OGRGeomFieldDefn, nullable int) {
	C.OGR_GFld_SetNullable(gfld.cValue, C.int(nullable))
}

func ogrGFldIsIgnored(gfld OGRGeomFieldDefn) (result int) {
	result = int(C.OGR_GFld_IsIgnored(gfld.cValue))
	return
}

func ogrGFldSetIgnored(gfld OGRGeomFieldDefn, ignored int) {
	C.OGR_GFld_SetIgnored(gfld.cValue, C.int(ignored))
}

func ogrGFldGetCoordinatePrecision(gfld OGRGeomFieldDefn) (result OGRGeomCoordinatePrecision) {
	result = OGRGeomCoordinatePrecision{cValue: C.OGR_GFld_GetCoordinatePrecision(gfld.cValue)}
	return
}

func ogrGFldSetCoordinatePrecision(gfld OGRGeomFieldDefn, precision OGRGeomCoordinatePrecision) {
	C.OGR_GFld_SetCoordinatePrecision(gfld.cValue, precision.cValue)
}

// /* OGRFeatureDefn */

func ogrFDCreate(name string) (result OGRFeatureDefn) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRFeatureDefn{cValue: C.OGR_FD_Create(cs)}
	return
}

func ogrFDDestroy(fd OGRFeatureDefn) {
	C.OGR_FD_Destroy(fd.cValue)
}

func ogrFDRelease(fd OGRFeatureDefn) {
	C.OGR_FD_Release(fd.cValue)
}

func ogrFDGetName(fd OGRFeatureDefn) (result string) {
	result = C.GoString(C.OGR_FD_GetName(fd.cValue))
	return
}

func ogrFDGetFieldCount(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_GetFieldCount(fd.cValue))
	return
}

func ogrFDGetFieldDefn(fd OGRFeatureDefn, field int) (result OGRFieldDefn) {
	result = OGRFieldDefn{cValue: C.OGR_FD_GetFieldDefn(fd.cValue, C.int(field))}
	return
}

func ogrFDGetFieldIndex(fd OGRFeatureDefn, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_FD_GetFieldIndex(fd.cValue, cs))
	return
}

func ogrFDAddFieldDefn(fd OGRFeatureDefn, fld OGRFieldDefn) {
	C.OGR_FD_AddFieldDefn(fd.cValue, fld.cValue)
}

func ogrFDDeleteFieldDefn(fd OGRFeatureDefn, field int) (result OGRErr) {
	result = OGRErr(C.OGR_FD_DeleteFieldDefn(fd.cValue, C.int(field)))
	return
}

func ogrFDReorderFieldDefns(fd OGRFeatureDefn, panMap []int) (result OGRErr) {
	cMap := make([]C.int, len(panMap))
	for i, v := range panMap {
		cMap[i] = C.int(v)
	}
	var p *C.int
	if len(cMap) > 0 {
		p = &cMap[0]
	}
	result = OGRErr(C.OGR_FD_ReorderFieldDefns(fd.cValue, p))
	return
}

func ogrFDGetGeomType(fd OGRFeatureDefn) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_FD_GetGeomType(fd.cValue))
	return
}

func ogrFDSetGeomType(fd OGRFeatureDefn, eType OGRwkbGeometryType) {
	C.OGR_FD_SetGeomType(fd.cValue, C.OGRwkbGeometryType(eType))
}

func ogrFDIsGeometryIgnored(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_IsGeometryIgnored(fd.cValue))
	return
}

func ogrFDSetGeometryIgnored(fd OGRFeatureDefn, ignored int) {
	C.OGR_FD_SetGeometryIgnored(fd.cValue, C.int(ignored))
}

func ogrFDIsStyleIgnored(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_IsStyleIgnored(fd.cValue))
	return
}

func ogrFDSetStyleIgnored(fd OGRFeatureDefn, ignored int) {
	C.OGR_FD_SetStyleIgnored(fd.cValue, C.int(ignored))
}

func ogrFDReference(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_Reference(fd.cValue))
	return
}

func ogrFDDereference(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_Dereference(fd.cValue))
	return
}

func ogrFDGetReferenceCount(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_GetReferenceCount(fd.cValue))
	return
}

func ogrFDGetGeomFieldCount(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_GetGeomFieldCount(fd.cValue))
	return
}

func ogrFDGetGeomFieldDefn(fd OGRFeatureDefn, geomField int) (result OGRGeomFieldDefn) {
	result = OGRGeomFieldDefn{cValue: C.OGR_FD_GetGeomFieldDefn(fd.cValue, C.int(geomField))}
	return
}

func ogrFDGetGeomFieldIndex(fd OGRFeatureDefn, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_FD_GetGeomFieldIndex(fd.cValue, cs))
	return
}

func ogrFDAddGeomFieldDefn(fd OGRFeatureDefn, gfld OGRGeomFieldDefn) {
	C.OGR_FD_AddGeomFieldDefn(fd.cValue, gfld.cValue)
}

func ogrFDDeleteGeomFieldDefn(fd OGRFeatureDefn, geomField int) (result OGRErr) {
	result = OGRErr(C.OGR_FD_DeleteGeomFieldDefn(fd.cValue, C.int(geomField)))
	return
}

func ogrFDIsSame(fd, other OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_IsSame(fd.cValue, other.cValue))
	return
}

// /* OGRFeature */

func ogrFCreate(fd OGRFeatureDefn) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_F_Create(fd.cValue)}
	return
}

func ogrFDestroy(feat OGRFeature) {
	C.OGR_F_Destroy(feat.cValue)
}

func ogrFGetDefnRef(feat OGRFeature) (result OGRFeatureDefn) {
	result = OGRFeatureDefn{cValue: C.OGR_F_GetDefnRef(feat.cValue)}
	return
}

func ogrFSetGeometryDirectly(feat OGRFeature, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeometryDirectly(feat.cValue, geom.cValue))
	return
}

func ogrFSetGeometry(feat OGRFeature, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeometry(feat.cValue, geom.cValue))
	return
}

func ogrFGetGeometryRef(feat OGRFeature) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_GetGeometryRef(feat.cValue)}
	return
}

func ogrFStealGeometry(feat OGRFeature) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_StealGeometry(feat.cValue)}
	return
}

func ogrFStealGeometryEx(feat OGRFeature, geomField int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_StealGeometryEx(feat.cValue, C.int(geomField))}
	return
}

func ogrFClone(feat OGRFeature) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_F_Clone(feat.cValue)}
	return
}

func ogrFEqual(feat, other OGRFeature) (result int) {
	result = int(C.OGR_F_Equal(feat.cValue, other.cValue))
	return
}

func ogrFGetFieldCount(feat OGRFeature) (result int) {
	result = int(C.OGR_F_GetFieldCount(feat.cValue))
	return
}

func ogrFGetFieldDefnRef(feat OGRFeature, field int) (result OGRFieldDefn) {
	result = OGRFieldDefn{cValue: C.OGR_F_GetFieldDefnRef(feat.cValue, C.int(field))}
	return
}

func ogrFGetFieldIndex(feat OGRFeature, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_F_GetFieldIndex(feat.cValue, cs))
	return
}

func ogrFIsFieldSet(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_IsFieldSet(feat.cValue, C.int(field)))
	return
}

func ogrFUnsetField(feat OGRFeature, field int) {
	C.OGR_F_UnsetField(feat.cValue, C.int(field))
}

func ogrFIsFieldNull(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_IsFieldNull(feat.cValue, C.int(field)))
	return
}

func ogrFIsFieldSetAndNotNull(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_IsFieldSetAndNotNull(feat.cValue, C.int(field)))
	return
}

func ogrFSetFieldNull(feat OGRFeature, field int) {
	C.OGR_F_SetFieldNull(feat.cValue, C.int(field))
}

func ogrFGetRawFieldRef(feat OGRFeature, field int) (result OGRField) {
	result = OGRField{cValue: C.OGR_F_GetRawFieldRef(feat.cValue, C.int(field))}
	return
}

func ogrRawFieldIsUnset(f OGRField) (result int) {
	result = int(C.OGR_RawField_IsUnset(f.cValue))
	return
}

func ogrRawFieldIsNull(f OGRField) (result int) {
	result = int(C.OGR_RawField_IsNull(f.cValue))
	return
}

func ogrRawFieldSetUnset(f OGRField) {
	C.OGR_RawField_SetUnset(f.cValue)
}

func ogrRawFieldSetNull(f OGRField) {
	C.OGR_RawField_SetNull(f.cValue)
}

func ogrFGetFieldAsInteger(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_GetFieldAsInteger(feat.cValue, C.int(field)))
	return
}

func ogrFGetFieldAsInteger64(feat OGRFeature, field int) (result int64) {
	result = int64(C.OGR_F_GetFieldAsInteger64(feat.cValue, C.int(field)))
	return
}

func ogrFGetFieldAsDouble(feat OGRFeature, field int) (result float64) {
	result = float64(C.OGR_F_GetFieldAsDouble(feat.cValue, C.int(field)))
	return
}

func ogrFGetFieldAsString(feat OGRFeature, field int) (result string) {
	result = C.GoString(C.OGR_F_GetFieldAsString(feat.cValue, C.int(field)))
	return
}

func ogrFGetFieldAsISO8601DateTime(feat OGRFeature, field int, options CSLConstList) (result string) {
	cOptions := options.cValue
	result = C.GoString(C.OGR_F_GetFieldAsISO8601DateTime(feat.cValue, C.int(field), cOptions))
	return
}

func ogrFGetFieldAsIntegerList(feat OGRFeature, field int) (result []int) {
	var count C.int
	p := C.OGR_F_GetFieldAsIntegerList(feat.cValue, C.int(field), &count)
	n := int(count)
	if n == 0 || p == nil {
		return
	}
	slice := unsafe.Slice(p, n)
	result = make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = int(slice[i])
	}
	return
}

func ogrFGetFieldAsInteger64List(feat OGRFeature, field int) (result []int64) {
	var count C.int
	p := C.OGR_F_GetFieldAsInteger64List(feat.cValue, C.int(field), &count)
	n := int(count)
	if n == 0 || p == nil {
		return
	}
	slice := unsafe.Slice(p, n)
	result = make([]int64, n)
	for i := 0; i < n; i++ {
		result[i] = int64(slice[i])
	}
	return
}

func ogrFGetFieldAsDoubleList(feat OGRFeature, field int) (result []float64) {
	var count C.int
	p := C.OGR_F_GetFieldAsDoubleList(feat.cValue, C.int(field), &count)
	n := int(count)
	if n == 0 || p == nil {
		return
	}
	slice := unsafe.Slice(p, n)
	result = make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = float64(slice[i])
	}
	return
}

func ogrFGetFieldAsStringList(feat OGRFeature, field int) (result CSLConstList) {
	raw := C.OGR_F_GetFieldAsStringList(feat.cValue, C.int(field))
	result = cslConstList(raw)
	return
}

func ogrFGetFieldAsBinary(feat OGRFeature, field int) (result []byte) {
	var count C.int
	p := C.OGR_F_GetFieldAsBinary(feat.cValue, C.int(field), &count)
	if p == nil || count == 0 {
		return
	}
	result = C.GoBytes(unsafe.Pointer(p), count)
	return
}

func ogrFGetFieldAsDateTime(feat OGRFeature, field int) (year, month, day, hour, minute, second, tzFlag int, ok bool) {
	var cy, cmo, cd, ch, cmi, cse, ctz C.int
	r := C.OGR_F_GetFieldAsDateTime(feat.cValue, C.int(field), &cy, &cmo, &cd, &ch, &cmi, &cse, &ctz)
	year, month, day, hour, minute, second, tzFlag = int(cy), int(cmo), int(cd), int(ch), int(cmi), int(cse), int(ctz)
	ok = r != 0
	return
}

func ogrFGetFieldAsDateTimeEx(feat OGRFeature, field int) (year, month, day, hour, minute int, second float32, tzFlag int, ok bool) {
	var cy, cmo, cd, ch, cmi, ctz C.int
	var cse C.float
	r := C.OGR_F_GetFieldAsDateTimeEx(feat.cValue, C.int(field), &cy, &cmo, &cd, &ch, &cmi, &cse, &ctz)
	year, month, day, hour, minute, tzFlag = int(cy), int(cmo), int(cd), int(ch), int(cmi), int(ctz)
	second = float32(cse)
	ok = r != 0
	return
}

func ogrFSetFieldInteger(feat OGRFeature, field, value int) {
	C.OGR_F_SetFieldInteger(feat.cValue, C.int(field), C.int(value))
}

func ogrFSetFieldInteger64(feat OGRFeature, field int, value int64) {
	C.OGR_F_SetFieldInteger64(feat.cValue, C.int(field), C.GIntBig(value))
}

func ogrFSetFieldDouble(feat OGRFeature, field int, value float64) {
	C.OGR_F_SetFieldDouble(feat.cValue, C.int(field), C.double(value))
}

func ogrFSetFieldString(feat OGRFeature, field int, value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetFieldString(feat.cValue, C.int(field), cs)
}

func ogrFSetFieldIntegerList(feat OGRFeature, field int, values []int) {
	cValues := make([]C.int, len(values))
	for i, v := range values {
		cValues[i] = C.int(v)
	}
	var p *C.int
	if len(cValues) > 0 {
		p = &cValues[0]
	}
	C.OGR_F_SetFieldIntegerList(feat.cValue, C.int(field), C.int(len(values)), p)
}

func ogrFSetFieldInteger64List(feat OGRFeature, field int, values []int64) {
	var p *C.GIntBig
	if len(values) > 0 {
		p = (*C.GIntBig)(unsafe.Pointer(&values[0]))
	}
	C.OGR_F_SetFieldInteger64List(feat.cValue, C.int(field), C.int(len(values)), p)
}

func ogrFSetFieldDoubleList(feat OGRFeature, field int, values []float64) {
	var p *C.double
	if len(values) > 0 {
		p = (*C.double)(unsafe.Pointer(&values[0]))
	}
	C.OGR_F_SetFieldDoubleList(feat.cValue, C.int(field), C.int(len(values)), p)
}

func ogrFSetFieldStringList(feat OGRFeature, field int, values CSLConstList) {
	C.OGR_F_SetFieldStringList(feat.cValue, C.int(field), values.cValue)
}

func ogrFSetFieldRaw(feat OGRFeature, field int, value OGRField) {
	C.OGR_F_SetFieldRaw(feat.cValue, C.int(field), value.cValue)
}

func ogrFSetFieldBinary(feat OGRFeature, field int, data []byte) {
	var p unsafe.Pointer
	if len(data) > 0 {
		p = unsafe.Pointer(&data[0])
	}
	C.OGR_F_SetFieldBinary(feat.cValue, C.int(field), C.int(len(data)), p)
}

func ogrFSetFieldDateTime(feat OGRFeature, field, year, month, day, hour, minute, second, tzFlag int) {
	C.OGR_F_SetFieldDateTime(feat.cValue, C.int(field), C.int(year), C.int(month), C.int(day), C.int(hour), C.int(minute), C.int(second), C.int(tzFlag))
}

func ogrFSetFieldDateTimeEx(feat OGRFeature, field, year, month, day, hour, minute int, second float32, tzFlag int) {
	C.OGR_F_SetFieldDateTimeEx(feat.cValue, C.int(field), C.int(year), C.int(month), C.int(day), C.int(hour), C.int(minute), C.float(second), C.int(tzFlag))
}

func ogrFGetGeomFieldCount(feat OGRFeature) (result int) {
	result = int(C.OGR_F_GetGeomFieldCount(feat.cValue))
	return
}

func ogrFGetGeomFieldDefnRef(feat OGRFeature, field int) (result OGRGeomFieldDefn) {
	result = OGRGeomFieldDefn{cValue: C.OGR_F_GetGeomFieldDefnRef(feat.cValue, C.int(field))}
	return
}

func ogrFGetGeomFieldIndex(feat OGRFeature, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_F_GetGeomFieldIndex(feat.cValue, cs))
	return
}

func ogrFGetGeomFieldRef(feat OGRFeature, field int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_GetGeomFieldRef(feat.cValue, C.int(field))}
	return
}

func ogrFSetGeomFieldDirectly(feat OGRFeature, field int, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeomFieldDirectly(feat.cValue, C.int(field), geom.cValue))
	return
}

func ogrFSetGeomField(feat OGRFeature, field int, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeomField(feat.cValue, C.int(field), geom.cValue))
	return
}

func ogrFGetFID(feat OGRFeature) (result int64) {
	result = int64(C.OGR_F_GetFID(feat.cValue))
	return
}

func ogrFSetFID(feat OGRFeature, fid int64) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetFID(feat.cValue, C.GIntBig(fid)))
	return
}

// void CPL_DLL OGR_F_DumpReadable(OGRFeatureH, FILE *);

func ogrFDumpReadableAsString(feat OGRFeature, options CSLConstList) (result string) {
	cOptions := options.cValue
	cs := C.OGR_F_DumpReadableAsString(feat.cValue, cOptions)
	if cs != nil {
		result = C.GoString(cs)
		vsiFree(unsafe.Pointer(cs))
	}
	return
}

func ogrFSetFrom(feat, other OGRFeature, forgiving int) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetFrom(feat.cValue, other.cValue, C.int(forgiving)))
	return
}

func ogrFSetFromWithMap(feat, other OGRFeature, forgiving int, panMap []int) (result OGRErr) {
	cMap := make([]C.int, len(panMap))
	for i, v := range panMap {
		cMap[i] = C.int(v)
	}
	var p *C.int
	if len(cMap) > 0 {
		p = &cMap[0]
	}
	result = OGRErr(C.OGR_F_SetFromWithMap(feat.cValue, other.cValue, C.int(forgiving), p))
	return
}

func ogrFGetStyleString(feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_F_GetStyleString(feat.cValue))
	return
}

func ogrFSetStyleString(feat OGRFeature, style string) {
	cs := C.CString(style)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetStyleString(feat.cValue, cs)
}

func ogrFSetStyleStringDirectly(feat OGRFeature, style string) {
	cs := C.CString(style)
	C.OGR_F_SetStyleStringDirectly(feat.cValue, cs)
}

// /** Return style table */
func ogrFGetStyleTable(feat OGRFeature) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_F_GetStyleTable(feat.cValue)}
	return
}

// /** Set style table and take ownership */
func ogrFSetStyleTableDirectly(feat OGRFeature, styleTable OGRStyleTable) {
	C.OGR_F_SetStyleTableDirectly(feat.cValue, styleTable.cValue)
}

// /** Set style table */
func ogrFSetStyleTable(feat OGRFeature, styleTable OGRStyleTable) {
	C.OGR_F_SetStyleTable(feat.cValue, styleTable.cValue)
}

func ogrFGetNativeData(feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_F_GetNativeData(feat.cValue))
	return
}

func ogrFSetNativeData(feat OGRFeature, data string) {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetNativeData(feat.cValue, cs)
}

func ogrFGetNativeMediaType(feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_F_GetNativeMediaType(feat.cValue))
	return
}

func ogrFSetNativeMediaType(feat OGRFeature, mediaType string) {
	cs := C.CString(mediaType)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetNativeMediaType(feat.cValue, cs)
}

func ogrFFillUnsetWithDefault(feat OGRFeature, notNullableOnly int, options CSLConstList) {
	cOptions := options.cValue
	C.OGR_F_FillUnsetWithDefault(feat.cValue, C.int(notNullableOnly), cOptions)
}

func ogrFValidate(feat OGRFeature, validateFlags, emitError int) (result int) {
	result = int(C.OGR_F_Validate(feat.cValue, C.int(validateFlags), C.int(emitError)))
	return
}

// /* OGRFieldDomain */

func ogrFldDomainDestroy(dom OGRFieldDomain) {
	C.OGR_FldDomain_Destroy(dom.cValue)
}

func ogrFldDomainGetName(dom OGRFieldDomain) (result string) {
	result = C.GoString(C.OGR_FldDomain_GetName(dom.cValue))
	return
}

func ogrFldDomainGetDescription(dom OGRFieldDomain) (result string) {
	result = C.GoString(C.OGR_FldDomain_GetDescription(dom.cValue))
	return
}

func ogrFldDomainGetDomainType(dom OGRFieldDomain) (result OGRFieldDomainType) {
	result = OGRFieldDomainType(C.OGR_FldDomain_GetDomainType(dom.cValue))
	return
}

func ogrFldDomainGetFieldType(dom OGRFieldDomain) (result OGRFieldType) {
	result = OGRFieldType(C.OGR_FldDomain_GetFieldType(dom.cValue))
	return
}

func ogrFldDomainGetFieldSubType(dom OGRFieldDomain) (result OGRFieldSubType) {
	result = OGRFieldSubType(C.OGR_FldDomain_GetFieldSubType(dom.cValue))
	return
}

func ogrFldDomainGetSplitPolicy(dom OGRFieldDomain) (result OGRFieldDomainSplitPolicy) {
	result = OGRFieldDomainSplitPolicy(C.OGR_FldDomain_GetSplitPolicy(dom.cValue))
	return
}

func ogrFldDomainSetSplitPolicy(dom OGRFieldDomain, policy OGRFieldDomainSplitPolicy) {
	C.OGR_FldDomain_SetSplitPolicy(dom.cValue, C.OGRFieldDomainSplitPolicy(policy))
}

func ogrFldDomainGetMergePolicy(dom OGRFieldDomain) (result OGRFieldDomainMergePolicy) {
	result = OGRFieldDomainMergePolicy(C.OGR_FldDomain_GetMergePolicy(dom.cValue))
	return
}

func ogrFldDomainSetMergePolicy(dom OGRFieldDomain, policy OGRFieldDomainMergePolicy) {
	C.OGR_FldDomain_SetMergePolicy(dom.cValue, C.OGRFieldDomainMergePolicy(policy))
}

func ogrCodedFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, enumeration OGRCodedValue) (result OGRFieldDomain) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csDesc := C.CString(description)
	defer C.free(unsafe.Pointer(csDesc))
	result = OGRFieldDomain{cValue: C.OGR_CodedFldDomain_Create(csName, csDesc, C.OGRFieldType(eFieldType), C.OGRFieldSubType(eFieldSubType), enumeration.cValue)}
	return
}

func ogrCodedFldDomainGetEnumeration(dom OGRFieldDomain) (result OGRCodedValue) {
	result = OGRCodedValue{cValue: C.OGR_CodedFldDomain_GetEnumeration(dom.cValue)}
	return
}

func ogrRangeFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, min OGRField, minInclusive bool, max OGRField, maxInclusive bool) (result OGRFieldDomain) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csDesc := C.CString(description)
	defer C.free(unsafe.Pointer(csDesc))
	result = OGRFieldDomain{cValue: C.OGR_RangeFldDomain_Create(csName, csDesc, C.OGRFieldType(eFieldType), C.OGRFieldSubType(eFieldSubType), min.cValue, C.bool(minInclusive), max.cValue, C.bool(maxInclusive))}
	return
}

func ogrRangeFldDomainGetMin(dom OGRFieldDomain) (result OGRField, inclusive bool) {
	var inc C.bool
	result = OGRField{cValue: C.OGR_RangeFldDomain_GetMin(dom.cValue, &inc)}
	inclusive = bool(inc)
	return
}

func ogrRangeFldDomainGetMax(dom OGRFieldDomain) (result OGRField, inclusive bool) {
	var inc C.bool
	result = OGRField{cValue: C.OGR_RangeFldDomain_GetMax(dom.cValue, &inc)}
	inclusive = bool(inc)
	return
}

func ogrGlobFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, glob string) (result OGRFieldDomain) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csDesc := C.CString(description)
	defer C.free(unsafe.Pointer(csDesc))
	csGlob := C.CString(glob)
	defer C.free(unsafe.Pointer(csGlob))
	result = OGRFieldDomain{cValue: C.OGR_GlobFldDomain_Create(csName, csDesc, C.OGRFieldType(eFieldType), C.OGRFieldSubType(eFieldSubType), csGlob)}
	return
}

func ogrGlobFldDomainGetGlob(dom OGRFieldDomain) (result string) {
	result = C.GoString(C.OGR_GlobFldDomain_GetGlob(dom.cValue))
	return
}

// /* -------------------------------------------------------------------- */
// /*      ogrsf_frmts.h                                                   */
// /* -------------------------------------------------------------------- */

// /* OGRLayer */

func ogrLGetName(l OGRLayer) (result string) {
	result = C.GoString(C.OGR_L_GetName(l.cValue))
	return
}

func ogrLGetGeomType(l OGRLayer) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_L_GetGeomType(l.cValue))
	return
}

// /* Defined in gdal.h to avoid circular dependency with ogr_api.h */
// /* GDALDatasetH CPL_DLL OGR_L_GetDataset(OGRLayerH hLayer); */

// /** Result item of OGR_L_GetGeometryTypes */
type OGRGeometryTypeCounter struct {
	cValue *C.OGRGeometryTypeCounter
}

// /** Flag for OGR_L_GetGeometryTypes() indicating that
//   - OGRGeometryTypeCounter::nCount value is not needed */
const OGRGGTCountNotNeeded = C.OGR_GGT_COUNT_NOT_NEEDED

// /** Flag for OGR_L_GetGeometryTypes() indicating that iteration might stop as
//   - sooon as 2 distinct geometry types are found. */
const OGRGGTStopIfMixed = C.OGR_GGT_STOP_IF_MIXED

// /** Flag for OGR_L_GetGeometryTypes() indicating that a GeometryCollectionZ
//   - whose first subgeometry is a TinZ should be reported as TinZ */
const OGRGGTGeomCollectionZTinZ = C.OGR_GGT_GEOMCOLLECTIONZ_TINZ

func ogrLGetGeometryTypes(l OGRLayer, iGeomField, flags int, progress GDALProgressFunc, progressData unsafe.Pointer) (result []OGRGeometryTypeCounter) {
	var count C.int
	p := C.OGR_L_GetGeometryTypes(l.cValue, C.int(iGeomField), C.int(flags), &count, progress.cValue, progressData)
	if p == nil || count == 0 {
		return
	}
	slice := unsafe.Slice(p, int(count))
	result = make([]OGRGeometryTypeCounter, int(count))
	for i := range slice {
		c := new(C.OGRGeometryTypeCounter)
		*c = slice[i]
		result[i] = OGRGeometryTypeCounter{cValue: c}
	}
	vsiFree(unsafe.Pointer(p))
	return
}

func ogrLGetSpatialFilter(l OGRLayer) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_L_GetSpatialFilter(l.cValue)}
	return
}

func ogrLSetSpatialFilter(l OGRLayer, geom OGRGeometry) {
	C.OGR_L_SetSpatialFilter(l.cValue, geom.cValue)
}

func ogrLSetSpatialFilterRect(l OGRLayer, minX, minY, maxX, maxY float64) {
	C.OGR_L_SetSpatialFilterRect(l.cValue, C.double(minX), C.double(minY), C.double(maxX), C.double(maxY))
}

func ogrLSetSpatialFilterEx(l OGRLayer, iGeomField int, geom OGRGeometry) {
	C.OGR_L_SetSpatialFilterEx(l.cValue, C.int(iGeomField), geom.cValue)
}

func ogrLSetSpatialFilterRectEx(l OGRLayer, iGeomField int, minX, minY, maxX, maxY float64) {
	C.OGR_L_SetSpatialFilterRectEx(l.cValue, C.int(iGeomField), C.double(minX), C.double(minY), C.double(maxX), C.double(maxY))
}

func ogrLSetAttributeFilter(l OGRLayer, query string) (result OGRErr) {
	var cs *C.char
	if query != "" {
		cs = C.CString(query)
		defer C.free(unsafe.Pointer(cs))
	}
	result = OGRErr(C.OGR_L_SetAttributeFilter(l.cValue, cs))
	return
}

func ogrLResetReading(l OGRLayer) {
	C.OGR_L_ResetReading(l.cValue)
}

func ogrLGetNextFeature(l OGRLayer) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_L_GetNextFeature(l.cValue)}
	return
}

// OGR_FOR_EACH_FEATURE_BEGIN(hFeat, hLayer) / OGR_FOR_EACH_FEATURE_END(hFeat):
// convenience C macros to iterate features; not wrapped (use ResetReading + GetNextFeature).

// /** Data type for a Arrow C stream. Include ogr_recordbatch.h to get the definition. */
// struct ArrowArrayStream;
// bool CPL_DLL OGR_L_GetArrowStream(OGRLayerH hLayer, struct ArrowArrayStream *out_stream, char **papszOptions);

// /** Data type for a Arrow C schema. Include ogr_recordbatch.h to get the definition. */
// struct ArrowSchema;
// bool CPL_DLL OGR_L_IsArrowSchemaSupported(OGRLayerH hLayer, const struct ArrowSchema *schema, char **papszOptions, char **ppszErrorMsg);
// bool CPL_DLL OGR_L_CreateFieldFromArrowSchema(OGRLayerH hLayer, const struct ArrowSchema *schema, char **papszOptions);

// /** Data type for a Arrow C array. Include ogr_recordbatch.h to get the definition. */
// struct ArrowArray;
// bool CPL_DLL OGR_L_WriteArrowBatch(OGRLayerH hLayer, const struct ArrowSchema *schema, struct ArrowArray *array, char **papszOptions);

func ogrLSetNextByIndex(l OGRLayer, index int64) (result OGRErr) {
	result = OGRErr(C.OGR_L_SetNextByIndex(l.cValue, C.GIntBig(index)))
	return
}

func ogrLGetFeature(l OGRLayer, fid int64) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_L_GetFeature(l.cValue, C.GIntBig(fid))}
	return
}

func ogrLSetFeature(l OGRLayer, feat OGRFeature) (result OGRErr) {
	result = OGRErr(C.OGR_L_SetFeature(l.cValue, feat.cValue))
	return
}

func ogrLCreateFeature(l OGRLayer, feat OGRFeature) (result OGRErr) {
	result = OGRErr(C.OGR_L_CreateFeature(l.cValue, feat.cValue))
	return
}

func ogrLDeleteFeature(l OGRLayer, fid int64) (result OGRErr) {
	result = OGRErr(C.OGR_L_DeleteFeature(l.cValue, C.GIntBig(fid)))
	return
}

func ogrLUpsertFeature(l OGRLayer, feat OGRFeature) (result OGRErr) {
	result = OGRErr(C.OGR_L_UpsertFeature(l.cValue, feat.cValue))
	return
}

func ogrLUpdateFeature(l OGRLayer, feat OGRFeature, updatedFieldsIdx, updatedGeomFieldsIdx []int, updateStyleString bool) (result OGRErr) {
	cFields := make([]C.int, len(updatedFieldsIdx))
	for i, v := range updatedFieldsIdx {
		cFields[i] = C.int(v)
	}
	var pFields *C.int
	if len(cFields) > 0 {
		pFields = &cFields[0]
	}
	cGeom := make([]C.int, len(updatedGeomFieldsIdx))
	for i, v := range updatedGeomFieldsIdx {
		cGeom[i] = C.int(v)
	}
	var pGeom *C.int
	if len(cGeom) > 0 {
		pGeom = &cGeom[0]
	}
	result = OGRErr(C.OGR_L_UpdateFeature(l.cValue, feat.cValue, C.int(len(updatedFieldsIdx)), pFields, C.int(len(updatedGeomFieldsIdx)), pGeom, C.bool(updateStyleString)))
	return
}

func ogrLGetLayerDefn(l OGRLayer) (result OGRFeatureDefn) {
	result = OGRFeatureDefn{cValue: C.OGR_L_GetLayerDefn(l.cValue)}
	return
}

func ogrLGetSpatialRef(l OGRLayer) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OGR_L_GetSpatialRef(l.cValue)}
	return
}

func ogrLGetSupportedSRSList(l OGRLayer, iGeomField int) (result []OGRSpatialReference) {
	var count C.int
	p := C.OGR_L_GetSupportedSRSList(l.cValue, C.int(iGeomField), &count)
	if p == nil || count == 0 {
		return
	}
	slice := unsafe.Slice(p, int(count))
	result = make([]OGRSpatialReference, int(count))
	for i := range slice {
		result[i] = OGRSpatialReference{cValue: slice[i]}
	}
	vsiFree(unsafe.Pointer(p))
	return
}

func ogrLSetActiveSRS(l OGRLayer, iGeomField int, sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OGR_L_SetActiveSRS(l.cValue, C.int(iGeomField), sr.cValue))
	return
}

func ogrLFindFieldIndex(l OGRLayer, name string, exactMatch int) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_L_FindFieldIndex(l.cValue, cs, C.int(exactMatch)))
	return
}

func ogrLGetFeatureCount(l OGRLayer, force int) (result int64) {
	result = int64(C.OGR_L_GetFeatureCount(l.cValue, C.int(force)))
	return
}

func ogrLGetExtent(l OGRLayer, force int) (result OGREnvelope, status OGRErr) {
	result = OGREnvelope{cValue: new(C.OGREnvelope)}
	status = OGRErr(C.OGR_L_GetExtent(l.cValue, result.cValue, C.int(force)))
	return
}

func ogrLGetExtentEx(l OGRLayer, iGeomField, force int) (result OGREnvelope, status OGRErr) {
	result = OGREnvelope{cValue: new(C.OGREnvelope)}
	status = OGRErr(C.OGR_L_GetExtentEx(l.cValue, C.int(iGeomField), result.cValue, C.int(force)))
	return
}

func ogrLGetExtent3D(l OGRLayer, iGeomField, force int) (result OGREnvelope3D, status OGRErr) {
	result = OGREnvelope3D{cValue: new(C.OGREnvelope3D)}
	status = OGRErr(C.OGR_L_GetExtent3D(l.cValue, C.int(iGeomField), result.cValue, C.int(force)))
	return
}

func ogrLTestCapability(l OGRLayer, capability string) (result int) {
	cs := C.CString(capability)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_L_TestCapability(l.cValue, cs))
	return
}

func ogrLCreateField(l OGRLayer, fld OGRFieldDefn, approxOK int) (result OGRErr) {
	result = OGRErr(C.OGR_L_CreateField(l.cValue, fld.cValue, C.int(approxOK)))
	return
}

func ogrLCreateGeomField(l OGRLayer, gfld OGRGeomFieldDefn, force int) (result OGRErr) {
	result = OGRErr(C.OGR_L_CreateGeomField(l.cValue, gfld.cValue, C.int(force)))
	return
}

func ogrLDeleteField(l OGRLayer, field int) (result OGRErr) {
	result = OGRErr(C.OGR_L_DeleteField(l.cValue, C.int(field)))
	return
}

func ogrLReorderFields(l OGRLayer, panMap []int) (result OGRErr) {
	cMap := make([]C.int, len(panMap))
	for i, v := range panMap {
		cMap[i] = C.int(v)
	}
	var p *C.int
	if len(cMap) > 0 {
		p = &cMap[0]
	}
	result = OGRErr(C.OGR_L_ReorderFields(l.cValue, p))
	return
}

func ogrLReorderField(l OGRLayer, oldFieldPos, newFieldPos int) (result OGRErr) {
	result = OGRErr(C.OGR_L_ReorderField(l.cValue, C.int(oldFieldPos), C.int(newFieldPos)))
	return
}

func ogrLAlterFieldDefn(l OGRLayer, field int, newFieldDefn OGRFieldDefn, flags int) (result OGRErr) {
	result = OGRErr(C.OGR_L_AlterFieldDefn(l.cValue, C.int(field), newFieldDefn.cValue, C.int(flags)))
	return
}

func ogrLAlterGeomFieldDefn(l OGRLayer, field int, newGeomFieldDefn OGRGeomFieldDefn, flags int) (result OGRErr) {
	result = OGRErr(C.OGR_L_AlterGeomFieldDefn(l.cValue, C.int(field), newGeomFieldDefn.cValue, C.int(flags)))
	return
}

func ogrLStartTransaction(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_StartTransaction(l.cValue))
	return
}

func ogrLCommitTransaction(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_CommitTransaction(l.cValue))
	return
}

func ogrLRollbackTransaction(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_RollbackTransaction(l.cValue))
	return
}

func ogrLRename(l OGRLayer, newName string) (result OGRErr) {
	cs := C.CString(newName)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OGR_L_Rename(l.cValue, cs))
	return
}

// /*! @cond Doxygen_Suppress */

func ogrLReference(l OGRLayer) (result int) {
	result = int(C.OGR_L_Reference(l.cValue))
	return
}

func ogrLDereference(l OGRLayer) (result int) {
	result = int(C.OGR_L_Dereference(l.cValue))
	return
}

func ogrLGetRefCount(l OGRLayer) (result int) {
	result = int(C.OGR_L_GetRefCount(l.cValue))
	return
}

// /*! @endcond */

func ogrLSyncToDisk(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_SyncToDisk(l.cValue))
	return
}

// /*! @cond Doxygen_Suppress */

func ogrLGetFeaturesRead(l OGRLayer) (result int64) {
	result = int64(C.OGR_L_GetFeaturesRead(l.cValue))
	return
}

// /*! @endcond */

func ogrLGetFIDColumn(l OGRLayer) (result string) {
	result = C.GoString(C.OGR_L_GetFIDColumn(l.cValue))
	return
}

func ogrLGetGeometryColumn(l OGRLayer) (result string) {
	result = C.GoString(C.OGR_L_GetGeometryColumn(l.cValue))
	return
}

// /** Get style table */
func ogrLGetStyleTable(l OGRLayer) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_L_GetStyleTable(l.cValue)}
	return
}

// /** Set style table (and take ownership) */
func ogrLSetStyleTableDirectly(l OGRLayer, styleTable OGRStyleTable) {
	C.OGR_L_SetStyleTableDirectly(l.cValue, styleTable.cValue)
}

// /** Set style table */
func ogrLSetStyleTable(l OGRLayer, styleTable OGRStyleTable) {
	C.OGR_L_SetStyleTable(l.cValue, styleTable.cValue)
}

func ogrLSetIgnoredFields(l OGRLayer, fields CSLConstList) (result OGRErr) {
	result = OGRErr(C.OGR_L_SetIgnoredFields(l.cValue, fields.cValue))
	return
}

func ogrLIntersection(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := options.cValue
	status = OGRErr(C.OGR_L_Intersection(input.cValue, method.cValue, result.cValue, cOptions, progress.cValue, progressData))
	return
}

func ogrLUnion(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := options.cValue
	status = OGRErr(C.OGR_L_Union(input.cValue, method.cValue, result.cValue, cOptions, progress.cValue, progressData))
	return
}

func ogrLSymDifference(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := options.cValue
	status = OGRErr(C.OGR_L_SymDifference(input.cValue, method.cValue, result.cValue, cOptions, progress.cValue, progressData))
	return
}

func ogrLIdentity(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := options.cValue
	status = OGRErr(C.OGR_L_Identity(input.cValue, method.cValue, result.cValue, cOptions, progress.cValue, progressData))
	return
}

func ogrLUpdate(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := options.cValue
	status = OGRErr(C.OGR_L_Update(input.cValue, method.cValue, result.cValue, cOptions, progress.cValue, progressData))
	return
}

func ogrLClip(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := options.cValue
	status = OGRErr(C.OGR_L_Clip(input.cValue, method.cValue, result.cValue, cOptions, progress.cValue, progressData))
	return
}

func ogrLErase(input, method, result OGRLayer, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := options.cValue
	status = OGRErr(C.OGR_L_Erase(input.cValue, method.cValue, result.cValue, cOptions, progress.cValue, progressData))
	return
}

// /* OGRDataSource */

func ogrDSDestroy(ds OGRDataSource) {
	C.OGR_DS_Destroy(ds.cValue)
}

func ogrDSGetName(ds OGRDataSource) (result string) {
	result = C.GoString(C.OGR_DS_GetName(ds.cValue))
	return
}

func ogrDSGetLayerCount(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_GetLayerCount(ds.cValue))
	return
}

func ogrDSGetLayer(ds OGRDataSource, layer int) (result OGRLayer) {
	result = OGRLayer{cValue: C.OGR_DS_GetLayer(ds.cValue, C.int(layer))}
	return
}

func ogrDSGetLayerByName(ds OGRDataSource, name string) (result OGRLayer) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRLayer{cValue: C.OGR_DS_GetLayerByName(ds.cValue, cs)}
	return
}

func ogrDSDeleteLayer(ds OGRDataSource, layer int) (result OGRErr) {
	result = OGRErr(C.OGR_DS_DeleteLayer(ds.cValue, C.int(layer)))
	return
}

func ogrDSGetDriver(ds OGRDataSource) (result OGRSFDriver) {
	result = OGRSFDriver{cValue: C.OGR_DS_GetDriver(ds.cValue)}
	return
}

func ogrDSCreateLayer(ds OGRDataSource, name string, sr OGRSpatialReference, geomType OGRwkbGeometryType, options CSLConstList) (result OGRLayer) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	cOptions := options.cValue
	result = OGRLayer{cValue: C.OGR_DS_CreateLayer(ds.cValue, cs, sr.cValue, C.OGRwkbGeometryType(geomType), cOptions)}
	return
}

func ogrDSCopyLayer(ds OGRDataSource, srcLayer OGRLayer, newName string, options CSLConstList) (result OGRLayer) {
	cs := C.CString(newName)
	defer C.free(unsafe.Pointer(cs))
	cOptions := options.cValue
	result = OGRLayer{cValue: C.OGR_DS_CopyLayer(ds.cValue, srcLayer.cValue, cs, cOptions)}
	return
}

func ogrDSTestCapability(ds OGRDataSource, capability string) (result int) {
	cs := C.CString(capability)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_DS_TestCapability(ds.cValue, cs))
	return
}

func ogrDSExecuteSQL(ds OGRDataSource, statement string, spatialFilter OGRGeometry, dialect string) (result OGRLayer) {
	csStatement := C.CString(statement)
	defer C.free(unsafe.Pointer(csStatement))
	var csDialect *C.char
	if dialect != "" {
		csDialect = C.CString(dialect)
		defer C.free(unsafe.Pointer(csDialect))
	}
	result = OGRLayer{cValue: C.OGR_DS_ExecuteSQL(ds.cValue, csStatement, spatialFilter.cValue, csDialect)}
	return
}

func ogrDSReleaseResultSet(ds OGRDataSource, layer OGRLayer) {
	C.OGR_DS_ReleaseResultSet(ds.cValue, layer.cValue)
}

// /*! @cond Doxygen_Suppress */

func ogrDSReference(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_Reference(ds.cValue))
	return
}

func ogrDSDereference(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_Dereference(ds.cValue))
	return
}

func ogrDSGetRefCount(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_GetRefCount(ds.cValue))
	return
}

func ogrDSGetSummaryRefCount(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_GetSummaryRefCount(ds.cValue))
	return
}

// /*! @endcond */

// /** Flush pending changes to disk. See GDALDataset::FlushCache() */
func ogrDSSyncToDisk(ds OGRDataSource) (result OGRErr) {
	result = OGRErr(C.OGR_DS_SyncToDisk(ds.cValue))
	return
}

// /** Get style table */
func ogrDSGetStyleTable(ds OGRDataSource) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_DS_GetStyleTable(ds.cValue)}
	return
}

// /** Set style table (and take ownership) */
func ogrDSSetStyleTableDirectly(ds OGRDataSource, styleTable OGRStyleTable) {
	C.OGR_DS_SetStyleTableDirectly(ds.cValue, styleTable.cValue)
}

// /** Set style table */
func ogrDSSetStyleTable(ds OGRDataSource, styleTable OGRStyleTable) {
	C.OGR_DS_SetStyleTable(ds.cValue, styleTable.cValue)
}

// /* OGRSFDriver */

func ogrDrGetName(dr OGRSFDriver) (result string) {
	result = C.GoString(C.OGR_Dr_GetName(dr.cValue))
	return
}

func ogrDrOpen(dr OGRSFDriver, name string, update int) (result OGRDataSource) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRDataSource{cValue: C.OGR_Dr_Open(dr.cValue, cs, C.int(update))}
	return
}

func ogrDrTestCapability(dr OGRSFDriver, capability string) (result int) {
	cs := C.CString(capability)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_Dr_TestCapability(dr.cValue, cs))
	return
}

func ogrDrCreateDataSource(dr OGRSFDriver, name string, options CSLConstList) (result OGRDataSource) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	cOptions := options.cValue
	result = OGRDataSource{cValue: C.OGR_Dr_CreateDataSource(dr.cValue, cs, cOptions)}
	return
}

func ogrDrCopyDataSource(dr OGRSFDriver, srcDS OGRDataSource, newName string, options CSLConstList) (result OGRDataSource) {
	cs := C.CString(newName)
	defer C.free(unsafe.Pointer(cs))
	cOptions := options.cValue
	result = OGRDataSource{cValue: C.OGR_Dr_CopyDataSource(dr.cValue, srcDS.cValue, cs, cOptions)}
	return
}

func ogrDrDeleteDataSource(dr OGRSFDriver, name string) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OGR_Dr_DeleteDataSource(dr.cValue, cs))
	return
}

// /* OGRSFDriverRegistrar */

func ogrOpen(name string, update int) (result OGRDataSource, driver OGRSFDriver) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	var hDriver C.OGRSFDriverH
	result = OGRDataSource{cValue: C.OGROpen(cs, C.int(update), &hDriver)}
	driver = OGRSFDriver{cValue: hDriver}
	return
}

func ogrOpenShared(name string, update int) (result OGRDataSource, driver OGRSFDriver) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	var hDriver C.OGRSFDriverH
	result = OGRDataSource{cValue: C.OGROpenShared(cs, C.int(update), &hDriver)}
	driver = OGRSFDriver{cValue: hDriver}
	return
}

func ogrReleaseDataSource(ds OGRDataSource) (result OGRErr) {
	result = OGRErr(C.OGRReleaseDataSource(ds.cValue))
	return
}

// /*! @cond Doxygen_Suppress */

func ogrRegisterDriver(dr OGRSFDriver) {
	C.OGRRegisterDriver(dr.cValue)
}

func ogrDeregisterDriver(dr OGRSFDriver) {
	C.OGRDeregisterDriver(dr.cValue)
}

// /*! @endcond */

func ogrGetDriverCount() (result int) {
	result = int(C.OGRGetDriverCount())
	return
}

func ogrGetDriver(driver int) (result OGRSFDriver) {
	result = OGRSFDriver{cValue: C.OGRGetDriver(C.int(driver))}
	return
}

func ogrGetDriverByName(name string) (result OGRSFDriver) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRSFDriver{cValue: C.OGRGetDriverByName(cs)}
	return
}

// /*! @cond Doxygen_Suppress */

func ogrGetOpenDSCount() (result int) {
	result = int(C.OGRGetOpenDSCount())
	return
}

func ogrGetOpenDS(ds int) (result OGRDataSource) {
	result = OGRDataSource{cValue: C.OGRGetOpenDS(C.int(ds))}
	return
}

// /*! @endcond */

func ogrRegisterAll() {
	C.OGRRegisterAll()
}

// /** Clean-up all drivers, including raster ones.
//   - See GDALDestroyDriverManager() */
func ogrCleanupAll() {
	C.OGRCleanupAll()
}

// /* -------------------------------------------------------------------- */
// /*      ogrsf_featurestyle.h                                            */
// /* -------------------------------------------------------------------- */

// /* OGRStyleMgr */

func ogrSMCreate(styleTable OGRStyleTable) (result OGRStyleMgr) {
	result = OGRStyleMgr{cValue: C.OGR_SM_Create(styleTable.cValue)}
	return
}

func ogrSMDestroy(sm OGRStyleMgr) {
	C.OGR_SM_Destroy(sm.cValue)
}

func ogrSMInitFromFeature(sm OGRStyleMgr, feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_SM_InitFromFeature(sm.cValue, feat.cValue))
	return
}

func ogrSMInitStyleString(sm OGRStyleMgr, styleString string) (result int) {
	var cs *C.char
	if styleString != "" {
		cs = C.CString(styleString)
		defer C.free(unsafe.Pointer(cs))
	}
	result = int(C.OGR_SM_InitStyleString(sm.cValue, cs))
	return
}

func ogrSMGetPartCount(sm OGRStyleMgr, styleString string) (result int) {
	var cs *C.char
	if styleString != "" {
		cs = C.CString(styleString)
		defer C.free(unsafe.Pointer(cs))
	}
	result = int(C.OGR_SM_GetPartCount(sm.cValue, cs))
	return
}

func ogrSMGetPart(sm OGRStyleMgr, partId int, styleString string) (result OGRStyleTool) {
	var cs *C.char
	if styleString != "" {
		cs = C.CString(styleString)
		defer C.free(unsafe.Pointer(cs))
	}
	result = OGRStyleTool{cValue: C.OGR_SM_GetPart(sm.cValue, C.int(partId), cs)}
	return
}

func ogrSMAddPart(sm OGRStyleMgr, st OGRStyleTool) (result int) {
	result = int(C.OGR_SM_AddPart(sm.cValue, st.cValue))
	return
}

func ogrSMAddStyle(sm OGRStyleMgr, styleName, styleString string) (result int) {
	csName := C.CString(styleName)
	defer C.free(unsafe.Pointer(csName))
	csStyle := C.CString(styleString)
	defer C.free(unsafe.Pointer(csStyle))
	result = int(C.OGR_SM_AddStyle(sm.cValue, csName, csStyle))
	return
}

// /* OGRStyleTool */

func ogrSTCreate(classId OGRSTClassId) (result OGRStyleTool) {
	result = OGRStyleTool{cValue: C.OGR_ST_Create(C.OGRSTClassId(classId))}
	return
}

func ogrSTDestroy(st OGRStyleTool) {
	C.OGR_ST_Destroy(st.cValue)
}

func ogrSTGetType(st OGRStyleTool) (result OGRSTClassId) {
	result = OGRSTClassId(C.OGR_ST_GetType(st.cValue))
	return
}

func ogrSTGetUnit(st OGRStyleTool) (result OGRSTUnitId) {
	result = OGRSTUnitId(C.OGR_ST_GetUnit(st.cValue))
	return
}

func ogrSTSetUnit(st OGRStyleTool, unit OGRSTUnitId, groundPaperScale float64) {
	C.OGR_ST_SetUnit(st.cValue, C.OGRSTUnitId(unit), C.double(groundPaperScale))
}

func ogrSTGetParamStr(st OGRStyleTool, param int) (result string, isNull bool) {
	var vNull C.int
	result = C.GoString(C.OGR_ST_GetParamStr(st.cValue, C.int(param), &vNull))
	isNull = vNull != 0
	return
}

func ogrSTGetParamNum(st OGRStyleTool, param int) (result int, isNull bool) {
	var vNull C.int
	result = int(C.OGR_ST_GetParamNum(st.cValue, C.int(param), &vNull))
	isNull = vNull != 0
	return
}

func ogrSTGetParamDbl(st OGRStyleTool, param int) (result float64, isNull bool) {
	var vNull C.int
	result = float64(C.OGR_ST_GetParamDbl(st.cValue, C.int(param), &vNull))
	isNull = vNull != 0
	return
}

func ogrSTSetParamStr(st OGRStyleTool, param int, value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_ST_SetParamStr(st.cValue, C.int(param), cs)
}

func ogrSTSetParamNum(st OGRStyleTool, param, value int) {
	C.OGR_ST_SetParamNum(st.cValue, C.int(param), C.int(value))
}

func ogrSTSetParamDbl(st OGRStyleTool, param int, value float64) {
	C.OGR_ST_SetParamDbl(st.cValue, C.int(param), C.double(value))
}

func ogrSTGetStyleString(st OGRStyleTool) (result string) {
	result = C.GoString(C.OGR_ST_GetStyleString(st.cValue))
	return
}

func ogrSTGetRGBFromString(st OGRStyleTool, color string) (red, green, blue, alpha int, ok bool) {
	cs := C.CString(color)
	defer C.free(unsafe.Pointer(cs))
	var cr, cg, cb, ca C.int
	r := C.OGR_ST_GetRGBFromString(st.cValue, cs, &cr, &cg, &cb, &ca)
	red, green, blue, alpha = int(cr), int(cg), int(cb), int(ca)
	ok = r != 0
	return
}

// /* OGRStyleTable */

func ogrSTBLCreate() (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_STBL_Create()}
	return
}

func ogrSTBLDestroy(stbl OGRStyleTable) {
	C.OGR_STBL_Destroy(stbl.cValue)
}

func ogrSTBLAddStyle(stbl OGRStyleTable, name, styleString string) (result int) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csStyle := C.CString(styleString)
	defer C.free(unsafe.Pointer(csStyle))
	result = int(C.OGR_STBL_AddStyle(stbl.cValue, csName, csStyle))
	return
}

func ogrSTBLSaveStyleTable(stbl OGRStyleTable, filename string) (result int) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_STBL_SaveStyleTable(stbl.cValue, cs))
	return
}

func ogrSTBLLoadStyleTable(stbl OGRStyleTable, filename string) (result int) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_STBL_LoadStyleTable(stbl.cValue, cs))
	return
}

func ogrSTBLFind(stbl OGRStyleTable, name string) (result string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OGR_STBL_Find(stbl.cValue, cs))
	return
}

func ogrSTBLResetStyleStringReading(stbl OGRStyleTable) {
	C.OGR_STBL_ResetStyleStringReading(stbl.cValue)
}

func ogrSTBLGetNextStyle(stbl OGRStyleTable) (result string) {
	result = C.GoString(C.OGR_STBL_GetNextStyle(stbl.cValue))
	return
}

func ogrSTBLGetLastStyleName(stbl OGRStyleTable) (result string) {
	result = C.GoString(C.OGR_STBL_GetLastStyleName(stbl.cValue))
	return
}

// CPL_C_END
