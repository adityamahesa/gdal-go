package gdal

/*
#include "gdal_preamble.h"
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

func gdalGetDataTypeSizeBits(dataType GDALDataType) (result int) {
	result = int(C.GDALGetDataTypeSizeBits(C.GDALDataType(dataType)))
	return
}

func gdalGetDataTypeSizeBytes(dataType GDALDataType) (result int) {
	result = int(C.GDALGetDataTypeSizeBytes(C.GDALDataType(dataType)))
	return
}

func gdalDataTypeIsComplex(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsComplex(C.GDALDataType(dataType)) != 0
	return
}

func gdalDataTypeIsInteger(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsInteger(C.GDALDataType(dataType)) != 0
	return
}

func gdalDataTypeIsFloating(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsFloating(C.GDALDataType(dataType)) != 0
	return
}

func gdalDataTypeIsSigned(dataType GDALDataType) (result bool) {
	result = C.GDALDataTypeIsSigned(C.GDALDataType(dataType)) != 0
	return
}

func gdalGetDataTypeName(dataType GDALDataType) (result string) {
	result = C.GoString(C.GDALGetDataTypeName(C.GDALDataType(dataType)))
	return
}

func gdalGetDataTypeByName(name string) (result GDALDataType) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDataType(C.GDALGetDataTypeByName(cName))
	return
}

func gdalDataTypeUnion(a, b GDALDataType) (result GDALDataType) {
	result = GDALDataType(C.GDALDataTypeUnion(C.GDALDataType(a), C.GDALDataType(b)))
	return
}

func gdalDataTypeUnionWithValue(dataType GDALDataType, value float64, complex int) (result GDALDataType) {
	result = GDALDataType(C.GDALDataTypeUnionWithValue(C.GDALDataType(dataType), C.double(value), C.int(complex)))
	return
}

func gdalFindDataType(bits, signed, floating, complex int) (result GDALDataType) {
	result = GDALDataType(C.GDALFindDataType(C.int(bits), C.int(signed), C.int(floating), C.int(complex)))
	return
}

func gdalFindDataTypeForValue(value float64, complex int) (result GDALDataType) {
	result = GDALDataType(C.GDALFindDataTypeForValue(C.double(value), C.int(complex)))
	return
}

func gdalAdjustValueToDataType(dataType GDALDataType, value float64, clamped, rounded *int) float64 {
	var cClamped, cRounded C.int
	r := C.GDALAdjustValueToDataType(C.GDALDataType(dataType), C.double(value), &cClamped, &cRounded)
	*clamped = int(cClamped)
	*rounded = int(cRounded)
	return float64(r)
}

func gdalIsValueExactAs(value float64, dataType GDALDataType) (result bool) {
	result = bool(C.GDALIsValueExactAs(C.double(value), C.GDALDataType(dataType)))
	return
}

func gdalIsValueInRangeOf(value float64, dataType GDALDataType) (result bool) {
	result = bool(C.GDALIsValueInRangeOf(C.double(value), C.GDALDataType(dataType)))
	return
}

func gdalGetNonComplexDataType(dataType GDALDataType) (result GDALDataType) {
	result = GDALDataType(C.GDALGetNonComplexDataType(C.GDALDataType(dataType)))
	return
}

func gdalDataTypeIsConversionLossy(from, to GDALDataType) (result bool) {
	result = C.GDALDataTypeIsConversionLossy(C.GDALDataType(from), C.GDALDataType(to)) != 0
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

func gdalGetAsyncStatusTypeByName(name string) (result GDALAsyncStatusType) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALAsyncStatusType(C.GDALGetAsyncStatusTypeByName(cName))
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

func gdalGetColorInterpretationByName(name string) (result GDALColorInterp) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALColorInterp(C.GDALGetColorInterpretationByName(cName))
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

// "well known" metadata items.
var (
	GDALMDAreaOrPoint = C.GoString(C._GDALMD_AREA_OR_POINT())
	GDALMDAOPArea     = C.GoString(C._GDALMD_AOP_AREA())
	GDALMDAOPPoint    = C.GoString(C._GDALMD_AOP_POINT())
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
	GDALDMDLongName                                  = C.GoString(C._GDAL_DMD_LONGNAME())
	GDALDMDHelpTopic                                 = C.GoString(C._GDAL_DMD_HELPTOPIC())
	GDALDMDMimeType                                  = C.GoString(C._GDAL_DMD_MIMETYPE())
	GDALDMDExtension                                 = C.GoString(C._GDAL_DMD_EXTENSION())
	GDALDMDConnectionPrefix                          = C.GoString(C._GDAL_DMD_CONNECTION_PREFIX())
	GDALDMDExtensions                                = C.GoString(C._GDAL_DMD_EXTENSIONS())
	GDALDMDCreationOptionList                        = C.GoString(C._GDAL_DMD_CREATIONOPTIONLIST())
	GDALDMDOverviewCreationOptionList                = C.GoString(C._GDAL_DMD_OVERVIEW_CREATIONOPTIONLIST())
	GDALDMDMultidimDatasetCreationOptionList         = C.GoString(C._GDAL_DMD_MULTIDIM_DATASET_CREATIONOPTIONLIST())
	GDALDMDMultidimGroupCreationOptionList           = C.GoString(C._GDAL_DMD_MULTIDIM_GROUP_CREATIONOPTIONLIST())
	GDALDMDMultidimDimensionCreationOptionList       = C.GoString(C._GDAL_DMD_MULTIDIM_DIMENSION_CREATIONOPTIONLIST())
	GDALDMDMultidimArrayCreationOptionList           = C.GoString(C._GDAL_DMD_MULTIDIM_ARRAY_CREATIONOPTIONLIST())
	GDALDMDMultidimArrayOpenOptionList               = C.GoString(C._GDAL_DMD_MULTIDIM_ARRAY_OPENOPTIONLIST())
	GDALDMDMultidimAttributeCreationOptionList       = C.GoString(C._GDAL_DMD_MULTIDIM_ATTRIBUTE_CREATIONOPTIONLIST())
	GDALDMDOpenOptionList                            = C.GoString(C._GDAL_DMD_OPENOPTIONLIST())
	GDALDMDCreationDataTypes                         = C.GoString(C._GDAL_DMD_CREATIONDATATYPES())
	GDALDMDCreationFieldDataTypes                    = C.GoString(C._GDAL_DMD_CREATIONFIELDDATATYPES())
	GDALDMDCreationFieldDataSubTypes                 = C.GoString(C._GDAL_DMD_CREATIONFIELDDATASUBTYPES())
	GDALDMDMaxStringLength                           = C.GoString(C._GDAL_DMD_MAX_STRING_LENGTH())
	GDALDMDCreationFieldDefnFlags                    = C.GoString(C._GDAL_DMD_CREATION_FIELD_DEFN_FLAGS())
	GDALDMDSubdatasets                               = C.GoString(C._GDAL_DMD_SUBDATASETS())
	GDALDCAPCreateSubdatasets                        = C.GoString(C._GDAL_DCAP_CREATE_SUBDATASETS())
	GDALDMDNumericFieldWidthIncludesDecimalSeparator = C.GoString(C._GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_DECIMAL_SEPARATOR())
	GDALDMDNumericFieldWidthIncludesSign             = C.GoString(C._GDAL_DMD_NUMERIC_FIELD_WIDTH_INCLUDES_SIGN())
	GDALDCAPOpen                                     = C.GoString(C._GDAL_DCAP_OPEN())
	GDALDCAPCreate                                   = C.GoString(C._GDAL_DCAP_CREATE())
	GDALDCAPCreateMultidimensional                   = C.GoString(C._GDAL_DCAP_CREATE_MULTIDIMENSIONAL())
	GDALDCAPCreateCopy                               = C.GoString(C._GDAL_DCAP_CREATECOPY())
	GDALDCAPCreateOnlyVisibleAtCloseTime             = C.GoString(C._GDAL_DCAP_CREATE_ONLY_VISIBLE_AT_CLOSE_TIME())
	GDALDCAPVectorTranslateFrom                      = C.GoString(C._GDAL_DCAP_VECTOR_TRANSLATE_FROM())
	GDALDCAPCreateCopyMultidimensional               = C.GoString(C._GDAL_DCAP_CREATECOPY_MULTIDIMENSIONAL())
	GDALDCAPMultidimRaster                           = C.GoString(C._GDAL_DCAP_MULTIDIM_RASTER())
	GDALDCAPSubCreateCopy                            = C.GoString(C._GDAL_DCAP_SUBCREATECOPY())
	GDALDCAPAppend                                   = C.GoString(C._GDAL_DCAP_APPEND())
	GDALDCAPUpdate                                   = C.GoString(C._GDAL_DCAP_UPDATE())
	GDALDCAPVirtualIO                                = C.GoString(C._GDAL_DCAP_VIRTUALIO())
	GDALDCAPRaster                                   = C.GoString(C._GDAL_DCAP_RASTER())
	GDALDCAPVector                                   = C.GoString(C._GDAL_DCAP_VECTOR())
	GDALDCAPGNM                                      = C.GoString(C._GDAL_DCAP_GNM())
	GDALDCAPCreateLayer                              = C.GoString(C._GDAL_DCAP_CREATE_LAYER())
	GDALDCAPDeleteLayer                              = C.GoString(C._GDAL_DCAP_DELETE_LAYER())
	GDALDCAPCreateField                              = C.GoString(C._GDAL_DCAP_CREATE_FIELD())
	GDALDCAPDeleteField                              = C.GoString(C._GDAL_DCAP_DELETE_FIELD())
	GDALDCAPReorderFields                            = C.GoString(C._GDAL_DCAP_REORDER_FIELDS())
	GDALDMDAlterFieldDefnFlags                       = C.GoString(C._GDAL_DMD_ALTER_FIELD_DEFN_FLAGS())
	GDALDMDIllegalFieldNames                         = C.GoString(C._GDAL_DMD_ILLEGAL_FIELD_NAMES())
	GDALDCAPNotNullFields                            = C.GoString(C._GDAL_DCAP_NOTNULL_FIELDS())
	GDALDCAPUniqueFields                             = C.GoString(C._GDAL_DCAP_UNIQUE_FIELDS())
	GDALDCAPDefaultFields                            = C.GoString(C._GDAL_DCAP_DEFAULT_FIELDS())
	GDALDCAPNotNullGeomFields                        = C.GoString(C._GDAL_DCAP_NOTNULL_GEOMFIELDS())
	GDALDCAPNonspatial                               = C.GoString(C._GDAL_DCAP_NONSPATIAL())
	GDALDCAPCurveGeometries                          = C.GoString(C._GDAL_DCAP_CURVE_GEOMETRIES())
	GDALDCAPMeasuredGeometries                       = C.GoString(C._GDAL_DCAP_MEASURED_GEOMETRIES())
	GDALDCAPZGeometries                              = C.GoString(C._GDAL_DCAP_Z_GEOMETRIES())
	GDALDMDGeometryFlags                             = C.GoString(C._GDAL_DMD_GEOMETRY_FLAGS())
	GDALDCAPFeatureStyles                            = C.GoString(C._GDAL_DCAP_FEATURE_STYLES())
	GDALDCAPFeatureStylesRead                        = C.GoString(C._GDAL_DCAP_FEATURE_STYLES_READ())
	GDALDCAPFeatureStylesWrite                       = C.GoString(C._GDAL_DCAP_FEATURE_STYLES_WRITE())
	GDALDCAPCoordinateEpoch                          = C.GoString(C._GDAL_DCAP_COORDINATE_EPOCH())
	GDALDCAPMultipleVectorLayers                     = C.GoString(C._GDAL_DCAP_MULTIPLE_VECTOR_LAYERS())
	GDALDCAPFieldDomains                             = C.GoString(C._GDAL_DCAP_FIELD_DOMAINS())
	GDALDCAPRelationships                            = C.GoString(C._GDAL_DCAP_RELATIONSHIPS())
	GDALDCAPCreateRelationship                       = C.GoString(C._GDAL_DCAP_CREATE_RELATIONSHIP())
	GDALDCAPDeleteRelationship                       = C.GoString(C._GDAL_DCAP_DELETE_RELATIONSHIP())
	GDALDCAPUpdateRelationship                       = C.GoString(C._GDAL_DCAP_UPDATE_RELATIONSHIP())
	GDALDCAPFlushCacheConsistentState                = C.GoString(C._GDAL_DCAP_FLUSHCACHE_CONSISTENT_STATE())
	GDALDCAPHonorGeomCoordinatePrecision             = C.GoString(C._GDAL_DCAP_HONOR_GEOM_COORDINATE_PRECISION())
	GDALDCAPUpsert                                   = C.GoString(C._GDAL_DCAP_UPSERT())
	GDALDMDRelationshipFlags                         = C.GoString(C._GDAL_DMD_RELATIONSHIP_FLAGS())
	GDALDMDRelationshipRelatedTableTypes             = C.GoString(C._GDAL_DMD_RELATIONSHIP_RELATED_TABLE_TYPES())
	GDALDCAPRenameLayers                             = C.GoString(C._GDAL_DCAP_RENAME_LAYERS())
	GDALDMDCreationFieldDomainTypes                  = C.GoString(C._GDAL_DMD_CREATION_FIELD_DOMAIN_TYPES())
	GDALDMDAlterGeomFieldDefnFlags                   = C.GoString(C._GDAL_DMD_ALTER_GEOM_FIELD_DEFN_FLAGS())
	GDALDMDSupportedSQLDialects                      = C.GoString(C._GDAL_DMD_SUPPORTED_SQL_DIALECTS())
	GDALDMDPluginInstallationMessage                 = C.GoString(C._GDAL_DMD_PLUGIN_INSTALLATION_MESSAGE())
	GDALDMDUpdateItems                               = C.GoString(C._GDAL_DMD_UPDATE_ITEMS())
	GDALDimTypeHorizontalX                           = C.GoString(C._GDAL_DIM_TYPE_HORIZONTAL_X())
	GDALDimTypeHorizontalY                           = C.GoString(C._GDAL_DIM_TYPE_HORIZONTAL_Y())
	GDALDimTypeVertical                              = C.GoString(C._GDAL_DIM_TYPE_VERTICAL())
	GDALDimTypeTemporal                              = C.GoString(C._GDAL_DIM_TYPE_TEMPORAL())
	GDALDimTypeParametric                            = C.GoString(C._GDAL_DIM_TYPE_PARAMETRIC())
	GDALDCAPReopenAfterWriteRequired                 = C.GoString(C._GDAL_DCAP_REOPEN_AFTER_WRITE_REQUIRED())
	GDALDCAPCanReadAfterDelete                       = C.GoString(C._GDAL_DCAP_CAN_READ_AFTER_DELETE())
	GDsCAddRelationship                              = C.GoString(C._GDsCAddRelationship())
	GDsCDeleteRelationship                           = C.GoString(C._GDsCDeleteRelationship())
	GDsCUpdateRelationship                           = C.GoString(C._GDsCUpdateRelationship())
	GDsCFastGetExtent                                = C.GoString(C._GDsCFastGetExtent())
	GDsCFastGetExtentWGS84LongLat                    = C.GoString(C._GDsCFastGetExtentWGS84LongLat())
)

func gdalAllRegister() {
	C.GDALAllRegister()
}

func gdalRegisterPlugins() {
	C.GDALRegisterPlugins()
}

func gdalRegisterPlugin(name string) (result CPLErr) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = CPLErr(C.GDALRegisterPlugin(cName))
	return
}

func gdalCreate(driver GDALDriver, name string, xSize, ySize, bands int, dataType GDALDataType, options CSLConstList) (result GDALDataset) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = GDALDataset{cValue: C.GDALCreate(driver.cValue, cName, C.int(xSize), C.int(ySize), C.int(bands), C.GDALDataType(dataType), opts)}
	return
}

func gdalCreateCopy(driver GDALDriver, name string, src GDALDataset, strict int, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result GDALDataset) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = GDALDataset{cValue: C.GDALCreateCopy(driver.cValue, cName, src.cValue, C.int(strict), opts, progress.cValue, progressData)}
	return
}

func gdalIdentifyDriver(filename string, fileList CSLConstList) (result GDALDriver) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	files := fileList.cValue
	result = GDALDriver{cValue: C.GDALIdentifyDriver(cName, files)}
	return
}

func gdalIdentifyDriverEx(filename string, identifyFlags uint, allowedDrivers, fileList CSLConstList) (result GDALDriver) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	allowed := allowedDrivers.cValue
	files := fileList.cValue
	result = GDALDriver{cValue: C.GDALIdentifyDriverEx(cName, C.uint(identifyFlags), allowed, files)}
	return
}

func gdalOpen(filename string, access GDALAccess) (result GDALDataset) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDataset{cValue: C.GDALOpen(cName, C.GDALAccess(access))}
	return
}

func gdalOpenShared(filename string, access GDALAccess) (result GDALDataset) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDataset{cValue: C.GDALOpenShared(cName, C.GDALAccess(access))}
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

func gdalOpenEx(filename string, openFlags GDALOpenFlag, allowedDrivers, openOptions, siblingFiles CSLConstList) (result GDALDataset) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	allowed := allowedDrivers.cValue
	options := openOptions.cValue
	siblings := siblingFiles.cValue
	result = GDALDataset{cValue: C.GDALOpenEx(cName, C.uint(openFlags), allowed, options, siblings)}
	return
}

func gdalDumpOpenDatasets(filename string) (result int, err error) {
	fp, closeFn, err := cFOpen(filename, "w")
	if err != nil {
		return
	}
	defer closeFn()
	result = int(C.GDALDumpOpenDatasets(fp))
	return
}

func gdalGetDriverByName(name string) (result GDALDriver) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALDriver{cValue: C.GDALGetDriverByName(cName)}
	return
}

func gdalGetDriverCount() (result int) {
	result = int(C.GDALGetDriverCount())
	return
}

func gdalGetDriver(index int) (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALGetDriver(C.int(index))}
	return
}

func gdalCreateDriver() (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALCreateDriver()}
	return
}

func gdalDestroyDriver(driver GDALDriver) {
	C.GDALDestroyDriver(driver.cValue)
}

func gdalRegisterDriver(driver GDALDriver) (result int) {
	result = int(C.GDALRegisterDriver(driver.cValue))
	return
}

func gdalDeregisterDriver(driver GDALDriver) {
	C.GDALDeregisterDriver(driver.cValue)
}

func gdalDestroyDriverManager() {
	C.GDALDestroyDriverManager()
}

func gdalDestroy() {
	C.GDALDestroy()
}

func gdalDeleteDataset(driver GDALDriver, filename string) (result CPLErr) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	result = CPLErr(C.GDALDeleteDataset(driver.cValue, cName))
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

func gdalCopyDatasetFiles(driver GDALDriver, newName, oldName string) (result CPLErr) {
	cNew := C.CString(newName)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOld))
	result = CPLErr(C.GDALCopyDatasetFiles(driver.cValue, cNew, cOld))
	return
}

func gdalValidateCreationOptions(driver GDALDriver, options CSLConstList) (result bool) {
	opts := options.cValue
	result = C.GDALValidateCreationOptions(driver.cValue, opts) != 0
	return
}

func gdalGetOutputDriversForDatasetName(destFilename string, flagRasterVector int, singleMatch, emitWarning bool) (result CSLConstList) {
	cName := C.CString(destFilename)
	defer C.free(unsafe.Pointer(cName))
	raw := C.GDALGetOutputDriversForDatasetName(cName, C.int(flagRasterVector), C.bool(singleMatch), C.bool(emitWarning))
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalDriverHasOpenOption(driver GDALDriver, openOptionName string) (result bool) {
	cName := C.CString(openOptionName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALDriverHasOpenOption(driver.cValue, cName))
	return
}

// The following are deprecated.

func gdalGetDriverShortName(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverShortName(driver.cValue))
	return
}

func gdalGetDriverLongName(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverLongName(driver.cValue))
	return
}

func gdalGetDriverHelpTopic(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverHelpTopic(driver.cValue))
	return
}

func gdalGetDriverCreationOptionList(driver GDALDriver) (result string) {
	result = C.GoString(C.GDALGetDriverCreationOptionList(driver.cValue))
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

func gdalDeinitGCPs(count int, gcps GDALGCPs) {
	C.GDALDeinitGCPs(C.int(count), gcps.cPtr())
}

func gdalDuplicateGCPs(count int, gcps GDALGCPs) (result GDALGCPs) {
	dup := C.GDALDuplicateGCPs(C.int(count), gcps.cPtr())
	if dup == nil {
		return
	}
	result = unsafe.Slice((*GDALGCP)(unsafe.Pointer(dup)), count)
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

func gdalGCPsToHomography(count int, gcps GDALGCPs, homography *[9]float64) int {
	var h [9]C.double
	r := C.GDALGCPsToHomography(C.int(count), gcps.cPtr(), &h[0])
	for i := range h {
		homography[i] = float64(h[i])
	}
	return int(r)
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

// /* ==================================================================== */
// /*      major objects (dataset, and, driver, drivermanager).            */
// /* ==================================================================== */

func gdalGetMetadataDomainList(object GDALMajorObject) (result CSLConstList) {
	raw := C.GDALGetMetadataDomainList(object.cValue)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalGetMetadata(object GDALMajorObject, domain string) (result CSLConstList) {
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	result = cslConstList(C.GDALGetMetadata(object.cValue, cDomain))
	return
}

func gdalSetMetadata(object GDALMajorObject, metadata CSLConstList, domain string) (result CPLErr) {
	md := metadata.cValue
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	result = CPLErr(C.GDALSetMetadata(object.cValue, md, cDomain))
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

func gdalGetDescription(object GDALMajorObject) (result string) {
	result = C.GoString(C.GDALGetDescription(object.cValue))
	return
}

func gdalSetDescription(object GDALMajorObject, description string) {
	cDesc := C.CString(description)
	defer C.free(unsafe.Pointer(cDesc))
	C.GDALSetDescription(object.cValue, cDesc)
}

// /* ==================================================================== */
// /*      GDALDataset class ... normally this represents one file.        */
// /* ==================================================================== */

// Name of driver metadata item for layer creation option list.
var GDALDSLayerCreationOptionList = C.GoString(C._GDAL_DS_LAYER_CREATIONOPTIONLIST())

func gdalGetDatasetDriver(dataset GDALDataset) (result GDALDriver) {
	result = GDALDriver{cValue: C.GDALGetDatasetDriver(dataset.cValue)}
	return
}

func gdalGetFileList(dataset GDALDataset) (result CSLConstList) {
	raw := C.GDALGetFileList(dataset.cValue)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalDatasetMarkSuppressOnClose(dataset GDALDataset) {
	C.GDALDatasetMarkSuppressOnClose(dataset.cValue)
}

func gdalClose(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALClose(dataset.cValue))
	return
}

func gdalDatasetRunCloseWithoutDestroying(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALDatasetRunCloseWithoutDestroying(dataset.cValue))
	return
}

func gdalGetRasterXSize(dataset GDALDataset) (result int) {
	result = int(C.GDALGetRasterXSize(dataset.cValue))
	return
}

func gdalGetRasterYSize(dataset GDALDataset) (result int) {
	result = int(C.GDALGetRasterYSize(dataset.cValue))
	return
}

func gdalGetRasterCount(dataset GDALDataset) (result int) {
	result = int(C.GDALGetRasterCount(dataset.cValue))
	return
}

func gdalGetRasterBand(dataset GDALDataset, band int) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetRasterBand(dataset.cValue, C.int(band))}
	return
}

func gdalDatasetIsThreadSafe(dataset GDALDataset, scopeFlags int, options CSLConstList) (result bool) {
	opts := options.cValue
	result = bool(C.GDALDatasetIsThreadSafe(dataset.cValue, C.int(scopeFlags), opts))
	return
}

func gdalGetThreadSafeDataset(dataset GDALDataset, scopeFlags int, options CSLConstList) (result GDALDataset) {
	opts := options.cValue
	result = GDALDataset{cValue: C.GDALGetThreadSafeDataset(dataset.cValue, C.int(scopeFlags), opts)}
	return
}

func gdalAddBand(dataset GDALDataset, dataType GDALDataType, options CSLConstList) (result CPLErr) {
	opts := options.cValue
	result = CPLErr(C.GDALAddBand(dataset.cValue, C.GDALDataType(dataType), opts))
	return
}

func gdalBeginAsyncReader(dataset GDALDataset, xOff, yOff, xSize, ySize int, buf unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandMap []int, pixelSpace, lineSpace, bandSpace int, options CSLConstList) (result GDALAsyncReader) {
	opts := options.cValue
	result = GDALAsyncReader{cValue: C.GDALBeginAsyncReader(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buf, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandMap), C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace), opts)}
	return
}

func gdalEndAsyncReader(dataset GDALDataset, reader GDALAsyncReader) {
	C.GDALEndAsyncReader(dataset.cValue, reader.cValue)
}

func gdalDatasetRasterIO(dataset GDALDataset, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int) (result CPLErr) {
	result = CPLErr(C.GDALDatasetRasterIO(dataset.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandList), C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace)))
	return
}

func gdalDatasetRasterIOEx(dataset GDALDataset, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, pixelSpace, lineSpace, bandSpace int64, extraArg GDALRasterIOExtraArg) (result CPLErr) {
	result = CPLErr(C.GDALDatasetRasterIOEx(dataset.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandList), C.GSpacing(pixelSpace), C.GSpacing(lineSpace), C.GSpacing(bandSpace), extraArg.cValue))
	return
}

func gdalDatasetAdviseRead(dataset GDALDataset, xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, bandCount int, bandList []int, options CSLConstList) (result CPLErr) {
	opts := options.cValue
	result = CPLErr(C.GDALDatasetAdviseRead(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(bandCount), cInts(bandList), opts))
	return
}

func gdalDatasetGetCompressionFormats(dataset GDALDataset, xOff, yOff, xSize, ySize, bandCount int, bandList []int) (result CSLConstList) {
	raw := C.GDALDatasetGetCompressionFormats(dataset.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bandCount), cInts(bandList))
	if raw == nil {
		return
	}
	result = cslConstList(raw)
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
		vsiFree(unsafe.Pointer(cDetailed))
	}
	return
}

func gdalGetProjectionRef(dataset GDALDataset) (result string) {
	result = C.GoString(C.GDALGetProjectionRef(dataset.cValue))
	return
}

func gdalGetSpatialRef(dataset GDALDataset) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.GDALGetSpatialRef(dataset.cValue)}
	return
}

func gdalSetProjection(dataset GDALDataset, projection string) (result CPLErr) {
	cProj := C.CString(projection)
	defer C.free(unsafe.Pointer(cProj))
	result = CPLErr(C.GDALSetProjection(dataset.cValue, cProj))
	return
}

func gdalSetSpatialRef(dataset GDALDataset, srs OGRSpatialReference) (result CPLErr) {
	result = CPLErr(C.GDALSetSpatialRef(dataset.cValue, srs.cValue))
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

func gdalSetGeoTransform(dataset GDALDataset, geoTransform [6]float64) (result CPLErr) {
	var gt [6]C.double
	for i, v := range geoTransform {
		gt[i] = C.double(v)
	}
	result = CPLErr(C.GDALSetGeoTransform(dataset.cValue, &gt[0]))
	return
}

func gdalGetExtent(dataset GDALDataset, envelope OGREnvelope, crs OGRSpatialReference) (result CPLErr) {
	result = CPLErr(C.GDALGetExtent(dataset.cValue, envelope.cValue, crs.cValue))
	return
}

func gdalGetExtentWGS84LongLat(dataset GDALDataset, envelope OGREnvelope) (result CPLErr) {
	result = CPLErr(C.GDALGetExtentWGS84LongLat(dataset.cValue, envelope.cValue))
	return
}

func gdalDatasetGeolocationToPixelLine(dataset GDALDataset, geolocX, geolocY float64, srs OGRSpatialReference, pixel, line *float64, transformerOptions CSLConstList) (result CPLErr) {
	opts := transformerOptions.cValue
	var cPixel, cLine C.double
	result = CPLErr(C.GDALDatasetGeolocationToPixelLine(dataset.cValue, C.double(geolocX), C.double(geolocY), srs.cValue, &cPixel, &cLine, opts))
	*pixel = float64(cPixel)
	*line = float64(cLine)
	return
}

func gdalGetGCPCount(dataset GDALDataset) (result int) {
	result = int(C.GDALGetGCPCount(dataset.cValue))
	return
}

func gdalGetGCPProjection(dataset GDALDataset) (result string) {
	result = C.GoString(C.GDALGetGCPProjection(dataset.cValue))
	return
}

func gdalGetGCPSpatialRef(dataset GDALDataset) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.GDALGetGCPSpatialRef(dataset.cValue)}
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

func gdalSetGCPs(dataset GDALDataset, count int, gcps GDALGCPs, projection string) (result CPLErr) {
	cProj := C.CString(projection)
	defer C.free(unsafe.Pointer(cProj))
	result = CPLErr(C.GDALSetGCPs(dataset.cValue, C.int(count), gcps.cPtr(), cProj))
	return
}

func gdalSetGCPs2(dataset GDALDataset, count int, gcps GDALGCPs, srs OGRSpatialReference) (result CPLErr) {
	result = CPLErr(C.GDALSetGCPs2(dataset.cValue, C.int(count), gcps.cPtr(), srs.cValue))
	return
}

func gdalGetInternalHandle(dataset GDALDataset, request string) unsafe.Pointer {
	cRequest := C.CString(request)
	defer C.free(unsafe.Pointer(cRequest))
	return C.GDALGetInternalHandle(dataset.cValue, cRequest)
}

func gdalReferenceDataset(dataset GDALDataset) (result int) {
	result = int(C.GDALReferenceDataset(dataset.cValue))
	return
}

func gdalDereferenceDataset(dataset GDALDataset) (result int) {
	result = int(C.GDALDereferenceDataset(dataset.cValue))
	return
}

func gdalReleaseDataset(dataset GDALDataset) (result int) {
	result = int(C.GDALReleaseDataset(dataset.cValue))
	return
}

func gdalBuildOverviews(dataset GDALDataset, resampling string, nOverviews int, overviewList []int, bandCount int, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	result = CPLErr(C.GDALBuildOverviews(dataset.cValue, cResampling, C.int(nOverviews), cInts(overviewList), C.int(bandCount), cInts(bandList), progress.cValue, progressData))
	return
}

func gdalBuildOverviewsEx(dataset GDALDataset, resampling string, nOverviews int, overviewList []int, bandCount int, bandList []int, progress GDALProgressFunc, progressData unsafe.Pointer, options CSLConstList) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	opts := options.cValue
	result = CPLErr(C.GDALBuildOverviewsEx(dataset.cValue, cResampling, C.int(nOverviews), cInts(overviewList), C.int(bandCount), cInts(bandList), progress.cValue, progressData, opts))
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

func gdalGetAccess(dataset GDALDataset) (result int) {
	result = int(C.GDALGetAccess(dataset.cValue))
	return
}

func gdalFlushCache(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALFlushCache(dataset.cValue))
	return
}

func gdalDropCache(dataset GDALDataset) (result CPLErr) {
	result = CPLErr(C.GDALDropCache(dataset.cValue))
	return
}

func gdalCreateDatasetMaskBand(dataset GDALDataset, flags int) (result CPLErr) {
	result = CPLErr(C.GDALCreateDatasetMaskBand(dataset.cValue, C.int(flags)))
	return
}

func gdalDatasetCopyWholeRaster(srcDataset, dstDataset GDALDataset, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	opts := options.cValue
	result = CPLErr(C.GDALDatasetCopyWholeRaster(srcDataset.cValue, dstDataset.cValue, opts, progress.cValue, progressData))
	return
}

func gdalRasterBandCopyWholeRaster(srcBand, dstBand GDALRasterBand, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	opts := options.cValue
	result = CPLErr(C.GDALRasterBandCopyWholeRaster(srcBand.cValue, dstBand.cValue, opts, progress.cValue, progressData))
	return
}

func gdalRegenerateOverviews(srcBand GDALRasterBand, overviewCount int, overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	result = CPLErr(C.GDALRegenerateOverviews(srcBand.cValue, C.int(overviewCount), overviewBands.cPtr(), cResampling, progress.cValue, progressData))
	return
}

func gdalRegenerateOverviewsEx(srcBand GDALRasterBand, overviewCount int, overviewBands GDALRasterBands, resampling string, progress GDALProgressFunc, progressData unsafe.Pointer, options CSLConstList) (result CPLErr) {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))
	opts := options.cValue
	result = CPLErr(C.GDALRegenerateOverviewsEx(srcBand.cValue, C.int(overviewCount), overviewBands.cPtr(), cResampling, progress.cValue, progressData, opts))
	return
}

func gdalDatasetGetLayerCount(dataset GDALDataset) (result int) {
	result = int(C.GDALDatasetGetLayerCount(dataset.cValue))
	return
}

func gdalDatasetGetLayer(dataset GDALDataset, index int) (result OGRLayer) {
	result = OGRLayer{cValue: C.GDALDatasetGetLayer(dataset.cValue, C.int(index))}
	return
}

// OGR_L_GetDataset is defined here to avoid a circular dependency with ogr_api.h.
func ogrLGetDataset(layer OGRLayer) (result GDALDataset) {
	result = GDALDataset{cValue: C.OGR_L_GetDataset(layer.cValue)}
	return
}

func gdalDatasetGetLayerByName(dataset GDALDataset, name string) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = OGRLayer{cValue: C.GDALDatasetGetLayerByName(dataset.cValue, cName)}
	return
}

func gdalDatasetIsLayerPrivate(dataset GDALDataset, index int) (result bool) {
	result = C.GDALDatasetIsLayerPrivate(dataset.cValue, C.int(index)) != 0
	return
}

func gdalDatasetDeleteLayer(dataset GDALDataset, index int) (result OGRErr) {
	result = OGRErr(C.GDALDatasetDeleteLayer(dataset.cValue, C.int(index)))
	return
}

func gdalDatasetCreateLayer(dataset GDALDataset, name string, srs OGRSpatialReference, geomType OGRwkbGeometryType, options CSLConstList) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = OGRLayer{cValue: C.GDALDatasetCreateLayer(dataset.cValue, cName, srs.cValue, C.OGRwkbGeometryType(geomType), opts)}
	return
}

func gdalDatasetCreateLayerFromGeomFieldDefn(dataset GDALDataset, name string, geomFieldDefn OGRGeomFieldDefn, options CSLConstList) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = OGRLayer{cValue: C.GDALDatasetCreateLayerFromGeomFieldDefn(dataset.cValue, cName, geomFieldDefn.cValue, opts)}
	return
}

func gdalDatasetCopyLayer(dataset GDALDataset, srcLayer OGRLayer, newName string, options CSLConstList) (result OGRLayer) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = OGRLayer{cValue: C.GDALDatasetCopyLayer(dataset.cValue, srcLayer.cValue, cName, opts)}
	return
}

func gdalDatasetResetReading(dataset GDALDataset) {
	C.GDALDatasetResetReading(dataset.cValue)
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

func gdalDatasetTestCapability(dataset GDALDataset, capability string) (result bool) {
	cCapability := C.CString(capability)
	defer C.free(unsafe.Pointer(cCapability))
	result = C.GDALDatasetTestCapability(dataset.cValue, cCapability) != 0
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

func gdalDatasetAbortSQL(dataset GDALDataset) (result OGRErr) {
	result = OGRErr(C.GDALDatasetAbortSQL(dataset.cValue))
	return
}

func gdalDatasetReleaseResultSet(dataset GDALDataset, layer OGRLayer) {
	C.GDALDatasetReleaseResultSet(dataset.cValue, layer.cValue)
}

func gdalDatasetGetStyleTable(dataset GDALDataset) (result OGRStyleTable) {
	result = OGRStyleTable{cValue: C.GDALDatasetGetStyleTable(dataset.cValue)}
	return
}

func gdalDatasetSetStyleTableDirectly(dataset GDALDataset, styleTable OGRStyleTable) {
	C.GDALDatasetSetStyleTableDirectly(dataset.cValue, styleTable.cValue)
}

func gdalDatasetSetStyleTable(dataset GDALDataset, styleTable OGRStyleTable) {
	C.GDALDatasetSetStyleTable(dataset.cValue, styleTable.cValue)
}

func gdalDatasetStartTransaction(dataset GDALDataset, force int) (result OGRErr) {
	result = OGRErr(C.GDALDatasetStartTransaction(dataset.cValue, C.int(force)))
	return
}

func gdalDatasetCommitTransaction(dataset GDALDataset) (result OGRErr) {
	result = OGRErr(C.GDALDatasetCommitTransaction(dataset.cValue))
	return
}

func gdalDatasetRollbackTransaction(dataset GDALDataset) (result OGRErr) {
	result = OGRErr(C.GDALDatasetRollbackTransaction(dataset.cValue))
	return
}

func gdalDatasetClearStatistics(dataset GDALDataset) {
	C.GDALDatasetClearStatistics(dataset.cValue)
}

func gdalDatasetAsMDArray(dataset GDALDataset, options CSLConstList) (result GDALMDArray) {
	opts := options.cValue
	result = GDALMDArray{cValue: C.GDALDatasetAsMDArray(dataset.cValue, opts)}
	return
}

func gdalDatasetGetFieldDomainNames(dataset GDALDataset, options CSLConstList) (result CSLConstList) {
	opts := options.cValue
	raw := C.GDALDatasetGetFieldDomainNames(dataset.cValue, opts)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalDatasetGetFieldDomain(dataset GDALDataset, name string) (result OGRFieldDomain) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = OGRFieldDomain{cValue: C.GDALDatasetGetFieldDomain(dataset.cValue, cName)}
	return
}

func gdalDatasetAddFieldDomain(dataset GDALDataset, fieldDomain OGRFieldDomain, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetAddFieldDomain(dataset.cValue, fieldDomain.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		vsiFree(unsafe.Pointer(cReason))
	}
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
		vsiFree(unsafe.Pointer(cReason))
	}
	return
}

func gdalDatasetUpdateFieldDomain(dataset GDALDataset, fieldDomain OGRFieldDomain, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetUpdateFieldDomain(dataset.cValue, fieldDomain.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		vsiFree(unsafe.Pointer(cReason))
	}
	return
}

func gdalDatasetGetRelationshipNames(dataset GDALDataset, options CSLConstList) (result CSLConstList) {
	opts := options.cValue
	raw := C.GDALDatasetGetRelationshipNames(dataset.cValue, opts)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalDatasetGetRelationship(dataset GDALDataset, name string) (result GDALRelationship) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALRelationship{cValue: C.GDALDatasetGetRelationship(dataset.cValue, cName)}
	return
}

func gdalDatasetAddRelationship(dataset GDALDataset, relationship GDALRelationship, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetAddRelationship(dataset.cValue, relationship.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		vsiFree(unsafe.Pointer(cReason))
	}
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
		vsiFree(unsafe.Pointer(cReason))
	}
	return
}

func gdalDatasetUpdateRelationship(dataset GDALDataset, relationship GDALRelationship, failureReason *string) (result bool) {
	var cReason *C.char
	result = bool(C.GDALDatasetUpdateRelationship(dataset.cValue, relationship.cValue, &cReason))
	if cReason != nil {
		if failureReason != nil {
			*failureReason = C.GoString(cReason)
		}
		vsiFree(unsafe.Pointer(cReason))
	}
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
		vsiFree(unsafe.Pointer(cResult))
	}
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
		vsiFree(unsafe.Pointer(cResult))
	}
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
		vsiFree(unsafe.Pointer(cResult))
	}
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

func gdalGetBlockSize(band GDALRasterBand, xSize, ySize *int) {
	var cXSize, cYSize C.int
	C.GDALGetBlockSize(band.cValue, &cXSize, &cYSize)
	*xSize = int(cXSize)
	*ySize = int(cYSize)
}

func gdalGetActualBlockSize(band GDALRasterBand, xBlockOff, yBlockOff int, xValid, yValid *int) (result CPLErr) {
	var cXValid, cYValid C.int
	result = CPLErr(C.GDALGetActualBlockSize(band.cValue, C.int(xBlockOff), C.int(yBlockOff), &cXValid, &cYValid))
	*xValid = int(cXValid)
	*yValid = int(cYValid)
	return
}

func gdalRasterAdviseRead(band GDALRasterBand, xOff, yOff, xSize, ySize, bufXSize, bufYSize int, bufType GDALDataType, options CSLConstList) (result CPLErr) {
	opts := options.cValue
	result = CPLErr(C.GDALRasterAdviseRead(band.cValue, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), opts))
	return
}

func gdalRasterIO(band GDALRasterBand, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int) (result CPLErr) {
	result = CPLErr(C.GDALRasterIO(band.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.int(pixelSpace), C.int(lineSpace)))
	return
}

func gdalRasterIOEx(band GDALRasterBand, rwFlag GDALRWFlag, xOff, yOff, xSize, ySize int, buffer unsafe.Pointer, bufXSize, bufYSize int, bufType GDALDataType, pixelSpace, lineSpace int64, extraArg GDALRasterIOExtraArg) (result CPLErr) {
	result = CPLErr(C.GDALRasterIOEx(band.cValue, C.GDALRWFlag(rwFlag), C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), buffer, C.int(bufXSize), C.int(bufYSize), C.GDALDataType(bufType), C.GSpacing(pixelSpace), C.GSpacing(lineSpace), extraArg.cValue))
	return
}

func gdalReadBlock(band GDALRasterBand, xBlockOff, yBlockOff int, buffer unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.GDALReadBlock(band.cValue, C.int(xBlockOff), C.int(yBlockOff), buffer))
	return
}

func gdalWriteBlock(band GDALRasterBand, xBlockOff, yBlockOff int, buffer unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.GDALWriteBlock(band.cValue, C.int(xBlockOff), C.int(yBlockOff), buffer))
	return
}

func gdalGetRasterBandXSize(band GDALRasterBand) (result int) {
	result = int(C.GDALGetRasterBandXSize(band.cValue))
	return
}

func gdalGetRasterBandYSize(band GDALRasterBand) (result int) {
	result = int(C.GDALGetRasterBandYSize(band.cValue))
	return
}

func gdalGetRasterAccess(band GDALRasterBand) (result GDALAccess) {
	result = GDALAccess(C.GDALGetRasterAccess(band.cValue))
	return
}

func gdalGetBandNumber(band GDALRasterBand) (result int) {
	result = int(C.GDALGetBandNumber(band.cValue))
	return
}

func gdalGetBandDataset(band GDALRasterBand) (result GDALDataset) {
	result = GDALDataset{cValue: C.GDALGetBandDataset(band.cValue)}
	return
}

func gdalGetRasterColorInterpretation(band GDALRasterBand) (result GDALColorInterp) {
	result = GDALColorInterp(C.GDALGetRasterColorInterpretation(band.cValue))
	return
}

func gdalSetRasterColorInterpretation(band GDALRasterBand, colorInterp GDALColorInterp) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterColorInterpretation(band.cValue, C.GDALColorInterp(colorInterp)))
	return
}

func gdalGetRasterColorTable(band GDALRasterBand) (result GDALColorTable) {
	result = GDALColorTable{cValue: C.GDALGetRasterColorTable(band.cValue)}
	return
}

func gdalSetRasterColorTable(band GDALRasterBand, colorTable GDALColorTable) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterColorTable(band.cValue, colorTable.cValue))
	return
}

func gdalHasArbitraryOverviews(band GDALRasterBand) (result bool) {
	result = C.GDALHasArbitraryOverviews(band.cValue) != 0
	return
}

func gdalGetOverviewCount(band GDALRasterBand) (result int) {
	result = int(C.GDALGetOverviewCount(band.cValue))
	return
}

func gdalGetOverview(band GDALRasterBand, index int) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetOverview(band.cValue, C.int(index))}
	return
}

func gdalGetRasterNoDataValue(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterNoDataValue(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func gdalGetRasterNoDataValueAsInt64(band GDALRasterBand, success *int) (result int64) {
	var cSuccess C.int
	result = int64(C.GDALGetRasterNoDataValueAsInt64(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func gdalGetRasterNoDataValueAsUInt64(band GDALRasterBand, success *int) (result uint64) {
	var cSuccess C.int
	result = uint64(C.GDALGetRasterNoDataValueAsUInt64(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func gdalSetRasterNoDataValue(band GDALRasterBand, value float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterNoDataValue(band.cValue, C.double(value)))
	return
}

func gdalSetRasterNoDataValueAsInt64(band GDALRasterBand, value int64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterNoDataValueAsInt64(band.cValue, C.int64_t(value)))
	return
}

func gdalSetRasterNoDataValueAsUInt64(band GDALRasterBand, value uint64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterNoDataValueAsUInt64(band.cValue, C.uint64_t(value)))
	return
}

func gdalDeleteRasterNoDataValue(band GDALRasterBand) (result CPLErr) {
	result = CPLErr(C.GDALDeleteRasterNoDataValue(band.cValue))
	return
}

func gdalGetRasterCategoryNames(band GDALRasterBand) (result CSLConstList) {
	result = cslConstList(C.GDALGetRasterCategoryNames(band.cValue))
	return
}

func gdalSetRasterCategoryNames(band GDALRasterBand, names CSLConstList) (result CPLErr) {
	n := names.cValue
	result = CPLErr(C.GDALSetRasterCategoryNames(band.cValue, n))
	return
}

func gdalGetRasterMinimum(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterMinimum(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func gdalGetRasterMaximum(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterMaximum(band.cValue, &cSuccess))
	*success = int(cSuccess)
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

func gdalComputeRasterStatistics(band GDALRasterBand, approxOK int, min, max, mean, stdDev *float64, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	var cMin, cMax, cMean, cStdDev C.double
	result = CPLErr(C.GDALComputeRasterStatistics(band.cValue, C.int(approxOK), &cMin, &cMax, &cMean, &cStdDev, progress.cValue, progressData))
	*min = float64(cMin)
	*max = float64(cMax)
	*mean = float64(cMean)
	*stdDev = float64(cStdDev)
	return
}

func gdalSetRasterStatistics(band GDALRasterBand, min, max, mean, stdDev float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterStatistics(band.cValue, C.double(min), C.double(max), C.double(mean), C.double(stdDev)))
	return
}

func gdalRasterBandAsMDArray(band GDALRasterBand) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALRasterBandAsMDArray(band.cValue)}
	return
}

func gdalGetRasterUnitType(band GDALRasterBand) (result string) {
	result = C.GoString(C.GDALGetRasterUnitType(band.cValue))
	return
}

func gdalSetRasterUnitType(band GDALRasterBand, newValue string) (result CPLErr) {
	cValue := C.CString(newValue)
	defer C.free(unsafe.Pointer(cValue))
	result = CPLErr(C.GDALSetRasterUnitType(band.cValue, cValue))
	return
}

func gdalGetRasterOffset(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterOffset(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func gdalSetRasterOffset(band GDALRasterBand, newOffset float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterOffset(band.cValue, C.double(newOffset)))
	return
}

func gdalGetRasterScale(band GDALRasterBand, success *int) (result float64) {
	var cSuccess C.int
	result = float64(C.GDALGetRasterScale(band.cValue, &cSuccess))
	*success = int(cSuccess)
	return
}

func gdalSetRasterScale(band GDALRasterBand, newScale float64) (result CPLErr) {
	result = CPLErr(C.GDALSetRasterScale(band.cValue, C.double(newScale)))
	return
}

func gdalComputeRasterMinMax(band GDALRasterBand, approxOK int, minMax *[2]float64) (result CPLErr) {
	var cMinMax [2]C.double
	result = CPLErr(C.GDALComputeRasterMinMax(band.cValue, C.int(approxOK), &cMinMax[0]))
	minMax[0] = float64(cMinMax[0])
	minMax[1] = float64(cMinMax[1])
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

func gdalFlushRasterCache(band GDALRasterBand) (result CPLErr) {
	result = CPLErr(C.GDALFlushRasterCache(band.cValue))
	return
}

func gdalDropRasterCache(band GDALRasterBand) (result CPLErr) {
	result = CPLErr(C.GDALDropRasterCache(band.cValue))
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
		vsiFree(unsafe.Pointer(cHist))
	}
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

func gdalGetRandomRasterSample(band GDALRasterBand, samples int, buffer []float32) (result int) {
	var ptr *C.float
	if len(buffer) > 0 {
		ptr = (*C.float)(unsafe.Pointer(&buffer[0]))
	}
	result = int(C.GDALGetRandomRasterSample(band.cValue, C.int(samples), ptr))
	return
}

func gdalGetRasterSampleOverview(band GDALRasterBand, desiredSamples int) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetRasterSampleOverview(band.cValue, C.int(desiredSamples))}
	return
}

func gdalGetRasterSampleOverviewEx(band GDALRasterBand, desiredSamples uint64) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetRasterSampleOverviewEx(band.cValue, C.GUIntBig(desiredSamples))}
	return
}

func gdalFillRaster(band GDALRasterBand, realValue, imaginaryValue float64) (result CPLErr) {
	result = CPLErr(C.GDALFillRaster(band.cValue, C.double(realValue), C.double(imaginaryValue)))
	return
}

func gdalComputeBandStats(band GDALRasterBand, sampleStep int, mean, stdDev *float64, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	var cMean, cStdDev C.double
	result = CPLErr(C.GDALComputeBandStats(band.cValue, C.int(sampleStep), &cMean, &cStdDev, progress.cValue, progressData))
	*mean = float64(cMean)
	*stdDev = float64(cStdDev)
	return
}

func gdalOverviewMagnitudeCorrection(baseBand GDALRasterBand, overviewCount int, overviews GDALRasterBands, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.GDALOverviewMagnitudeCorrection(baseBand.cValue, C.int(overviewCount), overviews.cPtr(), progress.cValue, progressData))
	return
}

func gdalGetDefaultRAT(band GDALRasterBand) (result GDALRasterAttributeTable) {
	result = GDALRasterAttributeTable{cValue: C.GDALGetDefaultRAT(band.cValue)}
	return
}

func gdalSetDefaultRAT(band GDALRasterBand, rat GDALRasterAttributeTable) (result CPLErr) {
	result = CPLErr(C.GDALSetDefaultRAT(band.cValue, rat.cValue))
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

func gdalRasterInterpolateAtGeolocation(band GDALRasterBand, geolocX, geolocY float64, srs OGRSpatialReference, interpolation GDALRIOResampleAlg, realValue, imagValue *float64, transformerOptions CSLConstList) (result CPLErr) {
	opts := transformerOptions.cValue
	var cReal, cImag C.double
	result = CPLErr(C.GDALRasterInterpolateAtGeolocation(band.cValue, C.double(geolocX), C.double(geolocY), srs.cValue, C.GDALRIOResampleAlg(interpolation), &cReal, &cImag, opts))
	*realValue = float64(cReal)
	*imagValue = float64(cImag)
	return
}

// The VRTProcessedDataset function API (VRTPDWorkingDataPtr, the Init/Free/
// Process callback typedefs and GDALVRTRegisterProcessedDatasetFunc) is
// deferred: it registers Go callbacks as C function pointers (//export).

func gdalGetMaskBand(band GDALRasterBand) (result GDALRasterBand) {
	result = GDALRasterBand{cValue: C.GDALGetMaskBand(band.cValue)}
	return
}

func gdalGetMaskFlags(band GDALRasterBand) (result int) {
	result = int(C.GDALGetMaskFlags(band.cValue))
	return
}

func gdalCreateMaskBand(band GDALRasterBand, flags int) (result CPLErr) {
	result = CPLErr(C.GDALCreateMaskBand(band.cValue, C.int(flags)))
	return
}

func gdalIsMaskBand(band GDALRasterBand) (result bool) {
	result = bool(C.GDALIsMaskBand(band.cValue))
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

func gdalComputedRasterBandRelease(band GDALComputedRasterBand) {
	C.GDALComputedRasterBandRelease(band.cValue)
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

func gdalRasterBandBinaryOpDouble(band GDALRasterBand, op GDALRasterAlgebraBinaryOperation, constant float64) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandBinaryOpDouble(band.cValue, C.GDALRasterAlgebraBinaryOperation(op), C.double(constant))}
	return
}

func gdalRasterBandBinaryOpDoubleToBand(constant float64, op GDALRasterAlgebraBinaryOperation, band GDALRasterBand) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandBinaryOpDoubleToBand(C.double(constant), C.GDALRasterAlgebraBinaryOperation(op), band.cValue)}
	return
}

func gdalRasterBandIfThenElse(condBand, thenBand, elseBand GDALRasterBand) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandIfThenElse(condBand.cValue, thenBand.cValue, elseBand.cValue)}
	return
}

func gdalRasterBandAsDataType(band GDALRasterBand, dataType GDALDataType) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandAsDataType(band.cValue, C.GDALDataType(dataType))}
	return
}

func gdalMaximumOfNBands(bandCount int, bands GDALRasterBands) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALMaximumOfNBands(C.size_t(bandCount), bands.cPtr())}
	return
}

func gdalRasterBandMaxConstant(band GDALRasterBand, constant float64) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandMaxConstant(band.cValue, C.double(constant))}
	return
}

func gdalMinimumOfNBands(bandCount int, bands GDALRasterBands) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALMinimumOfNBands(C.size_t(bandCount), bands.cPtr())}
	return
}

func gdalRasterBandMinConstant(band GDALRasterBand, constant float64) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALRasterBandMinConstant(band.cValue, C.double(constant))}
	return
}

func gdalMeanOfNBands(bandCount int, bands GDALRasterBands) (result GDALComputedRasterBand) {
	result = GDALComputedRasterBand{cValue: C.GDALMeanOfNBands(C.size_t(bandCount), bands.cPtr())}
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

func gdalARLockBuffer(reader GDALAsyncReader, timeout float64) (result bool) {
	result = C.GDALARLockBuffer(reader.cValue, C.double(timeout)) != 0
	return
}

func gdalARUnlockBuffer(reader GDALAsyncReader) {
	C.GDALARUnlockBuffer(reader.cValue)
}

// Helper functions.

// GDALGeneralCmdLineProcessor is deferred: its char*** in/out argv needs a
// dedicated design.

func gdalSwapWords(data unsafe.Pointer, wordSize, wordCount, wordSkip int) {
	C.GDALSwapWords(data, C.int(wordSize), C.int(wordCount), C.int(wordSkip))
}

func gdalSwapWordsEx(data unsafe.Pointer, wordSize int, wordCount int, wordSkip int) {
	C.GDALSwapWordsEx(data, C.int(wordSize), C.size_t(wordCount), C.int(wordSkip))
}

func gdalCopyWords(src unsafe.Pointer, srcType GDALDataType, srcPixelOffset int, dst unsafe.Pointer, dstType GDALDataType, dstPixelOffset, wordCount int) {
	C.GDALCopyWords(src, C.GDALDataType(srcType), C.int(srcPixelOffset), dst, C.GDALDataType(dstType), C.int(dstPixelOffset), C.int(wordCount))
}

func gdalCopyWords64(src unsafe.Pointer, srcType GDALDataType, srcPixelOffset int, dst unsafe.Pointer, dstType GDALDataType, dstPixelOffset int, wordCount int64) {
	C.GDALCopyWords64(src, C.GDALDataType(srcType), C.int(srcPixelOffset), dst, C.GDALDataType(dstType), C.int(dstPixelOffset), C.GPtrDiff_t(wordCount))
}

func gdalCopyBits(src unsafe.Pointer, srcOffset, srcStep int, dst unsafe.Pointer, dstOffset, dstStep, bitCount, stepCount int) {
	C.GDALCopyBits((*C.GByte)(src), C.int(srcOffset), C.int(srcStep), (*C.GByte)(dst), C.int(dstOffset), C.int(dstStep), C.int(bitCount), C.int(stepCount))
}

// GDALDeinterleave is deferred: its void **ppDestBuffer output array of buffers
// needs a dedicated design.

func gdalTranspose2D(src unsafe.Pointer, srcType GDALDataType, dst unsafe.Pointer, dstType GDALDataType, srcWidth, srcHeight int) {
	C.GDALTranspose2D(src, C.GDALDataType(srcType), dst, C.GDALDataType(dstType), C.size_t(srcWidth), C.size_t(srcHeight))
}

func gdalGetNoDataReplacementValue(dataType GDALDataType, value float64) (result float64) {
	result = float64(C.GDALGetNoDataReplacementValue(C.GDALDataType(dataType), C.double(value)))
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

// GDALLoadTabFile, GDALReadTabFile, GDALLoadOziMapFile and GDALReadOziMapFile
// are deferred: their char**/int*/GDAL_GCP** in/out parameters need a dedicated
// design.

func gdalDecToDMS(angle float64, axis string, precision int) (result string) {
	cAxis := C.CString(axis)
	defer C.free(unsafe.Pointer(cAxis))
	result = C.GoString(C.GDALDecToDMS(C.double(angle), cAxis, C.int(precision)))
	return
}

func gdalPackedDMSToDec(packed float64) (result float64) {
	result = float64(C.GDALPackedDMSToDec(C.double(packed)))
	return
}

func gdalDecToPackedDMS(dec float64) (result float64) {
	result = float64(C.GDALDecToPackedDMS(C.double(dec)))
	return
}

func gdalVersionInfo(request string) (result string) {
	cRequest := C.CString(request)
	defer C.free(unsafe.Pointer(cRequest))
	result = C.GoString(C.GDALVersionInfo(cRequest))
	return
}

func gdalCheckVersion(versionMajor, versionMinor int, callingComponentName string) (result bool) {
	cName := C.CString(callingComponentName)
	defer C.free(unsafe.Pointer(cName))
	result = C.GDALCheckVersion(C.int(versionMajor), C.int(versionMinor), cName) != 0
	return
}

// GDALRPCInfoV2 stores Rational Polynomial Coefficients / Rigorous Projection
// Model.
type GDALRPCInfoV2 struct {
	cValue C.GDALRPCInfoV2
}

// Note: gdal.h also declares the deprecated GDALExtractRPCInfoV1 /
// GDALRPCInfoV1 pair, but GDAL does not export a GDALExtractRPCInfoV1 symbol to
// external consumers — the V1 implementation is only reachable through the
// bare, macro-remapped GDALExtractRPCInfo name. Wrapping it would produce an
// undefined-symbol link error, so it is intentionally omitted; use V2.

func gdalExtractRPCInfoV2(metadata CSLConstList, rpcInfo *GDALRPCInfoV2) (result int) {
	md := metadata.cValue
	result = int(C.GDALExtractRPCInfoV2(md, &rpcInfo.cValue))
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

func gdalDestroyColorTable(colorTable GDALColorTable) {
	C.GDALDestroyColorTable(colorTable.cValue)
}

func gdalCloneColorTable(colorTable GDALColorTable) (result GDALColorTable) {
	result = GDALColorTable{cValue: C.GDALCloneColorTable(colorTable.cValue)}
	return
}

func gdalGetPaletteInterpretation(colorTable GDALColorTable) (result GDALPaletteInterp) {
	result = GDALPaletteInterp(C.GDALGetPaletteInterpretation(colorTable.cValue))
	return
}

func gdalGetColorEntryCount(colorTable GDALColorTable) (result int) {
	result = int(C.GDALGetColorEntryCount(colorTable.cValue))
	return
}

func gdalGetColorEntry(colorTable GDALColorTable, index int) (result GDALColorEntry) {
	entry := C.GDALGetColorEntry(colorTable.cValue, C.int(index))
	if entry != nil {
		result.cValue = *entry
	}
	return
}

func gdalGetColorEntryAsRGB(colorTable GDALColorTable, index int, result *GDALColorEntry) (ret int) {
	ret = int(C.GDALGetColorEntryAsRGB(colorTable.cValue, C.int(index), &result.cValue))
	return
}

func gdalSetColorEntry(colorTable GDALColorTable, index int, entry GDALColorEntry) {
	C.GDALSetColorEntry(colorTable.cValue, C.int(index), &entry.cValue)
}

func gdalCreateColorRamp(colorTable GDALColorTable, startIndex int, startColor GDALColorEntry, endIndex int, endColor GDALColorEntry) {
	C.GDALCreateColorRamp(colorTable.cValue, C.int(startIndex), &startColor.cValue, C.int(endIndex), &endColor.cValue)
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

func gdalDestroyRasterAttributeTable(rat GDALRasterAttributeTable) {
	C.GDALDestroyRasterAttributeTable(rat.cValue)
}

func gdalRATGetColumnCount(rat GDALRasterAttributeTable) (result int) {
	result = int(C.GDALRATGetColumnCount(rat.cValue))
	return
}

func gdalRATGetNameOfCol(rat GDALRasterAttributeTable, col int) (result string) {
	result = C.GoString(C.GDALRATGetNameOfCol(rat.cValue, C.int(col)))
	return
}

func gdalRATGetUsageOfCol(rat GDALRasterAttributeTable, col int) (result GDALRATFieldUsage) {
	result = GDALRATFieldUsage(C.GDALRATGetUsageOfCol(rat.cValue, C.int(col)))
	return
}

func gdalRATGetTypeOfCol(rat GDALRasterAttributeTable, col int) (result GDALRATFieldType) {
	result = GDALRATFieldType(C.GDALRATGetTypeOfCol(rat.cValue, C.int(col)))
	return
}

func gdalGetRATFieldTypeName(fieldType GDALRATFieldType) (result string) {
	result = C.GoString(C.GDALGetRATFieldTypeName(C.GDALRATFieldType(fieldType)))
	return
}

func gdalGetRATFieldUsageName(usage GDALRATFieldUsage) (result string) {
	result = C.GoString(C.GDALGetRATFieldUsageName(C.GDALRATFieldUsage(usage)))
	return
}

func gdalRATGetColOfUsage(rat GDALRasterAttributeTable, usage GDALRATFieldUsage) (result int) {
	result = int(C.GDALRATGetColOfUsage(rat.cValue, C.GDALRATFieldUsage(usage)))
	return
}

func gdalRATGetRowCount(rat GDALRasterAttributeTable) (result int) {
	result = int(C.GDALRATGetRowCount(rat.cValue))
	return
}

func gdalRATGetValueAsString(rat GDALRasterAttributeTable, row, field int) (result string) {
	result = C.GoString(C.GDALRATGetValueAsString(rat.cValue, C.int(row), C.int(field)))
	return
}

func gdalRATGetValueAsInt(rat GDALRasterAttributeTable, row, field int) (result int) {
	result = int(C.GDALRATGetValueAsInt(rat.cValue, C.int(row), C.int(field)))
	return
}

func gdalRATGetValueAsDouble(rat GDALRasterAttributeTable, row, field int) (result float64) {
	result = float64(C.GDALRATGetValueAsDouble(rat.cValue, C.int(row), C.int(field)))
	return
}

func gdalRATGetValueAsBoolean(rat GDALRasterAttributeTable, row, field int) (result bool) {
	result = bool(C.GDALRATGetValueAsBoolean(rat.cValue, C.int(row), C.int(field)))
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

func gdalRATGetValueAsWKBGeometry(rat GDALRasterAttributeTable, row, field int) (result []byte) {
	var cSize C.size_t
	ptr := C.GDALRATGetValueAsWKBGeometry(rat.cValue, C.int(row), C.int(field), &cSize)
	if ptr != nil {
		result = C.GoBytes(unsafe.Pointer(ptr), C.int(cSize))
	}
	return
}

func gdalRATSetValueAsString(rat GDALRasterAttributeTable, row, field int, value string) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.GDALRATSetValueAsString(rat.cValue, C.int(row), C.int(field), cValue)
}

func gdalRATSetValueAsInt(rat GDALRasterAttributeTable, row, field, value int) {
	C.GDALRATSetValueAsInt(rat.cValue, C.int(row), C.int(field), C.int(value))
}

func gdalRATSetValueAsDouble(rat GDALRasterAttributeTable, row, field int, value float64) {
	C.GDALRATSetValueAsDouble(rat.cValue, C.int(row), C.int(field), C.double(value))
}

func gdalRATSetValueAsBoolean(rat GDALRasterAttributeTable, row, field int, value bool) (result CPLErr) {
	result = CPLErr(C.GDALRATSetValueAsBoolean(rat.cValue, C.int(row), C.int(field), C.bool(value)))
	return
}

func gdalRATSetValueAsDateTime(rat GDALRasterAttributeTable, row, field int, dateTime GDALRATDateTime) (result CPLErr) {
	result = CPLErr(C.GDALRATSetValueAsDateTime(rat.cValue, C.int(row), C.int(field), &dateTime.cValue))
	return
}

func gdalRATSetValueAsWKBGeometry(rat GDALRasterAttributeTable, row, field int, wkb unsafe.Pointer, wkbSize int) (result CPLErr) {
	result = CPLErr(C.GDALRATSetValueAsWKBGeometry(rat.cValue, C.int(row), C.int(field), wkb, C.size_t(wkbSize)))
	return
}

func gdalRATChangesAreWrittenToFile(rat GDALRasterAttributeTable) (result bool) {
	result = C.GDALRATChangesAreWrittenToFile(rat.cValue) != 0
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

// GDALRATValuesIOAsString, GDALRATValuesIOAsBoolean, GDALRATValuesIOAsDateTime
// and GDALRATValuesIOAsWKBGeometry are deferred: their in/out char**, bool*,
// DateTime array and GByte** parameters need a dedicated design.

func gdalRATSetRowCount(rat GDALRasterAttributeTable, count int) {
	C.GDALRATSetRowCount(rat.cValue, C.int(count))
}

func gdalRATCreateColumn(rat GDALRasterAttributeTable, name string, fieldType GDALRATFieldType, fieldUsage GDALRATFieldUsage) (result CPLErr) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = CPLErr(C.GDALRATCreateColumn(rat.cValue, cName, C.GDALRATFieldType(fieldType), C.GDALRATFieldUsage(fieldUsage)))
	return
}

func gdalRATSetLinearBinning(rat GDALRasterAttributeTable, row0Min, binSize float64) (result CPLErr) {
	result = CPLErr(C.GDALRATSetLinearBinning(rat.cValue, C.double(row0Min), C.double(binSize)))
	return
}

func gdalRATGetLinearBinning(rat GDALRasterAttributeTable, row0Min, binSize *float64) (result int) {
	var cRow0Min, cBinSize C.double
	result = int(C.GDALRATGetLinearBinning(rat.cValue, &cRow0Min, &cBinSize))
	*row0Min = float64(cRow0Min)
	*binSize = float64(cBinSize)
	return
}

func gdalRATSetTableType(rat GDALRasterAttributeTable, tableType GDALRATTableType) (result CPLErr) {
	result = CPLErr(C.GDALRATSetTableType(rat.cValue, C.GDALRATTableType(tableType)))
	return
}

func gdalRATGetTableType(rat GDALRasterAttributeTable) (result GDALRATTableType) {
	result = GDALRATTableType(C.GDALRATGetTableType(rat.cValue))
	return
}

func gdalRATInitializeFromColorTable(rat GDALRasterAttributeTable, colorTable GDALColorTable) (result CPLErr) {
	result = CPLErr(C.GDALRATInitializeFromColorTable(rat.cValue, colorTable.cValue))
	return
}

func gdalRATTranslateToColorTable(rat GDALRasterAttributeTable, entryCount int) (result GDALColorTable) {
	result = GDALColorTable{cValue: C.GDALRATTranslateToColorTable(rat.cValue, C.int(entryCount))}
	return
}

func gdalRATDumpReadable(rat GDALRasterAttributeTable, filename string) (err error) {
	fp, closeFn, err := cFOpen(filename, "w")
	if err != nil {
		return
	}
	defer closeFn()
	C.GDALRATDumpReadable(rat.cValue, fp)
	return
}

func gdalRATClone(rat GDALRasterAttributeTable) (result GDALRasterAttributeTable) {
	result = GDALRasterAttributeTable{cValue: C.GDALRATClone(rat.cValue)}
	return
}

// GDALRATSerializeJSON is deferred: its void* return (a json-c object) needs a
// dedicated design.

func gdalRATGetRowOfValue(rat GDALRasterAttributeTable, value float64) (result int) {
	result = int(C.GDALRATGetRowOfValue(rat.cValue, C.double(value)))
	return
}

func gdalRATRemoveStatistics(rat GDALRasterAttributeTable) {
	C.GDALRATRemoveStatistics(rat.cValue)
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

func gdalDestroyRelationship(relationship GDALRelationship) {
	C.GDALDestroyRelationship(relationship.cValue)
}

func gdalRelationshipGetName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetName(relationship.cValue))
	return
}

func gdalRelationshipGetCardinality(relationship GDALRelationship) (result GDALRelationshipCardinality) {
	result = GDALRelationshipCardinality(C.GDALRelationshipGetCardinality(relationship.cValue))
	return
}

func gdalRelationshipGetLeftTableName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetLeftTableName(relationship.cValue))
	return
}

func gdalRelationshipGetRightTableName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetRightTableName(relationship.cValue))
	return
}

func gdalRelationshipGetMappingTableName(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetMappingTableName(relationship.cValue))
	return
}

func gdalRelationshipSetMappingTableName(relationship GDALRelationship, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.GDALRelationshipSetMappingTableName(relationship.cValue, cName)
}

func gdalRelationshipGetLeftTableFields(relationship GDALRelationship) (result CSLConstList) {
	raw := C.GDALRelationshipGetLeftTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalRelationshipGetRightTableFields(relationship GDALRelationship) (result CSLConstList) {
	raw := C.GDALRelationshipGetRightTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalRelationshipSetLeftTableFields(relationship GDALRelationship, fields CSLConstList) {
	f := fields.cValue
	C.GDALRelationshipSetLeftTableFields(relationship.cValue, f)
}

func gdalRelationshipSetRightTableFields(relationship GDALRelationship, fields CSLConstList) {
	f := fields.cValue
	C.GDALRelationshipSetRightTableFields(relationship.cValue, f)
}

func gdalRelationshipGetLeftMappingTableFields(relationship GDALRelationship) (result CSLConstList) {
	raw := C.GDALRelationshipGetLeftMappingTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalRelationshipGetRightMappingTableFields(relationship GDALRelationship) (result CSLConstList) {
	raw := C.GDALRelationshipGetRightMappingTableFields(relationship.cValue)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalRelationshipSetLeftMappingTableFields(relationship GDALRelationship, fields CSLConstList) {
	f := fields.cValue
	C.GDALRelationshipSetLeftMappingTableFields(relationship.cValue, f)
}

func gdalRelationshipSetRightMappingTableFields(relationship GDALRelationship, fields CSLConstList) {
	f := fields.cValue
	C.GDALRelationshipSetRightMappingTableFields(relationship.cValue, f)
}

func gdalRelationshipGetType(relationship GDALRelationship) (result GDALRelationshipType) {
	result = GDALRelationshipType(C.GDALRelationshipGetType(relationship.cValue))
	return
}

func gdalRelationshipSetType(relationship GDALRelationship, relationshipType GDALRelationshipType) {
	C.GDALRelationshipSetType(relationship.cValue, C.GDALRelationshipType(relationshipType))
}

func gdalRelationshipGetForwardPathLabel(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetForwardPathLabel(relationship.cValue))
	return
}

func gdalRelationshipSetForwardPathLabel(relationship GDALRelationship, label string) {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.GDALRelationshipSetForwardPathLabel(relationship.cValue, cLabel)
}

func gdalRelationshipGetBackwardPathLabel(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetBackwardPathLabel(relationship.cValue))
	return
}

func gdalRelationshipSetBackwardPathLabel(relationship GDALRelationship, label string) {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.GDALRelationshipSetBackwardPathLabel(relationship.cValue, cLabel)
}

func gdalRelationshipGetRelatedTableType(relationship GDALRelationship) (result string) {
	result = C.GoString(C.GDALRelationshipGetRelatedTableType(relationship.cValue))
	return
}

func gdalRelationshipSetRelatedTableType(relationship GDALRelationship, relatedTableType string) {
	cType := C.CString(relatedTableType)
	defer C.free(unsafe.Pointer(cType))
	C.GDALRelationshipSetRelatedTableType(relationship.cValue, cType)
}

// /* ==================================================================== */
// /*      GDAL Cache Management                                           */
// /* ==================================================================== */

func gdalSetCacheMax(bytes int) {
	C.GDALSetCacheMax(C.int(bytes))
}

func gdalGetCacheMax() (result int) {
	result = int(C.GDALGetCacheMax())
	return
}

func gdalGetCacheUsed() (result int) {
	result = int(C.GDALGetCacheUsed())
	return
}

func gdalSetCacheMax64(bytes int64) {
	C.GDALSetCacheMax64(C.GIntBig(bytes))
}

func gdalGetCacheMax64() (result int64) {
	result = int64(C.GDALGetCacheMax64())
	return
}

func gdalGetCacheUsed64() (result int64) {
	result = int64(C.GDALGetCacheUsed64())
	return
}

func gdalFlushCacheBlock() (result bool) {
	result = C.GDALFlushCacheBlock() != 0
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

func gdalGetJPEG2000Structure(filename string, options CSLConstList) (result CPLXMLNode) {
	cName := C.CString(filename)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = CPLXMLNode{cValue: C.GDALGetJPEG2000Structure(cName, opts)}
	return
}

// /* ==================================================================== */
// /*      Multidimensional API_api                                       */
// /* ==================================================================== */

func gdalCreateMultiDimensional(driver GDALDriver, name string, rootGroupOptions, options CSLConstList) (result GDALDataset) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	rootOpts := rootGroupOptions.cValue
	opts := options.cValue
	result = GDALDataset{cValue: C.GDALCreateMultiDimensional(driver.cValue, cName, rootOpts, opts)}
	return
}

func gdalExtendedDataTypeCreate(dataType GDALDataType) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALExtendedDataTypeCreate(C.GDALDataType(dataType))}
	return
}

func gdalExtendedDataTypeCreateString(maxStringLength int) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALExtendedDataTypeCreateString(C.size_t(maxStringLength))}
	return
}

func gdalExtendedDataTypeCreateStringEx(maxStringLength int, subType GDALExtendedDataTypeSubType) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALExtendedDataTypeCreateStringEx(C.size_t(maxStringLength), C.GDALExtendedDataTypeSubType(subType))}
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

func gdalExtendedDataTypeRelease(edt GDALExtendedDataType) {
	C.GDALExtendedDataTypeRelease(edt.cValue)
}

func gdalExtendedDataTypeGetName(edt GDALExtendedDataType) (result string) {
	result = C.GoString(C.GDALExtendedDataTypeGetName(edt.cValue))
	return
}

func gdalExtendedDataTypeGetClass(edt GDALExtendedDataType) (result GDALExtendedDataTypeClass) {
	result = GDALExtendedDataTypeClass(C.GDALExtendedDataTypeGetClass(edt.cValue))
	return
}

func gdalExtendedDataTypeGetNumericDataType(edt GDALExtendedDataType) (result GDALDataType) {
	result = GDALDataType(C.GDALExtendedDataTypeGetNumericDataType(edt.cValue))
	return
}

func gdalExtendedDataTypeGetSize(edt GDALExtendedDataType) (result int) {
	result = int(C.GDALExtendedDataTypeGetSize(edt.cValue))
	return
}

func gdalExtendedDataTypeGetMaxStringLength(edt GDALExtendedDataType) (result int) {
	result = int(C.GDALExtendedDataTypeGetMaxStringLength(edt.cValue))
	return
}

// GDALExtendedDataTypeGetComponents and GDALExtendedDataTypeFreeComponents are
// deferred: the returned GDALEDTComponentH* array has release/free ownership
// semantics that need a dedicated design.

func gdalExtendedDataTypeCanConvertTo(sourceEDT, targetEDT GDALExtendedDataType) (result bool) {
	result = C.GDALExtendedDataTypeCanConvertTo(sourceEDT.cValue, targetEDT.cValue) != 0
	return
}

func gdalExtendedDataTypeEquals(firstEDT, secondEDT GDALExtendedDataType) (result bool) {
	result = C.GDALExtendedDataTypeEquals(firstEDT.cValue, secondEDT.cValue) != 0
	return
}

func gdalExtendedDataTypeGetSubType(edt GDALExtendedDataType) (result GDALExtendedDataTypeSubType) {
	result = GDALExtendedDataTypeSubType(C.GDALExtendedDataTypeGetSubType(edt.cValue))
	return
}

func gdalExtendedDataTypeGetRAT(edt GDALExtendedDataType) (result GDALRasterAttributeTable) {
	result = GDALRasterAttributeTable{cValue: C.GDALExtendedDataTypeGetRAT(edt.cValue)}
	return
}

func gdalEDTComponentCreate(name string, offset int, dataType GDALExtendedDataType) (result GDALEDTComponent) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALEDTComponent{cValue: C.GDALEDTComponentCreate(cName, C.size_t(offset), dataType.cValue)}
	return
}

func gdalEDTComponentRelease(comp GDALEDTComponent) {
	C.GDALEDTComponentRelease(comp.cValue)
}

func gdalEDTComponentGetName(comp GDALEDTComponent) (result string) {
	result = C.GoString(C.GDALEDTComponentGetName(comp.cValue))
	return
}

func gdalEDTComponentGetOffset(comp GDALEDTComponent) (result int) {
	result = int(C.GDALEDTComponentGetOffset(comp.cValue))
	return
}

func gdalEDTComponentGetType(comp GDALEDTComponent) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALEDTComponentGetType(comp.cValue)}
	return
}

func gdalDatasetGetRootGroup(dataset GDALDataset) (result GDALGroup) {
	result = GDALGroup{cValue: C.GDALDatasetGetRootGroup(dataset.cValue)}
	return
}

func gdalGroupRelease(group GDALGroup) {
	C.GDALGroupRelease(group.cValue)
}

func gdalGroupGetName(group GDALGroup) (result string) {
	result = C.GoString(C.GDALGroupGetName(group.cValue))
	return
}

func gdalGroupGetFullName(group GDALGroup) (result string) {
	result = C.GoString(C.GDALGroupGetFullName(group.cValue))
	return
}

func gdalGroupGetMDArrayNames(group GDALGroup, options CSLConstList) (result CSLConstList) {
	opts := options.cValue
	raw := C.GDALGroupGetMDArrayNames(group.cValue, opts)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalGroupGetMDArrayFullNamesRecursive(group GDALGroup, groupOptions, arrayOptions CSLConstList) (result CSLConstList) {
	gOpts := groupOptions.cValue
	aOpts := arrayOptions.cValue
	raw := C.GDALGroupGetMDArrayFullNamesRecursive(group.cValue, gOpts, aOpts)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalGroupOpenMDArray(group GDALGroup, name string, options CSLConstList) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = GDALMDArray{cValue: C.GDALGroupOpenMDArray(group.cValue, cName, opts)}
	return
}

func gdalGroupOpenMDArrayFromFullname(group GDALGroup, name string, options CSLConstList) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = GDALMDArray{cValue: C.GDALGroupOpenMDArrayFromFullname(group.cValue, cName, opts)}
	return
}

func gdalGroupResolveMDArray(group GDALGroup, name, startingPoint string, options CSLConstList) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cStart := C.CString(startingPoint)
	defer C.free(unsafe.Pointer(cStart))
	opts := options.cValue
	result = GDALMDArray{cValue: C.GDALGroupResolveMDArray(group.cValue, cName, cStart, opts)}
	return
}

func gdalGroupGetGroupNames(group GDALGroup, options CSLConstList) (result CSLConstList) {
	opts := options.cValue
	raw := C.GDALGroupGetGroupNames(group.cValue, opts)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalGroupOpenGroup(group GDALGroup, name string, options CSLConstList) (result GDALGroup) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = GDALGroup{cValue: C.GDALGroupOpenGroup(group.cValue, cName, opts)}
	return
}

func gdalGroupOpenGroupFromFullname(group GDALGroup, name string, options CSLConstList) (result GDALGroup) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = GDALGroup{cValue: C.GDALGroupOpenGroupFromFullname(group.cValue, cName, opts)}
	return
}

func gdalGroupGetVectorLayerNames(group GDALGroup, options CSLConstList) (result CSLConstList) {
	opts := options.cValue
	raw := C.GDALGroupGetVectorLayerNames(group.cValue, opts)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
	return
}

func gdalGroupOpenVectorLayer(group GDALGroup, name string, options CSLConstList) (result OGRLayer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = OGRLayer{cValue: C.GDALGroupOpenVectorLayer(group.cValue, cName, opts)}
	return
}

func gdalGroupGetDimensions(group GDALGroup, options CSLConstList) (result []GDALDimension) {
	opts := options.cValue
	var count C.size_t
	arr := C.GDALGroupGetDimensions(group.cValue, &count, opts)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALDimension, int(count))
	for i := range result {
		result[i] = GDALDimension{cValue: src[i]}
	}
	vsiFree(unsafe.Pointer(arr))
	return
}

func gdalGroupGetAttribute(group GDALGroup, name string) (result GDALAttribute) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = GDALAttribute{cValue: C.GDALGroupGetAttribute(group.cValue, cName)}
	return
}

func gdalGroupGetAttributes(group GDALGroup, options CSLConstList) (result []GDALAttribute) {
	opts := options.cValue
	var count C.size_t
	arr := C.GDALGroupGetAttributes(group.cValue, &count, opts)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALAttribute, int(count))
	for i := range result {
		result[i] = GDALAttribute{cValue: src[i]}
	}
	vsiFree(unsafe.Pointer(arr))
	return
}

func gdalGroupGetStructuralInfo(group GDALGroup) (result CSLConstList) {
	result = cslConstList(C.GDALGroupGetStructuralInfo(group.cValue))
	return
}

func gdalGroupCreateGroup(group GDALGroup, name string, options CSLConstList) (result GDALGroup) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = GDALGroup{cValue: C.GDALGroupCreateGroup(group.cValue, cName, opts)}
	return
}

func gdalGroupDeleteGroup(group GDALGroup, name string, options CSLConstList) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = bool(C.GDALGroupDeleteGroup(group.cValue, cName, opts))
	return
}

func gdalGroupCreateDimension(group GDALGroup, name, dimType, direction string, size uint64, options CSLConstList) (result GDALDimension) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cType := C.CString(dimType)
	defer C.free(unsafe.Pointer(cType))
	cDir := C.CString(direction)
	defer C.free(unsafe.Pointer(cDir))
	opts := options.cValue
	result = GDALDimension{cValue: C.GDALGroupCreateDimension(group.cValue, cName, cType, cDir, C.GUInt64(size), opts)}
	return
}

func gdalGroupCreateMDArray(group GDALGroup, name string, nDimensions int, dimensions []GDALDimension, dataType GDALExtendedDataType, options CSLConstList) (result GDALMDArray) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	var dimsPtr *C.GDALDimensionH
	if len(dimensions) > 0 {
		dimsPtr = (*C.GDALDimensionH)(unsafe.Pointer(&dimensions[0]))
	}
	result = GDALMDArray{cValue: C.GDALGroupCreateMDArray(group.cValue, cName, C.size_t(nDimensions), dimsPtr, dataType.cValue, opts)}
	return
}

func gdalGroupDeleteMDArray(group GDALGroup, name string, options CSLConstList) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = bool(C.GDALGroupDeleteMDArray(group.cValue, cName, opts))
	return
}

func gdalGroupCreateAttribute(group GDALGroup, name string, nDimensions int, dimensions []uint64, dataType GDALExtendedDataType, options CSLConstList) (result GDALAttribute) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	cDims := make([]C.GUInt64, len(dimensions))
	for i, v := range dimensions {
		cDims[i] = C.GUInt64(v)
	}
	var dimsPtr *C.GUInt64
	if len(dimensions) > 0 {
		dimsPtr = &cDims[0]
	}
	result = GDALAttribute{cValue: C.GDALGroupCreateAttribute(group.cValue, cName, C.size_t(nDimensions), dimsPtr, dataType.cValue, opts)}
	return
}

func gdalGroupDeleteAttribute(group GDALGroup, name string, options CSLConstList) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = bool(C.GDALGroupDeleteAttribute(group.cValue, cName, opts))
	return
}

func gdalGroupRename(group GDALGroup, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALGroupRename(group.cValue, cName))
	return
}

func gdalGroupSubsetDimensionFromSelection(group GDALGroup, selection string, options CSLConstList) (result GDALGroup) {
	cSelection := C.CString(selection)
	defer C.free(unsafe.Pointer(cSelection))
	opts := options.cValue
	result = GDALGroup{cValue: C.GDALGroupSubsetDimensionFromSelection(group.cValue, cSelection, opts)}
	return
}

func gdalGroupGetDataTypeCount(group GDALGroup) (result int) {
	result = int(C.GDALGroupGetDataTypeCount(group.cValue))
	return
}

func gdalGroupGetDataType(group GDALGroup, index int) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALGroupGetDataType(group.cValue, C.size_t(index))}
	return
}

func gdalMDArrayRelease(array GDALMDArray) {
	C.GDALMDArrayRelease(array.cValue)
}

func gdalMDArrayGetName(array GDALMDArray) (result string) {
	result = C.GoString(C.GDALMDArrayGetName(array.cValue))
	return
}

func gdalMDArrayGetFullName(array GDALMDArray) (result string) {
	result = C.GoString(C.GDALMDArrayGetFullName(array.cValue))
	return
}

func gdalMDArrayGetTotalElementsCount(array GDALMDArray) (result uint64) {
	result = uint64(C.GDALMDArrayGetTotalElementsCount(array.cValue))
	return
}

func gdalMDArrayGetDimensionCount(array GDALMDArray) (result int) {
	result = int(C.GDALMDArrayGetDimensionCount(array.cValue))
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
	vsiFree(unsafe.Pointer(arr))
	return
}

func gdalMDArrayGetDataType(array GDALMDArray) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALMDArrayGetDataType(array.cValue)}
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

func gdalMDArrayGetAttributes(array GDALMDArray, options CSLConstList) (result []GDALAttribute) {
	opts := options.cValue
	var count C.size_t
	arr := C.GDALMDArrayGetAttributes(array.cValue, &count, opts)
	if arr == nil || count == 0 {
		return
	}
	src := unsafe.Slice(arr, int(count))
	result = make([]GDALAttribute, int(count))
	for i := range result {
		result[i] = GDALAttribute{cValue: src[i]}
	}
	vsiFree(unsafe.Pointer(arr))
	return
}

func gdalMDArrayCreateAttribute(array GDALMDArray, name string, nDimensions int, dimensions []uint64, dataType GDALExtendedDataType, options CSLConstList) (result GDALAttribute) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	cDims := make([]C.GUInt64, len(dimensions))
	for i, v := range dimensions {
		cDims[i] = C.GUInt64(v)
	}
	var dimsPtr *C.GUInt64
	if len(dimensions) > 0 {
		dimsPtr = &cDims[0]
	}
	result = GDALAttribute{cValue: C.GDALMDArrayCreateAttribute(array.cValue, cName, C.size_t(nDimensions), dimsPtr, dataType.cValue, opts)}
	return
}

func gdalMDArrayDeleteAttribute(array GDALMDArray, name string, options CSLConstList) (result bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	opts := options.cValue
	result = bool(C.GDALMDArrayDeleteAttribute(array.cValue, cName, opts))
	return
}

func gdalMDArrayResize(array GDALMDArray, newDimSizes []uint64, options CSLConstList) (result bool) {
	opts := options.cValue
	cSizes := make([]C.GUInt64, len(newDimSizes))
	for i, v := range newDimSizes {
		cSizes[i] = C.GUInt64(v)
	}
	var sizesPtr *C.GUInt64
	if len(newDimSizes) > 0 {
		sizesPtr = &cSizes[0]
	}
	result = bool(C.GDALMDArrayResize(array.cValue, sizesPtr, opts))
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

func gdalMDArrayGetNoDataValueAsInt64(array GDALMDArray, hasNoData *int) (result int64) {
	var cHas C.int
	result = int64(C.GDALMDArrayGetNoDataValueAsInt64(array.cValue, &cHas))
	*hasNoData = int(cHas)
	return
}

func gdalMDArrayGetNoDataValueAsUInt64(array GDALMDArray, hasNoData *int) (result uint64) {
	var cHas C.int
	result = uint64(C.GDALMDArrayGetNoDataValueAsUInt64(array.cValue, &cHas))
	*hasNoData = int(cHas)
	return
}

func gdalMDArraySetNoDataValueAsDouble(array GDALMDArray, value float64) (result bool) {
	result = C.GDALMDArraySetNoDataValueAsDouble(array.cValue, C.double(value)) != 0
	return
}

func gdalMDArraySetNoDataValueAsInt64(array GDALMDArray, value int64) (result bool) {
	result = C.GDALMDArraySetNoDataValueAsInt64(array.cValue, C.int64_t(value)) != 0
	return
}

func gdalMDArraySetNoDataValueAsUInt64(array GDALMDArray, value uint64) (result bool) {
	result = C.GDALMDArraySetNoDataValueAsUInt64(array.cValue, C.uint64_t(value)) != 0
	return
}

func gdalMDArraySetScale(array GDALMDArray, scale float64) (result bool) {
	result = C.GDALMDArraySetScale(array.cValue, C.double(scale)) != 0
	return
}

func gdalMDArraySetScaleEx(array GDALMDArray, scale float64, storageType GDALDataType) (result bool) {
	result = C.GDALMDArraySetScaleEx(array.cValue, C.double(scale), C.GDALDataType(storageType)) != 0
	return
}

func gdalMDArrayGetScale(array GDALMDArray, hasValue *int) (result float64) {
	var cHas C.int
	result = float64(C.GDALMDArrayGetScale(array.cValue, &cHas))
	*hasValue = int(cHas)
	return
}

func gdalMDArraySetOffset(array GDALMDArray, offset float64) (result bool) {
	result = C.GDALMDArraySetOffset(array.cValue, C.double(offset)) != 0
	return
}

func gdalMDArraySetOffsetEx(array GDALMDArray, offset float64, storageType GDALDataType) (result bool) {
	result = C.GDALMDArraySetOffsetEx(array.cValue, C.double(offset), C.GDALDataType(storageType)) != 0
	return
}

func gdalMDArrayGetOffset(array GDALMDArray, hasValue *int) (result float64) {
	var cHas C.int
	result = float64(C.GDALMDArrayGetOffset(array.cValue, &cHas))
	*hasValue = int(cHas)
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

func gdalMDArrayGetUnit(array GDALMDArray) (result string) {
	result = C.GoString(C.GDALMDArrayGetUnit(array.cValue))
	return
}

func gdalMDArraySetSpatialRef(array GDALMDArray, srs OGRSpatialReference) (result bool) {
	result = C.GDALMDArraySetSpatialRef(array.cValue, srs.cValue) != 0
	return
}

func gdalMDArrayGetSpatialRef(array GDALMDArray) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.GDALMDArrayGetSpatialRef(array.cValue)}
	return
}

func gdalMDArrayGetStructuralInfo(array GDALMDArray) (result CSLConstList) {
	result = cslConstList(C.GDALMDArrayGetStructuralInfo(array.cValue))
	return
}

func gdalMDArrayGetView(array GDALMDArray, viewExpr string) (result GDALMDArray) {
	cExpr := C.CString(viewExpr)
	defer C.free(unsafe.Pointer(cExpr))
	result = GDALMDArray{cValue: C.GDALMDArrayGetView(array.cValue, cExpr)}
	return
}

func gdalMDArrayTranspose(array GDALMDArray, newAxisCount int, mapNewAxisToOldAxis []int) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALMDArrayTranspose(array.cValue, C.size_t(newAxisCount), cInts(mapNewAxisToOldAxis))}
	return
}

func gdalMDArrayGetUnscaled(array GDALMDArray) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALMDArrayGetUnscaled(array.cValue)}
	return
}

func gdalMDArrayGetMask(array GDALMDArray, options CSLConstList) (result GDALMDArray) {
	opts := options.cValue
	result = GDALMDArray{cValue: C.GDALMDArrayGetMask(array.cValue, opts)}
	return
}

func gdalMDArrayAsClassicDataset(array GDALMDArray, xDim, yDim int) (result GDALDataset) {
	result = GDALDataset{cValue: C.GDALMDArrayAsClassicDataset(array.cValue, C.size_t(xDim), C.size_t(yDim))}
	return
}

func gdalMDArrayAsClassicDatasetEx(array GDALMDArray, xDim, yDim int, rootGroup GDALGroup, options CSLConstList) (result GDALDataset) {
	opts := options.cValue
	result = GDALDataset{cValue: C.GDALMDArrayAsClassicDatasetEx(array.cValue, C.size_t(xDim), C.size_t(yDim), rootGroup.cValue, opts)}
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

func gdalMDArrayGetResampled(array GDALMDArray, newDimCount int, newDims []GDALDimension, resampleAlg GDALRIOResampleAlg, targetSRS OGRSpatialReference, options CSLConstList) (result GDALMDArray) {
	opts := options.cValue
	var dimsPtr *C.GDALDimensionH
	if len(newDims) > 0 {
		dimsPtr = (*C.GDALDimensionH)(unsafe.Pointer(&newDims[0]))
	}
	result = GDALMDArray{cValue: C.GDALMDArrayGetResampled(array.cValue, C.size_t(newDimCount), dimsPtr, C.GDALRIOResampleAlg(resampleAlg), targetSRS.cValue, opts)}
	return
}

func gdalMDArrayGetGridded(array GDALMDArray, gridOptions string, xArray, yArray GDALMDArray, options CSLConstList) (result GDALMDArray) {
	cGridOptions := C.CString(gridOptions)
	defer C.free(unsafe.Pointer(cGridOptions))
	opts := options.cValue
	result = GDALMDArray{cValue: C.GDALMDArrayGetGridded(array.cValue, cGridOptions, xArray.cValue, yArray.cValue, opts)}
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
	vsiFree(unsafe.Pointer(arr))
	return
}

// GDALMDArrayGetMeshGrid (array-in/array-out), the GDALMDArrayRawBlockInfo
// struct + GDALMDArrayRawBlockInfoCreate/Release/GDALMDArrayGetRawBlockInfo, and
// GDALReleaseArrays are deferred.

func gdalMDArrayCache(array GDALMDArray, options CSLConstList) (result bool) {
	opts := options.cValue
	result = C.GDALMDArrayCache(array.cValue, opts) != 0
	return
}

func gdalMDArrayRename(array GDALMDArray, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALMDArrayRename(array.cValue, cName))
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

func gdalAttributeRelease(attr GDALAttribute) {
	C.GDALAttributeRelease(attr.cValue)
}

func gdalAttributeGetName(attr GDALAttribute) (result string) {
	result = C.GoString(C.GDALAttributeGetName(attr.cValue))
	return
}

func gdalAttributeGetFullName(attr GDALAttribute) (result string) {
	result = C.GoString(C.GDALAttributeGetFullName(attr.cValue))
	return
}

func gdalAttributeGetTotalElementsCount(attr GDALAttribute) (result uint64) {
	result = uint64(C.GDALAttributeGetTotalElementsCount(attr.cValue))
	return
}

func gdalAttributeGetDimensionCount(attr GDALAttribute) (result int) {
	result = int(C.GDALAttributeGetDimensionCount(attr.cValue))
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
	vsiFree(unsafe.Pointer(arr))
	return
}

func gdalAttributeGetDataType(attr GDALAttribute) (result GDALExtendedDataType) {
	result = GDALExtendedDataType{cValue: C.GDALAttributeGetDataType(attr.cValue)}
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

func gdalAttributeReadAsString(attr GDALAttribute) (result string) {
	result = C.GoString(C.GDALAttributeReadAsString(attr.cValue))
	return
}

func gdalAttributeReadAsInt(attr GDALAttribute) (result int) {
	result = int(C.GDALAttributeReadAsInt(attr.cValue))
	return
}

func gdalAttributeReadAsInt64(attr GDALAttribute) (result int64) {
	result = int64(C.GDALAttributeReadAsInt64(attr.cValue))
	return
}

func gdalAttributeReadAsDouble(attr GDALAttribute) (result float64) {
	result = float64(C.GDALAttributeReadAsDouble(attr.cValue))
	return
}

func gdalAttributeReadAsStringArray(attr GDALAttribute) (result CSLConstList) {
	raw := C.GDALAttributeReadAsStringArray(attr.cValue)
	if raw == nil {
		return
	}
	result = cslConstList(raw)
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
	vsiFree(unsafe.Pointer(arr))
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
	vsiFree(unsafe.Pointer(arr))
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
	vsiFree(unsafe.Pointer(arr))
	return
}

func gdalAttributeWriteRaw(attr GDALAttribute, data unsafe.Pointer, size int) (result bool) {
	result = C.GDALAttributeWriteRaw(attr.cValue, data, C.size_t(size)) != 0
	return
}

func gdalAttributeWriteString(attr GDALAttribute, value string) (result bool) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = C.GDALAttributeWriteString(attr.cValue, cValue) != 0
	return
}

func gdalAttributeWriteStringArray(attr GDALAttribute, values CSLConstList) (result bool) {
	v := values.cValue
	result = C.GDALAttributeWriteStringArray(attr.cValue, v) != 0
	return
}

func gdalAttributeWriteInt(attr GDALAttribute, value int) (result bool) {
	result = C.GDALAttributeWriteInt(attr.cValue, C.int(value)) != 0
	return
}

func gdalAttributeWriteIntArray(attr GDALAttribute, values []int, count int) (result bool) {
	result = C.GDALAttributeWriteIntArray(attr.cValue, cInts(values), C.size_t(count)) != 0
	return
}

func gdalAttributeWriteInt64(attr GDALAttribute, value int64) (result bool) {
	result = C.GDALAttributeWriteInt64(attr.cValue, C.int64_t(value)) != 0
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

func gdalAttributeWriteDouble(attr GDALAttribute, value float64) (result bool) {
	result = C.GDALAttributeWriteDouble(attr.cValue, C.double(value)) != 0
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

func gdalAttributeRename(attr GDALAttribute, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALAttributeRename(attr.cValue, cName))
	return
}

func gdalDimensionRelease(dim GDALDimension) {
	C.GDALDimensionRelease(dim.cValue)
}

func gdalDimensionGetName(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetName(dim.cValue))
	return
}

func gdalDimensionGetFullName(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetFullName(dim.cValue))
	return
}

func gdalDimensionGetType(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetType(dim.cValue))
	return
}

func gdalDimensionGetDirection(dim GDALDimension) (result string) {
	result = C.GoString(C.GDALDimensionGetDirection(dim.cValue))
	return
}

func gdalDimensionGetSize(dim GDALDimension) (result uint64) {
	result = uint64(C.GDALDimensionGetSize(dim.cValue))
	return
}

func gdalDimensionGetIndexingVariable(dim GDALDimension) (result GDALMDArray) {
	result = GDALMDArray{cValue: C.GDALDimensionGetIndexingVariable(dim.cValue)}
	return
}

func gdalDimensionSetIndexingVariable(dim GDALDimension, array GDALMDArray) (result bool) {
	result = C.GDALDimensionSetIndexingVariable(dim.cValue, array.cValue) != 0
	return
}

func gdalDimensionRename(dim GDALDimension, newName string) (result bool) {
	cName := C.CString(newName)
	defer C.free(unsafe.Pointer(cName))
	result = bool(C.GDALDimensionRename(dim.cValue, cName))
	return
}

// CPL_C_END
