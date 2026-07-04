package gdal

/*
#include "ogr_core.h"

const char* const _OLCRandomRead                        = OLCRandomRead;
const char* const _OLCSequentialWrite                   = OLCSequentialWrite;
const char* const _OLCRandomWrite                       = OLCRandomWrite;
const char* const _OLCFastSpatialFilter                 = OLCFastSpatialFilter;
const char* const _OLCFastFeatureCount                  = OLCFastFeatureCount;
const char* const _OLCFastGetExtent                     = OLCFastGetExtent;
const char* const _OLCFastGetExtent3D                   = OLCFastGetExtent3D;
const char* const _OLCCreateField                       = OLCCreateField;
const char* const _OLCDeleteField                       = OLCDeleteField;
const char* const _OLCReorderFields                     = OLCReorderFields;
const char* const _OLCAlterFieldDefn                    = OLCAlterFieldDefn;
const char* const _OLCAlterGeomFieldDefn                = OLCAlterGeomFieldDefn;
const char* const _OLCTransactions                      = OLCTransactions;
const char* const _OLCDeleteFeature                     = OLCDeleteFeature;
const char* const _OLCUpsertFeature                     = OLCUpsertFeature;
const char* const _OLCUpdateFeature                     = OLCUpdateFeature;
const char* const _OLCFastSetNextByIndex                = OLCFastSetNextByIndex;
const char* const _OLCStringsAsUTF8                     = OLCStringsAsUTF8;
const char* const _OLCIgnoreFields                      = OLCIgnoreFields;
const char* const _OLCCreateGeomField                   = OLCCreateGeomField;
const char* const _OLCCurveGeometries                   = OLCCurveGeometries;
const char* const _OLCMeasuredGeometries                = OLCMeasuredGeometries;
const char* const _OLCZGeometries                       = OLCZGeometries;
const char* const _OLCRename                            = OLCRename;
const char* const _OLCFastGetArrowStream                = OLCFastGetArrowStream;
const char* const _OLCFastWriteArrowBatch               = OLCFastWriteArrowBatch;

const char* const _ODsCCreateLayer                      = ODsCCreateLayer;
const char* const _ODsCDeleteLayer                      = ODsCDeleteLayer;
const char* const _ODsCCreateGeomFieldAfterCreateLayer  = ODsCCreateGeomFieldAfterCreateLayer;
const char* const _ODsCCurveGeometries                  = ODsCCurveGeometries;
const char* const _ODsCTransactions                     = ODsCTransactions;
const char* const _ODsCEmulatedTransactions             = ODsCEmulatedTransactions;
const char* const _ODsCMeasuredGeometries               = ODsCMeasuredGeometries;
const char* const _ODsCZGeometries                      = ODsCZGeometries;
const char* const _ODsCRandomLayerRead                  = ODsCRandomLayerRead;
const char* const _ODsCRandomLayerWrite                 = ODsCRandomLayerWrite;
const char* const _ODsCAddFieldDomain                   = ODsCAddFieldDomain;
const char* const _ODsCDeleteFieldDomain                = ODsCDeleteFieldDomain;
const char* const _ODsCUpdateFieldDomain                = ODsCUpdateFieldDomain;

const char* const _ODrCCreateDataSource                 = ODrCCreateDataSource;
const char* const _ODrCDeleteDataSource                 = ODrCDeleteDataSource;

const char* const _OLMD_FID64                           = OLMD_FID64;

const int _ALTER_ALL_FLAG                 = ALTER_ALL_FLAG;
const int _ALTER_GEOM_FIELD_DEFN_ALL_FLAG = ALTER_GEOM_FIELD_DEFN_ALL_FLAG;
const int _OGR_F_VAL_ALL                  = OGR_F_VAL_ALL;
*/
import "C"

import "unsafe"

// #include "cpl_port.h"
// #include "gdal_version.h"

/**
 * \file
 *
 * Core portability services for cross-platform OGR code.
 */

type OGREnvelope struct {
	cValue *C.OGREnvelope
}

type OGREnvelope3D struct {
	cValue *C.OGREnvelope3D
}

// CPL_C_START

// /*! @cond Doxygen_Suppress */
// void CPL_DLL *OGRMalloc(size_t) CPL_WARN_DEPRECATED("Use CPLMalloc instead.");
// void CPL_DLL *OGRCalloc(size_t, size_t)
//     CPL_WARN_DEPRECATED("Use CPLCalloc instead.");
// void CPL_DLL *OGRRealloc(void *, size_t)
//     CPL_WARN_DEPRECATED("Use CPLRealloc instead.");
// char CPL_DLL *OGRStrdup(const char *)
//     CPL_WARN_DEPRECATED("Use CPLStrdup instead.");
// void CPL_DLL OGRFree(void *) CPL_WARN_DEPRECATED("Use CPLFree instead.");
// /*! @endcond */

type OGRErr C.OGRErr

const (
	OGRErrNone                    OGRErr = C.OGRERR_NONE
	OGRErrNotEnoughData           OGRErr = C.OGRERR_NOT_ENOUGH_DATA
	OGRErrNotEnoughMemory         OGRErr = C.OGRERR_NOT_ENOUGH_MEMORY
	OGRErrUnsupportedGeometryType OGRErr = C.OGRERR_UNSUPPORTED_GEOMETRY_TYPE
	OGRErrUnsupportedOperation    OGRErr = C.OGRERR_UNSUPPORTED_OPERATION
	OGRErrCorruptData             OGRErr = C.OGRERR_CORRUPT_DATA
	OGRErrFailure                 OGRErr = C.OGRERR_FAILURE
	OGRErrUnsupportedSRS          OGRErr = C.OGRERR_UNSUPPORTED_SRS
	OGRErrInvalidHandle           OGRErr = C.OGRERR_INVALID_HANDLE
	OGRErrNonExistingFeature      OGRErr = C.OGRERR_NON_EXISTING_FEATURE
)

func (e OGRErr) String() string {
	switch e {
	case OGRErrNone:
		return "OGRERR_NONE"
	case OGRErrNotEnoughData:
		return "OGRERR_NOT_ENOUGH_DATA"
	case OGRErrNotEnoughMemory:
		return "OGRERR_NOT_ENOUGH_MEMORY"
	case OGRErrUnsupportedGeometryType:
		return "OGRERR_UNSUPPORTED_GEOMETRY_TYPE"
	case OGRErrUnsupportedOperation:
		return "OGRERR_UNSUPPORTED_OPERATION"
	case OGRErrCorruptData:
		return "OGRERR_CORRUPT_DATA"
	case OGRErrFailure:
		return "OGRERR_FAILURE"
	case OGRErrUnsupportedSRS:
		return "OGRERR_UNSUPPORTED_SRS"
	case OGRErrInvalidHandle:
		return "OGRERR_INVALID_HANDLE"
	case OGRErrNonExistingFeature:
		return "OGRERR_NON_EXISTING_FEATURE"
	default:
		return "OGRERR_UNKNOWN"
	}
}

type OGRBoolean = C.OGRBoolean

/* -------------------------------------------------------------------- */
/*      ogr_geometry.h related definitions.                             */
/* -------------------------------------------------------------------- */

type OGRwkbGeometryType C.OGRwkbGeometryType

const (
	WkbUnknown            OGRwkbGeometryType = C.wkbUnknown
	WkbPoint              OGRwkbGeometryType = C.wkbPoint
	WkbLineString         OGRwkbGeometryType = C.wkbLineString
	WkbPolygon            OGRwkbGeometryType = C.wkbPolygon
	WkbMultiPoint         OGRwkbGeometryType = C.wkbMultiPoint
	WkbMultiLineString    OGRwkbGeometryType = C.wkbMultiLineString
	WkbMultiPolygon       OGRwkbGeometryType = C.wkbMultiPolygon
	WkbGeometryCollection OGRwkbGeometryType = C.wkbGeometryCollection

	WkbCircularString    OGRwkbGeometryType = C.wkbCircularString
	WkbCompoundCurve     OGRwkbGeometryType = C.wkbCompoundCurve
	WkbCurvePolygon      OGRwkbGeometryType = C.wkbCurvePolygon
	WkbMultiCurve        OGRwkbGeometryType = C.wkbMultiCurve
	WkbMultiSurface      OGRwkbGeometryType = C.wkbMultiSurface
	WkbCurve             OGRwkbGeometryType = C.wkbCurve
	WkbSurface           OGRwkbGeometryType = C.wkbSurface
	WkbPolyhedralSurface OGRwkbGeometryType = C.wkbPolyhedralSurface
	WkbTIN               OGRwkbGeometryType = C.wkbTIN
	WkbTriangle          OGRwkbGeometryType = C.wkbTriangle

	WkbNone       OGRwkbGeometryType = C.wkbNone
	WkbLinearRing OGRwkbGeometryType = C.wkbLinearRing

	WkbCircularStringZ    OGRwkbGeometryType = C.wkbCircularStringZ
	WkbCompoundCurveZ     OGRwkbGeometryType = C.wkbCompoundCurveZ
	WkbCurvePolygonZ      OGRwkbGeometryType = C.wkbCurvePolygonZ
	WkbMultiCurveZ        OGRwkbGeometryType = C.wkbMultiCurveZ
	WkbMultiSurfaceZ      OGRwkbGeometryType = C.wkbMultiSurfaceZ
	WkbCurveZ             OGRwkbGeometryType = C.wkbCurveZ
	WkbSurfaceZ           OGRwkbGeometryType = C.wkbSurfaceZ
	WkbPolyhedralSurfaceZ OGRwkbGeometryType = C.wkbPolyhedralSurfaceZ
	WkbTINZ               OGRwkbGeometryType = C.wkbTINZ
	WkbTriangleZ          OGRwkbGeometryType = C.wkbTriangleZ

	WkbPointM              OGRwkbGeometryType = C.wkbPointM
	WkbLineStringM         OGRwkbGeometryType = C.wkbLineStringM
	WkbPolygonM            OGRwkbGeometryType = C.wkbPolygonM
	WkbMultiPointM         OGRwkbGeometryType = C.wkbMultiPointM
	WkbMultiLineStringM    OGRwkbGeometryType = C.wkbMultiLineStringM
	WkbMultiPolygonM       OGRwkbGeometryType = C.wkbMultiPolygonM
	WkbGeometryCollectionM OGRwkbGeometryType = C.wkbGeometryCollectionM
	WkbCircularStringM     OGRwkbGeometryType = C.wkbCircularStringM
	WkbCompoundCurveM      OGRwkbGeometryType = C.wkbCompoundCurveM
	WkbCurvePolygonM       OGRwkbGeometryType = C.wkbCurvePolygonM
	WkbMultiCurveM         OGRwkbGeometryType = C.wkbMultiCurveM
	WkbMultiSurfaceM       OGRwkbGeometryType = C.wkbMultiSurfaceM
	WkbCurveM              OGRwkbGeometryType = C.wkbCurveM
	WkbSurfaceM            OGRwkbGeometryType = C.wkbSurfaceM
	WkbPolyhedralSurfaceM  OGRwkbGeometryType = C.wkbPolyhedralSurfaceM
	WkbTINM                OGRwkbGeometryType = C.wkbTINM
	WkbTriangleM           OGRwkbGeometryType = C.wkbTriangleM

	WkbPointZM              OGRwkbGeometryType = C.wkbPointZM
	WkbLineStringZM         OGRwkbGeometryType = C.wkbLineStringZM
	WkbPolygonZM            OGRwkbGeometryType = C.wkbPolygonZM
	WkbMultiPointZM         OGRwkbGeometryType = C.wkbMultiPointZM
	WkbMultiLineStringZM    OGRwkbGeometryType = C.wkbMultiLineStringZM
	WkbMultiPolygonZM       OGRwkbGeometryType = C.wkbMultiPolygonZM
	WkbGeometryCollectionZM OGRwkbGeometryType = C.wkbGeometryCollectionZM
	WkbCircularStringZM     OGRwkbGeometryType = C.wkbCircularStringZM
	WkbCompoundCurveZM      OGRwkbGeometryType = C.wkbCompoundCurveZM
	WkbCurvePolygonZM       OGRwkbGeometryType = C.wkbCurvePolygonZM
	WkbMultiCurveZM         OGRwkbGeometryType = C.wkbMultiCurveZM
	WkbMultiSurfaceZM       OGRwkbGeometryType = C.wkbMultiSurfaceZM
	WkbCurveZM              OGRwkbGeometryType = C.wkbCurveZM
	WkbSurfaceZM            OGRwkbGeometryType = C.wkbSurfaceZM
	WkbPolyhedralSurfaceZM  OGRwkbGeometryType = C.wkbPolyhedralSurfaceZM
	WkbTINZM                OGRwkbGeometryType = C.wkbTINZM
	WkbTriangleZM           OGRwkbGeometryType = C.wkbTriangleZM

	WkbPoint25D              OGRwkbGeometryType = C.wkbPoint25D
	WkbLineString25D         OGRwkbGeometryType = C.wkbLineString25D
	WkbPolygon25D            OGRwkbGeometryType = C.wkbPolygon25D
	WkbMultiPoint25D         OGRwkbGeometryType = C.wkbMultiPoint25D
	WkbMultiLineString25D    OGRwkbGeometryType = C.wkbMultiLineString25D
	WkbMultiPolygon25D       OGRwkbGeometryType = C.wkbMultiPolygon25D
	WkbGeometryCollection25D OGRwkbGeometryType = C.wkbGeometryCollection25D
)

/* clang-format off */
/**
 * Output variants of WKB we support.
 *
 * 99-402 was a short-lived extension to SFSQL 1.1 that used a high-bit flag
 * to indicate the presence of Z coordinates in a WKB geometry.
 *
 * SQL/MM Part 3 and SFSQL 1.2 use offsets of 1000 (Z), 2000 (M) and 3000 (ZM)
 * to indicate the present of higher dimensional coordinates in a WKB geometry.
 * Reference: <a href="https://portal.opengeospatial.org/files/?artifact_id=320243">
 * 09-009_Committee_Draft_ISOIEC_CD_13249-3_SQLMM_Spatial.pdf</a>,
 * ISO/IEC JTC 1/SC 32 N 1820, ISO/IEC CD 13249-3:201x(E), Date: 2009-01-16.
 * The codes are also found in §8.2.3 of <a href="http://portal.opengeospatial.org/files/?artifact_id=25355"> OGC
 * 06-103r4 "OpenGIS® Implementation Standard for Geographic information -
 * Simple feature access - Part 1: Common architecture", v1.2.1</a>
 */
/* clang-format on */

type OGRwkbVariant C.OGRwkbVariant

const (
	WkbVariantOldOgc   OGRwkbVariant = C.wkbVariantOldOgc
	WkbVariantIso      OGRwkbVariant = C.wkbVariantIso
	WkbVariantPostGIS1 OGRwkbVariant = C.wkbVariantPostGIS1
)

// #ifndef GDAL_COMPILATION
// /** @deprecated Use wkbHasZ() or wkbSetZ() instead */
// #define wkb25DBit 0x80000000
// #endif

func wkbFlatten(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = ogrGTFlatten(x)
	return
}

func WkbFlatten(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = wkbFlatten(x)
	return
}

func wkbHasZ(x OGRwkbGeometryType) bool {
	return ogrGTHasZ(x) != 0
}

func WkbHasZ(x OGRwkbGeometryType) bool {
	return wkbHasZ(x)
}

func wkbSetZ(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = ogrGTSetZ(x)
	return
}

func WkbSetZ(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = wkbSetZ(x)
	return
}

func wkbHasM(x OGRwkbGeometryType) bool {
	return ogrGTHasM(x) != 0
}

func WkbHasM(x OGRwkbGeometryType) bool {
	return wkbHasM(x)
}

func wkbSetM(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = ogrGTSetM(x)
	return
}

func WkbSetM(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = wkbSetM(x)
	return
}

const OGRZMarker = C.ogrZMarker

func ogrGeometryTypeToName(eType OGRwkbGeometryType) (result string) {
	result = C.GoString(C.OGRGeometryTypeToName(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) ToName() (result string) {
	result = ogrGeometryTypeToName(gt)
	return
}

func ogrMergeGeometryTypes(eMain, eExtra OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGRMergeGeometryTypes(C.OGRwkbGeometryType(eMain), C.OGRwkbGeometryType(eExtra)))
	return
}

func (gt OGRwkbGeometryType) Merge(extra OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = ogrMergeGeometryTypes(gt, extra)
	return
}

func ogrMergeGeometryTypesEx(eMain, eExtra OGRwkbGeometryType, bAllowPromotingToCurves int) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGRMergeGeometryTypesEx(C.OGRwkbGeometryType(eMain), C.OGRwkbGeometryType(eExtra), C.int(bAllowPromotingToCurves)))
	return
}

func (gt OGRwkbGeometryType) MergeEx(extra OGRwkbGeometryType, allowPromotingToCurves int) (result OGRwkbGeometryType) {
	result = ogrMergeGeometryTypesEx(gt, extra, allowPromotingToCurves)
	return
}

func ogrGTFlatten(eType OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_Flatten(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) Flatten() (result OGRwkbGeometryType) {
	result = ogrGTFlatten(gt)
	return
}

func ogrGTSetZ(eType OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_SetZ(C.OGRwkbGeometryType(eType)))
	return
}

func (gt *OGRwkbGeometryType) SetZ() {
	*gt = ogrGTSetZ(*gt)
}

func ogrGTSetM(eType OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_SetM(C.OGRwkbGeometryType(eType)))
	return
}

func (gt *OGRwkbGeometryType) SetM() {
	*gt = ogrGTSetM(*gt)
}

func ogrGTSetModifier(eType OGRwkbGeometryType, bSetZ, bSetM int) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_SetModifier(C.OGRwkbGeometryType(eType), C.int(bSetZ), C.int(bSetM)))
	return
}

func (gt *OGRwkbGeometryType) SetModifier(bSetZ, bSetM int) {
	*gt = ogrGTSetModifier(*gt, bSetZ, bSetM)
}

func ogrGTHasZ(eType OGRwkbGeometryType) (result int) {
	result = int(C.OGR_GT_HasZ(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) HasZ() (result bool) {
	result = ogrGTHasZ(gt) != 0
	return
}

func ogrGTHasM(eType OGRwkbGeometryType) (result int) {
	result = int(C.OGR_GT_HasM(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) HasM() (result bool) {
	result = ogrGTHasM(gt) != 0
	return
}

func ogrGTIsSubClassOf(eType, eSuperType OGRwkbGeometryType) (result int) {
	result = int(C.OGR_GT_IsSubClassOf(C.OGRwkbGeometryType(eType), C.OGRwkbGeometryType(eSuperType)))
	return
}

func (gt OGRwkbGeometryType) IsSubClassOf(eSuperType OGRwkbGeometryType) (result bool) {
	result = ogrGTIsSubClassOf(gt, eSuperType) != 0
	return
}

func ogrGTIsCurve(eType OGRwkbGeometryType) (result int) {
	result = int(C.OGR_GT_IsCurve(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) IsCurve() (result bool) {
	result = ogrGTIsCurve(gt) != 0
	return
}

func ogrGTIsSurface(eType OGRwkbGeometryType) (result int) {
	result = int(C.OGR_GT_IsSurface(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) IsSurface() (result bool) {
	result = ogrGTIsSurface(gt) != 0
	return
}

func ogrGTIsNonLinear(eType OGRwkbGeometryType) (result int) {
	result = int(C.OGR_GT_IsNonLinear(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) IsNonLinear() (result bool) {
	result = ogrGTIsNonLinear(gt) != 0
	return
}

func ogrGTGetCollection(eType OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_GetCollection(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) GetCollection() (result OGRwkbGeometryType) {
	result = ogrGTGetCollection(gt)
	return
}

func ogrGTGetSingle(eType OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_GetSingle(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) GetSingle() (result OGRwkbGeometryType) {
	result = ogrGTGetSingle(gt)
	return
}

func ogrGTGetCurve(eType OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_GetCurve(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) GetCurve() (result OGRwkbGeometryType) {
	result = ogrGTGetCurve(gt)
	return
}

func ogrGTGetLinear(eType OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = OGRwkbGeometryType(C.OGR_GT_GetLinear(C.OGRwkbGeometryType(eType)))
	return
}

func (gt OGRwkbGeometryType) GetLinear() (result OGRwkbGeometryType) {
	result = ogrGTGetLinear(gt)
	return
}

type OGRwkbByteOrder C.OGRwkbByteOrder

const (
	WkbXDR OGRwkbByteOrder = C.wkbXDR
	WkbNDR OGRwkbByteOrder = C.wkbNDR
)

/*
These lines should be HACK_FOR_IBM_DB2_V72
*/

// /** Alter field name.
//   - Used by OGR_L_AlterFieldDefn().
//     */
const AlterNameFlag = C.ALTER_NAME_FLAG

// /** Alter field type.
//   - Used by OGR_L_AlterFieldDefn().
//     */
const AlterTypeFlag = C.ALTER_TYPE_FLAG

// /** Alter field width and precision.
//   - Used by OGR_L_AlterFieldDefn().
//     */
const AlterWidthPrecisionFlag = C.ALTER_WIDTH_PRECISION_FLAG

// /** Alter field NOT NULL constraint.
//   - Used by OGR_L_AlterFieldDefn().
//     */
const AlterNullableFlag = C.ALTER_NULLABLE_FLAG

// /** Alter field DEFAULT value.
//   - Used by OGR_L_AlterFieldDefn().
//     */
const AlterDefaultFlag = C.ALTER_DEFAULT_FLAG

// /** Alter field UNIQUE constraint.
//   - Used by OGR_L_AlterFieldDefn().
//   - @since GDAL 3.2
//     */
const AlterUniqueFlag = C.ALTER_UNIQUE_FLAG

// /** Alter field domain name.
//   - Used by OGR_L_AlterFieldDefn().
//   - @since GDAL 3.3
//     */
const AlterDomainFlag = C.ALTER_DOMAIN_FLAG

// /** Alter field alternative name.
//   - Used by OGR_L_AlterFieldDefn().
//   - @since GDAL 3.7
//     */
const AlterAlternativeNameFlag = C.ALTER_ALTERNATIVE_NAME_FLAG

// /** Alter field comment.
//   - Used by OGR_L_AlterFieldDefn().
//   - @since GDAL 3.7
//     */
const AlterCommentFlag = C.ALTER_COMMENT_FLAG

// /** Alter all parameters of field definition.
//   - Used by OGR_L_AlterFieldDefn().
//     */
var AlterAllFlag = int(C._ALTER_ALL_FLAG)

// /** Alter geometry field name.
//   - Used by OGR_L_AlterGeomFieldDefn().
//   - @since GDAL 3.6
//     */
const AlterGeomFieldDefnNameFlag = C.ALTER_GEOM_FIELD_DEFN_NAME_FLAG

// /** Alter geometry field type.
//   - Used by OGR_L_AlterGeomFieldDefn().
//   - @since GDAL 3.6
//     */
const AlterGeomFieldDefnTypeFlag = C.ALTER_GEOM_FIELD_DEFN_TYPE_FLAG

// /** Alter geometry field nullable state.
//   - Used by OGR_L_AlterGeomFieldDefn().
//   - @since GDAL 3.6
//     */
const AlterGeomFieldDefnNullableFlag = C.ALTER_GEOM_FIELD_DEFN_NULLABLE_FLAG

// /** Alter geometry field spatial reference system (except its coordinate epoch)
//   - Used by OGR_L_AlterGeomFieldDefn().
//   - @since GDAL 3.6
//     */
const AlterGeomFieldDefnSRSFlag = C.ALTER_GEOM_FIELD_DEFN_SRS_FLAG

// /** Alter geometry field coordinate epoch
//   - Used by OGR_L_AlterGeomFieldDefn().
//   - @since GDAL 3.6
//     */
const AlterGeomFieldDefnSRSCoordEpochFlag = C.ALTER_GEOM_FIELD_DEFN_SRS_COORD_EPOCH_FLAG

// /** Alter all parameters of field definition.
//   - Used by OGR_L_AlterGeomFieldDefn().
//   - @since GDAL 3.6
//     */
var AlterGeomFieldDefnAllFlag = int(C._ALTER_GEOM_FIELD_DEFN_ALL_FLAG)

// /** Validate that fields respect not-null constraints.
//   - Used by OGR_F_Validate().
//     */
const OGRFValNull = C.OGR_F_VAL_NULL

// /** Validate that geometries respect geometry column type.
//   - Used by OGR_F_Validate().
//     */
const OGRFValGeomType = C.OGR_F_VAL_GEOM_TYPE

// /** Validate that (string) fields respect field width.
//   - Used by OGR_F_Validate().
//     */
const OGRFValWidth = C.OGR_F_VAL_WIDTH

// /** Allow fields that are null when there's an associated default value.
//   - This can be used for drivers where the low-level layers will automatically
//   - set the field value to the associated default value. This flag only makes
//   - sense if OGR_F_VAL_NULL is set too. Used by OGR_F_Validate().
//     */
const OGRFValAllowNullWhenDefault = C.OGR_F_VAL_ALLOW_NULL_WHEN_DEFAULT

// /** Allow geometry fields to have a different coordinate dimension that their
//   - geometry column type.
//   - This flag only makes sense if OGR_F_VAL_GEOM_TYPE is set too.
//   - Used by OGR_F_Validate().
//     */
const OGRFValAllowDifferentGeomDim = C.OGR_F_VAL_ALLOW_DIFFERENT_GEOM_DIM

// /** Enable all validation tests (except OGR_F_VAL_ALLOW_DIFFERENT_GEOM_DIM)
//   - Used by OGR_F_Validate().
//     */
var OGRFValAll = int(C._OGR_F_VAL_ALL)

// /************************************************************************/
// /*                  ogr_feature.h related definitions.                  */
// /************************************************************************/

// /**
//  * List of feature field types.  This list is likely to be extended in the
//  * future ... avoid coding applications based on the assumption that all
//  * field types can be known.
//  */

type OGRFieldType C.OGRFieldType

const (
	OFTInteger        OGRFieldType = C.OFTInteger
	OFTIntegerList    OGRFieldType = C.OFTIntegerList
	OFTReal           OGRFieldType = C.OFTReal
	OFTRealList       OGRFieldType = C.OFTRealList
	OFTString         OGRFieldType = C.OFTString
	OFTStringList     OGRFieldType = C.OFTStringList
	OFTWideString     OGRFieldType = C.OFTWideString
	OFTWideStringList OGRFieldType = C.OFTWideStringList
	OFTBinary         OGRFieldType = C.OFTBinary
	OFTDate           OGRFieldType = C.OFTDate
	OFTTime           OGRFieldType = C.OFTTime
	OFTDateTime       OGRFieldType = C.OFTDateTime
	OFTInteger64      OGRFieldType = C.OFTInteger64
	OFTInteger64List  OGRFieldType = C.OFTInteger64List
	OFTMaxType        OGRFieldType = C.OFTMaxType
)

// /**
//   - List of field subtypes. A subtype represents a hint, a restriction of the
//   - main type, that is not strictly necessary to consult.
//   - This list is likely to be extended in the
//   - future ... avoid coding applications based on the assumption that all
//   - field types can be known.
//   - Most subtypes only make sense for a restricted set of main types.
//     */
type OGRFieldSubType C.OGRFieldSubType

const (
	OFSTNone       OGRFieldSubType = C.OFSTNone
	OFSTBoolean    OGRFieldSubType = C.OFSTBoolean
	OFSTInt16      OGRFieldSubType = C.OFSTInt16
	OFSTFloat32    OGRFieldSubType = C.OFSTFloat32
	OFSTJSON       OGRFieldSubType = C.OFSTJSON
	OFSTUUID       OGRFieldSubType = C.OFSTUUID
	OFSTMaxSubType OGRFieldSubType = C.OFSTMaxSubType
)

// /**
//  * Display justification for field values.
//  */

type OGRJustification C.OGRJustification

const (
	OJUndefined OGRJustification = C.OJUndefined
	OJLeft      OGRJustification = C.OJLeft
	OJRight     OGRJustification = C.OJRight
)

// /** Special value for a unset FID */
const OGRNullFID = C.OGRNullFID

// /** Special value set in OGRField.Set.nMarker1, nMarker2 and nMarker3 for
//   - a unset field.
//   - Direct use of this value is strongly discouraged.
//   - Use OGR_RawField_SetUnset() or OGR_RawField_IsUnset() instead.
//     */
const OGRUnsetMarker = C.OGRUnsetMarker

// /** Special value set in OGRField.Set.nMarker1, nMarker2 and nMarker3 for
//   - a null field.
//   - Direct use of this value is strongly discouraged.
//   - Use OGR_RawField_SetNull() or OGR_RawField_IsNull() instead.
//     */
const OGRNullMarker = C.OGRNullMarker

// /** Time zone flag indicating unknown timezone. For the
//   - OGRFieldDefn::GetTZFlag() property, this may also indicate a mix of
//   - unknown, localtime or known time zones in the same field.
//     */
const OGRTZFlagUnknown = C.OGR_TZFLAG_UNKNOWN

// /** Time zone flag indicating local time */
const OGRTZFlagLocaltime = C.OGR_TZFLAG_LOCALTIME

// /** Time zone flag only returned by OGRFieldDefn::GetTZFlag() to indicate
//   - that all values in the field have a known time zone (ie different from
//   - OGR_TZFLAG_UNKNOWN and OGR_TZFLAG_LOCALTIME), but it may be different among
//   - features. */
const OGRTZFlagMixedTZ = C.OGR_TZFLAG_MIXED_TZ

// /** Time zone flag indicating UTC.
//   - Used to derived other time zone flags with the following logic:
//   - - values above 100 indicate a 15 minute increment per unit.
//   - - values under 100 indicate a 15 minute decrement per unit.
//   - For example: a value of 101 indicates UTC+00:15, a value of 102 UTC+00:30,
//   - a value of 99 UTC-00:15 ,etc. */
const OGRTZFlagUTC = C.OGR_TZFLAG_UTC

type OGRField struct {
	cValue *C.OGRField
}

func ogrGetMS(fSec float32) (result int) {
	result = int(C.OGR_GET_MS(C.float(fSec)))
	return
}

func OGRGetMS(fSec float32) (result int) {
	result = ogrGetMS(fSec)
	return
}

// /** Option for OGRParseDate() to ask for lax checks on the input format */
const OGRParseDateOptionLax = C.OGRPARSEDATE_OPTION_LAX

func ogrParseDate(pszInput string, psOutput OGRField, nOptions int) (result int) {
	cs := C.CString(pszInput)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.OGRParseDate(cs, psOutput.cValue, C.int(nOptions)))
	return
}

func OGRParseDate(input string, options int) (result OGRField, ok bool) {
	result = OGRField{cValue: new(C.OGRField)}
	ok = ogrParseDate(input, result, options) != 0
	return
}

// Layer capabilities (OLC*)
var (
	OLCRandomRead          = C.GoString(C._OLCRandomRead)
	OLCSequentialWrite     = C.GoString(C._OLCSequentialWrite)
	OLCRandomWrite         = C.GoString(C._OLCRandomWrite)
	OLCFastSpatialFilter   = C.GoString(C._OLCFastSpatialFilter)
	OLCFastFeatureCount    = C.GoString(C._OLCFastFeatureCount)
	OLCFastGetExtent       = C.GoString(C._OLCFastGetExtent)
	OLCFastGetExtent3D     = C.GoString(C._OLCFastGetExtent3D)
	OLCCreateField         = C.GoString(C._OLCCreateField)
	OLCDeleteField         = C.GoString(C._OLCDeleteField)
	OLCReorderFields       = C.GoString(C._OLCReorderFields)
	OLCAlterFieldDefn      = C.GoString(C._OLCAlterFieldDefn)
	OLCAlterGeomFieldDefn  = C.GoString(C._OLCAlterGeomFieldDefn)
	OLCTransactions        = C.GoString(C._OLCTransactions)
	OLCDeleteFeature       = C.GoString(C._OLCDeleteFeature)
	OLCUpsertFeature       = C.GoString(C._OLCUpsertFeature)
	OLCUpdateFeature       = C.GoString(C._OLCUpdateFeature)
	OLCFastSetNextByIndex  = C.GoString(C._OLCFastSetNextByIndex)
	OLCStringsAsUTF8       = C.GoString(C._OLCStringsAsUTF8)
	OLCIgnoreFields        = C.GoString(C._OLCIgnoreFields)
	OLCCreateGeomField     = C.GoString(C._OLCCreateGeomField)
	OLCCurveGeometries     = C.GoString(C._OLCCurveGeometries)
	OLCMeasuredGeometries  = C.GoString(C._OLCMeasuredGeometries)
	OLCZGeometries         = C.GoString(C._OLCZGeometries)
	OLCRename              = C.GoString(C._OLCRename)
	OLCFastGetArrowStream  = C.GoString(C._OLCFastGetArrowStream)
	OLCFastWriteArrowBatch = C.GoString(C._OLCFastWriteArrowBatch)
)

// Dataset capabilities (ODs*)
var (
	ODsCCreateLayer                     = C.GoString(C._ODsCCreateLayer)
	ODsCDeleteLayer                     = C.GoString(C._ODsCDeleteLayer)
	ODsCCreateGeomFieldAfterCreateLayer = C.GoString(C._ODsCCreateGeomFieldAfterCreateLayer)
	ODsCCurveGeometries                 = C.GoString(C._ODsCCurveGeometries)
	ODsCTransactions                    = C.GoString(C._ODsCTransactions)
	ODsCEmulatedTransactions            = C.GoString(C._ODsCEmulatedTransactions)
	ODsCMeasuredGeometries              = C.GoString(C._ODsCMeasuredGeometries)
	ODsCZGeometries                     = C.GoString(C._ODsCZGeometries)
	ODsCRandomLayerRead                 = C.GoString(C._ODsCRandomLayerRead)
	ODsCRandomLayerWrite                = C.GoString(C._ODsCRandomLayerWrite)
	ODsCAddFieldDomain                  = C.GoString(C._ODsCAddFieldDomain)
	ODsCDeleteFieldDomain               = C.GoString(C._ODsCDeleteFieldDomain)
	ODsCUpdateFieldDomain               = C.GoString(C._ODsCUpdateFieldDomain)
)

// Driver capabilities (ODr*)
var (
	ODrCCreateDataSource = C.GoString(C._ODrCCreateDataSource)
	ODrCDeleteDataSource = C.GoString(C._ODrCDeleteDataSource)
)

// Layer metadata items
var OLMD_FID64 = C.GoString(C._OLMD_FID64)

// /************************************************************************/
// /*                  ogr_featurestyle.h related definitions.             */
// /************************************************************************/

// /**
//  * OGRStyleTool derived class types (returned by GetType()).
//  */

type OGRSTClassId C.OGRSTClassId

const (
	OGRSTCNone   OGRSTClassId = C.OGRSTCNone
	OGRSTCPen    OGRSTClassId = C.OGRSTCPen
	OGRSTCBrush  OGRSTClassId = C.OGRSTCBrush
	OGRSTCSymbol OGRSTClassId = C.OGRSTCSymbol
	OGRSTCLabel  OGRSTClassId = C.OGRSTCLabel
	OGRSTCVector OGRSTClassId = C.OGRSTCVector
)

// /**
//   - List of units supported by OGRStyleTools.
//     */
type OGRSTUnitId C.OGRSTUnitId

const (
	OGRSTUGround OGRSTUnitId = C.OGRSTUGround
	OGRSTUPixel  OGRSTUnitId = C.OGRSTUPixel
	OGRSTUPoints OGRSTUnitId = C.OGRSTUPoints
	OGRSTUMM     OGRSTUnitId = C.OGRSTUMM
	OGRSTUCM     OGRSTUnitId = C.OGRSTUCM
	OGRSTUInches OGRSTUnitId = C.OGRSTUInches
)

type OGRSTPenParam C.OGRSTPenParam

const (
	OGRSTPenColor     OGRSTPenParam = C.OGRSTPenColor
	OGRSTPenWidth     OGRSTPenParam = C.OGRSTPenWidth
	OGRSTPenPattern   OGRSTPenParam = C.OGRSTPenPattern
	OGRSTPenId        OGRSTPenParam = C.OGRSTPenId
	OGRSTPenPerOffset OGRSTPenParam = C.OGRSTPenPerOffset
	OGRSTPenCap       OGRSTPenParam = C.OGRSTPenCap
	OGRSTPenJoin      OGRSTPenParam = C.OGRSTPenJoin
	OGRSTPenPriority  OGRSTPenParam = C.OGRSTPenPriority
	OGRSTPenLast      OGRSTPenParam = C.OGRSTPenLast
)

type OGRSTBrushParam C.OGRSTBrushParam

const (
	OGRSTBrushFColor   OGRSTBrushParam = C.OGRSTBrushFColor
	OGRSTBrushBColor   OGRSTBrushParam = C.OGRSTBrushBColor
	OGRSTBrushId       OGRSTBrushParam = C.OGRSTBrushId
	OGRSTBrushAngle    OGRSTBrushParam = C.OGRSTBrushAngle
	OGRSTBrushSize     OGRSTBrushParam = C.OGRSTBrushSize
	OGRSTBrushDx       OGRSTBrushParam = C.OGRSTBrushDx
	OGRSTBrushDy       OGRSTBrushParam = C.OGRSTBrushDy
	OGRSTBrushPriority OGRSTBrushParam = C.OGRSTBrushPriority
	OGRSTBrushLast     OGRSTBrushParam = C.OGRSTBrushLast
)

type OGRSTSymbolParam C.OGRSTSymbolParam

const (
	OGRSTSymbolId       OGRSTSymbolParam = C.OGRSTSymbolId
	OGRSTSymbolAngle    OGRSTSymbolParam = C.OGRSTSymbolAngle
	OGRSTSymbolColor    OGRSTSymbolParam = C.OGRSTSymbolColor
	OGRSTSymbolSize     OGRSTSymbolParam = C.OGRSTSymbolSize
	OGRSTSymbolDx       OGRSTSymbolParam = C.OGRSTSymbolDx
	OGRSTSymbolDy       OGRSTSymbolParam = C.OGRSTSymbolDy
	OGRSTSymbolStep     OGRSTSymbolParam = C.OGRSTSymbolStep
	OGRSTSymbolPerp     OGRSTSymbolParam = C.OGRSTSymbolPerp
	OGRSTSymbolOffset   OGRSTSymbolParam = C.OGRSTSymbolOffset
	OGRSTSymbolPriority OGRSTSymbolParam = C.OGRSTSymbolPriority
	OGRSTSymbolFontName OGRSTSymbolParam = C.OGRSTSymbolFontName
	OGRSTSymbolOColor   OGRSTSymbolParam = C.OGRSTSymbolOColor
	OGRSTSymbolLast     OGRSTSymbolParam = C.OGRSTSymbolLast
)

type OGRSTLabelParam C.OGRSTLabelParam

const (
	OGRSTLabelFontName   OGRSTLabelParam = C.OGRSTLabelFontName
	OGRSTLabelSize       OGRSTLabelParam = C.OGRSTLabelSize
	OGRSTLabelTextString OGRSTLabelParam = C.OGRSTLabelTextString
	OGRSTLabelAngle      OGRSTLabelParam = C.OGRSTLabelAngle
	OGRSTLabelFColor     OGRSTLabelParam = C.OGRSTLabelFColor
	OGRSTLabelBColor     OGRSTLabelParam = C.OGRSTLabelBColor
	OGRSTLabelPlacement  OGRSTLabelParam = C.OGRSTLabelPlacement
	OGRSTLabelAnchor     OGRSTLabelParam = C.OGRSTLabelAnchor
	OGRSTLabelDx         OGRSTLabelParam = C.OGRSTLabelDx
	OGRSTLabelDy         OGRSTLabelParam = C.OGRSTLabelDy
	OGRSTLabelPerp       OGRSTLabelParam = C.OGRSTLabelPerp
	OGRSTLabelBold       OGRSTLabelParam = C.OGRSTLabelBold
	OGRSTLabelItalic     OGRSTLabelParam = C.OGRSTLabelItalic
	OGRSTLabelUnderline  OGRSTLabelParam = C.OGRSTLabelUnderline
	OGRSTLabelPriority   OGRSTLabelParam = C.OGRSTLabelPriority
	OGRSTLabelStrikeout  OGRSTLabelParam = C.OGRSTLabelStrikeout
	OGRSTLabelStretch    OGRSTLabelParam = C.OGRSTLabelStretch
	OGRSTLabelAdjHor     OGRSTLabelParam = C.OGRSTLabelAdjHor
	OGRSTLabelAdjVert    OGRSTLabelParam = C.OGRSTLabelAdjVert
	OGRSTLabelHColor     OGRSTLabelParam = C.OGRSTLabelHColor
	OGRSTLabelOColor     OGRSTLabelParam = C.OGRSTLabelOColor
	OGRSTLabelLast       OGRSTLabelParam = C.OGRSTLabelLast
)

// /* -------------------------------------------------------------------- */
// /*                          Field domains                               */
// /* -------------------------------------------------------------------- */

type OGRCodedValue struct {
	cValue *C.OGRCodedValue
}

type OGRFieldDomainType C.OGRFieldDomainType

const (
	OFDT_CODED OGRFieldDomainType = C.OFDT_CODED
	OFDT_RANGE OGRFieldDomainType = C.OFDT_RANGE
	OFDT_GLOB  OGRFieldDomainType = C.OFDT_GLOB
)

type OGRFieldDomainSplitPolicy C.OGRFieldDomainSplitPolicy

const (
	OFDSP_DEFAULT_VALUE  OGRFieldDomainSplitPolicy = C.OFDSP_DEFAULT_VALUE
	OFDSP_DUPLICATE      OGRFieldDomainSplitPolicy = C.OFDSP_DUPLICATE
	OFDSP_GEOMETRY_RATIO OGRFieldDomainSplitPolicy = C.OFDSP_GEOMETRY_RATIO
)

type OGRFieldDomainMergePolicy C.OGRFieldDomainMergePolicy

const (
	OFDMP_DEFAULT_VALUE     OGRFieldDomainMergePolicy = C.OFDMP_DEFAULT_VALUE
	OFDMP_SUM               OGRFieldDomainMergePolicy = C.OFDMP_SUM
	OFDMP_GEOMETRY_WEIGHTED OGRFieldDomainMergePolicy = C.OFDMP_GEOMETRY_WEIGHTED
)

// /* ------------------------------------------------------------------- */
// /*                        Version checking                             */
// /* -------------------------------------------------------------------- */

// /* Note to developers : please keep this section in sync with gdal.h */

func gdalVersionInfo(pszRequest string) (result string) {
	cs := C.CString(pszRequest)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.GDALVersionInfo(cs))
	return
}

func GDALVersionInfo(request string) (result string) {
	result = gdalVersionInfo(request)
	return
}

func gdalCheckVersion(nVersionMajor, nVersionMinor int, pszCallingComponentName string) (result int) {
	cs := C.CString(pszCallingComponentName)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.GDALCheckVersion(C.int(nVersionMajor), C.int(nVersionMinor), cs))
	return
}

func GDALCheckVersion(versionMajor, versionMinor int, callingComponentName string) (result int) {
	result = gdalCheckVersion(versionMajor, versionMinor, callingComponentName)
	return
}

// CPL_C_END

// #endif /* ndef OGR_CORE_H_INCLUDED */
