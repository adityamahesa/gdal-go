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

func gdalRegisterPlugin(name string) (err error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err = cplErr(CPLErr(C.GDALRegisterPlugin(cName)))
	return
}

func GDALRegisterPlugin(name string) (err error) {
	err = gdalRegisterPlugin(name)
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

func GDALIdentifyDriver(filename string, fileList []string) (result GDALDriver) {
	result = gdalIdentifyDriver(filename, fileList)
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

func GDALIdentifyDriverEx(filename string, identifyFlags uint, allowedDrivers, fileList []string) (result GDALDriver) {
	result = gdalIdentifyDriverEx(filename, identifyFlags, allowedDrivers, fileList)
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

func gdalDumpOpenDatasets(file *C.FILE) (result int) {
	result = int(C.GDALDumpOpenDatasets(file))
	return
}

func GDALDumpOpenDatasets(filename string) (result int, err error) {
	fp, closeFn, err := cFile(filename, "w")
	if err != nil {
		return
	}
	defer closeFn()
	result = gdalDumpOpenDatasets(fp)
	return
}

func gdalGetDriverByName(name string) (result GDALDriver) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDriver{cValue: C.GDALGetDriverByName(cName)}
	return
}

func GDALGetDriverByName(name string) (result GDALDriver) {
	result = gdalGetDriverByName(name)
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

func GDALGetDriver(index int) (result GDALDriver) {
	result = gdalGetDriver(index)
	return
}

func gdalCreateDriver() (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALCreateDriver()}
	return
}

func GDALCreateDriver() (result GDALDriver) {
	result = gdalCreateDriver()
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

func gdalDeleteDataset(driver GDALDriver, filename string) (err error) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	err = cplErr(CPLErr(C.GDALDeleteDataset(driver.cValue, cName)))
	return
}

func (d GDALDriver) DeleteDataset(filename string) (err error) {
	err = gdalDeleteDataset(d, filename)
	return
}

func gdalRenameDataset(driver GDALDriver, newName, oldName string) (err error) {
	cNew := C.CString(newName)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOld))
	err = cplErr(CPLErr(C.GDALRenameDataset(driver.cValue, cNew, cOld)))
	return
}

func (d GDALDriver) RenameDataset(newName, oldName string) (err error) {
	err = gdalRenameDataset(d, newName, oldName)
	return
}

func gdalCopyDatasetFiles(driver GDALDriver, newName, oldName string) (err error) {
	cNew := C.CString(newName)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOld))
	err = cplErr(CPLErr(C.GDALCopyDatasetFiles(driver.cValue, cNew, cOld)))
	return
}

func (d GDALDriver) CopyDatasetFiles(newName, oldName string) (err error) {
	err = gdalCopyDatasetFiles(d, newName, oldName)
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

func gdalSetMetadata(object GDALMajorObject, metadata []string, domain string) (err error) {
	md, free := cStrings(metadata)
	defer free()
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	err = cplErr(CPLErr(C.GDALSetMetadata(object.cValue, C.CSLConstList(unsafe.Pointer(md)), cDomain)))
	return
}

func (o GDALMajorObject) SetMetadata(metadata []string, domain string) (err error) {
	err = gdalSetMetadata(o, metadata, domain)
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

func gdalSetMetadataItem(object GDALMajorObject, name, value, domain string) (err error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	err = cplErr(CPLErr(C.GDALSetMetadataItem(object.cValue, cName, cValue, cDomain)))
	return
}

func (o GDALMajorObject) SetMetadataItem(name, value, domain string) (err error) {
	err = gdalSetMetadataItem(o, name, value, domain)
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

func gdalClose(dataset GDALDataset) (err error) {
	err = cplErr(CPLErr(C.GDALClose(dataset.cValue)))
	return
}

func (ds GDALDataset) Close() (err error) {
	err = gdalClose(ds)
	return
}

func gdalDatasetRunCloseWithoutDestroying(dataset GDALDataset) (err error) {
	err = cplErr(CPLErr(C.GDALDatasetRunCloseWithoutDestroying(dataset.cValue)))
	return
}

func (ds GDALDataset) RunCloseWithoutDestroying() (err error) {
	err = gdalDatasetRunCloseWithoutDestroying(ds)
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

func (ds GDALDataset) GetRasterBand(band int) (result GDALRasterBand) {
	result = gdalGetRasterBand(ds, band)
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

func gdalAddBand(dataset GDALDataset, dataType GDALDataType, options []string) (err error) {
	opts, free := cStrings(options)
	defer free()
	err = cplErr(CPLErr(C.GDALAddBand(dataset.cValue, C.GDALDataType(dataType), C.CSLConstList(unsafe.Pointer(opts)))))
	return
}

func (ds GDALDataset) AddBand(dataType GDALDataType, options []string) (err error) {
	err = gdalAddBand(ds, dataType, options)
	return
}

func gdalBeginAsyncReader(dataset GDALDataset, xOff, yOff, xSize, ySize int, buf unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandMap *C.int, pixelSpace, lineSpace, bandSpace int, options []string) (result GDALAsyncReader) {
	opts, free := cStrings(options)
	defer free()
	result = GDALAsyncReader{cValue: C.GDALBeginAsyncReader(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buf, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), bandMap, C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace), C.CSLConstList(unsafe.Pointer(opts)))}
	return
}

func (ds GDALDataset) BeginAsyncReader(xOff, yOff, xSize, ySize int, buf []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandMap []int, pixelSpace, lineSpace, bandSpace int, options []string) (result GDALAsyncReader, err error) {
	result = gdalBeginAsyncReader(ds, xOff, yOff, xSize, ySize, cBytes(buf), bufXSize, bufYSize, bufType, bandCount, cInts(bandMap), pixelSpace, lineSpace, bandSpace, options)
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

func gdalDatasetRasterIO(dataset GDALDataset, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList *C.int, pixelSpace, lineSpace, bandSpace int) (err error) {
	err = cplErr(CPLErr(C.GDALDatasetRasterIO(dataset.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), bandList, C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace))))
	return
}

func (ds GDALDataset) RasterIO(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int) (err error) {
	err = gdalDatasetRasterIO(ds, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, bandCount, cInts(bandList), pixelSpace, lineSpace, bandSpace)
	return
}

func gdalDatasetRasterIOEx(dataset GDALDataset, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList *C.int, pixelSpace, lineSpace, bandSpace int64, extraArg GDALRasterIOExtraArg) (err error) {
	err = cplErr(CPLErr(C.GDALDatasetRasterIOEx(dataset.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), bandList, C.GSpacing(pixelSpace), C.GSpacing(lineSpace), C.GSpacing(bandSpace), extraArg.cValue)))
	return
}

func (ds GDALDataset) RasterIOEx(rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer []byte, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int64, extraArg GDALRasterIOExtraArg) (err error) {
	err = gdalDatasetRasterIOEx(ds, rwFlag, xOff, yOff, xSize, ySize, cBytes(buffer), bufXSize, bufYSize, bufType, bandCount, cInts(bandList), pixelSpace, lineSpace, bandSpace, extraArg)
	return
}

func gdalDatasetAdviseRead(dataset GDALDataset, xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList *C.int, options []string) (err error) {
	opts, free := cStrings(options)
	defer free()
	err = cplErr(CPLErr(C.GDALDatasetAdviseRead(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), bandList, C.CSLConstList(unsafe.Pointer(opts)))))
	return
}

func (ds GDALDataset) AdviseRead(xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, options []string) (err error) {
	err = gdalDatasetAdviseRead(ds, xOff, yOff, xSize, ySize, bufXSize, bufYSize, bufType, bandCount, cInts(bandList), options)
	return
}

func gdalDatasetGetCompressionFormats(dataset GDALDataset, xOff, yOff, xSize, ySize, bandCount int, bandList *C.int) (result []string) {
	raw := C.GDALDatasetGetCompressionFormats(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bandCount), bandList)
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	result = goStrings(raw)
	return
}

func (ds GDALDataset) GetCompressionFormats(xOff, yOff, xSize, ySize, bandCount int, bandList []int) (result []string) {
	result = gdalDatasetGetCompressionFormats(ds, xOff, yOff, xSize, ySize, bandCount, cInts(bandList))
	return
}

func gdalDatasetReadCompressedData(dataset GDALDataset, format string, xOff, yOff, xSize, ySize, bandCount int, bandList *C.int, buffer *unsafe.Pointer, bufferSize *C.size_t, detailedFormat **C.char) (err error) {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	err = cplErr(CPLErr(C.GDALDatasetReadCompressedData(dataset.cValue, cFormat, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bandCount), bandList, buffer, bufferSize, detailedFormat)))
	return
}

func (ds GDALDataset) ReadCompressedData(format string, xOff, yOff, xSize, ySize, bandCount int, bandList []int) (buffer []byte, detailedFormat string, err error) {
	var cBuffer unsafe.Pointer
	var size C.size_t
	var cFormat *C.char
	err = gdalDatasetReadCompressedData(ds, format, xOff, yOff, xSize, ySize, bandCount, cInts(bandList), &cBuffer, &size, &cFormat)
	if err != nil {
		return
	}
	if cBuffer != nil {
		buffer = C.GoBytes(cBuffer, C.int(size))
		C.VSIFree(cBuffer)
	}
	if cFormat != nil {
		detailedFormat = C.GoString(cFormat)
		C.VSIFree(unsafe.Pointer(cFormat))
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

func gdalSetProjection(dataset GDALDataset, projection string) (err error) {
	cProj := C.CString(projection)
	defer C.free(unsafe.Pointer(cProj))
	err = cplErr(CPLErr(C.GDALSetProjection(dataset.cValue, cProj)))
	return
}

func (ds GDALDataset) SetProjection(projection string) (err error) {
	err = gdalSetProjection(ds, projection)
	return
}

func gdalSetSpatialRef(dataset GDALDataset, srs OGRSpatialReference) (err error) {
	err = cplErr(CPLErr(C.GDALSetSpatialRef(dataset.cValue, srs.cValue)))
	return
}

func (ds GDALDataset) SetSpatialRef(srs OGRSpatialReference) (err error) {
	err = gdalSetSpatialRef(ds, srs)
	return
}

func gdalGetGeoTransform(dataset GDALDataset, geoTransform *[6]float64) (err error) {
	var gt [6]C.double
	err = cplErr(CPLErr(C.GDALGetGeoTransform(dataset.cValue, &gt[0])))
	for i := range gt {
		geoTransform[i] = float64(gt[i])
	}
	return
}

func (ds GDALDataset) GetGeoTransform() (geoTransform [6]float64, err error) {
	err = gdalGetGeoTransform(ds, &geoTransform)
	return
}

func gdalSetGeoTransform(dataset GDALDataset, geoTransform [6]float64) (err error) {
	var gt [6]C.double
	for i, v := range geoTransform {
		gt[i] = C.double(v)
	}
	err = cplErr(CPLErr(C.GDALSetGeoTransform(dataset.cValue, &gt[0])))
	return
}

func (ds GDALDataset) SetGeoTransform(geoTransform [6]float64) (err error) {
	err = gdalSetGeoTransform(ds, geoTransform)
	return
}

func gdalGetExtent(dataset GDALDataset, envelope OGREnvelope, crs OGRSpatialReference) (err error) {
	err = cplErr(CPLErr(C.GDALGetExtent(dataset.cValue, envelope.cValue, crs.cValue)))
	return
}

func (ds GDALDataset) GetExtent(crs OGRSpatialReference) (envelope OGREnvelope, err error) {
	envelope = OGREnvelope{cValue: new(C.OGREnvelope)}
	err = gdalGetExtent(ds, envelope, crs)
	return
}

func gdalGetExtentWGS84LongLat(dataset GDALDataset, envelope OGREnvelope) (err error) {
	err = cplErr(CPLErr(C.GDALGetExtentWGS84LongLat(dataset.cValue, envelope.cValue)))
	return
}

func (ds GDALDataset) GetExtentWGS84LongLat() (envelope OGREnvelope, err error) {
	envelope = OGREnvelope{cValue: new(C.OGREnvelope)}
	err = gdalGetExtentWGS84LongLat(ds, envelope)
	return
}

func gdalDatasetGeolocationToPixelLine(dataset GDALDataset, geolocX, geolocY float64, srs OGRSpatialReference, pixel, line *float64, transformerOptions []string) (err error) {
	opts, free := cStrings(transformerOptions)
	defer free()
	var cPixel, cLine C.double
	err = cplErr(CPLErr(C.GDALDatasetGeolocationToPixelLine(dataset.cValue, C.double(geolocX), C.double(geolocY), srs.cValue, &cPixel, &cLine, C.CSLConstList(unsafe.Pointer(opts)))))
	*pixel = float64(cPixel)
	*line = float64(cLine)
	return
}

func (ds GDALDataset) GeolocationToPixelLine(geolocX, geolocY float64, srs OGRSpatialReference, transformerOptions []string) (pixel, line float64, err error) {
	err = gdalDatasetGeolocationToPixelLine(ds, geolocX, geolocY, srs, &pixel, &line, transformerOptions)
	return
}

// int CPL_DLL CPL_STDCALL GDALGetGCPCount(GDALDatasetH);
// const char CPL_DLL *CPL_STDCALL GDALGetGCPProjection(GDALDatasetH);
// OGRSpatialReferenceH CPL_DLL GDALGetGCPSpatialRef(GDALDatasetH);
// const GDAL_GCP CPL_DLL *CPL_STDCALL GDALGetGCPs(GDALDatasetH);
// CPLErr CPL_DLL CPL_STDCALL GDALSetGCPs(GDALDatasetH, int, const GDAL_GCP *,
//                                        const char *);
// CPLErr CPL_DLL GDALSetGCPs2(GDALDatasetH, int, const GDAL_GCP *,
//                             OGRSpatialReferenceH);

// void CPL_DLL *CPL_STDCALL GDALGetInternalHandle(GDALDatasetH, const char *);
// int CPL_DLL CPL_STDCALL GDALReferenceDataset(GDALDatasetH);
// int CPL_DLL CPL_STDCALL GDALDereferenceDataset(GDALDatasetH);
// int CPL_DLL CPL_STDCALL GDALReleaseDataset(GDALDatasetH);

// CPLErr CPL_DLL CPL_STDCALL GDALBuildOverviews(GDALDatasetH, const char *, int,
//                                               const int *, int, const int *,
//                                               GDALProgressFunc,
//                                               void *) CPL_WARN_UNUSED_RESULT;
// CPLErr CPL_DLL CPL_STDCALL GDALBuildOverviewsEx(
//     GDALDatasetH, const char *, int, const int *, int, const int *,
//     GDALProgressFunc, void *, CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// void CPL_DLL CPL_STDCALL GDALGetOpenDatasets(GDALDatasetH **hDS, int *pnCount);
// int CPL_DLL CPL_STDCALL GDALGetAccess(GDALDatasetH hDS);
// CPLErr CPL_DLL CPL_STDCALL GDALFlushCache(GDALDatasetH hDS);
// CPLErr CPL_DLL CPL_STDCALL GDALDropCache(GDALDatasetH hDS);

// CPLErr CPL_DLL CPL_STDCALL GDALCreateDatasetMaskBand(GDALDatasetH hDS,
//                                                      int nFlags);

// CPLErr CPL_DLL CPL_STDCALL GDALDatasetCopyWholeRaster(
//     GDALDatasetH hSrcDS, GDALDatasetH hDstDS, CSLConstList papszOptions,
//     GDALProgressFunc pfnProgress, void *pProgressData) CPL_WARN_UNUSED_RESULT;

// CPLErr CPL_DLL CPL_STDCALL GDALRasterBandCopyWholeRaster(
//     GDALRasterBandH hSrcBand, GDALRasterBandH hDstBand,
//     const char *const *constpapszOptions, GDALProgressFunc pfnProgress,
//     void *pProgressData) CPL_WARN_UNUSED_RESULT;

// CPLErr CPL_DLL GDALRegenerateOverviews(GDALRasterBandH hSrcBand,
//                                        int nOverviewCount,
//                                        GDALRasterBandH *pahOverviewBands,
//                                        const char *pszResampling,
//                                        GDALProgressFunc pfnProgress,
//                                        void *pProgressData);

// CPLErr CPL_DLL GDALRegenerateOverviewsEx(GDALRasterBandH hSrcBand,
//                                          int nOverviewCount,
//                                          GDALRasterBandH *pahOverviewBands,
//                                          const char *pszResampling,
//                                          GDALProgressFunc pfnProgress,
//                                          void *pProgressData,
//                                          CSLConstList papszOptions);

// int CPL_DLL GDALDatasetGetLayerCount(GDALDatasetH);
// OGRLayerH CPL_DLL GDALDatasetGetLayer(GDALDatasetH, int);

// /* Defined here to avoid circular dependency with ogr_api.h */
// GDALDatasetH CPL_DLL OGR_L_GetDataset(OGRLayerH hLayer);

// OGRLayerH CPL_DLL GDALDatasetGetLayerByName(GDALDatasetH, const char *);
// int CPL_DLL GDALDatasetIsLayerPrivate(GDALDatasetH, int);
// OGRErr CPL_DLL GDALDatasetDeleteLayer(GDALDatasetH, int);
// OGRLayerH CPL_DLL GDALDatasetCreateLayer(GDALDatasetH, const char *,
//                                          OGRSpatialReferenceH,
//                                          OGRwkbGeometryType, CSLConstList);
// OGRLayerH CPL_DLL GDALDatasetCreateLayerFromGeomFieldDefn(GDALDatasetH,
//                                                           const char *,
//                                                           OGRGeomFieldDefnH,
//                                                           CSLConstList);
// OGRLayerH CPL_DLL GDALDatasetCopyLayer(GDALDatasetH, OGRLayerH, const char *,
//                                        CSLConstList);
// void CPL_DLL GDALDatasetResetReading(GDALDatasetH);
// OGRFeatureH CPL_DLL GDALDatasetGetNextFeature(GDALDatasetH hDS,
//                                               OGRLayerH *phBelongingLayer,
//                                               double *pdfProgressPct,
//                                               GDALProgressFunc pfnProgress,
//                                               void *pProgressData);
// int CPL_DLL GDALDatasetTestCapability(GDALDatasetH, const char *);
// OGRLayerH CPL_DLL GDALDatasetExecuteSQL(GDALDatasetH, const char *,
//                                         OGRGeometryH, const char *);
// OGRErr CPL_DLL GDALDatasetAbortSQL(GDALDatasetH);
// void CPL_DLL GDALDatasetReleaseResultSet(GDALDatasetH, OGRLayerH);
// OGRStyleTableH CPL_DLL GDALDatasetGetStyleTable(GDALDatasetH);
// void CPL_DLL GDALDatasetSetStyleTableDirectly(GDALDatasetH, OGRStyleTableH);
// void CPL_DLL GDALDatasetSetStyleTable(GDALDatasetH, OGRStyleTableH);
// OGRErr CPL_DLL GDALDatasetStartTransaction(GDALDatasetH hDS, int bForce);
// OGRErr CPL_DLL GDALDatasetCommitTransaction(GDALDatasetH hDS);
// OGRErr CPL_DLL GDALDatasetRollbackTransaction(GDALDatasetH hDS);
// void CPL_DLL GDALDatasetClearStatistics(GDALDatasetH hDS);

// GDALMDArrayH CPL_DLL GDALDatasetAsMDArray(
//     GDALDatasetH hDS, CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// char CPL_DLL **GDALDatasetGetFieldDomainNames(GDALDatasetH, CSLConstList)
//     CPL_WARN_UNUSED_RESULT;
// OGRFieldDomainH CPL_DLL GDALDatasetGetFieldDomain(GDALDatasetH hDS,
//                                                   const char *pszName);
// bool CPL_DLL GDALDatasetAddFieldDomain(GDALDatasetH hDS,
//                                        OGRFieldDomainH hFieldDomain,
//                                        char **ppszFailureReason);
// bool CPL_DLL GDALDatasetDeleteFieldDomain(GDALDatasetH hDS, const char *pszName,
//                                           char **ppszFailureReason);
// bool CPL_DLL GDALDatasetUpdateFieldDomain(GDALDatasetH hDS,
//                                           OGRFieldDomainH hFieldDomain,
//                                           char **ppszFailureReason);

// char CPL_DLL **GDALDatasetGetRelationshipNames(GDALDatasetH, CSLConstList)
//     CPL_WARN_UNUSED_RESULT;
// GDALRelationshipH CPL_DLL GDALDatasetGetRelationship(GDALDatasetH hDS,
//                                                      const char *pszName);

// bool CPL_DLL GDALDatasetAddRelationship(GDALDatasetH hDS,
//                                         GDALRelationshipH hRelationship,
//                                         char **ppszFailureReason);
// bool CPL_DLL GDALDatasetDeleteRelationship(GDALDatasetH hDS,
//                                            const char *pszName,
//                                            char **ppszFailureReason);
// bool CPL_DLL GDALDatasetUpdateRelationship(GDALDatasetH hDS,
//                                            GDALRelationshipH hRelationship,
//                                            char **ppszFailureReason);

// /** Type of functions to pass to GDALDatasetSetQueryLoggerFunc
//  * @since GDAL 3.7 */
// typedef void (*GDALQueryLoggerFunc)(const char *pszSQL, const char *pszError,
//                                     int64_t lNumRecords,
//                                     int64_t lExecutionTimeMilliseconds,
//                                     void *pQueryLoggerArg);

// /**
//  * Sets the SQL query logger callback.
//  *
//  * When supported by the driver, the callback will be called with
//  * the executed SQL text, the error message, the execution time in milliseconds,
//  * the number of records fetched/affected and the client status data.
//  *
//  * A value of -1 in the execution time or in the number of records indicates
//  * that the values are unknown.
//  *
//  * @param hDS                   Dataset handle.
//  * @param pfnQueryLoggerFunc    Callback function
//  * @param poQueryLoggerArg      Opaque client status data
//  * @return                      true in case of success.
//  * @since                       GDAL 3.7
//  */
// bool CPL_DLL GDALDatasetSetQueryLoggerFunc(
//     GDALDatasetH hDS, GDALQueryLoggerFunc pfnQueryLoggerFunc,
//     void *poQueryLoggerArg);

// /* ==================================================================== */
// /*      Informational utilities about subdatasets in file names         */
// /* ==================================================================== */

// /**
//  * @brief Returns a new GDALSubdatasetInfo object with methods to extract
//  *        and manipulate subdataset information.
//  *        If the pszFileName argument is not recognized by any driver as
//  *        a subdataset descriptor, NULL is returned.
//  *        The returned object must be freed with GDALDestroySubdatasetInfo().
//  * @param pszFileName           File name with subdataset information
//  * @note                        This method does not check if the subdataset actually exists.
//  * @return                      Opaque pointer to a GDALSubdatasetInfo object or NULL if no drivers accepted the file name.
//  * @since                       GDAL 3.8
//  */
// GDALSubdatasetInfoH CPL_DLL GDALGetSubdatasetInfo(const char *pszFileName);

// /**
//  * @brief Returns the file path component of a
//  *        subdataset descriptor effectively stripping the information about the subdataset
//  *        and returning the "parent" dataset descriptor.
//  *        The returned string must be freed with CPLFree().
//  * @param hInfo                 Pointer to GDALSubdatasetInfo object
//  * @note                        This method does not check if the subdataset actually exists.
//  * @return                      The original string with the subdataset information removed.
//  * @since                       GDAL 3.8
//  */
// char CPL_DLL *GDALSubdatasetInfoGetPathComponent(GDALSubdatasetInfoH hInfo);

// /**
//  * @brief Returns the subdataset component of a subdataset descriptor descriptor.
//  *        The returned string must be freed with CPLFree().
//  * @param hInfo                 Pointer to GDALSubdatasetInfo object
//  * @note                        This method does not check if the subdataset actually exists.
//  * @return                      The subdataset name.
//  * @since                       GDAL 3.8
//  */
// char CPL_DLL *
// GDALSubdatasetInfoGetSubdatasetComponent(GDALSubdatasetInfoH hInfo);

// /**
//  * @brief Replaces the path component of a subdataset descriptor.
//  *        The returned string must be freed with CPLFree().
//  * @param hInfo                 Pointer to GDALSubdatasetInfo object
//  * @param pszNewPath            New path.
//  * @note                        This method does not check if the subdataset actually exists.
//  * @return                      The original subdataset descriptor with the old path component replaced by newPath.
//  * @since                       GDAL 3.8
//  */
// char CPL_DLL *GDALSubdatasetInfoModifyPathComponent(GDALSubdatasetInfoH hInfo,
//                                                     const char *pszNewPath);

// /**
//  * @brief Destroys a GDALSubdatasetInfo object.
//  * @param hInfo                 Pointer to GDALSubdatasetInfo object
//  * @since                       GDAL 3.8
//  */
// void CPL_DLL GDALDestroySubdatasetInfo(GDALSubdatasetInfoH hInfo);

// /* ==================================================================== */
// /*      GDALRasterBand ... one band/channel in a dataset.               */
// /* ==================================================================== */

// /* Note: the only user of SRCVAL() is alg/gdal_simplesurf.cpp */

// /* clang-format off */
// /**
//  * SRCVAL - Macro which may be used by pixel functions to obtain
//  *          a pixel from a source buffer.
//  */
// #define SRCVAL(papoSource, eSrcType, ii) \
//     (eSrcType == GDT_Byte ? CPL_REINTERPRET_CAST(const GByte *, papoSource)[ii] : \
//      eSrcType == GDT_Int8 ? CPL_REINTERPRET_CAST(const GInt8 *, papoSource)[ii] : \
//      eSrcType == GDT_UInt16 ? CPL_REINTERPRET_CAST(const GUInt16 *, papoSource)[ii] : \
//      eSrcType == GDT_Int16 ? CPL_REINTERPRET_CAST(const GInt16 *, papoSource)[ii] : \
//      eSrcType == GDT_UInt32 ? CPL_REINTERPRET_CAST(const GUInt32 *, papoSource)[ii] : \
//      eSrcType == GDT_Int32 ? CPL_REINTERPRET_CAST(const GInt32 *, papoSource)[ii] : \
//      eSrcType == GDT_UInt64 ? CPL_STATIC_CAST(double, CPL_REINTERPRET_CAST(const GUInt64 *, papoSource)[ii]) : \
//      eSrcType == GDT_Int64 ? CPL_STATIC_CAST(double, CPL_REINTERPRET_CAST(const GUInt64 *, papoSource)[ii]) : \
//      eSrcType == GDT_Float32 ? CPL_STATIC_CAST(double, CPL_REINTERPRET_CAST(const float *, papoSource)[ii]) : \
//      eSrcType == GDT_Float64 ? CPL_REINTERPRET_CAST(const double *, papoSource)[ii] : \
//      eSrcType == GDT_CInt16 ? CPL_REINTERPRET_CAST(const GInt16 *, papoSource)[(ii)*2] : \
//      eSrcType == GDT_CInt32 ? CPL_REINTERPRET_CAST(const GInt32 *, papoSource)[(ii)*2] : \
//      eSrcType == GDT_CFloat32 ? CPL_STATIC_CAST(double, CPL_REINTERPRET_CAST(const float *, papoSource)[ii*2]) : \
//      eSrcType == GDT_CFloat64 ? CPL_REINTERPRET_CAST(const double *, papoSource)[ii*2] : \
//      0)

// /* clang-format on */

// /** Type of functions to pass to GDALAddDerivedBandPixelFunc.
//  */
// typedef CPLErr (*GDALDerivedPixelFunc)(void **papoSources, int nSources,
//                                        void *pData, int nBufXSize,
//                                        int nBufYSize, GDALDataType eSrcType,
//                                        GDALDataType eBufType, int nPixelSpace,
//                                        int nLineSpace);

// /** Type of functions to pass to GDALAddDerivedBandPixelFuncWithArgs.
//  * @since GDAL 3.4 */
// typedef CPLErr (*GDALDerivedPixelFuncWithArgs)(
//     void **papoSources, int nSources, void *pData, int nBufXSize, int nBufYSize,
//     GDALDataType eSrcType, GDALDataType eBufType, int nPixelSpace,
//     int nLineSpace, CSLConstList papszFunctionArgs);

// GDALDataType CPL_DLL CPL_STDCALL GDALGetRasterDataType(GDALRasterBandH);
// void CPL_DLL CPL_STDCALL GDALGetBlockSize(GDALRasterBandH, int *pnXSize,
//                                           int *pnYSize);

// CPLErr CPL_DLL CPL_STDCALL GDALGetActualBlockSize(GDALRasterBandH,
//                                                   int nXBlockOff,
//                                                   int nYBlockOff, int *pnXValid,
//                                                   int *pnYValid);

// CPLErr CPL_DLL CPL_STDCALL GDALRasterAdviseRead(GDALRasterBandH hRB,
//                                                 int nDSXOff, int nDSYOff,
//                                                 int nDSXSize, int nDSYSize,
//                                                 int nBXSize, int nBYSize,
//                                                 GDALDataType eBDataType,
//                                                 CSLConstList papszOptions);

// CPLErr CPL_DLL CPL_STDCALL GDALRasterIO(GDALRasterBandH hRBand,
//                                         GDALRWFlag eRWFlag, int nDSXOff,
//                                         int nDSYOff, int nDSXSize, int nDSYSize,
//                                         void *pBuffer, int nBXSize, int nBYSize,
//                                         GDALDataType eBDataType,
//                                         int nPixelSpace,
//                                         int nLineSpace) CPL_WARN_UNUSED_RESULT;
// CPLErr CPL_DLL CPL_STDCALL GDALRasterIOEx(
//     GDALRasterBandH hRBand, GDALRWFlag eRWFlag, int nDSXOff, int nDSYOff,
//     int nDSXSize, int nDSYSize, void *pBuffer, int nBXSize, int nBYSize,
//     GDALDataType eBDataType, GSpacing nPixelSpace, GSpacing nLineSpace,
//     GDALRasterIOExtraArg *psExtraArg) CPL_WARN_UNUSED_RESULT;
// CPLErr CPL_DLL CPL_STDCALL GDALReadBlock(GDALRasterBandH, int, int,
//                                          void *) CPL_WARN_UNUSED_RESULT;
// CPLErr CPL_DLL CPL_STDCALL GDALWriteBlock(GDALRasterBandH, int, int,
//                                           void *) CPL_WARN_UNUSED_RESULT;
// int CPL_DLL CPL_STDCALL GDALGetRasterBandXSize(GDALRasterBandH);
// int CPL_DLL CPL_STDCALL GDALGetRasterBandYSize(GDALRasterBandH);
// GDALAccess CPL_DLL CPL_STDCALL GDALGetRasterAccess(GDALRasterBandH);
// int CPL_DLL CPL_STDCALL GDALGetBandNumber(GDALRasterBandH);
// GDALDatasetH CPL_DLL CPL_STDCALL GDALGetBandDataset(GDALRasterBandH);

// GDALColorInterp
//     CPL_DLL CPL_STDCALL GDALGetRasterColorInterpretation(GDALRasterBandH);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterColorInterpretation(GDALRasterBandH,
//                                                             GDALColorInterp);
// GDALColorTableH CPL_DLL CPL_STDCALL GDALGetRasterColorTable(GDALRasterBandH);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterColorTable(GDALRasterBandH,
//                                                    GDALColorTableH);
// int CPL_DLL CPL_STDCALL GDALHasArbitraryOverviews(GDALRasterBandH);
// int CPL_DLL CPL_STDCALL GDALGetOverviewCount(GDALRasterBandH);
// GDALRasterBandH CPL_DLL CPL_STDCALL GDALGetOverview(GDALRasterBandH, int);
// double CPL_DLL CPL_STDCALL GDALGetRasterNoDataValue(GDALRasterBandH, int *);
// int64_t CPL_DLL CPL_STDCALL GDALGetRasterNoDataValueAsInt64(GDALRasterBandH,
//                                                             int *);
// uint64_t CPL_DLL CPL_STDCALL GDALGetRasterNoDataValueAsUInt64(GDALRasterBandH,
//                                                               int *);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterNoDataValue(GDALRasterBandH, double);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterNoDataValueAsInt64(GDALRasterBandH,
//                                                            int64_t);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterNoDataValueAsUInt64(GDALRasterBandH,
//                                                             uint64_t);
// CPLErr CPL_DLL CPL_STDCALL GDALDeleteRasterNoDataValue(GDALRasterBandH);
// char CPL_DLL **CPL_STDCALL GDALGetRasterCategoryNames(GDALRasterBandH);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterCategoryNames(GDALRasterBandH,
//                                                       CSLConstList);
// double CPL_DLL CPL_STDCALL GDALGetRasterMinimum(GDALRasterBandH,
//                                                 int *pbSuccess);
// double CPL_DLL CPL_STDCALL GDALGetRasterMaximum(GDALRasterBandH,
//                                                 int *pbSuccess);
// CPLErr CPL_DLL CPL_STDCALL GDALGetRasterStatistics(
//     GDALRasterBandH, int bApproxOK, int bForce, double *pdfMin, double *pdfMax,
//     double *pdfMean, double *pdfStdDev);
// CPLErr CPL_DLL CPL_STDCALL
// GDALComputeRasterStatistics(GDALRasterBandH, int bApproxOK, double *pdfMin,
//                             double *pdfMax, double *pdfMean, double *pdfStdDev,
//                             GDALProgressFunc pfnProgress, void *pProgressData);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterStatistics(GDALRasterBandH hBand,
//                                                    double dfMin, double dfMax,
//                                                    double dfMean,
//                                                    double dfStdDev);

// GDALMDArrayH
//     CPL_DLL GDALRasterBandAsMDArray(GDALRasterBandH) CPL_WARN_UNUSED_RESULT;

// const char CPL_DLL *CPL_STDCALL GDALGetRasterUnitType(GDALRasterBandH);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterUnitType(GDALRasterBandH hBand,
//                                                  const char *pszNewValue);
// double CPL_DLL CPL_STDCALL GDALGetRasterOffset(GDALRasterBandH, int *pbSuccess);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterOffset(GDALRasterBandH hBand,
//                                                double dfNewOffset);
// double CPL_DLL CPL_STDCALL GDALGetRasterScale(GDALRasterBandH, int *pbSuccess);
// CPLErr CPL_DLL CPL_STDCALL GDALSetRasterScale(GDALRasterBandH hBand,
//                                               double dfNewOffset);
// CPLErr CPL_DLL CPL_STDCALL GDALComputeRasterMinMax(GDALRasterBandH hBand,
//                                                    int bApproxOK,
//                                                    double adfMinMax[2]);
// CPLErr CPL_DLL GDALComputeRasterMinMaxLocation(GDALRasterBandH hBand,
//                                                double *pdfMin, double *pdfMax,
//                                                int *pnMinX, int *pnMinY,
//                                                int *pnMaxX, int *pnMaxY);
// CPLErr CPL_DLL CPL_STDCALL GDALFlushRasterCache(GDALRasterBandH hBand);
// CPLErr CPL_DLL CPL_STDCALL GDALDropRasterCache(GDALRasterBandH hBand);
// CPLErr CPL_DLL CPL_STDCALL GDALGetRasterHistogram(
//     GDALRasterBandH hBand, double dfMin, double dfMax, int nBuckets,
//     int *panHistogram, int bIncludeOutOfRange, int bApproxOK,
//     GDALProgressFunc pfnProgress, void *pProgressData)
//     /*! @cond Doxygen_Suppress */
//     CPL_WARN_DEPRECATED("Use GDALGetRasterHistogramEx() instead")
//     /*! @endcond */
//     ;
// CPLErr CPL_DLL CPL_STDCALL GDALGetRasterHistogramEx(
//     GDALRasterBandH hBand, double dfMin, double dfMax, int nBuckets,
//     GUIntBig *panHistogram, int bIncludeOutOfRange, int bApproxOK,
//     GDALProgressFunc pfnProgress, void *pProgressData);
// CPLErr CPL_DLL CPL_STDCALL
// GDALGetDefaultHistogram(GDALRasterBandH hBand, double *pdfMin, double *pdfMax,
//                         int *pnBuckets, int **ppanHistogram, int bForce,
//                         GDALProgressFunc pfnProgress, void *pProgressData)
//     /*! @cond Doxygen_Suppress */
//     CPL_WARN_DEPRECATED("Use GDALGetDefaultHistogramEx() instead")
//     /*! @endcond */
//     ;
// CPLErr CPL_DLL CPL_STDCALL
// GDALGetDefaultHistogramEx(GDALRasterBandH hBand, double *pdfMin, double *pdfMax,
//                           int *pnBuckets, GUIntBig **ppanHistogram, int bForce,
//                           GDALProgressFunc pfnProgress, void *pProgressData);
// CPLErr CPL_DLL CPL_STDCALL GDALSetDefaultHistogram(GDALRasterBandH hBand,
//                                                    double dfMin, double dfMax,
//                                                    int nBuckets,
//                                                    int *panHistogram)
//     /*! @cond Doxygen_Suppress */
//     CPL_WARN_DEPRECATED("Use GDALSetDefaultHistogramEx() instead")
//     /*! @endcond */
//     ;
// CPLErr CPL_DLL CPL_STDCALL GDALSetDefaultHistogramEx(GDALRasterBandH hBand,
//                                                      double dfMin, double dfMax,
//                                                      int nBuckets,
//                                                      GUIntBig *panHistogram);
// int CPL_DLL CPL_STDCALL GDALGetRandomRasterSample(GDALRasterBandH, int,
//                                                   float *);
// GDALRasterBandH CPL_DLL CPL_STDCALL GDALGetRasterSampleOverview(GDALRasterBandH,
//                                                                 int);
// GDALRasterBandH CPL_DLL CPL_STDCALL
//     GDALGetRasterSampleOverviewEx(GDALRasterBandH, GUIntBig);
// CPLErr CPL_DLL CPL_STDCALL GDALFillRaster(GDALRasterBandH hBand,
//                                           double dfRealValue,
//                                           double dfImaginaryValue);
// CPLErr CPL_DLL CPL_STDCALL GDALComputeBandStats(
//     GDALRasterBandH hBand, int nSampleStep, double *pdfMean, double *pdfStdDev,
//     GDALProgressFunc pfnProgress, void *pProgressData);
// CPLErr CPL_DLL GDALOverviewMagnitudeCorrection(GDALRasterBandH hBaseBand,
//                                                int nOverviewCount,
//                                                GDALRasterBandH *pahOverviews,
//                                                GDALProgressFunc pfnProgress,
//                                                void *pProgressData);

// GDALRasterAttributeTableH CPL_DLL CPL_STDCALL
// GDALGetDefaultRAT(GDALRasterBandH hBand);
// CPLErr CPL_DLL CPL_STDCALL GDALSetDefaultRAT(GDALRasterBandH,
//                                              GDALRasterAttributeTableH);
// CPLErr CPL_DLL CPL_STDCALL GDALAddDerivedBandPixelFunc(
//     const char *pszName, GDALDerivedPixelFunc pfnPixelFunc);
// CPLErr CPL_DLL CPL_STDCALL GDALAddDerivedBandPixelFuncWithArgs(
//     const char *pszName, GDALDerivedPixelFuncWithArgs pfnPixelFunc,
//     const char *pszMetadata);

// CPLErr CPL_DLL GDALRasterInterpolateAtPoint(GDALRasterBandH hBand,
//                                             double dfPixel, double dfLine,
//                                             GDALRIOResampleAlg eInterpolation,
//                                             double *pdfRealValue,
//                                             double *pdfImagValue);

// CPLErr CPL_DLL GDALRasterInterpolateAtGeolocation(
//     GDALRasterBandH hBand, double dfGeolocX, double dfGeolocY,
//     OGRSpatialReferenceH hSRS, GDALRIOResampleAlg eInterpolation,
//     double *pdfRealValue, double *pdfImagValue,
//     CSLConstList papszTransformerOptions);

// /** Generic pointer for the working structure of VRTProcessedDataset
//  * function. */
// typedef void *VRTPDWorkingDataPtr;

// /** Initialization function to pass to GDALVRTRegisterProcessedDatasetFunc.
//  *
//  * This initialization function is called for each step of a VRTProcessedDataset
//  * that uses the related algorithm.
//  * The initialization function returns the output data type, output band count
//  * and potentially initializes a working structure, typically parsing arguments.
//  *
//  * @param pszFuncName Function name. Must be unique and not null.
//  * @param pUserData User data. May be nullptr. Must remain valid during the
//  *                  lifetime of GDAL.
//  * @param papszFunctionArgs Function arguments as a list of key=value pairs.
//  * @param nInBands Number of input bands.
//  * @param eInDT Input data type.
//  * @param[in,out] padfInNoData Array of nInBands values for the input nodata
//  *                             value. The init function may also override them.
//  * @param[in,out] pnOutBands Pointer whose value must be set to the number of
//  *                           output bands. This will be set to 0 by the caller
//  *                           when calling the function, unless this is the
//  *                           final step, in which case it will be initialized
//  *                           with the number of expected output bands.
//  * @param[out] peOutDT Pointer whose value must be set to the output
//  *                     data type.
//  * @param[in,out] ppadfOutNoData Pointer to an array of *pnOutBands values
//  *                               for the output nodata value that the
//  *                               function must set.
//  *                               For non-final steps, *ppadfOutNoData
//  *                               will be nullptr and it is the responsibility
//  *                               of the function to CPLMalloc()'ate it.
//  *                               If this is the final step, it will be
//  *                               already allocated and initialized with the
//  *                               expected nodata values from the output
//  *                               dataset (if the init function need to
//  *                               reallocate it, it must use CPLRealloc())
//  * @param pszVRTPath Directory of the VRT
//  * @param[out] ppWorkingData Pointer whose value must be set to a working
//  *                           structure, or nullptr.
//  * @return CE_None in case of success, error otherwise.
//  * @since GDAL 3.9 */
// typedef CPLErr (*GDALVRTProcessedDatasetFuncInit)(
//     const char *pszFuncName, void *pUserData, CSLConstList papszFunctionArgs,
//     int nInBands, GDALDataType eInDT, double *padfInNoData, int *pnOutBands,
//     GDALDataType *peOutDT, double **ppadfOutNoData, const char *pszVRTPath,
//     VRTPDWorkingDataPtr *ppWorkingData);

// /** Free function to pass to GDALVRTRegisterProcessedDatasetFunc.
//  *
//  * @param pszFuncName Function name. Must be unique and not null.
//  * @param pUserData User data. May be nullptr. Must remain valid during the
//  *                  lifetime of GDAL.
//  * @param pWorkingData Value of the *ppWorkingData output parameter of
//  *                     GDALVRTProcessedDatasetFuncInit.
//  * @since GDAL 3.9
//  */
// typedef void (*GDALVRTProcessedDatasetFuncFree)(
//     const char *pszFuncName, void *pUserData, VRTPDWorkingDataPtr pWorkingData);

// /** Processing function to pass to GDALVRTRegisterProcessedDatasetFunc.
//  * @param pszFuncName Function name. Must be unique and not null.
//  * @param pUserData User data. May be nullptr. Must remain valid during the
//  *                  lifetime of GDAL.
//  * @param pWorkingData Value of the *ppWorkingData output parameter of
//  *                     GDALVRTProcessedDatasetFuncInit.
//  * @param papszFunctionArgs Function arguments as a list of key=value pairs.
//  * @param nBufXSize Width in pixels of pInBuffer and pOutBuffer
//  * @param nBufYSize Height in pixels of pInBuffer and pOutBuffer
//  * @param pInBuffer Input buffer. It is pixel-interleaved
//  *                  (i.e. R00,G00,B00,R01,G01,B01, etc.)
//  * @param nInBufferSize Size in bytes of pInBuffer
//  * @param eInDT Data type of pInBuffer
//  * @param nInBands Number of bands in pInBuffer.
//  * @param padfInNoData Input nodata values.
//  * @param pOutBuffer Output buffer. It is pixel-interleaved
//  *                   (i.e. R00,G00,B00,R01,G01,B01, etc.)
//  * @param nOutBufferSize Size in bytes of pOutBuffer
//  * @param eOutDT Data type of pOutBuffer
//  * @param nOutBands Number of bands in pOutBuffer.
//  * @param padfOutNoData Input nodata values.
//  * @param dfSrcXOff Source X coordinate in pixel of the top-left of the region
//  * @param dfSrcYOff Source Y coordinate in pixel of the top-left of the region
//  * @param dfSrcXSize Width in pixels of the region
//  * @param dfSrcYSize Height in pixels of the region
//  * @param adfSrcGT Source geotransform
//  * @param pszVRTPath Directory of the VRT
//  * @param papszExtra Extra arguments (unused for now)
//  * @since GDAL 3.9
//  */
// typedef CPLErr (*GDALVRTProcessedDatasetFuncProcess)(
//     const char *pszFuncName, void *pUserData, VRTPDWorkingDataPtr pWorkingData,
//     CSLConstList papszFunctionArgs, int nBufXSize, int nBufYSize,
//     const void *pInBuffer, size_t nInBufferSize, GDALDataType eInDT,
//     int nInBands, const double *padfInNoData, void *pOutBuffer,
//     size_t nOutBufferSize, GDALDataType eOutDT, int nOutBands,
//     const double *padfOutNoData, double dfSrcXOff, double dfSrcYOff,
//     double dfSrcXSize, double dfSrcYSize, const double adfSrcGT[/*6*/],
//     const char *pszVRTPath, CSLConstList papszExtra);

// CPLErr CPL_DLL GDALVRTRegisterProcessedDatasetFunc(
//     const char *pszFuncName, void *pUserData, const char *pszXMLMetadata,
//     GDALDataType eRequestedInputDT, const GDALDataType *paeSupportedInputDT,
//     size_t nSupportedInputDTSize, const int *panSupportedInputBandCount,
//     size_t nSupportedInputBandCountSize,
//     GDALVRTProcessedDatasetFuncInit pfnInit,
//     GDALVRTProcessedDatasetFuncFree pfnFree,
//     GDALVRTProcessedDatasetFuncProcess pfnProcess, CSLConstList papszOptions);

// GDALRasterBandH CPL_DLL CPL_STDCALL GDALGetMaskBand(GDALRasterBandH hBand);
// int CPL_DLL CPL_STDCALL GDALGetMaskFlags(GDALRasterBandH hBand);
// CPLErr CPL_DLL CPL_STDCALL GDALCreateMaskBand(GDALRasterBandH hBand,
//                                               int nFlags);
// bool CPL_DLL GDALIsMaskBand(GDALRasterBandH hBand);

// /** Flag returned by GDALGetMaskFlags() to indicate that all pixels are valid */
// #define GMF_ALL_VALID 0x01
// /** Flag returned by GDALGetMaskFlags() to indicate that the mask band is
//  * valid for all bands */
// #define GMF_PER_DATASET 0x02
// /** Flag returned by GDALGetMaskFlags() to indicate that the mask band is
//  * an alpha band */
// #define GMF_ALPHA 0x04
// /** Flag returned by GDALGetMaskFlags() to indicate that the mask band is
//  * computed from nodata values */
// #define GMF_NODATA 0x08

// /** Flag returned by GDALGetDataCoverageStatus() when the driver does not
//  * implement GetDataCoverageStatus(). This flag should be returned together
//  * with GDAL_DATA_COVERAGE_STATUS_DATA */
// #define GDAL_DATA_COVERAGE_STATUS_UNIMPLEMENTED 0x01

// /** Flag returned by GDALGetDataCoverageStatus() when there is (potentially)
//  * data in the queried window. Can be combined with the binary or operator
//  * with GDAL_DATA_COVERAGE_STATUS_UNIMPLEMENTED or
//  * GDAL_DATA_COVERAGE_STATUS_EMPTY */
// #define GDAL_DATA_COVERAGE_STATUS_DATA 0x02

// /** Flag returned by GDALGetDataCoverageStatus() when there is nodata in the
//  * queried window. This is typically identified by the concept of missing block
//  * in formats that supports it.
//  * Can be combined with the binary or operator with
//  * GDAL_DATA_COVERAGE_STATUS_DATA */
// #define GDAL_DATA_COVERAGE_STATUS_EMPTY 0x04

// int CPL_DLL CPL_STDCALL GDALGetDataCoverageStatus(GDALRasterBandH hBand,
//                                                   int nXOff, int nYOff,
//                                                   int nXSize, int nYSize,
//                                                   int nMaskFlagStop,
//                                                   double *pdfDataPct);

// void CPL_DLL GDALComputedRasterBandRelease(GDALComputedRasterBandH hBand);

// /** Raster algebra unary operation */
// typedef enum
// {
//     /** Logical not */
//     GRAUO_LOGICAL_NOT,
//     /** Absolute value (module for complex data type) */
//     GRAUO_ABS,
//     /** Square root */
//     GRAUO_SQRT,
//     /** Natural logarithm (``ln``) */
//     GRAUO_LOG,
//     /** Logarithm base 10 */
//     GRAUO_LOG10,
// } GDALRasterAlgebraUnaryOperation;

// GDALComputedRasterBandH CPL_DLL GDALRasterBandUnaryOp(
//     GDALRasterBandH hBand,
//     GDALRasterAlgebraUnaryOperation eOp) CPL_WARN_UNUSED_RESULT;

// /** Raster algebra binary operation */
// typedef enum
// {
//     /** Addition */
//     GRABO_ADD,
//     /** Subtraction */
//     GRABO_SUB,
//     /** Multiplication */
//     GRABO_MUL,
//     /** Division */
//     GRABO_DIV,
//     /** Power */
//     GRABO_POW,
//     /** Strictly greater than test*/
//     GRABO_GT,
//     /** Greater or equal to test */
//     GRABO_GE,
//     /** Strictly lesser than test */
//     GRABO_LT,
//     /** Lesser or equal to test */
//     GRABO_LE,
//     /** Equality test */
//     GRABO_EQ,
//     /** Non-equality test */
//     GRABO_NE,
//     /** Logical and */
//     GRABO_LOGICAL_AND,
//     /** Logical or */
//     GRABO_LOGICAL_OR
// } GDALRasterAlgebraBinaryOperation;

// GDALComputedRasterBandH CPL_DLL GDALRasterBandBinaryOpBand(
//     GDALRasterBandH hBand, GDALRasterAlgebraBinaryOperation eOp,
//     GDALRasterBandH hOtherBand) CPL_WARN_UNUSED_RESULT;
// GDALComputedRasterBandH CPL_DLL GDALRasterBandBinaryOpDouble(
//     GDALRasterBandH hBand, GDALRasterAlgebraBinaryOperation eOp,
//     double constant) CPL_WARN_UNUSED_RESULT;
// GDALComputedRasterBandH CPL_DLL GDALRasterBandBinaryOpDoubleToBand(
//     double constant, GDALRasterAlgebraBinaryOperation eOp,
//     GDALRasterBandH hBand) CPL_WARN_UNUSED_RESULT;

// GDALComputedRasterBandH CPL_DLL
// GDALRasterBandIfThenElse(GDALRasterBandH hCondBand, GDALRasterBandH hThenBand,
//                          GDALRasterBandH hElseBand) CPL_WARN_UNUSED_RESULT;

// GDALComputedRasterBandH CPL_DLL GDALRasterBandAsDataType(
//     GDALRasterBandH hBand, GDALDataType eDT) CPL_WARN_UNUSED_RESULT;

// GDALComputedRasterBandH CPL_DLL GDALMaximumOfNBands(
//     size_t nBandCount, GDALRasterBandH *pahBands) CPL_WARN_UNUSED_RESULT;
// GDALComputedRasterBandH CPL_DLL GDALRasterBandMaxConstant(
//     GDALRasterBandH hBand, double dfConstant) CPL_WARN_UNUSED_RESULT;
// GDALComputedRasterBandH CPL_DLL GDALMinimumOfNBands(
//     size_t nBandCount, GDALRasterBandH *pahBands) CPL_WARN_UNUSED_RESULT;
// GDALComputedRasterBandH CPL_DLL GDALRasterBandMinConstant(
//     GDALRasterBandH hBand, double dfConstant) CPL_WARN_UNUSED_RESULT;
// GDALComputedRasterBandH CPL_DLL GDALMeanOfNBands(
//     size_t nBandCount, GDALRasterBandH *pahBands) CPL_WARN_UNUSED_RESULT;

// /* ==================================================================== */
// /*     GDALAsyncReader                                                  */
// /* ==================================================================== */

// GDALAsyncStatusType CPL_DLL CPL_STDCALL GDALARGetNextUpdatedRegion(
//     GDALAsyncReaderH hARIO, double dfTimeout, int *pnXBufOff, int *pnYBufOff,
//     int *pnXBufSize, int *pnYBufSize);
// int CPL_DLL CPL_STDCALL GDALARLockBuffer(GDALAsyncReaderH hARIO,
//                                          double dfTimeout);
// void CPL_DLL CPL_STDCALL GDALARUnlockBuffer(GDALAsyncReaderH hARIO);

// /* -------------------------------------------------------------------- */
// /*      Helper functions.                                               */
// /* -------------------------------------------------------------------- */
// int CPL_DLL CPL_STDCALL GDALGeneralCmdLineProcessor(int nArgc,
//                                                     char ***ppapszArgv,
//                                                     int nOptions);
// void CPL_DLL CPL_STDCALL GDALSwapWords(void *pData, int nWordSize,
//                                        int nWordCount, int nWordSkip);
// void CPL_DLL CPL_STDCALL GDALSwapWordsEx(void *pData, int nWordSize,
//                                          size_t nWordCount, int nWordSkip);

// void CPL_DLL CPL_STDCALL GDALCopyWords(const void *CPL_RESTRICT pSrcData,
//                                        GDALDataType eSrcType,
//                                        int nSrcPixelOffset,
//                                        void *CPL_RESTRICT pDstData,
//                                        GDALDataType eDstType,
//                                        int nDstPixelOffset, int nWordCount);

// void CPL_DLL CPL_STDCALL GDALCopyWords64(
//     const void *CPL_RESTRICT pSrcData, GDALDataType eSrcType,
//     int nSrcPixelOffset, void *CPL_RESTRICT pDstData, GDALDataType eDstType,
//     int nDstPixelOffset, GPtrDiff_t nWordCount);

// void CPL_DLL GDALCopyBits(const GByte *pabySrcData, int nSrcOffset,
//                           int nSrcStep, GByte *pabyDstData, int nDstOffset,
//                           int nDstStep, int nBitCount, int nStepCount);

// void CPL_DLL GDALDeinterleave(const void *pSourceBuffer, GDALDataType eSourceDT,
//                               int nComponents, void **ppDestBuffer,
//                               GDALDataType eDestDT, size_t nIters);

// void CPL_DLL GDALTranspose2D(const void *pSrc, GDALDataType eSrcType,
//                              void *pDst, GDALDataType eDstType,
//                              size_t nSrcWidth, size_t nSrcHeight);

// double CPL_DLL GDALGetNoDataReplacementValue(GDALDataType, double);

// int CPL_DLL CPL_STDCALL GDALLoadWorldFile(const char *, double *);
// int CPL_DLL CPL_STDCALL GDALReadWorldFile(const char *, const char *, double *);
// int CPL_DLL CPL_STDCALL GDALWriteWorldFile(const char *, const char *,
//                                            double *);
// int CPL_DLL CPL_STDCALL GDALLoadTabFile(const char *, double *, char **, int *,
//                                         GDAL_GCP **);
// int CPL_DLL CPL_STDCALL GDALReadTabFile(const char *, double *, char **, int *,
//                                         GDAL_GCP **);
// int CPL_DLL CPL_STDCALL GDALLoadOziMapFile(const char *, double *, char **,
//                                            int *, GDAL_GCP **);
// int CPL_DLL CPL_STDCALL GDALReadOziMapFile(const char *, double *, char **,
//                                            int *, GDAL_GCP **);

// const char CPL_DLL *CPL_STDCALL GDALDecToDMS(double, const char *, int);
// double CPL_DLL CPL_STDCALL GDALPackedDMSToDec(double);
// double CPL_DLL CPL_STDCALL GDALDecToPackedDMS(double);

// /* Note to developers : please keep this section in sync with ogr_core.h */

// #ifndef GDAL_VERSION_INFO_DEFINED
// #ifndef DOXYGEN_SKIP
// #define GDAL_VERSION_INFO_DEFINED
// #endif
// const char CPL_DLL *CPL_STDCALL GDALVersionInfo(const char *);
// #endif

// #ifndef GDAL_CHECK_VERSION

// int CPL_DLL CPL_STDCALL GDALCheckVersion(int nVersionMajor, int nVersionMinor,
//                                          const char *pszCallingComponentName);

// /** Helper macro for GDALCheckVersion()
//   @see GDALCheckVersion()
//   */
// #define GDAL_CHECK_VERSION(pszCallingComponentName)                            \
//     GDALCheckVersion(GDAL_VERSION_MAJOR, GDAL_VERSION_MINOR,                   \
//                      pszCallingComponentName)

// #endif

// /*! @cond Doxygen_Suppress */
// #ifdef GDAL_COMPILATION
// #define GDALExtractRPCInfoV1 GDALExtractRPCInfo
// #else
// #define GDALRPCInfo GDALRPCInfoV2
// #define GDALExtractRPCInfo GDALExtractRPCInfoV2
// #endif

// /* Deprecated: use GDALRPCInfoV2 */
// typedef struct
// {
//     double dfLINE_OFF;   /*!< Line offset */
//     double dfSAMP_OFF;   /*!< Sample/Pixel offset */
//     double dfLAT_OFF;    /*!< Latitude offset */
//     double dfLONG_OFF;   /*!< Longitude offset */
//     double dfHEIGHT_OFF; /*!< Height offset */

//     double dfLINE_SCALE;   /*!< Line scale */
//     double dfSAMP_SCALE;   /*!< Sample/Pixel scale */
//     double dfLAT_SCALE;    /*!< Latitude scale */
//     double dfLONG_SCALE;   /*!< Longitude scale */
//     double dfHEIGHT_SCALE; /*!< Height scale */

//     double adfLINE_NUM_COEFF[20]; /*!< Line Numerator Coefficients */
//     double adfLINE_DEN_COEFF[20]; /*!< Line Denominator Coefficients */
//     double adfSAMP_NUM_COEFF[20]; /*!< Sample/Pixel Numerator Coefficients */
//     double adfSAMP_DEN_COEFF[20]; /*!< Sample/Pixel Denominator Coefficients */

//     double dfMIN_LONG; /*!< Minimum longitude */
//     double dfMIN_LAT;  /*!< Minimum latitude */
//     double dfMAX_LONG; /*!< Maximum longitude */
//     double dfMAX_LAT;  /*!< Maximum latitude */
// } GDALRPCInfoV1;

// /*! @endcond */

// /** Structure to store Rational Polynomial Coefficients / Rigorous Projection
//  * Model. See http://geotiff.maptools.org/rpc_prop.html */
// typedef struct
// {
//     double dfLINE_OFF;   /*!< Line offset */
//     double dfSAMP_OFF;   /*!< Sample/Pixel offset */
//     double dfLAT_OFF;    /*!< Latitude offset */
//     double dfLONG_OFF;   /*!< Longitude offset */
//     double dfHEIGHT_OFF; /*!< Height offset */

//     double dfLINE_SCALE;   /*!< Line scale */
//     double dfSAMP_SCALE;   /*!< Sample/Pixel scale */
//     double dfLAT_SCALE;    /*!< Latitude scale */
//     double dfLONG_SCALE;   /*!< Longitude scale */
//     double dfHEIGHT_SCALE; /*!< Height scale */

//     double adfLINE_NUM_COEFF[20]; /*!< Line Numerator Coefficients */
//     double adfLINE_DEN_COEFF[20]; /*!< Line Denominator Coefficients */
//     double adfSAMP_NUM_COEFF[20]; /*!< Sample/Pixel Numerator Coefficients */
//     double adfSAMP_DEN_COEFF[20]; /*!< Sample/Pixel Denominator Coefficients */

//     double dfMIN_LONG; /*!< Minimum longitude */
//     double dfMIN_LAT;  /*!< Minimum latitude */
//     double dfMAX_LONG; /*!< Maximum longitude */
//     double dfMAX_LAT;  /*!< Maximum latitude */

//     /* Those fields should be at the end. And all above fields should be the
//      * same as in GDALRPCInfoV1 */
//     double dfERR_BIAS; /*!< Bias error */
//     double dfERR_RAND; /*!< Random error */
// } GDALRPCInfoV2;

// /*! @cond Doxygen_Suppress */
// int CPL_DLL CPL_STDCALL GDALExtractRPCInfoV1(CSLConstList, GDALRPCInfoV1 *);
// /*! @endcond */
// int CPL_DLL CPL_STDCALL GDALExtractRPCInfoV2(CSLConstList, GDALRPCInfoV2 *);

// /* ==================================================================== */
// /*      Color tables.                                                   */
// /* ==================================================================== */

// /** Color tuple */
// typedef struct
// {
//     /*! gray, red, cyan or hue */
//     short c1;

//     /*! green, magenta, or lightness */
//     short c2;

//     /*! blue, yellow, or saturation */
//     short c3;

//     /*! alpha or blackband */
//     short c4;
// } GDALColorEntry;

// GDALColorTableH CPL_DLL CPL_STDCALL GDALCreateColorTable(GDALPaletteInterp)
//     CPL_WARN_UNUSED_RESULT;
// void CPL_DLL CPL_STDCALL GDALDestroyColorTable(GDALColorTableH);
// GDALColorTableH CPL_DLL CPL_STDCALL GDALCloneColorTable(GDALColorTableH);
// GDALPaletteInterp
//     CPL_DLL CPL_STDCALL GDALGetPaletteInterpretation(GDALColorTableH);
// int CPL_DLL CPL_STDCALL GDALGetColorEntryCount(GDALColorTableH);
// const GDALColorEntry CPL_DLL *CPL_STDCALL GDALGetColorEntry(GDALColorTableH,
//                                                             int);
// int CPL_DLL CPL_STDCALL GDALGetColorEntryAsRGB(GDALColorTableH, int,
//                                                GDALColorEntry *);
// void CPL_DLL CPL_STDCALL GDALSetColorEntry(GDALColorTableH, int,
//                                            const GDALColorEntry *);
// void CPL_DLL CPL_STDCALL GDALCreateColorRamp(GDALColorTableH hTable,
//                                              int nStartIndex,
//                                              const GDALColorEntry *psStartColor,
//                                              int nEndIndex,
//                                              const GDALColorEntry *psEndColor);

// /* ==================================================================== */
// /*      Raster Attribute Table                                          */
// /* ==================================================================== */

// /** Field type of raster attribute table */
// typedef enum
// {
//     /*! Integer field */ GFT_Integer,
//     /*! Floating point (double) field */ GFT_Real,
//     /*! String field */ GFT_String,
//     /*! Boolean field (GDAL >= 3.12) */ GFT_Boolean,
//     /*! DateTime field (GDAL >= 3.12) */ GFT_DateTime,
//     /*! Geometry field, as WKB (GDAL >= 3.12) */ GFT_WKBGeometry
// } GDALRATFieldType;

// /** First invalid value for the GDALRATFieldType enumeration */
// #define GFT_MaxCount (GFT_WKBGeometry + 1)

// /** Field usage of raster attribute table */
// typedef enum
// {
//     /*! General purpose field. */ GFU_Generic = 0,
//     /*! Histogram pixel count */ GFU_PixelCount = 1,
//     /*! Class name */ GFU_Name = 2,
//     /*! Class range minimum */ GFU_Min = 3,
//     /*! Class range maximum */ GFU_Max = 4,
//     /*! Class value (min=max) */ GFU_MinMax = 5,
//     /*! Red class color (0-255) */ GFU_Red = 6,
//     /*! Green class color (0-255) */ GFU_Green = 7,
//     /*! Blue class color (0-255) */ GFU_Blue = 8,
//     /*! Alpha (0=transparent,255=opaque)*/ GFU_Alpha = 9,
//     /*! Color Range Red Minimum */ GFU_RedMin = 10,
//     /*! Color Range Green Minimum */ GFU_GreenMin = 11,
//     /*! Color Range Blue Minimum */ GFU_BlueMin = 12,
//     /*! Color Range Alpha Minimum */ GFU_AlphaMin = 13,
//     /*! Color Range Red Maximum */ GFU_RedMax = 14,
//     /*! Color Range Green Maximum */ GFU_GreenMax = 15,
//     /*! Color Range Blue Maximum */ GFU_BlueMax = 16,
//     /*! Color Range Alpha Maximum */ GFU_AlphaMax = 17,
//     /*! Maximum GFU value (equals to GFU_AlphaMax+1 currently) */ GFU_MaxCount
// } GDALRATFieldUsage;

// /** RAT table type (thematic or athematic)
//  */
// typedef enum
// {
//     /*! Thematic table type */ GRTT_THEMATIC,
//     /*! Athematic table type */ GRTT_ATHEMATIC
// } GDALRATTableType;

// GDALRasterAttributeTableH CPL_DLL CPL_STDCALL
// GDALCreateRasterAttributeTable(void) CPL_WARN_UNUSED_RESULT;

// void CPL_DLL CPL_STDCALL
//     GDALDestroyRasterAttributeTable(GDALRasterAttributeTableH);

// int CPL_DLL CPL_STDCALL GDALRATGetColumnCount(GDALRasterAttributeTableH);

// const char CPL_DLL *CPL_STDCALL GDALRATGetNameOfCol(GDALRasterAttributeTableH,
//                                                     int);
// GDALRATFieldUsage CPL_DLL CPL_STDCALL
// GDALRATGetUsageOfCol(GDALRasterAttributeTableH, int);
// GDALRATFieldType CPL_DLL CPL_STDCALL
// GDALRATGetTypeOfCol(GDALRasterAttributeTableH, int);

// const char CPL_DLL *GDALGetRATFieldTypeName(GDALRATFieldType);
// const char CPL_DLL *GDALGetRATFieldUsageName(GDALRATFieldUsage);

// int CPL_DLL CPL_STDCALL GDALRATGetColOfUsage(GDALRasterAttributeTableH,
//                                              GDALRATFieldUsage);
// int CPL_DLL CPL_STDCALL GDALRATGetRowCount(GDALRasterAttributeTableH);

// const char CPL_DLL *CPL_STDCALL
// GDALRATGetValueAsString(GDALRasterAttributeTableH, int iRow, int iField);
// int CPL_DLL CPL_STDCALL GDALRATGetValueAsInt(GDALRasterAttributeTableH,
//                                              int iRow, int iField);
// double CPL_DLL CPL_STDCALL GDALRATGetValueAsDouble(GDALRasterAttributeTableH,
//                                                    int iRow, int iField);
// bool CPL_DLL GDALRATGetValueAsBoolean(GDALRasterAttributeTableH, int iRow,
//                                       int iField);

// #ifdef __cplusplus
// extern "C++"
// {
// #endif

//     /** Structure encoding a DateTime field for a GDAL Raster Attribute Table.
//  *
//  * @since 3.12
//  */
//     struct GDALRATDateTime
//     {
//         /*! Year */ int nYear;
//         /*! Month [1, 12] */ int nMonth;
//         /*! Day [1, 31] */ int nDay;
//         /*! Hour [0, 23] */ int nHour;
//         /*! Minute [0, 59] */ int nMinute;
//         /*! Second [0, 61) */ float fSecond;
//         /*! Time zone hour [0, 23] */ int nTimeZoneHour;
//         /*! Time zone minute: 0, 15, 30, 45 */ int nTimeZoneMinute;
//         /*! Whether time zone is positive (or null) */ bool bPositiveTimeZone;
//         /*! Whether this object is valid */ bool bIsValid;

// #ifdef __cplusplus
//         GDALRATDateTime()
//             : nYear(0), nMonth(0), nDay(0), nHour(0), nMinute(0), fSecond(0),
//               nTimeZoneHour(0), nTimeZoneMinute(0), bPositiveTimeZone(false),
//               bIsValid(false)
//         {
//         }
// #endif
//     };

// #ifdef __cplusplus
// }
// #endif

// /*! @cond Doxygen_Suppress */
// typedef struct GDALRATDateTime GDALRATDateTime;
// /*! @endcond */

// CPLErr CPL_DLL GDALRATGetValueAsDateTime(GDALRasterAttributeTableH, int iRow,
//                                          int iField,
//                                          GDALRATDateTime *psDateTime);
// const GByte CPL_DLL *GDALRATGetValueAsWKBGeometry(GDALRasterAttributeTableH,
//                                                   int iRow, int iField,
//                                                   size_t *pnWKBSize);

// void CPL_DLL CPL_STDCALL GDALRATSetValueAsString(GDALRasterAttributeTableH,
//                                                  int iRow, int iField,
//                                                  const char *);
// void CPL_DLL CPL_STDCALL GDALRATSetValueAsInt(GDALRasterAttributeTableH,
//                                               int iRow, int iField, int);
// void CPL_DLL CPL_STDCALL GDALRATSetValueAsDouble(GDALRasterAttributeTableH,
//                                                  int iRow, int iField, double);
// CPLErr CPL_DLL GDALRATSetValueAsBoolean(GDALRasterAttributeTableH, int iRow,
//                                         int iField, bool);
// CPLErr CPL_DLL GDALRATSetValueAsDateTime(GDALRasterAttributeTableH, int iRow,
//                                          int iField,
//                                          const GDALRATDateTime *psDateTime);
// CPLErr CPL_DLL GDALRATSetValueAsWKBGeometry(GDALRasterAttributeTableH, int iRow,
//                                             int iField, const void *pabyWKB,
//                                             size_t nWKBSize);

// int CPL_DLL CPL_STDCALL
// GDALRATChangesAreWrittenToFile(GDALRasterAttributeTableH hRAT);

// CPLErr CPL_DLL CPL_STDCALL GDALRATValuesIOAsDouble(
//     GDALRasterAttributeTableH hRAT, GDALRWFlag eRWFlag, int iField,
//     int iStartRow, int iLength, double *pdfData);
// CPLErr CPL_DLL CPL_STDCALL
// GDALRATValuesIOAsInteger(GDALRasterAttributeTableH hRAT, GDALRWFlag eRWFlag,
//                          int iField, int iStartRow, int iLength, int *pnData);
// CPLErr CPL_DLL CPL_STDCALL GDALRATValuesIOAsString(
//     GDALRasterAttributeTableH hRAT, GDALRWFlag eRWFlag, int iField,
//     int iStartRow, int iLength, char **papszStrList);
// CPLErr CPL_DLL GDALRATValuesIOAsBoolean(GDALRasterAttributeTableH hRAT,
//                                         GDALRWFlag eRWFlag, int iField,
//                                         int iStartRow, int iLength,
//                                         bool *pbData);
// CPLErr CPL_DLL GDALRATValuesIOAsDateTime(GDALRasterAttributeTableH hRAT,
//                                          GDALRWFlag eRWFlag, int iField,
//                                          int iStartRow, int iLength,
//                                          GDALRATDateTime *pasDateTime);
// CPLErr CPL_DLL GDALRATValuesIOAsWKBGeometry(GDALRasterAttributeTableH hRAT,
//                                             GDALRWFlag eRWFlag, int iField,
//                                             int iStartRow, int iLength,
//                                             GByte **ppabyWKB,
//                                             size_t *pnWKBSize);

// void CPL_DLL CPL_STDCALL GDALRATSetRowCount(GDALRasterAttributeTableH, int);
// CPLErr CPL_DLL CPL_STDCALL GDALRATCreateColumn(GDALRasterAttributeTableH,
//                                                const char *, GDALRATFieldType,
//                                                GDALRATFieldUsage);
// CPLErr CPL_DLL CPL_STDCALL GDALRATSetLinearBinning(GDALRasterAttributeTableH,
//                                                    double, double);
// int CPL_DLL CPL_STDCALL GDALRATGetLinearBinning(GDALRasterAttributeTableH,
//                                                 double *, double *);
// CPLErr CPL_DLL CPL_STDCALL GDALRATSetTableType(
//     GDALRasterAttributeTableH hRAT, const GDALRATTableType eInTableType);
// GDALRATTableType CPL_DLL CPL_STDCALL
// GDALRATGetTableType(GDALRasterAttributeTableH hRAT);
// CPLErr CPL_DLL CPL_STDCALL
//     GDALRATInitializeFromColorTable(GDALRasterAttributeTableH, GDALColorTableH);
// GDALColorTableH CPL_DLL CPL_STDCALL
// GDALRATTranslateToColorTable(GDALRasterAttributeTableH, int nEntryCount);
// void CPL_DLL CPL_STDCALL GDALRATDumpReadable(GDALRasterAttributeTableH, FILE *);
// GDALRasterAttributeTableH CPL_DLL CPL_STDCALL
// GDALRATClone(const GDALRasterAttributeTableH);

// void CPL_DLL *CPL_STDCALL GDALRATSerializeJSON(GDALRasterAttributeTableH)
//     CPL_WARN_UNUSED_RESULT;

// int CPL_DLL CPL_STDCALL GDALRATGetRowOfValue(GDALRasterAttributeTableH, double);
// void CPL_DLL CPL_STDCALL GDALRATRemoveStatistics(GDALRasterAttributeTableH);

// /* -------------------------------------------------------------------- */
// /*                          Relationships                               */
// /* -------------------------------------------------------------------- */

// /** Cardinality of relationship.
//  *
//  * @since GDAL 3.6
//  */
// typedef enum
// {
//     /** One-to-one */
//     GRC_ONE_TO_ONE,
//     /** One-to-many */
//     GRC_ONE_TO_MANY,
//     /** Many-to-one */
//     GRC_MANY_TO_ONE,
//     /** Many-to-many */
//     GRC_MANY_TO_MANY,
// } GDALRelationshipCardinality;

// /** Type of relationship.
//  *
//  * @since GDAL 3.6
//  */
// typedef enum
// {
//     /** Composite relationship */
//     GRT_COMPOSITE,
//     /** Association relationship */
//     GRT_ASSOCIATION,
//     /** Aggregation relationship */
//     GRT_AGGREGATION,
// } GDALRelationshipType;

// GDALRelationshipH CPL_DLL GDALRelationshipCreate(const char *, const char *,
//                                                  const char *,
//                                                  GDALRelationshipCardinality);
// void CPL_DLL CPL_STDCALL GDALDestroyRelationship(GDALRelationshipH);
// const char CPL_DLL *GDALRelationshipGetName(GDALRelationshipH);
// GDALRelationshipCardinality
//     CPL_DLL GDALRelationshipGetCardinality(GDALRelationshipH);
// const char CPL_DLL *GDALRelationshipGetLeftTableName(GDALRelationshipH);
// const char CPL_DLL *GDALRelationshipGetRightTableName(GDALRelationshipH);
// const char CPL_DLL *GDALRelationshipGetMappingTableName(GDALRelationshipH);
// void CPL_DLL GDALRelationshipSetMappingTableName(GDALRelationshipH,
//                                                  const char *);
// char CPL_DLL **GDALRelationshipGetLeftTableFields(GDALRelationshipH);
// char CPL_DLL **GDALRelationshipGetRightTableFields(GDALRelationshipH);
// void CPL_DLL GDALRelationshipSetLeftTableFields(GDALRelationshipH,
//                                                 CSLConstList);
// void CPL_DLL GDALRelationshipSetRightTableFields(GDALRelationshipH,
//                                                  CSLConstList);
// char CPL_DLL **GDALRelationshipGetLeftMappingTableFields(GDALRelationshipH);
// char CPL_DLL **GDALRelationshipGetRightMappingTableFields(GDALRelationshipH);
// void CPL_DLL GDALRelationshipSetLeftMappingTableFields(GDALRelationshipH,
//                                                        CSLConstList);
// void CPL_DLL GDALRelationshipSetRightMappingTableFields(GDALRelationshipH,
//                                                         CSLConstList);
// GDALRelationshipType CPL_DLL GDALRelationshipGetType(GDALRelationshipH);
// void CPL_DLL GDALRelationshipSetType(GDALRelationshipH, GDALRelationshipType);
// const char CPL_DLL *GDALRelationshipGetForwardPathLabel(GDALRelationshipH);
// void CPL_DLL GDALRelationshipSetForwardPathLabel(GDALRelationshipH,
//                                                  const char *);
// const char CPL_DLL *GDALRelationshipGetBackwardPathLabel(GDALRelationshipH);
// void CPL_DLL GDALRelationshipSetBackwardPathLabel(GDALRelationshipH,
//                                                   const char *);
// const char CPL_DLL *GDALRelationshipGetRelatedTableType(GDALRelationshipH);
// void CPL_DLL GDALRelationshipSetRelatedTableType(GDALRelationshipH,
//                                                  const char *);

// /* ==================================================================== */
// /*      GDAL Cache Management                                           */
// /* ==================================================================== */

// void CPL_DLL CPL_STDCALL GDALSetCacheMax(int nBytes);
// int CPL_DLL CPL_STDCALL GDALGetCacheMax(void);
// int CPL_DLL CPL_STDCALL GDALGetCacheUsed(void);
// void CPL_DLL CPL_STDCALL GDALSetCacheMax64(GIntBig nBytes);
// GIntBig CPL_DLL CPL_STDCALL GDALGetCacheMax64(void);
// GIntBig CPL_DLL CPL_STDCALL GDALGetCacheUsed64(void);

// int CPL_DLL CPL_STDCALL GDALFlushCacheBlock(void);

// /* ==================================================================== */
// /*      GDAL virtual memory                                             */
// /* ==================================================================== */

// CPLVirtualMem CPL_DLL *GDALDatasetGetVirtualMem(
//     GDALDatasetH hDS, GDALRWFlag eRWFlag, int nXOff, int nYOff, int nXSize,
//     int nYSize, int nBufXSize, int nBufYSize, GDALDataType eBufType,
//     int nBandCount, int *panBandMap, int nPixelSpace, GIntBig nLineSpace,
//     GIntBig nBandSpace, size_t nCacheSize, size_t nPageSizeHint,
//     int bSingleThreadUsage, CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// CPLVirtualMem CPL_DLL *GDALRasterBandGetVirtualMem(
//     GDALRasterBandH hBand, GDALRWFlag eRWFlag, int nXOff, int nYOff, int nXSize,
//     int nYSize, int nBufXSize, int nBufYSize, GDALDataType eBufType,
//     int nPixelSpace, GIntBig nLineSpace, size_t nCacheSize,
//     size_t nPageSizeHint, int bSingleThreadUsage,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// CPLVirtualMem CPL_DLL *
// GDALGetVirtualMemAuto(GDALRasterBandH hBand, GDALRWFlag eRWFlag,
//                       int *pnPixelSpace, GIntBig *pnLineSpace,
//                       CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// /**! Enumeration to describe the tile organization */
// typedef enum
// {
//     /*! Tile Interleaved by Pixel: tile (0,0) with internal band interleaved by
//        pixel organization, tile (1, 0), ...  */
//     GTO_TIP,
//     /*! Band Interleaved by Tile : tile (0,0) of first band, tile (0,0) of
//        second band, ... tile (1,0) of first band, tile (1,0) of second band, ...
//      */
//     GTO_BIT,
//     /*! Band SeQuential : all the tiles of first band, all the tiles of
//        following band... */
//     GTO_BSQ
// } GDALTileOrganization;

// CPLVirtualMem CPL_DLL *GDALDatasetGetTiledVirtualMem(
//     GDALDatasetH hDS, GDALRWFlag eRWFlag, int nXOff, int nYOff, int nXSize,
//     int nYSize, int nTileXSize, int nTileYSize, GDALDataType eBufType,
//     int nBandCount, int *panBandMap, GDALTileOrganization eTileOrganization,
//     size_t nCacheSize, int bSingleThreadUsage,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// CPLVirtualMem CPL_DLL *GDALRasterBandGetTiledVirtualMem(
//     GDALRasterBandH hBand, GDALRWFlag eRWFlag, int nXOff, int nYOff, int nXSize,
//     int nYSize, int nTileXSize, int nTileYSize, GDALDataType eBufType,
//     size_t nCacheSize, int bSingleThreadUsage,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// /* ==================================================================== */
// /*      VRTPansharpenedDataset class.                                   */
// /* ==================================================================== */

// GDALDatasetH CPL_DLL GDALCreatePansharpenedVRT(
//     const char *pszXML, GDALRasterBandH hPanchroBand, int nInputSpectralBands,
//     GDALRasterBandH *pahInputSpectralBands) CPL_WARN_UNUSED_RESULT;

// /* =================================================================== */
// /*      Misc API                                                        */
// /* ==================================================================== */

// CPLXMLNode CPL_DLL *
// GDALGetJPEG2000Structure(const char *pszFilename,
//                          CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// /* ==================================================================== */
// /*      Multidimensional API_api                                       */
// /* ==================================================================== */

// GDALDatasetH CPL_DLL
// GDALCreateMultiDimensional(GDALDriverH hDriver, const char *pszName,
//                            CSLConstList papszRootGroupOptions,
//                            CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// GDALExtendedDataTypeH CPL_DLL GDALExtendedDataTypeCreate(GDALDataType eType)
//     CPL_WARN_UNUSED_RESULT;
// GDALExtendedDataTypeH CPL_DLL GDALExtendedDataTypeCreateString(
//     size_t nMaxStringLength) CPL_WARN_UNUSED_RESULT;
// GDALExtendedDataTypeH CPL_DLL GDALExtendedDataTypeCreateStringEx(
//     size_t nMaxStringLength,
//     GDALExtendedDataTypeSubType eSubType) CPL_WARN_UNUSED_RESULT;
// GDALExtendedDataTypeH CPL_DLL GDALExtendedDataTypeCreateCompound(
//     const char *pszName, size_t nTotalSize, size_t nComponents,
//     const GDALEDTComponentH *comps) CPL_WARN_UNUSED_RESULT;
// void CPL_DLL GDALExtendedDataTypeRelease(GDALExtendedDataTypeH hEDT);
// const char CPL_DLL *GDALExtendedDataTypeGetName(GDALExtendedDataTypeH hEDT);
// GDALExtendedDataTypeClass CPL_DLL
// GDALExtendedDataTypeGetClass(GDALExtendedDataTypeH hEDT);
// GDALDataType CPL_DLL
// GDALExtendedDataTypeGetNumericDataType(GDALExtendedDataTypeH hEDT);
// size_t CPL_DLL GDALExtendedDataTypeGetSize(GDALExtendedDataTypeH hEDT);
// size_t CPL_DLL
// GDALExtendedDataTypeGetMaxStringLength(GDALExtendedDataTypeH hEDT);
// GDALEDTComponentH CPL_DLL *
// GDALExtendedDataTypeGetComponents(GDALExtendedDataTypeH hEDT,
//                                   size_t *pnCount) CPL_WARN_UNUSED_RESULT;
// void CPL_DLL GDALExtendedDataTypeFreeComponents(GDALEDTComponentH *components,
//                                                 size_t nCount);
// int CPL_DLL GDALExtendedDataTypeCanConvertTo(GDALExtendedDataTypeH hSourceEDT,
//                                              GDALExtendedDataTypeH hTargetEDT);
// int CPL_DLL GDALExtendedDataTypeEquals(GDALExtendedDataTypeH hFirstEDT,
//                                        GDALExtendedDataTypeH hSecondEDT);
// GDALExtendedDataTypeSubType CPL_DLL
// GDALExtendedDataTypeGetSubType(GDALExtendedDataTypeH hEDT);
// GDALRasterAttributeTableH CPL_DLL
// GDALExtendedDataTypeGetRAT(GDALExtendedDataTypeH hEDT) CPL_WARN_UNUSED_RESULT;

// GDALEDTComponentH CPL_DLL
// GDALEDTComponentCreate(const char *pszName, size_t nOffset,
//                        GDALExtendedDataTypeH hType) CPL_WARN_UNUSED_RESULT;
// void CPL_DLL GDALEDTComponentRelease(GDALEDTComponentH hComp);
// const char CPL_DLL *GDALEDTComponentGetName(GDALEDTComponentH hComp);
// size_t CPL_DLL GDALEDTComponentGetOffset(GDALEDTComponentH hComp);
// GDALExtendedDataTypeH CPL_DLL GDALEDTComponentGetType(GDALEDTComponentH hComp)
//     CPL_WARN_UNUSED_RESULT;

// GDALGroupH CPL_DLL GDALDatasetGetRootGroup(GDALDatasetH hDS)
//     CPL_WARN_UNUSED_RESULT;
// void CPL_DLL GDALGroupRelease(GDALGroupH hGroup);
// const char CPL_DLL *GDALGroupGetName(GDALGroupH hGroup);
// const char CPL_DLL *GDALGroupGetFullName(GDALGroupH hGroup);
// char CPL_DLL **
// GDALGroupGetMDArrayNames(GDALGroupH hGroup,
//                          CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// char CPL_DLL **GDALGroupGetMDArrayFullNamesRecursive(
//     GDALGroupH hGroup, CSLConstList papszGroupOptions,
//     CSLConstList papszArrayOptions) CPL_WARN_UNUSED_RESULT;
// GDALMDArrayH CPL_DLL
// GDALGroupOpenMDArray(GDALGroupH hGroup, const char *pszMDArrayName,
//                      CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALMDArrayH CPL_DLL GDALGroupOpenMDArrayFromFullname(
//     GDALGroupH hGroup, const char *pszMDArrayName,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALMDArrayH CPL_DLL GDALGroupResolveMDArray(
//     GDALGroupH hGroup, const char *pszName, const char *pszStartingPoint,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// char CPL_DLL **
// GDALGroupGetGroupNames(GDALGroupH hGroup,
//                        CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALGroupH CPL_DLL
// GDALGroupOpenGroup(GDALGroupH hGroup, const char *pszSubGroupName,
//                    CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALGroupH CPL_DLL GDALGroupOpenGroupFromFullname(
//     GDALGroupH hGroup, const char *pszMDArrayName,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// char CPL_DLL **
// GDALGroupGetVectorLayerNames(GDALGroupH hGroup,
//                              CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// OGRLayerH CPL_DLL
// GDALGroupOpenVectorLayer(GDALGroupH hGroup, const char *pszVectorLayerName,
//                          CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALDimensionH CPL_DLL *
// GDALGroupGetDimensions(GDALGroupH hGroup, size_t *pnCount,
//                        CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALAttributeH CPL_DLL GDALGroupGetAttribute(
//     GDALGroupH hGroup, const char *pszName) CPL_WARN_UNUSED_RESULT;
// GDALAttributeH CPL_DLL *
// GDALGroupGetAttributes(GDALGroupH hGroup, size_t *pnCount,
//                        CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// CSLConstList CPL_DLL GDALGroupGetStructuralInfo(GDALGroupH hGroup);
// GDALGroupH CPL_DLL
// GDALGroupCreateGroup(GDALGroupH hGroup, const char *pszSubGroupName,
//                      CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// bool CPL_DLL GDALGroupDeleteGroup(GDALGroupH hGroup, const char *pszName,
//                                   CSLConstList papszOptions);
// GDALDimensionH CPL_DLL GDALGroupCreateDimension(
//     GDALGroupH hGroup, const char *pszName, const char *pszType,
//     const char *pszDirection, GUInt64 nSize,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALMDArrayH CPL_DLL GDALGroupCreateMDArray(
//     GDALGroupH hGroup, const char *pszName, size_t nDimensions,
//     GDALDimensionH *pahDimensions, GDALExtendedDataTypeH hEDT,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// bool CPL_DLL GDALGroupDeleteMDArray(GDALGroupH hGroup, const char *pszName,
//                                     CSLConstList papszOptions);
// GDALAttributeH CPL_DLL GDALGroupCreateAttribute(
//     GDALGroupH hGroup, const char *pszName, size_t nDimensions,
//     const GUInt64 *panDimensions, GDALExtendedDataTypeH hEDT,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// bool CPL_DLL GDALGroupDeleteAttribute(GDALGroupH hGroup, const char *pszName,
//                                       CSLConstList papszOptions);
// bool CPL_DLL GDALGroupRename(GDALGroupH hGroup, const char *pszNewName);
// GDALGroupH CPL_DLL GDALGroupSubsetDimensionFromSelection(
//     GDALGroupH hGroup, const char *pszSelection, CSLConstList papszOptions);
// size_t CPL_DLL GDALGroupGetDataTypeCount(GDALGroupH hGroup);
// GDALExtendedDataTypeH CPL_DLL GDALGroupGetDataType(GDALGroupH hGroup,
//                                                    size_t nIdx);

// void CPL_DLL GDALMDArrayRelease(GDALMDArrayH hMDArray);
// const char CPL_DLL *GDALMDArrayGetName(GDALMDArrayH hArray);
// const char CPL_DLL *GDALMDArrayGetFullName(GDALMDArrayH hArray);
// GUInt64 CPL_DLL GDALMDArrayGetTotalElementsCount(GDALMDArrayH hArray);
// size_t CPL_DLL GDALMDArrayGetDimensionCount(GDALMDArrayH hArray);
// GDALDimensionH CPL_DLL *
// GDALMDArrayGetDimensions(GDALMDArrayH hArray,
//                          size_t *pnCount) CPL_WARN_UNUSED_RESULT;
// GDALExtendedDataTypeH CPL_DLL GDALMDArrayGetDataType(GDALMDArrayH hArray)
//     CPL_WARN_UNUSED_RESULT;
// int CPL_DLL GDALMDArrayRead(GDALMDArrayH hArray, const GUInt64 *arrayStartIdx,
//                             const size_t *count, const GInt64 *arrayStep,
//                             const GPtrDiff_t *bufferStride,
//                             GDALExtendedDataTypeH bufferDatatype,
//                             void *pDstBuffer, const void *pDstBufferAllocStart,
//                             size_t nDstBufferllocSize);
// int CPL_DLL GDALMDArrayWrite(GDALMDArrayH hArray, const GUInt64 *arrayStartIdx,
//                              const size_t *count, const GInt64 *arrayStep,
//                              const GPtrDiff_t *bufferStride,
//                              GDALExtendedDataTypeH bufferDatatype,
//                              const void *pSrcBuffer,
//                              const void *psrcBufferAllocStart,
//                              size_t nSrcBufferllocSize);
// int CPL_DLL GDALMDArrayAdviseRead(GDALMDArrayH hArray,
//                                   const GUInt64 *arrayStartIdx,
//                                   const size_t *count);
// int CPL_DLL GDALMDArrayAdviseReadEx(GDALMDArrayH hArray,
//                                     const GUInt64 *arrayStartIdx,
//                                     const size_t *count,
//                                     CSLConstList papszOptions);
// GDALAttributeH CPL_DLL GDALMDArrayGetAttribute(
//     GDALMDArrayH hArray, const char *pszName) CPL_WARN_UNUSED_RESULT;
// GDALAttributeH CPL_DLL *
// GDALMDArrayGetAttributes(GDALMDArrayH hArray, size_t *pnCount,
//                          CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// GDALAttributeH CPL_DLL GDALMDArrayCreateAttribute(
//     GDALMDArrayH hArray, const char *pszName, size_t nDimensions,
//     const GUInt64 *panDimensions, GDALExtendedDataTypeH hEDT,
//     CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;
// bool CPL_DLL GDALMDArrayDeleteAttribute(GDALMDArrayH hArray,
//                                         const char *pszName,
//                                         CSLConstList papszOptions);
// bool CPL_DLL GDALMDArrayResize(GDALMDArrayH hArray,
//                                const GUInt64 *panNewDimSizes,
//                                CSLConstList papszOptions);
// const void CPL_DLL *GDALMDArrayGetRawNoDataValue(GDALMDArrayH hArray);
// double CPL_DLL GDALMDArrayGetNoDataValueAsDouble(GDALMDArrayH hArray,
//                                                  int *pbHasNoDataValue);
// int64_t CPL_DLL GDALMDArrayGetNoDataValueAsInt64(GDALMDArrayH hArray,
//                                                  int *pbHasNoDataValue);
// uint64_t CPL_DLL GDALMDArrayGetNoDataValueAsUInt64(GDALMDArrayH hArray,
//                                                    int *pbHasNoDataValue);
// int CPL_DLL GDALMDArraySetRawNoDataValue(GDALMDArrayH hArray, const void *);
// int CPL_DLL GDALMDArraySetNoDataValueAsDouble(GDALMDArrayH hArray,
//                                               double dfNoDataValue);
// int CPL_DLL GDALMDArraySetNoDataValueAsInt64(GDALMDArrayH hArray,
//                                              int64_t nNoDataValue);
// int CPL_DLL GDALMDArraySetNoDataValueAsUInt64(GDALMDArrayH hArray,
//                                               uint64_t nNoDataValue);
// int CPL_DLL GDALMDArraySetScale(GDALMDArrayH hArray, double dfScale);
// int CPL_DLL GDALMDArraySetScaleEx(GDALMDArrayH hArray, double dfScale,
//                                   GDALDataType eStorageType);
// double CPL_DLL GDALMDArrayGetScale(GDALMDArrayH hArray, int *pbHasValue);
// double CPL_DLL GDALMDArrayGetScaleEx(GDALMDArrayH hArray, int *pbHasValue,
//                                      GDALDataType *peStorageType);
// int CPL_DLL GDALMDArraySetOffset(GDALMDArrayH hArray, double dfOffset);
// int CPL_DLL GDALMDArraySetOffsetEx(GDALMDArrayH hArray, double dfOffset,
//                                    GDALDataType eStorageType);
// double CPL_DLL GDALMDArrayGetOffset(GDALMDArrayH hArray, int *pbHasValue);
// double CPL_DLL GDALMDArrayGetOffsetEx(GDALMDArrayH hArray, int *pbHasValue,
//                                       GDALDataType *peStorageType);
// GUInt64 CPL_DLL *GDALMDArrayGetBlockSize(GDALMDArrayH hArray, size_t *pnCount);
// int CPL_DLL GDALMDArraySetUnit(GDALMDArrayH hArray, const char *);
// const char CPL_DLL *GDALMDArrayGetUnit(GDALMDArrayH hArray);
// int CPL_DLL GDALMDArraySetSpatialRef(GDALMDArrayH, OGRSpatialReferenceH);
// OGRSpatialReferenceH CPL_DLL GDALMDArrayGetSpatialRef(GDALMDArrayH hArray);
// size_t CPL_DLL *GDALMDArrayGetProcessingChunkSize(GDALMDArrayH hArray,
//                                                   size_t *pnCount,
//                                                   size_t nMaxChunkMemory);
// CSLConstList CPL_DLL GDALMDArrayGetStructuralInfo(GDALMDArrayH hArray);
// GDALMDArrayH CPL_DLL GDALMDArrayGetView(GDALMDArrayH hArray,
//                                         const char *pszViewExpr);
// GDALMDArrayH CPL_DLL GDALMDArrayTranspose(GDALMDArrayH hArray,
//                                           size_t nNewAxisCount,
//                                           const int *panMapNewAxisToOldAxis);
// GDALMDArrayH CPL_DLL GDALMDArrayGetUnscaled(GDALMDArrayH hArray);
// GDALMDArrayH CPL_DLL GDALMDArrayGetMask(GDALMDArrayH hArray,
//                                         CSLConstList papszOptions);
// GDALDatasetH CPL_DLL GDALMDArrayAsClassicDataset(GDALMDArrayH hArray,
//                                                  size_t iXDim, size_t iYDim);
// GDALDatasetH CPL_DLL GDALMDArrayAsClassicDatasetEx(GDALMDArrayH hArray,
//                                                    size_t iXDim, size_t iYDim,
//                                                    GDALGroupH hRootGroup,
//                                                    CSLConstList papszOptions);
// CPLErr CPL_DLL GDALMDArrayGetStatistics(
//     GDALMDArrayH hArray, GDALDatasetH, int bApproxOK, int bForce,
//     double *pdfMin, double *pdfMax, double *pdfMean, double *pdfStdDev,
//     GUInt64 *pnValidCount, GDALProgressFunc pfnProgress, void *pProgressData);
// int CPL_DLL GDALMDArrayComputeStatistics(GDALMDArrayH hArray, GDALDatasetH,
//                                          int bApproxOK, double *pdfMin,
//                                          double *pdfMax, double *pdfMean,
//                                          double *pdfStdDev,
//                                          GUInt64 *pnValidCount,
//                                          GDALProgressFunc, void *pProgressData);
// int CPL_DLL GDALMDArrayComputeStatisticsEx(
//     GDALMDArrayH hArray, GDALDatasetH, int bApproxOK, double *pdfMin,
//     double *pdfMax, double *pdfMean, double *pdfStdDev, GUInt64 *pnValidCount,
//     GDALProgressFunc, void *pProgressData, CSLConstList papszOptions);
// GDALMDArrayH CPL_DLL GDALMDArrayGetResampled(GDALMDArrayH hArray,
//                                              size_t nNewDimCount,
//                                              const GDALDimensionH *pahNewDims,
//                                              GDALRIOResampleAlg resampleAlg,
//                                              OGRSpatialReferenceH hTargetSRS,
//                                              CSLConstList papszOptions);
// GDALMDArrayH CPL_DLL GDALMDArrayGetGridded(
//     GDALMDArrayH hArray, const char *pszGridOptions, GDALMDArrayH hXArray,
//     GDALMDArrayH hYArray, CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// GDALMDArrayH CPL_DLL *
// GDALMDArrayGetCoordinateVariables(GDALMDArrayH hArray,
//                                   size_t *pnCount) CPL_WARN_UNUSED_RESULT;

// GDALMDArrayH CPL_DLL *
// GDALMDArrayGetMeshGrid(const GDALMDArrayH *pahInputArrays,
//                        size_t nCountInputArrays, size_t *pnCountOutputArrays,
//                        CSLConstList papszOptions) CPL_WARN_UNUSED_RESULT;

// #ifdef __cplusplus
// extern "C++"
// {
// #endif

//     /** Information on a raw block of a GDALMDArray
//      *
//      * @since 3.12
//      */
//     struct
// #ifdef __cplusplus
//         CPL_DLL
// #endif
//             GDALMDArrayRawBlockInfo
//     {
//         /** Filename into which the raw block is found */
//         char *pszFilename;
//         /** File offset within pszFilename of the start of the raw block */
//         uint64_t nOffset;
//         /** Size in bytes of the raw block */
//         uint64_t nSize;
//         /** NULL or Null-terminated list of driver specific information on the
//          * raw block */
//         char **papszInfo;
//         /** In-memory buffer of nSize bytes. When this is set, pszFilename and
//          * nOffset are set to NULL.
//          *
//          * When using C++ copy constructor or copy-assignment operator, if
//          * a memory allocation fails, a CPLError() will be emitted and this
//          * field will be NULL, but nSize not zero.
//          */
//         GByte *pabyInlineData;

// #ifdef __cplusplus
//         /*! @cond Doxygen_Suppress */
//         inline GDALMDArrayRawBlockInfo()
//             : pszFilename(nullptr), nOffset(0), nSize(0), papszInfo(nullptr),
//               pabyInlineData(nullptr)
//         {
//         }

//         ~GDALMDArrayRawBlockInfo();

//         void clear();

//         GDALMDArrayRawBlockInfo(const GDALMDArrayRawBlockInfo &);
//         GDALMDArrayRawBlockInfo &operator=(const GDALMDArrayRawBlockInfo &);
//         GDALMDArrayRawBlockInfo(GDALMDArrayRawBlockInfo &&);
//         GDALMDArrayRawBlockInfo &operator=(GDALMDArrayRawBlockInfo &&);
//         /*! @endcond */
// #endif
//     };

// #ifdef __cplusplus
// }
// #endif

// /*! @cond Doxygen_Suppress */
// typedef struct GDALMDArrayRawBlockInfo GDALMDArrayRawBlockInfo;
// /*! @endcond */

// GDALMDArrayRawBlockInfo CPL_DLL *GDALMDArrayRawBlockInfoCreate(void);
// void CPL_DLL
// GDALMDArrayRawBlockInfoRelease(GDALMDArrayRawBlockInfo *psBlockInfo);
// bool CPL_DLL GDALMDArrayGetRawBlockInfo(GDALMDArrayH hArray,
//                                         const uint64_t *panBlockCoordinates,
//                                         GDALMDArrayRawBlockInfo *psBlockInfo);

// void CPL_DLL GDALReleaseArrays(GDALMDArrayH *arrays, size_t nCount);
// int CPL_DLL GDALMDArrayCache(GDALMDArrayH hArray, CSLConstList papszOptions);
// bool CPL_DLL GDALMDArrayRename(GDALMDArrayH hArray, const char *pszNewName);

// GDALRasterAttributeTableH CPL_DLL GDALCreateRasterAttributeTableFromMDArrays(
//     GDALRATTableType eTableType, int nArrays, const GDALMDArrayH *ahArrays,
//     const GDALRATFieldUsage *paeUsages);

// void CPL_DLL GDALAttributeRelease(GDALAttributeH hAttr);
// void CPL_DLL GDALReleaseAttributes(GDALAttributeH *attributes, size_t nCount);
// const char CPL_DLL *GDALAttributeGetName(GDALAttributeH hAttr);
// const char CPL_DLL *GDALAttributeGetFullName(GDALAttributeH hAttr);
// GUInt64 CPL_DLL GDALAttributeGetTotalElementsCount(GDALAttributeH hAttr);
// size_t CPL_DLL GDALAttributeGetDimensionCount(GDALAttributeH hAttr);
// GUInt64 CPL_DLL *
// GDALAttributeGetDimensionsSize(GDALAttributeH hAttr,
//                                size_t *pnCount) CPL_WARN_UNUSED_RESULT;
// GDALExtendedDataTypeH CPL_DLL GDALAttributeGetDataType(GDALAttributeH hAttr)
//     CPL_WARN_UNUSED_RESULT;
// GByte CPL_DLL *GDALAttributeReadAsRaw(GDALAttributeH hAttr,
//                                       size_t *pnSize) CPL_WARN_UNUSED_RESULT;
// void CPL_DLL GDALAttributeFreeRawResult(GDALAttributeH hAttr, GByte *raw,
//                                         size_t nSize);
// const char CPL_DLL *GDALAttributeReadAsString(GDALAttributeH hAttr);
// int CPL_DLL GDALAttributeReadAsInt(GDALAttributeH hAttr);
// int64_t CPL_DLL GDALAttributeReadAsInt64(GDALAttributeH hAttr);
// double CPL_DLL GDALAttributeReadAsDouble(GDALAttributeH hAttr);
// char CPL_DLL **
// GDALAttributeReadAsStringArray(GDALAttributeH hAttr) CPL_WARN_UNUSED_RESULT;
// int CPL_DLL *GDALAttributeReadAsIntArray(GDALAttributeH hAttr, size_t *pnCount)
//     CPL_WARN_UNUSED_RESULT;
// int64_t CPL_DLL *
// GDALAttributeReadAsInt64Array(GDALAttributeH hAttr,
//                               size_t *pnCount) CPL_WARN_UNUSED_RESULT;
// double CPL_DLL *
// GDALAttributeReadAsDoubleArray(GDALAttributeH hAttr,
//                                size_t *pnCount) CPL_WARN_UNUSED_RESULT;
// int CPL_DLL GDALAttributeWriteRaw(GDALAttributeH hAttr, const void *, size_t);
// int CPL_DLL GDALAttributeWriteString(GDALAttributeH hAttr, const char *);
// int CPL_DLL GDALAttributeWriteStringArray(GDALAttributeH hAttr, CSLConstList);
// int CPL_DLL GDALAttributeWriteInt(GDALAttributeH hAttr, int);
// int CPL_DLL GDALAttributeWriteIntArray(GDALAttributeH hAttr, const int *,
//                                        size_t);
// int CPL_DLL GDALAttributeWriteInt64(GDALAttributeH hAttr, int64_t);
// int CPL_DLL GDALAttributeWriteInt64Array(GDALAttributeH hAttr, const int64_t *,
//                                          size_t);
// int CPL_DLL GDALAttributeWriteDouble(GDALAttributeH hAttr, double);
// int CPL_DLL GDALAttributeWriteDoubleArray(GDALAttributeH hAttr, const double *,
//                                           size_t);
// bool CPL_DLL GDALAttributeRename(GDALAttributeH hAttr, const char *pszNewName);

// void CPL_DLL GDALDimensionRelease(GDALDimensionH hDim);
// void CPL_DLL GDALReleaseDimensions(GDALDimensionH *dims, size_t nCount);
// const char CPL_DLL *GDALDimensionGetName(GDALDimensionH hDim);
// const char CPL_DLL *GDALDimensionGetFullName(GDALDimensionH hDim);
// const char CPL_DLL *GDALDimensionGetType(GDALDimensionH hDim);
// const char CPL_DLL *GDALDimensionGetDirection(GDALDimensionH hDim);
// GUInt64 CPL_DLL GDALDimensionGetSize(GDALDimensionH hDim);
// GDALMDArrayH CPL_DLL GDALDimensionGetIndexingVariable(GDALDimensionH hDim)
//     CPL_WARN_UNUSED_RESULT;
// int CPL_DLL GDALDimensionSetIndexingVariable(GDALDimensionH hDim,
//                                              GDALMDArrayH hArray);
// bool CPL_DLL GDALDimensionRename(GDALDimensionH hDim, const char *pszNewName);

// CPL_C_END
