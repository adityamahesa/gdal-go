package gdal

/*
#include "ogr_api.h"
#include "cpl_string.h" // TODO: implement cpl_string.go
#include "cpl_conv.h" // TODO: implement cpl_conv.go
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

func OGRGetGEOSVersion() (major, minor, patch int) {
	major, minor, patch = ogrGetGEOSVersion()
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

func OGRGeomCoordinatePrecisionCreate() (result OGRGeomCoordinatePrecision, err error) {
	result = ogrGeomCoordinatePrecisionCreate()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGeomCoordinatePrecisionDestroy(p OGRGeomCoordinatePrecision) {
	C.OGRGeomCoordinatePrecisionDestroy(p.cValue)
}

func (p OGRGeomCoordinatePrecision) Destroy() {
	ogrGeomCoordinatePrecisionDestroy(p)
}

func ogrGeomCoordinatePrecisionGetXYResolution(p OGRGeomCoordinatePrecision) (result float64) {
	result = float64(C.OGRGeomCoordinatePrecisionGetXYResolution(p.cValue))
	return
}

func (p OGRGeomCoordinatePrecision) GetXYResolution() (result float64) {
	result = ogrGeomCoordinatePrecisionGetXYResolution(p)
	return
}

func ogrGeomCoordinatePrecisionGetZResolution(p OGRGeomCoordinatePrecision) (result float64) {
	result = float64(C.OGRGeomCoordinatePrecisionGetZResolution(p.cValue))
	return
}

func (p OGRGeomCoordinatePrecision) GetZResolution() (result float64) {
	result = ogrGeomCoordinatePrecisionGetZResolution(p)
	return
}

func ogrGeomCoordinatePrecisionGetMResolution(p OGRGeomCoordinatePrecision) (result float64) {
	result = float64(C.OGRGeomCoordinatePrecisionGetMResolution(p.cValue))
	return
}

func (p OGRGeomCoordinatePrecision) GetMResolution() (result float64) {
	result = ogrGeomCoordinatePrecisionGetMResolution(p)
	return
}

func ogrGeomCoordinatePrecisionGetFormats(p OGRGeomCoordinatePrecision) (result []string) {
	raw := C.OGRGeomCoordinatePrecisionGetFormats(p.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	n := int(C.CSLCount(raw))
	for i := 0; i < n; i++ {
		result = append(result, C.GoString(C.CSLGetField(raw, C.int(i))))
	}
	return
}

func (p OGRGeomCoordinatePrecision) GetFormats() (result []string) {
	result = ogrGeomCoordinatePrecisionGetFormats(p)
	return
}

func ogrGeomCoordinatePrecisionGetFormatSpecificOptions(p OGRGeomCoordinatePrecision, formatName string) (result []string) {
	cs := C.CString(formatName)
	defer C.free(unsafe.Pointer(cs))
	raw := C.OGRGeomCoordinatePrecisionGetFormatSpecificOptions(p.cValue, cs)
	n := int(C.CSLCount(raw))
	for i := 0; i < n; i++ {
		result = append(result, C.GoString(C.CSLGetField(raw, C.int(i))))
	}
	return
}

func (p OGRGeomCoordinatePrecision) GetFormatSpecificOptions(formatName string) (result []string) {
	result = ogrGeomCoordinatePrecisionGetFormatSpecificOptions(p, formatName)
	return
}

func ogrGeomCoordinatePrecisionSet(p OGRGeomCoordinatePrecision, xyResolution, zResolution, mResolution float64) {
	C.OGRGeomCoordinatePrecisionSet(p.cValue, C.double(xyResolution), C.double(zResolution), C.double(mResolution))
}

func (p OGRGeomCoordinatePrecision) Set(xyResolution, zResolution, mResolution float64) {
	ogrGeomCoordinatePrecisionSet(p, xyResolution, zResolution, mResolution)
}

func ogrGeomCoordinatePrecisionSetFromMeter(p OGRGeomCoordinatePrecision, sr OGRSpatialReference, xyMeterResolution, zMeterResolution, mResolution float64) {
	C.OGRGeomCoordinatePrecisionSetFromMeter(p.cValue, sr.cValue, C.double(xyMeterResolution), C.double(zMeterResolution), C.double(mResolution))
}

func (p OGRGeomCoordinatePrecision) SetFromMeter(sr OGRSpatialReference, xyMeterResolution, zMeterResolution, mResolution float64) {
	ogrGeomCoordinatePrecisionSetFromMeter(p, sr, xyMeterResolution, zMeterResolution, mResolution)
}

func ogrGeomCoordinatePrecisionSetFormatSpecificOptions(p OGRGeomCoordinatePrecision, formatName string, options []string) {
	csName := C.CString(formatName)
	defer C.free(unsafe.Pointer(csName))
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	C.OGRGeomCoordinatePrecisionSetFormatSpecificOptions(p.cValue, csName, &cOptions[0])
}

func (p OGRGeomCoordinatePrecision) SetFormatSpecificOptions(formatName string, options []string) {
	ogrGeomCoordinatePrecisionSetFormatSpecificOptions(p, formatName, options)
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

func OGRGCreateFromWkb(data []byte, sr OGRSpatialReference) (result OGRGeometry, err error) {
	var status OGRErr
	result, status = ogrGCreateFromWkb(data, sr)
	err = ogrError(status)
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

func OGRGCreateFromWkbEx(data []byte, sr OGRSpatialReference) (result OGRGeometry, err error) {
	var status OGRErr
	result, status = ogrGCreateFromWkbEx(data, sr)
	err = ogrError(status)
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

func OGRGCreateFromWkt(wkt string, sr OGRSpatialReference) (result OGRGeometry, err error) {
	var status OGRErr
	result, status = ogrGCreateFromWkt(wkt, sr)
	err = ogrError(status)
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

func OGRGCreateFromFgf(data []byte, sr OGRSpatialReference) (result OGRGeometry, bytesConsumed int, err error) {
	var status OGRErr
	result, bytesConsumed, status = ogrGCreateFromFgf(data, sr)
	err = ogrError(status)
	return
}

func ogrGCreateFromEnvelope(minX, maxX, minY, maxY float64, sr OGRSpatialReference) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_CreateFromEnvelope(C.double(minX), C.double(maxX), C.double(minY), C.double(maxY), sr.cValue)}
	return
}

func OGRGCreateFromEnvelope(minX, maxX, minY, maxY float64, sr OGRSpatialReference) (result OGRGeometry, err error) {
	result = ogrGCreateFromEnvelope(minX, maxX, minY, maxY, sr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGDestroyGeometry(g OGRGeometry) {
	C.OGR_G_DestroyGeometry(g.cValue)
}

func (g OGRGeometry) Destroy() {
	ogrGDestroyGeometry(g)
}

func ogrGCreateGeometry(eType OGRwkbGeometryType) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_CreateGeometry(C.OGRwkbGeometryType(eType))}
	return
}

func OGRGCreateGeometry(eType OGRwkbGeometryType) (result OGRGeometry, err error) {
	result = ogrGCreateGeometry(eType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGApproximateArcAngles(centerX, centerY, z, primaryRadius, secondaryAxis, rotation, startAngle, endAngle, maxAngleStepSizeDegrees float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ApproximateArcAngles(C.double(centerX), C.double(centerY), C.double(z), C.double(primaryRadius), C.double(secondaryAxis), C.double(rotation), C.double(startAngle), C.double(endAngle), C.double(maxAngleStepSizeDegrees))}
	return
}

func OGRGApproximateArcAngles(centerX, centerY, z, primaryRadius, secondaryAxis, rotation, startAngle, endAngle, maxAngleStepSizeDegrees float64) (result OGRGeometry, err error) {
	result = ogrGApproximateArcAngles(centerX, centerY, z, primaryRadius, secondaryAxis, rotation, startAngle, endAngle, maxAngleStepSizeDegrees)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGForceToPolygon(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToPolygon(g.cValue)}
	return
}

func (g OGRGeometry) ForceToPolygon() (result OGRGeometry, err error) {
	result = ogrGForceToPolygon(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGForceToLineString(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToLineString(g.cValue)}
	return
}

func (g OGRGeometry) ForceToLineString() (result OGRGeometry, err error) {
	result = ogrGForceToLineString(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGForceToMultiPolygon(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToMultiPolygon(g.cValue)}
	return
}

func (g OGRGeometry) ForceToMultiPolygon() (result OGRGeometry, err error) {
	result = ogrGForceToMultiPolygon(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGForceToMultiPoint(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToMultiPoint(g.cValue)}
	return
}

func (g OGRGeometry) ForceToMultiPoint() (result OGRGeometry, err error) {
	result = ogrGForceToMultiPoint(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGForceToMultiLineString(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ForceToMultiLineString(g.cValue)}
	return
}

func (g OGRGeometry) ForceToMultiLineString() (result OGRGeometry, err error) {
	result = ogrGForceToMultiLineString(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGForceTo(g OGRGeometry, eTargetType OGRwkbGeometryType, options []string) (result OGRGeometry) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRGeometry{cValue: C.OGR_G_ForceTo(g.cValue, C.OGRwkbGeometryType(eTargetType), &cOptions[0])}
	return
}

func (g OGRGeometry) ForceTo(eTargetType OGRwkbGeometryType, options []string) (result OGRGeometry, err error) {
	result = ogrGForceTo(g, eTargetType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGRemoveLowerDimensionSubGeoms(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_RemoveLowerDimensionSubGeoms(g.cValue)}
	return
}

func (g OGRGeometry) RemoveLowerDimensionSubGeoms() (result OGRGeometry, err error) {
	result = ogrGRemoveLowerDimensionSubGeoms(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGGetDimension(g OGRGeometry) (result int) {
	result = int(C.OGR_G_GetDimension(g.cValue))
	return
}

func (g OGRGeometry) GetDimension() (result int) {
	result = ogrGGetDimension(g)
	return
}

func ogrGGetCoordinateDimension(g OGRGeometry) (result int) {
	result = int(C.OGR_G_GetCoordinateDimension(g.cValue))
	return
}

func (g OGRGeometry) GetCoordinateDimension() (result int) {
	result = ogrGGetCoordinateDimension(g)
	return
}

func ogrGCoordinateDimension(g OGRGeometry) (result int) {
	result = int(C.OGR_G_CoordinateDimension(g.cValue))
	return
}

func (g OGRGeometry) CoordinateDimension() (result int) {
	result = ogrGCoordinateDimension(g)
	return
}

func ogrGSetCoordinateDimension(g OGRGeometry, dimension int) {
	C.OGR_G_SetCoordinateDimension(g.cValue, C.int(dimension))
}

func (g OGRGeometry) SetCoordinateDimension(dimension int) {
	ogrGSetCoordinateDimension(g, dimension)
}

func ogrGIs3D(g OGRGeometry) (result int) {
	result = int(C.OGR_G_Is3D(g.cValue))
	return
}

func (g OGRGeometry) Is3D() (result bool) {
	result = ogrGIs3D(g) != 0
	return
}

func ogrGIsMeasured(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsMeasured(g.cValue))
	return
}

func (g OGRGeometry) IsMeasured() (result bool) {
	result = ogrGIsMeasured(g) != 0
	return
}

func ogrGSet3D(g OGRGeometry, is3D int) {
	C.OGR_G_Set3D(g.cValue, C.int(is3D))
}

func (g OGRGeometry) Set3D(is3D int) {
	ogrGSet3D(g, is3D)
}

func ogrGSetMeasured(g OGRGeometry, isMeasured int) {
	C.OGR_G_SetMeasured(g.cValue, C.int(isMeasured))
}

func (g OGRGeometry) SetMeasured(isMeasured int) {
	ogrGSetMeasured(g, isMeasured)
}

func ogrGClone(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Clone(g.cValue)}
	return
}

func (g OGRGeometry) Clone() (result OGRGeometry, err error) {
	result = ogrGClone(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGGetEnvelope(g OGRGeometry) (result OGREnvelope) {
	result = OGREnvelope{cValue: new(C.OGREnvelope)}
	C.OGR_G_GetEnvelope(g.cValue, result.cValue)
	return
}

func (g OGRGeometry) GetEnvelope() (result OGREnvelope) {
	result = ogrGGetEnvelope(g)
	return
}

func ogrGGetEnvelope3D(g OGRGeometry) (result OGREnvelope3D) {
	result = OGREnvelope3D{cValue: new(C.OGREnvelope3D)}
	C.OGR_G_GetEnvelope3D(g.cValue, result.cValue)
	return
}

func (g OGRGeometry) GetEnvelope3D() (result OGREnvelope3D) {
	result = ogrGGetEnvelope3D(g)
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

func (g OGRGeometry) ImportFromWkb(data []byte) (err error) {
	err = ogrError(ogrGImportFromWkb(g, data))
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

func (g OGRGeometry) ExportToWkb(order OGRwkbByteOrder) (result []byte, err error) {
	var status OGRErr
	result, status = ogrGExportToWkb(g, order)
	err = ogrError(status)
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

func (g OGRGeometry) ExportToIsoWkb(order OGRwkbByteOrder) (result []byte, err error) {
	var status OGRErr
	result, status = ogrGExportToIsoWkb(g, order)
	err = ogrError(status)
	return
}

func ogrwkbExportOptionsCreate() (result OGRwkbExportOptions) {
	result = OGRwkbExportOptions{cValue: C.OGRwkbExportOptionsCreate()}
	return
}

func OGRwkbExportOptionsCreate() (result OGRwkbExportOptions, err error) {
	result = ogrwkbExportOptionsCreate()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrwkbExportOptionsDestroy(o OGRwkbExportOptions) {
	C.OGRwkbExportOptionsDestroy(o.cValue)
}

func (o OGRwkbExportOptions) Destroy() {
	ogrwkbExportOptionsDestroy(o)
}

func ogrwkbExportOptionsSetByteOrder(o OGRwkbExportOptions, order OGRwkbByteOrder) {
	C.OGRwkbExportOptionsSetByteOrder(o.cValue, C.OGRwkbByteOrder(order))
}

func (o OGRwkbExportOptions) SetByteOrder(order OGRwkbByteOrder) {
	ogrwkbExportOptionsSetByteOrder(o, order)
}

func ogrwkbExportOptionsSetVariant(o OGRwkbExportOptions, variant OGRwkbVariant) {
	C.OGRwkbExportOptionsSetVariant(o.cValue, C.OGRwkbVariant(variant))
}

func (o OGRwkbExportOptions) SetVariant(variant OGRwkbVariant) {
	ogrwkbExportOptionsSetVariant(o, variant)
}

func ogrwkbExportOptionsSetPrecision(o OGRwkbExportOptions, precision OGRGeomCoordinatePrecision) {
	C.OGRwkbExportOptionsSetPrecision(o.cValue, precision.cValue)
}

func (o OGRwkbExportOptions) SetPrecision(precision OGRGeomCoordinatePrecision) {
	ogrwkbExportOptionsSetPrecision(o, precision)
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

func (g OGRGeometry) ExportToWkbEx(opts OGRwkbExportOptions) (result []byte, err error) {
	var status OGRErr
	result, status = ogrGExportToWkbEx(g, opts)
	err = ogrError(status)
	return
}

func ogrGWkbSize(g OGRGeometry) (result int) {
	result = int(C.OGR_G_WkbSize(g.cValue))
	return
}

func (g OGRGeometry) WkbSize() (result int) {
	result = ogrGWkbSize(g)
	return
}

func ogrGWkbSizeEx(g OGRGeometry) (result int) {
	result = int(C.OGR_G_WkbSizeEx(g.cValue))
	return
}

func (g OGRGeometry) WkbSizeEx() (result int) {
	result = ogrGWkbSizeEx(g)
	return
}

func ogrGImportFromWkt(g OGRGeometry, wkt string) (result OGRErr) {
	cs := C.CString(wkt)
	defer C.free(unsafe.Pointer(cs))
	pcs := cs
	result = OGRErr(C.OGR_G_ImportFromWkt(g.cValue, &pcs))
	return
}

func (g OGRGeometry) ImportFromWkt(wkt string) (err error) {
	err = ogrError(ogrGImportFromWkt(g, wkt))
	return
}

func ogrGExportToWkt(g OGRGeometry) (result string, status OGRErr) {
	var cs *C.char
	status = OGRErr(C.OGR_G_ExportToWkt(g.cValue, &cs))
	if cs != nil {
		result = C.GoString(cs)
		C.CPLFree(unsafe.Pointer(cs))
	}
	return
}

func (g OGRGeometry) ExportToWkt() (result string, err error) {
	var status OGRErr
	result, status = ogrGExportToWkt(g)
	err = ogrError(status)
	return
}

func ogrGExportToIsoWkt(g OGRGeometry) (result string, status OGRErr) {
	var cs *C.char
	status = OGRErr(C.OGR_G_ExportToIsoWkt(g.cValue, &cs))
	if cs != nil {
		result = C.GoString(cs)
		C.CPLFree(unsafe.Pointer(cs))
	}
	return
}

func (g OGRGeometry) ExportToIsoWkt() (result string, err error) {
	var status OGRErr
	result, status = ogrGExportToIsoWkt(g)
	err = ogrError(status)
	return
}

func ogrGGetGeometryType(g OGRGeometry) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_G_GetGeometryType(g.cValue))
	return
}

func (g OGRGeometry) GetGeometryType() (result OGRwkbGeometryType) {
	result = ogrGGetGeometryType(g)
	return
}

func ogrGGetGeometryName(g OGRGeometry) (result string) {
	result = C.GoString(C.OGR_G_GetGeometryName(g.cValue))
	return
}

func (g OGRGeometry) GetGeometryName() (result string) {
	result = ogrGGetGeometryName(g)
	return
}

// void CPL_DLL OGR_G_DumpReadable(OGRGeometryH, FILE *, const char *);

func ogrGFlattenTo2D(g OGRGeometry) {
	C.OGR_G_FlattenTo2D(g.cValue)
}

func (g OGRGeometry) FlattenTo2D() {
	ogrGFlattenTo2D(g)
}

func ogrGCloseRings(g OGRGeometry) {
	C.OGR_G_CloseRings(g.cValue)
}

func (g OGRGeometry) CloseRings() {
	ogrGCloseRings(g)
}

func ogrGCreateFromGML(gml string) (result OGRGeometry) {
	cs := C.CString(gml)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeometry{cValue: C.OGR_G_CreateFromGML(cs)}
	return
}

func OGRGCreateFromGML(gml string) (result OGRGeometry, err error) {
	result = ogrGCreateFromGML(gml)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGExportToGML(g OGRGeometry) (result string) {
	cs := C.OGR_G_ExportToGML(g.cValue)
	if cs != nil {
		result = C.GoString(cs)
		C.CPLFree(unsafe.Pointer(cs))
	}
	return
}

func (g OGRGeometry) ExportToGML() (result string) {
	result = ogrGExportToGML(g)
	return
}

func ogrGExportToGMLEx(g OGRGeometry, options []string) (result string) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	cs := C.OGR_G_ExportToGMLEx(g.cValue, &cOptions[0])
	if cs != nil {
		result = C.GoString(cs)
		C.CPLFree(unsafe.Pointer(cs))
	}
	return
}

func (g OGRGeometry) ExportToGMLEx(options []string) (result string) {
	result = ogrGExportToGMLEx(g, options)
	return
}

func ogrGCreateFromGMLTree(tree CPLXMLNode) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_CreateFromGMLTree(tree.cValue)}
	return
}

func OGRGCreateFromGMLTree(tree CPLXMLNode) (result OGRGeometry, err error) {
	result = ogrGCreateFromGMLTree(tree)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGExportToGMLTree(g OGRGeometry) (result CPLXMLNode) {
	result = CPLXMLNode{cValue: C.OGR_G_ExportToGMLTree(g.cValue)}
	return
}

func (g OGRGeometry) ExportToGMLTree() (result CPLXMLNode, err error) {
	result = ogrGExportToGMLTree(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGExportEnvelopeToGMLTree(g OGRGeometry) (result CPLXMLNode) {
	result = CPLXMLNode{cValue: C.OGR_G_ExportEnvelopeToGMLTree(g.cValue)}
	return
}

func (g OGRGeometry) ExportEnvelopeToGMLTree() (result CPLXMLNode, err error) {
	result = ogrGExportEnvelopeToGMLTree(g)
	if result.cValue == nil {
		err = lastError()
	}
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
		C.CPLFree(unsafe.Pointer(raw))
	}
	return
}

func (g OGRGeometry) ExportToKML(altitudeMode string) (result string) {
	result = ogrGExportToKML(g, altitudeMode)
	return
}

func ogrGExportToJson(g OGRGeometry) (result string) {
	cs := C.OGR_G_ExportToJson(g.cValue)
	if cs != nil {
		result = C.GoString(cs)
		C.CPLFree(unsafe.Pointer(cs))
	}
	return
}

func (g OGRGeometry) ExportToJson() (result string) {
	result = ogrGExportToJson(g)
	return
}

func ogrGExportToJsonEx(g OGRGeometry, options []string) (result string) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	cs := C.OGR_G_ExportToJsonEx(g.cValue, &cOptions[0])
	if cs != nil {
		result = C.GoString(cs)
		C.CPLFree(unsafe.Pointer(cs))
	}
	return
}

func (g OGRGeometry) ExportToJsonEx(options []string) (result string) {
	result = ogrGExportToJsonEx(g, options)
	return
}

// /** Create a OGR geometry from a GeoJSON geometry object */
func ogrGCreateGeometryFromJson(json string) (result OGRGeometry) {
	cs := C.CString(json)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeometry{cValue: C.OGR_G_CreateGeometryFromJson(cs)}
	return
}

// /** Create a OGR geometry from a GeoJSON geometry object */
func OGRGCreateGeometryFromJson(json string) (result OGRGeometry, err error) {
	result = ogrGCreateGeometryFromJson(json)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// /** Create a OGR geometry from a ESRI JSON geometry object */
func ogrGCreateGeometryFromEsriJson(json string) (result OGRGeometry) {
	cs := C.CString(json)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeometry{cValue: C.OGR_G_CreateGeometryFromEsriJson(cs)}
	return
}

// /** Create a OGR geometry from a ESRI JSON geometry object */
func OGRGCreateGeometryFromEsriJson(json string) (result OGRGeometry, err error) {
	result = ogrGCreateGeometryFromEsriJson(json)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGAssignSpatialReference(g OGRGeometry, sr OGRSpatialReference) {
	C.OGR_G_AssignSpatialReference(g.cValue, sr.cValue)
}

func (g OGRGeometry) AssignSpatialReference(sr OGRSpatialReference) {
	ogrGAssignSpatialReference(g, sr)
}

func ogrGGetSpatialReference(g OGRGeometry) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OGR_G_GetSpatialReference(g.cValue)}
	return
}

func (g OGRGeometry) GetSpatialReference() (result OGRSpatialReference) {
	result = ogrGGetSpatialReference(g)
	return
}

func ogrGTransform(g OGRGeometry, ct OGRCoordinateTransformation) (result OGRErr) {
	result = OGRErr(C.OGR_G_Transform(g.cValue, ct.cValue))
	return
}

func (g OGRGeometry) Transform(ct OGRCoordinateTransformation) (err error) {
	err = ogrError(ogrGTransform(g, ct))
	return
}

func ogrGTransformTo(g OGRGeometry, sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OGR_G_TransformTo(g.cValue, sr.cValue))
	return
}

func (g OGRGeometry) TransformTo(sr OGRSpatialReference) (err error) {
	err = ogrError(ogrGTransformTo(g, sr))
	return
}

func ogrGeomTransformerCreate(ct OGRCoordinateTransformation, options []string) (result OGRGeomTransformer) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRGeomTransformer{cValue: C.OGR_GeomTransformer_Create(ct.cValue, &cOptions[0])}
	return
}

func OGRGeomTransformerCreate(ct OGRCoordinateTransformation, options []string) (result OGRGeomTransformer, err error) {
	result = ogrGeomTransformerCreate(ct, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGeomTransformerTransform(t OGRGeomTransformer, g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_GeomTransformer_Transform(t.cValue, g.cValue)}
	return
}

func (t OGRGeomTransformer) Transform(g OGRGeometry) (result OGRGeometry, err error) {
	result = ogrGeomTransformerTransform(t, g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGeomTransformerDestroy(t OGRGeomTransformer) {
	C.OGR_GeomTransformer_Destroy(t.cValue)
}

func (t OGRGeomTransformer) Destroy() {
	ogrGeomTransformerDestroy(t)
}

func ogrGSimplify(g OGRGeometry, tolerance float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Simplify(g.cValue, C.double(tolerance))}
	return
}

func (g OGRGeometry) Simplify(tolerance float64) (result OGRGeometry, err error) {
	result = ogrGSimplify(g, tolerance)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGSimplifyPreserveTopology(g OGRGeometry, tolerance float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_SimplifyPreserveTopology(g.cValue, C.double(tolerance))}
	return
}

func (g OGRGeometry) SimplifyPreserveTopology(tolerance float64) (result OGRGeometry, err error) {
	result = ogrGSimplifyPreserveTopology(g, tolerance)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGDelaunayTriangulation(g OGRGeometry, tolerance float64, onlyEdges int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_DelaunayTriangulation(g.cValue, C.double(tolerance), C.int(onlyEdges))}
	return
}

func (g OGRGeometry) DelaunayTriangulation(tolerance float64, onlyEdges int) (result OGRGeometry, err error) {
	result = ogrGDelaunayTriangulation(g, tolerance, onlyEdges)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGConstrainedDelaunayTriangulation(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ConstrainedDelaunayTriangulation(g.cValue)}
	return
}

func (g OGRGeometry) ConstrainedDelaunayTriangulation() (result OGRGeometry, err error) {
	result = ogrGConstrainedDelaunayTriangulation(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGSegmentize(g OGRGeometry, maxLength float64) {
	C.OGR_G_Segmentize(g.cValue, C.double(maxLength))
}

func (g OGRGeometry) Segmentize(maxLength float64) {
	ogrGSegmentize(g, maxLength)
}

func ogrGIntersects(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Intersects(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Intersects(other OGRGeometry) (result bool) {
	result = ogrGIntersects(g, other) != 0
	return
}

func ogrGEquals(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Equals(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Equals(other OGRGeometry) (result bool) {
	result = ogrGEquals(g, other) != 0
	return
}

func ogrGDisjoint(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Disjoint(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Disjoint(other OGRGeometry) (result bool) {
	result = ogrGDisjoint(g, other) != 0
	return
}

func ogrGTouches(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Touches(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Touches(other OGRGeometry) (result bool) {
	result = ogrGTouches(g, other) != 0
	return
}

func ogrGCrosses(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Crosses(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Crosses(other OGRGeometry) (result bool) {
	result = ogrGCrosses(g, other) != 0
	return
}

func ogrGWithin(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Within(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Within(other OGRGeometry) (result bool) {
	result = ogrGWithin(g, other) != 0
	return
}

func ogrGContains(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Contains(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Contains(other OGRGeometry) (result bool) {
	result = ogrGContains(g, other) != 0
	return
}

func ogrGOverlaps(g, other OGRGeometry) (result int) {
	result = int(C.OGR_G_Overlaps(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Overlaps(other OGRGeometry) (result bool) {
	result = ogrGOverlaps(g, other) != 0
	return
}

func ogrGBoundary(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Boundary(g.cValue)}
	return
}

func (g OGRGeometry) Boundary() (result OGRGeometry, err error) {
	result = ogrGBoundary(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGConvexHull(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ConvexHull(g.cValue)}
	return
}

func (g OGRGeometry) ConvexHull() (result OGRGeometry, err error) {
	result = ogrGConvexHull(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGConcaveHull(g OGRGeometry, ratio float64, allowHoles bool) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_ConcaveHull(g.cValue, C.double(ratio), C.bool(allowHoles))}
	return
}

func (g OGRGeometry) ConcaveHull(ratio float64, allowHoles bool) (result OGRGeometry, err error) {
	result = ogrGConcaveHull(g, ratio, allowHoles)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGBuffer(g OGRGeometry, dist float64, quadSegs int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Buffer(g.cValue, C.double(dist), C.int(quadSegs))}
	return
}

func (g OGRGeometry) Buffer(dist float64, quadSegs int) (result OGRGeometry, err error) {
	result = ogrGBuffer(g, dist, quadSegs)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGBufferEx(g OGRGeometry, dist float64, options []string) (result OGRGeometry) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRGeometry{cValue: C.OGR_G_BufferEx(g.cValue, C.double(dist), &cOptions[0])}
	return
}

func (g OGRGeometry) BufferEx(dist float64, options []string) (result OGRGeometry, err error) {
	result = ogrGBufferEx(g, dist, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGIntersection(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Intersection(g.cValue, other.cValue)}
	return
}

func (g OGRGeometry) Intersection(other OGRGeometry) (result OGRGeometry, err error) {
	result = ogrGIntersection(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGUnion(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Union(g.cValue, other.cValue)}
	return
}

func (g OGRGeometry) Union(other OGRGeometry) (result OGRGeometry, err error) {
	result = ogrGUnion(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGUnionCascaded(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_UnionCascaded(g.cValue)}
	return
}

func (g OGRGeometry) UnionCascaded() (result OGRGeometry, err error) {
	result = ogrGUnionCascaded(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGUnaryUnion(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_UnaryUnion(g.cValue)}
	return
}

func (g OGRGeometry) UnaryUnion() (result OGRGeometry, err error) {
	result = ogrGUnaryUnion(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGPointOnSurface(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_PointOnSurface(g.cValue)}
	return
}

func (g OGRGeometry) PointOnSurface() (result OGRGeometry, err error) {
	result = ogrGPointOnSurface(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGDifference(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Difference(g.cValue, other.cValue)}
	return
}

func (g OGRGeometry) Difference(other OGRGeometry) (result OGRGeometry, err error) {
	result = ogrGDifference(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGSymDifference(g, other OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_SymDifference(g.cValue, other.cValue)}
	return
}

func (g OGRGeometry) SymDifference(other OGRGeometry) (result OGRGeometry, err error) {
	result = ogrGSymDifference(g, other)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGDistance(g, other OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Distance(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Distance(other OGRGeometry) (result float64) {
	result = ogrGDistance(g, other)
	return
}

func ogrGDistance3D(g, other OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Distance3D(g.cValue, other.cValue))
	return
}

func (g OGRGeometry) Distance3D(other OGRGeometry) (result float64) {
	result = ogrGDistance3D(g, other)
	return
}

func ogrGLength(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Length(g.cValue))
	return
}

func (g OGRGeometry) Length() (result float64) {
	result = ogrGLength(g)
	return
}

func ogrGGeodesicLength(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_GeodesicLength(g.cValue))
	return
}

func (g OGRGeometry) GeodesicLength() (result float64) {
	result = ogrGGeodesicLength(g)
	return
}

func ogrGArea(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_Area(g.cValue))
	return
}

func (g OGRGeometry) Area() (result float64) {
	result = ogrGArea(g)
	return
}

func ogrGGeodesicArea(g OGRGeometry) (result float64) {
	result = float64(C.OGR_G_GeodesicArea(g.cValue))
	return
}

func (g OGRGeometry) GeodesicArea() (result float64) {
	result = ogrGGeodesicArea(g)
	return
}

func ogrGIsClockwise(g OGRGeometry) (result bool) {
	result = bool(C.OGR_G_IsClockwise(g.cValue))
	return
}

func (g OGRGeometry) IsClockwise() (result bool) {
	result = ogrGIsClockwise(g)
	return
}

func ogrGCentroid(g, centroid OGRGeometry) (result int) {
	result = int(C.OGR_G_Centroid(g.cValue, centroid.cValue))
	return
}

func (g OGRGeometry) Centroid() (result OGRGeometry, err error) {
	result = ogrGCreateGeometry(WkbPoint)
	err = ogrError(OGRErr(ogrGCentroid(g, result)))
	return
}

func ogrGValue(g OGRGeometry, distance float64) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Value(g.cValue, C.double(distance))}
	return
}

func (g OGRGeometry) Value(distance float64) (result OGRGeometry, err error) {
	result = ogrGValue(g, distance)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGEmpty(g OGRGeometry) {
	C.OGR_G_Empty(g.cValue)
}

func (g OGRGeometry) Empty() {
	ogrGEmpty(g)
}

func ogrGIsEmpty(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsEmpty(g.cValue))
	return
}

func (g OGRGeometry) IsEmpty() (result bool) {
	result = ogrGIsEmpty(g) != 0
	return
}

func ogrGIsValid(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsValid(g.cValue))
	return
}

func (g OGRGeometry) IsValid() (result bool) {
	result = ogrGIsValid(g) != 0
	return
}

// /*char    CPL_DLL *OGR_G_IsValidReason( OGRGeometryH );*/

func ogrGMakeValid(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_MakeValid(g.cValue)}
	return
}

func (g OGRGeometry) MakeValid() (result OGRGeometry, err error) {
	result = ogrGMakeValid(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGMakeValidEx(g OGRGeometry, options []string) (result OGRGeometry) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRGeometry{cValue: C.OGR_G_MakeValidEx(g.cValue, &cOptions[0])}
	return
}

func (g OGRGeometry) MakeValidEx(options []string) (result OGRGeometry, err error) {
	result = ogrGMakeValidEx(g, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGNormalize(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Normalize(g.cValue)}
	return
}

func (g OGRGeometry) Normalize() (result OGRGeometry, err error) {
	result = ogrGNormalize(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGIsSimple(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsSimple(g.cValue))
	return
}

func (g OGRGeometry) IsSimple() (result bool) {
	result = ogrGIsSimple(g) != 0
	return
}

func ogrGIsRing(g OGRGeometry) (result int) {
	result = int(C.OGR_G_IsRing(g.cValue))
	return
}

func (g OGRGeometry) IsRing() (result bool) {
	result = ogrGIsRing(g) != 0
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

func (g OGRGeometry) SetPrecision(gridSize float64, flags int) (result OGRGeometry, err error) {
	result = ogrGSetPrecision(g, gridSize, flags)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGPolygonize(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_Polygonize(g.cValue)}
	return
}

func (g OGRGeometry) Polygonize() (result OGRGeometry, err error) {
	result = ogrGPolygonize(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGBuildArea(g OGRGeometry) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_BuildArea(g.cValue)}
	return
}

func (g OGRGeometry) BuildArea() (result OGRGeometry, err error) {
	result = ogrGBuildArea(g)
	if result.cValue == nil {
		err = lastError()
	}
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

func (g OGRGeometry) GetPointCount() (result int) {
	result = ogrGGetPointCount(g)
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

func (g OGRGeometry) GetPoints() (x, y, z []float64) {
	x, y, z = ogrGGetPoints(g)
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

func (g OGRGeometry) GetPointsZM() (x, y, z, m []float64) {
	x, y, z, m = ogrGGetPointsZM(g)
	return
}

func ogrGGetX(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetX(g.cValue, C.int(point)))
	return
}

func (g OGRGeometry) GetX(point int) (result float64) {
	result = ogrGGetX(g, point)
	return
}

func ogrGGetY(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetY(g.cValue, C.int(point)))
	return
}

func (g OGRGeometry) GetY(point int) (result float64) {
	result = ogrGGetY(g, point)
	return
}

func ogrGGetZ(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetZ(g.cValue, C.int(point)))
	return
}

func (g OGRGeometry) GetZ(point int) (result float64) {
	result = ogrGGetZ(g, point)
	return
}

func ogrGGetM(g OGRGeometry, point int) (result float64) {
	result = float64(C.OGR_G_GetM(g.cValue, C.int(point)))
	return
}

func (g OGRGeometry) GetM(point int) (result float64) {
	result = ogrGGetM(g, point)
	return
}

func ogrGGetPoint(g OGRGeometry, point int) (x, y, z float64) {
	var cx, cy, cz C.double
	C.OGR_G_GetPoint(g.cValue, C.int(point), &cx, &cy, &cz)
	x, y, z = float64(cx), float64(cy), float64(cz)
	return
}

func (g OGRGeometry) GetPoint(point int) (x, y, z float64) {
	x, y, z = ogrGGetPoint(g, point)
	return
}

func ogrGGetPointZM(g OGRGeometry, point int) (x, y, z, m float64) {
	var cx, cy, cz, cm C.double
	C.OGR_G_GetPointZM(g.cValue, C.int(point), &cx, &cy, &cz, &cm)
	x, y, z, m = float64(cx), float64(cy), float64(cz), float64(cm)
	return
}

func (g OGRGeometry) GetPointZM(point int) (x, y, z, m float64) {
	x, y, z, m = ogrGGetPointZM(g, point)
	return
}

func ogrGSetPointCount(g OGRGeometry, count int) {
	C.OGR_G_SetPointCount(g.cValue, C.int(count))
}

func (g OGRGeometry) SetPointCount(count int) {
	ogrGSetPointCount(g, count)
}

func ogrGSetPoint(g OGRGeometry, point int, x, y, z float64) {
	C.OGR_G_SetPoint(g.cValue, C.int(point), C.double(x), C.double(y), C.double(z))
}

func (g OGRGeometry) SetPoint(point int, x, y, z float64) {
	ogrGSetPoint(g, point, x, y, z)
}

func ogrGSetPoint2D(g OGRGeometry, point int, x, y float64) {
	C.OGR_G_SetPoint_2D(g.cValue, C.int(point), C.double(x), C.double(y))
}

func (g OGRGeometry) SetPoint2D(point int, x, y float64) {
	ogrGSetPoint2D(g, point, x, y)
}

func ogrGSetPointM(g OGRGeometry, point int, x, y, m float64) {
	C.OGR_G_SetPointM(g.cValue, C.int(point), C.double(x), C.double(y), C.double(m))
}

func (g OGRGeometry) SetPointM(point int, x, y, m float64) {
	ogrGSetPointM(g, point, x, y, m)
}

func ogrGSetPointZM(g OGRGeometry, point int, x, y, z, m float64) {
	C.OGR_G_SetPointZM(g.cValue, C.int(point), C.double(x), C.double(y), C.double(z), C.double(m))
}

func (g OGRGeometry) SetPointZM(point int, x, y, z, m float64) {
	ogrGSetPointZM(g, point, x, y, z, m)
}

func ogrGAddPoint(g OGRGeometry, x, y, z float64) {
	C.OGR_G_AddPoint(g.cValue, C.double(x), C.double(y), C.double(z))
}

func (g OGRGeometry) AddPoint(x, y, z float64) {
	ogrGAddPoint(g, x, y, z)
}

func ogrGAddPoint2D(g OGRGeometry, x, y float64) {
	C.OGR_G_AddPoint_2D(g.cValue, C.double(x), C.double(y))
}

func (g OGRGeometry) AddPoint2D(x, y float64) {
	ogrGAddPoint2D(g, x, y)
}

func ogrGAddPointM(g OGRGeometry, x, y, m float64) {
	C.OGR_G_AddPointM(g.cValue, C.double(x), C.double(y), C.double(m))
}

func (g OGRGeometry) AddPointM(x, y, m float64) {
	ogrGAddPointM(g, x, y, m)
}

func ogrGAddPointZM(g OGRGeometry, x, y, z, m float64) {
	C.OGR_G_AddPointZM(g.cValue, C.double(x), C.double(y), C.double(z), C.double(m))
}

func (g OGRGeometry) AddPointZM(x, y, z, m float64) {
	ogrGAddPointZM(g, x, y, z, m)
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

func (g OGRGeometry) SetPoints(x, y, z []float64) {
	ogrGSetPoints(g, x, y, z)
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

func (g OGRGeometry) SetPointsZM(x, y, z, m []float64) {
	ogrGSetPointsZM(g, x, y, z, m)
}

func ogrGSwapXY(g OGRGeometry) {
	C.OGR_G_SwapXY(g.cValue)
}

func (g OGRGeometry) SwapXY() {
	ogrGSwapXY(g)
}

// /* Methods for getting/setting rings and members collections */

func ogrGGetGeometryCount(g OGRGeometry) (result int) {
	result = int(C.OGR_G_GetGeometryCount(g.cValue))
	return
}

func (g OGRGeometry) GetGeometryCount() (result int) {
	result = ogrGGetGeometryCount(g)
	return
}

func ogrGGetGeometryRef(g OGRGeometry, subGeom int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_G_GetGeometryRef(g.cValue, C.int(subGeom))}
	return
}

func (g OGRGeometry) GetGeometryRef(subGeom int) (result OGRGeometry) {
	result = ogrGGetGeometryRef(g, subGeom)
	return
}

func ogrGAddGeometry(g, subGeom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_G_AddGeometry(g.cValue, subGeom.cValue))
	return
}

func (g OGRGeometry) AddGeometry(subGeom OGRGeometry) (err error) {
	err = ogrError(ogrGAddGeometry(g, subGeom))
	return
}

func ogrGAddGeometryDirectly(g, subGeom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_G_AddGeometryDirectly(g.cValue, subGeom.cValue))
	return
}

func (g OGRGeometry) AddGeometryDirectly(subGeom OGRGeometry) (err error) {
	err = ogrError(ogrGAddGeometryDirectly(g, subGeom))
	return
}

func ogrGRemoveGeometry(g OGRGeometry, subGeom, delete int) (result OGRErr) {
	result = OGRErr(C.OGR_G_RemoveGeometry(g.cValue, C.int(subGeom), C.int(delete)))
	return
}

func (g OGRGeometry) RemoveGeometry(subGeom, delete int) (err error) {
	err = ogrError(ogrGRemoveGeometry(g, subGeom, delete))
	return
}

func ogrGHasCurveGeometry(g OGRGeometry, lookForNonLinear int) (result int) {
	result = int(C.OGR_G_HasCurveGeometry(g.cValue, C.int(lookForNonLinear)))
	return
}

func (g OGRGeometry) HasCurveGeometry(lookForNonLinear int) (result bool) {
	result = ogrGHasCurveGeometry(g, lookForNonLinear) != 0
	return
}

func ogrGGetLinearGeometry(g OGRGeometry, maxAngleStepSizeDegrees float64, options []string) (result OGRGeometry) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRGeometry{cValue: C.OGR_G_GetLinearGeometry(g.cValue, C.double(maxAngleStepSizeDegrees), &cOptions[0])}
	return
}

func (g OGRGeometry) GetLinearGeometry(maxAngleStepSizeDegrees float64, options []string) (result OGRGeometry, err error) {
	result = ogrGGetLinearGeometry(g, maxAngleStepSizeDegrees, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGGetCurveGeometry(g OGRGeometry, options []string) (result OGRGeometry) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRGeometry{cValue: C.OGR_G_GetCurveGeometry(g.cValue, &cOptions[0])}
	return
}

func (g OGRGeometry) GetCurveGeometry(options []string) (result OGRGeometry, err error) {
	result = ogrGGetCurveGeometry(g, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrBuildPolygonFromEdges(lines OGRGeometry, bestEffort, autoClose int, tolerance float64) (result OGRGeometry, status OGRErr) {
	var e C.OGRErr
	result = OGRGeometry{cValue: C.OGRBuildPolygonFromEdges(lines.cValue, C.int(bestEffort), C.int(autoClose), C.double(tolerance), &e)}
	status = OGRErr(e)
	return
}

func OGRBuildPolygonFromEdges(lines OGRGeometry, bestEffort, autoClose int, tolerance float64) (result OGRGeometry, err error) {
	var status OGRErr
	result, status = ogrBuildPolygonFromEdges(lines, bestEffort, autoClose, tolerance)
	err = ogrError(status)
	return
}

// /*! @cond Doxygen_Suppress */

func ogrSetGenerateDB2V72ByteOrder(generate int) (result OGRErr) {
	result = OGRErr(C.OGRSetGenerate_DB2_V72_BYTE_ORDER(C.int(generate)))
	return
}

func OGRSetGenerateDB2V72ByteOrder(generate int) (err error) {
	err = ogrError(ogrSetGenerateDB2V72ByteOrder(generate))
	return
}

func ogrGetGenerateDB2V72ByteOrder() (result int) {
	result = int(C.OGRGetGenerate_DB2_V72_BYTE_ORDER())
	return
}

func OGRGetGenerateDB2V72ByteOrder() (result int) {
	result = ogrGetGenerateDB2V72ByteOrder()
	return
}

// /*! @endcond */

func ogrSetNonLinearGeometriesEnabledFlag(flag int) {
	C.OGRSetNonLinearGeometriesEnabledFlag(C.int(flag))
}

func OGRSetNonLinearGeometriesEnabledFlag(flag int) {
	ogrSetNonLinearGeometriesEnabledFlag(flag)
}

func ogrGetNonLinearGeometriesEnabledFlag() (result int) {
	result = int(C.OGRGetNonLinearGeometriesEnabledFlag())
	return
}

func OGRGetNonLinearGeometriesEnabledFlag() (result int) {
	result = ogrGetNonLinearGeometriesEnabledFlag()
	return
}

func ogrHasPreparedGeometrySupport() (result bool) {
	result = C.OGRHasPreparedGeometrySupport() != 0
	return
}

func OGRHasPreparedGeometrySupport() (result bool) {
	result = ogrHasPreparedGeometrySupport()
	return
}

func ogrCreatePreparedGeometry(g OGRGeometry) (result OGRPreparedGeometry) {
	result = OGRPreparedGeometry{cValue: C.OGRCreatePreparedGeometry(g.cValue)}
	return
}

func OGRCreatePreparedGeometry(g OGRGeometry) (result OGRPreparedGeometry, err error) {
	result = ogrCreatePreparedGeometry(g)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDestroyPreparedGeometry(p OGRPreparedGeometry) {
	C.OGRDestroyPreparedGeometry(p.cValue)
}

func (p OGRPreparedGeometry) Destroy() {
	ogrDestroyPreparedGeometry(p)
}

func ogrPreparedGeometryIntersects(p OGRPreparedGeometry, other OGRGeometry) (result int) {
	result = int(C.OGRPreparedGeometryIntersects(p.cValue, other.cValue))
	return
}

func (p OGRPreparedGeometry) Intersects(other OGRGeometry) (result bool) {
	result = ogrPreparedGeometryIntersects(p, other) != 0
	return
}

func ogrPreparedGeometryContains(p OGRPreparedGeometry, other OGRGeometry) (result int) {
	result = int(C.OGRPreparedGeometryContains(p.cValue, other.cValue))
	return
}

func (p OGRPreparedGeometry) Contains(other OGRGeometry) (result bool) {
	result = ogrPreparedGeometryContains(p, other) != 0
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

func OGRFldCreate(name string, eType OGRFieldType) (result OGRFieldDefn, err error) {
	result = ogrFldCreate(name, eType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrFldDestroy(fld OGRFieldDefn) {
	C.OGR_Fld_Destroy(fld.cValue)
}

func (fld OGRFieldDefn) Destroy() {
	ogrFldDestroy(fld)
}

func ogrFldSetName(fld OGRFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetName(fld.cValue, cs)
}

func (fld OGRFieldDefn) SetName(name string) {
	ogrFldSetName(fld, name)
}

func ogrFldGetNameRef(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetNameRef(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetNameRef() (result string) {
	result = ogrFldGetNameRef(fld)
	return
}

func ogrFldSetAlternativeName(fld OGRFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetAlternativeName(fld.cValue, cs)
}

func (fld OGRFieldDefn) SetAlternativeName(name string) {
	ogrFldSetAlternativeName(fld, name)
}

func ogrFldGetAlternativeNameRef(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetAlternativeNameRef(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetAlternativeNameRef() (result string) {
	result = ogrFldGetAlternativeNameRef(fld)
	return
}

func ogrFldGetType(fld OGRFieldDefn) (result OGRFieldType) {
	result = OGRFieldType(C.OGR_Fld_GetType(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetType() (result OGRFieldType) {
	result = ogrFldGetType(fld)
	return
}

func ogrFldSetType(fld OGRFieldDefn, eType OGRFieldType) {
	C.OGR_Fld_SetType(fld.cValue, C.OGRFieldType(eType))
}

func (fld OGRFieldDefn) SetType(eType OGRFieldType) {
	ogrFldSetType(fld, eType)
}

func ogrFldGetSubType(fld OGRFieldDefn) (result OGRFieldSubType) {
	result = OGRFieldSubType(C.OGR_Fld_GetSubType(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetSubType() (result OGRFieldSubType) {
	result = ogrFldGetSubType(fld)
	return
}

func ogrFldSetSubType(fld OGRFieldDefn, eSubType OGRFieldSubType) {
	C.OGR_Fld_SetSubType(fld.cValue, C.OGRFieldSubType(eSubType))
}

func (fld OGRFieldDefn) SetSubType(eSubType OGRFieldSubType) {
	ogrFldSetSubType(fld, eSubType)
}

func ogrFldGetJustify(fld OGRFieldDefn) (result OGRJustification) {
	result = OGRJustification(C.OGR_Fld_GetJustify(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetJustify() (result OGRJustification) {
	result = ogrFldGetJustify(fld)
	return
}

func ogrFldSetJustify(fld OGRFieldDefn, eJustify OGRJustification) {
	C.OGR_Fld_SetJustify(fld.cValue, C.OGRJustification(eJustify))
}

func (fld OGRFieldDefn) SetJustify(eJustify OGRJustification) {
	ogrFldSetJustify(fld, eJustify)
}

func ogrFldGetWidth(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_GetWidth(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetWidth() (result int) {
	result = ogrFldGetWidth(fld)
	return
}

func ogrFldSetWidth(fld OGRFieldDefn, width int) {
	C.OGR_Fld_SetWidth(fld.cValue, C.int(width))
}

func (fld OGRFieldDefn) SetWidth(width int) {
	ogrFldSetWidth(fld, width)
}

func ogrFldGetPrecision(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_GetPrecision(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetPrecision() (result int) {
	result = ogrFldGetPrecision(fld)
	return
}

func ogrFldSetPrecision(fld OGRFieldDefn, precision int) {
	C.OGR_Fld_SetPrecision(fld.cValue, C.int(precision))
}

func (fld OGRFieldDefn) SetPrecision(precision int) {
	ogrFldSetPrecision(fld, precision)
}

func ogrFldGetTZFlag(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_GetTZFlag(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetTZFlag() (result int) {
	result = ogrFldGetTZFlag(fld)
	return
}

func ogrFldSetTZFlag(fld OGRFieldDefn, tzFlag int) {
	C.OGR_Fld_SetTZFlag(fld.cValue, C.int(tzFlag))
}

func (fld OGRFieldDefn) SetTZFlag(tzFlag int) {
	ogrFldSetTZFlag(fld, tzFlag)
}

func ogrFldSet(fld OGRFieldDefn, name string, eType OGRFieldType, width, precision int, justify OGRJustification) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_Set(fld.cValue, cs, C.OGRFieldType(eType), C.int(width), C.int(precision), C.OGRJustification(justify))
}

func (fld OGRFieldDefn) Set(name string, eType OGRFieldType, width, precision int, justify OGRJustification) {
	ogrFldSet(fld, name, eType, width, precision, justify)
}

func ogrFldIsIgnored(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsIgnored(fld.cValue))
	return
}

func (fld OGRFieldDefn) IsIgnored() (result bool) {
	result = ogrFldIsIgnored(fld) != 0
	return
}

func ogrFldSetIgnored(fld OGRFieldDefn, ignored int) {
	C.OGR_Fld_SetIgnored(fld.cValue, C.int(ignored))
}

func (fld OGRFieldDefn) SetIgnored(ignored int) {
	ogrFldSetIgnored(fld, ignored)
}

func ogrFldIsNullable(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsNullable(fld.cValue))
	return
}

func (fld OGRFieldDefn) IsNullable() (result bool) {
	result = ogrFldIsNullable(fld) != 0
	return
}

func ogrFldSetNullable(fld OGRFieldDefn, nullable int) {
	C.OGR_Fld_SetNullable(fld.cValue, C.int(nullable))
}

func (fld OGRFieldDefn) SetNullable(nullable int) {
	ogrFldSetNullable(fld, nullable)
}

func ogrFldSetGenerated(fld OGRFieldDefn, generated int) {
	C.OGR_Fld_SetGenerated(fld.cValue, C.int(generated))
}

func (fld OGRFieldDefn) SetGenerated(generated int) {
	ogrFldSetGenerated(fld, generated)
}

func ogrFldIsGenerated(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsGenerated(fld.cValue))
	return
}

func (fld OGRFieldDefn) IsGenerated() (result bool) {
	result = ogrFldIsGenerated(fld) != 0
	return
}

func ogrFldIsUnique(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsUnique(fld.cValue))
	return
}

func (fld OGRFieldDefn) IsUnique() (result bool) {
	result = ogrFldIsUnique(fld) != 0
	return
}

func ogrFldSetUnique(fld OGRFieldDefn, unique int) {
	C.OGR_Fld_SetUnique(fld.cValue, C.int(unique))
}

func (fld OGRFieldDefn) SetUnique(unique int) {
	ogrFldSetUnique(fld, unique)
}

func ogrFldGetDefault(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetDefault(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetDefault() (result string) {
	result = ogrFldGetDefault(fld)
	return
}

func ogrFldSetDefault(fld OGRFieldDefn, value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetDefault(fld.cValue, cs)
}

func (fld OGRFieldDefn) SetDefault(value string) {
	ogrFldSetDefault(fld, value)
}

func ogrFldIsDefaultDriverSpecific(fld OGRFieldDefn) (result int) {
	result = int(C.OGR_Fld_IsDefaultDriverSpecific(fld.cValue))
	return
}

func (fld OGRFieldDefn) IsDefaultDriverSpecific() (result bool) {
	result = ogrFldIsDefaultDriverSpecific(fld) != 0
	return
}

func ogrFldGetDomainName(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetDomainName(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetDomainName() (result string) {
	result = ogrFldGetDomainName(fld)
	return
}

func ogrFldSetDomainName(fld OGRFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetDomainName(fld.cValue, cs)
}

func (fld OGRFieldDefn) SetDomainName(name string) {
	ogrFldSetDomainName(fld, name)
}

func ogrFldGetComment(fld OGRFieldDefn) (result string) {
	result = C.GoString(C.OGR_Fld_GetComment(fld.cValue))
	return
}

func (fld OGRFieldDefn) GetComment() (result string) {
	result = ogrFldGetComment(fld)
	return
}

func ogrFldSetComment(fld OGRFieldDefn, comment string) {
	cs := C.CString(comment)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_Fld_SetComment(fld.cValue, cs)
}

func (fld OGRFieldDefn) SetComment(comment string) {
	ogrFldSetComment(fld, comment)
}

func ogrGetFieldTypeName(eType OGRFieldType) (result string) {
	result = C.GoString(C.OGR_GetFieldTypeName(C.OGRFieldType(eType)))
	return
}

func OGRGetFieldTypeName(eType OGRFieldType) (result string) {
	result = ogrGetFieldTypeName(eType)
	return
}

func ogrGetFieldTypeByName(name string) (result OGRFieldType) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRFieldType(C.OGR_GetFieldTypeByName(cs))
	return
}

func OGRGetFieldTypeByName(name string) (result OGRFieldType) {
	result = ogrGetFieldTypeByName(name)
	return
}

func ogrGetFieldSubTypeName(eSubType OGRFieldSubType) (result string) {
	result = C.GoString(C.OGR_GetFieldSubTypeName(C.OGRFieldSubType(eSubType)))
	return
}

func OGRGetFieldSubTypeName(eSubType OGRFieldSubType) (result string) {
	result = ogrGetFieldSubTypeName(eSubType)
	return
}

func ogrGetFieldSubTypeByName(name string) (result OGRFieldSubType) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRFieldSubType(C.OGR_GetFieldSubTypeByName(cs))
	return
}

func OGRGetFieldSubTypeByName(name string) (result OGRFieldSubType) {
	result = ogrGetFieldSubTypeByName(name)
	return
}

func ogrAreTypeSubTypeCompatible(eType OGRFieldType, eSubType OGRFieldSubType) (result int) {
	result = int(C.OGR_AreTypeSubTypeCompatible(C.OGRFieldType(eType), C.OGRFieldSubType(eSubType)))
	return
}

func OGRAreTypeSubTypeCompatible(eType OGRFieldType, eSubType OGRFieldSubType) (result bool) {
	result = ogrAreTypeSubTypeCompatible(eType, eSubType) != 0
	return
}

// /* OGRGeomFieldDefnH */

func ogrGFldCreate(name string, eType OGRwkbGeometryType) (result OGRGeomFieldDefn) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRGeomFieldDefn{cValue: C.OGR_GFld_Create(cs, C.OGRwkbGeometryType(eType))}
	return
}

func OGRGFldCreate(name string, eType OGRwkbGeometryType) (result OGRGeomFieldDefn, err error) {
	result = ogrGFldCreate(name, eType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGFldDestroy(gfld OGRGeomFieldDefn) {
	C.OGR_GFld_Destroy(gfld.cValue)
}

func (gfld OGRGeomFieldDefn) Destroy() {
	ogrGFldDestroy(gfld)
}

func ogrGFldSetName(gfld OGRGeomFieldDefn, name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_GFld_SetName(gfld.cValue, cs)
}

func (gfld OGRGeomFieldDefn) SetName(name string) {
	ogrGFldSetName(gfld, name)
}

func ogrGFldGetNameRef(gfld OGRGeomFieldDefn) (result string) {
	result = C.GoString(C.OGR_GFld_GetNameRef(gfld.cValue))
	return
}

func (gfld OGRGeomFieldDefn) GetNameRef() (result string) {
	result = ogrGFldGetNameRef(gfld)
	return
}

func ogrGFldGetType(gfld OGRGeomFieldDefn) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GFld_GetType(gfld.cValue))
	return
}

func (gfld OGRGeomFieldDefn) GetType() (result OGRwkbGeometryType) {
	result = ogrGFldGetType(gfld)
	return
}

func ogrGFldSetType(gfld OGRGeomFieldDefn, eType OGRwkbGeometryType) {
	C.OGR_GFld_SetType(gfld.cValue, C.OGRwkbGeometryType(eType))
}

func (gfld OGRGeomFieldDefn) SetType(eType OGRwkbGeometryType) {
	ogrGFldSetType(gfld, eType)
}

func ogrGFldGetSpatialRef(gfld OGRGeomFieldDefn) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OGR_GFld_GetSpatialRef(gfld.cValue)}
	return
}

func (gfld OGRGeomFieldDefn) GetSpatialRef() (result OGRSpatialReference) {
	result = ogrGFldGetSpatialRef(gfld)
	return
}

func ogrGFldSetSpatialRef(gfld OGRGeomFieldDefn, sr OGRSpatialReference) {
	C.OGR_GFld_SetSpatialRef(gfld.cValue, sr.cValue)
}

func (gfld OGRGeomFieldDefn) SetSpatialRef(sr OGRSpatialReference) {
	ogrGFldSetSpatialRef(gfld, sr)
}

func ogrGFldIsNullable(gfld OGRGeomFieldDefn) (result int) {
	result = int(C.OGR_GFld_IsNullable(gfld.cValue))
	return
}

func (gfld OGRGeomFieldDefn) IsNullable() (result bool) {
	result = ogrGFldIsNullable(gfld) != 0
	return
}

func ogrGFldSetNullable(gfld OGRGeomFieldDefn, nullable int) {
	C.OGR_GFld_SetNullable(gfld.cValue, C.int(nullable))
}

func (gfld OGRGeomFieldDefn) SetNullable(nullable int) {
	ogrGFldSetNullable(gfld, nullable)
}

func ogrGFldIsIgnored(gfld OGRGeomFieldDefn) (result int) {
	result = int(C.OGR_GFld_IsIgnored(gfld.cValue))
	return
}

func (gfld OGRGeomFieldDefn) IsIgnored() (result bool) {
	result = ogrGFldIsIgnored(gfld) != 0
	return
}

func ogrGFldSetIgnored(gfld OGRGeomFieldDefn, ignored int) {
	C.OGR_GFld_SetIgnored(gfld.cValue, C.int(ignored))
}

func (gfld OGRGeomFieldDefn) SetIgnored(ignored int) {
	ogrGFldSetIgnored(gfld, ignored)
}

func ogrGFldGetCoordinatePrecision(gfld OGRGeomFieldDefn) (result OGRGeomCoordinatePrecision) {
	result = OGRGeomCoordinatePrecision{cValue: C.OGR_GFld_GetCoordinatePrecision(gfld.cValue)}
	return
}

func (gfld OGRGeomFieldDefn) GetCoordinatePrecision() (result OGRGeomCoordinatePrecision) {
	result = ogrGFldGetCoordinatePrecision(gfld)
	return
}

func ogrGFldSetCoordinatePrecision(gfld OGRGeomFieldDefn, precision OGRGeomCoordinatePrecision) {
	C.OGR_GFld_SetCoordinatePrecision(gfld.cValue, precision.cValue)
}

func (gfld OGRGeomFieldDefn) SetCoordinatePrecision(precision OGRGeomCoordinatePrecision) {
	ogrGFldSetCoordinatePrecision(gfld, precision)
}

// /* OGRFeatureDefn */

func ogrFDCreate(name string) (result OGRFeatureDefn) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRFeatureDefn{cValue: C.OGR_FD_Create(cs)}
	return
}

func OGRFDCreate(name string) (result OGRFeatureDefn, err error) {
	result = ogrFDCreate(name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrFDDestroy(fd OGRFeatureDefn) {
	C.OGR_FD_Destroy(fd.cValue)
}

func (fd OGRFeatureDefn) Destroy() {
	ogrFDDestroy(fd)
}

func ogrFDRelease(fd OGRFeatureDefn) {
	C.OGR_FD_Release(fd.cValue)
}

func (fd OGRFeatureDefn) Release() {
	ogrFDRelease(fd)
}

func ogrFDGetName(fd OGRFeatureDefn) (result string) {
	result = C.GoString(C.OGR_FD_GetName(fd.cValue))
	return
}

func (fd OGRFeatureDefn) GetName() (result string) {
	result = ogrFDGetName(fd)
	return
}

func ogrFDGetFieldCount(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_GetFieldCount(fd.cValue))
	return
}

func (fd OGRFeatureDefn) GetFieldCount() (result int) {
	result = ogrFDGetFieldCount(fd)
	return
}

func ogrFDGetFieldDefn(fd OGRFeatureDefn, field int) (result OGRFieldDefn) {
	result = OGRFieldDefn{cValue: C.OGR_FD_GetFieldDefn(fd.cValue, C.int(field))}
	return
}

func (fd OGRFeatureDefn) GetFieldDefn(field int) (result OGRFieldDefn) {
	result = ogrFDGetFieldDefn(fd, field)
	return
}

func ogrFDGetFieldIndex(fd OGRFeatureDefn, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_FD_GetFieldIndex(fd.cValue, cs))
	return
}

func (fd OGRFeatureDefn) GetFieldIndex(name string) (result int) {
	result = ogrFDGetFieldIndex(fd, name)
	return
}

func ogrFDAddFieldDefn(fd OGRFeatureDefn, fld OGRFieldDefn) {
	C.OGR_FD_AddFieldDefn(fd.cValue, fld.cValue)
}

func (fd OGRFeatureDefn) AddFieldDefn(fld OGRFieldDefn) {
	ogrFDAddFieldDefn(fd, fld)
}

func ogrFDDeleteFieldDefn(fd OGRFeatureDefn, field int) (result OGRErr) {
	result = OGRErr(C.OGR_FD_DeleteFieldDefn(fd.cValue, C.int(field)))
	return
}

func (fd OGRFeatureDefn) DeleteFieldDefn(field int) (err error) {
	err = ogrError(ogrFDDeleteFieldDefn(fd, field))
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

func (fd OGRFeatureDefn) ReorderFieldDefns(panMap []int) (err error) {
	err = ogrError(ogrFDReorderFieldDefns(fd, panMap))
	return
}

func ogrFDGetGeomType(fd OGRFeatureDefn) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_FD_GetGeomType(fd.cValue))
	return
}

func (fd OGRFeatureDefn) GetGeomType() (result OGRwkbGeometryType) {
	result = ogrFDGetGeomType(fd)
	return
}

func ogrFDSetGeomType(fd OGRFeatureDefn, eType OGRwkbGeometryType) {
	C.OGR_FD_SetGeomType(fd.cValue, C.OGRwkbGeometryType(eType))
}

func (fd OGRFeatureDefn) SetGeomType(eType OGRwkbGeometryType) {
	ogrFDSetGeomType(fd, eType)
}

func ogrFDIsGeometryIgnored(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_IsGeometryIgnored(fd.cValue))
	return
}

func (fd OGRFeatureDefn) IsGeometryIgnored() (result bool) {
	result = ogrFDIsGeometryIgnored(fd) != 0
	return
}

func ogrFDSetGeometryIgnored(fd OGRFeatureDefn, ignored int) {
	C.OGR_FD_SetGeometryIgnored(fd.cValue, C.int(ignored))
}

func (fd OGRFeatureDefn) SetGeometryIgnored(ignored int) {
	ogrFDSetGeometryIgnored(fd, ignored)
}

func ogrFDIsStyleIgnored(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_IsStyleIgnored(fd.cValue))
	return
}

func (fd OGRFeatureDefn) IsStyleIgnored() (result bool) {
	result = ogrFDIsStyleIgnored(fd) != 0
	return
}

func ogrFDSetStyleIgnored(fd OGRFeatureDefn, ignored int) {
	C.OGR_FD_SetStyleIgnored(fd.cValue, C.int(ignored))
}

func (fd OGRFeatureDefn) SetStyleIgnored(ignored int) {
	ogrFDSetStyleIgnored(fd, ignored)
}

func ogrFDReference(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_Reference(fd.cValue))
	return
}

func (fd OGRFeatureDefn) Reference() (result int) {
	result = ogrFDReference(fd)
	return
}

func ogrFDDereference(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_Dereference(fd.cValue))
	return
}

func (fd OGRFeatureDefn) Dereference() (result int) {
	result = ogrFDDereference(fd)
	return
}

func ogrFDGetReferenceCount(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_GetReferenceCount(fd.cValue))
	return
}

func (fd OGRFeatureDefn) GetReferenceCount() (result int) {
	result = ogrFDGetReferenceCount(fd)
	return
}

func ogrFDGetGeomFieldCount(fd OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_GetGeomFieldCount(fd.cValue))
	return
}

func (fd OGRFeatureDefn) GetGeomFieldCount() (result int) {
	result = ogrFDGetGeomFieldCount(fd)
	return
}

func ogrFDGetGeomFieldDefn(fd OGRFeatureDefn, geomField int) (result OGRGeomFieldDefn) {
	result = OGRGeomFieldDefn{cValue: C.OGR_FD_GetGeomFieldDefn(fd.cValue, C.int(geomField))}
	return
}

func (fd OGRFeatureDefn) GetGeomFieldDefn(geomField int) (result OGRGeomFieldDefn) {
	result = ogrFDGetGeomFieldDefn(fd, geomField)
	return
}

func ogrFDGetGeomFieldIndex(fd OGRFeatureDefn, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_FD_GetGeomFieldIndex(fd.cValue, cs))
	return
}

func (fd OGRFeatureDefn) GetGeomFieldIndex(name string) (result int) {
	result = ogrFDGetGeomFieldIndex(fd, name)
	return
}

func ogrFDAddGeomFieldDefn(fd OGRFeatureDefn, gfld OGRGeomFieldDefn) {
	C.OGR_FD_AddGeomFieldDefn(fd.cValue, gfld.cValue)
}

func (fd OGRFeatureDefn) AddGeomFieldDefn(gfld OGRGeomFieldDefn) {
	ogrFDAddGeomFieldDefn(fd, gfld)
}

func ogrFDDeleteGeomFieldDefn(fd OGRFeatureDefn, geomField int) (result OGRErr) {
	result = OGRErr(C.OGR_FD_DeleteGeomFieldDefn(fd.cValue, C.int(geomField)))
	return
}

func (fd OGRFeatureDefn) DeleteGeomFieldDefn(geomField int) (err error) {
	err = ogrError(ogrFDDeleteGeomFieldDefn(fd, geomField))
	return
}

func ogrFDIsSame(fd, other OGRFeatureDefn) (result int) {
	result = int(C.OGR_FD_IsSame(fd.cValue, other.cValue))
	return
}

func (fd OGRFeatureDefn) IsSame(other OGRFeatureDefn) (result bool) {
	result = ogrFDIsSame(fd, other) != 0
	return
}

// /* OGRFeature */

func ogrFCreate(fd OGRFeatureDefn) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_F_Create(fd.cValue)}
	return
}

func OGRFCreate(fd OGRFeatureDefn) (result OGRFeature, err error) {
	result = ogrFCreate(fd)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrFDestroy(feat OGRFeature) {
	C.OGR_F_Destroy(feat.cValue)
}

func (feat OGRFeature) Destroy() {
	ogrFDestroy(feat)
}

func ogrFGetDefnRef(feat OGRFeature) (result OGRFeatureDefn) {
	result = OGRFeatureDefn{cValue: C.OGR_F_GetDefnRef(feat.cValue)}
	return
}

func (feat OGRFeature) GetDefnRef() (result OGRFeatureDefn) {
	result = ogrFGetDefnRef(feat)
	return
}

func ogrFSetGeometryDirectly(feat OGRFeature, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeometryDirectly(feat.cValue, geom.cValue))
	return
}

func (feat OGRFeature) SetGeometryDirectly(geom OGRGeometry) (err error) {
	err = ogrError(ogrFSetGeometryDirectly(feat, geom))
	return
}

func ogrFSetGeometry(feat OGRFeature, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeometry(feat.cValue, geom.cValue))
	return
}

func (feat OGRFeature) SetGeometry(geom OGRGeometry) (err error) {
	err = ogrError(ogrFSetGeometry(feat, geom))
	return
}

func ogrFGetGeometryRef(feat OGRFeature) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_GetGeometryRef(feat.cValue)}
	return
}

func (feat OGRFeature) GetGeometryRef() (result OGRGeometry) {
	result = ogrFGetGeometryRef(feat)
	return
}

func ogrFStealGeometry(feat OGRFeature) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_StealGeometry(feat.cValue)}
	return
}

func (feat OGRFeature) StealGeometry() (result OGRGeometry) {
	result = ogrFStealGeometry(feat)
	return
}

func ogrFStealGeometryEx(feat OGRFeature, geomField int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_StealGeometryEx(feat.cValue, C.int(geomField))}
	return
}

func (feat OGRFeature) StealGeometryEx(geomField int) (result OGRGeometry) {
	result = ogrFStealGeometryEx(feat, geomField)
	return
}

func ogrFClone(feat OGRFeature) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_F_Clone(feat.cValue)}
	return
}

func (feat OGRFeature) Clone() (result OGRFeature, err error) {
	result = ogrFClone(feat)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrFEqual(feat, other OGRFeature) (result int) {
	result = int(C.OGR_F_Equal(feat.cValue, other.cValue))
	return
}

func (feat OGRFeature) Equal(other OGRFeature) (result bool) {
	result = ogrFEqual(feat, other) != 0
	return
}

func ogrFGetFieldCount(feat OGRFeature) (result int) {
	result = int(C.OGR_F_GetFieldCount(feat.cValue))
	return
}

func (feat OGRFeature) GetFieldCount() (result int) {
	result = ogrFGetFieldCount(feat)
	return
}

func ogrFGetFieldDefnRef(feat OGRFeature, field int) (result OGRFieldDefn) {
	result = OGRFieldDefn{cValue: C.OGR_F_GetFieldDefnRef(feat.cValue, C.int(field))}
	return
}

func (feat OGRFeature) GetFieldDefnRef(field int) (result OGRFieldDefn) {
	result = ogrFGetFieldDefnRef(feat, field)
	return
}

func ogrFGetFieldIndex(feat OGRFeature, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_F_GetFieldIndex(feat.cValue, cs))
	return
}

func (feat OGRFeature) GetFieldIndex(name string) (result int) {
	result = ogrFGetFieldIndex(feat, name)
	return
}

func ogrFIsFieldSet(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_IsFieldSet(feat.cValue, C.int(field)))
	return
}

func (feat OGRFeature) IsFieldSet(field int) (result bool) {
	result = ogrFIsFieldSet(feat, field) != 0
	return
}

func ogrFUnsetField(feat OGRFeature, field int) {
	C.OGR_F_UnsetField(feat.cValue, C.int(field))
}

func (feat OGRFeature) UnsetField(field int) {
	ogrFUnsetField(feat, field)
}

func ogrFIsFieldNull(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_IsFieldNull(feat.cValue, C.int(field)))
	return
}

func (feat OGRFeature) IsFieldNull(field int) (result bool) {
	result = ogrFIsFieldNull(feat, field) != 0
	return
}

func ogrFIsFieldSetAndNotNull(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_IsFieldSetAndNotNull(feat.cValue, C.int(field)))
	return
}

func (feat OGRFeature) IsFieldSetAndNotNull(field int) (result bool) {
	result = ogrFIsFieldSetAndNotNull(feat, field) != 0
	return
}

func ogrFSetFieldNull(feat OGRFeature, field int) {
	C.OGR_F_SetFieldNull(feat.cValue, C.int(field))
}

func (feat OGRFeature) SetFieldNull(field int) {
	ogrFSetFieldNull(feat, field)
}

func ogrFGetRawFieldRef(feat OGRFeature, field int) (result OGRField) {
	result = OGRField{cValue: C.OGR_F_GetRawFieldRef(feat.cValue, C.int(field))}
	return
}

func (feat OGRFeature) GetRawFieldRef(field int) (result OGRField) {
	result = ogrFGetRawFieldRef(feat, field)
	return
}

func ogrRawFieldIsUnset(f OGRField) (result int) {
	result = int(C.OGR_RawField_IsUnset(f.cValue))
	return
}

func (f OGRField) IsUnset() (result bool) {
	result = ogrRawFieldIsUnset(f) != 0
	return
}

func ogrRawFieldIsNull(f OGRField) (result int) {
	result = int(C.OGR_RawField_IsNull(f.cValue))
	return
}

func (f OGRField) IsNull() (result bool) {
	result = ogrRawFieldIsNull(f) != 0
	return
}

func ogrRawFieldSetUnset(f OGRField) {
	C.OGR_RawField_SetUnset(f.cValue)
}

func (f OGRField) SetUnset() {
	ogrRawFieldSetUnset(f)
}

func ogrRawFieldSetNull(f OGRField) {
	C.OGR_RawField_SetNull(f.cValue)
}

func (f OGRField) SetNull() {
	ogrRawFieldSetNull(f)
}

func ogrFGetFieldAsInteger(feat OGRFeature, field int) (result int) {
	result = int(C.OGR_F_GetFieldAsInteger(feat.cValue, C.int(field)))
	return
}

func (feat OGRFeature) GetFieldAsInteger(field int) (result int) {
	result = ogrFGetFieldAsInteger(feat, field)
	return
}

func ogrFGetFieldAsInteger64(feat OGRFeature, field int) (result int64) {
	result = int64(C.OGR_F_GetFieldAsInteger64(feat.cValue, C.int(field)))
	return
}

func (feat OGRFeature) GetFieldAsInteger64(field int) (result int64) {
	result = ogrFGetFieldAsInteger64(feat, field)
	return
}

func ogrFGetFieldAsDouble(feat OGRFeature, field int) (result float64) {
	result = float64(C.OGR_F_GetFieldAsDouble(feat.cValue, C.int(field)))
	return
}

func (feat OGRFeature) GetFieldAsDouble(field int) (result float64) {
	result = ogrFGetFieldAsDouble(feat, field)
	return
}

func ogrFGetFieldAsString(feat OGRFeature, field int) (result string) {
	result = C.GoString(C.OGR_F_GetFieldAsString(feat.cValue, C.int(field)))
	return
}

func (feat OGRFeature) GetFieldAsString(field int) (result string) {
	result = ogrFGetFieldAsString(feat, field)
	return
}

func ogrFGetFieldAsISO8601DateTime(feat OGRFeature, field int, options []string) (result string) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = C.GoString(C.OGR_F_GetFieldAsISO8601DateTime(feat.cValue, C.int(field), &cOptions[0]))
	return
}

func (feat OGRFeature) GetFieldAsISO8601DateTime(field int, options []string) (result string) {
	result = ogrFGetFieldAsISO8601DateTime(feat, field, options)
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

func (feat OGRFeature) GetFieldAsIntegerList(field int) (result []int) {
	result = ogrFGetFieldAsIntegerList(feat, field)
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

func (feat OGRFeature) GetFieldAsInteger64List(field int) (result []int64) {
	result = ogrFGetFieldAsInteger64List(feat, field)
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

func (feat OGRFeature) GetFieldAsDoubleList(field int) (result []float64) {
	result = ogrFGetFieldAsDoubleList(feat, field)
	return
}

func ogrFGetFieldAsStringList(feat OGRFeature, field int) (result []string) {
	raw := C.OGR_F_GetFieldAsStringList(feat.cValue, C.int(field))
	n := int(C.CSLCount(raw))
	for i := 0; i < n; i++ {
		result = append(result, C.GoString(C.CSLGetField(raw, C.int(i))))
	}
	return
}

func (feat OGRFeature) GetFieldAsStringList(field int) (result []string) {
	result = ogrFGetFieldAsStringList(feat, field)
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

func (feat OGRFeature) GetFieldAsBinary(field int) (result []byte) {
	result = ogrFGetFieldAsBinary(feat, field)
	return
}

func ogrFGetFieldAsDateTime(feat OGRFeature, field int) (year, month, day, hour, minute, second, tzFlag int, ok bool) {
	var cy, cmo, cd, ch, cmi, cse, ctz C.int
	r := C.OGR_F_GetFieldAsDateTime(feat.cValue, C.int(field), &cy, &cmo, &cd, &ch, &cmi, &cse, &ctz)
	year, month, day, hour, minute, second, tzFlag = int(cy), int(cmo), int(cd), int(ch), int(cmi), int(cse), int(ctz)
	ok = r != 0
	return
}

func (feat OGRFeature) GetFieldAsDateTime(field int) (year, month, day, hour, minute, second, tzFlag int, ok bool) {
	year, month, day, hour, minute, second, tzFlag, ok = ogrFGetFieldAsDateTime(feat, field)
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

func (feat OGRFeature) GetFieldAsDateTimeEx(field int) (year, month, day, hour, minute int, second float32, tzFlag int, ok bool) {
	year, month, day, hour, minute, second, tzFlag, ok = ogrFGetFieldAsDateTimeEx(feat, field)
	return
}

func ogrFSetFieldInteger(feat OGRFeature, field, value int) {
	C.OGR_F_SetFieldInteger(feat.cValue, C.int(field), C.int(value))
}

func (feat OGRFeature) SetFieldInteger(field, value int) {
	ogrFSetFieldInteger(feat, field, value)
}

func ogrFSetFieldInteger64(feat OGRFeature, field int, value int64) {
	C.OGR_F_SetFieldInteger64(feat.cValue, C.int(field), C.GIntBig(value))
}

func (feat OGRFeature) SetFieldInteger64(field int, value int64) {
	ogrFSetFieldInteger64(feat, field, value)
}

func ogrFSetFieldDouble(feat OGRFeature, field int, value float64) {
	C.OGR_F_SetFieldDouble(feat.cValue, C.int(field), C.double(value))
}

func (feat OGRFeature) SetFieldDouble(field int, value float64) {
	ogrFSetFieldDouble(feat, field, value)
}

func ogrFSetFieldString(feat OGRFeature, field int, value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetFieldString(feat.cValue, C.int(field), cs)
}

func (feat OGRFeature) SetFieldString(field int, value string) {
	ogrFSetFieldString(feat, field, value)
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

func (feat OGRFeature) SetFieldIntegerList(field int, values []int) {
	ogrFSetFieldIntegerList(feat, field, values)
}

func ogrFSetFieldInteger64List(feat OGRFeature, field int, values []int64) {
	var p *C.GIntBig
	if len(values) > 0 {
		p = (*C.GIntBig)(unsafe.Pointer(&values[0]))
	}
	C.OGR_F_SetFieldInteger64List(feat.cValue, C.int(field), C.int(len(values)), p)
}

func (feat OGRFeature) SetFieldInteger64List(field int, values []int64) {
	ogrFSetFieldInteger64List(feat, field, values)
}

func ogrFSetFieldDoubleList(feat OGRFeature, field int, values []float64) {
	var p *C.double
	if len(values) > 0 {
		p = (*C.double)(unsafe.Pointer(&values[0]))
	}
	C.OGR_F_SetFieldDoubleList(feat.cValue, C.int(field), C.int(len(values)), p)
}

func (feat OGRFeature) SetFieldDoubleList(field int, values []float64) {
	ogrFSetFieldDoubleList(feat, field, values)
}

func ogrFSetFieldStringList(feat OGRFeature, field int, values []string) {
	cValues := make([]*C.char, len(values)+1)
	for i, v := range values {
		cValues[i] = C.CString(v)
		defer C.free(unsafe.Pointer(cValues[i]))
	}
	cValues[len(values)] = nil
	C.OGR_F_SetFieldStringList(feat.cValue, C.int(field), &cValues[0])
}

func (feat OGRFeature) SetFieldStringList(field int, values []string) {
	ogrFSetFieldStringList(feat, field, values)
}

func ogrFSetFieldRaw(feat OGRFeature, field int, value OGRField) {
	C.OGR_F_SetFieldRaw(feat.cValue, C.int(field), value.cValue)
}

func (feat OGRFeature) SetFieldRaw(field int, value OGRField) {
	ogrFSetFieldRaw(feat, field, value)
}

func ogrFSetFieldBinary(feat OGRFeature, field int, data []byte) {
	var p unsafe.Pointer
	if len(data) > 0 {
		p = unsafe.Pointer(&data[0])
	}
	C.OGR_F_SetFieldBinary(feat.cValue, C.int(field), C.int(len(data)), p)
}

func (feat OGRFeature) SetFieldBinary(field int, data []byte) {
	ogrFSetFieldBinary(feat, field, data)
}

func ogrFSetFieldDateTime(feat OGRFeature, field, year, month, day, hour, minute, second, tzFlag int) {
	C.OGR_F_SetFieldDateTime(feat.cValue, C.int(field), C.int(year), C.int(month), C.int(day), C.int(hour), C.int(minute), C.int(second), C.int(tzFlag))
}

func (feat OGRFeature) SetFieldDateTime(field, year, month, day, hour, minute, second, tzFlag int) {
	ogrFSetFieldDateTime(feat, field, year, month, day, hour, minute, second, tzFlag)
}

func ogrFSetFieldDateTimeEx(feat OGRFeature, field, year, month, day, hour, minute int, second float32, tzFlag int) {
	C.OGR_F_SetFieldDateTimeEx(feat.cValue, C.int(field), C.int(year), C.int(month), C.int(day), C.int(hour), C.int(minute), C.float(second), C.int(tzFlag))
}

func (feat OGRFeature) SetFieldDateTimeEx(field, year, month, day, hour, minute int, second float32, tzFlag int) {
	ogrFSetFieldDateTimeEx(feat, field, year, month, day, hour, minute, second, tzFlag)
}

func ogrFGetGeomFieldCount(feat OGRFeature) (result int) {
	result = int(C.OGR_F_GetGeomFieldCount(feat.cValue))
	return
}

func (feat OGRFeature) GetGeomFieldCount() (result int) {
	result = ogrFGetGeomFieldCount(feat)
	return
}

func ogrFGetGeomFieldDefnRef(feat OGRFeature, field int) (result OGRGeomFieldDefn) {
	result = OGRGeomFieldDefn{cValue: C.OGR_F_GetGeomFieldDefnRef(feat.cValue, C.int(field))}
	return
}

func (feat OGRFeature) GetGeomFieldDefnRef(field int) (result OGRGeomFieldDefn) {
	result = ogrFGetGeomFieldDefnRef(feat, field)
	return
}

func ogrFGetGeomFieldIndex(feat OGRFeature, name string) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_F_GetGeomFieldIndex(feat.cValue, cs))
	return
}

func (feat OGRFeature) GetGeomFieldIndex(name string) (result int) {
	result = ogrFGetGeomFieldIndex(feat, name)
	return
}

func ogrFGetGeomFieldRef(feat OGRFeature, field int) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_F_GetGeomFieldRef(feat.cValue, C.int(field))}
	return
}

func (feat OGRFeature) GetGeomFieldRef(field int) (result OGRGeometry) {
	result = ogrFGetGeomFieldRef(feat, field)
	return
}

func ogrFSetGeomFieldDirectly(feat OGRFeature, field int, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeomFieldDirectly(feat.cValue, C.int(field), geom.cValue))
	return
}

func (feat OGRFeature) SetGeomFieldDirectly(field int, geom OGRGeometry) (err error) {
	err = ogrError(ogrFSetGeomFieldDirectly(feat, field, geom))
	return
}

func ogrFSetGeomField(feat OGRFeature, field int, geom OGRGeometry) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetGeomField(feat.cValue, C.int(field), geom.cValue))
	return
}

func (feat OGRFeature) SetGeomField(field int, geom OGRGeometry) (err error) {
	err = ogrError(ogrFSetGeomField(feat, field, geom))
	return
}

func ogrFGetFID(feat OGRFeature) (result int64) {
	result = int64(C.OGR_F_GetFID(feat.cValue))
	return
}

func (feat OGRFeature) GetFID() (result int64) {
	result = ogrFGetFID(feat)
	return
}

func ogrFSetFID(feat OGRFeature, fid int64) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetFID(feat.cValue, C.GIntBig(fid)))
	return
}

func (feat OGRFeature) SetFID(fid int64) (err error) {
	err = ogrError(ogrFSetFID(feat, fid))
	return
}

// void CPL_DLL OGR_F_DumpReadable(OGRFeatureH, FILE *);

func ogrFDumpReadableAsString(feat OGRFeature, options []string) (result string) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	cs := C.OGR_F_DumpReadableAsString(feat.cValue, &cOptions[0])
	if cs != nil {
		result = C.GoString(cs)
		C.CPLFree(unsafe.Pointer(cs))
	}
	return
}

func (feat OGRFeature) DumpReadableAsString(options []string) (result string) {
	result = ogrFDumpReadableAsString(feat, options)
	return
}

func ogrFSetFrom(feat, other OGRFeature, forgiving int) (result OGRErr) {
	result = OGRErr(C.OGR_F_SetFrom(feat.cValue, other.cValue, C.int(forgiving)))
	return
}

func (feat OGRFeature) SetFrom(other OGRFeature, forgiving int) (err error) {
	err = ogrError(ogrFSetFrom(feat, other, forgiving))
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

func (feat OGRFeature) SetFromWithMap(other OGRFeature, forgiving int, panMap []int) (err error) {
	err = ogrError(ogrFSetFromWithMap(feat, other, forgiving, panMap))
	return
}

func ogrFGetStyleString(feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_F_GetStyleString(feat.cValue))
	return
}

func (feat OGRFeature) GetStyleString() (result string) {
	result = ogrFGetStyleString(feat)
	return
}

func ogrFSetStyleString(feat OGRFeature, style string) {
	cs := C.CString(style)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetStyleString(feat.cValue, cs)
}

func (feat OGRFeature) SetStyleString(style string) {
	ogrFSetStyleString(feat, style)
}

func ogrFSetStyleStringDirectly(feat OGRFeature, style string) {
	cs := C.CString(style)
	C.OGR_F_SetStyleStringDirectly(feat.cValue, cs)
}

func (feat OGRFeature) SetStyleStringDirectly(style string) {
	ogrFSetStyleStringDirectly(feat, style)
}

// /** Return style table */
func ogrFGetStyleTable(feat OGRFeature) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_F_GetStyleTable(feat.cValue)}
	return
}

// /** Return style table */
func (feat OGRFeature) GetStyleTable() (result OGRStyleTable) {
	result = ogrFGetStyleTable(feat)
	return
}

// /** Set style table and take ownership */
func ogrFSetStyleTableDirectly(feat OGRFeature, styleTable OGRStyleTable) {
	C.OGR_F_SetStyleTableDirectly(feat.cValue, styleTable.cValue)
}

// /** Set style table and take ownership */
func (feat OGRFeature) SetStyleTableDirectly(styleTable OGRStyleTable) {
	ogrFSetStyleTableDirectly(feat, styleTable)
}

// /** Set style table */
func ogrFSetStyleTable(feat OGRFeature, styleTable OGRStyleTable) {
	C.OGR_F_SetStyleTable(feat.cValue, styleTable.cValue)
}

// /** Set style table */
func (feat OGRFeature) SetStyleTable(styleTable OGRStyleTable) {
	ogrFSetStyleTable(feat, styleTable)
}

func ogrFGetNativeData(feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_F_GetNativeData(feat.cValue))
	return
}

func (feat OGRFeature) GetNativeData() (result string) {
	result = ogrFGetNativeData(feat)
	return
}

func ogrFSetNativeData(feat OGRFeature, data string) {
	cs := C.CString(data)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetNativeData(feat.cValue, cs)
}

func (feat OGRFeature) SetNativeData(data string) {
	ogrFSetNativeData(feat, data)
}

func ogrFGetNativeMediaType(feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_F_GetNativeMediaType(feat.cValue))
	return
}

func (feat OGRFeature) GetNativeMediaType() (result string) {
	result = ogrFGetNativeMediaType(feat)
	return
}

func ogrFSetNativeMediaType(feat OGRFeature, mediaType string) {
	cs := C.CString(mediaType)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_F_SetNativeMediaType(feat.cValue, cs)
}

func (feat OGRFeature) SetNativeMediaType(mediaType string) {
	ogrFSetNativeMediaType(feat, mediaType)
}

func ogrFFillUnsetWithDefault(feat OGRFeature, notNullableOnly int, options []string) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	C.OGR_F_FillUnsetWithDefault(feat.cValue, C.int(notNullableOnly), &cOptions[0])
}

func (feat OGRFeature) FillUnsetWithDefault(notNullableOnly int, options []string) {
	ogrFFillUnsetWithDefault(feat, notNullableOnly, options)
}

func ogrFValidate(feat OGRFeature, validateFlags, emitError int) (result int) {
	result = int(C.OGR_F_Validate(feat.cValue, C.int(validateFlags), C.int(emitError)))
	return
}

func (feat OGRFeature) Validate(validateFlags, emitError int) (result bool) {
	result = ogrFValidate(feat, validateFlags, emitError) != 0
	return
}

// /* OGRFieldDomain */

func ogrFldDomainDestroy(dom OGRFieldDomain) {
	C.OGR_FldDomain_Destroy(dom.cValue)
}

func (dom OGRFieldDomain) Destroy() {
	ogrFldDomainDestroy(dom)
}

func ogrFldDomainGetName(dom OGRFieldDomain) (result string) {
	result = C.GoString(C.OGR_FldDomain_GetName(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetName() (result string) {
	result = ogrFldDomainGetName(dom)
	return
}

func ogrFldDomainGetDescription(dom OGRFieldDomain) (result string) {
	result = C.GoString(C.OGR_FldDomain_GetDescription(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetDescription() (result string) {
	result = ogrFldDomainGetDescription(dom)
	return
}

func ogrFldDomainGetDomainType(dom OGRFieldDomain) (result OGRFieldDomainType) {
	result = OGRFieldDomainType(C.OGR_FldDomain_GetDomainType(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetDomainType() (result OGRFieldDomainType) {
	result = ogrFldDomainGetDomainType(dom)
	return
}

func ogrFldDomainGetFieldType(dom OGRFieldDomain) (result OGRFieldType) {
	result = OGRFieldType(C.OGR_FldDomain_GetFieldType(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetFieldType() (result OGRFieldType) {
	result = ogrFldDomainGetFieldType(dom)
	return
}

func ogrFldDomainGetFieldSubType(dom OGRFieldDomain) (result OGRFieldSubType) {
	result = OGRFieldSubType(C.OGR_FldDomain_GetFieldSubType(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetFieldSubType() (result OGRFieldSubType) {
	result = ogrFldDomainGetFieldSubType(dom)
	return
}

func ogrFldDomainGetSplitPolicy(dom OGRFieldDomain) (result OGRFieldDomainSplitPolicy) {
	result = OGRFieldDomainSplitPolicy(C.OGR_FldDomain_GetSplitPolicy(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetSplitPolicy() (result OGRFieldDomainSplitPolicy) {
	result = ogrFldDomainGetSplitPolicy(dom)
	return
}

func ogrFldDomainSetSplitPolicy(dom OGRFieldDomain, policy OGRFieldDomainSplitPolicy) {
	C.OGR_FldDomain_SetSplitPolicy(dom.cValue, C.OGRFieldDomainSplitPolicy(policy))
}

func (dom OGRFieldDomain) SetSplitPolicy(policy OGRFieldDomainSplitPolicy) {
	ogrFldDomainSetSplitPolicy(dom, policy)
}

func ogrFldDomainGetMergePolicy(dom OGRFieldDomain) (result OGRFieldDomainMergePolicy) {
	result = OGRFieldDomainMergePolicy(C.OGR_FldDomain_GetMergePolicy(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetMergePolicy() (result OGRFieldDomainMergePolicy) {
	result = ogrFldDomainGetMergePolicy(dom)
	return
}

func ogrFldDomainSetMergePolicy(dom OGRFieldDomain, policy OGRFieldDomainMergePolicy) {
	C.OGR_FldDomain_SetMergePolicy(dom.cValue, C.OGRFieldDomainMergePolicy(policy))
}

func (dom OGRFieldDomain) SetMergePolicy(policy OGRFieldDomainMergePolicy) {
	ogrFldDomainSetMergePolicy(dom, policy)
}

func ogrCodedFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, enumeration OGRCodedValue) (result OGRFieldDomain) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csDesc := C.CString(description)
	defer C.free(unsafe.Pointer(csDesc))
	result = OGRFieldDomain{cValue: C.OGR_CodedFldDomain_Create(csName, csDesc, C.OGRFieldType(eFieldType), C.OGRFieldSubType(eFieldSubType), enumeration.cValue)}
	return
}

func OGRCodedFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, enumeration OGRCodedValue) (result OGRFieldDomain, err error) {
	result = ogrCodedFldDomainCreate(name, description, eFieldType, eFieldSubType, enumeration)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrCodedFldDomainGetEnumeration(dom OGRFieldDomain) (result OGRCodedValue) {
	result = OGRCodedValue{cValue: C.OGR_CodedFldDomain_GetEnumeration(dom.cValue)}
	return
}

func (dom OGRFieldDomain) GetEnumeration() (result OGRCodedValue) {
	result = ogrCodedFldDomainGetEnumeration(dom)
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

func OGRRangeFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, min OGRField, minInclusive bool, max OGRField, maxInclusive bool) (result OGRFieldDomain, err error) {
	result = ogrRangeFldDomainCreate(name, description, eFieldType, eFieldSubType, min, minInclusive, max, maxInclusive)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrRangeFldDomainGetMin(dom OGRFieldDomain) (result OGRField, inclusive bool) {
	var inc C.bool
	result = OGRField{cValue: C.OGR_RangeFldDomain_GetMin(dom.cValue, &inc)}
	inclusive = bool(inc)
	return
}

func (dom OGRFieldDomain) GetMin() (result OGRField, inclusive bool) {
	result, inclusive = ogrRangeFldDomainGetMin(dom)
	return
}

func ogrRangeFldDomainGetMax(dom OGRFieldDomain) (result OGRField, inclusive bool) {
	var inc C.bool
	result = OGRField{cValue: C.OGR_RangeFldDomain_GetMax(dom.cValue, &inc)}
	inclusive = bool(inc)
	return
}

func (dom OGRFieldDomain) GetMax() (result OGRField, inclusive bool) {
	result, inclusive = ogrRangeFldDomainGetMax(dom)
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

func OGRGlobFldDomainCreate(name, description string, eFieldType OGRFieldType, eFieldSubType OGRFieldSubType, glob string) (result OGRFieldDomain, err error) {
	result = ogrGlobFldDomainCreate(name, description, eFieldType, eFieldSubType, glob)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGlobFldDomainGetGlob(dom OGRFieldDomain) (result string) {
	result = C.GoString(C.OGR_GlobFldDomain_GetGlob(dom.cValue))
	return
}

func (dom OGRFieldDomain) GetGlob() (result string) {
	result = ogrGlobFldDomainGetGlob(dom)
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

func (l OGRLayer) GetName() (result string) {
	result = ogrLGetName(l)
	return
}

func ogrLGetGeomType(l OGRLayer) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_L_GetGeomType(l.cValue))
	return
}

func (l OGRLayer) GetGeomType() (result OGRwkbGeometryType) {
	result = ogrLGetGeomType(l)
	return
}

// /* Defined in gdal.h to avoid circular dependency with ogr_api.h */
// /* GDALDatasetH CPL_DLL OGR_L_GetDataset(OGRLayerH hLayer); */

// /** Result item of OGR_L_GetGeometryTypes */
type OGRGeometryTypeCounter struct {
	cValue *C.OGRGeometryTypeCounter
}

func (c OGRGeometryTypeCounter) GeomType() (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(c.cValue.eGeomType)
	return
}

func (c OGRGeometryTypeCounter) Count() (result int64) {
	result = int64(c.cValue.nCount)
	return
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
	C.CPLFree(unsafe.Pointer(p))
	return
}

func (l OGRLayer) GetGeometryTypes(iGeomField, flags int, progress GDALProgressFunc, progressData unsafe.Pointer) (result []OGRGeometryTypeCounter) {
	result = ogrLGetGeometryTypes(l, iGeomField, flags, progress, progressData)
	return
}

func ogrLGetSpatialFilter(l OGRLayer) (result OGRGeometry) {
	result = OGRGeometry{cValue: C.OGR_L_GetSpatialFilter(l.cValue)}
	return
}

func (l OGRLayer) GetSpatialFilter() (result OGRGeometry) {
	result = ogrLGetSpatialFilter(l)
	return
}

func ogrLSetSpatialFilter(l OGRLayer, geom OGRGeometry) {
	C.OGR_L_SetSpatialFilter(l.cValue, geom.cValue)
}

func (l OGRLayer) SetSpatialFilter(geom OGRGeometry) {
	ogrLSetSpatialFilter(l, geom)
}

func ogrLSetSpatialFilterRect(l OGRLayer, minX, minY, maxX, maxY float64) {
	C.OGR_L_SetSpatialFilterRect(l.cValue, C.double(minX), C.double(minY), C.double(maxX), C.double(maxY))
}

func (l OGRLayer) SetSpatialFilterRect(minX, minY, maxX, maxY float64) {
	ogrLSetSpatialFilterRect(l, minX, minY, maxX, maxY)
}

func ogrLSetSpatialFilterEx(l OGRLayer, iGeomField int, geom OGRGeometry) {
	C.OGR_L_SetSpatialFilterEx(l.cValue, C.int(iGeomField), geom.cValue)
}

func (l OGRLayer) SetSpatialFilterEx(iGeomField int, geom OGRGeometry) {
	ogrLSetSpatialFilterEx(l, iGeomField, geom)
}

func ogrLSetSpatialFilterRectEx(l OGRLayer, iGeomField int, minX, minY, maxX, maxY float64) {
	C.OGR_L_SetSpatialFilterRectEx(l.cValue, C.int(iGeomField), C.double(minX), C.double(minY), C.double(maxX), C.double(maxY))
}

func (l OGRLayer) SetSpatialFilterRectEx(iGeomField int, minX, minY, maxX, maxY float64) {
	ogrLSetSpatialFilterRectEx(l, iGeomField, minX, minY, maxX, maxY)
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

func (l OGRLayer) SetAttributeFilter(query string) (err error) {
	err = ogrError(ogrLSetAttributeFilter(l, query))
	return
}

func ogrLResetReading(l OGRLayer) {
	C.OGR_L_ResetReading(l.cValue)
}

func (l OGRLayer) ResetReading() {
	ogrLResetReading(l)
}

func ogrLGetNextFeature(l OGRLayer) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_L_GetNextFeature(l.cValue)}
	return
}

func (l OGRLayer) GetNextFeature() (result OGRFeature) {
	result = ogrLGetNextFeature(l)
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

func (l OGRLayer) SetNextByIndex(index int64) (err error) {
	err = ogrError(ogrLSetNextByIndex(l, index))
	return
}

func ogrLGetFeature(l OGRLayer, fid int64) (result OGRFeature) {
	result = OGRFeature{cValue: C.OGR_L_GetFeature(l.cValue, C.GIntBig(fid))}
	return
}

func (l OGRLayer) GetFeature(fid int64) (result OGRFeature, err error) {
	result = ogrLGetFeature(l, fid)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrLSetFeature(l OGRLayer, feat OGRFeature) (result OGRErr) {
	result = OGRErr(C.OGR_L_SetFeature(l.cValue, feat.cValue))
	return
}

func (l OGRLayer) SetFeature(feat OGRFeature) (err error) {
	err = ogrError(ogrLSetFeature(l, feat))
	return
}

func ogrLCreateFeature(l OGRLayer, feat OGRFeature) (result OGRErr) {
	result = OGRErr(C.OGR_L_CreateFeature(l.cValue, feat.cValue))
	return
}

func (l OGRLayer) CreateFeature(feat OGRFeature) (err error) {
	err = ogrError(ogrLCreateFeature(l, feat))
	return
}

func ogrLDeleteFeature(l OGRLayer, fid int64) (result OGRErr) {
	result = OGRErr(C.OGR_L_DeleteFeature(l.cValue, C.GIntBig(fid)))
	return
}

func (l OGRLayer) DeleteFeature(fid int64) (err error) {
	err = ogrError(ogrLDeleteFeature(l, fid))
	return
}

func ogrLUpsertFeature(l OGRLayer, feat OGRFeature) (result OGRErr) {
	result = OGRErr(C.OGR_L_UpsertFeature(l.cValue, feat.cValue))
	return
}

func (l OGRLayer) UpsertFeature(feat OGRFeature) (err error) {
	err = ogrError(ogrLUpsertFeature(l, feat))
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

func (l OGRLayer) UpdateFeature(feat OGRFeature, updatedFieldsIdx, updatedGeomFieldsIdx []int, updateStyleString bool) (err error) {
	err = ogrError(ogrLUpdateFeature(l, feat, updatedFieldsIdx, updatedGeomFieldsIdx, updateStyleString))
	return
}

func ogrLGetLayerDefn(l OGRLayer) (result OGRFeatureDefn) {
	result = OGRFeatureDefn{cValue: C.OGR_L_GetLayerDefn(l.cValue)}
	return
}

func (l OGRLayer) GetLayerDefn() (result OGRFeatureDefn) {
	result = ogrLGetLayerDefn(l)
	return
}

func ogrLGetSpatialRef(l OGRLayer) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OGR_L_GetSpatialRef(l.cValue)}
	return
}

func (l OGRLayer) GetSpatialRef() (result OGRSpatialReference) {
	result = ogrLGetSpatialRef(l)
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
	C.CPLFree(unsafe.Pointer(p))
	return
}

func (l OGRLayer) GetSupportedSRSList(iGeomField int) (result []OGRSpatialReference) {
	result = ogrLGetSupportedSRSList(l, iGeomField)
	return
}

func ogrLSetActiveSRS(l OGRLayer, iGeomField int, sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OGR_L_SetActiveSRS(l.cValue, C.int(iGeomField), sr.cValue))
	return
}

func (l OGRLayer) SetActiveSRS(iGeomField int, sr OGRSpatialReference) (err error) {
	err = ogrError(ogrLSetActiveSRS(l, iGeomField, sr))
	return
}

func ogrLFindFieldIndex(l OGRLayer, name string, exactMatch int) (result int) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_L_FindFieldIndex(l.cValue, cs, C.int(exactMatch)))
	return
}

func (l OGRLayer) FindFieldIndex(name string, exactMatch int) (result int) {
	result = ogrLFindFieldIndex(l, name, exactMatch)
	return
}

func ogrLGetFeatureCount(l OGRLayer, force int) (result int64) {
	result = int64(C.OGR_L_GetFeatureCount(l.cValue, C.int(force)))
	return
}

func (l OGRLayer) GetFeatureCount(force int) (result int64) {
	result = ogrLGetFeatureCount(l, force)
	return
}

func ogrLGetExtent(l OGRLayer, force int) (result OGREnvelope, status OGRErr) {
	result = OGREnvelope{cValue: new(C.OGREnvelope)}
	status = OGRErr(C.OGR_L_GetExtent(l.cValue, result.cValue, C.int(force)))
	return
}

func (l OGRLayer) GetExtent(force int) (result OGREnvelope, err error) {
	var status OGRErr
	result, status = ogrLGetExtent(l, force)
	err = ogrError(status)
	return
}

func ogrLGetExtentEx(l OGRLayer, iGeomField, force int) (result OGREnvelope, status OGRErr) {
	result = OGREnvelope{cValue: new(C.OGREnvelope)}
	status = OGRErr(C.OGR_L_GetExtentEx(l.cValue, C.int(iGeomField), result.cValue, C.int(force)))
	return
}

func (l OGRLayer) GetExtentEx(iGeomField, force int) (result OGREnvelope, err error) {
	var status OGRErr
	result, status = ogrLGetExtentEx(l, iGeomField, force)
	err = ogrError(status)
	return
}

func ogrLGetExtent3D(l OGRLayer, iGeomField, force int) (result OGREnvelope3D, status OGRErr) {
	result = OGREnvelope3D{cValue: new(C.OGREnvelope3D)}
	status = OGRErr(C.OGR_L_GetExtent3D(l.cValue, C.int(iGeomField), result.cValue, C.int(force)))
	return
}

func (l OGRLayer) GetExtent3D(iGeomField, force int) (result OGREnvelope3D, err error) {
	var status OGRErr
	result, status = ogrLGetExtent3D(l, iGeomField, force)
	err = ogrError(status)
	return
}

func ogrLTestCapability(l OGRLayer, capability string) (result int) {
	cs := C.CString(capability)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_L_TestCapability(l.cValue, cs))
	return
}

func (l OGRLayer) TestCapability(capability string) (result bool) {
	result = ogrLTestCapability(l, capability) != 0
	return
}

func ogrLCreateField(l OGRLayer, fld OGRFieldDefn, approxOK int) (result OGRErr) {
	result = OGRErr(C.OGR_L_CreateField(l.cValue, fld.cValue, C.int(approxOK)))
	return
}

func (l OGRLayer) CreateField(fld OGRFieldDefn, approxOK int) (err error) {
	err = ogrError(ogrLCreateField(l, fld, approxOK))
	return
}

func ogrLCreateGeomField(l OGRLayer, gfld OGRGeomFieldDefn, force int) (result OGRErr) {
	result = OGRErr(C.OGR_L_CreateGeomField(l.cValue, gfld.cValue, C.int(force)))
	return
}

func (l OGRLayer) CreateGeomField(gfld OGRGeomFieldDefn, force int) (err error) {
	err = ogrError(ogrLCreateGeomField(l, gfld, force))
	return
}

func ogrLDeleteField(l OGRLayer, field int) (result OGRErr) {
	result = OGRErr(C.OGR_L_DeleteField(l.cValue, C.int(field)))
	return
}

func (l OGRLayer) DeleteField(field int) (err error) {
	err = ogrError(ogrLDeleteField(l, field))
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

func (l OGRLayer) ReorderFields(panMap []int) (err error) {
	err = ogrError(ogrLReorderFields(l, panMap))
	return
}

func ogrLReorderField(l OGRLayer, oldFieldPos, newFieldPos int) (result OGRErr) {
	result = OGRErr(C.OGR_L_ReorderField(l.cValue, C.int(oldFieldPos), C.int(newFieldPos)))
	return
}

func (l OGRLayer) ReorderField(oldFieldPos, newFieldPos int) (err error) {
	err = ogrError(ogrLReorderField(l, oldFieldPos, newFieldPos))
	return
}

func ogrLAlterFieldDefn(l OGRLayer, field int, newFieldDefn OGRFieldDefn, flags int) (result OGRErr) {
	result = OGRErr(C.OGR_L_AlterFieldDefn(l.cValue, C.int(field), newFieldDefn.cValue, C.int(flags)))
	return
}

func (l OGRLayer) AlterFieldDefn(field int, newFieldDefn OGRFieldDefn, flags int) (err error) {
	err = ogrError(ogrLAlterFieldDefn(l, field, newFieldDefn, flags))
	return
}

func ogrLAlterGeomFieldDefn(l OGRLayer, field int, newGeomFieldDefn OGRGeomFieldDefn, flags int) (result OGRErr) {
	result = OGRErr(C.OGR_L_AlterGeomFieldDefn(l.cValue, C.int(field), newGeomFieldDefn.cValue, C.int(flags)))
	return
}

func (l OGRLayer) AlterGeomFieldDefn(field int, newGeomFieldDefn OGRGeomFieldDefn, flags int) (err error) {
	err = ogrError(ogrLAlterGeomFieldDefn(l, field, newGeomFieldDefn, flags))
	return
}

func ogrLStartTransaction(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_StartTransaction(l.cValue))
	return
}

func (l OGRLayer) StartTransaction() (err error) {
	err = ogrError(ogrLStartTransaction(l))
	return
}

func ogrLCommitTransaction(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_CommitTransaction(l.cValue))
	return
}

func (l OGRLayer) CommitTransaction() (err error) {
	err = ogrError(ogrLCommitTransaction(l))
	return
}

func ogrLRollbackTransaction(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_RollbackTransaction(l.cValue))
	return
}

func (l OGRLayer) RollbackTransaction() (err error) {
	err = ogrError(ogrLRollbackTransaction(l))
	return
}

func ogrLRename(l OGRLayer, newName string) (result OGRErr) {
	cs := C.CString(newName)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OGR_L_Rename(l.cValue, cs))
	return
}

func (l OGRLayer) Rename(newName string) (err error) {
	err = ogrError(ogrLRename(l, newName))
	return
}

// /*! @cond Doxygen_Suppress */

func ogrLReference(l OGRLayer) (result int) {
	result = int(C.OGR_L_Reference(l.cValue))
	return
}

func (l OGRLayer) Reference() (result int) {
	result = ogrLReference(l)
	return
}

func ogrLDereference(l OGRLayer) (result int) {
	result = int(C.OGR_L_Dereference(l.cValue))
	return
}

func (l OGRLayer) Dereference() (result int) {
	result = ogrLDereference(l)
	return
}

func ogrLGetRefCount(l OGRLayer) (result int) {
	result = int(C.OGR_L_GetRefCount(l.cValue))
	return
}

func (l OGRLayer) GetRefCount() (result int) {
	result = ogrLGetRefCount(l)
	return
}

// /*! @endcond */

func ogrLSyncToDisk(l OGRLayer) (result OGRErr) {
	result = OGRErr(C.OGR_L_SyncToDisk(l.cValue))
	return
}

func (l OGRLayer) SyncToDisk() (err error) {
	err = ogrError(ogrLSyncToDisk(l))
	return
}

// /*! @cond Doxygen_Suppress */

func ogrLGetFeaturesRead(l OGRLayer) (result int64) {
	result = int64(C.OGR_L_GetFeaturesRead(l.cValue))
	return
}

func (l OGRLayer) GetFeaturesRead() (result int64) {
	result = ogrLGetFeaturesRead(l)
	return
}

// /*! @endcond */

func ogrLGetFIDColumn(l OGRLayer) (result string) {
	result = C.GoString(C.OGR_L_GetFIDColumn(l.cValue))
	return
}

func (l OGRLayer) GetFIDColumn() (result string) {
	result = ogrLGetFIDColumn(l)
	return
}

func ogrLGetGeometryColumn(l OGRLayer) (result string) {
	result = C.GoString(C.OGR_L_GetGeometryColumn(l.cValue))
	return
}

func (l OGRLayer) GetGeometryColumn() (result string) {
	result = ogrLGetGeometryColumn(l)
	return
}

// /** Get style table */
func ogrLGetStyleTable(l OGRLayer) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_L_GetStyleTable(l.cValue)}
	return
}

// /** Get style table */
func (l OGRLayer) GetStyleTable() (result OGRStyleTable) {
	result = ogrLGetStyleTable(l)
	return
}

// /** Set style table (and take ownership) */
func ogrLSetStyleTableDirectly(l OGRLayer, styleTable OGRStyleTable) {
	C.OGR_L_SetStyleTableDirectly(l.cValue, styleTable.cValue)
}

// /** Set style table (and take ownership) */
func (l OGRLayer) SetStyleTableDirectly(styleTable OGRStyleTable) {
	ogrLSetStyleTableDirectly(l, styleTable)
}

// /** Set style table */
func ogrLSetStyleTable(l OGRLayer, styleTable OGRStyleTable) {
	C.OGR_L_SetStyleTable(l.cValue, styleTable.cValue)
}

// /** Set style table */
func (l OGRLayer) SetStyleTable(styleTable OGRStyleTable) {
	ogrLSetStyleTable(l, styleTable)
}

func ogrLSetIgnoredFields(l OGRLayer, fields []string) (result OGRErr) {
	cFields := make([]*C.char, len(fields)+1)
	for i, f := range fields {
		cFields[i] = C.CString(f)
		defer C.free(unsafe.Pointer(cFields[i]))
	}
	cFields[len(fields)] = nil
	result = OGRErr(C.OGR_L_SetIgnoredFields(l.cValue, &cFields[0]))
	return
}

func (l OGRLayer) SetIgnoredFields(fields []string) (err error) {
	err = ogrError(ogrLSetIgnoredFields(l, fields))
	return
}

func ogrLIntersection(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	status = OGRErr(C.OGR_L_Intersection(input.cValue, method.cValue, result.cValue, &cOptions[0], progress.cValue, progressData))
	return
}

func OGRLIntersection(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = ogrError(ogrLIntersection(input, method, result, options, progress, progressData))
	return
}

func ogrLUnion(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	status = OGRErr(C.OGR_L_Union(input.cValue, method.cValue, result.cValue, &cOptions[0], progress.cValue, progressData))
	return
}

func OGRLUnion(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = ogrError(ogrLUnion(input, method, result, options, progress, progressData))
	return
}

func ogrLSymDifference(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	status = OGRErr(C.OGR_L_SymDifference(input.cValue, method.cValue, result.cValue, &cOptions[0], progress.cValue, progressData))
	return
}

func OGRLSymDifference(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = ogrError(ogrLSymDifference(input, method, result, options, progress, progressData))
	return
}

func ogrLIdentity(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	status = OGRErr(C.OGR_L_Identity(input.cValue, method.cValue, result.cValue, &cOptions[0], progress.cValue, progressData))
	return
}

func OGRLIdentity(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = ogrError(ogrLIdentity(input, method, result, options, progress, progressData))
	return
}

func ogrLUpdate(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	status = OGRErr(C.OGR_L_Update(input.cValue, method.cValue, result.cValue, &cOptions[0], progress.cValue, progressData))
	return
}

func OGRLUpdate(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = ogrError(ogrLUpdate(input, method, result, options, progress, progressData))
	return
}

func ogrLClip(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	status = OGRErr(C.OGR_L_Clip(input.cValue, method.cValue, result.cValue, &cOptions[0], progress.cValue, progressData))
	return
}

func OGRLClip(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = ogrError(ogrLClip(input, method, result, options, progress, progressData))
	return
}

func ogrLErase(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (status OGRErr) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	status = OGRErr(C.OGR_L_Erase(input.cValue, method.cValue, result.cValue, &cOptions[0], progress.cValue, progressData))
	return
}

func OGRLErase(input, method, result OGRLayer, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = ogrError(ogrLErase(input, method, result, options, progress, progressData))
	return
}

// /* OGRDataSource */

func ogrDSDestroy(ds OGRDataSource) {
	C.OGR_DS_Destroy(ds.cValue)
}

func (ds OGRDataSource) Destroy() {
	ogrDSDestroy(ds)
}

func ogrDSGetName(ds OGRDataSource) (result string) {
	result = C.GoString(C.OGR_DS_GetName(ds.cValue))
	return
}

func (ds OGRDataSource) GetName() (result string) {
	result = ogrDSGetName(ds)
	return
}

func ogrDSGetLayerCount(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_GetLayerCount(ds.cValue))
	return
}

func (ds OGRDataSource) GetLayerCount() (result int) {
	result = ogrDSGetLayerCount(ds)
	return
}

func ogrDSGetLayer(ds OGRDataSource, layer int) (result OGRLayer) {
	result = OGRLayer{cValue: C.OGR_DS_GetLayer(ds.cValue, C.int(layer))}
	return
}

func (ds OGRDataSource) GetLayer(layer int) (result OGRLayer, err error) {
	result = ogrDSGetLayer(ds, layer)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDSGetLayerByName(ds OGRDataSource, name string) (result OGRLayer) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRLayer{cValue: C.OGR_DS_GetLayerByName(ds.cValue, cs)}
	return
}

func (ds OGRDataSource) GetLayerByName(name string) (result OGRLayer, err error) {
	result = ogrDSGetLayerByName(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDSDeleteLayer(ds OGRDataSource, layer int) (result OGRErr) {
	result = OGRErr(C.OGR_DS_DeleteLayer(ds.cValue, C.int(layer)))
	return
}

func (ds OGRDataSource) DeleteLayer(layer int) (err error) {
	err = ogrError(ogrDSDeleteLayer(ds, layer))
	return
}

func ogrDSGetDriver(ds OGRDataSource) (result OGRSFDriver) {
	result = OGRSFDriver{cValue: C.OGR_DS_GetDriver(ds.cValue)}
	return
}

func (ds OGRDataSource) GetDriver() (result OGRSFDriver, err error) {
	result = ogrDSGetDriver(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDSCreateLayer(ds OGRDataSource, name string, sr OGRSpatialReference, geomType OGRwkbGeometryType, options []string) (result OGRLayer) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRLayer{cValue: C.OGR_DS_CreateLayer(ds.cValue, cs, sr.cValue, C.OGRwkbGeometryType(geomType), &cOptions[0])}
	return
}

func (ds OGRDataSource) CreateLayer(name string, sr OGRSpatialReference, geomType OGRwkbGeometryType, options []string) (result OGRLayer, err error) {
	result = ogrDSCreateLayer(ds, name, sr, geomType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDSCopyLayer(ds OGRDataSource, srcLayer OGRLayer, newName string, options []string) (result OGRLayer) {
	cs := C.CString(newName)
	defer C.free(unsafe.Pointer(cs))
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRLayer{cValue: C.OGR_DS_CopyLayer(ds.cValue, srcLayer.cValue, cs, &cOptions[0])}
	return
}

func (ds OGRDataSource) CopyLayer(srcLayer OGRLayer, newName string, options []string) (result OGRLayer, err error) {
	result = ogrDSCopyLayer(ds, srcLayer, newName, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDSTestCapability(ds OGRDataSource, capability string) (result int) {
	cs := C.CString(capability)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_DS_TestCapability(ds.cValue, cs))
	return
}

func (ds OGRDataSource) TestCapability(capability string) (result bool) {
	result = ogrDSTestCapability(ds, capability) != 0
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

func (ds OGRDataSource) ExecuteSQL(statement string, spatialFilter OGRGeometry, dialect string) (result OGRLayer, err error) {
	result = ogrDSExecuteSQL(ds, statement, spatialFilter, dialect)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDSReleaseResultSet(ds OGRDataSource, layer OGRLayer) {
	C.OGR_DS_ReleaseResultSet(ds.cValue, layer.cValue)
}

func (ds OGRDataSource) ReleaseResultSet(layer OGRLayer) {
	ogrDSReleaseResultSet(ds, layer)
}

// /*! @cond Doxygen_Suppress */

func ogrDSReference(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_Reference(ds.cValue))
	return
}

func (ds OGRDataSource) Reference() (result int) {
	result = ogrDSReference(ds)
	return
}

func ogrDSDereference(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_Dereference(ds.cValue))
	return
}

func (ds OGRDataSource) Dereference() (result int) {
	result = ogrDSDereference(ds)
	return
}

func ogrDSGetRefCount(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_GetRefCount(ds.cValue))
	return
}

func (ds OGRDataSource) GetRefCount() (result int) {
	result = ogrDSGetRefCount(ds)
	return
}

func ogrDSGetSummaryRefCount(ds OGRDataSource) (result int) {
	result = int(C.OGR_DS_GetSummaryRefCount(ds.cValue))
	return
}

func (ds OGRDataSource) GetSummaryRefCount() (result int) {
	result = ogrDSGetSummaryRefCount(ds)
	return
}

// /*! @endcond */

// /** Flush pending changes to disk. See GDALDataset::FlushCache() */
func ogrDSSyncToDisk(ds OGRDataSource) (result OGRErr) {
	result = OGRErr(C.OGR_DS_SyncToDisk(ds.cValue))
	return
}

// /** Flush pending changes to disk. See GDALDataset::FlushCache() */
func (ds OGRDataSource) SyncToDisk() (err error) {
	err = ogrError(ogrDSSyncToDisk(ds))
	return
}

// /** Get style table */
func ogrDSGetStyleTable(ds OGRDataSource) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_DS_GetStyleTable(ds.cValue)}
	return
}

// /** Get style table */
func (ds OGRDataSource) GetStyleTable() (result OGRStyleTable) {
	result = ogrDSGetStyleTable(ds)
	return
}

// /** Set style table (and take ownership) */
func ogrDSSetStyleTableDirectly(ds OGRDataSource, styleTable OGRStyleTable) {
	C.OGR_DS_SetStyleTableDirectly(ds.cValue, styleTable.cValue)
}

// /** Set style table (and take ownership) */
func (ds OGRDataSource) SetStyleTableDirectly(styleTable OGRStyleTable) {
	ogrDSSetStyleTableDirectly(ds, styleTable)
}

// /** Set style table */
func ogrDSSetStyleTable(ds OGRDataSource, styleTable OGRStyleTable) {
	C.OGR_DS_SetStyleTable(ds.cValue, styleTable.cValue)
}

// /** Set style table */
func (ds OGRDataSource) SetStyleTable(styleTable OGRStyleTable) {
	ogrDSSetStyleTable(ds, styleTable)
}

// /* OGRSFDriver */

func ogrDrGetName(dr OGRSFDriver) (result string) {
	result = C.GoString(C.OGR_Dr_GetName(dr.cValue))
	return
}

func (dr OGRSFDriver) GetName() (result string) {
	result = ogrDrGetName(dr)
	return
}

func ogrDrOpen(dr OGRSFDriver, name string, update int) (result OGRDataSource) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRDataSource{cValue: C.OGR_Dr_Open(dr.cValue, cs, C.int(update))}
	return
}

func (dr OGRSFDriver) Open(name string, update int) (result OGRDataSource, err error) {
	result = ogrDrOpen(dr, name, update)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDrTestCapability(dr OGRSFDriver, capability string) (result int) {
	cs := C.CString(capability)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_Dr_TestCapability(dr.cValue, cs))
	return
}

func (dr OGRSFDriver) TestCapability(capability string) (result bool) {
	result = ogrDrTestCapability(dr, capability) != 0
	return
}

func ogrDrCreateDataSource(dr OGRSFDriver, name string, options []string) (result OGRDataSource) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRDataSource{cValue: C.OGR_Dr_CreateDataSource(dr.cValue, cs, &cOptions[0])}
	return
}

func (dr OGRSFDriver) CreateDataSource(name string, options []string) (result OGRDataSource, err error) {
	result = ogrDrCreateDataSource(dr, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDrCopyDataSource(dr OGRSFDriver, srcDS OGRDataSource, newName string, options []string) (result OGRDataSource) {
	cs := C.CString(newName)
	defer C.free(unsafe.Pointer(cs))
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRDataSource{cValue: C.OGR_Dr_CopyDataSource(dr.cValue, srcDS.cValue, cs, &cOptions[0])}
	return
}

func (dr OGRSFDriver) CopyDataSource(srcDS OGRDataSource, newName string, options []string) (result OGRDataSource, err error) {
	result = ogrDrCopyDataSource(dr, srcDS, newName, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrDrDeleteDataSource(dr OGRSFDriver, name string) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OGR_Dr_DeleteDataSource(dr.cValue, cs))
	return
}

func (dr OGRSFDriver) DeleteDataSource(name string) (err error) {
	err = ogrError(ogrDrDeleteDataSource(dr, name))
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

func OGROpen(name string, update int) (result OGRDataSource, driver OGRSFDriver, err error) {
	result, driver = ogrOpen(name, update)
	if result.cValue == nil {
		err = lastError()
	}
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

func OGROpenShared(name string, update int) (result OGRDataSource, driver OGRSFDriver, err error) {
	result, driver = ogrOpenShared(name, update)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrReleaseDataSource(ds OGRDataSource) (result OGRErr) {
	result = OGRErr(C.OGRReleaseDataSource(ds.cValue))
	return
}

func (ds OGRDataSource) Release() (err error) {
	err = ogrError(ogrReleaseDataSource(ds))
	return
}

// /*! @cond Doxygen_Suppress */

func ogrRegisterDriver(dr OGRSFDriver) {
	C.OGRRegisterDriver(dr.cValue)
}

func (dr OGRSFDriver) Register() {
	ogrRegisterDriver(dr)
}

func ogrDeregisterDriver(dr OGRSFDriver) {
	C.OGRDeregisterDriver(dr.cValue)
}

func (dr OGRSFDriver) Deregister() {
	ogrDeregisterDriver(dr)
}

// /*! @endcond */

func ogrGetDriverCount() (result int) {
	result = int(C.OGRGetDriverCount())
	return
}

func OGRGetDriverCount() (result int) {
	result = ogrGetDriverCount()
	return
}

func ogrGetDriver(driver int) (result OGRSFDriver) {
	result = OGRSFDriver{cValue: C.OGRGetDriver(C.int(driver))}
	return
}

func OGRGetDriver(driver int) (result OGRSFDriver, err error) {
	result = ogrGetDriver(driver)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrGetDriverByName(name string) (result OGRSFDriver) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRSFDriver{cValue: C.OGRGetDriverByName(cs)}
	return
}

func OGRGetDriverByName(name string) (result OGRSFDriver, err error) {
	result = ogrGetDriverByName(name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// /*! @cond Doxygen_Suppress */

func ogrGetOpenDSCount() (result int) {
	result = int(C.OGRGetOpenDSCount())
	return
}

func OGRGetOpenDSCount() (result int) {
	result = ogrGetOpenDSCount()
	return
}

func ogrGetOpenDS(ds int) (result OGRDataSource) {
	result = OGRDataSource{cValue: C.OGRGetOpenDS(C.int(ds))}
	return
}

func OGRGetOpenDS(ds int) (result OGRDataSource, err error) {
	result = ogrGetOpenDS(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// /*! @endcond */

func ogrRegisterAll() {
	C.OGRRegisterAll()
}

func OGRRegisterAll() {
	ogrRegisterAll()
}

// /** Clean-up all drivers, including raster ones.
//   - See GDALDestroyDriverManager() */
func ogrCleanupAll() {
	C.OGRCleanupAll()
}

// /** Clean-up all drivers, including raster ones.
//   - See GDALDestroyDriverManager() */
func OGRCleanupAll() {
	ogrCleanupAll()
}

// /* -------------------------------------------------------------------- */
// /*      ogrsf_featurestyle.h                                            */
// /* -------------------------------------------------------------------- */

// /* OGRStyleMgr */

func ogrSMCreate(styleTable OGRStyleTable) (result OGRStyleMgr) {
	result = OGRStyleMgr{cValue: C.OGR_SM_Create(styleTable.cValue)}
	return
}

func OGRSMCreate(styleTable OGRStyleTable) (result OGRStyleMgr, err error) {
	result = ogrSMCreate(styleTable)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrSMDestroy(sm OGRStyleMgr) {
	C.OGR_SM_Destroy(sm.cValue)
}

func (sm OGRStyleMgr) Destroy() {
	ogrSMDestroy(sm)
}

func ogrSMInitFromFeature(sm OGRStyleMgr, feat OGRFeature) (result string) {
	result = C.GoString(C.OGR_SM_InitFromFeature(sm.cValue, feat.cValue))
	return
}

func (sm OGRStyleMgr) InitFromFeature(feat OGRFeature) (result string) {
	result = ogrSMInitFromFeature(sm, feat)
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

func (sm OGRStyleMgr) InitStyleString(styleString string) (result bool) {
	result = ogrSMInitStyleString(sm, styleString) != 0
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

func (sm OGRStyleMgr) GetPartCount(styleString string) (result int) {
	result = ogrSMGetPartCount(sm, styleString)
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

func (sm OGRStyleMgr) GetPart(partId int, styleString string) (result OGRStyleTool) {
	result = ogrSMGetPart(sm, partId, styleString)
	return
}

func ogrSMAddPart(sm OGRStyleMgr, st OGRStyleTool) (result int) {
	result = int(C.OGR_SM_AddPart(sm.cValue, st.cValue))
	return
}

func (sm OGRStyleMgr) AddPart(st OGRStyleTool) (result bool) {
	result = ogrSMAddPart(sm, st) != 0
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

func (sm OGRStyleMgr) AddStyle(styleName, styleString string) (result bool) {
	result = ogrSMAddStyle(sm, styleName, styleString) != 0
	return
}

// /* OGRStyleTool */

func ogrSTCreate(classId OGRSTClassId) (result OGRStyleTool) {
	result = OGRStyleTool{cValue: C.OGR_ST_Create(C.OGRSTClassId(classId))}
	return
}

func OGRSTCreate(classId OGRSTClassId) (result OGRStyleTool, err error) {
	result = ogrSTCreate(classId)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrSTDestroy(st OGRStyleTool) {
	C.OGR_ST_Destroy(st.cValue)
}

func (st OGRStyleTool) Destroy() {
	ogrSTDestroy(st)
}

func ogrSTGetType(st OGRStyleTool) (result OGRSTClassId) {
	result = OGRSTClassId(C.OGR_ST_GetType(st.cValue))
	return
}

func (st OGRStyleTool) GetType() (result OGRSTClassId) {
	result = ogrSTGetType(st)
	return
}

func ogrSTGetUnit(st OGRStyleTool) (result OGRSTUnitId) {
	result = OGRSTUnitId(C.OGR_ST_GetUnit(st.cValue))
	return
}

func (st OGRStyleTool) GetUnit() (result OGRSTUnitId) {
	result = ogrSTGetUnit(st)
	return
}

func ogrSTSetUnit(st OGRStyleTool, unit OGRSTUnitId, groundPaperScale float64) {
	C.OGR_ST_SetUnit(st.cValue, C.OGRSTUnitId(unit), C.double(groundPaperScale))
}

func (st OGRStyleTool) SetUnit(unit OGRSTUnitId, groundPaperScale float64) {
	ogrSTSetUnit(st, unit, groundPaperScale)
}

func ogrSTGetParamStr(st OGRStyleTool, param int) (result string, isNull bool) {
	var vNull C.int
	result = C.GoString(C.OGR_ST_GetParamStr(st.cValue, C.int(param), &vNull))
	isNull = vNull != 0
	return
}

func (st OGRStyleTool) GetParamStr(param int) (result string, isNull bool) {
	result, isNull = ogrSTGetParamStr(st, param)
	return
}

func ogrSTGetParamNum(st OGRStyleTool, param int) (result int, isNull bool) {
	var vNull C.int
	result = int(C.OGR_ST_GetParamNum(st.cValue, C.int(param), &vNull))
	isNull = vNull != 0
	return
}

func (st OGRStyleTool) GetParamNum(param int) (result int, isNull bool) {
	result, isNull = ogrSTGetParamNum(st, param)
	return
}

func ogrSTGetParamDbl(st OGRStyleTool, param int) (result float64, isNull bool) {
	var vNull C.int
	result = float64(C.OGR_ST_GetParamDbl(st.cValue, C.int(param), &vNull))
	isNull = vNull != 0
	return
}

func (st OGRStyleTool) GetParamDbl(param int) (result float64, isNull bool) {
	result, isNull = ogrSTGetParamDbl(st, param)
	return
}

func ogrSTSetParamStr(st OGRStyleTool, param int, value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))
	C.OGR_ST_SetParamStr(st.cValue, C.int(param), cs)
}

func (st OGRStyleTool) SetParamStr(param int, value string) {
	ogrSTSetParamStr(st, param, value)
}

func ogrSTSetParamNum(st OGRStyleTool, param, value int) {
	C.OGR_ST_SetParamNum(st.cValue, C.int(param), C.int(value))
}

func (st OGRStyleTool) SetParamNum(param, value int) {
	ogrSTSetParamNum(st, param, value)
}

func ogrSTSetParamDbl(st OGRStyleTool, param int, value float64) {
	C.OGR_ST_SetParamDbl(st.cValue, C.int(param), C.double(value))
}

func (st OGRStyleTool) SetParamDbl(param int, value float64) {
	ogrSTSetParamDbl(st, param, value)
}

func ogrSTGetStyleString(st OGRStyleTool) (result string) {
	result = C.GoString(C.OGR_ST_GetStyleString(st.cValue))
	return
}

func (st OGRStyleTool) GetStyleString() (result string) {
	result = ogrSTGetStyleString(st)
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

func (st OGRStyleTool) GetRGBFromString(color string) (red, green, blue, alpha int, ok bool) {
	red, green, blue, alpha, ok = ogrSTGetRGBFromString(st, color)
	return
}

// /* OGRStyleTable */

func ogrSTBLCreate() (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.OGR_STBL_Create()}
	return
}

func OGRSTBLCreate() (result OGRStyleTable, err error) {
	result = ogrSTBLCreate()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func ogrSTBLDestroy(stbl OGRStyleTable) {
	C.OGR_STBL_Destroy(stbl.cValue)
}

func (stbl OGRStyleTable) Destroy() {
	ogrSTBLDestroy(stbl)
}

func ogrSTBLAddStyle(stbl OGRStyleTable, name, styleString string) (result int) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csStyle := C.CString(styleString)
	defer C.free(unsafe.Pointer(csStyle))
	result = int(C.OGR_STBL_AddStyle(stbl.cValue, csName, csStyle))
	return
}

func (stbl OGRStyleTable) AddStyle(name, styleString string) (result bool) {
	result = ogrSTBLAddStyle(stbl, name, styleString) != 0
	return
}

func ogrSTBLSaveStyleTable(stbl OGRStyleTable, filename string) (result int) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_STBL_SaveStyleTable(stbl.cValue, cs))
	return
}

func (stbl OGRStyleTable) SaveStyleTable(filename string) (result bool) {
	result = ogrSTBLSaveStyleTable(stbl, filename) != 0
	return
}

func ogrSTBLLoadStyleTable(stbl OGRStyleTable, filename string) (result int) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGR_STBL_LoadStyleTable(stbl.cValue, cs))
	return
}

func (stbl OGRStyleTable) LoadStyleTable(filename string) (result bool) {
	result = ogrSTBLLoadStyleTable(stbl, filename) != 0
	return
}

func ogrSTBLFind(stbl OGRStyleTable, name string) (result string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OGR_STBL_Find(stbl.cValue, cs))
	return
}

func (stbl OGRStyleTable) Find(name string) (result string) {
	result = ogrSTBLFind(stbl, name)
	return
}

func ogrSTBLResetStyleStringReading(stbl OGRStyleTable) {
	C.OGR_STBL_ResetStyleStringReading(stbl.cValue)
}

func (stbl OGRStyleTable) ResetStyleStringReading() {
	ogrSTBLResetStyleStringReading(stbl)
}

func ogrSTBLGetNextStyle(stbl OGRStyleTable) (result string) {
	result = C.GoString(C.OGR_STBL_GetNextStyle(stbl.cValue))
	return
}

func (stbl OGRStyleTable) GetNextStyle() (result string) {
	result = ogrSTBLGetNextStyle(stbl)
	return
}

func ogrSTBLGetLastStyleName(stbl OGRStyleTable) (result string) {
	result = C.GoString(C.OGR_STBL_GetLastStyleName(stbl.cValue))
	return
}

func (stbl OGRStyleTable) GetLastStyleName() (result string) {
	result = ogrSTBLGetLastStyleName(stbl)
	return
}

// CPL_C_END
