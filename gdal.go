package gdal

/*
#include "gdal.h"
#include "cpl_string.h" // TODO: implement cpl_string.go
#include "cpl_vsi.h" // TODO: implement cpl_vsi.go

const char* const _GDALMD_AREA_OR_POINT = GDALMD_AREA_OR_POINT;
const char* const _GDALMD_AOP_AREA      = GDALMD_AOP_AREA;
const char* const _GDALMD_AOP_POINT     = GDALMD_AOP_POINT;

const char* const _GDAL_DS_LAYER_CREATIONOPTIONLIST = GDAL_DS_LAYER_CREATIONOPTIONLIST;

const char* const _GDAL_DMD_LONGNAME                                       = GDAL_DMD_LONGNAME;
const char* const _GDAL_DMD_HELPTOPIC                                      = GDAL_DMD_HELPTOPIC;
const char* const _GDAL_DMD_MIMETYPE                                       = GDAL_DMD_MIMETYPE;
const char* const _GDAL_DMD_EXTENSION                                      = GDAL_DMD_EXTENSION;
const char* const _GDAL_DMD_CONNECTION_PREFIX                              = GDAL_DMD_CONNECTION_PREFIX;
const char* const _GDAL_DMD_EXTENSIONS                                     = GDAL_DMD_EXTENSIONS;
const char* const _GDAL_DMD_CREATIONOPTIONLIST                             = GDAL_DMD_CREATIONOPTIONLIST;
const char* const _GDAL_DMD_OVERVIEW_CREATIONOPTIONLIST                    = GDAL_DMD_OVERVIEW_CREATIONOPTIONLIST;
const char* const _GDAL_DMD_MULTIDIM_DATASET_CREATIONOPTIONLIST            = GDAL_DMD_MULTIDIM_DATASET_CREATIONOPTIONLIST;
const char* const _GDAL_DMD_MULTIDIM_GROUP_CREATIONOPTIONLIST              = GDAL_DMD_MULTIDIM_GROUP_CREATIONOPTIONLIST;
const char* const _GDAL_DMD_MULTIDIM_DIMENSION_CREATIONOPTIONLIST          = GDAL_DMD_MULTIDIM_DIMENSION_CREATIONOPTIONLIST;
const char* const _GDAL_DMD_MULTIDIM_ARRAY_CREATIONOPTIONLIST              = GDAL_DMD_MULTIDIM_ARRAY_CREATIONOPTIONLIST;
const char* const _GDAL_DMD_MULTIDIM_ARRAY_OPENOPTIONLIST                  = GDAL_DMD_MULTIDIM_ARRAY_OPENOPTIONLIST;
const char* const _GDAL_DMD_MULTIDIM_ATTRIBUTE_CREATIONOPTIONLIST          = GDAL_DMD_MULTIDIM_ATTRIBUTE_CREATIONOPTIONLIST;
const char* const _GDAL_DMD_OPENOPTIONLIST                                 = GDAL_DMD_OPENOPTIONLIST;
const char* const _GDAL_DMD_CREATIONDATATYPES                              = GDAL_DMD_CREATIONDATATYPES;
const char* const _GDAL_DMD_CREATIONFIELDDATATYPES                         = GDAL_DMD_CREATIONFIELDDATATYPES;
const char* const _GDAL_DMD_CREATIONFIELDDATASUBTYPES                      = GDAL_DMD_CREATIONFIELDDATASUBTYPES;
const char* const _GDAL_DMD_MAX_STRING_LENGTH                              = GDAL_DMD_MAX_STRING_LENGTH;
const char* const _GDAL_DMD_CREATION_FIELD_DEFN_FLAGS                      = GDAL_DMD_CREATION_FIELD_DEFN_FLAGS;
const char* const _GDAL_DMD_SUBDATASETS                                    = GDAL_DMD_SUBDATASETS;
const char* const _GDAL_DCAP_CREATE_SUBDATASETS                            = GDAL_DCAP_CREATE_SUBDATASETS;
const char* const _GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_DECIMAL_SEPARATOR = GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_DECIMAL_SEPARATOR;
const char* const _GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_SIGN              = GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_SIGN;
const char* const _GDAL_DCAP_OPEN                                          = GDAL_DCAP_OPEN;
const char* const _GDAL_DCAP_CREATE                                        = GDAL_DCAP_CREATE;
const char* const _GDAL_DCAP_CREATE_MULTIDIMENSIONAL                       = GDAL_DCAP_CREATE_MULTIDIMENSIONAL;
const char* const _GDAL_DCAP_CREATECOPY                                    = GDAL_DCAP_CREATECOPY;
const char* const _GDAL_DCAP_CREATE_ONLY_VISIBLE_AT_CLOSE_TIME             = GDAL_DCAP_CREATE_ONLY_VISIBLE_AT_CLOSE_TIME;
const char* const _GDAL_DCAP_VECTOR_TRANSLATE_FROM                         = GDAL_DCAP_VECTOR_TRANSLATE_FROM;
const char* const _GDAL_DCAP_CREATECOPY_MULTIDIMENSIONAL                   = GDAL_DCAP_CREATECOPY_MULTIDIMENSIONAL;
const char* const _GDAL_DCAP_MULTIDIM_RASTER                               = GDAL_DCAP_MULTIDIM_RASTER;
const char* const _GDAL_DCAP_SUBCREATECOPY                                 = GDAL_DCAP_SUBCREATECOPY;
const char* const _GDAL_DCAP_APPEND                                        = GDAL_DCAP_APPEND;
const char* const _GDAL_DCAP_UPDATE                                        = GDAL_DCAP_UPDATE;
const char* const _GDAL_DCAP_VIRTUALIO                                     = GDAL_DCAP_VIRTUALIO;
const char* const _GDAL_DCAP_RASTER                                        = GDAL_DCAP_RASTER;
const char* const _GDAL_DCAP_VECTOR                                        = GDAL_DCAP_VECTOR;
const char* const _GDAL_DCAP_GNM                                           = GDAL_DCAP_GNM;
const char* const _GDAL_DCAP_CREATE_LAYER                                  = GDAL_DCAP_CREATE_LAYER;
const char* const _GDAL_DCAP_DELETE_LAYER                                  = GDAL_DCAP_DELETE_LAYER;
const char* const _GDAL_DCAP_CREATE_FIELD                                  = GDAL_DCAP_CREATE_FIELD;
const char* const _GDAL_DCAP_DELETE_FIELD                                  = GDAL_DCAP_DELETE_FIELD;
const char* const _GDAL_DCAP_REORDER_FIELDS                                = GDAL_DCAP_REORDER_FIELDS;
const char* const _GDAL_DMD_ALTER_FIELD_DEFN_FLAGS                         = GDAL_DMD_ALTER_FIELD_DEFN_FLAGS;
const char* const _GDAL_DMD_ILLEGAL_FIELD_NAMES                            = GDAL_DMD_ILLEGAL_FIELD_NAMES;
const char* const _GDAL_DCAP_NOTNULL_FIELDS                                = GDAL_DCAP_NOTNULL_FIELDS;
const char* const _GDAL_DCAP_UNIQUE_FIELDS                                 = GDAL_DCAP_UNIQUE_FIELDS;
const char* const _GDAL_DCAP_DEFAULT_FIELDS                                = GDAL_DCAP_DEFAULT_FIELDS;
const char* const _GDAL_DCAP_NOTNULL_GEOMFIELDS                            = GDAL_DCAP_NOTNULL_GEOMFIELDS;
const char* const _GDAL_DCAP_NONSPATIAL                                    = GDAL_DCAP_NONSPATIAL;
const char* const _GDAL_DCAP_CURVE_GEOMETRIES                              = GDAL_DCAP_CURVE_GEOMETRIES;
const char* const _GDAL_DCAP_MEASURED_GEOMETRIES                           = GDAL_DCAP_MEASURED_GEOMETRIES;
const char* const _GDAL_DCAP_Z_GEOMETRIES                                  = GDAL_DCAP_Z_GEOMETRIES;
const char* const _GDAL_DMD_GEOMETRY_FLAGS                                 = GDAL_DMD_GEOMETRY_FLAGS;
const char* const _GDAL_DCAP_FEATURE_STYLES                                = GDAL_DCAP_FEATURE_STYLES;
const char* const _GDAL_DCAP_FEATURE_STYLES_READ                           = GDAL_DCAP_FEATURE_STYLES_READ;
const char* const _GDAL_DCAP_FEATURE_STYLES_WRITE                          = GDAL_DCAP_FEATURE_STYLES_WRITE;
const char* const _GDAL_DCAP_COORDINATE_EPOCH                              = GDAL_DCAP_COORDINATE_EPOCH;
const char* const _GDAL_DCAP_MULTIPLE_VECTOR_LAYERS                        = GDAL_DCAP_MULTIPLE_VECTOR_LAYERS;
const char* const _GDAL_DCAP_FIELD_DOMAINS                                 = GDAL_DCAP_FIELD_DOMAINS;
const char* const _GDAL_DCAP_RELATIONSHIPS                                 = GDAL_DCAP_RELATIONSHIPS;
const char* const _GDAL_DCAP_CREATE_RELATIONSHIP                           = GDAL_DCAP_CREATE_RELATIONSHIP;
const char* const _GDAL_DCAP_DELETE_RELATIONSHIP                           = GDAL_DCAP_DELETE_RELATIONSHIP;
const char* const _GDAL_DCAP_UPDATE_RELATIONSHIP                           = GDAL_DCAP_UPDATE_RELATIONSHIP;
const char* const _GDAL_DCAP_FLUSHCACHE_CONSISTENT_STATE                   = GDAL_DCAP_FLUSHCACHE_CONSISTENT_STATE;
const char* const _GDAL_DCAP_HONOR_GEOM_COORDINATE_PRECISION               = GDAL_DCAP_HONOR_GEOM_COORDINATE_PRECISION;
const char* const _GDAL_DCAP_UPSERT                                        = GDAL_DCAP_UPSERT;
const char* const _GDAL_DMD_RELATIONSHIP_FLAGS                             = GDAL_DMD_RELATIONSHIP_FLAGS;
const char* const _GDAL_DMD_RELATIONSHIP_RELATED_TABLE_TYPES               = GDAL_DMD_RELATIONSHIP_RELATED_TABLE_TYPES;
const char* const _GDAL_DCAP_RENAME_LAYERS                                 = GDAL_DCAP_RENAME_LAYERS;
const char* const _GDAL_DMD_CREATION_FIELD_DOMAIN_TYPES                    = GDAL_DMD_CREATION_FIELD_DOMAIN_TYPES;
const char* const _GDAL_DMD_ALTER_GEOM_FIELD_DEFN_FLAGS                    = GDAL_DMD_ALTER_GEOM_FIELD_DEFN_FLAGS;
const char* const _GDAL_DMD_SUPPORTED_SQL_DIALECTS                         = GDAL_DMD_SUPPORTED_SQL_DIALECTS;
const char* const _GDAL_DMD_PLUGIN_INSTALLATION_MESSAGE                    = GDAL_DMD_PLUGIN_INSTALLATION_MESSAGE;
const char* const _GDAL_DMD_UPDATE_ITEMS                                   = GDAL_DMD_UPDATE_ITEMS;
const char* const _GDAL_DIM_TYPE_HORIZONTAL_X                              = GDAL_DIM_TYPE_HORIZONTAL_X;
const char* const _GDAL_DIM_TYPE_HORIZONTAL_Y                              = GDAL_DIM_TYPE_HORIZONTAL_Y;
const char* const _GDAL_DIM_TYPE_VERTICAL                                  = GDAL_DIM_TYPE_VERTICAL;
const char* const _GDAL_DIM_TYPE_TEMPORAL                                  = GDAL_DIM_TYPE_TEMPORAL;
const char* const _GDAL_DIM_TYPE_PARAMETRIC                                = GDAL_DIM_TYPE_PARAMETRIC;
const char* const _GDAL_DCAP_REOPEN_AFTER_WRITE_REQUIRED                   = GDAL_DCAP_REOPEN_AFTER_WRITE_REQUIRED;
const char* const _GDAL_DCAP_CAN_READ_AFTER_DELETE                         = GDAL_DCAP_CAN_READ_AFTER_DELETE;
const char* const _GDsCAddRelationship                                     = GDsCAddRelationship;
const char* const _GDsCDeleteRelationship                                  = GDsCDeleteRelationship;
const char* const _GDsCUpdateRelationship                                  = GDsCUpdateRelationship;
const char* const _GDsCFastGetExtent                                       = GDsCFastGetExtent;
const char* const _GDsCFastGetExtentWGS84LongLat                           = GDsCFastGetExtentWGS84LongLat;
*/
import "C"
import "unsafe"

// CPL_C_START

// Pixel data types.
type GDALDataType C.GDALDataType

const (
	GDTUnknown   GDALDataType = C.GDT_Unknown
	GDTByte      GDALDataType = C.GDT_Byte
	GDTInt8      GDALDataType = C.GDT_Int8
	GDTUInt16    GDALDataType = C.GDT_UInt16
	GDTInt16     GDALDataType = C.GDT_Int16
	GDTUInt32    GDALDataType = C.GDT_UInt32
	GDTInt32     GDALDataType = C.GDT_Int32
	GDTUInt64    GDALDataType = C.GDT_UInt64
	GDTInt64     GDALDataType = C.GDT_Int64
	GDTFloat16   GDALDataType = C.GDT_Float16
	GDTFloat32   GDALDataType = C.GDT_Float32
	GDTFloat64   GDALDataType = C.GDT_Float64
	GDTCInt16    GDALDataType = C.GDT_CInt16
	GDTCInt32    GDALDataType = C.GDT_CInt32
	GDTCFloat16  GDALDataType = C.GDT_CFloat16
	GDTCFloat32  GDALDataType = C.GDT_CFloat32
	GDTCFloat64  GDALDataType = C.GDT_CFloat64
	GDTTypeCount GDALDataType = C.GDT_TypeCount
)

func gdalGetDataTypeSize(dataType GDALDataType) (result int) {
	result = int(C.GDALGetDataTypeSize(C.GDALDataType(dataType)))
	return
}

// Deprecated: use SizeBits or SizeBytes instead.
func (dt GDALDataType) Size() (result int) {
	result = gdalGetDataTypeSize(dt)
	return
}

func gdalGetDataTypeSizeBits(dataType GDALDataType) (result int) {
	result = int(C.GDALGetDataTypeSizeBits(C.GDALDataType(dataType)))
	return
}

func (dt GDALDataType) SizeBits() (result int) {
	result = gdalGetDataTypeSizeBits(dt)
	return
}

func gdalGetDataTypeSizeBytes(dataType GDALDataType) (result int) {
	result = int(C.GDALGetDataTypeSizeBytes(C.GDALDataType(dataType)))
	return
}

func (dt GDALDataType) SizeBytes() (result int) {
	result = gdalGetDataTypeSizeBytes(dt)
	return
}

func gdalDataTypeIsComplex(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsComplex(C.GDALDataType(dataType)) != 0
	return
}

func (dt GDALDataType) IsComplex() (result bool) {
	result = gdalDataTypeIsComplex(dt)
	return
}

func gdalDataTypeIsInteger(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsInteger(C.GDALDataType(dataType)) != 0
	return
}

func (dt GDALDataType) IsInteger() (result bool) {
	result = gdalDataTypeIsInteger(dt)
	return
}

func gdalDataTypeIsFloating(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsFloating(C.GDALDataType(dataType)) != 0
	return
}

func (dt GDALDataType) IsFloating() (result bool) {
	result = gdalDataTypeIsFloating(dt)
	return
}

func gdalDataTypeIsSigned(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsSigned(C.GDALDataType(dataType)) != 0
	return
}

func (dt GDALDataType) IsSigned() (result bool) {
	result = gdalDataTypeIsSigned(dt)
	return
}

func gdalGetDataTypeName(dataType GDALDataType) (result string) {
	result = C.GoString(C.GDALGetDataTypeName(C.GDALDataType(dataType)))
	return
}

func (dt GDALDataType) GetName() (result string) {
	result = gdalGetDataTypeName(dt)
	return
}

func gdalGetDataTypeByName(name string) (result GDALDataType) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDataType(C.GDALGetDataTypeByName(cName))
	return
}

func GDALGetDataTypeByName(name string) (result GDALDataType) {
	result = gdalGetDataTypeByName(name)
	return
}

func gdalDataTypeUnion(a, b GDALDataType) (result GDALDataType) {
	result = GDALDataType(C.GDALDataTypeUnion(C.GDALDataType(a), C.GDALDataType(b)))
	return
}

func (dt GDALDataType) Union(other GDALDataType) (result GDALDataType) {
	result = gdalDataTypeUnion(dt, other)
	return
}

func gdalDataTypeUnionWithValue(dataType GDALDataType, value float64, complex int) (result GDALDataType) {
	result = GDALDataType(C.GDALDataTypeUnionWithValue(C.GDALDataType(dataType), C.double(value), C.int(complex)))
	return
}

func (dt GDALDataType) UnionWithValue(value float64, complex int) (result GDALDataType) {
	result = gdalDataTypeUnionWithValue(dt, value, complex)
	return
}

func gdalFindDataType(bits, signed, floating, complex int) (result GDALDataType) {
	result = GDALDataType(C.GDALFindDataType(C.int(bits), C.int(signed), C.int(floating), C.int(complex)))
	return
}

func GDALFindDataType(bits, signed, floating, complex int) (result GDALDataType) {
	result = gdalFindDataType(bits, signed, floating, complex)
	return
}

func gdalFindDataTypeForValue(value float64, complex int) (result GDALDataType) {
	result = GDALDataType(C.GDALFindDataTypeForValue(C.double(value), C.int(complex)))
	return
}

func GDALFindDataTypeForValue(value float64, complex int) (result GDALDataType) {
	result = gdalFindDataTypeForValue(value, complex)
	return
}

func gdalAdjustValueToDataType(dataType GDALDataType, value float64, clamped, rounded *int) float64 {
	var cClamped, cRounded C.int
	r := C.GDALAdjustValueToDataType(C.GDALDataType(dataType), C.double(value), &cClamped, &cRounded)
	*clamped = int(cClamped)
	*rounded = int(cRounded)
	return float64(r)
}

func (dt GDALDataType) AdjustValue(value float64) (result float64, clamped, rounded int) {
	result = gdalAdjustValueToDataType(dt, value, &clamped, &rounded)
	return
}

func gdalIsValueExactAs(value float64, dataType GDALDataType) (result bool) {
	result = bool(C.GDALIsValueExactAs(C.double(value), C.GDALDataType(dataType)))
	return
}

func (dt GDALDataType) IsValueExactAs(value float64) (result bool) {
	result = gdalIsValueExactAs(value, dt)
	return
}

func gdalIsValueInRangeOf(value float64, dataType GDALDataType) (result bool) {
	result = bool(C.GDALIsValueInRangeOf(C.double(value), C.GDALDataType(dataType)))
	return
}

func (dt GDALDataType) IsValueInRangeOf(value float64) (result bool) {
	result = gdalIsValueInRangeOf(value, dt)
	return
}

func gdalGetNonComplexDataType(dataType GDALDataType) (result GDALDataType) {
	result = GDALDataType(C.GDALGetNonComplexDataType(C.GDALDataType(dataType)))
	return
}

func (dt GDALDataType) GetNonComplex() (result GDALDataType) {
	result = gdalGetNonComplexDataType(dt)
	return
}

func gdalDataTypeIsConversionLossy(from, to GDALDataType) (result bool) {
	result = C.GDALDataTypeIsConversionLossy(C.GDALDataType(from), C.GDALDataType(to)) != 0
	return
}

func (dt GDALDataType) IsConversionLossyTo(to GDALDataType) (result bool) {
	result = gdalDataTypeIsConversionLossy(dt, to)
	return
}

// Status of the asynchronous stream.
type GDALAsyncStatusType C.GDALAsyncStatusType

const (
	GARIOPending   GDALAsyncStatusType = C.GARIO_PENDING
	GARIOUpdate    GDALAsyncStatusType = C.GARIO_UPDATE
	GARIOError     GDALAsyncStatusType = C.GARIO_ERROR
	GARIOComplete  GDALAsyncStatusType = C.GARIO_COMPLETE
	GARIOTypeCount GDALAsyncStatusType = C.GARIO_TypeCount
)

func gdalGetAsyncStatusTypeName(status GDALAsyncStatusType) (result string) {
	result = C.GoString(C.GDALGetAsyncStatusTypeName(C.GDALAsyncStatusType(status)))
	return
}

func (s GDALAsyncStatusType) GetName() (result string) {
	result = gdalGetAsyncStatusTypeName(s)
	return
}

func gdalGetAsyncStatusTypeByName(name string) (result GDALAsyncStatusType) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALAsyncStatusType(C.GDALGetAsyncStatusTypeByName(cName))
	return
}

func GDALGetAsyncStatusTypeByName(name string) (result GDALAsyncStatusType) {
	result = gdalGetAsyncStatusTypeByName(name)
	return
}

// Flag indicating read/write, or read-only access to data.
type GDALAccess C.GDALAccess

const (
	GAReadOnly GDALAccess = C.GA_ReadOnly
	GAUpdate   GDALAccess = C.GA_Update
)

// Read/Write flag for RasterIO() method.
type GDALRWFlag C.GDALRWFlag

const (
	GFRead  GDALRWFlag = C.GF_Read
	GFWrite GDALRWFlag = C.GF_Write
)

// RasterIO() resampling method.
type GDALRIOResampleAlg C.GDALRIOResampleAlg

const (
	GRIORANearestNeighbour GDALRIOResampleAlg = C.GRIORA_NearestNeighbour
	GRIORABilinear         GDALRIOResampleAlg = C.GRIORA_Bilinear
	GRIORACubic            GDALRIOResampleAlg = C.GRIORA_Cubic
	GRIORACubicSpline      GDALRIOResampleAlg = C.GRIORA_CubicSpline
	GRIORALanczos          GDALRIOResampleAlg = C.GRIORA_Lanczos
	GRIORAAverage          GDALRIOResampleAlg = C.GRIORA_Average
	GRIORAMode             GDALRIOResampleAlg = C.GRIORA_Mode
	GRIORAGauss            GDALRIOResampleAlg = C.GRIORA_Gauss
	// RMS: Root Mean Square / Quadratic Mean.
	GRIORARMS GDALRIOResampleAlg = C.GRIORA_RMS
)

// Structure to pass extra arguments to RasterIO() method.
type GDALRasterIOExtraArg struct {
	cValue *C.GDALRasterIOExtraArg
}

const RasterIOExtraArgCurrentVersion = C.RASTERIO_EXTRA_ARG_CURRENT_VERSION

// GCIIRStart/GCIIREnd bound the InfraRed (IR) domain color interpretations.
// GCISARStart/GCISAREnd bound the Synthetic Aperture Radar (SAR) domain.
const (
	GCIIRStart  = C.GCI_IR_Start
	GCIIREnd    = C.GCI_IR_End
	GCISARStart = C.GCI_SAR_Start
	GCISAREnd   = C.GCI_SAR_End
)

// Types of color interpretation for raster bands.
type GDALColorInterp C.GDALColorInterp

const (
	GCIUndefined      GDALColorInterp = C.GCI_Undefined
	GCIGrayIndex      GDALColorInterp = C.GCI_GrayIndex
	GCIPaletteIndex   GDALColorInterp = C.GCI_PaletteIndex
	GCIRedBand        GDALColorInterp = C.GCI_RedBand
	GCIGreenBand      GDALColorInterp = C.GCI_GreenBand
	GCIBlueBand       GDALColorInterp = C.GCI_BlueBand
	GCIAlphaBand      GDALColorInterp = C.GCI_AlphaBand
	GCIHueBand        GDALColorInterp = C.GCI_HueBand
	GCISaturationBand GDALColorInterp = C.GCI_SaturationBand
	GCILightnessBand  GDALColorInterp = C.GCI_LightnessBand
	GCICyanBand       GDALColorInterp = C.GCI_CyanBand
	GCIMagentaBand    GDALColorInterp = C.GCI_MagentaBand
	GCIYellowBand     GDALColorInterp = C.GCI_YellowBand
	GCIBlackBand      GDALColorInterp = C.GCI_BlackBand
	GCIYCbCrYBand     GDALColorInterp = C.GCI_YCbCr_YBand
	GCIYCbCrCbBand    GDALColorInterp = C.GCI_YCbCr_CbBand
	GCIYCbCrCrBand    GDALColorInterp = C.GCI_YCbCr_CrBand
	GCIPanBand        GDALColorInterp = C.GCI_PanBand
	GCICoastalBand    GDALColorInterp = C.GCI_CoastalBand
	GCIRedEdgeBand    GDALColorInterp = C.GCI_RedEdgeBand
	GCINIRBand        GDALColorInterp = C.GCI_NIRBand
	GCISWIRBand       GDALColorInterp = C.GCI_SWIRBand
	GCIMWIRBand       GDALColorInterp = C.GCI_MWIRBand
	GCILWIRBand       GDALColorInterp = C.GCI_LWIRBand
	GCITIRBand        GDALColorInterp = C.GCI_TIRBand
	GCIOtherIRBand    GDALColorInterp = C.GCI_OtherIRBand
	GCIIRReserved1    GDALColorInterp = C.GCI_IR_Reserved_1
	GCIIRReserved2    GDALColorInterp = C.GCI_IR_Reserved_2
	GCIIRReserved3    GDALColorInterp = C.GCI_IR_Reserved_3
	GCIIRReserved4    GDALColorInterp = C.GCI_IR_Reserved_4
	GCISARKaBand      GDALColorInterp = C.GCI_SAR_Ka_Band
	GCISARKBand       GDALColorInterp = C.GCI_SAR_K_Band
	GCISARKuBand      GDALColorInterp = C.GCI_SAR_Ku_Band
	GCISARXBand       GDALColorInterp = C.GCI_SAR_X_Band
	GCISARCBand       GDALColorInterp = C.GCI_SAR_C_Band
	GCISARSBand       GDALColorInterp = C.GCI_SAR_S_Band
	GCISARLBand       GDALColorInterp = C.GCI_SAR_L_Band
	GCISARPBand       GDALColorInterp = C.GCI_SAR_P_Band
	GCISARReserved1   GDALColorInterp = C.GCI_SAR_Reserved_1
	GCISARReserved2   GDALColorInterp = C.GCI_SAR_Reserved_2
	GCIMax            GDALColorInterp = C.GCI_Max
)

func gdalGetColorInterpretationName(colorInterp GDALColorInterp) (result string) {
	result = C.GoString(C.GDALGetColorInterpretationName(C.GDALColorInterp(colorInterp)))
	return
}

func (ci GDALColorInterp) GetName() (result string) {
	result = gdalGetColorInterpretationName(ci)
	return
}

func gdalGetColorInterpretationByName(name string) (result GDALColorInterp) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALColorInterp(C.GDALGetColorInterpretationByName(cName))
	return
}

func GDALGetColorInterpretationByName(name string) (result GDALColorInterp) {
	result = gdalGetColorInterpretationByName(name)
	return
}

// Types of color interpretations for a GDALColorTable.
type GDALPaletteInterp C.GDALPaletteInterp

const (
	GPIGray GDALPaletteInterp = C.GPI_Gray
	GPIRGB  GDALPaletteInterp = C.GPI_RGB
	GPICMYK GDALPaletteInterp = C.GPI_CMYK
	GPIHLS  GDALPaletteInterp = C.GPI_HLS
)

func gdalGetPaletteInterpretationName(paletteInterp GDALPaletteInterp) (result string) {
	result = C.GoString(C.GDALGetPaletteInterpretationName(C.GDALPaletteInterp(paletteInterp)))
	return
}

func (pi GDALPaletteInterp) GetName() (result string) {
	result = gdalGetPaletteInterpretationName(pi)
	return
}

// "well known" metadata items.
var (
	GDALMDAreaOrPoint = C.GoString(C._GDALMD_AREA_OR_POINT)
	GDALMDAOPArea     = C.GoString(C._GDALMD_AOP_AREA)
	GDALMDAOPPoint    = C.GoString(C._GDALMD_AOP_POINT)
)

// GDAL-specific error codes (100 to 299 reserved for GDAL).
const CPLEWrongFormat CPLErrorNum = C.CPLE_WrongFormat

// GSpacing expresses pixel, line or band spacing. Signed 64 bit integer.
type GSpacing C.GSpacing

// Class of a GDALExtendedDataType.
type GDALExtendedDataTypeClass C.GDALExtendedDataTypeClass

const (
	GEDTCNumeric  GDALExtendedDataTypeClass = C.GEDTC_NUMERIC
	GEDTCString   GDALExtendedDataTypeClass = C.GEDTC_STRING
	GEDTCCompound GDALExtendedDataTypeClass = C.GEDTC_COMPOUND
)

// Subtype of a GDALExtendedDataType.
type GDALExtendedDataTypeSubType C.GDALExtendedDataTypeSubType

const (
	GEDTSTNone GDALExtendedDataTypeSubType = C.GEDTST_NONE
	GEDTSTJSON GDALExtendedDataTypeSubType = C.GEDTST_JSON
)

// Driver metadata (GDAL_DMD_*), capability (GDAL_DCAP_*), dimension type
// (GDAL_DIM_TYPE_*) and driver capability (GDsC*) item names.
var (
	GDALDMDLongName                                  = C.GoString(C._GDAL_DMD_LONGNAME)
	GDALDMDHelpTopic                                 = C.GoString(C._GDAL_DMD_HELPTOPIC)
	GDALDMDMimeType                                  = C.GoString(C._GDAL_DMD_MIMETYPE)
	GDALDMDExtension                                 = C.GoString(C._GDAL_DMD_EXTENSION)
	GDALDMDConnectionPrefix                          = C.GoString(C._GDAL_DMD_CONNECTION_PREFIX)
	GDALDMDExtensions                                = C.GoString(C._GDAL_DMD_EXTENSIONS)
	GDALDMDCreationOptionList                        = C.GoString(C._GDAL_DMD_CREATIONOPTIONLIST)
	GDALDMDOverviewCreationOptionList                = C.GoString(C._GDAL_DMD_OVERVIEW_CREATIONOPTIONLIST)
	GDALDMDMultidimDatasetCreationOptionList         = C.GoString(C._GDAL_DMD_MULTIDIM_DATASET_CREATIONOPTIONLIST)
	GDALDMDMultidimGroupCreationOptionList           = C.GoString(C._GDAL_DMD_MULTIDIM_GROUP_CREATIONOPTIONLIST)
	GDALDMDMultidimDimensionCreationOptionList       = C.GoString(C._GDAL_DMD_MULTIDIM_DIMENSION_CREATIONOPTIONLIST)
	GDALDMDMultidimArrayCreationOptionList           = C.GoString(C._GDAL_DMD_MULTIDIM_ARRAY_CREATIONOPTIONLIST)
	GDALDMDMultidimArrayOpenOptionList               = C.GoString(C._GDAL_DMD_MULTIDIM_ARRAY_OPENOPTIONLIST)
	GDALDMDMultidimAttributeCreationOptionList       = C.GoString(C._GDAL_DMD_MULTIDIM_ATTRIBUTE_CREATIONOPTIONLIST)
	GDALDMDOpenOptionList                            = C.GoString(C._GDAL_DMD_OPENOPTIONLIST)
	GDALDMDCreationDataTypes                         = C.GoString(C._GDAL_DMD_CREATIONDATATYPES)
	GDALDMDCreationFieldDataTypes                    = C.GoString(C._GDAL_DMD_CREATIONFIELDDATATYPES)
	GDALDMDCreationFieldDataSubTypes                 = C.GoString(C._GDAL_DMD_CREATIONFIELDDATASUBTYPES)
	GDALDMDMaxStringLength                           = C.GoString(C._GDAL_DMD_MAX_STRING_LENGTH)
	GDALDMDCreationFieldDefnFlags                    = C.GoString(C._GDAL_DMD_CREATION_FIELD_DEFN_FLAGS)
	GDALDMDSubdatasets                               = C.GoString(C._GDAL_DMD_SUBDATASETS)
	GDALDCAPCreateSubdatasets                        = C.GoString(C._GDAL_DCAP_CREATE_SUBDATASETS)
	GDALDMDNumericFieldWidthIncludesDecimalSeparator = C.GoString(C._GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_DECIMAL_SEPARATOR)
	GDALDMDNumericFieldWidthIncludesSign             = C.GoString(C._GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_SIGN)
	GDALDCAPOpen                                     = C.GoString(C._GDAL_DCAP_OPEN)
	GDALDCAPCreate                                   = C.GoString(C._GDAL_DCAP_CREATE)
	GDALDCAPCreateMultidimensional                   = C.GoString(C._GDAL_DCAP_CREATE_MULTIDIMENSIONAL)
	GDALDCAPCreateCopy                               = C.GoString(C._GDAL_DCAP_CREATECOPY)
	GDALDCAPCreateOnlyVisibleAtCloseTime             = C.GoString(C._GDAL_DCAP_CREATE_ONLY_VISIBLE_AT_CLOSE_TIME)
	GDALDCAPVectorTranslateFrom                      = C.GoString(C._GDAL_DCAP_VECTOR_TRANSLATE_FROM)
	GDALDCAPCreateCopyMultidimensional               = C.GoString(C._GDAL_DCAP_CREATECOPY_MULTIDIMENSIONAL)
	GDALDCAPMultidimRaster                           = C.GoString(C._GDAL_DCAP_MULTIDIM_RASTER)
	GDALDCAPSubCreateCopy                            = C.GoString(C._GDAL_DCAP_SUBCREATECOPY)
	GDALDCAPAppend                                   = C.GoString(C._GDAL_DCAP_APPEND)
	GDALDCAPUpdate                                   = C.GoString(C._GDAL_DCAP_UPDATE)
	GDALDCAPVirtualIO                                = C.GoString(C._GDAL_DCAP_VIRTUALIO)
	GDALDCAPRaster                                   = C.GoString(C._GDAL_DCAP_RASTER)
	GDALDCAPVector                                   = C.GoString(C._GDAL_DCAP_VECTOR)
	GDALDCAPGNM                                      = C.GoString(C._GDAL_DCAP_GNM)
	GDALDCAPCreateLayer                              = C.GoString(C._GDAL_DCAP_CREATE_LAYER)
	GDALDCAPDeleteLayer                              = C.GoString(C._GDAL_DCAP_DELETE_LAYER)
	GDALDCAPCreateField                              = C.GoString(C._GDAL_DCAP_CREATE_FIELD)
	GDALDCAPDeleteField                              = C.GoString(C._GDAL_DCAP_DELETE_FIELD)
	GDALDCAPReorderFields                            = C.GoString(C._GDAL_DCAP_REORDER_FIELDS)
	GDALDMDAlterFieldDefnFlags                       = C.GoString(C._GDAL_DMD_ALTER_FIELD_DEFN_FLAGS)
	GDALDMDIllegalFieldNames                         = C.GoString(C._GDAL_DMD_ILLEGAL_FIELD_NAMES)
	GDALDCAPNotNullFields                            = C.GoString(C._GDAL_DCAP_NOTNULL_FIELDS)
	GDALDCAPUniqueFields                             = C.GoString(C._GDAL_DCAP_UNIQUE_FIELDS)
	GDALDCAPDefaultFields                            = C.GoString(C._GDAL_DCAP_DEFAULT_FIELDS)
	GDALDCAPNotNullGeomFields                        = C.GoString(C._GDAL_DCAP_NOTNULL_GEOMFIELDS)
	GDALDCAPNonspatial                               = C.GoString(C._GDAL_DCAP_NONSPATIAL)
	GDALDCAPCurveGeometries                          = C.GoString(C._GDAL_DCAP_CURVE_GEOMETRIES)
	GDALDCAPMeasuredGeometries                       = C.GoString(C._GDAL_DCAP_MEASURED_GEOMETRIES)
	GDALDCAPZGeometries                              = C.GoString(C._GDAL_DCAP_Z_GEOMETRIES)
	GDALDMDGeometryFlags                             = C.GoString(C._GDAL_DMD_GEOMETRY_FLAGS)
	GDALDCAPFeatureStyles                            = C.GoString(C._GDAL_DCAP_FEATURE_STYLES)
	GDALDCAPFeatureStylesRead                        = C.GoString(C._GDAL_DCAP_FEATURE_STYLES_READ)
	GDALDCAPFeatureStylesWrite                       = C.GoString(C._GDAL_DCAP_FEATURE_STYLES_WRITE)
	GDALDCAPCoordinateEpoch                          = C.GoString(C._GDAL_DCAP_COORDINATE_EPOCH)
	GDALDCAPMultipleVectorLayers                     = C.GoString(C._GDAL_DCAP_MULTIPLE_VECTOR_LAYERS)
	GDALDCAPFieldDomains                             = C.GoString(C._GDAL_DCAP_FIELD_DOMAINS)
	GDALDCAPRelationships                            = C.GoString(C._GDAL_DCAP_RELATIONSHIPS)
	GDALDCAPCreateRelationship                       = C.GoString(C._GDAL_DCAP_CREATE_RELATIONSHIP)
	GDALDCAPDeleteRelationship                       = C.GoString(C._GDAL_DCAP_DELETE_RELATIONSHIP)
	GDALDCAPUpdateRelationship                       = C.GoString(C._GDAL_DCAP_UPDATE_RELATIONSHIP)
	GDALDCAPFlushCacheConsistentState                = C.GoString(C._GDAL_DCAP_FLUSHCACHE_CONSISTENT_STATE)
	GDALDCAPHonorGeomCoordinatePrecision             = C.GoString(C._GDAL_DCAP_HONOR_GEOM_COORDINATE_PRECISION)
	GDALDCAPUpsert                                   = C.GoString(C._GDAL_DCAP_UPSERT)
	GDALDMDRelationshipFlags                         = C.GoString(C._GDAL_DMD_RELATIONSHIP_FLAGS)
	GDALDMDRelationshipRelatedTableTypes             = C.GoString(C._GDAL_DMD_RELATIONSHIP_RELATED_TABLE_TYPES)
	GDALDCAPRenameLayers                             = C.GoString(C._GDAL_DCAP_RENAME_LAYERS)
	GDALDMDCreationFieldDomainTypes                  = C.GoString(C._GDAL_DMD_CREATION_FIELD_DOMAIN_TYPES)
	GDALDMDAlterGeomFieldDefnFlags                   = C.GoString(C._GDAL_DMD_ALTER_GEOM_FIELD_DEFN_FLAGS)
	GDALDMDSupportedSQLDialects                      = C.GoString(C._GDAL_DMD_SUPPORTED_SQL_DIALECTS)
	GDALDMDPluginInstallationMessage                 = C.GoString(C._GDAL_DMD_PLUGIN_INSTALLATION_MESSAGE)
	GDALDMDUpdateItems                               = C.GoString(C._GDAL_DMD_UPDATE_ITEMS)
	GDALDimTypeHorizontalX                           = C.GoString(C._GDAL_DIM_TYPE_HORIZONTAL_X)
	GDALDimTypeHorizontalY                           = C.GoString(C._GDAL_DIM_TYPE_HORIZONTAL_Y)
	GDALDimTypeVertical                              = C.GoString(C._GDAL_DIM_TYPE_VERTICAL)
	GDALDimTypeTemporal                              = C.GoString(C._GDAL_DIM_TYPE_TEMPORAL)
	GDALDimTypeParametric                            = C.GoString(C._GDAL_DIM_TYPE_PARAMETRIC)
	GDALDCAPReopenAfterWriteRequired                 = C.GoString(C._GDAL_DCAP_REOPEN_AFTER_WRITE_REQUIRED)
	GDALDCAPCanReadAfterDelete                       = C.GoString(C._GDAL_DCAP_CAN_READ_AFTER_DELETE)
	GDsCAddRelationship                              = C.GoString(C._GDsCAddRelationship)
	GDsCDeleteRelationship                           = C.GoString(C._GDsCDeleteRelationship)
	GDsCUpdateRelationship                           = C.GoString(C._GDsCUpdateRelationship)
	GDsCFastGetExtent                                = C.GoString(C._GDsCFastGetExtent)
	GDsCFastGetExtentWGS84LongLat                    = C.GoString(C._GDsCFastGetExtentWGS84LongLat)
)

func gdalAllRegister() {
	C.GDALAllRegister()
}

func GDALAllRegister() {
	gdalAllRegister()
}

func gdalRegisterPlugins() {
	C.GDALRegisterPlugins()
}

func GDALRegisterPlugins() {
	gdalRegisterPlugins()
}

func gdalRegisterPlugin(name string) (result CPLErr) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = CPLErr(C.GDALRegisterPlugin(cName))
	return
}

func GDALRegisterPlugin(name string) (err error) {
	err = cplErr(gdalRegisterPlugin(name))
	return
}

func gdalCreate(driver GDALDriver, name string, xSize, ySize, bands int, dataType GDALDataType, options []string) (result GDALDataset) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = GDALDataset{cValue: C.GDALCreate(driver.cValue, cName, C.int(xSize), C.int(ySize), C.int(bands), C.GDALDataType(dataType), C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (d GDALDriver) Create(name string, xSize, ySize, bands int, dataType GDALDataType, options []string) (result GDALDataset, err error) {
	result = gdalCreate(d, name, xSize, ySize, bands, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalCreateCopy(driver GDALDriver, name string, src GDALDataset, strict int, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (result GDALDataset) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = GDALDataset{cValue: C.GDALCreateCopy(driver.cValue, cName, src.cValue, C.int(strict), C.CSLConstList(unsafe.Pointer(opts)), progress.cValue, progressData)}
	return
}

func (d GDALDriver) CreateCopy(name string, src GDALDataset, strict int, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (result GDALDataset, err error) {
	result = gdalCreateCopy(d, name, src, strict, options, progress, progressData)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalIdentifyDriver(filename string, fileList []string) (result GDALDriver) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	files, free := cStrings(fileList)
	defer free()
	result = GDALDriver{cValue: C.GDALIdentifyDriver(cName, C.CSLConstList(unsafe.Pointer(files)))}
	return
}

func GDALIdentifyDriver(filename string, fileList []string) (result GDALDriver, err error) {
	result = gdalIdentifyDriver(filename, fileList)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalIdentifyDriverEx(filename string, identifyFlags uint, allowedDrivers, fileList []string) (result GDALDriver) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	allowed, freeA := cStrings(allowedDrivers)
	defer freeA()
	files, freeF := cStrings(fileList)
	defer freeF()
	result = GDALDriver{cValue: C.GDALIdentifyDriverEx(cName, C.uint(identifyFlags), allowed, files)}
	return
}

func GDALIdentifyDriverEx(filename string, identifyFlags uint, allowedDrivers, fileList []string) (result GDALDriver, err error) {
	result = gdalIdentifyDriverEx(filename, identifyFlags, allowedDrivers, fileList)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalOpen(filename string, access GDALAccess) (result GDALDataset) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDataset{cValue: C.GDALOpen(cName, C.GDALAccess(access))}
	return
}

func GDALOpen(filename string, access GDALAccess) (result GDALDataset, err error) {
	result = gdalOpen(filename, access)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalOpenShared(filename string, access GDALAccess) (result GDALDataset) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDataset{cValue: C.GDALOpenShared(cName, C.GDALAccess(access))}
	return
}

func GDALOpenShared(filename string, access GDALAccess) (result GDALDataset, err error) {
	result = gdalOpenShared(filename, access)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// Open flags for GDALOpenEx().
type GDALOpenFlag uint

const (
	GDALOfReadonly           GDALOpenFlag = C.GDAL_OF_READONLY
	GDALOfUpdate             GDALOpenFlag = C.GDAL_OF_UPDATE
	GDALOfAll                GDALOpenFlag = C.GDAL_OF_ALL
	GDALOfRaster             GDALOpenFlag = C.GDAL_OF_RASTER
	GDALOfVector             GDALOpenFlag = C.GDAL_OF_VECTOR
	GDALOfGNM                GDALOpenFlag = C.GDAL_OF_GNM
	GDALOfMultidimRaster     GDALOpenFlag = C.GDAL_OF_MULTIDIM_RASTER
	GDALOfKindMask           GDALOpenFlag = C.GDAL_OF_KIND_MASK
	GDALOfShared             GDALOpenFlag = C.GDAL_OF_SHARED
	GDALOfVerboseError       GDALOpenFlag = C.GDAL_OF_VERBOSE_ERROR
	GDALOfInternal           GDALOpenFlag = C.GDAL_OF_INTERNAL
	GDALOfDefaultBlockAccess GDALOpenFlag = C.GDAL_OF_DEFAULT_BLOCK_ACCESS
	GDALOfArrayBlockAccess   GDALOpenFlag = C.GDAL_OF_ARRAY_BLOCK_ACCESS
	GDALOfHashsetBlockAccess GDALOpenFlag = C.GDAL_OF_HASHSET_BLOCK_ACCESS
	GDALOfReserved1          GDALOpenFlag = C.GDAL_OF_RESERVED_1
	GDALOfBlockAccessMask    GDALOpenFlag = C.GDAL_OF_BLOCK_ACCESS_MASK
	GDALOfFromGDALOpen       GDALOpenFlag = C.GDAL_OF_FROM_GDALOPEN
	GDALOfThreadSafe         GDALOpenFlag = C.GDAL_OF_THREAD_SAFE
)

func gdalOpenEx(filename string, openFlags GDALOpenFlag, allowedDrivers, openOptions, siblingFiles []string) (result GDALDataset) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	allowed, freeA := cStrings(allowedDrivers)
	defer freeA()
	options, freeO := cStrings(openOptions)
	defer freeO()
	siblings, freeS := cStrings(siblingFiles)
	defer freeS()
	result = GDALDataset{cValue: C.GDALOpenEx(cName, C.uint(openFlags), allowed, options, siblings)}
	return
}

func GDALOpenEx(filename string, openFlags GDALOpenFlag, allowedDrivers, openOptions, siblingFiles []string) (result GDALDataset, err error) {
	result = gdalOpenEx(filename, openFlags, allowedDrivers, openOptions, siblingFiles)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDumpOpenDatasets(filename string) (result int, err error) {
	fp, closeFn, err := cFile(filename, "w")
	if err != nil {
		return
	}
	defer closeFn()
	result = int(C.GDALDumpOpenDatasets(fp))
	return
}

func GDALDumpOpenDatasets(filename string) (result int, err error) {
	return gdalDumpOpenDatasets(filename)
}

func gdalGetDriverByName(name string) (result GDALDriver) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDriver{cValue: C.GDALGetDriverByName(cName)}
	return
}

func GDALGetDriverByName(name string) (result GDALDriver, err error) {
	result = gdalGetDriverByName(name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetDriverCount() (result int) {
	result = int(C.GDALGetDriverCount())
	return
}

func GDALGetDriverCount() (result int) {
	result = gdalGetDriverCount()
	return
}

func gdalGetDriver(index int) (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALGetDriver(C.int(index))}
	return
}

func GDALGetDriver(index int) (result GDALDriver, err error) {
	result = gdalGetDriver(index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalCreateDriver() (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALCreateDriver()}
	return
}

func GDALCreateDriver() (result GDALDriver, err error) {
	result = gdalCreateDriver()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDestroyDriver(driver GDALDriver) {
	C.GDALDestroyDriver(driver.cValue)
}

func (d GDALDriver) Destroy() {
	gdalDestroyDriver(d)
}

func gdalRegisterDriver(driver GDALDriver) (result int) {
	result = int(C.GDALRegisterDriver(driver.cValue))
	return
}

func (d GDALDriver) Register() (result int) {
	result = gdalRegisterDriver(d)
	return
}

func gdalDeregisterDriver(driver GDALDriver) {
	C.GDALDeregisterDriver(driver.cValue)
}

func (d GDALDriver) Deregister() {
	gdalDeregisterDriver(d)
}

func gdalDestroyDriverManager() {
	C.GDALDestroyDriverManager()
}

func GDALDestroyDriverManager() {
	gdalDestroyDriverManager()
}

func gdalDestroy() {
	C.GDALDestroy()
}

func GDALDestroy() {
	gdalDestroy()
}

func gdalDeleteDataset(driver GDALDriver, filename string) (result CPLErr) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	result = CPLErr(C.GDALDeleteDataset(driver.cValue, cName))
	return
}

func (d GDALDriver) DeleteDataset(filename string) (err error) {
	err = cplErr(gdalDeleteDataset(d, filename))
	return
}

func gdalRenameDataset(driver GDALDriver, newName, oldName string) (result CPLErr) {
	cNew := C.CString(newName)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOld))
	result = CPLErr(C.GDALRenameDataset(driver.cValue, cNew, cOld))
	return
}

func (d GDALDriver) RenameDataset(newName, oldName string) (err error) {
	err = cplErr(gdalRenameDataset(d, newName, oldName))
	return
}

func gdalCopyDatasetFiles(driver GDALDriver, newName, oldName string) (result CPLErr) {
	cNew := C.CString(newName)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOld))
	result = CPLErr(C.GDALCopyDatasetFiles(driver.cValue, cNew, cOld))
	return
}

func (d GDALDriver) CopyDatasetFiles(newName, oldName string) (err error) {
	err = cplErr(gdalCopyDatasetFiles(d, newName, oldName))
	return
}

func gdalValidateCreationOptions(driver GDALDriver, options []string) (result bool) {
	opts, free := cStrings(options)
	defer free()
	result = C.GDALValidateCreationOptions(driver.cValue, C.CSLConstList(unsafe.Pointer(opts))) != 0
	return
}

func (d GDALDriver) ValidateCreationOptions(options []string) (result bool) {
	result = gdalValidateCreationOptions(d, options)
	return
}

func gdalGetOutputDriversForDatasetName(destFilename string, flagRasterVector int, singleMatch, emitWarning bool) (result []string) {
	cName := C.CString(destFilename)
	defer C.free(unsafe.Pointer(cName))
	raw := C.GDALGetOutputDriversForDatasetName(cName, C.int(flagRasterVector), C.bool(singleMatch), C.bool(emitWarning))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func GDALGetOutputDriversForDatasetName(destFilename string, flagRasterVector int, singleMatch, emitWarning bool) (result []string) {
	result = gdalGetOutputDriversForDatasetName(destFilename, flagRasterVector, singleMatch, emitWarning)
	return
}

func gdalDriverHasOpenOption(driver GDALDriver, openOptionName string) (result bool) {
	cName := C.CString(openOptionName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALDriverHasOpenOption(driver.cValue, cName))
	return
}

func (d GDALDriver) HasOpenOption(openOptionName string) (result bool) {
	result = gdalDriverHasOpenOption(d, openOptionName)
	return
}

// The following are deprecated.

func gdalGetDriverShortName(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverShortName(driver.cValue))
	return
}

func (d GDALDriver) GetShortName() (result string) {
	result = gdalGetDriverShortName(d)
	return
}

func gdalGetDriverLongName(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverLongName(driver.cValue))
	return
}

func (d GDALDriver) GetLongName() (result string) {
	result = gdalGetDriverLongName(d)
	return
}

func gdalGetDriverHelpTopic(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverHelpTopic(driver.cValue))
	return
}

func (d GDALDriver) GetHelpTopic() (result string) {
	result = gdalGetDriverHelpTopic(d)
	return
}

func gdalGetDriverCreationOptionList(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverCreationOptionList(driver.cValue))
	return
}

func (d GDALDriver) GetCreationOptionList() (result string) {
	result = gdalGetDriverCreationOptionList(d)
	return
}

// /* ==================================================================== */
// /*      GDAL_GCP                                                        */
// /* ==================================================================== */

// GDALGCP is a Ground Control Point.
type GDALGCP struct {
	cValue C.GDAL_GCP
}

// GDALGCPs is a contiguous list of Ground Control Points, matching a C
// GDAL_GCP array (its length carries the C nGCPCount argument).
type GDALGCPs []GDALGCP

// cPtr returns a pointer to the first C GDAL_GCP, or nil for an empty list.
// GDALGCP is a single-field struct, so the slice is a contiguous C array.
func (g GDALGCPs) cPtr() *C.GDAL_GCP {
	if len(g) == 0 {
		return nil
	}
	return &g[0].cValue
}

func gdalInitGCPs(count int, gcps GDALGCPs) {
	C.GDALInitGCPs(C.int(count), gcps.cPtr())
}

// GDALInitGCPs allocates and initializes a list of count Ground Control Points.
func GDALInitGCPs(count int) (result GDALGCPs) {
	result = make(GDALGCPs, count)
	gdalInitGCPs(count, result)
	return
}

func gdalDeinitGCPs(count int, gcps GDALGCPs) {
	C.GDALDeinitGCPs(C.int(count), gcps.cPtr())
}

func (g GDALGCPs) Deinit() {
	gdalDeinitGCPs(len(g), g)
}

func gdalDuplicateGCPs(count int, gcps GDALGCPs) (result GDALGCPs) {
	dup := C.GDALDuplicateGCPs(C.int(count), gcps.cPtr())
	if dup == nil {
		return
	}
	result = unsafe.Slice((*GDALGCP)(unsafe.Pointer(dup)), count)
	return
}

func (g GDALGCPs) Duplicate() (result GDALGCPs) {
	result = gdalDuplicateGCPs(len(g), g)
	return
}

func gdalGCPsToGeoTransform(count int, gcps GDALGCPs, geoTransform *[6]float64, approxOK int) int {
	var gt [6]C.double
	r := C.GDALGCPsToGeoTransform(C.int(count), gcps.cPtr(), &gt[0], C.int(approxOK))
	for i := range gt {
		geoTransform[i] = float64(gt[i])
	}
	return int(r)
}

func (g GDALGCPs) ToGeoTransform(approxOK int) (geoTransform [6]float64, ok bool) {
	ok = gdalGCPsToGeoTransform(len(g), g, &geoTransform, approxOK) != 0
	return
}

func gdalInvGeoTransform(geoTransform [6]float64, result *[6]float64) int {
	var in, out [6]C.double
	for i, v := range geoTransform {
		in[i] = C.double(v)
	}
	r := C.GDALInvGeoTransform(&in[0], &out[0])
	for i := range out {
		result[i] = float64(out[i])
	}
	return int(r)
}

func GDALInvGeoTransform(geoTransform [6]float64) (result [6]float64, ok bool) {
	ok = gdalInvGeoTransform(geoTransform, &result) != 0
	return
}

func gdalApplyGeoTransform(geoTransform [6]float64, pixel, line float64, geoX, geoY *float64) {
	var gt [6]C.double
	for i, v := range geoTransform {
		gt[i] = C.double(v)
	}
	var x, y C.double
	C.GDALApplyGeoTransform(&gt[0], C.double(pixel), C.double(line), &x, &y)
	*geoX = float64(x)
	*geoY = float64(y)
}

func GDALApplyGeoTransform(geoTransform [6]float64, pixel, line float64) (geoX, geoY float64) {
	gdalApplyGeoTransform(geoTransform, pixel, line, &geoX, &geoY)
	return
}

func gdalComposeGeoTransforms(a, b [6]float64, result *[6]float64) {
	var ca, cb, out [6]C.double
	for i := range a {
		ca[i] = C.double(a[i])
		cb[i] = C.double(b[i])
	}
	C.GDALComposeGeoTransforms(&ca[0], &cb[0], &out[0])
	for i := range out {
		result[i] = float64(out[i])
	}
}

func GDALComposeGeoTransforms(a, b [6]float64) (result [6]float64) {
	gdalComposeGeoTransforms(a, b, &result)
	return
}

func gdalGCPsToHomography(count int, gcps GDALGCPs, homography *[9]float64) int {
	var h [9]C.double
	r := C.GDALGCPsToHomography(C.int(count), gcps.cPtr(), &h[0])
	for i := range h {
		homography[i] = float64(h[i])
	}
	return int(r)
}

func (g GDALGCPs) ToHomography() (homography [9]float64, ok bool) {
	ok = gdalGCPsToHomography(len(g), g, &homography) != 0
	return
}

func gdalInvHomography(homography [9]float64, result *[9]float64) int {
	var in, out [9]C.double
	for i, v := range homography {
		in[i] = C.double(v)
	}
	r := C.GDALInvHomography(&in[0], &out[0])
	for i := range out {
		result[i] = float64(out[i])
	}
	return int(r)
}

func GDALInvHomography(homography [9]float64) (result [9]float64, ok bool) {
	ok = gdalInvHomography(homography, &result) != 0
	return
}

func gdalApplyHomography(homography [9]float64, x, y float64, outX, outY *float64) int {
	var h [9]C.double
	for i, v := range homography {
		h[i] = C.double(v)
	}
	var cx, cy C.double
	r := C.GDALApplyHomography(&h[0], C.double(x), C.double(y), &cx, &cy)
	*outX = float64(cx)
	*outY = float64(cy)
	return int(r)
}

func GDALApplyHomography(homography [9]float64, x, y float64) (outX, outY float64, ok bool) {
	ok = gdalApplyHomography(homography, x, y, &outX, &outY) != 0
	return
}

func gdalComposeHomographies(a, b [9]float64, result *[9]float64) {
	var ca, cb, out [9]C.double
	for i := range a {
		ca[i] = C.double(a[i])
		cb[i] = C.double(b[i])
	}
	C.GDALComposeHomographies(&ca[0], &cb[0], &out[0])
	for i := range out {
		result[i] = float64(out[i])
	}
}

func GDALComposeHomographies(a, b [9]float64) (result [9]float64) {
	gdalComposeHomographies(a, b, &result)
	return
}

// /* ==================================================================== */
// /*      major objects (dataset, and, driver, drivermanager).            */
// /* ==================================================================== */

func gdalGetMetadataDomainList(object GDALMajorObject) (result []string) {
	raw := C.GDALGetMetadataDomainList(object.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (o GDALMajorObject) GetMetadataDomainList() (result []string) {
	result = gdalGetMetadataDomainList(o)
	return
}

func gdalGetMetadata(object GDALMajorObject, domain string) (result []string) {
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	result = goStrings(C.GDALGetMetadata(object.cValue, cDomain))
	return
}

func (o GDALMajorObject) GetMetadata(domain string) (result []string) {
	result = gdalGetMetadata(o, domain)
	return
}

func gdalSetMetadata(object GDALMajorObject, metadata []string, domain string) (result CPLErr) {
	md, free := cStrings(metadata)
	defer free()
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	result = CPLErr(C.GDALSetMetadata(object.cValue, C.CSLConstList(unsafe.Pointer(md)), cDomain))
	return
}

func (o GDALMajorObject) SetMetadata(metadata []string, domain string) (err error) {
	err = cplErr(gdalSetMetadata(o, metadata, domain))
	return
}

func gdalGetMetadataItem(object GDALMajorObject, name, domain string) (result string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	result = C.GoString(C.GDALGetMetadataItem(object.cValue, cName, cDomain))
	return
}

func (o GDALMajorObject) GetMetadataItem(name, domain string) (result string) {
	result = gdalGetMetadataItem(o, name, domain)
	return
}

func gdalSetMetadataItem(object GDALMajorObject, name, value, domain string) (result CPLErr) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	result = CPLErr(C.GDALSetMetadataItem(object.cValue, cName, cValue, cDomain))
	return
}

func (o GDALMajorObject) SetMetadataItem(name, value, domain string) (err error) {
	err = cplErr(gdalSetMetadataItem(o, name, value, domain))
	return
}

func gdalGetDescription(object GDALMajorObject) (result string) {
	result = C.GoString(C.GDALGetDescription(object.cValue))
	return
}

func (o GDALMajorObject) GetDescription() (result string) {
	result = gdalGetDescription(o)
	return
}

func gdalSetDescription(object GDALMajorObject, description string) {
	cDesc := C.CString(description)
	defer C.free(unsafe.Pointer(cDesc))
	C.GDALSetDescription(object.cValue, cDesc)
}

func (o GDALMajorObject) SetDescription(description string) {
	gdalSetDescription(o, description)
}

// /* ==================================================================== */
// /*      GDALDataset class ... normally this represents one file.        */
// /* ==================================================================== */

// Name of driver metadata item for layer creation option list.
var GDALDSLayerCreationOptionList = C.GoString(C._GDAL_DS_LAYER_CREATIONOPTIONLIST)

func gdalGetDatasetDriver(dataset GDALDataset) (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALGetDatasetDriver(dataset.cValue)}
	return
}

func (ds GDALDataset) GetDriver() (result GDALDriver) {
	result = gdalGetDatasetDriver(ds)
	return
}

func gdalGetFileList(dataset GDALDataset) (result []string) {
	raw := C.GDALGetFileList(dataset.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (ds GDALDataset) GetFileList() (result []string) {
	result = gdalGetFileList(ds)
	return
}

func gdalDatasetMarkSuppressOnClose(dataset GDALDataset) {
	C.GDALDatasetMarkSuppressOnClose(dataset.cValue)
}

func (ds GDALDataset) MarkSuppressOnClose() {
	gdalDatasetMarkSuppressOnClose(ds)
}

func gdalClose(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALClose(dataset.cValue))
	return
}

func (ds GDALDataset) Close() (err error) {
	err = cplErr(gdalClose(ds))
	return
}

func gdalDatasetRunCloseWithoutDestroying(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALDatasetRunCloseWithoutDestroying(dataset.cValue))
	return
}

func (ds GDALDataset) RunCloseWithoutDestroying() (err error) {
	err = cplErr(gdalDatasetRunCloseWithoutDestroying(ds))
	return
}

func gdalGetRasterXSize(dataset GDALDataset) (result int) {
	result = int(C.GDALGetRasterXSize(dataset.cValue))
	return
}

func (ds GDALDataset) GetRasterXSize() (result int) {
	result = gdalGetRasterXSize(ds)
	return
}

func gdalGetRasterYSize(dataset GDALDataset) (result int) {
	result = int(C.GDALGetRasterYSize(dataset.cValue))
	return
}

func (ds GDALDataset) GetRasterYSize() (result int) {
	result = gdalGetRasterYSize(ds)
	return
}

func gdalGetRasterCount(dataset GDALDataset) (result int) {
	result = int(C.GDALGetRasterCount(dataset.cValue))
	return
}

func (ds GDALDataset) GetRasterCount() (result int) {
	result = gdalGetRasterCount(ds)
	return
}

func gdalGetRasterBand(dataset GDALDataset, band int) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetRasterBand(dataset.cValue, C.int(band))}
	return
}

func (ds GDALDataset) GetRasterBand(band int) (result GDALRasterBand, err error) {
	result = gdalGetRasterBand(ds, band)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetIsThreadSafe(dataset GDALDataset, scopeFlags int, options []string) (result bool) {
	opts, free := cStrings(options)
	defer free()
	result = bool(C.GDALDatasetIsThreadSafe(dataset.cValue, C.int(scopeFlags), C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (ds GDALDataset) IsThreadSafe(scopeFlags int, options []string) (result bool) {
	result = gdalDatasetIsThreadSafe(ds, scopeFlags, options)
	return
}

func gdalGetThreadSafeDataset(dataset GDALDataset, scopeFlags int, options []string) (result GDALDataset) {
	opts, free := cStrings(options)
	defer free()
	result = GDALDataset{cValue: C.GDALGetThreadSafeDataset(dataset.cValue, C.int(scopeFlags), C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (ds GDALDataset) GetThreadSafeDataset(scopeFlags int, options []string) (result GDALDataset, err error) {
	result = gdalGetThreadSafeDataset(ds, scopeFlags, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalAddBand(dataset GDALDataset, dataType GDALDataType, options []string) (result CPLErr) {
	opts, free := cStrings(options)
	defer free()
	result = CPLErr(C.GDALAddBand(dataset.cValue, C.GDALDataType(dataType), C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (ds GDALDataset) AddBand(dataType GDALDataType, options []string) (err error) {
	err = cplErr(gdalAddBand(ds, dataType, options))
	return
}

func gdalBeginAsyncReader(dataset GDALDataset, xOff, yOff, xSize, ySize int, buf unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandMap []int, pixelSpace, lineSpace, bandSpace int, options []string) (result GDALAsyncReader) {
	opts, free := cStrings(options)
	defer free()
	result = GDALAsyncReader{cValue: C.GDALBeginAsyncReader(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buf, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandMap), C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace), C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (ds GDALDataset) BeginAsyncReader(xOff, yOff, xSize, ySize int, buf []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandMap []int, pixelSpace, lineSpace, bandSpace int, options []string) (result GDALAsyncReader, err error) {
	result = gdalBeginAsyncReader(ds, xOff, yOff, xSize, ySize, cBytes(buf), bufXSize, bufYSize, bufType, bandCount, bandMap, pixelSpace, lineSpace, bandSpace, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalEndAsyncReader(dataset GDALDataset, reader GDALAsyncReader) {
	C.GDALEndAsyncReader(dataset.cValue, reader.cValue)
}

func (ds GDALDataset) EndAsyncReader(reader GDALAsyncReader) {
	gdalEndAsyncReader(ds, reader)
}

func gdalDatasetRasterIO(dataset GDALDataset, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int) (result CPLErr) {
	result = CPLErr(C.GDALDatasetRasterIO(dataset.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandList), C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace)))
	return
}

func (ds GDALDataset) RasterIO(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int) (err error) {
	err = cplErr(gdalDatasetRasterIO(ds, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, bandCount, bandList, pixelSpace, lineSpace, bandSpace))
	return
}

func gdalDatasetRasterIOEx(dataset GDALDataset, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int64, extraArg GDALRasterIOExtraArg) (result CPLErr) {
	result = CPLErr(C.GDALDatasetRasterIOEx(dataset.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandList), C.GSpacing(pixelSpace), C.GSpacing(lineSpace), C.GSpacing(bandSpace), extraArg.cValue))
	return
}

func (ds GDALDataset) RasterIOEx(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int64, extraArg GDALRasterIOExtraArg) (err error) {
	err = cplErr(gdalDatasetRasterIOEx(ds, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, bandCount, bandList, pixelSpace, lineSpace, bandSpace, extraArg))
	return
}

func gdalDatasetAdviseRead(dataset GDALDataset, xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, options []string) (result CPLErr) {
	opts, free := cStrings(options)
	defer free()
	result = CPLErr(C.GDALDatasetAdviseRead(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandList), C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (ds GDALDataset) AdviseRead(xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, options []string) (err error) {
	err = cplErr(gdalDatasetAdviseRead(ds, xOff, yOff, xSize, ySize, bufXSize, bufYSize, bufType, bandCount, bandList, options))
	return
}

func gdalDatasetGetCompressionFormats(dataset GDALDataset, xOff, yOff, xSize, ySize, bandCount int, bandList []int) (result []string) {
	raw := C.GDALDatasetGetCompressionFormats(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bandCount), cInts(bandList))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (ds GDALDataset) GetCompressionFormats(xOff, yOff, xSize, ySize, bandCount int, bandList []int) (result []string) {
	result = gdalDatasetGetCompressionFormats(ds, xOff, yOff, xSize, ySize, bandCount, bandList)
	return
}

func gdalDatasetReadCompressedData(dataset GDALDataset, format string, xOff, yOff, xSize, ySize, bandCount int, bandList []int, buffer *unsafe.Pointer, bufferSize *int, detailedFormat *string) (result CPLErr) {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	var cBuffer unsafe.Pointer
	var size C.size_t
	var cDetailed *C.char
	result = CPLErr(C.GDALDatasetReadCompressedData(dataset.cValue, cFormat, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bandCount), cInts(bandList), &cBuffer, &size, &cDetailed))
	if buffer != nil {
		*buffer = cBuffer
	}
	if bufferSize != nil {
		*bufferSize = int(size)
	}
	if cDetailed != nil {
		if detailedFormat != nil {
			*detailedFormat = C.GoString(cDetailed)
		}
		C.VSIFree(unsafe.Pointer(cDetailed))
	}
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
		buffer = C.GoBytes(cBuffer, C.int(size))
		C.VSIFree(cBuffer)
	}
	return
}

func gdalGetProjectionRef(dataset GDALDataset) (result string) {
	result = C.GoString(C.GDALGetProjectionRef(dataset.cValue))
	return
}

func (ds GDALDataset) GetProjectionRef() (result string) {
	result = gdalGetProjectionRef(ds)
	return
}

func gdalGetSpatialRef(dataset GDALDataset) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.GDALGetSpatialRef(dataset.cValue)}
	return
}

func (ds GDALDataset) GetSpatialRef() (result OGRSpatialReference) {
	result = gdalGetSpatialRef(ds)
	return
}

func gdalSetProjection(dataset GDALDataset, projection string) (result CPLErr) {
	cProj := C.CString(projection)
	defer C.free(unsafe.Pointer(cProj))
	result = CPLErr(C.GDALSetProjection(dataset.cValue, cProj))
	return
}

func (ds GDALDataset) SetProjection(projection string) (err error) {
	err = cplErr(gdalSetProjection(ds, projection))
	return
}

func gdalSetSpatialRef(dataset GDALDataset, srs OGRSpatialReference) (result CPLErr) {
	result = CPLErr(C.GDALSetSpatialRef(dataset.cValue, srs.cValue))
	return
}

func (ds GDALDataset) SetSpatialRef(srs OGRSpatialReference) (err error) {
	err = cplErr(gdalSetSpatialRef(ds, srs))
	return
}

func gdalGetGeoTransform(dataset GDALDataset, geoTransform *[6]float64) (result CPLErr) {
	var gt [6]C.double
	result = CPLErr(C.GDALGetGeoTransform(dataset.cValue, &gt[0]))
	for i := range gt {
		geoTransform[i] = float64(gt[i])
	}
	return
}

func (ds GDALDataset) GetGeoTransform() (geoTransform [6]float64, err error) {
	err = cplErr(gdalGetGeoTransform(ds, &geoTransform))
	return
}

func gdalSetGeoTransform(dataset GDALDataset, geoTransform [6]float64) (result CPLErr) {
	var gt [6]C.double
	for i, v := range geoTransform {
		gt[i] = C.double(v)
	}
	result = CPLErr(C.GDALSetGeoTransform(dataset.cValue, &gt[0]))
	return
}

func (ds GDALDataset) SetGeoTransform(geoTransform [6]float64) (err error) {
	err = cplErr(gdalSetGeoTransform(ds, geoTransform))
	return
}

func gdalGetExtent(dataset GDALDataset, envelope OGREnvelope, crs OGRSpatialReference) (result CPLErr) {
	result = CPLErr(C.GDALGetExtent(dataset.cValue, envelope.cValue, crs.cValue))
	return
}

func (ds GDALDataset) GetExtent(crs OGRSpatialReference) (envelope OGREnvelope, err error) {
	envelope = OGREnvelope{cValue: new(C.OGREnvelope)}
	err = cplErr(gdalGetExtent(ds, envelope, crs))
	return
}

func gdalGetExtentWGS84LongLat(dataset GDALDataset, envelope OGREnvelope) (result CPLErr) {
	result = CPLErr(C.GDALGetExtentWGS84LongLat(dataset.cValue, envelope.cValue))
	return
}

func (ds GDALDataset) GetExtentWGS84LongLat() (envelope OGREnvelope, err error) {
	envelope = OGREnvelope{cValue: new(C.OGREnvelope)}
	err = cplErr(gdalGetExtentWGS84LongLat(ds, envelope))
	return
}

func gdalDatasetGeolocationToPixelLine(dataset GDALDataset, geolocX, geolocY float64, srs OGRSpatialReference, pixel, line *float64, transformerOptions []string) (result CPLErr) {
	opts, free := cStrings(transformerOptions)
	defer free()
	var cPixel, cLine C.double
	result = CPLErr(C.GDALDatasetGeolocationToPixelLine(dataset.cValue, C.double(geolocX), C.double(geolocY), srs.cValue, &cPixel, &cLine, C.CSLConstList(unsafe.Pointer(opts))))
	*pixel = float64(cPixel)
	*line = float64(cLine)
	return
}

func (ds GDALDataset) GeolocationToPixelLine(geolocX, geolocY float64, srs OGRSpatialReference, transformerOptions []string) (pixel, line float64, err error) {
	err = cplErr(gdalDatasetGeolocationToPixelLine(ds, geolocX, geolocY, srs, &pixel, &line, transformerOptions))
	return
}

func gdalGetGCPCount(dataset GDALDataset) (result int) {
	result = int(C.GDALGetGCPCount(dataset.cValue))
	return
}

func (ds GDALDataset) GetGCPCount() (result int) {
	result = gdalGetGCPCount(ds)
	return
}

func gdalGetGCPProjection(dataset GDALDataset) (result string) {
	result = C.GoString(C.GDALGetGCPProjection(dataset.cValue))
	return
}

func (ds GDALDataset) GetGCPProjection() (result string) {
	result = gdalGetGCPProjection(ds)
	return
}

func gdalGetGCPSpatialRef(dataset GDALDataset) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.GDALGetGCPSpatialRef(dataset.cValue)}
	return
}

func (ds GDALDataset) GetGCPSpatialRef() (result OGRSpatialReference) {
	result = gdalGetGCPSpatialRef(ds)
	return
}

func gdalGetGCPs(dataset GDALDataset) (result GDALGCPs) {
	raw := C.GDALGetGCPs(dataset.cValue)
	count := int(C.GDALGetGCPCount(dataset.cValue))
	if raw == nil || count == 0 {
		return
	}
	result = unsafe.Slice((*GDALGCP)(unsafe.Pointer(raw)), count)
	return
}

func (ds GDALDataset) GetGCPs() (result GDALGCPs) {
	result = gdalGetGCPs(ds)
	return
}

func gdalSetGCPs(dataset GDALDataset, count int, gcps GDALGCPs, projection string) (result CPLErr) {
	cProj := C.CString(projection)
	defer C.free(unsafe.Pointer(cProj))
	result = CPLErr(C.GDALSetGCPs(dataset.cValue, C.int(count), gcps.cPtr(), cProj))
	return
}

func (ds GDALDataset) SetGCPs(gcps GDALGCPs, projection string) (err error) {
	err = cplErr(gdalSetGCPs(ds, len(gcps), gcps, projection))
	return
}

func gdalSetGCPs2(dataset GDALDataset, count int, gcps GDALGCPs, srs OGRSpatialReference) (result CPLErr) {
	result = CPLErr(C.GDALSetGCPs2(dataset.cValue, C.int(count), gcps.cPtr(), srs.cValue))
	return
}

func (ds GDALDataset) SetGCPs2(gcps GDALGCPs, srs OGRSpatialReference) (err error) {
	err = cplErr(gdalSetGCPs2(ds, len(gcps), gcps, srs))
	return
}

func gdalGetInternalHandle(dataset GDALDataset, request string) unsafe.Pointer {
	cRequest := C.CString(request)
	defer C.free(unsafe.Pointer(cRequest))
	return C.GDALGetInternalHandle(dataset.cValue, cRequest)
}

func (ds GDALDataset) GetInternalHandle(request string) unsafe.Pointer {
	return gdalGetInternalHandle(ds, request)
}

func gdalReferenceDataset(dataset GDALDataset) (result int) {
	result = int(C.GDALReferenceDataset(dataset.cValue))
	return
}

func (ds GDALDataset) Reference() (result int) {
	result = gdalReferenceDataset(ds)
	return
}

func gdalDereferenceDataset(dataset GDALDataset) (result int) {
	result = int(C.GDALDereferenceDataset(dataset.cValue))
	return
}

func (ds GDALDataset) Dereference() (result int) {
	result = gdalDereferenceDataset(ds)
	return
}

func gdalReleaseDataset(dataset GDALDataset) (result int) {
	result = int(C.GDALReleaseDataset(dataset.cValue))
	return
}

func (ds GDALDataset) Release() (result int) {
	result = gdalReleaseDataset(ds)
	return
}

func gdalBuildOverviews(dataset GDALDataset, resampling string, nOverviews int, overviewList []int, bandCount int, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	result = CPLErr(C.GDALBuildOverviews(dataset.cValue, cResampling, C.int(nOverviews), cInts(overviewList), C.int(bandCount), cInts(bandList), progress.cValue, progressData))
	return
}

func (ds GDALDataset) BuildOverviews(resampling string, overviewList, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalBuildOverviews(ds, resampling, len(overviewList), overviewList, len(bandList), bandList, progress, progressData))
	return
}

func gdalBuildOverviewsEx(dataset GDALDataset, resampling string, nOverviews int, overviewList []int, bandCount int, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer, options []string) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	opts, free := cStrings(options)
	defer free()
	result = CPLErr(C.GDALBuildOverviewsEx(dataset.cValue, cResampling, C.int(nOverviews), cInts(overviewList), C.int(bandCount), cInts(bandList), progress.cValue, progressData, C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (ds GDALDataset) BuildOverviewsEx(resampling string, overviewList, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer, options []string) (err error) {
	err = cplErr(gdalBuildOverviewsEx(ds, resampling, len(overviewList), overviewList, len(bandList), bandList, progress, progressData, options))
	return
}

func gdalGetOpenDatasets(datasets *GDALDatasets, count *int) {
	var arr *C.GDALDatasetH
	var n C.int
	C.GDALGetOpenDatasets(&arr, &n)
	if count != nil {
		*count = int(n)
	}
	if datasets != nil {
		datasets.cValue = arr
	}
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

func gdalGetAccess(dataset GDALDataset) (result int) {
	result = int(C.GDALGetAccess(dataset.cValue))
	return
}

func (ds GDALDataset) GetAccess() (result int) {
	result = gdalGetAccess(ds)
	return
}

func gdalFlushCache(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALFlushCache(dataset.cValue))
	return
}

func (ds GDALDataset) FlushCache() (err error) {
	err = cplErr(gdalFlushCache(ds))
	return
}

func gdalDropCache(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALDropCache(dataset.cValue))
	return
}

func (ds GDALDataset) DropCache() (err error) {
	err = cplErr(gdalDropCache(ds))
	return
}

func gdalCreateDatasetMaskBand(dataset GDALDataset, flags int) (result CPLErr) {
	result = CPLErr(C.GDALCreateDatasetMaskBand(dataset.cValue, C.int(flags)))
	return
}

func (ds GDALDataset) CreateMaskBand(flags int) (err error) {
	err = cplErr(gdalCreateDatasetMaskBand(ds, flags))
	return
}

func gdalDatasetCopyWholeRaster(srcDataset, dstDataset GDALDataset, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	opts, free := cStrings(options)
	defer free()
	result = CPLErr(C.GDALDatasetCopyWholeRaster(srcDataset.cValue, dstDataset.cValue, C.CSLConstList(unsafe.Pointer(opts)), progress.cValue, progressData))
	return
}

func (src GDALDataset) CopyWholeRaster(dst GDALDataset, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalDatasetCopyWholeRaster(src, dst, options, progress, progressData))
	return
}

func gdalRasterBandCopyWholeRaster(srcBand, dstBand GDALRasterBand, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	opts, free := cStrings(options)
	defer free()
	result = CPLErr(C.GDALRasterBandCopyWholeRaster(srcBand.cValue, dstBand.cValue, opts, progress.cValue, progressData))
	return
}

func (src GDALRasterBand) CopyWholeRaster(dst GDALRasterBand, options []string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalRasterBandCopyWholeRaster(src, dst, options, progress, progressData))
	return
}

func gdalRegenerateOverviews(srcBand GDALRasterBand, overviewCount int, overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	result = CPLErr(C.GDALRegenerateOverviews(srcBand.cValue, C.int(overviewCount), overviewBands.cPtr(), cResampling, progress.cValue, progressData))
	return
}

func (src GDALRasterBand) RegenerateOverviews(overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalRegenerateOverviews(src, len(overviewBands), overviewBands, resampling, progress, progressData))
	return
}

func gdalRegenerateOverviewsEx(srcBand GDALRasterBand, overviewCount int, overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer, options []string) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	opts, free := cStrings(options)
	defer free()
	result = CPLErr(C.GDALRegenerateOverviewsEx(srcBand.cValue, C.int(overviewCount), overviewBands.cPtr(), cResampling, progress.cValue, progressData, C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (src GDALRasterBand) RegenerateOverviewsEx(overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer, options []string) (err error) {
	err = cplErr(gdalRegenerateOverviewsEx(src, len(overviewBands), overviewBands, resampling, progress, progressData, options))
	return
}

func gdalDatasetGetLayerCount(dataset GDALDataset) (result int) {
	result = int(C.GDALDatasetGetLayerCount(dataset.cValue))
	return
}

func (ds GDALDataset) GetLayerCount() (result int) {
	result = gdalDatasetGetLayerCount(ds)
	return
}

func gdalDatasetGetLayer(dataset GDALDataset, index int) (result OGRLayer) {
	result = OGRLayer{cValue: C.GDALDatasetGetLayer(dataset.cValue, C.int(index))}
	return
}

func (ds GDALDataset) GetLayer(index int) (result OGRLayer, err error) {
	result = gdalDatasetGetLayer(ds, index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// OGR_L_GetDataset is defined here to avoid a circular dependency with ogr_api.h.
func ogrLGetDataset(layer OGRLayer) (result GDALDataset) {
	result = GDALDataset{cValue: C.OGR_L_GetDataset(layer.cValue)}
	return
}

func (l OGRLayer) GetDataset() (result GDALDataset, err error) {
	result = ogrLGetDataset(l)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetGetLayerByName(dataset GDALDataset, name string) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = OGRLayer{cValue: C.GDALDatasetGetLayerByName(dataset.cValue, cName)}
	return
}

func (ds GDALDataset) GetLayerByName(name string) (result OGRLayer, err error) {
	result = gdalDatasetGetLayerByName(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetIsLayerPrivate(dataset GDALDataset, index int) (result bool) {
	result = C.GDALDatasetIsLayerPrivate(dataset.cValue, C.int(index)) != 0
	return
}

func (ds GDALDataset) IsLayerPrivate(index int) (result bool) {
	result = gdalDatasetIsLayerPrivate(ds, index)
	return
}

func gdalDatasetDeleteLayer(dataset GDALDataset, index int) (result OGRErr) {
	result = OGRErr(C.GDALDatasetDeleteLayer(dataset.cValue, C.int(index)))
	return
}

func (ds GDALDataset) DeleteLayer(index int) (err error) {
	err = ogrError(gdalDatasetDeleteLayer(ds, index))
	return
}

func gdalDatasetCreateLayer(dataset GDALDataset, name string, srs OGRSpatialReference, geomType OGRwkbGeometryType, options []string) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = OGRLayer{cValue: C.GDALDatasetCreateLayer(dataset.cValue, cName, srs.cValue, C.OGRwkbGeometryType(geomType), C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (ds GDALDataset) CreateLayer(name string, srs OGRSpatialReference, geomType OGRwkbGeometryType, options []string) (result OGRLayer, err error) {
	result = gdalDatasetCreateLayer(ds, name, srs, geomType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetCreateLayerFromGeomFieldDefn(dataset GDALDataset, name string, geomFieldDefn OGRGeomFieldDefn, options []string) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = OGRLayer{cValue: C.GDALDatasetCreateLayerFromGeomFieldDefn(dataset.cValue, cName, geomFieldDefn.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (ds GDALDataset) CreateLayerFromGeomFieldDefn(name string, geomFieldDefn OGRGeomFieldDefn, options []string) (result OGRLayer, err error) {
	result = gdalDatasetCreateLayerFromGeomFieldDefn(ds, name, geomFieldDefn, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetCopyLayer(dataset GDALDataset, srcLayer OGRLayer, newName string, options []string) (result OGRLayer) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = OGRLayer{cValue: C.GDALDatasetCopyLayer(dataset.cValue, srcLayer.cValue, cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (ds GDALDataset) CopyLayer(srcLayer OGRLayer, newName string, options []string) (result OGRLayer, err error) {
	result = gdalDatasetCopyLayer(ds, srcLayer, newName, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetResetReading(dataset GDALDataset) {
	C.GDALDatasetResetReading(dataset.cValue)
}

func (ds GDALDataset) ResetReading() {
	gdalDatasetResetReading(ds)
}

func gdalDatasetGetNextFeature(dataset GDALDataset, belongingLayer *OGRLayer, progressPct *float64, progress GDALProgressFunc, progressData unsafe.Pointer) (result OGRFeature) {
	var cLayer C.OGRLayerH
	var cPct C.double
	result = OGRFeature{cValue: C.GDALDatasetGetNextFeature(dataset.cValue, &cLayer, &cPct, progress.cValue, progressData)}
	if belongingLayer != nil {
		*belongingLayer = OGRLayer{cValue: cLayer}
	}
	if progressPct != nil {
		*progressPct = float64(cPct)
	}
	return
}

func (ds GDALDataset) GetNextFeature(progress GDALProgressFunc, progressData unsafe.Pointer) (feature OGRFeature, belongingLayer OGRLayer, progressPct float64, err error) {
	feature = gdalDatasetGetNextFeature(ds, &belongingLayer, &progressPct, progress, progressData)
	if feature.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetTestCapability(dataset GDALDataset, capability string) (result bool) {
	cCapability := C.CString(capability)
	defer C.free(unsafe.Pointer(cCapability))
	result = C.GDALDatasetTestCapability(dataset.cValue, cCapability) != 0
	return
}

func (ds GDALDataset) TestCapability(capability string) (result bool) {
	result = gdalDatasetTestCapability(ds, capability)
	return
}

func gdalDatasetExecuteSQL(dataset GDALDataset, statement string, spatialFilter OGRGeometry, dialect string) (result OGRLayer) {
	cStatement := C.CString(statement)
	defer C.free(unsafe.Pointer(cStatement))
	cDialect := C.CString(dialect)
	defer C.free(unsafe.Pointer(cDialect))
	result = OGRLayer{cValue: C.GDALDatasetExecuteSQL(dataset.cValue, cStatement, spatialFilter.cValue, cDialect)}
	return
}

func (ds GDALDataset) ExecuteSQL(statement string, spatialFilter OGRGeometry, dialect string) (result OGRLayer, err error) {
	result = gdalDatasetExecuteSQL(ds, statement, spatialFilter, dialect)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetAbortSQL(dataset GDALDataset) (result OGRErr) {
	result = OGRErr(C.GDALDatasetAbortSQL(dataset.cValue))
	return
}

func (ds GDALDataset) AbortSQL() (err error) {
	err = ogrError(gdalDatasetAbortSQL(ds))
	return
}

func gdalDatasetReleaseResultSet(dataset GDALDataset, layer OGRLayer) {
	C.GDALDatasetReleaseResultSet(dataset.cValue, layer.cValue)
}

func (ds GDALDataset) ReleaseResultSet(layer OGRLayer) {
	gdalDatasetReleaseResultSet(ds, layer)
}

func gdalDatasetGetStyleTable(dataset GDALDataset) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.GDALDatasetGetStyleTable(dataset.cValue)}
	return
}

func (ds GDALDataset) GetStyleTable() (result OGRStyleTable, err error) {
	result = gdalDatasetGetStyleTable(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetSetStyleTableDirectly(dataset GDALDataset, styleTable OGRStyleTable) {
	C.GDALDatasetSetStyleTableDirectly(dataset.cValue, styleTable.cValue)
}

func (ds GDALDataset) SetStyleTableDirectly(styleTable OGRStyleTable) {
	gdalDatasetSetStyleTableDirectly(ds, styleTable)
}

func gdalDatasetSetStyleTable(dataset GDALDataset, styleTable OGRStyleTable) {
	C.GDALDatasetSetStyleTable(dataset.cValue, styleTable.cValue)
}

func (ds GDALDataset) SetStyleTable(styleTable OGRStyleTable) {
	gdalDatasetSetStyleTable(ds, styleTable)
}

func gdalDatasetStartTransaction(dataset GDALDataset, force int) (result OGRErr) {
	result = OGRErr(C.GDALDatasetStartTransaction(dataset.cValue, C.int(force)))
	return
}

func (ds GDALDataset) StartTransaction(force int) (err error) {
	err = ogrError(gdalDatasetStartTransaction(ds, force))
	return
}

func gdalDatasetCommitTransaction(dataset GDALDataset) (result OGRErr) {
	result = OGRErr(C.GDALDatasetCommitTransaction(dataset.cValue))
	return
}

func (ds GDALDataset) CommitTransaction() (err error) {
	err = ogrError(gdalDatasetCommitTransaction(ds))
	return
}

func gdalDatasetRollbackTransaction(dataset GDALDataset) (result OGRErr) {
	result = OGRErr(C.GDALDatasetRollbackTransaction(dataset.cValue))
	return
}

func (ds GDALDataset) RollbackTransaction() (err error) {
	err = ogrError(gdalDatasetRollbackTransaction(ds))
	return
}

func gdalDatasetClearStatistics(dataset GDALDataset) {
	C.GDALDatasetClearStatistics(dataset.cValue)
}

func (ds GDALDataset) ClearStatistics() {
	gdalDatasetClearStatistics(ds)
}

func gdalDatasetAsMDArray(dataset GDALDataset, options []string) (result GDALMDArray) {
	opts, free := cStrings(options)
	defer free()
	result = GDALMDArray{cValue: C.GDALDatasetAsMDArray(dataset.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (ds GDALDataset) AsMDArray(options []string) (result GDALMDArray, err error) {
	result = gdalDatasetAsMDArray(ds, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetGetFieldDomainNames(dataset GDALDataset, options []string) (result []string) {
	opts, free := cStrings(options)
	defer free()
	raw := C.GDALDatasetGetFieldDomainNames(dataset.cValue, C.CSLConstList(unsafe.Pointer(opts)))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (ds GDALDataset) GetFieldDomainNames(options []string) (result []string) {
	result = gdalDatasetGetFieldDomainNames(ds, options)
	return
}

func gdalDatasetGetFieldDomain(dataset GDALDataset, name string) (result OGRFieldDomain) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = OGRFieldDomain{cValue: C.GDALDatasetGetFieldDomain(dataset.cValue, cName)}
	return
}

func (ds GDALDataset) GetFieldDomain(name string) (result OGRFieldDomain, err error) {
	result = gdalDatasetGetFieldDomain(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetAddFieldDomain(dataset GDALDataset, fieldDomain OGRFieldDomain, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetAddFieldDomain(dataset.cValue, fieldDomain.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		C.VSIFree(unsafe.Pointer(cReason))
	}
	return
}

func (ds GDALDataset) AddFieldDomain(fieldDomain OGRFieldDomain) (ok bool, failureReason string) {
	ok = gdalDatasetAddFieldDomain(ds, fieldDomain, &failureReason)
	return
}

func gdalDatasetDeleteFieldDomain(dataset GDALDataset, name string, failureReason *string) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cReason *C.char
	result = bool(C.GDALDatasetDeleteFieldDomain(dataset.cValue, cName, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		C.VSIFree(unsafe.Pointer(cReason))
	}
	return
}

func (ds GDALDataset) DeleteFieldDomain(name string) (ok bool, failureReason string) {
	ok = gdalDatasetDeleteFieldDomain(ds, name, &failureReason)
	return
}

func gdalDatasetUpdateFieldDomain(dataset GDALDataset, fieldDomain OGRFieldDomain, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetUpdateFieldDomain(dataset.cValue, fieldDomain.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		C.VSIFree(unsafe.Pointer(cReason))
	}
	return
}

func (ds GDALDataset) UpdateFieldDomain(fieldDomain OGRFieldDomain) (ok bool, failureReason string) {
	ok = gdalDatasetUpdateFieldDomain(ds, fieldDomain, &failureReason)
	return
}

func gdalDatasetGetRelationshipNames(dataset GDALDataset, options []string) (result []string) {
	opts, free := cStrings(options)
	defer free()
	raw := C.GDALDatasetGetRelationshipNames(dataset.cValue, C.CSLConstList(unsafe.Pointer(opts)))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (ds GDALDataset) GetRelationshipNames(options []string) (result []string) {
	result = gdalDatasetGetRelationshipNames(ds, options)
	return
}

func gdalDatasetGetRelationship(dataset GDALDataset, name string) (result GDALRelationship) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALRelationship{cValue: C.GDALDatasetGetRelationship(dataset.cValue, cName)}
	return
}

func (ds GDALDataset) GetRelationship(name string) (result GDALRelationship, err error) {
	result = gdalDatasetGetRelationship(ds, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetAddRelationship(dataset GDALDataset, relationship GDALRelationship, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetAddRelationship(dataset.cValue, relationship.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		C.VSIFree(unsafe.Pointer(cReason))
	}
	return
}

func (ds GDALDataset) AddRelationship(relationship GDALRelationship) (ok bool, failureReason string) {
	ok = gdalDatasetAddRelationship(ds, relationship, &failureReason)
	return
}

func gdalDatasetDeleteRelationship(dataset GDALDataset, name string, failureReason *string) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cReason *C.char
	result = bool(C.GDALDatasetDeleteRelationship(dataset.cValue, cName, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		C.VSIFree(unsafe.Pointer(cReason))
	}
	return
}

func (ds GDALDataset) DeleteRelationship(name string) (ok bool, failureReason string) {
	ok = gdalDatasetDeleteRelationship(ds, name, &failureReason)
	return
}

func gdalDatasetUpdateRelationship(dataset GDALDataset, relationship GDALRelationship, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetUpdateRelationship(dataset.cValue, relationship.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		C.VSIFree(unsafe.Pointer(cReason))
	}
	return
}

func (ds GDALDataset) UpdateRelationship(relationship GDALRelationship) (ok bool, failureReason string) {
	ok = gdalDatasetUpdateRelationship(ds, relationship, &failureReason)
	return
}

// /** Type of functions to pass to GDALDatasetSetQueryLoggerFunc
//  * @since GDAL 3.7 */
// GDALDatasetSetQueryLoggerFunc (with the GDALQueryLoggerFunc callback) is
// deferred: passing a Go function as a C callback needs //export plumbing and a
// registration shim, to be designed later.

// /* ==================================================================== */
// /*      Informational utilities about subdatasets in file names         */
// /* ==================================================================== */

// /**
//   - @brief Returns a new GDALSubdatasetInfo object with methods to extract
//   - and manipulate subdataset information.
//   - If the pszFileName argument is not recognized by any driver as
//   - a subdataset descriptor, NULL is returned.
//   - The returned object must be freed with GDALDestroySubdatasetInfo().
//   - @param pszFileName           File name with subdataset information
//   - @note                        This method does not check if the subdataset actually exists.
//   - @return                      Opaque pointer to a GDALSubdatasetInfo object or NULL if no drivers accepted the file name.
//   - @since                       GDAL 3.8
//     */
func gdalGetSubdatasetInfo(fileName string) (result GDALSubdatasetInfo) {
	cName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cName))
	result = GDALSubdatasetInfo{cValue: C.GDALGetSubdatasetInfo(cName)}
	return
}

func GDALGetSubdatasetInfo(fileName string) (result GDALSubdatasetInfo, err error) {
	result = gdalGetSubdatasetInfo(fileName)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// /**
//   - @brief Returns the file path component of a
//   - subdataset descriptor effectively stripping the information about the subdataset
//   - and returning the "parent" dataset descriptor.
//   - The returned string must be freed with CPLFree().
//   - @param hInfo                 Pointer to GDALSubdatasetInfo object
//   - @note                        This method does not check if the subdataset actually exists.
//   - @return                      The original string with the subdataset information removed.
//   - @since                       GDAL 3.8
//     */
func gdalSubdatasetInfoGetPathComponent(info GDALSubdatasetInfo) (result string) {
	cResult := C.GDALSubdatasetInfoGetPathComponent(info.cValue)
	if cResult != nil {
		result = C.GoString(cResult)
		C.VSIFree(unsafe.Pointer(cResult))
	}
	return
}

func (i GDALSubdatasetInfo) GetPathComponent() (result string) {
	result = gdalSubdatasetInfoGetPathComponent(i)
	return
}

// /**
//   - @brief Returns the subdataset component of a subdataset descriptor descriptor.
//   - The returned string must be freed with CPLFree().
//   - @param hInfo                 Pointer to GDALSubdatasetInfo object
//   - @note                        This method does not check if the subdataset actually exists.
//   - @return                      The subdataset name.
//   - @since                       GDAL 3.8
//     */
func gdalSubdatasetInfoGetSubdatasetComponent(info GDALSubdatasetInfo) (result string) {
	cResult := C.GDALSubdatasetInfoGetSubdatasetComponent(info.cValue)
	if cResult != nil {
		result = C.GoString(cResult)
		C.VSIFree(unsafe.Pointer(cResult))
	}
	return
}

func (i GDALSubdatasetInfo) GetSubdatasetComponent() (result string) {
	result = gdalSubdatasetInfoGetSubdatasetComponent(i)
	return
}

// /**
//   - @brief Replaces the path component of a subdataset descriptor.
//   - The returned string must be freed with CPLFree().
//   - @param hInfo                 Pointer to GDALSubdatasetInfo object
//   - @param pszNewPath            New path.
//   - @note                        This method does not check if the subdataset actually exists.
//   - @return                      The original subdataset descriptor with the old path component replaced by newPath.
//   - @since                       GDAL 3.8
//     */
func gdalSubdatasetInfoModifyPathComponent(info GDALSubdatasetInfo, newPath string) (result string) {
	cNewPath := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNewPath))
	cResult := C.GDALSubdatasetInfoModifyPathComponent(info.cValue, cNewPath)
	if cResult != nil {
		result = C.GoString(cResult)
		C.VSIFree(unsafe.Pointer(cResult))
	}
	return
}

func (i GDALSubdatasetInfo) ModifyPathComponent(newPath string) (result string) {
	result = gdalSubdatasetInfoModifyPathComponent(i, newPath)
	return
}

// /**
//   - @brief Destroys a GDALSubdatasetInfo object.
//   - @param hInfo                 Pointer to GDALSubdatasetInfo object
//   - @since                       GDAL 3.8
//     */
func gdalDestroySubdatasetInfo(info GDALSubdatasetInfo) {
	C.GDALDestroySubdatasetInfo(info.cValue)
}

func (i GDALSubdatasetInfo) Destroy() {
	gdalDestroySubdatasetInfo(i)
}

// /* ==================================================================== */
// /*      GDALRasterBand ... one band/channel in a dataset.               */
// /* ==================================================================== */

// The SRCVAL pixel macro and the GDALDerivedPixelFunc / GDALDerivedPixelFuncWithArgs
// callback typedefs are omitted/deferred (a code macro and Go->C callback
// plumbing respectively).

func gdalGetRasterDataType(band GDALRasterBand) (result GDALDataType) {
	result = GDALDataType(C.GDALGetRasterDataType(band.cValue))
	return
}

func (b GDALRasterBand) GetRasterDataType() (result GDALDataType) {
	result = gdalGetRasterDataType(b)
	return
}

func gdalGetBlockSize(band GDALRasterBand, xSize, ySize *int) {
	var cXSize, cYSize C.int
	C.GDALGetBlockSize(band.cValue, &cXSize, &cYSize)
	*xSize = int(cXSize)
	*ySize = int(cYSize)
}

func (b GDALRasterBand) GetBlockSize() (xSize, ySize int) {
	gdalGetBlockSize(b, &xSize, &ySize)
	return
}

func gdalGetActualBlockSize(band GDALRasterBand, xBlockOff, yBlockOff int, xValid, yValid *int) (result CPLErr) {
	var cXValid, cYValid C.int
	result = CPLErr(C.GDALGetActualBlockSize(band.cValue, C.int(xBlockOff), C.int(yBlockOff), &cXValid, &cYValid))
	*xValid = int(cXValid)
	*yValid = int(cYValid)
	return
}

func (b GDALRasterBand) GetActualBlockSize(xBlockOff, yBlockOff int) (xValid, yValid int, err error) {
	err = cplErr(gdalGetActualBlockSize(b, xBlockOff, yBlockOff, &xValid, &yValid))
	return
}

func gdalRasterAdviseRead(band GDALRasterBand, xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, options []string) (result CPLErr) {
	opts, free := cStrings(options)
	defer free()
	result = CPLErr(C.GDALRasterAdviseRead(band.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (b GDALRasterBand) AdviseRead(xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, options []string) (err error) {
	err = cplErr(gdalRasterAdviseRead(b, xOff, yOff, xSize, ySize, bufXSize, bufYSize, bufType, options))
	return
}

func gdalRasterIO(band GDALRasterBand, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int) (result CPLErr) {
	result = CPLErr(C.GDALRasterIO(band.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(pixelSpace), C.int(lineSpace)))
	return
}

func (b GDALRasterBand) RasterIO(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int) (err error) {
	err = cplErr(gdalRasterIO(b, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, pixelSpace, lineSpace))
	return
}

func gdalRasterIOEx(band GDALRasterBand, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int64, extraArg GDALRasterIOExtraArg) (result CPLErr) {
	result = CPLErr(C.GDALRasterIOEx(band.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.GSpacing(pixelSpace), C.GSpacing(lineSpace), extraArg.cValue))
	return
}

func (b GDALRasterBand) RasterIOEx(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int64, extraArg GDALRasterIOExtraArg) (err error) {
	err = cplErr(gdalRasterIOEx(b, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, pixelSpace, lineSpace, extraArg))
	return
}

func gdalReadBlock(band GDALRasterBand, xBlockOff, yBlockOff int, buffer unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.GDALReadBlock(band.cValue, C.int(xBlockOff), C.int(yBlockOff), buffer))
	return
}

func (b GDALRasterBand) ReadBlock(xBlockOff, yBlockOff int, buffer []byte) (err error) {
	err = cplErr(gdalReadBlock(b, xBlockOff, yBlockOff, cBytes(buffer)))
	return
}

func gdalWriteBlock(band GDALRasterBand, xBlockOff, yBlockOff int, buffer unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.GDALWriteBlock(band.cValue, C.int(xBlockOff), C.int(yBlockOff), buffer))
	return
}

func (b GDALRasterBand) WriteBlock(xBlockOff, yBlockOff int, buffer []byte) (err error) {
	err = cplErr(gdalWriteBlock(b, xBlockOff, yBlockOff, cBytes(buffer)))
	return
}

func gdalGetRasterBandXSize(band GDALRasterBand) (result int) {
	result = int(C.GDALGetRasterBandXSize(band.cValue))
	return
}

func (b GDALRasterBand) GetXSize() (result int) {
	result = gdalGetRasterBandXSize(b)
	return
}

func gdalGetRasterBandYSize(band GDALRasterBand) (result int) {
	result = int(C.GDALGetRasterBandYSize(band.cValue))
	return
}

func (b GDALRasterBand) GetYSize() (result int) {
	result = gdalGetRasterBandYSize(b)
	return
}

func gdalGetRasterAccess(band GDALRasterBand) (result GDALAccess) {
	result = GDALAccess(C.GDALGetRasterAccess(band.cValue))
	return
}

func (b GDALRasterBand) GetAccess() (result GDALAccess) {
	result = gdalGetRasterAccess(b)
	return
}

func gdalGetBandNumber(band GDALRasterBand) (result int) {
	result = int(C.GDALGetBandNumber(band.cValue))
	return
}

func (b GDALRasterBand) GetBandNumber() (result int) {
	result = gdalGetBandNumber(b)
	return
}

func gdalGetBandDataset(band GDALRasterBand) (result GDALDataset) {
	result = GDALDataset{cValue: C.GDALGetBandDataset(band.cValue)}
	return
}

func (b GDALRasterBand) GetDataset() (result GDALDataset, err error) {
	result = gdalGetBandDataset(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetRasterColorInterpretation(band GDALRasterBand) (result GDALColorInterp) {
	result = GDALColorInterp(C.GDALGetRasterColorInterpretation(band.cValue))
	return
}

func (b GDALRasterBand) GetColorInterpretation() (result GDALColorInterp) {
	result = gdalGetRasterColorInterpretation(b)
	return
}

func gdalSetRasterColorInterpretation(band GDALRasterBand, colorInterp GDALColorInterp) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterColorInterpretation(band.cValue, C.GDALColorInterp(colorInterp)))
	return
}

func (b GDALRasterBand) SetColorInterpretation(colorInterp GDALColorInterp) (err error) {
	err = cplErr(gdalSetRasterColorInterpretation(b, colorInterp))
	return
}

func gdalGetRasterColorTable(band GDALRasterBand) (result GDALColorTable) {
	result = GDALColorTable{cValue: C.GDALGetRasterColorTable(band.cValue)}
	return
}

func (b GDALRasterBand) GetColorTable() (result GDALColorTable, err error) {
	result = gdalGetRasterColorTable(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalSetRasterColorTable(band GDALRasterBand, colorTable GDALColorTable) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterColorTable(band.cValue, colorTable.cValue))
	return
}

func (b GDALRasterBand) SetColorTable(colorTable GDALColorTable) (err error) {
	err = cplErr(gdalSetRasterColorTable(b, colorTable))
	return
}

func gdalHasArbitraryOverviews(band GDALRasterBand) (result bool) {
	result = C.GDALHasArbitraryOverviews(band.cValue) != 0
	return
}

func (b GDALRasterBand) HasArbitraryOverviews() (result bool) {
	result = gdalHasArbitraryOverviews(b)
	return
}

func gdalGetOverviewCount(band GDALRasterBand) (result int) {
	result = int(C.GDALGetOverviewCount(band.cValue))
	return
}

func (b GDALRasterBand) GetOverviewCount() (result int) {
	result = gdalGetOverviewCount(b)
	return
}

func gdalGetOverview(band GDALRasterBand, index int) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetOverview(band.cValue, C.int(index))}
	return
}

func (b GDALRasterBand) GetOverview(index int) (result GDALRasterBand, err error) {
	result = gdalGetOverview(b, index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetRasterNoDataValue(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterNoDataValue(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func (b GDALRasterBand) GetNoDataValue() (value float64, ok bool) {
	var success int
	value = gdalGetRasterNoDataValue(b, &success)
	ok = success != 0
	return
}

func gdalGetRasterNoDataValueAsInt64(band GDALRasterBand, success *int) (result int64) {
	var cSuccess C.int
	result = int64(C.GDALGetRasterNoDataValueAsInt64(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func (b GDALRasterBand) GetNoDataValueAsInt64() (value int64, ok bool) {
	var success int
	value = gdalGetRasterNoDataValueAsInt64(b, &success)
	ok = success != 0
	return
}

func gdalGetRasterNoDataValueAsUInt64(band GDALRasterBand, success *int) (result uint64) {
	var cSuccess C.int
	result = uint64(C.GDALGetRasterNoDataValueAsUInt64(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func (b GDALRasterBand) GetNoDataValueAsUInt64() (value uint64, ok bool) {
	var success int
	value = gdalGetRasterNoDataValueAsUInt64(b, &success)
	ok = success != 0
	return
}

func gdalSetRasterNoDataValue(band GDALRasterBand, value float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterNoDataValue(band.cValue, C.double(value)))
	return
}

func (b GDALRasterBand) SetNoDataValue(value float64) (err error) {
	err = cplErr(gdalSetRasterNoDataValue(b, value))
	return
}

func gdalSetRasterNoDataValueAsInt64(band GDALRasterBand, value int64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterNoDataValueAsInt64(band.cValue, C.int64_t(value)))
	return
}

func (b GDALRasterBand) SetNoDataValueAsInt64(value int64) (err error) {
	err = cplErr(gdalSetRasterNoDataValueAsInt64(b, value))
	return
}

func gdalSetRasterNoDataValueAsUInt64(band GDALRasterBand, value uint64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterNoDataValueAsUInt64(band.cValue, C.uint64_t(value)))
	return
}

func (b GDALRasterBand) SetNoDataValueAsUInt64(value uint64) (err error) {
	err = cplErr(gdalSetRasterNoDataValueAsUInt64(b, value))
	return
}

func gdalDeleteRasterNoDataValue(band GDALRasterBand) (result CPLErr) {
	result = CPLErr(C.GDALDeleteRasterNoDataValue(band.cValue))
	return
}

func (b GDALRasterBand) DeleteNoDataValue() (err error) {
	err = cplErr(gdalDeleteRasterNoDataValue(b))
	return
}

func gdalGetRasterCategoryNames(band GDALRasterBand) (result []string) {
	result = goStrings(C.GDALGetRasterCategoryNames(band.cValue))
	return
}

func (b GDALRasterBand) GetCategoryNames() (result []string) {
	result = gdalGetRasterCategoryNames(b)
	return
}

func gdalSetRasterCategoryNames(band GDALRasterBand, names []string) (result CPLErr) {
	n, free := cStrings(names)
	defer free()
	result = CPLErr(C.GDALSetRasterCategoryNames(band.cValue, C.CSLConstList(unsafe.Pointer(n))))
	return
}

func (b GDALRasterBand) SetCategoryNames(names []string) (err error) {
	err = cplErr(gdalSetRasterCategoryNames(b, names))
	return
}

func gdalGetRasterMinimum(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterMinimum(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func (b GDALRasterBand) GetMinimum() (value float64, ok bool) {
	var success int
	value = gdalGetRasterMinimum(b, &success)
	ok = success != 0
	return
}

func gdalGetRasterMaximum(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterMaximum(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func (b GDALRasterBand) GetMaximum() (value float64, ok bool) {
	var success int
	value = gdalGetRasterMaximum(b, &success)
	ok = success != 0
	return
}

func gdalGetRasterStatistics(band GDALRasterBand, approxOK, force int, min, max, mean, stdDev *float64) (result CPLErr) {
	var cMin, cMax, cMean, cStdDev C.double
	result = CPLErr(C.GDALGetRasterStatistics(band.cValue, C.int(approxOK), C.int(force), &cMin, &cMax, &cMean, &cStdDev))
	*min = float64(cMin)
	*max = float64(cMax)
	*mean = float64(cMean)
	*stdDev = float64(cStdDev)
	return
}

func (b GDALRasterBand) GetStatistics(approxOK, force int) (min, max, mean, stdDev float64, err error) {
	err = cplErr(gdalGetRasterStatistics(b, approxOK, force, &min, &max, &mean, &stdDev))
	return
}

func gdalComputeRasterStatistics(band GDALRasterBand, approxOK int, min, max, mean, stdDev *float64, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	var cMin, cMax, cMean, cStdDev C.double
	result = CPLErr(C.GDALComputeRasterStatistics(band.cValue, C.int(approxOK), &cMin, &cMax, &cMean, &cStdDev, progress.cValue, progressData))
	*min = float64(cMin)
	*max = float64(cMax)
	*mean = float64(cMean)
	*stdDev = float64(cStdDev)
	return
}

func (b GDALRasterBand) ComputeStatistics(approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max, mean, stdDev float64, err error) {
	err = cplErr(gdalComputeRasterStatistics(b, approxOK, &min, &max, &mean, &stdDev, progress, progressData))
	return
}

func gdalSetRasterStatistics(band GDALRasterBand, min, max, mean, stdDev float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterStatistics(band.cValue, C.double(min), C.double(max), C.double(mean), C.double(stdDev)))
	return
}

func (b GDALRasterBand) SetStatistics(min, max, mean, stdDev float64) (err error) {
	err = cplErr(gdalSetRasterStatistics(b, min, max, mean, stdDev))
	return
}

func gdalRasterBandAsMDArray(band GDALRasterBand) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALRasterBandAsMDArray(band.cValue)}
	return
}

func (b GDALRasterBand) AsMDArray() (result GDALMDArray, err error) {
	result = gdalRasterBandAsMDArray(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetRasterUnitType(band GDALRasterBand) (result string) {
	result = C.GoString(C.GDALGetRasterUnitType(band.cValue))
	return
}

func (b GDALRasterBand) GetUnitType() (result string) {
	result = gdalGetRasterUnitType(b)
	return
}

func gdalSetRasterUnitType(band GDALRasterBand, newValue string) (result CPLErr) {
	cValue := C.CString(newValue)
	defer C.free(unsafe.Pointer(cValue))
	result = CPLErr(C.GDALSetRasterUnitType(band.cValue, cValue))
	return
}

func (b GDALRasterBand) SetUnitType(newValue string) (err error) {
	err = cplErr(gdalSetRasterUnitType(b, newValue))
	return
}

func gdalGetRasterOffset(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterOffset(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func (b GDALRasterBand) GetOffset() (value float64, ok bool) {
	var success int
	value = gdalGetRasterOffset(b, &success)
	ok = success != 0
	return
}

func gdalSetRasterOffset(band GDALRasterBand, newOffset float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterOffset(band.cValue, C.double(newOffset)))
	return
}

func (b GDALRasterBand) SetOffset(newOffset float64) (err error) {
	err = cplErr(gdalSetRasterOffset(b, newOffset))
	return
}

func gdalGetRasterScale(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterScale(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func (b GDALRasterBand) GetScale() (value float64, ok bool) {
	var success int
	value = gdalGetRasterScale(b, &success)
	ok = success != 0
	return
}

func gdalSetRasterScale(band GDALRasterBand, newScale float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterScale(band.cValue, C.double(newScale)))
	return
}

func (b GDALRasterBand) SetScale(newScale float64) (err error) {
	err = cplErr(gdalSetRasterScale(b, newScale))
	return
}

func gdalComputeRasterMinMax(band GDALRasterBand, approxOK int, minMax *[2]float64) (result CPLErr) {
	var cMinMax [2]C.double
	result = CPLErr(C.GDALComputeRasterMinMax(band.cValue, C.int(approxOK), &cMinMax[0]))
	minMax[0] = float64(cMinMax[0])
	minMax[1] = float64(cMinMax[1])
	return
}

func (b GDALRasterBand) ComputeMinMax(approxOK int) (minMax [2]float64, err error) {
	err = cplErr(gdalComputeRasterMinMax(b, approxOK, &minMax))
	return
}

func gdalComputeRasterMinMaxLocation(band GDALRasterBand, min, max *float64, minX, minY, maxX, maxY *int) (result CPLErr) {
	var cMin, cMax C.double
	var cMinX, cMinY, cMaxX, cMaxY C.int
	result = CPLErr(C.GDALComputeRasterMinMaxLocation(band.cValue, &cMin, &cMax, &cMinX, &cMinY, &cMaxX, &cMaxY))
	*min = float64(cMin)
	*max = float64(cMax)
	*minX = int(cMinX)
	*minY = int(cMinY)
	*maxX = int(cMaxX)
	*maxY = int(cMaxY)
	return
}

func (b GDALRasterBand) ComputeMinMaxLocation() (min, max float64, minX, minY, maxX, maxY int, err error) {
	err = cplErr(gdalComputeRasterMinMaxLocation(b, &min, &max, &minX, &minY, &maxX, &maxY))
	return
}

func gdalFlushRasterCache(band GDALRasterBand) (result CPLErr) {
	result = CPLErr(C.GDALFlushRasterCache(band.cValue))
	return
}

func (b GDALRasterBand) FlushCache() (err error) {
	err = cplErr(gdalFlushRasterCache(b))
	return
}

func gdalDropRasterCache(band GDALRasterBand) (result CPLErr) {
	result = CPLErr(C.GDALDropRasterCache(band.cValue))
	return
}

func (b GDALRasterBand) DropCache() (err error) {
	err = cplErr(gdalDropRasterCache(b))
	return
}

// Deprecated: use GetHistogramEx.
func gdalGetRasterHistogram(band GDALRasterBand, min, max float64, nBuckets int, histogram []int, includeOutOfRange, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cHist := make([]C.int, nBuckets)
	var ptr *C.int
	if nBuckets > 0 {
		ptr = &cHist[0]
	}
	result = CPLErr(C.GDALGetRasterHistogram(band.cValue, C.double(min), C.double(max), C.int(nBuckets), ptr, C.int(includeOutOfRange), C.int(approxOK), progress.cValue, progressData))
	for i := 0; i < nBuckets; i++ {
		histogram[i] = int(cHist[i])
	}
	return
}

// Deprecated: use GetHistogramEx.
func (b GDALRasterBand) GetHistogram(min, max float64, nBuckets, includeOutOfRange, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (histogram []int, err error) {
	histogram = make([]int, nBuckets)
	err = cplErr(gdalGetRasterHistogram(b, min, max, nBuckets, histogram, includeOutOfRange, approxOK, progress, progressData))
	return
}

func gdalGetRasterHistogramEx(band GDALRasterBand, min, max float64, nBuckets int, histogram []uint64, includeOutOfRange, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cHist := make([]C.GUIntBig, nBuckets)
	var ptr *C.GUIntBig
	if nBuckets > 0 {
		ptr = &cHist[0]
	}
	result = CPLErr(C.GDALGetRasterHistogramEx(band.cValue, C.double(min), C.double(max), C.int(nBuckets), ptr, C.int(includeOutOfRange), C.int(approxOK), progress.cValue, progressData))
	for i := 0; i < nBuckets; i++ {
		histogram[i] = uint64(cHist[i])
	}
	return
}

func (b GDALRasterBand) GetHistogramEx(min, max float64, nBuckets, includeOutOfRange, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (histogram []uint64, err error) {
	histogram = make([]uint64, nBuckets)
	err = cplErr(gdalGetRasterHistogramEx(b, min, max, nBuckets, histogram, includeOutOfRange, approxOK, progress, progressData))
	return
}

func gdalGetDefaultHistogramEx(band GDALRasterBand, min, max *float64, buckets *int, histogram *[]uint64, force int, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	var cMin, cMax C.double
	var cBuckets C.int
	var cHist *C.GUIntBig
	result = CPLErr(C.GDALGetDefaultHistogramEx(band.cValue, &cMin, &cMax, &cBuckets, &cHist, C.int(force), progress.cValue, progressData))
	*min = float64(cMin)
	*max = float64(cMax)
	*buckets = int(cBuckets)
	if cHist != nil {
		src := unsafe.Slice(cHist, int(cBuckets))
		h := make([]uint64, int(cBuckets))
		for i := range h {
			h[i] = uint64(src[i])
		}
		*histogram = h
		C.VSIFree(unsafe.Pointer(cHist))
	}
	return
}

func (b GDALRasterBand) GetDefaultHistogramEx(force int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max float64, buckets int, histogram []uint64, err error) {
	err = cplErr(gdalGetDefaultHistogramEx(b, &min, &max, &buckets, &histogram, force, progress, progressData))
	return
}

func gdalSetDefaultHistogramEx(band GDALRasterBand, min, max float64, nBuckets int, histogram []uint64) (result CPLErr) {
	cHist := make([]C.GUIntBig, nBuckets)
	for i := 0; i < nBuckets && i < len(histogram); i++ {
		cHist[i] = C.GUIntBig(histogram[i])
	}
	var ptr *C.GUIntBig
	if nBuckets > 0 {
		ptr = &cHist[0]
	}
	result = CPLErr(C.GDALSetDefaultHistogramEx(band.cValue, C.double(min), C.double(max), C.int(nBuckets), ptr))
	return
}

func (b GDALRasterBand) SetDefaultHistogramEx(min, max float64, histogram []uint64) (err error) {
	err = cplErr(gdalSetDefaultHistogramEx(b, min, max, len(histogram), histogram))
	return
}

func gdalGetRandomRasterSample(band GDALRasterBand, samples int, buffer []float32) (result int) {
	var ptr *C.float
	if len(buffer) > 0 {
		ptr = (*C.float)(unsafe.Pointer(&buffer[0]))
	}
	result = int(C.GDALGetRandomRasterSample(band.cValue, C.int(samples), ptr))
	return
}

func (b GDALRasterBand) GetRandomSample(samples int) (buffer []float32, count int) {
	buffer = make([]float32, samples)
	count = gdalGetRandomRasterSample(b, samples, buffer)
	return
}

func gdalGetRasterSampleOverview(band GDALRasterBand, desiredSamples int) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetRasterSampleOverview(band.cValue, C.int(desiredSamples))}
	return
}

func (b GDALRasterBand) GetSampleOverview(desiredSamples int) (result GDALRasterBand, err error) {
	result = gdalGetRasterSampleOverview(b, desiredSamples)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetRasterSampleOverviewEx(band GDALRasterBand, desiredSamples uint64) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetRasterSampleOverviewEx(band.cValue, C.GUIntBig(desiredSamples))}
	return
}

func (b GDALRasterBand) GetSampleOverviewEx(desiredSamples uint64) (result GDALRasterBand, err error) {
	result = gdalGetRasterSampleOverviewEx(b, desiredSamples)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalFillRaster(band GDALRasterBand, realValue, imaginaryValue float64) (result CPLErr) {
	result = CPLErr(C.GDALFillRaster(band.cValue, C.double(realValue), C.double(imaginaryValue)))
	return
}

func (b GDALRasterBand) Fill(realValue, imaginaryValue float64) (err error) {
	err = cplErr(gdalFillRaster(b, realValue, imaginaryValue))
	return
}

func gdalComputeBandStats(band GDALRasterBand, sampleStep int, mean, stdDev *float64, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	var cMean, cStdDev C.double
	result = CPLErr(C.GDALComputeBandStats(band.cValue, C.int(sampleStep), &cMean, &cStdDev, progress.cValue, progressData))
	*mean = float64(cMean)
	*stdDev = float64(cStdDev)
	return
}

func (b GDALRasterBand) ComputeBandStats(sampleStep int, progress GDALProgressFunc, progressData unsafe.Pointer) (mean, stdDev float64, err error) {
	err = cplErr(gdalComputeBandStats(b, sampleStep, &mean, &stdDev, progress, progressData))
	return
}

func gdalOverviewMagnitudeCorrection(baseBand GDALRasterBand, overviewCount int, overviews GDALRasterBands, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.GDALOverviewMagnitudeCorrection(baseBand.cValue, C.int(overviewCount), overviews.cPtr(), progress.cValue, progressData))
	return
}

func (b GDALRasterBand) OverviewMagnitudeCorrection(overviews GDALRasterBands, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	err = cplErr(gdalOverviewMagnitudeCorrection(b, len(overviews), overviews, progress, progressData))
	return
}

func gdalGetDefaultRAT(band GDALRasterBand) (result GDALRasterAttributeTable) {
	result = GDALRasterAttributeTable{cValue: C.GDALGetDefaultRAT(band.cValue)}
	return
}

func (b GDALRasterBand) GetDefaultRAT() (result GDALRasterAttributeTable, err error) {
	result = gdalGetDefaultRAT(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalSetDefaultRAT(band GDALRasterBand, rat GDALRasterAttributeTable) (result CPLErr) {
	result = CPLErr(C.GDALSetDefaultRAT(band.cValue, rat.cValue))
	return
}

func (b GDALRasterBand) SetDefaultRAT(rat GDALRasterAttributeTable) (err error) {
	err = cplErr(gdalSetDefaultRAT(b, rat))
	return
}

// GDALAddDerivedBandPixelFunc and GDALAddDerivedBandPixelFuncWithArgs are
// deferred: registering a Go pixel function as a C callback needs //export
// plumbing.

func gdalRasterInterpolateAtPoint(band GDALRasterBand, pixel, line float64, interpolation GDALRIOResampleAlg, realValue, imagValue *float64) (result CPLErr) {
	var cReal, cImag C.double
	result = CPLErr(C.GDALRasterInterpolateAtPoint(band.cValue, C.double(pixel), C.double(line), C.GDALRIOResampleAlg(interpolation), &cReal, &cImag))
	*realValue = float64(cReal)
	*imagValue = float64(cImag)
	return
}

func (b GDALRasterBand) InterpolateAtPoint(pixel, line float64, interpolation GDALRIOResampleAlg) (realValue, imagValue float64, err error) {
	err = cplErr(gdalRasterInterpolateAtPoint(b, pixel, line, interpolation, &realValue, &imagValue))
	return
}

func gdalRasterInterpolateAtGeolocation(band GDALRasterBand, geolocX, geolocY float64, srs OGRSpatialReference, interpolation GDALRIOResampleAlg, realValue, imagValue *float64, transformerOptions []string) (result CPLErr) {
	opts, free := cStrings(transformerOptions)
	defer free()
	var cReal, cImag C.double
	result = CPLErr(C.GDALRasterInterpolateAtGeolocation(band.cValue, C.double(geolocX), C.double(geolocY), srs.cValue, C.GDALRIOResampleAlg(interpolation), &cReal, &cImag, C.CSLConstList(unsafe.Pointer(opts))))
	*realValue = float64(cReal)
	*imagValue = float64(cImag)
	return
}

func (b GDALRasterBand) InterpolateAtGeolocation(geolocX, geolocY float64, srs OGRSpatialReference, interpolation GDALRIOResampleAlg, transformerOptions []string) (realValue, imagValue float64, err error) {
	err = cplErr(gdalRasterInterpolateAtGeolocation(b, geolocX, geolocY, srs, interpolation, &realValue, &imagValue, transformerOptions))
	return
}

// The VRTProcessedDataset function API (VRTPDWorkingDataPtr, the Init/Free/
// Process callback typedefs and GDALVRTRegisterProcessedDatasetFunc) is
// deferred: it registers Go callbacks as C function pointers (//export).

func gdalGetMaskBand(band GDALRasterBand) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetMaskBand(band.cValue)}
	return
}

func (b GDALRasterBand) GetMaskBand() (result GDALRasterBand, err error) {
	result = gdalGetMaskBand(b)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetMaskFlags(band GDALRasterBand) (result int) {
	result = int(C.GDALGetMaskFlags(band.cValue))
	return
}

func (b GDALRasterBand) GetMaskFlags() (result int) {
	result = gdalGetMaskFlags(b)
	return
}

func gdalCreateMaskBand(band GDALRasterBand, flags int) (result CPLErr) {
	result = CPLErr(C.GDALCreateMaskBand(band.cValue, C.int(flags)))
	return
}

func (b GDALRasterBand) CreateMaskBand(flags int) (err error) {
	err = cplErr(gdalCreateMaskBand(b, flags))
	return
}

func gdalIsMaskBand(band GDALRasterBand) (result bool) {
	result = bool(C.GDALIsMaskBand(band.cValue))
	return
}

func (b GDALRasterBand) IsMaskBand() (result bool) {
	result = gdalIsMaskBand(b)
	return
}

// Flags returned by GetMaskFlags.
const (
	GMFAllValid   = C.GMF_ALL_VALID
	GMFPerDataset = C.GMF_PER_DATASET
	GMFAlpha      = C.GMF_ALPHA
	GMFNoData     = C.GMF_NODATA
)

// Flags returned by GetDataCoverageStatus.
const (
	GDALDataCoverageStatusUnimplemented = C.GDAL_DATA_COVERAGE_STATUS_UNIMPLEMENTED
	GDALDataCoverageStatusData          = C.GDAL_DATA_COVERAGE_STATUS_DATA
	GDALDataCoverageStatusEmpty         = C.GDAL_DATA_COVERAGE_STATUS_EMPTY
)

func gdalGetDataCoverageStatus(band GDALRasterBand, xOff, yOff, xSize, ySize, maskFlagStop int, dataPct *float64) (result int) {
	var cDataPct C.double
	result = int(C.GDALGetDataCoverageStatus(band.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(maskFlagStop), &cDataPct))
	*dataPct = float64(cDataPct)
	return
}

func (b GDALRasterBand) GetDataCoverageStatus(xOff, yOff, xSize, ySize, maskFlagStop int) (status int, dataPct float64) {
	status = gdalGetDataCoverageStatus(b, xOff, yOff, xSize, ySize, maskFlagStop, &dataPct)
	return
}

func gdalComputedRasterBandRelease(band GDALComputedRasterBand) {
	C.GDALComputedRasterBandRelease(band.cValue)
}

func (b GDALComputedRasterBand) Release() {
	gdalComputedRasterBandRelease(b)
}

// Raster algebra unary operation.
type GDALRasterAlgebraUnaryOperation C.GDALRasterAlgebraUnaryOperation

const (
	GRAUOLogicalNot GDALRasterAlgebraUnaryOperation = C.GRAUO_LOGICAL_NOT
	GRAUOAbs        GDALRasterAlgebraUnaryOperation = C.GRAUO_ABS
	GRAUOSqrt       GDALRasterAlgebraUnaryOperation = C.GRAUO_SQRT
	GRAUOLog        GDALRasterAlgebraUnaryOperation = C.GRAUO_LOG
	GRAUOLog10      GDALRasterAlgebraUnaryOperation = C.GRAUO_LOG10
)

func gdalRasterBandUnaryOp(band GDALRasterBand, op GDALRasterAlgebraUnaryOperation) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandUnaryOp(band.cValue, C.GDALRasterAlgebraUnaryOperation(op))}
	return
}

func (b GDALRasterBand) UnaryOp(op GDALRasterAlgebraUnaryOperation) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandUnaryOp(b, op)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// Raster algebra binary operation.
type GDALRasterAlgebraBinaryOperation C.GDALRasterAlgebraBinaryOperation

const (
	GRABOAdd        GDALRasterAlgebraBinaryOperation = C.GRABO_ADD
	GRABOSub        GDALRasterAlgebraBinaryOperation = C.GRABO_SUB
	GRABOMul        GDALRasterAlgebraBinaryOperation = C.GRABO_MUL
	GRABODiv        GDALRasterAlgebraBinaryOperation = C.GRABO_DIV
	GRABOPow        GDALRasterAlgebraBinaryOperation = C.GRABO_POW
	GRABOGt         GDALRasterAlgebraBinaryOperation = C.GRABO_GT
	GRABOGe         GDALRasterAlgebraBinaryOperation = C.GRABO_GE
	GRABOLt         GDALRasterAlgebraBinaryOperation = C.GRABO_LT
	GRABOLe         GDALRasterAlgebraBinaryOperation = C.GRABO_LE
	GRABOEq         GDALRasterAlgebraBinaryOperation = C.GRABO_EQ
	GRABONe         GDALRasterAlgebraBinaryOperation = C.GRABO_NE
	GRABOLogicalAnd GDALRasterAlgebraBinaryOperation = C.GRABO_LOGICAL_AND
	GRABOLogicalOr  GDALRasterAlgebraBinaryOperation = C.GRABO_LOGICAL_OR
)

func gdalRasterBandBinaryOpBand(band GDALRasterBand, op GDALRasterAlgebraBinaryOperation, otherBand GDALRasterBand) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandBinaryOpBand(band.cValue, C.GDALRasterAlgebraBinaryOperation(op), otherBand.cValue)}
	return
}

func (b GDALRasterBand) BinaryOpBand(op GDALRasterAlgebraBinaryOperation, otherBand GDALRasterBand) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandBinaryOpBand(b, op, otherBand)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalRasterBandBinaryOpDouble(band GDALRasterBand, op GDALRasterAlgebraBinaryOperation, constant float64) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandBinaryOpDouble(band.cValue, C.GDALRasterAlgebraBinaryOperation(op), C.double(constant))}
	return
}

func (b GDALRasterBand) BinaryOpDouble(op GDALRasterAlgebraBinaryOperation, constant float64) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandBinaryOpDouble(b, op, constant)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalRasterBandBinaryOpDoubleToBand(constant float64, op GDALRasterAlgebraBinaryOperation, band GDALRasterBand) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandBinaryOpDoubleToBand(C.double(constant), C.GDALRasterAlgebraBinaryOperation(op), band.cValue)}
	return
}

func GDALRasterBandBinaryOpDoubleToBand(constant float64, op GDALRasterAlgebraBinaryOperation, band GDALRasterBand) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandBinaryOpDoubleToBand(constant, op, band)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalRasterBandIfThenElse(condBand, thenBand, elseBand GDALRasterBand) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandIfThenElse(condBand.cValue, thenBand.cValue, elseBand.cValue)}
	return
}

func (b GDALRasterBand) IfThenElse(thenBand, elseBand GDALRasterBand) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandIfThenElse(b, thenBand, elseBand)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalRasterBandAsDataType(band GDALRasterBand, dataType GDALDataType) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandAsDataType(band.cValue, C.GDALDataType(dataType))}
	return
}

func (b GDALRasterBand) AsDataType(dataType GDALDataType) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandAsDataType(b, dataType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMaximumOfNBands(bandCount int, bands GDALRasterBands) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALMaximumOfNBands(C.size_t(bandCount), bands.cPtr())}
	return
}

func GDALMaximumOfNBands(bands GDALRasterBands) (result GDALComputedRasterBand, err error) {
	result = gdalMaximumOfNBands(len(bands), bands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalRasterBandMaxConstant(band GDALRasterBand, constant float64) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandMaxConstant(band.cValue, C.double(constant))}
	return
}

func (b GDALRasterBand) MaxConstant(constant float64) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandMaxConstant(b, constant)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMinimumOfNBands(bandCount int, bands GDALRasterBands) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALMinimumOfNBands(C.size_t(bandCount), bands.cPtr())}
	return
}

func GDALMinimumOfNBands(bands GDALRasterBands) (result GDALComputedRasterBand, err error) {
	result = gdalMinimumOfNBands(len(bands), bands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalRasterBandMinConstant(band GDALRasterBand, constant float64) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandMinConstant(band.cValue, C.double(constant))}
	return
}

func (b GDALRasterBand) MinConstant(constant float64) (result GDALComputedRasterBand, err error) {
	result = gdalRasterBandMinConstant(b, constant)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMeanOfNBands(bandCount int, bands GDALRasterBands) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALMeanOfNBands(C.size_t(bandCount), bands.cPtr())}
	return
}

func GDALMeanOfNBands(bands GDALRasterBands) (result GDALComputedRasterBand, err error) {
	result = gdalMeanOfNBands(len(bands), bands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// /* ==================================================================== */
// /*     GDALAsyncReader                                                  */
// /* ==================================================================== */

func gdalARGetNextUpdatedRegion(reader GDALAsyncReader, timeout float64, xBufOff, yBufOff, xBufSize, yBufSize *int) (result GDALAsyncStatusType) {
	var cXOff, cYOff, cXSize, cYSize C.int
	result = GDALAsyncStatusType(C.GDALARGetNextUpdatedRegion(reader.cValue, C.double(timeout), &cXOff, &cYOff, &cXSize, &cYSize))
	*xBufOff = int(cXOff)
	*yBufOff = int(cYOff)
	*xBufSize = int(cXSize)
	*yBufSize = int(cYSize)
	return
}

func (r GDALAsyncReader) GetNextUpdatedRegion(timeout float64) (status GDALAsyncStatusType, xBufOff, yBufOff, xBufSize, yBufSize int) {
	status = gdalARGetNextUpdatedRegion(r, timeout, &xBufOff, &yBufOff, &xBufSize, &yBufSize)
	return
}

func gdalARLockBuffer(reader GDALAsyncReader, timeout float64) (result bool) {
	result = C.GDALARLockBuffer(reader.cValue, C.double(timeout)) != 0
	return
}

func (r GDALAsyncReader) LockBuffer(timeout float64) (result bool) {
	result = gdalARLockBuffer(r, timeout)
	return
}

func gdalARUnlockBuffer(reader GDALAsyncReader) {
	C.GDALARUnlockBuffer(reader.cValue)
}

func (r GDALAsyncReader) UnlockBuffer() {
	gdalARUnlockBuffer(r)
}

// Helper functions.

// GDALGeneralCmdLineProcessor is deferred: its char*** in/out argv needs a
// dedicated design.

func gdalSwapWords(data unsafe.Pointer, wordSize, wordCount, wordSkip int) {
	C.GDALSwapWords(data, C.int(wordSize), C.int(wordCount), C.int(wordSkip))
}

func GDALSwapWords(data []byte, wordSize, wordCount, wordSkip int) {
	gdalSwapWords(cBytes(data), wordSize, wordCount, wordSkip)
}

func gdalSwapWordsEx(data unsafe.Pointer, wordSize int, wordCount int, wordSkip int) {
	C.GDALSwapWordsEx(data, C.int(wordSize), C.size_t(wordCount), C.int(wordSkip))
}

func GDALSwapWordsEx(data []byte, wordSize, wordCount, wordSkip int) {
	gdalSwapWordsEx(cBytes(data), wordSize, wordCount, wordSkip)
}

func gdalCopyWords(src unsafe.Pointer, srcType GDALDataType, srcPixelOffset int, dst unsafe.Pointer, dstType GDALDataType, dstPixelOffset, wordCount int) {
	C.GDALCopyWords(src, C.GDALDataType(srcType), C.int(srcPixelOffset), dst, C.GDALDataType(dstType), C.int(dstPixelOffset), C.int(wordCount))
}

func GDALCopyWords(src []byte, srcType GDALDataType, srcPixelOffset int, dst []byte, dstType GDALDataType, dstPixelOffset, wordCount int) {
	gdalCopyWords(cBytes(src), srcType, srcPixelOffset, cBytes(dst), dstType, dstPixelOffset, wordCount)
}

func gdalCopyWords64(src unsafe.Pointer, srcType GDALDataType, srcPixelOffset int, dst unsafe.Pointer, dstType GDALDataType, dstPixelOffset int, wordCount int64) {
	C.GDALCopyWords64(src, C.GDALDataType(srcType), C.int(srcPixelOffset), dst, C.GDALDataType(dstType), C.int(dstPixelOffset), C.GPtrDiff_t(wordCount))
}

func GDALCopyWords64(src []byte, srcType GDALDataType, srcPixelOffset int, dst []byte, dstType GDALDataType, dstPixelOffset int, wordCount int64) {
	gdalCopyWords64(cBytes(src), srcType, srcPixelOffset, cBytes(dst), dstType, dstPixelOffset, wordCount)
}

func gdalCopyBits(src unsafe.Pointer, srcOffset, srcStep int, dst unsafe.Pointer, dstOffset, dstStep, bitCount, stepCount int) {
	C.GDALCopyBits((*C.GByte)(src), C.int(srcOffset), C.int(srcStep), (*C.GByte)(dst), C.int(dstOffset), C.int(dstStep), C.int(bitCount), C.int(stepCount))
}

func GDALCopyBits(src []byte, srcOffset, srcStep int, dst []byte, dstOffset, dstStep, bitCount, stepCount int) {
	gdalCopyBits(cBytes(src), srcOffset, srcStep, cBytes(dst), dstOffset, dstStep, bitCount, stepCount)
}

// GDALDeinterleave is deferred: its void **ppDestBuffer output array of buffers
// needs a dedicated design.

func gdalTranspose2D(src unsafe.Pointer, srcType GDALDataType, dst unsafe.Pointer, dstType GDALDataType, srcWidth, srcHeight int) {
	C.GDALTranspose2D(src, C.GDALDataType(srcType), dst, C.GDALDataType(dstType), C.size_t(srcWidth), C.size_t(srcHeight))
}

func GDALTranspose2D(src []byte, srcType GDALDataType, dst []byte, dstType GDALDataType, srcWidth, srcHeight int) {
	gdalTranspose2D(cBytes(src), srcType, cBytes(dst), dstType, srcWidth, srcHeight)
}

func gdalGetNoDataReplacementValue(dataType GDALDataType, value float64) (result float64) {
	result = float64(C.GDALGetNoDataReplacementValue(C.GDALDataType(dataType), C.double(value)))
	return
}

func (dt GDALDataType) GetNoDataReplacementValue(value float64) (result float64) {
	result = gdalGetNoDataReplacementValue(dt, value)
	return
}

func gdalLoadWorldFile(filename string, geoTransform *[6]float64) (result int) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	var gt [6]C.double
	result = int(C.GDALLoadWorldFile(cName, &gt[0]))
	for i := range gt {
		geoTransform[i] = float64(gt[i])
	}
	return
}

func GDALLoadWorldFile(filename string) (geoTransform [6]float64, ok bool) {
	ok = gdalLoadWorldFile(filename, &geoTransform) != 0
	return
}

func gdalReadWorldFile(baseFilename, extension string, geoTransform *[6]float64) (result int) {
	cBase := C.CString(baseFilename)
	defer C.free(unsafe.Pointer(cBase))
	cExt := C.CString(extension)
	defer C.free(unsafe.Pointer(cExt))
	var gt [6]C.double
	result = int(C.GDALReadWorldFile(cBase, cExt, &gt[0]))
	for i := range gt {
		geoTransform[i] = float64(gt[i])
	}
	return
}

func GDALReadWorldFile(baseFilename, extension string) (geoTransform [6]float64, ok bool) {
	ok = gdalReadWorldFile(baseFilename, extension, &geoTransform) != 0
	return
}

func gdalWriteWorldFile(baseFilename, extension string, geoTransform [6]float64) (result int) {
	cBase := C.CString(baseFilename)
	defer C.free(unsafe.Pointer(cBase))
	cExt := C.CString(extension)
	defer C.free(unsafe.Pointer(cExt))
	var gt [6]C.double
	for i, v := range geoTransform {
		gt[i] = C.double(v)
	}
	result = int(C.GDALWriteWorldFile(cBase, cExt, &gt[0]))
	return
}

func GDALWriteWorldFile(baseFilename, extension string, geoTransform [6]float64) (ok bool) {
	ok = gdalWriteWorldFile(baseFilename, extension, geoTransform) != 0
	return
}

// GDALLoadTabFile, GDALReadTabFile, GDALLoadOziMapFile and GDALReadOziMapFile
// are deferred: their char**/int*/GDAL_GCP** in/out parameters need a dedicated
// design.

func gdalDecToDMS(angle float64, axis string, precision int) (result string) {
	cAxis := C.CString(axis)
	defer C.free(unsafe.Pointer(cAxis))
	result = C.GoString(C.GDALDecToDMS(C.double(angle), cAxis, C.int(precision)))
	return
}

func GDALDecToDMS(angle float64, axis string, precision int) (result string) {
	result = gdalDecToDMS(angle, axis, precision)
	return
}

func gdalPackedDMSToDec(packed float64) (result float64) {
	result = float64(C.GDALPackedDMSToDec(C.double(packed)))
	return
}

func GDALPackedDMSToDec(packed float64) (result float64) {
	result = gdalPackedDMSToDec(packed)
	return
}

func gdalDecToPackedDMS(dec float64) (result float64) {
	result = float64(C.GDALDecToPackedDMS(C.double(dec)))
	return
}

func GDALDecToPackedDMS(dec float64) (result float64) {
	result = gdalDecToPackedDMS(dec)
	return
}

func gdalVersionInfo(request string) (result string) {
	cRequest := C.CString(request)
	defer C.free(unsafe.Pointer(cRequest))
	result = C.GoString(C.GDALVersionInfo(cRequest))
	return
}

func GDALVersionInfo(request string) (result string) {
	result = gdalVersionInfo(request)
	return
}

func gdalCheckVersion(versionMajor, versionMinor int, callingComponentName string) (result bool) {
	cName := C.CString(callingComponentName)
	defer C.free(unsafe.Pointer(cName))
	result = C.GDALCheckVersion(C.int(versionMajor), C.int(versionMinor), cName) != 0
	return
}

func GDALCheckVersion(versionMajor, versionMinor int, callingComponentName string) (result bool) {
	result = gdalCheckVersion(versionMajor, versionMinor, callingComponentName)
	return
}

// GDALRPCInfoV1 is a deprecated Rational Polynomial Coefficients structure.
type GDALRPCInfoV1 struct {
	cValue C.GDALRPCInfoV1
}

// GDALRPCInfoV2 stores Rational Polynomial Coefficients / Rigorous Projection
// Model.
type GDALRPCInfoV2 struct {
	cValue C.GDALRPCInfoV2
}

func gdalExtractRPCInfoV1(metadata []string, rpcInfo *GDALRPCInfoV1) (result int) {
	md, free := cStrings(metadata)
	defer free()
	result = int(C.GDALExtractRPCInfoV1(C.CSLConstList(unsafe.Pointer(md)), &rpcInfo.cValue))
	return
}

func GDALExtractRPCInfoV1(metadata []string) (rpcInfo GDALRPCInfoV1, ok bool) {
	ok = gdalExtractRPCInfoV1(metadata, &rpcInfo) != 0
	return
}

func gdalExtractRPCInfoV2(metadata []string, rpcInfo *GDALRPCInfoV2) (result int) {
	md, free := cStrings(metadata)
	defer free()
	result = int(C.GDALExtractRPCInfoV2(C.CSLConstList(unsafe.Pointer(md)), &rpcInfo.cValue))
	return
}

func GDALExtractRPCInfoV2(metadata []string) (rpcInfo GDALRPCInfoV2, ok bool) {
	ok = gdalExtractRPCInfoV2(metadata, &rpcInfo) != 0
	return
}

// /* ==================================================================== */
// /*      Color tables.                                                   */
// /* ==================================================================== */

// GDALColorEntry is a color tuple (c1..c4).
type GDALColorEntry struct {
	cValue C.GDALColorEntry
}

func gdalCreateColorTable(paletteInterp GDALPaletteInterp) (result GDALColorTable) {
	result = GDALColorTable{cValue: C.GDALCreateColorTable(C.GDALPaletteInterp(paletteInterp))}
	return
}

func GDALCreateColorTable(paletteInterp GDALPaletteInterp) (result GDALColorTable, err error) {
	result = gdalCreateColorTable(paletteInterp)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDestroyColorTable(colorTable GDALColorTable) {
	C.GDALDestroyColorTable(colorTable.cValue)
}

func (ct GDALColorTable) Destroy() {
	gdalDestroyColorTable(ct)
}

func gdalCloneColorTable(colorTable GDALColorTable) (result GDALColorTable) {
	result = GDALColorTable{cValue: C.GDALCloneColorTable(colorTable.cValue)}
	return
}

func (ct GDALColorTable) Clone() (result GDALColorTable, err error) {
	result = gdalCloneColorTable(ct)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetPaletteInterpretation(colorTable GDALColorTable) (result GDALPaletteInterp) {
	result = GDALPaletteInterp(C.GDALGetPaletteInterpretation(colorTable.cValue))
	return
}

func (ct GDALColorTable) GetPaletteInterpretation() (result GDALPaletteInterp) {
	result = gdalGetPaletteInterpretation(ct)
	return
}

func gdalGetColorEntryCount(colorTable GDALColorTable) (result int) {
	result = int(C.GDALGetColorEntryCount(colorTable.cValue))
	return
}

func (ct GDALColorTable) GetColorEntryCount() (result int) {
	result = gdalGetColorEntryCount(ct)
	return
}

func gdalGetColorEntry(colorTable GDALColorTable, index int) (result GDALColorEntry) {
	entry := C.GDALGetColorEntry(colorTable.cValue, C.int(index))
	if entry != nil {
		result.cValue = *entry
	}
	return
}

func (ct GDALColorTable) GetColorEntry(index int) (result GDALColorEntry) {
	result = gdalGetColorEntry(ct, index)
	return
}

func gdalGetColorEntryAsRGB(colorTable GDALColorTable, index int, result *GDALColorEntry) (ret int) {
	ret = int(C.GDALGetColorEntryAsRGB(colorTable.cValue, C.int(index), &result.cValue))
	return
}

func (ct GDALColorTable) GetColorEntryAsRGB(index int) (result GDALColorEntry, ok bool) {
	ok = gdalGetColorEntryAsRGB(ct, index, &result) != 0
	return
}

func gdalSetColorEntry(colorTable GDALColorTable, index int, entry GDALColorEntry) {
	C.GDALSetColorEntry(colorTable.cValue, C.int(index), &entry.cValue)
}

func (ct GDALColorTable) SetColorEntry(index int, entry GDALColorEntry) {
	gdalSetColorEntry(ct, index, entry)
}

func gdalCreateColorRamp(colorTable GDALColorTable, startIndex int, startColor GDALColorEntry, endIndex int, endColor GDALColorEntry) {
	C.GDALCreateColorRamp(colorTable.cValue, C.int(startIndex), &startColor.cValue, C.int(endIndex), &endColor.cValue)
}

func (ct GDALColorTable) CreateColorRamp(startIndex int, startColor GDALColorEntry, endIndex int, endColor GDALColorEntry) {
	gdalCreateColorRamp(ct, startIndex, startColor, endIndex, endColor)
}

// /* ==================================================================== */
// /*      Raster Attribute Table                                          */
// /* ==================================================================== */

// Field type of raster attribute table.
type GDALRATFieldType C.GDALRATFieldType

const (
	GFTInteger     GDALRATFieldType = C.GFT_Integer
	GFTReal        GDALRATFieldType = C.GFT_Real
	GFTString      GDALRATFieldType = C.GFT_String
	GFTBoolean     GDALRATFieldType = C.GFT_Boolean
	GFTDateTime    GDALRATFieldType = C.GFT_DateTime
	GFTWKBGeometry GDALRATFieldType = C.GFT_WKBGeometry
)

const GFTMaxCount = C.GFT_MaxCount

// Field usage of raster attribute table.
type GDALRATFieldUsage C.GDALRATFieldUsage

const (
	GFUGeneric    GDALRATFieldUsage = C.GFU_Generic
	GFUPixelCount GDALRATFieldUsage = C.GFU_PixelCount
	GFUName       GDALRATFieldUsage = C.GFU_Name
	GFUMin        GDALRATFieldUsage = C.GFU_Min
	GFUMax        GDALRATFieldUsage = C.GFU_Max
	GFUMinMax     GDALRATFieldUsage = C.GFU_MinMax
	GFURed        GDALRATFieldUsage = C.GFU_Red
	GFUGreen      GDALRATFieldUsage = C.GFU_Green
	GFUBlue       GDALRATFieldUsage = C.GFU_Blue
	GFUAlpha      GDALRATFieldUsage = C.GFU_Alpha
	GFURedMin     GDALRATFieldUsage = C.GFU_RedMin
	GFUGreenMin   GDALRATFieldUsage = C.GFU_GreenMin
	GFUBlueMin    GDALRATFieldUsage = C.GFU_BlueMin
	GFUAlphaMin   GDALRATFieldUsage = C.GFU_AlphaMin
	GFURedMax     GDALRATFieldUsage = C.GFU_RedMax
	GFUGreenMax   GDALRATFieldUsage = C.GFU_GreenMax
	GFUBlueMax    GDALRATFieldUsage = C.GFU_BlueMax
	GFUAlphaMax   GDALRATFieldUsage = C.GFU_AlphaMax
	GFUMaxCount   GDALRATFieldUsage = C.GFU_MaxCount
)

// RAT table type (thematic or athematic).
type GDALRATTableType C.GDALRATTableType

const (
	GRTTThematic  GDALRATTableType = C.GRTT_THEMATIC
	GRTTAthematic GDALRATTableType = C.GRTT_ATHEMATIC
)

func gdalCreateRasterAttributeTable() (result GDALRasterAttributeTable) {
	result = GDALRasterAttributeTable{cValue: C.GDALCreateRasterAttributeTable()}
	return
}

func GDALCreateRasterAttributeTable() (result GDALRasterAttributeTable, err error) {
	result = gdalCreateRasterAttributeTable()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDestroyRasterAttributeTable(rat GDALRasterAttributeTable) {
	C.GDALDestroyRasterAttributeTable(rat.cValue)
}

func (rat GDALRasterAttributeTable) Destroy() {
	gdalDestroyRasterAttributeTable(rat)
}

func gdalRATGetColumnCount(rat GDALRasterAttributeTable) (result int) {
	result = int(C.GDALRATGetColumnCount(rat.cValue))
	return
}

func (rat GDALRasterAttributeTable) GetColumnCount() (result int) {
	result = gdalRATGetColumnCount(rat)
	return
}

func gdalRATGetNameOfCol(rat GDALRasterAttributeTable, col int) (result string) {
	result = C.GoString(C.GDALRATGetNameOfCol(rat.cValue, C.int(col)))
	return
}

func (rat GDALRasterAttributeTable) GetNameOfCol(col int) (result string) {
	result = gdalRATGetNameOfCol(rat, col)
	return
}

func gdalRATGetUsageOfCol(rat GDALRasterAttributeTable, col int) (result GDALRATFieldUsage) {
	result = GDALRATFieldUsage(C.GDALRATGetUsageOfCol(rat.cValue, C.int(col)))
	return
}

func (rat GDALRasterAttributeTable) GetUsageOfCol(col int) (result GDALRATFieldUsage) {
	result = gdalRATGetUsageOfCol(rat, col)
	return
}

func gdalRATGetTypeOfCol(rat GDALRasterAttributeTable, col int) (result GDALRATFieldType) {
	result = GDALRATFieldType(C.GDALRATGetTypeOfCol(rat.cValue, C.int(col)))
	return
}

func (rat GDALRasterAttributeTable) GetTypeOfCol(col int) (result GDALRATFieldType) {
	result = gdalRATGetTypeOfCol(rat, col)
	return
}

func gdalGetRATFieldTypeName(fieldType GDALRATFieldType) (result string) {
	result = C.GoString(C.GDALGetRATFieldTypeName(C.GDALRATFieldType(fieldType)))
	return
}

func (ft GDALRATFieldType) GetName() (result string) {
	result = gdalGetRATFieldTypeName(ft)
	return
}

func gdalGetRATFieldUsageName(usage GDALRATFieldUsage) (result string) {
	result = C.GoString(C.GDALGetRATFieldUsageName(C.GDALRATFieldUsage(usage)))
	return
}

func (fu GDALRATFieldUsage) GetName() (result string) {
	result = gdalGetRATFieldUsageName(fu)
	return
}

func gdalRATGetColOfUsage(rat GDALRasterAttributeTable, usage GDALRATFieldUsage) (result int) {
	result = int(C.GDALRATGetColOfUsage(rat.cValue, C.GDALRATFieldUsage(usage)))
	return
}

func (rat GDALRasterAttributeTable) GetColOfUsage(usage GDALRATFieldUsage) (result int) {
	result = gdalRATGetColOfUsage(rat, usage)
	return
}

func gdalRATGetRowCount(rat GDALRasterAttributeTable) (result int) {
	result = int(C.GDALRATGetRowCount(rat.cValue))
	return
}

func (rat GDALRasterAttributeTable) GetRowCount() (result int) {
	result = gdalRATGetRowCount(rat)
	return
}

func gdalRATGetValueAsString(rat GDALRasterAttributeTable, row, field int) (result string) {
	result = C.GoString(C.GDALRATGetValueAsString(rat.cValue, C.int(row), C.int(field)))
	return
}

func (rat GDALRasterAttributeTable) GetValueAsString(row, field int) (result string) {
	result = gdalRATGetValueAsString(rat, row, field)
	return
}

func gdalRATGetValueAsInt(rat GDALRasterAttributeTable, row, field int) (result int) {
	result = int(C.GDALRATGetValueAsInt(rat.cValue, C.int(row), C.int(field)))
	return
}

func (rat GDALRasterAttributeTable) GetValueAsInt(row, field int) (result int) {
	result = gdalRATGetValueAsInt(rat, row, field)
	return
}

func gdalRATGetValueAsDouble(rat GDALRasterAttributeTable, row, field int) (result float64) {
	result = float64(C.GDALRATGetValueAsDouble(rat.cValue, C.int(row), C.int(field)))
	return
}

func (rat GDALRasterAttributeTable) GetValueAsDouble(row, field int) (result float64) {
	result = gdalRATGetValueAsDouble(rat, row, field)
	return
}

func gdalRATGetValueAsBoolean(rat GDALRasterAttributeTable, row, field int) (result bool) {
	result = bool(C.GDALRATGetValueAsBoolean(rat.cValue, C.int(row), C.int(field)))
	return
}

func (rat GDALRasterAttributeTable) GetValueAsBoolean(row, field int) (result bool) {
	result = gdalRATGetValueAsBoolean(rat, row, field)
	return
}

// GDALRATDateTime encodes a DateTime field for a GDAL Raster Attribute Table.
type GDALRATDateTime struct {
	cValue C.GDALRATDateTime
}

func gdalRATGetValueAsDateTime(rat GDALRasterAttributeTable, row, field int, dateTime *GDALRATDateTime) (result CPLErr) {
	result = CPLErr(C.GDALRATGetValueAsDateTime(rat.cValue, C.int(row), C.int(field), &dateTime.cValue))
	return
}

func (rat GDALRasterAttributeTable) GetValueAsDateTime(row, field int) (dateTime GDALRATDateTime, err error) {
	err = cplErr(gdalRATGetValueAsDateTime(rat, row, field, &dateTime))
	return
}

func gdalRATGetValueAsWKBGeometry(rat GDALRasterAttributeTable, row, field int) (result []byte) {
	var cSize C.size_t
	ptr := C.GDALRATGetValueAsWKBGeometry(rat.cValue, C.int(row), C.int(field), &cSize)
	if ptr != nil {
		result = C.GoBytes(unsafe.Pointer(ptr), C.int(cSize))
	}
	return
}

func (rat GDALRasterAttributeTable) GetValueAsWKBGeometry(row, field int) (result []byte) {
	result = gdalRATGetValueAsWKBGeometry(rat, row, field)
	return
}

func gdalRATSetValueAsString(rat GDALRasterAttributeTable, row, field int, value string) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.GDALRATSetValueAsString(rat.cValue, C.int(row), C.int(field), cValue)
}

func (rat GDALRasterAttributeTable) SetValueAsString(row, field int, value string) {
	gdalRATSetValueAsString(rat, row, field, value)
}

func gdalRATSetValueAsInt(rat GDALRasterAttributeTable, row, field, value int) {
	C.GDALRATSetValueAsInt(rat.cValue, C.int(row), C.int(field), C.int(value))
}

func (rat GDALRasterAttributeTable) SetValueAsInt(row, field, value int) {
	gdalRATSetValueAsInt(rat, row, field, value)
}

func gdalRATSetValueAsDouble(rat GDALRasterAttributeTable, row, field int, value float64) {
	C.GDALRATSetValueAsDouble(rat.cValue, C.int(row), C.int(field), C.double(value))
}

func (rat GDALRasterAttributeTable) SetValueAsDouble(row, field int, value float64) {
	gdalRATSetValueAsDouble(rat, row, field, value)
}

func gdalRATSetValueAsBoolean(rat GDALRasterAttributeTable, row, field int, value bool) (result CPLErr) {
	result = CPLErr(C.GDALRATSetValueAsBoolean(rat.cValue, C.int(row), C.int(field), C.bool(value)))
	return
}

func (rat GDALRasterAttributeTable) SetValueAsBoolean(row, field int, value bool) (err error) {
	err = cplErr(gdalRATSetValueAsBoolean(rat, row, field, value))
	return
}

func gdalRATSetValueAsDateTime(rat GDALRasterAttributeTable, row, field int, dateTime GDALRATDateTime) (result CPLErr) {
	result = CPLErr(C.GDALRATSetValueAsDateTime(rat.cValue, C.int(row), C.int(field), &dateTime.cValue))
	return
}

func (rat GDALRasterAttributeTable) SetValueAsDateTime(row, field int, dateTime GDALRATDateTime) (err error) {
	err = cplErr(gdalRATSetValueAsDateTime(rat, row, field, dateTime))
	return
}

func gdalRATSetValueAsWKBGeometry(rat GDALRasterAttributeTable, row, field int, wkb unsafe.Pointer, wkbSize int) (result CPLErr) {
	result = CPLErr(C.GDALRATSetValueAsWKBGeometry(rat.cValue, C.int(row), C.int(field), wkb, C.size_t(wkbSize)))
	return
}

func (rat GDALRasterAttributeTable) SetValueAsWKBGeometry(row, field int, wkb []byte) (err error) {
	err = cplErr(gdalRATSetValueAsWKBGeometry(rat, row, field, cBytes(wkb), len(wkb)))
	return
}

func gdalRATChangesAreWrittenToFile(rat GDALRasterAttributeTable) (result bool) {
	result = C.GDALRATChangesAreWrittenToFile(rat.cValue) != 0
	return
}

func (rat GDALRasterAttributeTable) ChangesAreWrittenToFile() (result bool) {
	result = gdalRATChangesAreWrittenToFile(rat)
	return
}

func gdalRATValuesIOAsDouble(rat GDALRasterAttributeTable, rwFlag GDALRWFlag, field, startRow, length int, data []float64) (result CPLErr) {
	cData := make([]C.double, length)
	for i := 0; i < length && i < len(data); i++ {
		cData[i] = C.double(data[i])
	}
	var ptr *C.double
	if length > 0 {
		ptr = &cData[0]
	}
	result = CPLErr(C.GDALRATValuesIOAsDouble(rat.cValue, C.GDALRWFlag(rwFlag), C.int(field), C.int(startRow), C.int(length), ptr))
	for i := 0; i < length && i < len(data); i++ {
		data[i] = float64(cData[i])
	}
	return
}

func (rat GDALRasterAttributeTable) ValuesIOAsDouble(rwFlag GDALRWFlag, field, startRow int, data []float64) (err error) {
	err = cplErr(gdalRATValuesIOAsDouble(rat, rwFlag, field, startRow, len(data), data))
	return
}

func gdalRATValuesIOAsInteger(rat GDALRasterAttributeTable, rwFlag GDALRWFlag, field, startRow, length int, data []int) (result CPLErr) {
	cData := make([]C.int, length)
	for i := 0; i < length && i < len(data); i++ {
		cData[i] = C.int(data[i])
	}
	var ptr *C.int
	if length > 0 {
		ptr = &cData[0]
	}
	result = CPLErr(C.GDALRATValuesIOAsInteger(rat.cValue, C.GDALRWFlag(rwFlag), C.int(field), C.int(startRow), C.int(length), ptr))
	for i := 0; i < length && i < len(data); i++ {
		data[i] = int(cData[i])
	}
	return
}

func (rat GDALRasterAttributeTable) ValuesIOAsInteger(rwFlag GDALRWFlag, field, startRow int, data []int) (err error) {
	err = cplErr(gdalRATValuesIOAsInteger(rat, rwFlag, field, startRow, len(data), data))
	return
}

// GDALRATValuesIOAsString, GDALRATValuesIOAsBoolean, GDALRATValuesIOAsDateTime
// and GDALRATValuesIOAsWKBGeometry are deferred: their in/out char**, bool*,
// DateTime array and GByte** parameters need a dedicated design.

func gdalRATSetRowCount(rat GDALRasterAttributeTable, count int) {
	C.GDALRATSetRowCount(rat.cValue, C.int(count))
}

func (rat GDALRasterAttributeTable) SetRowCount(count int) {
	gdalRATSetRowCount(rat, count)
}

func gdalRATCreateColumn(rat GDALRasterAttributeTable, name string, fieldType GDALRATFieldType, fieldUsage GDALRATFieldUsage) (result CPLErr) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = CPLErr(C.GDALRATCreateColumn(rat.cValue, cName, C.GDALRATFieldType(fieldType), C.GDALRATFieldUsage(fieldUsage)))
	return
}

func (rat GDALRasterAttributeTable) CreateColumn(name string, fieldType GDALRATFieldType, fieldUsage GDALRATFieldUsage) (err error) {
	err = cplErr(gdalRATCreateColumn(rat, name, fieldType, fieldUsage))
	return
}

func gdalRATSetLinearBinning(rat GDALRasterAttributeTable, row0Min, binSize float64) (result CPLErr) {
	result = CPLErr(C.GDALRATSetLinearBinning(rat.cValue, C.double(row0Min), C.double(binSize)))
	return
}

func (rat GDALRasterAttributeTable) SetLinearBinning(row0Min, binSize float64) (err error) {
	err = cplErr(gdalRATSetLinearBinning(rat, row0Min, binSize))
	return
}

func gdalRATGetLinearBinning(rat GDALRasterAttributeTable, row0Min, binSize *float64) (result int) {
	var cRow0Min, cBinSize C.double
	result = int(C.GDALRATGetLinearBinning(rat.cValue, &cRow0Min, &cBinSize))
	*row0Min = float64(cRow0Min)
	*binSize = float64(cBinSize)
	return
}

func (rat GDALRasterAttributeTable) GetLinearBinning() (row0Min, binSize float64, ok bool) {
	ok = gdalRATGetLinearBinning(rat, &row0Min, &binSize) != 0
	return
}

func gdalRATSetTableType(rat GDALRasterAttributeTable, tableType GDALRATTableType) (result CPLErr) {
	result = CPLErr(C.GDALRATSetTableType(rat.cValue, C.GDALRATTableType(tableType)))
	return
}

func (rat GDALRasterAttributeTable) SetTableType(tableType GDALRATTableType) (err error) {
	err = cplErr(gdalRATSetTableType(rat, tableType))
	return
}

func gdalRATGetTableType(rat GDALRasterAttributeTable) (result GDALRATTableType) {
	result = GDALRATTableType(C.GDALRATGetTableType(rat.cValue))
	return
}

func (rat GDALRasterAttributeTable) GetTableType() (result GDALRATTableType) {
	result = gdalRATGetTableType(rat)
	return
}

func gdalRATInitializeFromColorTable(rat GDALRasterAttributeTable, colorTable GDALColorTable) (result CPLErr) {
	result = CPLErr(C.GDALRATInitializeFromColorTable(rat.cValue, colorTable.cValue))
	return
}

func (rat GDALRasterAttributeTable) InitializeFromColorTable(colorTable GDALColorTable) (err error) {
	err = cplErr(gdalRATInitializeFromColorTable(rat, colorTable))
	return
}

func gdalRATTranslateToColorTable(rat GDALRasterAttributeTable, entryCount int) (result GDALColorTable) {
	result = GDALColorTable{cValue: C.GDALRATTranslateToColorTable(rat.cValue, C.int(entryCount))}
	return
}

func (rat GDALRasterAttributeTable) TranslateToColorTable(entryCount int) (result GDALColorTable, err error) {
	result = gdalRATTranslateToColorTable(rat, entryCount)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalRATDumpReadable(rat GDALRasterAttributeTable, filename string) (err error) {
	fp, closeFn, err := cFile(filename, "w")
	if err != nil {
		return
	}
	defer closeFn()
	C.GDALRATDumpReadable(rat.cValue, fp)
	return
}

func (rat GDALRasterAttributeTable) DumpReadable(filename string) (err error) {
	return gdalRATDumpReadable(rat, filename)
}

func gdalRATClone(rat GDALRasterAttributeTable) (result GDALRasterAttributeTable) {
	result = GDALRasterAttributeTable{cValue: C.GDALRATClone(rat.cValue)}
	return
}

func (rat GDALRasterAttributeTable) Clone() (result GDALRasterAttributeTable, err error) {
	result = gdalRATClone(rat)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// GDALRATSerializeJSON is deferred: its void* return (a json-c object) needs a
// dedicated design.

func gdalRATGetRowOfValue(rat GDALRasterAttributeTable, value float64) (result int) {
	result = int(C.GDALRATGetRowOfValue(rat.cValue, C.double(value)))
	return
}

func (rat GDALRasterAttributeTable) GetRowOfValue(value float64) (result int) {
	result = gdalRATGetRowOfValue(rat, value)
	return
}

func gdalRATRemoveStatistics(rat GDALRasterAttributeTable) {
	C.GDALRATRemoveStatistics(rat.cValue)
}

func (rat GDALRasterAttributeTable) RemoveStatistics() {
	gdalRATRemoveStatistics(rat)
}

// /* -------------------------------------------------------------------- */
// /*                          Relationships                               */
// /* -------------------------------------------------------------------- */

// Cardinality of relationship.
type GDALRelationshipCardinality C.GDALRelationshipCardinality

const (
	GRCOneToOne   GDALRelationshipCardinality = C.GRC_ONE_TO_ONE
	GRCOneToMany  GDALRelationshipCardinality = C.GRC_ONE_TO_MANY
	GRCManyToOne  GDALRelationshipCardinality = C.GRC_MANY_TO_ONE
	GRCManyToMany GDALRelationshipCardinality = C.GRC_MANY_TO_MANY
)

// Type of relationship.
type GDALRelationshipType C.GDALRelationshipType

const (
	GRTComposite   GDALRelationshipType = C.GRT_COMPOSITE
	GRTAssociation GDALRelationshipType = C.GRT_ASSOCIATION
	GRTAggregation GDALRelationshipType = C.GRT_AGGREGATION
)

func gdalRelationshipCreate(name, leftTableName, rightTableName string, cardinality GDALRelationshipCardinality) (result GDALRelationship) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cLeft := C.CString(leftTableName)
	defer C.free(unsafe.Pointer(cLeft))
	cRight := C.CString(rightTableName)
	defer C.free(unsafe.Pointer(cRight))
	result = GDALRelationship{cValue: C.GDALRelationshipCreate(cName, cLeft, cRight, C.GDALRelationshipCardinality(cardinality))}
	return
}

func GDALRelationshipCreate(name, leftTableName, rightTableName string, cardinality GDALRelationshipCardinality) (result GDALRelationship, err error) {
	result = gdalRelationshipCreate(name, leftTableName, rightTableName, cardinality)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDestroyRelationship(relationship GDALRelationship) {
	C.GDALDestroyRelationship(relationship.cValue)
}

func (r GDALRelationship) Destroy() {
	gdalDestroyRelationship(r)
}

func gdalRelationshipGetName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetName(relationship.cValue))
	return
}

func (r GDALRelationship) GetName() (result string) {
	result = gdalRelationshipGetName(r)
	return
}

func gdalRelationshipGetCardinality(relationship GDALRelationship) (result GDALRelationshipCardinality) {
	result = GDALRelationshipCardinality(C.GDALRelationshipGetCardinality(relationship.cValue))
	return
}

func (r GDALRelationship) GetCardinality() (result GDALRelationshipCardinality) {
	result = gdalRelationshipGetCardinality(r)
	return
}

func gdalRelationshipGetLeftTableName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetLeftTableName(relationship.cValue))
	return
}

func (r GDALRelationship) GetLeftTableName() (result string) {
	result = gdalRelationshipGetLeftTableName(r)
	return
}

func gdalRelationshipGetRightTableName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetRightTableName(relationship.cValue))
	return
}

func (r GDALRelationship) GetRightTableName() (result string) {
	result = gdalRelationshipGetRightTableName(r)
	return
}

func gdalRelationshipGetMappingTableName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetMappingTableName(relationship.cValue))
	return
}

func (r GDALRelationship) GetMappingTableName() (result string) {
	result = gdalRelationshipGetMappingTableName(r)
	return
}

func gdalRelationshipSetMappingTableName(relationship GDALRelationship, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.GDALRelationshipSetMappingTableName(relationship.cValue, cName)
}

func (r GDALRelationship) SetMappingTableName(name string) {
	gdalRelationshipSetMappingTableName(r, name)
}

func gdalRelationshipGetLeftTableFields(relationship GDALRelationship) (result []string) {
	raw := C.GDALRelationshipGetLeftTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (r GDALRelationship) GetLeftTableFields() (result []string) {
	result = gdalRelationshipGetLeftTableFields(r)
	return
}

func gdalRelationshipGetRightTableFields(relationship GDALRelationship) (result []string) {
	raw := C.GDALRelationshipGetRightTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (r GDALRelationship) GetRightTableFields() (result []string) {
	result = gdalRelationshipGetRightTableFields(r)
	return
}

func gdalRelationshipSetLeftTableFields(relationship GDALRelationship, fields []string) {
	f, free := cStrings(fields)
	defer free()
	C.GDALRelationshipSetLeftTableFields(relationship.cValue, C.CSLConstList(unsafe.Pointer(f)))
}

func (r GDALRelationship) SetLeftTableFields(fields []string) {
	gdalRelationshipSetLeftTableFields(r, fields)
}

func gdalRelationshipSetRightTableFields(relationship GDALRelationship, fields []string) {
	f, free := cStrings(fields)
	defer free()
	C.GDALRelationshipSetRightTableFields(relationship.cValue, C.CSLConstList(unsafe.Pointer(f)))
}

func (r GDALRelationship) SetRightTableFields(fields []string) {
	gdalRelationshipSetRightTableFields(r, fields)
}

func gdalRelationshipGetLeftMappingTableFields(relationship GDALRelationship) (result []string) {
	raw := C.GDALRelationshipGetLeftMappingTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (r GDALRelationship) GetLeftMappingTableFields() (result []string) {
	result = gdalRelationshipGetLeftMappingTableFields(r)
	return
}

func gdalRelationshipGetRightMappingTableFields(relationship GDALRelationship) (result []string) {
	raw := C.GDALRelationshipGetRightMappingTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (r GDALRelationship) GetRightMappingTableFields() (result []string) {
	result = gdalRelationshipGetRightMappingTableFields(r)
	return
}

func gdalRelationshipSetLeftMappingTableFields(relationship GDALRelationship, fields []string) {
	f, free := cStrings(fields)
	defer free()
	C.GDALRelationshipSetLeftMappingTableFields(relationship.cValue, C.CSLConstList(unsafe.Pointer(f)))
}

func (r GDALRelationship) SetLeftMappingTableFields(fields []string) {
	gdalRelationshipSetLeftMappingTableFields(r, fields)
}

func gdalRelationshipSetRightMappingTableFields(relationship GDALRelationship, fields []string) {
	f, free := cStrings(fields)
	defer free()
	C.GDALRelationshipSetRightMappingTableFields(relationship.cValue, C.CSLConstList(unsafe.Pointer(f)))
}

func (r GDALRelationship) SetRightMappingTableFields(fields []string) {
	gdalRelationshipSetRightMappingTableFields(r, fields)
}

func gdalRelationshipGetType(relationship GDALRelationship) (result GDALRelationshipType) {
	result = GDALRelationshipType(C.GDALRelationshipGetType(relationship.cValue))
	return
}

func (r GDALRelationship) GetType() (result GDALRelationshipType) {
	result = gdalRelationshipGetType(r)
	return
}

func gdalRelationshipSetType(relationship GDALRelationship, relationshipType GDALRelationshipType) {
	C.GDALRelationshipSetType(relationship.cValue, C.GDALRelationshipType(relationshipType))
}

func (r GDALRelationship) SetType(relationshipType GDALRelationshipType) {
	gdalRelationshipSetType(r, relationshipType)
}

func gdalRelationshipGetForwardPathLabel(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetForwardPathLabel(relationship.cValue))
	return
}

func (r GDALRelationship) GetForwardPathLabel() (result string) {
	result = gdalRelationshipGetForwardPathLabel(r)
	return
}

func gdalRelationshipSetForwardPathLabel(relationship GDALRelationship, label string) {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.GDALRelationshipSetForwardPathLabel(relationship.cValue, cLabel)
}

func (r GDALRelationship) SetForwardPathLabel(label string) {
	gdalRelationshipSetForwardPathLabel(r, label)
}

func gdalRelationshipGetBackwardPathLabel(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetBackwardPathLabel(relationship.cValue))
	return
}

func (r GDALRelationship) GetBackwardPathLabel() (result string) {
	result = gdalRelationshipGetBackwardPathLabel(r)
	return
}

func gdalRelationshipSetBackwardPathLabel(relationship GDALRelationship, label string) {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.GDALRelationshipSetBackwardPathLabel(relationship.cValue, cLabel)
}

func (r GDALRelationship) SetBackwardPathLabel(label string) {
	gdalRelationshipSetBackwardPathLabel(r, label)
}

func gdalRelationshipGetRelatedTableType(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetRelatedTableType(relationship.cValue))
	return
}

func (r GDALRelationship) GetRelatedTableType() (result string) {
	result = gdalRelationshipGetRelatedTableType(r)
	return
}

func gdalRelationshipSetRelatedTableType(relationship GDALRelationship, relatedTableType string) {
	cType := C.CString(relatedTableType)
	defer C.free(unsafe.Pointer(cType))
	C.GDALRelationshipSetRelatedTableType(relationship.cValue, cType)
}

func (r GDALRelationship) SetRelatedTableType(relatedTableType string) {
	gdalRelationshipSetRelatedTableType(r, relatedTableType)
}

// /* ==================================================================== */
// /*      GDAL Cache Management                                           */
// /* ==================================================================== */

func gdalSetCacheMax(bytes int) {
	C.GDALSetCacheMax(C.int(bytes))
}

func GDALSetCacheMax(bytes int) {
	gdalSetCacheMax(bytes)
}

func gdalGetCacheMax() (result int) {
	result = int(C.GDALGetCacheMax())
	return
}

func GDALGetCacheMax() (result int) {
	result = gdalGetCacheMax()
	return
}

func gdalGetCacheUsed() (result int) {
	result = int(C.GDALGetCacheUsed())
	return
}

func GDALGetCacheUsed() (result int) {
	result = gdalGetCacheUsed()
	return
}

func gdalSetCacheMax64(bytes int64) {
	C.GDALSetCacheMax64(C.GIntBig(bytes))
}

func GDALSetCacheMax64(bytes int64) {
	gdalSetCacheMax64(bytes)
}

func gdalGetCacheMax64() (result int64) {
	result = int64(C.GDALGetCacheMax64())
	return
}

func GDALGetCacheMax64() (result int64) {
	result = gdalGetCacheMax64()
	return
}

func gdalGetCacheUsed64() (result int64) {
	result = int64(C.GDALGetCacheUsed64())
	return
}

func GDALGetCacheUsed64() (result int64) {
	result = gdalGetCacheUsed64()
	return
}

func gdalFlushCacheBlock() (result bool) {
	result = C.GDALFlushCacheBlock() != 0
	return
}

func GDALFlushCacheBlock() (result bool) {
	result = gdalFlushCacheBlock()
	return
}

// The GDAL virtual memory API (GDALDatasetGetVirtualMem,
// GDALRasterBandGetVirtualMem, GDALGetVirtualMemAuto,
// GDALDatasetGetTiledVirtualMem, GDALRasterBandGetTiledVirtualMem) is deferred:
// it returns CPLVirtualMem*, which needs cpl_virtualmem.go.

// Enumeration to describe the tile organization.
type GDALTileOrganization C.GDALTileOrganization

const (
	GTOTIP GDALTileOrganization = C.GTO_TIP
	GTOBIT GDALTileOrganization = C.GTO_BIT
	GTOBSQ GDALTileOrganization = C.GTO_BSQ
)

func gdalCreatePansharpenedVRT(xml string, panchroBand GDALRasterBand, inputSpectralBandCount int, inputSpectralBands GDALRasterBands) (result GDALDataset) {
	cXML := C.CString(xml)
	defer C.free(unsafe.Pointer(cXML))
	result = GDALDataset{cValue: C.GDALCreatePansharpenedVRT(cXML, panchroBand.cValue, C.int(inputSpectralBandCount), inputSpectralBands.cPtr())}
	return
}

func GDALCreatePansharpenedVRT(xml string, panchroBand GDALRasterBand, inputSpectralBands GDALRasterBands) (result GDALDataset, err error) {
	result = gdalCreatePansharpenedVRT(xml, panchroBand, len(inputSpectralBands), inputSpectralBands)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGetJPEG2000Structure(filename string, options []string) (result CPLXMLNode) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = CPLXMLNode{cValue: C.GDALGetJPEG2000Structure(cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func GDALGetJPEG2000Structure(filename string, options []string) (result CPLXMLNode, err error) {
	result = gdalGetJPEG2000Structure(filename, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// /* ==================================================================== */
// /*      Multidimensional API_api                                       */
// /* ==================================================================== */

func gdalCreateMultiDimensional(driver GDALDriver, name string, rootGroupOptions, options []string) (result GDALDataset) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	rootOpts, freeRoot := cStrings(rootGroupOptions)
	defer freeRoot()
	opts, free := cStrings(options)
	defer free()
	result = GDALDataset{cValue: C.GDALCreateMultiDimensional(driver.cValue, cName, C.CSLConstList(unsafe.Pointer(rootOpts)), C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (d GDALDriver) CreateMultiDimensional(name string, rootGroupOptions, options []string) (result GDALDataset, err error) {
	result = gdalCreateMultiDimensional(d, name, rootGroupOptions, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalExtendedDataTypeCreate(dataType GDALDataType) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALExtendedDataTypeCreate(C.GDALDataType(dataType))}
	return
}

func GDALExtendedDataTypeCreate(dataType GDALDataType) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreate(dataType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalExtendedDataTypeCreateString(maxStringLength int) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALExtendedDataTypeCreateString(C.size_t(maxStringLength))}
	return
}

func GDALExtendedDataTypeCreateString(maxStringLength int) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreateString(maxStringLength)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalExtendedDataTypeCreateStringEx(maxStringLength int, subType GDALExtendedDataTypeSubType) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALExtendedDataTypeCreateStringEx(C.size_t(maxStringLength), C.GDALExtendedDataTypeSubType(subType))}
	return
}

func GDALExtendedDataTypeCreateStringEx(maxStringLength int, subType GDALExtendedDataTypeSubType) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreateStringEx(maxStringLength, subType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalExtendedDataTypeCreateCompound(name string, totalSize, nComponents int, comps []GDALEDTComponent) (result GDALExtendedDataType) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var ptr *C.GDALEDTComponentH
	if len(comps) > 0 {
		ptr = (*C.GDALEDTComponentH)(unsafe.Pointer(&comps[0]))
	}
	result = GDALExtendedDataType{cValue: C.GDALExtendedDataTypeCreateCompound(cName, C.size_t(totalSize), C.size_t(nComponents), ptr)}
	return
}

func GDALExtendedDataTypeCreateCompound(name string, totalSize int, comps []GDALEDTComponent) (result GDALExtendedDataType, err error) {
	result = gdalExtendedDataTypeCreateCompound(name, totalSize, len(comps), comps)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalExtendedDataTypeRelease(edt GDALExtendedDataType) {
	C.GDALExtendedDataTypeRelease(edt.cValue)
}

func (edt GDALExtendedDataType) Release() {
	gdalExtendedDataTypeRelease(edt)
}

func gdalExtendedDataTypeGetName(edt GDALExtendedDataType) (result string) {
	result = C.GoString(C.GDALExtendedDataTypeGetName(edt.cValue))
	return
}

func (edt GDALExtendedDataType) GetName() (result string) {
	result = gdalExtendedDataTypeGetName(edt)
	return
}

func gdalExtendedDataTypeGetClass(edt GDALExtendedDataType) (result GDALExtendedDataTypeClass) {
	result = GDALExtendedDataTypeClass(C.GDALExtendedDataTypeGetClass(edt.cValue))
	return
}

func (edt GDALExtendedDataType) GetClass() (result GDALExtendedDataTypeClass) {
	result = gdalExtendedDataTypeGetClass(edt)
	return
}

func gdalExtendedDataTypeGetNumericDataType(edt GDALExtendedDataType) (result GDALDataType) {
	result = GDALDataType(C.GDALExtendedDataTypeGetNumericDataType(edt.cValue))
	return
}

func (edt GDALExtendedDataType) GetNumericDataType() (result GDALDataType) {
	result = gdalExtendedDataTypeGetNumericDataType(edt)
	return
}

func gdalExtendedDataTypeGetSize(edt GDALExtendedDataType) (result int) {
	result = int(C.GDALExtendedDataTypeGetSize(edt.cValue))
	return
}

func (edt GDALExtendedDataType) GetSize() (result int) {
	result = gdalExtendedDataTypeGetSize(edt)
	return
}

func gdalExtendedDataTypeGetMaxStringLength(edt GDALExtendedDataType) (result int) {
	result = int(C.GDALExtendedDataTypeGetMaxStringLength(edt.cValue))
	return
}

func (edt GDALExtendedDataType) GetMaxStringLength() (result int) {
	result = gdalExtendedDataTypeGetMaxStringLength(edt)
	return
}

// GDALExtendedDataTypeGetComponents and GDALExtendedDataTypeFreeComponents are
// deferred: the returned GDALEDTComponentH* array has release/free ownership
// semantics that need a dedicated design.

func gdalExtendedDataTypeCanConvertTo(sourceEDT, targetEDT GDALExtendedDataType) (result bool) {
	result = C.GDALExtendedDataTypeCanConvertTo(sourceEDT.cValue, targetEDT.cValue) != 0
	return
}

func (edt GDALExtendedDataType) CanConvertTo(targetEDT GDALExtendedDataType) (result bool) {
	result = gdalExtendedDataTypeCanConvertTo(edt, targetEDT)
	return
}

func gdalExtendedDataTypeEquals(firstEDT, secondEDT GDALExtendedDataType) (result bool) {
	result = C.GDALExtendedDataTypeEquals(firstEDT.cValue, secondEDT.cValue) != 0
	return
}

func (edt GDALExtendedDataType) Equals(other GDALExtendedDataType) (result bool) {
	result = gdalExtendedDataTypeEquals(edt, other)
	return
}

func gdalExtendedDataTypeGetSubType(edt GDALExtendedDataType) (result GDALExtendedDataTypeSubType) {
	result = GDALExtendedDataTypeSubType(C.GDALExtendedDataTypeGetSubType(edt.cValue))
	return
}

func (edt GDALExtendedDataType) GetSubType() (result GDALExtendedDataTypeSubType) {
	result = gdalExtendedDataTypeGetSubType(edt)
	return
}

func gdalExtendedDataTypeGetRAT(edt GDALExtendedDataType) (result GDALRasterAttributeTable) {
	result = GDALRasterAttributeTable{cValue: C.GDALExtendedDataTypeGetRAT(edt.cValue)}
	return
}

func (edt GDALExtendedDataType) GetRAT() (result GDALRasterAttributeTable, err error) {
	result = gdalExtendedDataTypeGetRAT(edt)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalEDTComponentCreate(name string, offset int, dataType GDALExtendedDataType) (result GDALEDTComponent) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALEDTComponent{cValue: C.GDALEDTComponentCreate(cName, C.size_t(offset), dataType.cValue)}
	return
}

func GDALEDTComponentCreate(name string, offset int, dataType GDALExtendedDataType) (result GDALEDTComponent, err error) {
	result = gdalEDTComponentCreate(name, offset, dataType)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalEDTComponentRelease(comp GDALEDTComponent) {
	C.GDALEDTComponentRelease(comp.cValue)
}

func (comp GDALEDTComponent) Release() {
	gdalEDTComponentRelease(comp)
}

func gdalEDTComponentGetName(comp GDALEDTComponent) (result string) {
	result = C.GoString(C.GDALEDTComponentGetName(comp.cValue))
	return
}

func (comp GDALEDTComponent) GetName() (result string) {
	result = gdalEDTComponentGetName(comp)
	return
}

func gdalEDTComponentGetOffset(comp GDALEDTComponent) (result int) {
	result = int(C.GDALEDTComponentGetOffset(comp.cValue))
	return
}

func (comp GDALEDTComponent) GetOffset() (result int) {
	result = gdalEDTComponentGetOffset(comp)
	return
}

func gdalEDTComponentGetType(comp GDALEDTComponent) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALEDTComponentGetType(comp.cValue)}
	return
}

func (comp GDALEDTComponent) GetType() (result GDALExtendedDataType, err error) {
	result = gdalEDTComponentGetType(comp)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDatasetGetRootGroup(dataset GDALDataset) (result GDALGroup) {
	result = GDALGroup{cValue: C.GDALDatasetGetRootGroup(dataset.cValue)}
	return
}

func (ds GDALDataset) GetRootGroup() (result GDALGroup, err error) {
	result = gdalDatasetGetRootGroup(ds)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupRelease(group GDALGroup) {
	C.GDALGroupRelease(group.cValue)
}

func (g GDALGroup) Release() {
	gdalGroupRelease(g)
}

func gdalGroupGetName(group GDALGroup) (result string) {
	result = C.GoString(C.GDALGroupGetName(group.cValue))
	return
}

func (g GDALGroup) GetName() (result string) {
	result = gdalGroupGetName(g)
	return
}

func gdalGroupGetFullName(group GDALGroup) (result string) {
	result = C.GoString(C.GDALGroupGetFullName(group.cValue))
	return
}

func (g GDALGroup) GetFullName() (result string) {
	result = gdalGroupGetFullName(g)
	return
}

func gdalGroupGetMDArrayNames(group GDALGroup, options []string) (result []string) {
	opts, free := cStrings(options)
	defer free()
	raw := C.GDALGroupGetMDArrayNames(group.cValue, C.CSLConstList(unsafe.Pointer(opts)))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (g GDALGroup) GetMDArrayNames(options []string) (result []string) {
	result = gdalGroupGetMDArrayNames(g, options)
	return
}

func gdalGroupGetMDArrayFullNamesRecursive(group GDALGroup, groupOptions, arrayOptions []string) (result []string) {
	gOpts, freeG := cStrings(groupOptions)
	defer freeG()
	aOpts, freeA := cStrings(arrayOptions)
	defer freeA()
	raw := C.GDALGroupGetMDArrayFullNamesRecursive(group.cValue, C.CSLConstList(unsafe.Pointer(gOpts)), C.CSLConstList(unsafe.Pointer(aOpts)))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (g GDALGroup) GetMDArrayFullNamesRecursive(groupOptions, arrayOptions []string) (result []string) {
	result = gdalGroupGetMDArrayFullNamesRecursive(g, groupOptions, arrayOptions)
	return
}

func gdalGroupOpenMDArray(group GDALGroup, name string, options []string) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = GDALMDArray{cValue: C.GDALGroupOpenMDArray(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) OpenMDArray(name string, options []string) (result GDALMDArray, err error) {
	result = gdalGroupOpenMDArray(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupOpenMDArrayFromFullname(group GDALGroup, name string, options []string) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = GDALMDArray{cValue: C.GDALGroupOpenMDArrayFromFullname(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) OpenMDArrayFromFullname(name string, options []string) (result GDALMDArray, err error) {
	result = gdalGroupOpenMDArrayFromFullname(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupResolveMDArray(group GDALGroup, name, startingPoint string, options []string) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cStart := C.CString(startingPoint)
	defer C.free(unsafe.Pointer(cStart))
	opts, free := cStrings(options)
	defer free()
	result = GDALMDArray{cValue: C.GDALGroupResolveMDArray(group.cValue, cName, cStart, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) ResolveMDArray(name, startingPoint string, options []string) (result GDALMDArray, err error) {
	result = gdalGroupResolveMDArray(g, name, startingPoint, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupGetGroupNames(group GDALGroup, options []string) (result []string) {
	opts, free := cStrings(options)
	defer free()
	raw := C.GDALGroupGetGroupNames(group.cValue, C.CSLConstList(unsafe.Pointer(opts)))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (g GDALGroup) GetGroupNames(options []string) (result []string) {
	result = gdalGroupGetGroupNames(g, options)
	return
}

func gdalGroupOpenGroup(group GDALGroup, name string, options []string) (result GDALGroup) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = GDALGroup{cValue: C.GDALGroupOpenGroup(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) OpenGroup(name string, options []string) (result GDALGroup, err error) {
	result = gdalGroupOpenGroup(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupOpenGroupFromFullname(group GDALGroup, name string, options []string) (result GDALGroup) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = GDALGroup{cValue: C.GDALGroupOpenGroupFromFullname(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) OpenGroupFromFullname(name string, options []string) (result GDALGroup, err error) {
	result = gdalGroupOpenGroupFromFullname(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupGetVectorLayerNames(group GDALGroup, options []string) (result []string) {
	opts, free := cStrings(options)
	defer free()
	raw := C.GDALGroupGetVectorLayerNames(group.cValue, C.CSLConstList(unsafe.Pointer(opts)))
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (g GDALGroup) GetVectorLayerNames(options []string) (result []string) {
	result = gdalGroupGetVectorLayerNames(g, options)
	return
}

func gdalGroupOpenVectorLayer(group GDALGroup, name string, options []string) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = OGRLayer{cValue: C.GDALGroupOpenVectorLayer(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) OpenVectorLayer(name string, options []string) (result OGRLayer, err error) {
	result = gdalGroupOpenVectorLayer(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupGetDimensions(group GDALGroup, options []string) (result []GDALDimension) {
	opts, free := cStrings(options)
	defer free()
	var count C.size_t
	arr := C.GDALGroupGetDimensions(group.cValue, &count, C.CSLConstList(unsafe.Pointer(opts)))
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALDimension, int(count))
	for i := range result {
		result[i] = GDALDimension{cValue: src[i]}
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (g GDALGroup) GetDimensions(options []string) (result []GDALDimension) {
	result = gdalGroupGetDimensions(g, options)
	return
}

func gdalGroupGetAttribute(group GDALGroup, name string) (result GDALAttribute) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALAttribute{cValue: C.GDALGroupGetAttribute(group.cValue, cName)}
	return
}

func (g GDALGroup) GetAttribute(name string) (result GDALAttribute, err error) {
	result = gdalGroupGetAttribute(g, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupGetAttributes(group GDALGroup, options []string) (result []GDALAttribute) {
	opts, free := cStrings(options)
	defer free()
	var count C.size_t
	arr := C.GDALGroupGetAttributes(group.cValue, &count, C.CSLConstList(unsafe.Pointer(opts)))
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALAttribute, int(count))
	for i := range result {
		result[i] = GDALAttribute{cValue: src[i]}
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (g GDALGroup) GetAttributes(options []string) (result []GDALAttribute) {
	result = gdalGroupGetAttributes(g, options)
	return
}

func gdalGroupGetStructuralInfo(group GDALGroup) (result []string) {
	result = goStrings(C.GDALGroupGetStructuralInfo(group.cValue))
	return
}

func (g GDALGroup) GetStructuralInfo() (result []string) {
	result = gdalGroupGetStructuralInfo(g)
	return
}

func gdalGroupCreateGroup(group GDALGroup, name string, options []string) (result GDALGroup) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = GDALGroup{cValue: C.GDALGroupCreateGroup(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) CreateGroup(name string, options []string) (result GDALGroup, err error) {
	result = gdalGroupCreateGroup(g, name, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupDeleteGroup(group GDALGroup, name string, options []string) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = bool(C.GDALGroupDeleteGroup(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (g GDALGroup) DeleteGroup(name string, options []string) (result bool) {
	result = gdalGroupDeleteGroup(g, name, options)
	return
}

func gdalGroupCreateDimension(group GDALGroup, name, dimType, direction string, size uint64, options []string) (result GDALDimension) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cType := C.CString(dimType)
	defer C.free(unsafe.Pointer(cType))
	cDir := C.CString(direction)
	defer C.free(unsafe.Pointer(cDir))
	opts, free := cStrings(options)
	defer free()
	result = GDALDimension{cValue: C.GDALGroupCreateDimension(group.cValue, cName, cType, cDir, C.GUInt64(size), C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) CreateDimension(name, dimType, direction string, size uint64, options []string) (result GDALDimension, err error) {
	result = gdalGroupCreateDimension(g, name, dimType, direction, size, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupCreateMDArray(group GDALGroup, name string, nDimensions int, dimensions []GDALDimension, dataType GDALExtendedDataType, options []string) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	var dimsPtr *C.GDALDimensionH
	if len(dimensions) > 0 {
		dimsPtr = (*C.GDALDimensionH)(unsafe.Pointer(&dimensions[0]))
	}
	result = GDALMDArray{cValue: C.GDALGroupCreateMDArray(group.cValue, cName, C.size_t(nDimensions), dimsPtr, dataType.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) CreateMDArray(name string, dimensions []GDALDimension, dataType GDALExtendedDataType, options []string) (result GDALMDArray, err error) {
	result = gdalGroupCreateMDArray(g, name, len(dimensions), dimensions, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupDeleteMDArray(group GDALGroup, name string, options []string) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = bool(C.GDALGroupDeleteMDArray(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (g GDALGroup) DeleteMDArray(name string, options []string) (result bool) {
	result = gdalGroupDeleteMDArray(g, name, options)
	return
}

func gdalGroupCreateAttribute(group GDALGroup, name string, nDimensions int, dimensions []uint64, dataType GDALExtendedDataType, options []string) (result GDALAttribute) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	cDims := make([]C.GUInt64, len(dimensions))
	for i, v := range dimensions {
		cDims[i] = C.GUInt64(v)
	}
	var dimsPtr *C.GUInt64
	if len(dimensions) > 0 {
		dimsPtr = &cDims[0]
	}
	result = GDALAttribute{cValue: C.GDALGroupCreateAttribute(group.cValue, cName, C.size_t(nDimensions), dimsPtr, dataType.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) CreateAttribute(name string, dimensions []uint64, dataType GDALExtendedDataType, options []string) (result GDALAttribute, err error) {
	result = gdalGroupCreateAttribute(g, name, len(dimensions), dimensions, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupDeleteAttribute(group GDALGroup, name string, options []string) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = bool(C.GDALGroupDeleteAttribute(group.cValue, cName, C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (g GDALGroup) DeleteAttribute(name string, options []string) (result bool) {
	result = gdalGroupDeleteAttribute(g, name, options)
	return
}

func gdalGroupRename(group GDALGroup, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALGroupRename(group.cValue, cName))
	return
}

func (g GDALGroup) Rename(newName string) (result bool) {
	result = gdalGroupRename(g, newName)
	return
}

func gdalGroupSubsetDimensionFromSelection(group GDALGroup, selection string, options []string) (result GDALGroup) {
	cSelection := C.CString(selection)
	defer C.free(unsafe.Pointer(cSelection))
	opts, free := cStrings(options)
	defer free()
	result = GDALGroup{cValue: C.GDALGroupSubsetDimensionFromSelection(group.cValue, cSelection, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (g GDALGroup) SubsetDimensionFromSelection(selection string, options []string) (result GDALGroup, err error) {
	result = gdalGroupSubsetDimensionFromSelection(g, selection, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalGroupGetDataTypeCount(group GDALGroup) (result int) {
	result = int(C.GDALGroupGetDataTypeCount(group.cValue))
	return
}

func (g GDALGroup) GetDataTypeCount() (result int) {
	result = gdalGroupGetDataTypeCount(g)
	return
}

func gdalGroupGetDataType(group GDALGroup, index int) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALGroupGetDataType(group.cValue, C.size_t(index))}
	return
}

func (g GDALGroup) GetDataType(index int) (result GDALExtendedDataType, err error) {
	result = gdalGroupGetDataType(g, index)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayRelease(array GDALMDArray) {
	C.GDALMDArrayRelease(array.cValue)
}

func (a GDALMDArray) Release() {
	gdalMDArrayRelease(a)
}

func gdalMDArrayGetName(array GDALMDArray) (result string) {
	result = C.GoString(C.GDALMDArrayGetName(array.cValue))
	return
}

func (a GDALMDArray) GetName() (result string) {
	result = gdalMDArrayGetName(a)
	return
}

func gdalMDArrayGetFullName(array GDALMDArray) (result string) {
	result = C.GoString(C.GDALMDArrayGetFullName(array.cValue))
	return
}

func (a GDALMDArray) GetFullName() (result string) {
	result = gdalMDArrayGetFullName(a)
	return
}

func gdalMDArrayGetTotalElementsCount(array GDALMDArray) (result uint64) {
	result = uint64(C.GDALMDArrayGetTotalElementsCount(array.cValue))
	return
}

func (a GDALMDArray) GetTotalElementsCount() (result uint64) {
	result = gdalMDArrayGetTotalElementsCount(a)
	return
}

func gdalMDArrayGetDimensionCount(array GDALMDArray) (result int) {
	result = int(C.GDALMDArrayGetDimensionCount(array.cValue))
	return
}

func (a GDALMDArray) GetDimensionCount() (result int) {
	result = gdalMDArrayGetDimensionCount(a)
	return
}

func gdalMDArrayGetDimensions(array GDALMDArray) (result []GDALDimension) {
	var count C.size_t
	arr := C.GDALMDArrayGetDimensions(array.cValue, &count)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALDimension, int(count))
	for i := range result {
		result[i] = GDALDimension{cValue: src[i]}
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (a GDALMDArray) GetDimensions() (result []GDALDimension) {
	result = gdalMDArrayGetDimensions(a)
	return
}

func gdalMDArrayGetDataType(array GDALMDArray) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALMDArrayGetDataType(array.cValue)}
	return
}

func (a GDALMDArray) GetDataType() (result GDALExtendedDataType, err error) {
	result = gdalMDArrayGetDataType(a)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// GDALMDArrayRead, GDALMDArrayWrite, GDALMDArrayAdviseRead and
// GDALMDArrayAdviseReadEx are deferred: their typed index/step/stride arrays
// plus void* buffer need a dedicated design.

func gdalMDArrayGetAttribute(array GDALMDArray, name string) (result GDALAttribute) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALAttribute{cValue: C.GDALMDArrayGetAttribute(array.cValue, cName)}
	return
}

func (a GDALMDArray) GetAttribute(name string) (result GDALAttribute, err error) {
	result = gdalMDArrayGetAttribute(a, name)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayGetAttributes(array GDALMDArray, options []string) (result []GDALAttribute) {
	opts, free := cStrings(options)
	defer free()
	var count C.size_t
	arr := C.GDALMDArrayGetAttributes(array.cValue, &count, C.CSLConstList(unsafe.Pointer(opts)))
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALAttribute, int(count))
	for i := range result {
		result[i] = GDALAttribute{cValue: src[i]}
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (a GDALMDArray) GetAttributes(options []string) (result []GDALAttribute) {
	result = gdalMDArrayGetAttributes(a, options)
	return
}

func gdalMDArrayCreateAttribute(array GDALMDArray, name string, nDimensions int, dimensions []uint64, dataType GDALExtendedDataType, options []string) (result GDALAttribute) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	cDims := make([]C.GUInt64, len(dimensions))
	for i, v := range dimensions {
		cDims[i] = C.GUInt64(v)
	}
	var dimsPtr *C.GUInt64
	if len(dimensions) > 0 {
		dimsPtr = &cDims[0]
	}
	result = GDALAttribute{cValue: C.GDALMDArrayCreateAttribute(array.cValue, cName, C.size_t(nDimensions), dimsPtr, dataType.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (a GDALMDArray) CreateAttribute(name string, dimensions []uint64, dataType GDALExtendedDataType, options []string) (result GDALAttribute, err error) {
	result = gdalMDArrayCreateAttribute(a, name, len(dimensions), dimensions, dataType, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayDeleteAttribute(array GDALMDArray, name string, options []string) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts, free := cStrings(options)
	defer free()
	result = bool(C.GDALMDArrayDeleteAttribute(array.cValue, cName, C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (a GDALMDArray) DeleteAttribute(name string, options []string) (result bool) {
	result = gdalMDArrayDeleteAttribute(a, name, options)
	return
}

func gdalMDArrayResize(array GDALMDArray, newDimSizes []uint64, options []string) (result bool) {
	opts, free := cStrings(options)
	defer free()
	cSizes := make([]C.GUInt64, len(newDimSizes))
	for i, v := range newDimSizes {
		cSizes[i] = C.GUInt64(v)
	}
	var sizesPtr *C.GUInt64
	if len(newDimSizes) > 0 {
		sizesPtr = &cSizes[0]
	}
	result = bool(C.GDALMDArrayResize(array.cValue, sizesPtr, C.CSLConstList(unsafe.Pointer(opts))))
	return
}

func (a GDALMDArray) Resize(newDimSizes []uint64, options []string) (result bool) {
	result = gdalMDArrayResize(a, newDimSizes, options)
	return
}

// GDALMDArrayGetRawNoDataValue and GDALMDArraySetRawNoDataValue are deferred:
// the raw void* nodata value needs a dedicated design.

func gdalMDArrayGetNoDataValueAsDouble(array GDALMDArray, hasNoData *int) (result float64) {
	var cHas C.int
	result = float64(C.GDALMDArrayGetNoDataValueAsDouble(array.cValue, &cHas))
	*hasNoData = int(cHas)
	return
}

func (a GDALMDArray) GetNoDataValueAsDouble() (value float64, ok bool) {
	var has int
	value = gdalMDArrayGetNoDataValueAsDouble(a, &has)
	ok = has != 0
	return
}

func gdalMDArrayGetNoDataValueAsInt64(array GDALMDArray, hasNoData *int) (result int64) {
	var cHas C.int
	result = int64(C.GDALMDArrayGetNoDataValueAsInt64(array.cValue, &cHas))
	*hasNoData = int(cHas)
	return
}

func (a GDALMDArray) GetNoDataValueAsInt64() (value int64, ok bool) {
	var has int
	value = gdalMDArrayGetNoDataValueAsInt64(a, &has)
	ok = has != 0
	return
}

func gdalMDArrayGetNoDataValueAsUInt64(array GDALMDArray, hasNoData *int) (result uint64) {
	var cHas C.int
	result = uint64(C.GDALMDArrayGetNoDataValueAsUInt64(array.cValue, &cHas))
	*hasNoData = int(cHas)
	return
}

func (a GDALMDArray) GetNoDataValueAsUInt64() (value uint64, ok bool) {
	var has int
	value = gdalMDArrayGetNoDataValueAsUInt64(a, &has)
	ok = has != 0
	return
}

func gdalMDArraySetNoDataValueAsDouble(array GDALMDArray, value float64) (result bool) {
	result = C.GDALMDArraySetNoDataValueAsDouble(array.cValue, C.double(value)) != 0
	return
}

func (a GDALMDArray) SetNoDataValueAsDouble(value float64) (result bool) {
	result = gdalMDArraySetNoDataValueAsDouble(a, value)
	return
}

func gdalMDArraySetNoDataValueAsInt64(array GDALMDArray, value int64) (result bool) {
	result = C.GDALMDArraySetNoDataValueAsInt64(array.cValue, C.int64_t(value)) != 0
	return
}

func (a GDALMDArray) SetNoDataValueAsInt64(value int64) (result bool) {
	result = gdalMDArraySetNoDataValueAsInt64(a, value)
	return
}

func gdalMDArraySetNoDataValueAsUInt64(array GDALMDArray, value uint64) (result bool) {
	result = C.GDALMDArraySetNoDataValueAsUInt64(array.cValue, C.uint64_t(value)) != 0
	return
}

func (a GDALMDArray) SetNoDataValueAsUInt64(value uint64) (result bool) {
	result = gdalMDArraySetNoDataValueAsUInt64(a, value)
	return
}

func gdalMDArraySetScale(array GDALMDArray, scale float64) (result bool) {
	result = C.GDALMDArraySetScale(array.cValue, C.double(scale)) != 0
	return
}

func (a GDALMDArray) SetScale(scale float64) (result bool) {
	result = gdalMDArraySetScale(a, scale)
	return
}

func gdalMDArraySetScaleEx(array GDALMDArray, scale float64, storageType GDALDataType) (result bool) {
	result = C.GDALMDArraySetScaleEx(array.cValue, C.double(scale), C.GDALDataType(storageType)) != 0
	return
}

func (a GDALMDArray) SetScaleEx(scale float64, storageType GDALDataType) (result bool) {
	result = gdalMDArraySetScaleEx(a, scale, storageType)
	return
}

func gdalMDArrayGetScale(array GDALMDArray, hasValue *int) (result float64) {
	var cHas C.int
	result = float64(C.GDALMDArrayGetScale(array.cValue, &cHas))
	*hasValue = int(cHas)
	return
}

func (a GDALMDArray) GetScale() (value float64, ok bool) {
	var has int
	value = gdalMDArrayGetScale(a, &has)
	ok = has != 0
	return
}

func gdalMDArraySetOffset(array GDALMDArray, offset float64) (result bool) {
	result = C.GDALMDArraySetOffset(array.cValue, C.double(offset)) != 0
	return
}

func (a GDALMDArray) SetOffset(offset float64) (result bool) {
	result = gdalMDArraySetOffset(a, offset)
	return
}

func gdalMDArraySetOffsetEx(array GDALMDArray, offset float64, storageType GDALDataType) (result bool) {
	result = C.GDALMDArraySetOffsetEx(array.cValue, C.double(offset), C.GDALDataType(storageType)) != 0
	return
}

func (a GDALMDArray) SetOffsetEx(offset float64, storageType GDALDataType) (result bool) {
	result = gdalMDArraySetOffsetEx(a, offset, storageType)
	return
}

func gdalMDArrayGetOffset(array GDALMDArray, hasValue *int) (result float64) {
	var cHas C.int
	result = float64(C.GDALMDArrayGetOffset(array.cValue, &cHas))
	*hasValue = int(cHas)
	return
}

func (a GDALMDArray) GetOffset() (value float64, ok bool) {
	var has int
	value = gdalMDArrayGetOffset(a, &has)
	ok = has != 0
	return
}

// GDALMDArrayGetScaleEx, GDALMDArrayGetOffsetEx (with storage-type out-param),
// GDALMDArrayGetBlockSize and GDALMDArrayGetProcessingChunkSize (GUInt64*/size_t*
// array returns) are deferred.

func gdalMDArraySetUnit(array GDALMDArray, unit string) (result bool) {
	cUnit := C.CString(unit)
	defer C.free(unsafe.Pointer(cUnit))
	result = C.GDALMDArraySetUnit(array.cValue, cUnit) != 0
	return
}

func (a GDALMDArray) SetUnit(unit string) (result bool) {
	result = gdalMDArraySetUnit(a, unit)
	return
}

func gdalMDArrayGetUnit(array GDALMDArray) (result string) {
	result = C.GoString(C.GDALMDArrayGetUnit(array.cValue))
	return
}

func (a GDALMDArray) GetUnit() (result string) {
	result = gdalMDArrayGetUnit(a)
	return
}

func gdalMDArraySetSpatialRef(array GDALMDArray, srs OGRSpatialReference) (result bool) {
	result = C.GDALMDArraySetSpatialRef(array.cValue, srs.cValue) != 0
	return
}

func (a GDALMDArray) SetSpatialRef(srs OGRSpatialReference) (result bool) {
	result = gdalMDArraySetSpatialRef(a, srs)
	return
}

func gdalMDArrayGetSpatialRef(array GDALMDArray) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.GDALMDArrayGetSpatialRef(array.cValue)}
	return
}

func (a GDALMDArray) GetSpatialRef() (result OGRSpatialReference) {
	result = gdalMDArrayGetSpatialRef(a)
	return
}

func gdalMDArrayGetStructuralInfo(array GDALMDArray) (result []string) {
	result = goStrings(C.GDALMDArrayGetStructuralInfo(array.cValue))
	return
}

func (a GDALMDArray) GetStructuralInfo() (result []string) {
	result = gdalMDArrayGetStructuralInfo(a)
	return
}

func gdalMDArrayGetView(array GDALMDArray, viewExpr string) (result GDALMDArray) {
	cExpr := C.CString(viewExpr)
	defer C.free(unsafe.Pointer(cExpr))
	result = GDALMDArray{cValue: C.GDALMDArrayGetView(array.cValue, cExpr)}
	return
}

func (a GDALMDArray) GetView(viewExpr string) (result GDALMDArray, err error) {
	result = gdalMDArrayGetView(a, viewExpr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayTranspose(array GDALMDArray, newAxisCount int, mapNewAxisToOldAxis []int) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALMDArrayTranspose(array.cValue, C.size_t(newAxisCount), cInts(mapNewAxisToOldAxis))}
	return
}

func (a GDALMDArray) Transpose(mapNewAxisToOldAxis []int) (result GDALMDArray, err error) {
	result = gdalMDArrayTranspose(a, len(mapNewAxisToOldAxis), mapNewAxisToOldAxis)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayGetUnscaled(array GDALMDArray) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALMDArrayGetUnscaled(array.cValue)}
	return
}

func (a GDALMDArray) GetUnscaled() (result GDALMDArray, err error) {
	result = gdalMDArrayGetUnscaled(a)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayGetMask(array GDALMDArray, options []string) (result GDALMDArray) {
	opts, free := cStrings(options)
	defer free()
	result = GDALMDArray{cValue: C.GDALMDArrayGetMask(array.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (a GDALMDArray) GetMask(options []string) (result GDALMDArray, err error) {
	result = gdalMDArrayGetMask(a, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayAsClassicDataset(array GDALMDArray, xDim, yDim int) (result GDALDataset) {
	result = GDALDataset{cValue: C.GDALMDArrayAsClassicDataset(array.cValue, C.size_t(xDim), C.size_t(yDim))}
	return
}

func (a GDALMDArray) AsClassicDataset(xDim, yDim int) (result GDALDataset, err error) {
	result = gdalMDArrayAsClassicDataset(a, xDim, yDim)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayAsClassicDatasetEx(array GDALMDArray, xDim, yDim int, rootGroup GDALGroup, options []string) (result GDALDataset) {
	opts, free := cStrings(options)
	defer free()
	result = GDALDataset{cValue: C.GDALMDArrayAsClassicDatasetEx(array.cValue, C.size_t(xDim), C.size_t(yDim), rootGroup.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (a GDALMDArray) AsClassicDatasetEx(xDim, yDim int, rootGroup GDALGroup, options []string) (result GDALDataset, err error) {
	result = gdalMDArrayAsClassicDatasetEx(a, xDim, yDim, rootGroup, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayGetStatistics(array GDALMDArray, dataset GDALDataset, approxOK, force int, min, max, mean, stdDev *float64, validCount *uint64, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	var cMin, cMax, cMean, cStdDev C.double
	var cValidCount C.GUInt64
	result = CPLErr(C.GDALMDArrayGetStatistics(array.cValue, dataset.cValue, C.int(approxOK), C.int(force), &cMin, &cMax, &cMean, &cStdDev, &cValidCount, progress.cValue, progressData))
	*min = float64(cMin)
	*max = float64(cMax)
	*mean = float64(cMean)
	*stdDev = float64(cStdDev)
	*validCount = uint64(cValidCount)
	return
}

func (a GDALMDArray) GetStatistics(dataset GDALDataset, approxOK, force int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max, mean, stdDev float64, validCount uint64, err error) {
	err = cplErr(gdalMDArrayGetStatistics(a, dataset, approxOK, force, &min, &max, &mean, &stdDev, &validCount, progress, progressData))
	return
}

func gdalMDArrayComputeStatistics(array GDALMDArray, dataset GDALDataset, approxOK int, min, max, mean, stdDev *float64, validCount *uint64, progress GDALProgressFunc, progressData unsafe.Pointer) (result bool) {
	var cMin, cMax, cMean, cStdDev C.double
	var cValidCount C.GUInt64
	result = C.GDALMDArrayComputeStatistics(array.cValue, dataset.cValue, C.int(approxOK), &cMin, &cMax, &cMean, &cStdDev, &cValidCount, progress.cValue, progressData) != 0
	*min = float64(cMin)
	*max = float64(cMax)
	*mean = float64(cMean)
	*stdDev = float64(cStdDev)
	*validCount = uint64(cValidCount)
	return
}

func (a GDALMDArray) ComputeStatistics(dataset GDALDataset, approxOK int, progress GDALProgressFunc, progressData unsafe.Pointer) (min, max, mean, stdDev float64, validCount uint64, ok bool) {
	ok = gdalMDArrayComputeStatistics(a, dataset, approxOK, &min, &max, &mean, &stdDev, &validCount, progress, progressData)
	return
}

func gdalMDArrayGetResampled(array GDALMDArray, newDimCount int, newDims []GDALDimension, resampleAlg GDALRIOResampleAlg, targetSRS OGRSpatialReference, options []string) (result GDALMDArray) {
	opts, free := cStrings(options)
	defer free()
	var dimsPtr *C.GDALDimensionH
	if len(newDims) > 0 {
		dimsPtr = (*C.GDALDimensionH)(unsafe.Pointer(&newDims[0]))
	}
	result = GDALMDArray{cValue: C.GDALMDArrayGetResampled(array.cValue, C.size_t(newDimCount), dimsPtr, C.GDALRIOResampleAlg(resampleAlg), targetSRS.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (a GDALMDArray) GetResampled(newDims []GDALDimension, resampleAlg GDALRIOResampleAlg, targetSRS OGRSpatialReference, options []string) (result GDALMDArray, err error) {
	result = gdalMDArrayGetResampled(a, len(newDims), newDims, resampleAlg, targetSRS, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayGetGridded(array GDALMDArray, gridOptions string, xArray, yArray GDALMDArray, options []string) (result GDALMDArray) {
	cGridOptions := C.CString(gridOptions)
	defer C.free(unsafe.Pointer(cGridOptions))
	opts, free := cStrings(options)
	defer free()
	result = GDALMDArray{cValue: C.GDALMDArrayGetGridded(array.cValue, cGridOptions, xArray.cValue, yArray.cValue, C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (a GDALMDArray) GetGridded(gridOptions string, xArray, yArray GDALMDArray, options []string) (result GDALMDArray, err error) {
	result = gdalMDArrayGetGridded(a, gridOptions, xArray, yArray, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalMDArrayGetCoordinateVariables(array GDALMDArray) (result []GDALMDArray) {
	var count C.size_t
	arr := C.GDALMDArrayGetCoordinateVariables(array.cValue, &count)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALMDArray, int(count))
	for i := range result {
		result[i] = GDALMDArray{cValue: src[i]}
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (a GDALMDArray) GetCoordinateVariables() (result []GDALMDArray) {
	result = gdalMDArrayGetCoordinateVariables(a)
	return
}

// GDALMDArrayGetMeshGrid (array-in/array-out), the GDALMDArrayRawBlockInfo
// struct + GDALMDArrayRawBlockInfoCreate/Release/GDALMDArrayGetRawBlockInfo, and
// GDALReleaseArrays are deferred.

func gdalMDArrayCache(array GDALMDArray, options []string) (result bool) {
	opts, free := cStrings(options)
	defer free()
	result = C.GDALMDArrayCache(array.cValue, C.CSLConstList(unsafe.Pointer(opts))) != 0
	return
}

func (a GDALMDArray) Cache(options []string) (result bool) {
	result = gdalMDArrayCache(a, options)
	return
}

func gdalMDArrayRename(array GDALMDArray, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALMDArrayRename(array.cValue, cName))
	return
}

func (a GDALMDArray) Rename(newName string) (result bool) {
	result = gdalMDArrayRename(a, newName)
	return
}

func gdalCreateRasterAttributeTableFromMDArrays(tableType GDALRATTableType, nArrays int, arrays []GDALMDArray, usages []GDALRATFieldUsage) (result GDALRasterAttributeTable) {
	var arraysPtr *C.GDALMDArrayH
	if len(arrays) > 0 {
		arraysPtr = (*C.GDALMDArrayH)(unsafe.Pointer(&arrays[0]))
	}
	var usagesPtr *C.GDALRATFieldUsage
	if len(usages) > 0 {
		usagesPtr = (*C.GDALRATFieldUsage)(unsafe.Pointer(&usages[0]))
	}
	result = GDALRasterAttributeTable{cValue: C.GDALCreateRasterAttributeTableFromMDArrays(C.GDALRATTableType(tableType), C.int(nArrays), arraysPtr, usagesPtr)}
	return
}

func GDALCreateRasterAttributeTableFromMDArrays(tableType GDALRATTableType, arrays []GDALMDArray, usages []GDALRATFieldUsage) (result GDALRasterAttributeTable, err error) {
	result = gdalCreateRasterAttributeTableFromMDArrays(tableType, len(arrays), arrays, usages)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalAttributeRelease(attr GDALAttribute) {
	C.GDALAttributeRelease(attr.cValue)
}

func (attr GDALAttribute) Release() {
	gdalAttributeRelease(attr)
}

func gdalAttributeGetName(attr GDALAttribute) (result string) {
	result = C.GoString(C.GDALAttributeGetName(attr.cValue))
	return
}

func (attr GDALAttribute) GetName() (result string) {
	result = gdalAttributeGetName(attr)
	return
}

func gdalAttributeGetFullName(attr GDALAttribute) (result string) {
	result = C.GoString(C.GDALAttributeGetFullName(attr.cValue))
	return
}

func (attr GDALAttribute) GetFullName() (result string) {
	result = gdalAttributeGetFullName(attr)
	return
}

func gdalAttributeGetTotalElementsCount(attr GDALAttribute) (result uint64) {
	result = uint64(C.GDALAttributeGetTotalElementsCount(attr.cValue))
	return
}

func (attr GDALAttribute) GetTotalElementsCount() (result uint64) {
	result = gdalAttributeGetTotalElementsCount(attr)
	return
}

func gdalAttributeGetDimensionCount(attr GDALAttribute) (result int) {
	result = int(C.GDALAttributeGetDimensionCount(attr.cValue))
	return
}

func (attr GDALAttribute) GetDimensionCount() (result int) {
	result = gdalAttributeGetDimensionCount(attr)
	return
}

func gdalAttributeGetDimensionsSize(attr GDALAttribute) (result []uint64) {
	var count C.size_t
	arr := C.GDALAttributeGetDimensionsSize(attr.cValue, &count)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]uint64, int(count))
	for i := range result {
		result[i] = uint64(src[i])
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (attr GDALAttribute) GetDimensionsSize() (result []uint64) {
	result = gdalAttributeGetDimensionsSize(attr)
	return
}

func gdalAttributeGetDataType(attr GDALAttribute) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALAttributeGetDataType(attr.cValue)}
	return
}

func (attr GDALAttribute) GetDataType() (result GDALExtendedDataType, err error) {
	result = gdalAttributeGetDataType(attr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalAttributeReadAsRaw(attr GDALAttribute) (result []byte) {
	var size C.size_t
	raw := C.GDALAttributeReadAsRaw(attr.cValue, &size)
	if raw != nil {
		result = C.GoBytes(unsafe.Pointer(raw), C.int(size))
		C.GDALAttributeFreeRawResult(attr.cValue, raw, size)
	}
	return
}

func (attr GDALAttribute) ReadAsRaw() (result []byte) {
	result = gdalAttributeReadAsRaw(attr)
	return
}

func gdalAttributeReadAsString(attr GDALAttribute) (result string) {
	result = C.GoString(C.GDALAttributeReadAsString(attr.cValue))
	return
}

func (attr GDALAttribute) ReadAsString() (result string) {
	result = gdalAttributeReadAsString(attr)
	return
}

func gdalAttributeReadAsInt(attr GDALAttribute) (result int) {
	result = int(C.GDALAttributeReadAsInt(attr.cValue))
	return
}

func (attr GDALAttribute) ReadAsInt() (result int) {
	result = gdalAttributeReadAsInt(attr)
	return
}

func gdalAttributeReadAsInt64(attr GDALAttribute) (result int64) {
	result = int64(C.GDALAttributeReadAsInt64(attr.cValue))
	return
}

func (attr GDALAttribute) ReadAsInt64() (result int64) {
	result = gdalAttributeReadAsInt64(attr)
	return
}

func gdalAttributeReadAsDouble(attr GDALAttribute) (result float64) {
	result = float64(C.GDALAttributeReadAsDouble(attr.cValue))
	return
}

func (attr GDALAttribute) ReadAsDouble() (result float64) {
	result = gdalAttributeReadAsDouble(attr)
	return
}

func gdalAttributeReadAsStringArray(attr GDALAttribute) (result []string) {
	raw := C.GDALAttributeReadAsStringArray(attr.cValue)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (attr GDALAttribute) ReadAsStringArray() (result []string) {
	result = gdalAttributeReadAsStringArray(attr)
	return
}

func gdalAttributeReadAsIntArray(attr GDALAttribute) (result []int) {
	var count C.size_t
	arr := C.GDALAttributeReadAsIntArray(attr.cValue, &count)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]int, int(count))
	for i := range result {
		result[i] = int(src[i])
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (attr GDALAttribute) ReadAsIntArray() (result []int) {
	result = gdalAttributeReadAsIntArray(attr)
	return
}

func gdalAttributeReadAsInt64Array(attr GDALAttribute) (result []int64) {
	var count C.size_t
	arr := C.GDALAttributeReadAsInt64Array(attr.cValue, &count)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]int64, int(count))
	for i := range result {
		result[i] = int64(src[i])
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (attr GDALAttribute) ReadAsInt64Array() (result []int64) {
	result = gdalAttributeReadAsInt64Array(attr)
	return
}

func gdalAttributeReadAsDoubleArray(attr GDALAttribute) (result []float64) {
	var count C.size_t
	arr := C.GDALAttributeReadAsDoubleArray(attr.cValue, &count)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]float64, int(count))
	for i := range result {
		result[i] = float64(src[i])
	}
	C.VSIFree(unsafe.Pointer(arr))
	return
}

func (attr GDALAttribute) ReadAsDoubleArray() (result []float64) {
	result = gdalAttributeReadAsDoubleArray(attr)
	return
}

func gdalAttributeWriteRaw(attr GDALAttribute, data unsafe.Pointer, size int) (result bool) {
	result = C.GDALAttributeWriteRaw(attr.cValue, data, C.size_t(size)) != 0
	return
}

func (attr GDALAttribute) WriteRaw(data []byte) (result bool) {
	result = gdalAttributeWriteRaw(attr, cBytes(data), len(data))
	return
}

func gdalAttributeWriteString(attr GDALAttribute, value string) (result bool) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = C.GDALAttributeWriteString(attr.cValue, cValue) != 0
	return
}

func (attr GDALAttribute) WriteString(value string) (result bool) {
	result = gdalAttributeWriteString(attr, value)
	return
}

func gdalAttributeWriteStringArray(attr GDALAttribute, values []string) (result bool) {
	v, free := cStrings(values)
	defer free()
	result = C.GDALAttributeWriteStringArray(attr.cValue, C.CSLConstList(unsafe.Pointer(v))) != 0
	return
}

func (attr GDALAttribute) WriteStringArray(values []string) (result bool) {
	result = gdalAttributeWriteStringArray(attr, values)
	return
}

func gdalAttributeWriteInt(attr GDALAttribute, value int) (result bool) {
	result = C.GDALAttributeWriteInt(attr.cValue, C.int(value)) != 0
	return
}

func (attr GDALAttribute) WriteInt(value int) (result bool) {
	result = gdalAttributeWriteInt(attr, value)
	return
}

func gdalAttributeWriteIntArray(attr GDALAttribute, values []int, count int) (result bool) {
	result = C.GDALAttributeWriteIntArray(attr.cValue, cInts(values), C.size_t(count)) != 0
	return
}

func (attr GDALAttribute) WriteIntArray(values []int) (result bool) {
	result = gdalAttributeWriteIntArray(attr, values, len(values))
	return
}

func gdalAttributeWriteInt64(attr GDALAttribute, value int64) (result bool) {
	result = C.GDALAttributeWriteInt64(attr.cValue, C.int64_t(value)) != 0
	return
}

func (attr GDALAttribute) WriteInt64(value int64) (result bool) {
	result = gdalAttributeWriteInt64(attr, value)
	return
}

func gdalAttributeWriteInt64Array(attr GDALAttribute, values []int64, count int) (result bool) {
	cValues := make([]C.int64_t, len(values))
	for i, v := range values {
		cValues[i] = C.int64_t(v)
	}
	var ptr *C.int64_t
	if len(values) > 0 {
		ptr = &cValues[0]
	}
	result = C.GDALAttributeWriteInt64Array(attr.cValue, ptr, C.size_t(count)) != 0
	return
}

func (attr GDALAttribute) WriteInt64Array(values []int64) (result bool) {
	result = gdalAttributeWriteInt64Array(attr, values, len(values))
	return
}

func gdalAttributeWriteDouble(attr GDALAttribute, value float64) (result bool) {
	result = C.GDALAttributeWriteDouble(attr.cValue, C.double(value)) != 0
	return
}

func (attr GDALAttribute) WriteDouble(value float64) (result bool) {
	result = gdalAttributeWriteDouble(attr, value)
	return
}

func gdalAttributeWriteDoubleArray(attr GDALAttribute, values []float64, count int) (result bool) {
	cValues := make([]C.double, len(values))
	for i, v := range values {
		cValues[i] = C.double(v)
	}
	var ptr *C.double
	if len(values) > 0 {
		ptr = &cValues[0]
	}
	result = C.GDALAttributeWriteDoubleArray(attr.cValue, ptr, C.size_t(count)) != 0
	return
}

func (attr GDALAttribute) WriteDoubleArray(values []float64) (result bool) {
	result = gdalAttributeWriteDoubleArray(attr, values, len(values))
	return
}

func gdalAttributeRename(attr GDALAttribute, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALAttributeRename(attr.cValue, cName))
	return
}

func (attr GDALAttribute) Rename(newName string) (result bool) {
	result = gdalAttributeRename(attr, newName)
	return
}

func gdalDimensionRelease(dim GDALDimension) {
	C.GDALDimensionRelease(dim.cValue)
}

func (dim GDALDimension) Release() {
	gdalDimensionRelease(dim)
}

func gdalDimensionGetName(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetName(dim.cValue))
	return
}

func (dim GDALDimension) GetName() (result string) {
	result = gdalDimensionGetName(dim)
	return
}

func gdalDimensionGetFullName(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetFullName(dim.cValue))
	return
}

func (dim GDALDimension) GetFullName() (result string) {
	result = gdalDimensionGetFullName(dim)
	return
}

func gdalDimensionGetType(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetType(dim.cValue))
	return
}

func (dim GDALDimension) GetType() (result string) {
	result = gdalDimensionGetType(dim)
	return
}

func gdalDimensionGetDirection(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetDirection(dim.cValue))
	return
}

func (dim GDALDimension) GetDirection() (result string) {
	result = gdalDimensionGetDirection(dim)
	return
}

func gdalDimensionGetSize(dim GDALDimension) (result uint64) {
	result = uint64(C.GDALDimensionGetSize(dim.cValue))
	return
}

func (dim GDALDimension) GetSize() (result uint64) {
	result = gdalDimensionGetSize(dim)
	return
}

func gdalDimensionGetIndexingVariable(dim GDALDimension) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALDimensionGetIndexingVariable(dim.cValue)}
	return
}

func (dim GDALDimension) GetIndexingVariable() (result GDALMDArray, err error) {
	result = gdalDimensionGetIndexingVariable(dim)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func gdalDimensionSetIndexingVariable(dim GDALDimension, array GDALMDArray) (result bool) {
	result = C.GDALDimensionSetIndexingVariable(dim.cValue, array.cValue) != 0
	return
}

func (dim GDALDimension) SetIndexingVariable(array GDALMDArray) (result bool) {
	result = gdalDimensionSetIndexingVariable(dim, array)
	return
}

func gdalDimensionRename(dim GDALDimension, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALDimensionRename(dim.cValue, cName))
	return
}

func (dim GDALDimension) Rename(newName string) (result bool) {
	result = gdalDimensionRename(dim, newName)
	return
}

// CPL_C_END
