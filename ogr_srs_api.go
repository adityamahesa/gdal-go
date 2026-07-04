package gdal

/*
#include "ogr_srs_api.h"
#include "cpl_string.h" // TODO: implement cpl_string.go

const char* const _SRS_WKT_WGS84_LAT_LONG = SRS_WKT_WGS84_LAT_LONG;

const char* const _SRS_PT_ALBERS_CONIC_EQUAL_AREA                       = SRS_PT_ALBERS_CONIC_EQUAL_AREA;
const char* const _SRS_PT_AZIMUTHAL_EQUIDISTANT                         = SRS_PT_AZIMUTHAL_EQUIDISTANT;
const char* const _SRS_PT_CASSINI_SOLDNER                               = SRS_PT_CASSINI_SOLDNER;
const char* const _SRS_PT_CYLINDRICAL_EQUAL_AREA                        = SRS_PT_CYLINDRICAL_EQUAL_AREA;
const char* const _SRS_PT_BONNE                                         = SRS_PT_BONNE;
const char* const _SRS_PT_ECKERT_I                                      = SRS_PT_ECKERT_I;
const char* const _SRS_PT_ECKERT_II                                     = SRS_PT_ECKERT_II;
const char* const _SRS_PT_ECKERT_III                                    = SRS_PT_ECKERT_III;
const char* const _SRS_PT_ECKERT_IV                                     = SRS_PT_ECKERT_IV;
const char* const _SRS_PT_ECKERT_V                                      = SRS_PT_ECKERT_V;
const char* const _SRS_PT_ECKERT_VI                                     = SRS_PT_ECKERT_VI;
const char* const _SRS_PT_EQUIDISTANT_CONIC                             = SRS_PT_EQUIDISTANT_CONIC;
const char* const _SRS_PT_EQUIRECTANGULAR                               = SRS_PT_EQUIRECTANGULAR;
const char* const _SRS_PT_GALL_STEREOGRAPHIC                            = SRS_PT_GALL_STEREOGRAPHIC;
const char* const _SRS_PT_GAUSSSCHREIBERTMERCATOR                       = SRS_PT_GAUSSSCHREIBERTMERCATOR;
const char* const _SRS_PT_GEOSTATIONARY_SATELLITE                       = SRS_PT_GEOSTATIONARY_SATELLITE;
const char* const _SRS_PT_GOODE_HOMOLOSINE                              = SRS_PT_GOODE_HOMOLOSINE;
const char* const _SRS_PT_IGH                                           = SRS_PT_IGH;
const char* const _SRS_PT_GNOMONIC                                      = SRS_PT_GNOMONIC;
const char* const _SRS_PT_HOTINE_OBLIQUE_MERCATOR_AZIMUTH_CENTER        = SRS_PT_HOTINE_OBLIQUE_MERCATOR_AZIMUTH_CENTER;
const char* const _SRS_PT_HOTINE_OBLIQUE_MERCATOR                       = SRS_PT_HOTINE_OBLIQUE_MERCATOR;
const char* const _SRS_PT_HOTINE_OBLIQUE_MERCATOR_TWO_POINT_NATURAL_ORIGIN = SRS_PT_HOTINE_OBLIQUE_MERCATOR_TWO_POINT_NATURAL_ORIGIN;
const char* const _SRS_PT_LABORDE_OBLIQUE_MERCATOR                      = SRS_PT_LABORDE_OBLIQUE_MERCATOR;
const char* const _SRS_PT_LAMBERT_CONFORMAL_CONIC_1SP                   = SRS_PT_LAMBERT_CONFORMAL_CONIC_1SP;
const char* const _SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP                   = SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP;
const char* const _SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP_BELGIUM           = SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP_BELGIUM;
const char* const _SRS_PT_LAMBERT_AZIMUTHAL_EQUAL_AREA                  = SRS_PT_LAMBERT_AZIMUTHAL_EQUAL_AREA;
const char* const _SRS_PT_MERCATOR_1SP                                  = SRS_PT_MERCATOR_1SP;
const char* const _SRS_PT_MERCATOR_2SP                                  = SRS_PT_MERCATOR_2SP;
const char* const _SRS_PT_MERCATOR_AUXILIARY_SPHERE                     = SRS_PT_MERCATOR_AUXILIARY_SPHERE;
const char* const _SRS_PT_MILLER_CYLINDRICAL                            = SRS_PT_MILLER_CYLINDRICAL;
const char* const _SRS_PT_MOLLWEIDE                                     = SRS_PT_MOLLWEIDE;
const char* const _SRS_PT_NEW_ZEALAND_MAP_GRID                          = SRS_PT_NEW_ZEALAND_MAP_GRID;
const char* const _SRS_PT_OBLIQUE_STEREOGRAPHIC                         = SRS_PT_OBLIQUE_STEREOGRAPHIC;
const char* const _SRS_PT_ORTHOGRAPHIC                                  = SRS_PT_ORTHOGRAPHIC;
const char* const _SRS_PT_POLAR_STEREOGRAPHIC                           = SRS_PT_POLAR_STEREOGRAPHIC;
const char* const _SRS_PT_POLYCONIC                                     = SRS_PT_POLYCONIC;
const char* const _SRS_PT_ROBINSON                                      = SRS_PT_ROBINSON;
const char* const _SRS_PT_SINUSOIDAL                                    = SRS_PT_SINUSOIDAL;
const char* const _SRS_PT_STEREOGRAPHIC                                 = SRS_PT_STEREOGRAPHIC;
const char* const _SRS_PT_SWISS_OBLIQUE_CYLINDRICAL                     = SRS_PT_SWISS_OBLIQUE_CYLINDRICAL;
const char* const _SRS_PT_TRANSVERSE_MERCATOR                           = SRS_PT_TRANSVERSE_MERCATOR;
const char* const _SRS_PT_TRANSVERSE_MERCATOR_SOUTH_ORIENTED            = SRS_PT_TRANSVERSE_MERCATOR_SOUTH_ORIENTED;
const char* const _SRS_PT_TRANSVERSE_MERCATOR_MI_21                     = SRS_PT_TRANSVERSE_MERCATOR_MI_21;
const char* const _SRS_PT_TRANSVERSE_MERCATOR_MI_22                     = SRS_PT_TRANSVERSE_MERCATOR_MI_22;
const char* const _SRS_PT_TRANSVERSE_MERCATOR_MI_23                     = SRS_PT_TRANSVERSE_MERCATOR_MI_23;
const char* const _SRS_PT_TRANSVERSE_MERCATOR_MI_24                     = SRS_PT_TRANSVERSE_MERCATOR_MI_24;
const char* const _SRS_PT_TRANSVERSE_MERCATOR_MI_25                     = SRS_PT_TRANSVERSE_MERCATOR_MI_25;
const char* const _SRS_PT_TUNISIA_MINING_GRID                           = SRS_PT_TUNISIA_MINING_GRID;
const char* const _SRS_PT_TWO_POINT_EQUIDISTANT                         = SRS_PT_TWO_POINT_EQUIDISTANT;
const char* const _SRS_PT_VANDERGRINTEN                                 = SRS_PT_VANDERGRINTEN;
const char* const _SRS_PT_KROVAK                                        = SRS_PT_KROVAK;
const char* const _SRS_PT_IMW_POLYCONIC                                 = SRS_PT_IMW_POLYCONIC;
const char* const _SRS_PT_WAGNER_I                                      = SRS_PT_WAGNER_I;
const char* const _SRS_PT_WAGNER_II                                     = SRS_PT_WAGNER_II;
const char* const _SRS_PT_WAGNER_III                                    = SRS_PT_WAGNER_III;
const char* const _SRS_PT_WAGNER_IV                                     = SRS_PT_WAGNER_IV;
const char* const _SRS_PT_WAGNER_V                                      = SRS_PT_WAGNER_V;
const char* const _SRS_PT_WAGNER_VI                                     = SRS_PT_WAGNER_VI;
const char* const _SRS_PT_WAGNER_VII                                    = SRS_PT_WAGNER_VII;
const char* const _SRS_PT_QSC                                           = SRS_PT_QSC;
const char* const _SRS_PT_AITOFF                                        = SRS_PT_AITOFF;
const char* const _SRS_PT_WINKEL_I                                      = SRS_PT_WINKEL_I;
const char* const _SRS_PT_WINKEL_II                                     = SRS_PT_WINKEL_II;
const char* const _SRS_PT_WINKEL_TRIPEL                                 = SRS_PT_WINKEL_TRIPEL;
const char* const _SRS_PT_CRASTER_PARABOLIC                             = SRS_PT_CRASTER_PARABOLIC;
const char* const _SRS_PT_LOXIMUTHAL                                    = SRS_PT_LOXIMUTHAL;
const char* const _SRS_PT_QUARTIC_AUTHALIC                              = SRS_PT_QUARTIC_AUTHALIC;
const char* const _SRS_PT_SCH                                           = SRS_PT_SCH;

const char* const _SRS_PP_CENTRAL_MERIDIAN        = SRS_PP_CENTRAL_MERIDIAN;
const char* const _SRS_PP_SCALE_FACTOR            = SRS_PP_SCALE_FACTOR;
const char* const _SRS_PP_STANDARD_PARALLEL_1     = SRS_PP_STANDARD_PARALLEL_1;
const char* const _SRS_PP_STANDARD_PARALLEL_2     = SRS_PP_STANDARD_PARALLEL_2;
const char* const _SRS_PP_PSEUDO_STD_PARALLEL_1   = SRS_PP_PSEUDO_STD_PARALLEL_1;
const char* const _SRS_PP_LONGITUDE_OF_CENTER     = SRS_PP_LONGITUDE_OF_CENTER;
const char* const _SRS_PP_LATITUDE_OF_CENTER      = SRS_PP_LATITUDE_OF_CENTER;
const char* const _SRS_PP_LONGITUDE_OF_ORIGIN     = SRS_PP_LONGITUDE_OF_ORIGIN;
const char* const _SRS_PP_LATITUDE_OF_ORIGIN      = SRS_PP_LATITUDE_OF_ORIGIN;
const char* const _SRS_PP_FALSE_EASTING           = SRS_PP_FALSE_EASTING;
const char* const _SRS_PP_FALSE_NORTHING          = SRS_PP_FALSE_NORTHING;
const char* const _SRS_PP_AZIMUTH                 = SRS_PP_AZIMUTH;
const char* const _SRS_PP_LONGITUDE_OF_POINT_1    = SRS_PP_LONGITUDE_OF_POINT_1;
const char* const _SRS_PP_LATITUDE_OF_POINT_1     = SRS_PP_LATITUDE_OF_POINT_1;
const char* const _SRS_PP_LONGITUDE_OF_POINT_2    = SRS_PP_LONGITUDE_OF_POINT_2;
const char* const _SRS_PP_LATITUDE_OF_POINT_2     = SRS_PP_LATITUDE_OF_POINT_2;
const char* const _SRS_PP_LONGITUDE_OF_POINT_3    = SRS_PP_LONGITUDE_OF_POINT_3;
const char* const _SRS_PP_LATITUDE_OF_POINT_3     = SRS_PP_LATITUDE_OF_POINT_3;
const char* const _SRS_PP_RECTIFIED_GRID_ANGLE    = SRS_PP_RECTIFIED_GRID_ANGLE;
const char* const _SRS_PP_LANDSAT_NUMBER          = SRS_PP_LANDSAT_NUMBER;
const char* const _SRS_PP_PATH_NUMBER             = SRS_PP_PATH_NUMBER;
const char* const _SRS_PP_PERSPECTIVE_POINT_HEIGHT = SRS_PP_PERSPECTIVE_POINT_HEIGHT;
const char* const _SRS_PP_SATELLITE_HEIGHT        = SRS_PP_SATELLITE_HEIGHT;
const char* const _SRS_PP_FIPSZONE                = SRS_PP_FIPSZONE;
const char* const _SRS_PP_ZONE                    = SRS_PP_ZONE;
const char* const _SRS_PP_LATITUDE_OF_1ST_POINT   = SRS_PP_LATITUDE_OF_1ST_POINT;
const char* const _SRS_PP_LONGITUDE_OF_1ST_POINT  = SRS_PP_LONGITUDE_OF_1ST_POINT;
const char* const _SRS_PP_LATITUDE_OF_2ND_POINT   = SRS_PP_LATITUDE_OF_2ND_POINT;
const char* const _SRS_PP_LONGITUDE_OF_2ND_POINT  = SRS_PP_LONGITUDE_OF_2ND_POINT;
const char* const _SRS_PP_PEG_POINT_LATITUDE      = SRS_PP_PEG_POINT_LATITUDE;
const char* const _SRS_PP_PEG_POINT_LONGITUDE     = SRS_PP_PEG_POINT_LONGITUDE;
const char* const _SRS_PP_PEG_POINT_HEADING       = SRS_PP_PEG_POINT_HEADING;
const char* const _SRS_PP_PEG_POINT_HEIGHT        = SRS_PP_PEG_POINT_HEIGHT;

const char* const _SRS_UL_METER              = SRS_UL_METER;
const char* const _SRS_UL_FOOT              = SRS_UL_FOOT;
const char* const _SRS_UL_FOOT_CONV         = SRS_UL_FOOT_CONV;
const char* const _SRS_UL_US_FOOT           = SRS_UL_US_FOOT;
const char* const _SRS_UL_US_FOOT_CONV      = SRS_UL_US_FOOT_CONV;
const char* const _SRS_UL_NAUTICAL_MILE     = SRS_UL_NAUTICAL_MILE;
const char* const _SRS_UL_NAUTICAL_MILE_CONV = SRS_UL_NAUTICAL_MILE_CONV;
const char* const _SRS_UL_LINK              = SRS_UL_LINK;
const char* const _SRS_UL_LINK_CONV         = SRS_UL_LINK_CONV;
const char* const _SRS_UL_CHAIN             = SRS_UL_CHAIN;
const char* const _SRS_UL_CHAIN_CONV        = SRS_UL_CHAIN_CONV;
const char* const _SRS_UL_ROD               = SRS_UL_ROD;
const char* const _SRS_UL_ROD_CONV          = SRS_UL_ROD_CONV;
const char* const _SRS_UL_LINK_Clarke       = SRS_UL_LINK_Clarke;
const char* const _SRS_UL_LINK_Clarke_CONV  = SRS_UL_LINK_Clarke_CONV;
const char* const _SRS_UL_KILOMETER         = SRS_UL_KILOMETER;
const char* const _SRS_UL_KILOMETER_CONV    = SRS_UL_KILOMETER_CONV;
const char* const _SRS_UL_DECIMETER         = SRS_UL_DECIMETER;
const char* const _SRS_UL_DECIMETER_CONV    = SRS_UL_DECIMETER_CONV;
const char* const _SRS_UL_CENTIMETER        = SRS_UL_CENTIMETER;
const char* const _SRS_UL_CENTIMETER_CONV   = SRS_UL_CENTIMETER_CONV;
const char* const _SRS_UL_MILLIMETER        = SRS_UL_MILLIMETER;
const char* const _SRS_UL_MILLIMETER_CONV   = SRS_UL_MILLIMETER_CONV;
const char* const _SRS_UL_INTL_NAUT_MILE    = SRS_UL_INTL_NAUT_MILE;
const char* const _SRS_UL_INTL_NAUT_MILE_CONV = SRS_UL_INTL_NAUT_MILE_CONV;
const char* const _SRS_UL_INTL_INCH         = SRS_UL_INTL_INCH;
const char* const _SRS_UL_INTL_INCH_CONV    = SRS_UL_INTL_INCH_CONV;
const char* const _SRS_UL_INTL_FOOT         = SRS_UL_INTL_FOOT;
const char* const _SRS_UL_INTL_FOOT_CONV    = SRS_UL_INTL_FOOT_CONV;
const char* const _SRS_UL_INTL_YARD         = SRS_UL_INTL_YARD;
const char* const _SRS_UL_INTL_YARD_CONV    = SRS_UL_INTL_YARD_CONV;
const char* const _SRS_UL_INTL_STAT_MILE    = SRS_UL_INTL_STAT_MILE;
const char* const _SRS_UL_INTL_STAT_MILE_CONV = SRS_UL_INTL_STAT_MILE_CONV;
const char* const _SRS_UL_INTL_FATHOM       = SRS_UL_INTL_FATHOM;
const char* const _SRS_UL_INTL_FATHOM_CONV  = SRS_UL_INTL_FATHOM_CONV;
const char* const _SRS_UL_INTL_CHAIN        = SRS_UL_INTL_CHAIN;
const char* const _SRS_UL_INTL_CHAIN_CONV   = SRS_UL_INTL_CHAIN_CONV;
const char* const _SRS_UL_INTL_LINK         = SRS_UL_INTL_LINK;
const char* const _SRS_UL_INTL_LINK_CONV    = SRS_UL_INTL_LINK_CONV;
const char* const _SRS_UL_US_INCH           = SRS_UL_US_INCH;
const char* const _SRS_UL_US_INCH_CONV      = SRS_UL_US_INCH_CONV;
const char* const _SRS_UL_US_YARD           = SRS_UL_US_YARD;
const char* const _SRS_UL_US_YARD_CONV      = SRS_UL_US_YARD_CONV;
const char* const _SRS_UL_US_CHAIN          = SRS_UL_US_CHAIN;
const char* const _SRS_UL_US_CHAIN_CONV     = SRS_UL_US_CHAIN_CONV;
const char* const _SRS_UL_US_STAT_MILE      = SRS_UL_US_STAT_MILE;
const char* const _SRS_UL_US_STAT_MILE_CONV = SRS_UL_US_STAT_MILE_CONV;
const char* const _SRS_UL_INDIAN_YARD       = SRS_UL_INDIAN_YARD;
const char* const _SRS_UL_INDIAN_YARD_CONV  = SRS_UL_INDIAN_YARD_CONV;
const char* const _SRS_UL_INDIAN_FOOT       = SRS_UL_INDIAN_FOOT;
const char* const _SRS_UL_INDIAN_FOOT_CONV  = SRS_UL_INDIAN_FOOT_CONV;
const char* const _SRS_UL_INDIAN_CHAIN      = SRS_UL_INDIAN_CHAIN;
const char* const _SRS_UL_INDIAN_CHAIN_CONV = SRS_UL_INDIAN_CHAIN_CONV;

const char* const _SRS_UA_DEGREE      = SRS_UA_DEGREE;
const char* const _SRS_UA_DEGREE_CONV = SRS_UA_DEGREE_CONV;
const char* const _SRS_UA_RADIAN      = SRS_UA_RADIAN;

const char* const _SRS_PM_GREENWICH = SRS_PM_GREENWICH;

const char* const _SRS_DN_NAD27 = SRS_DN_NAD27;
const char* const _SRS_DN_NAD83 = SRS_DN_NAD83;
const char* const _SRS_DN_WGS72 = SRS_DN_WGS72;
const char* const _SRS_DN_WGS84 = SRS_DN_WGS84;
*/
import "C"
import "unsafe"

type OGRAxisOrientation C.OGRAxisOrientation

const (
	OAOOther OGRAxisOrientation = C.OAO_Other
	OAONorth OGRAxisOrientation = C.OAO_North
	OAOSouth OGRAxisOrientation = C.OAO_South
	OAOEast  OGRAxisOrientation = C.OAO_East
	OAOWest  OGRAxisOrientation = C.OAO_West
	OAOUp    OGRAxisOrientation = C.OAO_Up
	OAODown  OGRAxisOrientation = C.OAO_Down
)

func osrAxisEnumToName(e OGRAxisOrientation) (result string) {
	result = C.GoString(C.OSRAxisEnumToName(C.OGRAxisOrientation(e)))
	return
}

func (e OGRAxisOrientation) Name() (result string) {
	result = osrAxisEnumToName(e)
	return
}

// #endif  // ndef SWIG

// /* ==================================================================== */
// /*      Some standard WKT geographic coordinate systems.                */
// /* ==================================================================== */

// #ifdef USE_DEPRECATED_SRS_WKT_WGS84
// #define SRS_WKT_WGS84                                                          \
//     "GEOGCS[\"WGS 84\",DATUM[\"WGS_1984\",SPHEROID[\"WGS "                     \
//     "84\",6378137,298.257223563,AUTHORITY[\"EPSG\",\"7030\"]],AUTHORITY["      \
//     "\"EPSG\",\"6326\"]],PRIMEM[\"Greenwich\",0,AUTHORITY[\"EPSG\",\"8901\"]]" \
//     ",UNIT[\"degree\",0.0174532925199433,AUTHORITY[\"EPSG\",\"9122\"]],"       \
//     "AUTHORITY[\"EPSG\",\"4326\"]]"
// #endif

// /** WGS 84 geodetic (lat/long) WKT / EPSG:4326 with lat,long ordering */
var SRSWKTWgs84LatLong = C.GoString(C._SRS_WKT_WGS84_LAT_LONG)

// /* ==================================================================== */
// /*      Some "standard" strings.                                        */
// /* ==================================================================== */

// /** Albers_Conic_Equal_Area projection */
var SRSPTAlbersConicEqualArea = C.GoString(C._SRS_PT_ALBERS_CONIC_EQUAL_AREA)

// /** Azimuthal_Equidistant projection */
var SRSPTAzimuthalEquidistant = C.GoString(C._SRS_PT_AZIMUTHAL_EQUIDISTANT)

// /** Cassini_Soldner projection */
var SRSPTCassiniSoldner = C.GoString(C._SRS_PT_CASSINI_SOLDNER)

// /** Cylindrical_Equal_Area projection */
var SRSPTCylindricalEqualArea = C.GoString(C._SRS_PT_CYLINDRICAL_EQUAL_AREA)

// /** Cylindrical_Equal_Area projection */
var SRSPTBonne = C.GoString(C._SRS_PT_BONNE)

// /** Eckert_I projection */
var SRSPTEckertI = C.GoString(C._SRS_PT_ECKERT_I)

// /** Eckert_II projection */
var SRSPTEckertII = C.GoString(C._SRS_PT_ECKERT_II)

// /** Eckert_III projection */
var SRSPTEckertIII = C.GoString(C._SRS_PT_ECKERT_III)

// /** Eckert_IV projection */
var SRSPTEckertIV = C.GoString(C._SRS_PT_ECKERT_IV)

// /** Eckert_V projection */
var SRSPTEckertV = C.GoString(C._SRS_PT_ECKERT_V)

// /** Eckert_VI projection */
var SRSPTEckertVI = C.GoString(C._SRS_PT_ECKERT_VI)

// /** Equidistant_Conic projection */
var SRSPTEquidistantConic = C.GoString(C._SRS_PT_EQUIDISTANT_CONIC)

// /** Equirectangular projection */
var SRSPTEquirectangular = C.GoString(C._SRS_PT_EQUIRECTANGULAR)

// /** Gall_Stereographic projection */
var SRSPTGallStereographic = C.GoString(C._SRS_PT_GALL_STEREOGRAPHIC)

// /** Gauss_Schreiber_Transverse_Mercator projection */
var SRSPTGaussschreibertmercator = C.GoString(C._SRS_PT_GAUSSSCHREIBERTMERCATOR)

// /** Geostationary_Satellite projection */
var SRSPTGeostationarySatellite = C.GoString(C._SRS_PT_GEOSTATIONARY_SATELLITE)

// /** Goode_Homolosine projection */
var SRSPTGoodeHomolosine = C.GoString(C._SRS_PT_GOODE_HOMOLOSINE)

// /** Interrupted_Goode_Homolosine projection */
var SRSPTIgh = C.GoString(C._SRS_PT_IGH)

// /** Gnomonic projection */
var SRSPTGnomonic = C.GoString(C._SRS_PT_GNOMONIC)

// /** Hotine_Oblique_Mercator_Azimuth_Center projection */
var SRSPTHotineObliqueMercatorAzimuthCenter = C.GoString(C._SRS_PT_HOTINE_OBLIQUE_MERCATOR_AZIMUTH_CENTER)

// /** Hotine_Oblique_Mercator projection */
var SRSPTHotineObliqueMercator = C.GoString(C._SRS_PT_HOTINE_OBLIQUE_MERCATOR)

// /** Hotine_Oblique_Mercator_Two_Point_Natural_Origin projection */
var SRSPTHotineObliqueMercatorTwoPointNaturalOrigin = C.GoString(C._SRS_PT_HOTINE_OBLIQUE_MERCATOR_TWO_POINT_NATURAL_ORIGIN)

// /** Laborde_Oblique_Mercator projection */
var SRSPTLabardeObliqueMercator = C.GoString(C._SRS_PT_LABORDE_OBLIQUE_MERCATOR)

// /** Lambert_Conformal_Conic_1SP projection */
var SRSPTLambertConformalConic1SP = C.GoString(C._SRS_PT_LAMBERT_CONFORMAL_CONIC_1SP)

// /** Lambert_Conformal_Conic_2SP projection */
var SRSPTLambertConformalConic2SP = C.GoString(C._SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP)

// /** Lambert_Conformal_Conic_2SP_Belgium projection */
var SRSPTLambertConformalConic2SPBelgium = C.GoString(C._SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP_BELGIUM)

// /** Lambert_Azimuthal_Equal_Area projection */
var SRSPTLambertAzimuthalEqualArea = C.GoString(C._SRS_PT_LAMBERT_AZIMUTHAL_EQUAL_AREA)

// /** Mercator_1SP projection */
var SRSPTMercator1SP = C.GoString(C._SRS_PT_MERCATOR_1SP)

// /** Mercator_2SP projection */
var SRSPTMercator2SP = C.GoString(C._SRS_PT_MERCATOR_2SP)

// /** Mercator_Auxiliary_Sphere is used used by ESRI to mean EPSG:3875 */
var SRSPTMercatorAuxiliarySphere = C.GoString(C._SRS_PT_MERCATOR_AUXILIARY_SPHERE)

// /** Miller_Cylindrical projection */
var SRSPTMillerCylindrical = C.GoString(C._SRS_PT_MILLER_CYLINDRICAL)

// /** Mollweide projection */
var SRSPTMollweide = C.GoString(C._SRS_PT_MOLLWEIDE)

// /** New_Zealand_Map_Grid projection */
var SRSPTNewZealandMapGrid = C.GoString(C._SRS_PT_NEW_ZEALAND_MAP_GRID)

// /** Oblique_Stereographic projection */
var SRSPTObliqueStereographic = C.GoString(C._SRS_PT_OBLIQUE_STEREOGRAPHIC)

// /** Orthographic projection */
var SRSPTOrthographic = C.GoString(C._SRS_PT_ORTHOGRAPHIC)

// /** Polar_Stereographic projection */
var SRSPTPolarStereographic = C.GoString(C._SRS_PT_POLAR_STEREOGRAPHIC)

// /** Polyconic projection */
var SRSPTPolyconic = C.GoString(C._SRS_PT_POLYCONIC)

// /** Robinson projection */
var SRSPTRobinson = C.GoString(C._SRS_PT_ROBINSON)

// /** Sinusoidal projection */
var SRSPTSinusoidal = C.GoString(C._SRS_PT_SINUSOIDAL)

// /** Stereographic projection */
var SRSPTStereographic = C.GoString(C._SRS_PT_STEREOGRAPHIC)

// /** Swiss_Oblique_Cylindrical projection */
var SRSPTSwissObliqueCylindrical = C.GoString(C._SRS_PT_SWISS_OBLIQUE_CYLINDRICAL)

// /** Transverse_Mercator projection */
var SRSPTTransverseMercator = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR)

// /** Transverse_Mercator_South_Orientated projection */
var SRSPTTransverseMercatorSouthOriented = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_SOUTH_ORIENTED)

// /* special mapinfo variants on Transverse Mercator */
// /** Transverse_Mercator_MapInfo_21 projection */
var SRSPTTransverseMercatorMI21 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_21)

// /** Transverse_Mercator_MapInfo_22 projection */
var SRSPTTransverseMercatorMI22 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_22)

// /** Transverse_Mercator_MapInfo_23 projection */
var SRSPTTransverseMercatorMI23 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_23)

// /** Transverse_Mercator_MapInfo_24 projection */
var SRSPTTransverseMercatorMI24 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_24)

// /** Transverse_Mercator_MapInfo_25 projection */
var SRSPTTransverseMercatorMI25 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_25)

// /** Tunisia_Mining_Grid projection */
var SRSPTTunisiaMiningGrid = C.GoString(C._SRS_PT_TUNISIA_MINING_GRID)

// /** Two_Point_Equidistant projection */
var SRSPTTwoPointEquidistant = C.GoString(C._SRS_PT_TWO_POINT_EQUIDISTANT)

// /** VanDerGrinten projection */
var SRSPTVandergrinten = C.GoString(C._SRS_PT_VANDERGRINTEN)

// /** Krovak projection */
var SRSPTKrovak = C.GoString(C._SRS_PT_KROVAK)

// /** International_Map_of_the_World_Polyconic projection */
var SRSPTIMWPolyconic = C.GoString(C._SRS_PT_IMW_POLYCONIC)

// /** Wagner_I projection */
var SRSPTWagnerI = C.GoString(C._SRS_PT_WAGNER_I)

// /** Wagner_II projection */
var SRSPTWagnerII = C.GoString(C._SRS_PT_WAGNER_II)

// /** Wagner_III projection */
var SRSPTWagnerIII = C.GoString(C._SRS_PT_WAGNER_III)

// /** Wagner_IV projection */
var SRSPTWagnerIV = C.GoString(C._SRS_PT_WAGNER_IV)

// /** Wagner_V projection */
var SRSPTWagnerV = C.GoString(C._SRS_PT_WAGNER_V)

// /** Wagner_VI projection */
var SRSPTWagnerVI = C.GoString(C._SRS_PT_WAGNER_VI)

// /** Wagner_VII projection */
var SRSPTWagnerVII = C.GoString(C._SRS_PT_WAGNER_VII)

// /** Quadrilateralized_Spherical_Cube projection */
var SRSPTQsc = C.GoString(C._SRS_PT_QSC)

// /** Aitoff projection */
var SRSPTAitoff = C.GoString(C._SRS_PT_AITOFF)

// /** Winkel_I projection */
var SRSPTWinkelI = C.GoString(C._SRS_PT_WINKEL_I)

// /** Winkel_II projection */
var SRSPTWinkelII = C.GoString(C._SRS_PT_WINKEL_II)

// /** Winkel_Tripel projection */
var SRSPTWinkelTripel = C.GoString(C._SRS_PT_WINKEL_TRIPEL)

// /** Craster_Parabolic projection */
var SRSPTCrasterParabolic = C.GoString(C._SRS_PT_CRASTER_PARABOLIC)

// /** Loximuthal projection */
var SRSPTLoximuthal = C.GoString(C._SRS_PT_LOXIMUTHAL)

// /** Quartic_Authalic projection */
var SRSPTQuarticAuthalic = C.GoString(C._SRS_PT_QUARTIC_AUTHALIC)

// /** Spherical_Cross_Track_Height projection */
var SRSPTSch = C.GoString(C._SRS_PT_SCH)

// /** central_meridian projection parameter */
var SRSPPCentralMeridian = C.GoString(C._SRS_PP_CENTRAL_MERIDIAN)

// /** scale_factor projection parameter */
var SRSPPScaleFactor = C.GoString(C._SRS_PP_SCALE_FACTOR)

// /** standard_parallel_1 projection parameter */
var SRSPPStandardParallel1 = C.GoString(C._SRS_PP_STANDARD_PARALLEL_1)

// /** standard_parallel_2 projection parameter */
var SRSPPStandardParallel2 = C.GoString(C._SRS_PP_STANDARD_PARALLEL_2)

// /** pseudo_standard_parallel_1 projection parameter */
var SRSPPPseudoStdParallel1 = C.GoString(C._SRS_PP_PSEUDO_STD_PARALLEL_1)

// /** longitude_of_center projection parameter */
var SRSPPLongitudeOfCenter = C.GoString(C._SRS_PP_LONGITUDE_OF_CENTER)

// /** latitude_of_center projection parameter */
var SRSPPLatitudeOfCenter = C.GoString(C._SRS_PP_LATITUDE_OF_CENTER)

// /** longitude_of_origin projection parameter */
var SRSPPLongitudeOfOrigin = C.GoString(C._SRS_PP_LONGITUDE_OF_ORIGIN)

// /** latitude_of_origin projection parameter */
var SRSPPLatitudeOfOrigin = C.GoString(C._SRS_PP_LATITUDE_OF_ORIGIN)

// /** false_easting projection parameter */
var SRSPPFalseEasting = C.GoString(C._SRS_PP_FALSE_EASTING)

// /** false_northing projection parameter */
var SRSPPFalseNorthing = C.GoString(C._SRS_PP_FALSE_NORTHING)

// /** azimuth projection parameter */
var SRSPPAzimuth = C.GoString(C._SRS_PP_AZIMUTH)

// /** longitude_of_point_1 projection parameter */
var SRSPPLongitudeOfPoint1 = C.GoString(C._SRS_PP_LONGITUDE_OF_POINT_1)

// /** latitude_of_point_1 projection parameter */
var SRSPPLatitudeOfPoint1 = C.GoString(C._SRS_PP_LATITUDE_OF_POINT_1)

// /** longitude_of_point_2 projection parameter */
var SRSPPLongitudeOfPoint2 = C.GoString(C._SRS_PP_LONGITUDE_OF_POINT_2)

// /** latitude_of_point_2 projection parameter */
var SRSPPLatitudeOfPoint2 = C.GoString(C._SRS_PP_LATITUDE_OF_POINT_2)

// /** longitude_of_point_3 projection parameter */
var SRSPPLongitudeOfPoint3 = C.GoString(C._SRS_PP_LONGITUDE_OF_POINT_3)

// /** latitude_of_point_3 projection parameter */
var SRSPPLatitudeOfPoint3 = C.GoString(C._SRS_PP_LATITUDE_OF_POINT_3)

// /** rectified_grid_angle projection parameter */
var SRSPPRectifiedGridAngle = C.GoString(C._SRS_PP_RECTIFIED_GRID_ANGLE)

// /** landsat_number projection parameter */
var SRSPPLandsatNumber = C.GoString(C._SRS_PP_LANDSAT_NUMBER)

// /** path_number projection parameter */
var SRSPPPathNumber = C.GoString(C._SRS_PP_PATH_NUMBER)

// /** perspective_point_height projection parameter */
var SRSPPPerspectivePointHeight = C.GoString(C._SRS_PP_PERSPECTIVE_POINT_HEIGHT)

// /** satellite_height projection parameter */
var SRSPPSatelliteHeight = C.GoString(C._SRS_PP_SATELLITE_HEIGHT)

// /** fipszone projection parameter */
var SRSPPFipszone = C.GoString(C._SRS_PP_FIPSZONE)

// /** zone projection parameter */
var SRSPPZone = C.GoString(C._SRS_PP_ZONE)

// /** Latitude_Of_1st_Point projection parameter */
var SRSPPLatitudeOf1stPoint = C.GoString(C._SRS_PP_LATITUDE_OF_1ST_POINT)

// /** Longitude_Of_1st_Point projection parameter */
var SRSPPLongitudeOf1stPoint = C.GoString(C._SRS_PP_LONGITUDE_OF_1ST_POINT)

// /** Latitude_Of_2nd_Point projection parameter */
var SRSPPLatitudeOf2ndPoint = C.GoString(C._SRS_PP_LATITUDE_OF_2ND_POINT)

// /** Longitude_Of_2nd_Point projection parameter */
var SRSPPLongitudeOf2ndPoint = C.GoString(C._SRS_PP_LONGITUDE_OF_2ND_POINT)

// /** peg_point_latitude projection parameter */
var SRSPPPegPointLatitude = C.GoString(C._SRS_PP_PEG_POINT_LATITUDE)

// /** peg_point_longitude projection parameter */
var SRSPPPegPointLongitude = C.GoString(C._SRS_PP_PEG_POINT_LONGITUDE)

// /** peg_point_heading projection parameter */
var SRSPPPegPointHeading = C.GoString(C._SRS_PP_PEG_POINT_HEADING)

// /** peg_point_height projection parameter */
var SRSPPPegPointHeight = C.GoString(C._SRS_PP_PEG_POINT_HEIGHT)

// /** Linear unit Meter */
var SRSULMeter = C.GoString(C._SRS_UL_METER)

// /** Linear unit Foot (International) */
var SRSULFoot = C.GoString(C._SRS_UL_FOOT)

// /** Linear unit Foot (International) conversion factor to meter*/
var SRSULFootConv = C.GoString(C._SRS_UL_FOOT_CONV)

// /** Linear unit Foot */
var SRSULUSFoot = C.GoString(C._SRS_UL_US_FOOT)

// /** Linear unit Foot conversion factor to meter */
var SRSULUSFootConv = C.GoString(C._SRS_UL_US_FOOT_CONV)

// /** Linear unit Nautical Mile */
var SRSULNauticalMile = C.GoString(C._SRS_UL_NAUTICAL_MILE)

// /** Linear unit Nautical Mile conversion factor to meter */
var SRSULNauticalMileConv = C.GoString(C._SRS_UL_NAUTICAL_MILE_CONV)

// /** Linear unit Link */
var SRSULLink = C.GoString(C._SRS_UL_LINK)

// /** Linear unit Link conversion factor to meter */
var SRSULLinkConv = C.GoString(C._SRS_UL_LINK_CONV)

// /** Linear unit Chain */
var SRSULChain = C.GoString(C._SRS_UL_CHAIN)

// /** Linear unit Chain conversion factor to meter */
var SRSULChainConv = C.GoString(C._SRS_UL_CHAIN_CONV)

// /** Linear unit Rod */
var SRSULRod = C.GoString(C._SRS_UL_ROD)

// /** Linear unit Rod conversion factor to meter */
var SRSULRodConv = C.GoString(C._SRS_UL_ROD_CONV)

// /** Linear unit Link_Clarke */
var SRSULLinkClarke = C.GoString(C._SRS_UL_LINK_Clarke)

// /** Linear unit Link_Clarke conversion factor to meter */
var SRSULLinkClarkeConv = C.GoString(C._SRS_UL_LINK_Clarke_CONV)

// /** Linear unit Kilometer */
var SRSULKilometer = C.GoString(C._SRS_UL_KILOMETER)

// /** Linear unit Kilometer conversion factor to meter */
var SRSULKilometerConv = C.GoString(C._SRS_UL_KILOMETER_CONV)

// /** Linear unit Decimeter */
var SRSULDecimeter = C.GoString(C._SRS_UL_DECIMETER)

// /** Linear unit Decimeter conversion factor to meter */
var SRSULDecimeterConv = C.GoString(C._SRS_UL_DECIMETER_CONV)

// /** Linear unit Decimeter */
var SRSULCentimeter = C.GoString(C._SRS_UL_CENTIMETER)

// /** Linear unit Decimeter conversion factor to meter */
var SRSULCentimeterConv = C.GoString(C._SRS_UL_CENTIMETER_CONV)

// /** Linear unit Millimeter */
var SRSULMillimeter = C.GoString(C._SRS_UL_MILLIMETER)

// /** Linear unit Millimeter conversion factor to meter */
var SRSULMillimeterConv = C.GoString(C._SRS_UL_MILLIMETER_CONV)

// /** Linear unit Nautical_Mile_International */
var SRSULIntlNautMile = C.GoString(C._SRS_UL_INTL_NAUT_MILE)

// /** Linear unit Nautical_Mile_International conversion factor to meter */
var SRSULIntlNautMileConv = C.GoString(C._SRS_UL_INTL_NAUT_MILE_CONV)

// /** Linear unit Inch_International */
var SRSULIntlInch = C.GoString(C._SRS_UL_INTL_INCH)

// /** Linear unit Inch_International conversion factor to meter */
var SRSULIntlInchConv = C.GoString(C._SRS_UL_INTL_INCH_CONV)

// /** Linear unit Foot_International */
var SRSULIntlFoot = C.GoString(C._SRS_UL_INTL_FOOT)

// /** Linear unit Foot_International conversion factor to meter */
var SRSULIntlFootConv = C.GoString(C._SRS_UL_INTL_FOOT_CONV)

// /** Linear unit Yard_International */
var SRSULIntlYard = C.GoString(C._SRS_UL_INTL_YARD)

// /** Linear unit Yard_International conversion factor to meter */
var SRSULIntlYardConv = C.GoString(C._SRS_UL_INTL_YARD_CONV)

// /** Linear unit Statute_Mile_International */
var SRSULIntlStatMile = C.GoString(C._SRS_UL_INTL_STAT_MILE)

// /** Linear unit Statute_Mile_Internationalconversion factor to meter */
var SRSULIntlStatMileConv = C.GoString(C._SRS_UL_INTL_STAT_MILE_CONV)

// /** Linear unit Fathom_International */
var SRSULIntlFathom = C.GoString(C._SRS_UL_INTL_FATHOM)

// /** Linear unit Fathom_International conversion factor to meter */
var SRSULIntlFathomConv = C.GoString(C._SRS_UL_INTL_FATHOM_CONV)

// /** Linear unit Chain_International */
var SRSULIntlChain = C.GoString(C._SRS_UL_INTL_CHAIN)

// /** Linear unit Chain_International conversion factor to meter */
var SRSULIntlChainConv = C.GoString(C._SRS_UL_INTL_CHAIN_CONV)

// /** Linear unit Link_International */
var SRSULIntlLink = C.GoString(C._SRS_UL_INTL_LINK)

// /** Linear unit Link_International conversion factor to meter */
var SRSULIntlLinkConv = C.GoString(C._SRS_UL_INTL_LINK_CONV)

// /** Linear unit Inch_US_Surveyor */
var SRSULUSInch = C.GoString(C._SRS_UL_US_INCH)

// /** Linear unit Inch_US_Surveyor conversion factor to meter */
var SRSULUSInchConv = C.GoString(C._SRS_UL_US_INCH_CONV)

// /** Linear unit Yard_US_Surveyor */
var SRSULUSYard = C.GoString(C._SRS_UL_US_YARD)

// /** Linear unit Yard_US_Surveyor conversion factor to meter */
var SRSULUSYardConv = C.GoString(C._SRS_UL_US_YARD_CONV)

// /** Linear unit Chain_US_Surveyor */
var SRSULUSChain = C.GoString(C._SRS_UL_US_CHAIN)

// /** Linear unit Chain_US_Surveyor conversion factor to meter */
var SRSULUSChainConv = C.GoString(C._SRS_UL_US_CHAIN_CONV)

// /** Linear unit Statute_Mile_US_Surveyor */
var SRSULUSStatMile = C.GoString(C._SRS_UL_US_STAT_MILE)

// /** Linear unit Statute_Mile_US_Surveyor conversion factor to meter */
var SRSULUSStatMileConv = C.GoString(C._SRS_UL_US_STAT_MILE_CONV)

// /** Linear unit Yard_Indian */
var SRSULIndianYard = C.GoString(C._SRS_UL_INDIAN_YARD)

// /** Linear unit Yard_Indian conversion factor to meter */
var SRSULIndianYardConv = C.GoString(C._SRS_UL_INDIAN_YARD_CONV)

// /** Linear unit Foot_Indian */
var SRSULIndianFoot = C.GoString(C._SRS_UL_INDIAN_FOOT)

// /** Linear unit Foot_Indian conversion factor to meter */
var SRSULIndianFootConv = C.GoString(C._SRS_UL_INDIAN_FOOT_CONV)

// /** Linear unit Chain_Indian */
var SRSULIndianChain = C.GoString(C._SRS_UL_INDIAN_CHAIN)

// /** Linear unit Chain_Indian conversion factor to meter */
var SRSULIndianChainConv = C.GoString(C._SRS_UL_INDIAN_CHAIN_CONV)

// /** Angular unit degree */
var SRSUADegree = C.GoString(C._SRS_UA_DEGREE)

// /** Angular unit degree conversion factor to radians */
var SRSUADegreeConv = C.GoString(C._SRS_UA_DEGREE_CONV)

// /** Angular unit radian */
var SRSUARadian = C.GoString(C._SRS_UA_RADIAN)

// /** Prime meridian Greenwich */
var SRSPMGreenwich = C.GoString(C._SRS_PM_GREENWICH)

// /** North_American_Datum_1927 datum name */
var SRSDNNad27 = C.GoString(C._SRS_DN_NAD27)

// /** North_American_Datum_1983 datum name */
var SRSDNNad83 = C.GoString(C._SRS_DN_NAD83)

// /** WGS_1972 datum name */
var SRSDNWgs72 = C.GoString(C._SRS_DN_WGS72)

// /** WGS_1984 datum name */
var SRSDNWgs84 = C.GoString(C._SRS_DN_WGS84)

// /** Semi-major axis of the WGS84 ellipsoid */
const SRSWgs84SemiMajor = C.SRS_WGS84_SEMIMAJOR

// /** Inverse flattening of the WGS84 ellipsoid */
const SRSWgs84InvFlattening = C.SRS_WGS84_INVFLATTENING

// #ifndef SWIG
// /* -------------------------------------------------------------------- */
// /*      C Wrappers for C++ objects and methods.                         */
// /* -------------------------------------------------------------------- */

func osrSetPROJSearchPaths(paths []string) {
	cPaths := make([]*C.char, len(paths)+1)
	for i, p := range paths {
		cPaths[i] = C.CString(p)
		defer C.free(unsafe.Pointer(cPaths[i]))
	}
	cPaths[len(paths)] = nil
	C.OSRSetPROJSearchPaths(&cPaths[0])
}

func OSRSetPROJSearchPaths(paths []string) {
	osrSetPROJSearchPaths(paths)
}

func osrGetPROJSearchPaths() (result []string) {
	raw := C.OSRGetPROJSearchPaths()
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	for i := 0; ; i++ {
		p := C.CSLGetField(raw, C.int(i))
		if p == nil {
			break
		}
		result = append(result, C.GoString(p))
	}
	return
}

func OSRGetPROJSearchPaths() (result []string) {
	result = osrGetPROJSearchPaths()
	return
}

func osrSetPROJAuxDbPaths(paths []string) {
	cPaths := make([]*C.char, len(paths)+1)
	for i, p := range paths {
		cPaths[i] = C.CString(p)
		defer C.free(unsafe.Pointer(cPaths[i]))
	}
	cPaths[len(paths)] = nil
	C.OSRSetPROJAuxDbPaths(&cPaths[0])
}

func OSRSetPROJAuxDbPaths(paths []string) {
	osrSetPROJAuxDbPaths(paths)
}

func osrGetPROJAuxDbPaths() (result []string) {
	raw := C.OSRGetPROJAuxDbPaths()
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	for i := 0; ; i++ {
		p := C.CSLGetField(raw, C.int(i))
		if p == nil {
			break
		}
		result = append(result, C.GoString(p))
	}
	return
}

func OSRGetPROJAuxDbPaths() (result []string) {
	result = osrGetPROJAuxDbPaths()
	return
}

func osrSetPROJEnableNetwork(enabled int) {
	C.OSRSetPROJEnableNetwork(C.int(enabled))
}

func OSRSetPROJEnableNetwork(enabled int) {
	osrSetPROJEnableNetwork(enabled)
}

func osrGetPROJEnableNetwork() (result int) {
	result = int(C.OSRGetPROJEnableNetwork())
	return
}

func OSRGetPROJEnableNetwork() (result int) {
	result = osrGetPROJEnableNetwork()
	return
}

func osrGetPROJVersion() (major, minor, patch int) {
	var cMajor, cMinor, cPatch C.int
	C.OSRGetPROJVersion(&cMajor, &cMinor, &cPatch)
	major = int(cMajor)
	minor = int(cMinor)
	patch = int(cPatch)
	return
}

func OSRGetPROJVersion() (major, minor, patch int) {
	major, minor, patch = osrGetPROJVersion()
	return
}

func osrNewSpatialReference(wkt string) (result OGRSpatialReference) {
	var cs *C.char
	if wkt != "" {
		cs = C.CString(wkt)
		defer C.free(unsafe.Pointer(cs))
	}
	result = OGRSpatialReference{cValue: C.OSRNewSpatialReference(cs)}
	return
}

func OSRNewSpatialReference(wkt string) (result OGRSpatialReference, err error) {
	result = osrNewSpatialReference(wkt)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func osrCloneGeogCS(sr OGRSpatialReference) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OSRCloneGeogCS(sr.cValue)}
	return
}

func (sr OGRSpatialReference) CloneGeogCS() (result OGRSpatialReference, err error) {
	result = osrCloneGeogCS(sr)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func osrClone(sr OGRSpatialReference) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OSRClone(sr.cValue)}
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
	C.OSRDestroySpatialReference(sr.cValue)
}

func osrReference(sr OGRSpatialReference) (result int) {
	result = int(C.OSRReference(sr.cValue))
	return
}

func (sr OGRSpatialReference) Reference() (result int) {
	result = osrReference(sr)
	return
}

func osrDereference(sr OGRSpatialReference) (result int) {
	result = int(C.OSRDereference(sr.cValue))
	return
}

func (sr OGRSpatialReference) Dereference() (result int) {
	result = osrDereference(sr)
	return
}

func (sr OGRSpatialReference) Release() {
	C.OSRRelease(sr.cValue)
}

func osrValidate(sr OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRValidate(sr.cValue)))
	return
}

func (sr OGRSpatialReference) Validate() (err error) {
	err = osrValidate(sr)
	return
}

func osrImportFromEPSG(sr OGRSpatialReference, nCode int) (err error) {
	err = ogrError(OGRErr(C.OSRImportFromEPSG(sr.cValue, C.int(nCode))))
	return
}

func (sr OGRSpatialReference) ImportFromEPSG(nCode int) (err error) {
	err = osrImportFromEPSG(sr, nCode)
	return
}

func osrImportFromEPSGA(sr OGRSpatialReference, nCode int) (err error) {
	err = ogrError(OGRErr(C.OSRImportFromEPSGA(sr.cValue, C.int(nCode))))
	return
}

func (sr OGRSpatialReference) ImportFromEPSGA(nCode int) (err error) {
	err = osrImportFromEPSGA(sr, nCode)
	return
}

func osrImportFromWkt(sr OGRSpatialReference, wkt string) (err error) {
	cs := C.CString(wkt)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRImportFromWkt(sr.cValue, &cs)))
	return
}

func (sr OGRSpatialReference) ImportFromWkt(wkt string) (err error) {
	err = osrImportFromWkt(sr, wkt)
	return
}

func osrImportFromProj4(sr OGRSpatialReference, proj4 string) (err error) {
	cs := C.CString(proj4)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRImportFromProj4(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) ImportFromProj4(proj4 string) (err error) {
	err = osrImportFromProj4(sr, proj4)
	return
}

func osrImportFromESRI(sr OGRSpatialReference, lines []string) (err error) {
	cLines := make([]*C.char, len(lines)+1)
	for i, l := range lines {
		cLines[i] = C.CString(l)
		defer C.free(unsafe.Pointer(cLines[i]))
	}
	cLines[len(lines)] = nil
	err = ogrError(OGRErr(C.OSRImportFromESRI(sr.cValue, &cLines[0])))
	return
}

func (sr OGRSpatialReference) ImportFromESRI(lines []string) (err error) {
	err = osrImportFromESRI(sr, lines)
	return
}

func osrImportFromPCI(sr OGRSpatialReference, proj, units string, arParams []float64) (err error) {
	csProj := C.CString(proj)
	defer C.free(unsafe.Pointer(csProj))
	csUnits := C.CString(units)
	defer C.free(unsafe.Pointer(csUnits))
	var pParams *C.double
	if len(arParams) > 0 {
		pParams = (*C.double)(unsafe.Pointer(&arParams[0]))
	}
	err = ogrError(OGRErr(C.OSRImportFromPCI(sr.cValue, csProj, csUnits, pParams)))
	return
}

func (sr OGRSpatialReference) ImportFromPCI(proj, units string, arParams []float64) (err error) {
	err = osrImportFromPCI(sr, proj, units, arParams)
	return
}

func osrImportFromUSGS(sr OGRSpatialReference, projSys, zone int, arParams []float64, datum int) (err error) {
	var pParams *C.double
	if len(arParams) > 0 {
		pParams = (*C.double)(unsafe.Pointer(&arParams[0]))
	}
	err = ogrError(OGRErr(C.OSRImportFromUSGS(sr.cValue, C.long(projSys), C.long(zone), pParams, C.long(datum))))
	return
}

func (sr OGRSpatialReference) ImportFromUSGS(projSys, zone int, arParams []float64, datum int) (err error) {
	err = osrImportFromUSGS(sr, projSys, zone, arParams, datum)
	return
}

func osrImportFromXML(sr OGRSpatialReference, xmlString string) (err error) {
	cs := C.CString(xmlString)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRImportFromXML(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) ImportFromXML(xmlString string) (err error) {
	err = osrImportFromXML(sr, xmlString)
	return
}

func osrImportFromDict(sr OGRSpatialReference, dictFile, code string) (err error) {
	csDictFile := C.CString(dictFile)
	defer C.free(unsafe.Pointer(csDictFile))
	csCode := C.CString(code)
	defer C.free(unsafe.Pointer(csCode))
	err = ogrError(OGRErr(C.OSRImportFromDict(sr.cValue, csDictFile, csCode)))
	return
}

func (sr OGRSpatialReference) ImportFromDict(dictFile, code string) (err error) {
	err = osrImportFromDict(sr, dictFile, code)
	return
}

func osrImportFromPanorama(sr OGRSpatialReference, projSys, datum, ellipsoid int, arParams []float64) (err error) {
	var pParams *C.double
	if len(arParams) > 0 {
		pParams = (*C.double)(unsafe.Pointer(&arParams[0]))
	}
	err = ogrError(OGRErr(C.OSRImportFromPanorama(sr.cValue, C.long(projSys), C.long(datum), C.long(ellipsoid), pParams)))
	return
}

func (sr OGRSpatialReference) ImportFromPanorama(projSys, datum, ellipsoid int, arParams []float64) (err error) {
	err = osrImportFromPanorama(sr, projSys, datum, ellipsoid, arParams)
	return
}

func osrImportFromOzi(sr OGRSpatialReference, lines []string) (err error) {
	cLines := make([]*C.char, len(lines)+1)
	for i, l := range lines {
		cLines[i] = C.CString(l)
		defer C.free(unsafe.Pointer(cLines[i]))
	}
	cLines[len(lines)] = nil
	err = ogrError(OGRErr(C.OSRImportFromOzi(sr.cValue, &cLines[0])))
	return
}

func (sr OGRSpatialReference) ImportFromOzi(lines []string) (err error) {
	err = osrImportFromOzi(sr, lines)
	return
}

func osrImportFromMICoordSys(sr OGRSpatialReference, coordSys string) (err error) {
	cs := C.CString(coordSys)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRImportFromMICoordSys(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) ImportFromMICoordSys(coordSys string) (err error) {
	err = osrImportFromMICoordSys(sr, coordSys)
	return
}

func osrImportFromERM(sr OGRSpatialReference, proj, datum, units string) (err error) {
	csProj := C.CString(proj)
	defer C.free(unsafe.Pointer(csProj))
	csDatum := C.CString(datum)
	defer C.free(unsafe.Pointer(csDatum))
	csUnits := C.CString(units)
	defer C.free(unsafe.Pointer(csUnits))
	err = ogrError(OGRErr(C.OSRImportFromERM(sr.cValue, csProj, csDatum, csUnits)))
	return
}

func (sr OGRSpatialReference) ImportFromERM(proj, datum, units string) (err error) {
	err = osrImportFromERM(sr, proj, datum, units)
	return
}

func osrImportFromUrl(sr OGRSpatialReference, url string) (err error) {
	cs := C.CString(url)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRImportFromUrl(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) ImportFromUrl(url string) (err error) {
	err = osrImportFromUrl(sr, url)
	return
}

func osrImportFromCF1(sr OGRSpatialReference, keyValues []string, units string) (err error) {
	cKeyValues := make([]*C.char, len(keyValues)+1)
	for i, kv := range keyValues {
		cKeyValues[i] = C.CString(kv)
		defer C.free(unsafe.Pointer(cKeyValues[i]))
	}
	cKeyValues[len(keyValues)] = nil
	csUnits := C.CString(units)
	defer C.free(unsafe.Pointer(csUnits))
	err = ogrError(OGRErr(C.OSRImportFromCF1(sr.cValue, &cKeyValues[0], csUnits)))
	return
}

func (sr OGRSpatialReference) ImportFromCF1(keyValues []string, units string) (err error) {
	err = osrImportFromCF1(sr, keyValues, units)
	return
}

func osrExportToWkt(sr OGRSpatialReference) (result string, err error) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToWkt(sr.cValue, &p))
	defer C.CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToWkt() (result string, err error) {
	result, err = osrExportToWkt(sr)
	return
}

func osrExportToWktEx(sr OGRSpatialReference, options []string) (result string, err error) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToWktEx(sr.cValue, &p, &cOptions[0]))
	defer C.CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToWktEx(options []string) (result string, err error) {
	result, err = osrExportToWktEx(sr, options)
	return
}

func osrExportToPrettyWkt(sr OGRSpatialReference, simplify int) (result string, err error) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToPrettyWkt(sr.cValue, &p, C.int(simplify)))
	defer C.CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToPrettyWkt(simplify int) (result string, err error) {
	result, err = osrExportToPrettyWkt(sr, simplify)
	return
}

func osrExportToPROJJSON(sr OGRSpatialReference, options []string) (result string, err error) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToPROJJSON(sr.cValue, &p, &cOptions[0]))
	defer C.CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToPROJJSON(options []string) (result string, err error) {
	result, err = osrExportToPROJJSON(sr, options)
	return
}

func osrExportToProj4(sr OGRSpatialReference) (result string, err error) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToProj4(sr.cValue, &p))
	defer C.CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToProj4() (result string, err error) {
	result, err = osrExportToProj4(sr)
	return
}

func osrExportToPCI(sr OGRSpatialReference) (proj, units string, params []float64, err error) {
	var cProj, cUnits *C.char
	var cParams *C.double
	ogrErr := OGRErr(C.OSRExportToPCI(sr.cValue, &cProj, &cUnits, &cParams))
	defer C.CPLFree(unsafe.Pointer(cProj))
	defer C.CPLFree(unsafe.Pointer(cUnits))
	defer C.CPLFree(unsafe.Pointer(cParams))
	proj = C.GoString(cProj)
	units = C.GoString(cUnits)
	const nPCIParams = 17
	params = make([]float64, nPCIParams)
	for i := 0; i < nPCIParams; i++ {
		params[i] = float64(*(*C.double)(unsafe.Add(unsafe.Pointer(cParams), uintptr(i)*unsafe.Sizeof(*cParams))))
	}
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToPCI() (proj, units string, params []float64, err error) {
	proj, units, params, err = osrExportToPCI(sr)
	return
}

func osrExportToUSGS(sr OGRSpatialReference) (projSys, zone int, params []float64, datum int, err error) {
	var cProjSys, cZone, cDatum C.long
	var cParams *C.double
	ogrErr := OGRErr(C.OSRExportToUSGS(sr.cValue, &cProjSys, &cZone, &cParams, &cDatum))
	defer C.CPLFree(unsafe.Pointer(cParams))
	projSys = int(cProjSys)
	zone = int(cZone)
	datum = int(cDatum)
	const nUSGSParams = 15
	params = make([]float64, nUSGSParams)
	for i := 0; i < nUSGSParams; i++ {
		params[i] = float64(*(*C.double)(unsafe.Add(unsafe.Pointer(cParams), uintptr(i)*unsafe.Sizeof(*cParams))))
	}
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToUSGS() (projSys, zone int, params []float64, datum int, err error) {
	projSys, zone, params, datum, err = osrExportToUSGS(sr)
	return
}

func osrExportToXML(sr OGRSpatialReference, dialect string) (result string, err error) {
	var cs *C.char
	if dialect != "" {
		cs = C.CString(dialect)
		defer C.free(unsafe.Pointer(cs))
	}
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToXML(sr.cValue, &p, cs))
	defer C.CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToXML(dialect string) (result string, err error) {
	result, err = osrExportToXML(sr, dialect)
	return
}

func osrExportToPanorama(sr OGRSpatialReference) (projSys, datum, ellipsoid, zone int, params []float64, err error) {
	var cProjSys, cDatum, cEllipsoid, cZone C.long
	const nPanoramaParams = 7
	cParams := make([]C.double, nPanoramaParams)
	ogrErr := OGRErr(C.OSRExportToPanorama(sr.cValue, &cProjSys, &cDatum, &cEllipsoid, &cZone, &cParams[0]))
	projSys = int(cProjSys)
	datum = int(cDatum)
	ellipsoid = int(cEllipsoid)
	zone = int(cZone)
	params = make([]float64, nPanoramaParams)
	for i := 0; i < nPanoramaParams; i++ {
		params[i] = float64(cParams[i])
	}
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToPanorama() (projSys, datum, ellipsoid, zone int, params []float64, err error) {
	projSys, datum, ellipsoid, zone, params, err = osrExportToPanorama(sr)
	return
}

func osrExportToMICoordSys(sr OGRSpatialReference) (result string, err error) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToMICoordSys(sr.cValue, &p))
	defer C.CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToMICoordSys() (result string, err error) {
	result, err = osrExportToMICoordSys(sr)
	return
}

func osrExportToERM(sr OGRSpatialReference) (proj, datum, units string, err error) {
	cProj := make([]C.char, 32)
	cDatum := make([]C.char, 32)
	cUnits := make([]C.char, 32)
	ogrErr := OGRErr(C.OSRExportToERM(sr.cValue, &cProj[0], &cDatum[0], &cUnits[0]))
	proj = C.GoString(&cProj[0])
	datum = C.GoString(&cDatum[0])
	units = C.GoString(&cUnits[0])
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToERM() (proj, datum, units string, err error) {
	proj, datum, units, err = osrExportToERM(sr)
	return
}

func osrExportToCF1(sr OGRSpatialReference, options []string) (gridMappingName string, keyValues []string, units string, err error) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	var cGridMappingName *C.char
	var cKeyValues **C.char
	var cUnits *C.char
	ogrErr := OGRErr(C.OSRExportToCF1(sr.cValue, &cGridMappingName, &cKeyValues, &cUnits, &cOptions[0]))
	defer C.CPLFree(unsafe.Pointer(cGridMappingName))
	defer C.CSLDestroy(cKeyValues)
	defer C.CPLFree(unsafe.Pointer(cUnits))
	gridMappingName = C.GoString(cGridMappingName)
	units = C.GoString(cUnits)
	for i := 0; ; i++ {
		p := C.CSLGetField(cKeyValues, C.int(i))
		if p == nil {
			break
		}
		keyValues = append(keyValues, C.GoString(p))
	}
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) ExportToCF1(options []string) (gridMappingName string, keyValues []string, units string, err error) {
	gridMappingName, keyValues, units, err = osrExportToCF1(sr, options)
	return
}

func osrMorphToESRI(sr OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRMorphToESRI(sr.cValue)))
	return
}

func (sr OGRSpatialReference) MorphToESRI() (err error) {
	err = osrMorphToESRI(sr)
	return
}

func osrMorphFromESRI(sr OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRMorphFromESRI(sr.cValue)))
	return
}

func (sr OGRSpatialReference) MorphFromESRI() (err error) {
	err = osrMorphFromESRI(sr)
	return
}

func osrStripVertical(sr OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRStripVertical(sr.cValue)))
	return
}

func (sr OGRSpatialReference) StripVertical() (err error) {
	err = osrStripVertical(sr)
	return
}

func osrConvertToOtherProjection(sr OGRSpatialReference, targetProjection string, options []string) (result OGRSpatialReference) {
	cs := C.CString(targetProjection)
	defer C.free(unsafe.Pointer(cs))
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRSpatialReference{cValue: C.OSRConvertToOtherProjection(sr.cValue, cs, &cOptions[0])}
	return
}

func (sr OGRSpatialReference) ConvertToOtherProjection(targetProjection string, options []string) (result OGRSpatialReference, err error) {
	result = osrConvertToOtherProjection(sr, targetProjection, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func osrGetName(sr OGRSpatialReference) (result string) {
	result = C.GoString(C.OSRGetName(sr.cValue))
	return
}

func (sr OGRSpatialReference) GetName() (result string) {
	result = osrGetName(sr)
	return
}

func osrGetCelestialBodyName(sr OGRSpatialReference) (result string) {
	result = C.GoString(C.OSRGetCelestialBodyName(sr.cValue))
	return
}

func (sr OGRSpatialReference) GetCelestialBodyName() (result string) {
	result = osrGetCelestialBodyName(sr)
	return
}

func osrSetAttrValue(sr OGRSpatialReference, nodePath, newValue string) (err error) {
	csPath := C.CString(nodePath)
	defer C.free(unsafe.Pointer(csPath))
	csValue := C.CString(newValue)
	defer C.free(unsafe.Pointer(csValue))
	err = ogrError(OGRErr(C.OSRSetAttrValue(sr.cValue, csPath, csValue)))
	return
}

func (sr OGRSpatialReference) SetAttrValue(nodePath, newValue string) (err error) {
	err = osrSetAttrValue(sr, nodePath, newValue)
	return
}

func osrGetAttrValue(sr OGRSpatialReference, name string, iChild int) (result string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OSRGetAttrValue(sr.cValue, cs, C.int(iChild)))
	return
}

func (sr OGRSpatialReference) GetAttrValue(name string, iChild int) (result string) {
	result = osrGetAttrValue(sr, name, iChild)
	return
}

func osrSetAngularUnits(sr OGRSpatialReference, name string, inDegrees float64) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetAngularUnits(sr.cValue, cs, C.double(inDegrees))))
	return
}

func (sr OGRSpatialReference) SetAngularUnits(name string, inDegrees float64) (err error) {
	err = osrSetAngularUnits(sr, name, inDegrees)
	return
}

func osrGetAngularUnits(sr OGRSpatialReference) (result float64, name string) {
	var pName *C.char
	result = float64(C.OSRGetAngularUnits(sr.cValue, &pName))
	name = C.GoString(pName)
	return
}

func (sr OGRSpatialReference) GetAngularUnits() (result float64, name string) {
	result, name = osrGetAngularUnits(sr)
	return
}

func osrSetLinearUnits(sr OGRSpatialReference, name string, inMeters float64) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetLinearUnits(sr.cValue, cs, C.double(inMeters))))
	return
}

func (sr OGRSpatialReference) SetLinearUnits(name string, inMeters float64) (err error) {
	err = osrSetLinearUnits(sr, name, inMeters)
	return
}

func osrSetTargetLinearUnits(sr OGRSpatialReference, targetKey, name string, inMeters float64) (err error) {
	csKey := C.CString(targetKey)
	defer C.free(unsafe.Pointer(csKey))
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	err = ogrError(OGRErr(C.OSRSetTargetLinearUnits(sr.cValue, csKey, csName, C.double(inMeters))))
	return
}

func (sr OGRSpatialReference) SetTargetLinearUnits(targetKey, name string, inMeters float64) (err error) {
	err = osrSetTargetLinearUnits(sr, targetKey, name, inMeters)
	return
}

func osrSetLinearUnitsAndUpdateParameters(sr OGRSpatialReference, name string, inMeters float64) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetLinearUnitsAndUpdateParameters(sr.cValue, cs, C.double(inMeters))))
	return
}

func (sr OGRSpatialReference) SetLinearUnitsAndUpdateParameters(name string, inMeters float64) (err error) {
	err = osrSetLinearUnitsAndUpdateParameters(sr, name, inMeters)
	return
}

func osrGetLinearUnits(sr OGRSpatialReference) (result float64, name string) {
	var pName *C.char
	result = float64(C.OSRGetLinearUnits(sr.cValue, &pName))
	name = C.GoString(pName)
	return
}

func (sr OGRSpatialReference) GetLinearUnits() (result float64, name string) {
	result, name = osrGetLinearUnits(sr)
	return
}

func osrGetTargetLinearUnits(sr OGRSpatialReference, targetKey string) (result float64, name string) {
	csKey := C.CString(targetKey)
	defer C.free(unsafe.Pointer(csKey))
	var pName *C.char
	result = float64(C.OSRGetTargetLinearUnits(sr.cValue, csKey, &pName))
	name = C.GoString(pName)
	return
}

func (sr OGRSpatialReference) GetTargetLinearUnits(targetKey string) (result float64, name string) {
	result, name = osrGetTargetLinearUnits(sr, targetKey)
	return
}

func osrGetPrimeMeridian(sr OGRSpatialReference) (result float64, name string) {
	var pName *C.char
	result = float64(C.OSRGetPrimeMeridian(sr.cValue, &pName))
	name = C.GoString(pName)
	return
}

func (sr OGRSpatialReference) GetPrimeMeridian() (result float64, name string) {
	result, name = osrGetPrimeMeridian(sr)
	return
}

func osrIsGeographic(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsGeographic(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsGeographic() (result bool) {
	result = osrIsGeographic(sr)
	return
}

func osrIsDerivedGeographic(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsDerivedGeographic(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsDerivedGeographic() (result bool) {
	result = osrIsDerivedGeographic(sr)
	return
}

func osrIsLocal(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsLocal(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsLocal() (result bool) {
	result = osrIsLocal(sr)
	return
}

func osrIsProjected(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsProjected(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsProjected() (result bool) {
	result = osrIsProjected(sr)
	return
}

func osrIsDerivedProjected(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsDerivedProjected(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsDerivedProjected() (result bool) {
	result = osrIsDerivedProjected(sr)
	return
}

func osrIsCompound(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsCompound(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsCompound() (result bool) {
	result = osrIsCompound(sr)
	return
}

func osrIsGeocentric(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsGeocentric(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsGeocentric() (result bool) {
	result = osrIsGeocentric(sr)
	return
}

func osrIsVertical(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsVertical(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsVertical() (result bool) {
	result = osrIsVertical(sr)
	return
}

func osrIsDynamic(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsDynamic(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsDynamic() (result bool) {
	result = osrIsDynamic(sr)
	return
}

func osrHasPointMotionOperation(sr OGRSpatialReference) (result bool) {
	result = C.OSRHasPointMotionOperation(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) HasPointMotionOperation() (result bool) {
	result = osrHasPointMotionOperation(sr)
	return
}

func osrIsSameGeogCS(sr, other OGRSpatialReference) (result bool) {
	result = C.OSRIsSameGeogCS(sr.cValue, other.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsSameGeogCS(other OGRSpatialReference) (result bool) {
	result = osrIsSameGeogCS(sr, other)
	return
}

func osrIsSameVertCS(sr, other OGRSpatialReference) (result bool) {
	result = C.OSRIsSameVertCS(sr.cValue, other.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsSameVertCS(other OGRSpatialReference) (result bool) {
	result = osrIsSameVertCS(sr, other)
	return
}

func osrIsSame(sr, other OGRSpatialReference) (result bool) {
	result = C.OSRIsSame(sr.cValue, other.cValue) != 0
	return
}

func (sr OGRSpatialReference) IsSame(other OGRSpatialReference) (result bool) {
	result = osrIsSame(sr, other)
	return
}

func osrIsSameEx(sr, other OGRSpatialReference, options []string) (result bool) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = C.OSRIsSameEx(sr.cValue, other.cValue, &cOptions[0]) != 0
	return
}

func (sr OGRSpatialReference) IsSameEx(other OGRSpatialReference, options []string) (result bool) {
	result = osrIsSameEx(sr, other, options)
	return
}

func osrSetCoordinateEpoch(sr OGRSpatialReference, epoch float64) {
	C.OSRSetCoordinateEpoch(sr.cValue, C.double(epoch))
}

func (sr OGRSpatialReference) SetCoordinateEpoch(epoch float64) {
	osrSetCoordinateEpoch(sr, epoch)
}

func osrGetCoordinateEpoch(sr OGRSpatialReference) (result float64) {
	result = float64(C.OSRGetCoordinateEpoch(sr.cValue))
	return
}

func (sr OGRSpatialReference) GetCoordinateEpoch() (result float64) {
	result = osrGetCoordinateEpoch(sr)
	return
}

func osrSetLocalCS(sr OGRSpatialReference, name string) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetLocalCS(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) SetLocalCS(name string) (err error) {
	err = osrSetLocalCS(sr, name)
	return
}

func osrSetProjCS(sr OGRSpatialReference, name string) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetProjCS(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) SetProjCS(name string) (err error) {
	err = osrSetProjCS(sr, name)
	return
}

func osrSetGeocCS(sr OGRSpatialReference, name string) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetGeocCS(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) SetGeocCS(name string) (err error) {
	err = osrSetGeocCS(sr, name)
	return
}

func osrSetWellKnownGeogCS(sr OGRSpatialReference, name string) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetWellKnownGeogCS(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) SetWellKnownGeogCS(name string) (err error) {
	err = osrSetWellKnownGeogCS(sr, name)
	return
}

func osrSetFromUserInput(sr OGRSpatialReference, definition string) (err error) {
	cs := C.CString(definition)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetFromUserInput(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) SetFromUserInput(definition string) (err error) {
	err = osrSetFromUserInput(sr, definition)
	return
}

func osrSetFromUserInputEx(sr OGRSpatialReference, definition string, options []string) (err error) {
	cs := C.CString(definition)
	defer C.free(unsafe.Pointer(cs))
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	err = ogrError(OGRErr(C.OSRSetFromUserInputEx(sr.cValue, cs, &cOptions[0])))
	return
}

func (sr OGRSpatialReference) SetFromUserInputEx(definition string, options []string) (err error) {
	err = osrSetFromUserInputEx(sr, definition, options)
	return
}

func osrCopyGeogCSFrom(sr, src OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRCopyGeogCSFrom(sr.cValue, src.cValue)))
	return
}

func (sr OGRSpatialReference) CopyGeogCSFrom(src OGRSpatialReference) (err error) {
	err = osrCopyGeogCSFrom(sr, src)
	return
}

func osrSetTOWGS84(sr OGRSpatialReference, dx, dy, dz, ex, ey, ez, ppm float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetTOWGS84(sr.cValue, C.double(dx), C.double(dy), C.double(dz), C.double(ex), C.double(ey), C.double(ez), C.double(ppm))))
	return
}

func (sr OGRSpatialReference) SetTOWGS84(dx, dy, dz, ex, ey, ez, ppm float64) (err error) {
	err = osrSetTOWGS84(sr, dx, dy, dz, ex, ey, ez, ppm)
	return
}

func osrGetTOWGS84(sr OGRSpatialReference) (params []float64, err error) {
	const n = 7
	cParams := make([]C.double, n)
	ogrErr := OGRErr(C.OSRGetTOWGS84(sr.cValue, &cParams[0], C.int(n)))
	params = make([]float64, n)
	for i := 0; i < n; i++ {
		params[i] = float64(cParams[i])
	}
	err = ogrError(ogrErr)
	return
}

func (sr OGRSpatialReference) GetTOWGS84() (params []float64, err error) {
	params, err = osrGetTOWGS84(sr)
	return
}

func osrAddGuessedTOWGS84(sr OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRAddGuessedTOWGS84(sr.cValue)))
	return
}

func (sr OGRSpatialReference) AddGuessedTOWGS84() (err error) {
	err = osrAddGuessedTOWGS84(sr)
	return
}

func osrSetCompoundCS(sr OGRSpatialReference, name string, horizSRS, vertSRS OGRSpatialReference) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetCompoundCS(sr.cValue, cs, horizSRS.cValue, vertSRS.cValue)))
	return
}

func (sr OGRSpatialReference) SetCompoundCS(name string, horizSRS, vertSRS OGRSpatialReference) (err error) {
	err = osrSetCompoundCS(sr, name, horizSRS, vertSRS)
	return
}

func osrPromoteTo3D(sr OGRSpatialReference, name string) (err error) {
	var cs *C.char
	if name != "" {
		cs = C.CString(name)
		defer C.free(unsafe.Pointer(cs))
	}
	err = ogrError(OGRErr(C.OSRPromoteTo3D(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) PromoteTo3D(name string) (err error) {
	err = osrPromoteTo3D(sr, name)
	return
}

func osrDemoteTo2D(sr OGRSpatialReference, name string) (err error) {
	var cs *C.char
	if name != "" {
		cs = C.CString(name)
		defer C.free(unsafe.Pointer(cs))
	}
	err = ogrError(OGRErr(C.OSRDemoteTo2D(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) DemoteTo2D(name string) (err error) {
	err = osrDemoteTo2D(sr, name)
	return
}

func osrSetGeogCS(sr OGRSpatialReference, geogName, datumName, ellipsoidName string, semiMajor, invFlattening float64, pmName string, pmOffset float64, units string, convertToRadians float64) (err error) {
	csGeog := C.CString(geogName)
	defer C.free(unsafe.Pointer(csGeog))
	csDatum := C.CString(datumName)
	defer C.free(unsafe.Pointer(csDatum))
	csEllipsoid := C.CString(ellipsoidName)
	defer C.free(unsafe.Pointer(csEllipsoid))
	var csPM *C.char
	if pmName != "" {
		csPM = C.CString(pmName)
		defer C.free(unsafe.Pointer(csPM))
	}
	var csUnits *C.char
	if units != "" {
		csUnits = C.CString(units)
		defer C.free(unsafe.Pointer(csUnits))
	}
	err = ogrError(OGRErr(C.OSRSetGeogCS(sr.cValue, csGeog, csDatum, csEllipsoid, C.double(semiMajor), C.double(invFlattening), csPM, C.double(pmOffset), csUnits, C.double(convertToRadians))))
	return
}

func (sr OGRSpatialReference) SetGeogCS(geogName, datumName, ellipsoidName string, semiMajor, invFlattening float64, pmName string, pmOffset float64, units string, convertToRadians float64) (err error) {
	err = osrSetGeogCS(sr, geogName, datumName, ellipsoidName, semiMajor, invFlattening, pmName, pmOffset, units, convertToRadians)
	return
}

func osrSetVertCS(sr OGRSpatialReference, vertCSName, vertDatumName string, vertDatumType int) (err error) {
	csName := C.CString(vertCSName)
	defer C.free(unsafe.Pointer(csName))
	csDatum := C.CString(vertDatumName)
	defer C.free(unsafe.Pointer(csDatum))
	err = ogrError(OGRErr(C.OSRSetVertCS(sr.cValue, csName, csDatum, C.int(vertDatumType))))
	return
}

func (sr OGRSpatialReference) SetVertCS(vertCSName, vertDatumName string, vertDatumType int) (err error) {
	err = osrSetVertCS(sr, vertCSName, vertDatumName, vertDatumType)
	return
}

func osrGetSemiMajor(sr OGRSpatialReference) (result float64, err error) {
	var ogrErr C.OGRErr
	result = float64(C.OSRGetSemiMajor(sr.cValue, &ogrErr))
	err = ogrError(OGRErr(ogrErr))
	return
}

func (sr OGRSpatialReference) GetSemiMajor() (result float64, err error) {
	result, err = osrGetSemiMajor(sr)
	return
}

func osrGetSemiMinor(sr OGRSpatialReference) (result float64, err error) {
	var ogrErr C.OGRErr
	result = float64(C.OSRGetSemiMinor(sr.cValue, &ogrErr))
	err = ogrError(OGRErr(ogrErr))
	return
}

func (sr OGRSpatialReference) GetSemiMinor() (result float64, err error) {
	result, err = osrGetSemiMinor(sr)
	return
}

func osrGetInvFlattening(sr OGRSpatialReference) (result float64, err error) {
	var ogrErr C.OGRErr
	result = float64(C.OSRGetInvFlattening(sr.cValue, &ogrErr))
	err = ogrError(OGRErr(ogrErr))
	return
}

func (sr OGRSpatialReference) GetInvFlattening() (result float64, err error) {
	result, err = osrGetInvFlattening(sr)
	return
}

func osrSetAuthority(sr OGRSpatialReference, targetKey, authority string, code int) (err error) {
	csKey := C.CString(targetKey)
	defer C.free(unsafe.Pointer(csKey))
	csAuth := C.CString(authority)
	defer C.free(unsafe.Pointer(csAuth))
	err = ogrError(OGRErr(C.OSRSetAuthority(sr.cValue, csKey, csAuth, C.int(code))))
	return
}

func (sr OGRSpatialReference) SetAuthority(targetKey, authority string, code int) (err error) {
	err = osrSetAuthority(sr, targetKey, authority, code)
	return
}

func osrGetAuthorityCode(sr OGRSpatialReference, targetKey string) (result string) {
	cs := C.CString(targetKey)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OSRGetAuthorityCode(sr.cValue, cs))
	return
}

func (sr OGRSpatialReference) GetAuthorityCode(targetKey string) (result string) {
	result = osrGetAuthorityCode(sr, targetKey)
	return
}

func osrGetAuthorityName(sr OGRSpatialReference, targetKey string) (result string) {
	cs := C.CString(targetKey)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OSRGetAuthorityName(sr.cValue, cs))
	return
}

func (sr OGRSpatialReference) GetAuthorityName(targetKey string) (result string) {
	result = osrGetAuthorityName(sr, targetKey)
	return
}

func osrGetAreaOfUse(sr OGRSpatialReference) (westLon, southLat, eastLon, northLat float64, areaName string, result bool) {
	var cWest, cSouth, cEast, cNorth C.double
	var cName *C.char
	result = C.OSRGetAreaOfUse(sr.cValue, &cWest, &cSouth, &cEast, &cNorth, &cName) != 0
	westLon = float64(cWest)
	southLat = float64(cSouth)
	eastLon = float64(cEast)
	northLat = float64(cNorth)
	areaName = C.GoString(cName)
	return
}

func (sr OGRSpatialReference) GetAreaOfUse() (westLon, southLat, eastLon, northLat float64, areaName string, result bool) {
	westLon, southLat, eastLon, northLat, areaName, result = osrGetAreaOfUse(sr)
	return
}

func osrSetProjection(sr OGRSpatialReference, projection string) (err error) {
	cs := C.CString(projection)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetProjection(sr.cValue, cs)))
	return
}

func (sr OGRSpatialReference) SetProjection(projection string) (err error) {
	err = osrSetProjection(sr, projection)
	return
}

func osrSetProjParm(sr OGRSpatialReference, name string, value float64) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetProjParm(sr.cValue, cs, C.double(value))))
	return
}

func (sr OGRSpatialReference) SetProjParm(name string, value float64) (err error) {
	err = osrSetProjParm(sr, name, value)
	return
}

func osrGetProjParm(sr OGRSpatialReference, name string, dfDefault float64) (result float64, err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	var ogrErr C.OGRErr
	result = float64(C.OSRGetProjParm(sr.cValue, cs, C.double(dfDefault), &ogrErr))
	err = ogrError(OGRErr(ogrErr))
	return
}

func (sr OGRSpatialReference) GetProjParm(name string, dfDefault float64) (result float64, err error) {
	result, err = osrGetProjParm(sr, name, dfDefault)
	return
}

func osrSetNormProjParm(sr OGRSpatialReference, name string, value float64) (err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetNormProjParm(sr.cValue, cs, C.double(value))))
	return
}

func (sr OGRSpatialReference) SetNormProjParm(name string, value float64) (err error) {
	err = osrSetNormProjParm(sr, name, value)
	return
}

func osrGetNormProjParm(sr OGRSpatialReference, name string, dfDefault float64) (result float64, err error) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	var ogrErr C.OGRErr
	result = float64(C.OSRGetNormProjParm(sr.cValue, cs, C.double(dfDefault), &ogrErr))
	err = ogrError(OGRErr(ogrErr))
	return
}

func (sr OGRSpatialReference) GetNormProjParm(name string, dfDefault float64) (result float64, err error) {
	result, err = osrGetNormProjParm(sr, name, dfDefault)
	return
}

func osrSetUTM(sr OGRSpatialReference, zone, north int) (err error) {
	err = ogrError(OGRErr(C.OSRSetUTM(sr.cValue, C.int(zone), C.int(north))))
	return
}

func (sr OGRSpatialReference) SetUTM(zone int, north bool) (err error) {
	bNorth := 0
	if north {
		bNorth = 1
	}
	err = osrSetUTM(sr, zone, bNorth)
	return
}

func osrGetUTMZone(sr OGRSpatialReference) (zone int, north bool) {
	var bNorth C.int
	zone = int(C.OSRGetUTMZone(sr.cValue, &bNorth))
	north = bNorth != 0
	return
}

func (sr OGRSpatialReference) GetUTMZone() (zone int, north bool) {
	zone, north = osrGetUTMZone(sr)
	return
}

func osrSetStatePlane(sr OGRSpatialReference, zone, nad83 int) (err error) {
	err = ogrError(OGRErr(C.OSRSetStatePlane(sr.cValue, C.int(zone), C.int(nad83))))
	return
}

func (sr OGRSpatialReference) SetStatePlane(zone int, nad83 bool) (err error) {
	bNad83 := 0
	if nad83 {
		bNad83 = 1
	}
	err = osrSetStatePlane(sr, zone, bNad83)
	return
}

func osrSetStatePlaneWithUnits(sr OGRSpatialReference, zone, nad83 int, unitName string, unit float64) (err error) {
	cs := C.CString(unitName)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetStatePlaneWithUnits(sr.cValue, C.int(zone), C.int(nad83), cs, C.double(unit))))
	return
}

func (sr OGRSpatialReference) SetStatePlaneWithUnits(zone int, nad83 bool, unitName string, unit float64) (err error) {
	bNad83 := 0
	if nad83 {
		bNad83 = 1
	}
	err = osrSetStatePlaneWithUnits(sr, zone, bNad83, unitName, unit)
	return
}

func osrAutoIdentifyEPSG(sr OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRAutoIdentifyEPSG(sr.cValue)))
	return
}

func (sr OGRSpatialReference) AutoIdentifyEPSG() (err error) {
	err = osrAutoIdentifyEPSG(sr)
	return
}

func osrFindMatches(sr OGRSpatialReference, options []string) (result OGRSpatialReferences) {
	cOptions := make([]*C.char, len(options)+1)
	for i, o := range options {
		cOptions[i] = C.CString(o)
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[len(options)] = nil
	result = OGRSpatialReferences{cValue: C.OSRFindMatches(sr.cValue, &cOptions[0], nil, nil)}
	return
}

func (sr OGRSpatialReference) FindMatches(options []string) (result OGRSpatialReferences) {
	result = osrFindMatches(sr, options)
	return
}

func osrFreeSRSArray(pahSRS OGRSpatialReferences) {
	C.OSRFreeSRSArray(pahSRS.cValue)
}

func (srs OGRSpatialReferences) FreeSRSArray() {
	osrFreeSRSArray(srs)
}

func osrEPSGTreatsAsLatLong(sr OGRSpatialReference) (result bool) {
	result = C.OSREPSGTreatsAsLatLong(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) EPSGTreatsAsLatLong() (result bool) {
	result = osrEPSGTreatsAsLatLong(sr)
	return
}

func osrEPSGTreatsAsNorthingEasting(sr OGRSpatialReference) (result bool) {
	result = C.OSREPSGTreatsAsNorthingEasting(sr.cValue) != 0
	return
}

func (sr OGRSpatialReference) EPSGTreatsAsNorthingEasting() (result bool) {
	result = osrEPSGTreatsAsNorthingEasting(sr)
	return
}

func osrGetAxis(sr OGRSpatialReference, targetKey string, iAxis int) (name string, orientation OGRAxisOrientation) {
	cs := C.CString(targetKey)
	defer C.free(unsafe.Pointer(cs))
	var cOrientation C.OGRAxisOrientation
	name = C.GoString(C.OSRGetAxis(sr.cValue, cs, C.int(iAxis), &cOrientation))
	orientation = OGRAxisOrientation(cOrientation)
	return
}

func (sr OGRSpatialReference) GetAxis(targetKey string, iAxis int) (name string, orientation OGRAxisOrientation) {
	name, orientation = osrGetAxis(sr, targetKey, iAxis)
	return
}

func osrGetAxesCount(sr OGRSpatialReference) (result int) {
	result = int(C.OSRGetAxesCount(sr.cValue))
	return
}

func (sr OGRSpatialReference) GetAxesCount() (result int) {
	result = osrGetAxesCount(sr)
	return
}

func osrSetAxes(sr OGRSpatialReference, targetKey, xAxisName string, xOrientation OGRAxisOrientation, yAxisName string, yOrientation OGRAxisOrientation) (err error) {
	csKey := C.CString(targetKey)
	defer C.free(unsafe.Pointer(csKey))
	csX := C.CString(xAxisName)
	defer C.free(unsafe.Pointer(csX))
	csY := C.CString(yAxisName)
	defer C.free(unsafe.Pointer(csY))
	err = ogrError(OGRErr(C.OSRSetAxes(sr.cValue, csKey, csX, C.OGRAxisOrientation(xOrientation), csY, C.OGRAxisOrientation(yOrientation))))
	return
}

func (sr OGRSpatialReference) SetAxes(targetKey, xAxisName string, xOrientation OGRAxisOrientation, yAxisName string, yOrientation OGRAxisOrientation) (err error) {
	err = osrSetAxes(sr, targetKey, xAxisName, xOrientation, yAxisName, yOrientation)
	return
}

type OSRAxisMappingStrategy C.OSRAxisMappingStrategy

const (
	OAMSTraditionalGisOrder OSRAxisMappingStrategy = C.OAMS_TRADITIONAL_GIS_ORDER
	OAMSAuthorityCompliant  OSRAxisMappingStrategy = C.OAMS_AUTHORITY_COMPLIANT
	OAMSCustom              OSRAxisMappingStrategy = C.OAMS_CUSTOM
)

func osrGetAxisMappingStrategy(sr OGRSpatialReference) (result OSRAxisMappingStrategy) {
	result = OSRAxisMappingStrategy(C.OSRGetAxisMappingStrategy(sr.cValue))
	return
}

func (sr OGRSpatialReference) GetAxisMappingStrategy() (result OSRAxisMappingStrategy) {
	result = osrGetAxisMappingStrategy(sr)
	return
}

func osrSetAxisMappingStrategy(sr OGRSpatialReference, strategy OSRAxisMappingStrategy) {
	C.OSRSetAxisMappingStrategy(sr.cValue, C.OSRAxisMappingStrategy(strategy))
}

func (sr OGRSpatialReference) SetAxisMappingStrategy(strategy OSRAxisMappingStrategy) {
	osrSetAxisMappingStrategy(sr, strategy)
}

func osrGetDataAxisToSRSAxisMapping(sr OGRSpatialReference) (result []int) {
	var count C.int
	raw := C.OSRGetDataAxisToSRSAxisMapping(sr.cValue, &count)
	if raw == nil {
		return
	}
	n := int(count)
	result = make([]int, n)
	for i := 0; i < n; i++ {
		p := (*C.int)(unsafe.Add(unsafe.Pointer(raw), uintptr(i)*unsafe.Sizeof(*raw)))
		result[i] = int(*p)
	}
	return
}

func (sr OGRSpatialReference) GetDataAxisToSRSAxisMapping() (result []int) {
	result = osrGetDataAxisToSRSAxisMapping(sr)
	return
}

func osrSetDataAxisToSRSAxisMapping(sr OGRSpatialReference, mapping []int) (err error) {
	if len(mapping) == 0 {
		return
	}
	cMapping := make([]C.int, len(mapping))
	for i, m := range mapping {
		cMapping[i] = C.int(m)
	}
	err = ogrError(OGRErr(C.OSRSetDataAxisToSRSAxisMapping(sr.cValue, C.int(len(mapping)), &cMapping[0])))
	return
}

func (sr OGRSpatialReference) SetDataAxisToSRSAxisMapping(mapping []int) (err error) {
	err = osrSetDataAxisToSRSAxisMapping(sr, mapping)
	return
}

/** Albers Conic Equal Area */
func osrSetACEA(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetACEA(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetACEA(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetACEA(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Azimuthal Equidistant */
func osrSetAE(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetAE(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetAE(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetAE(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Bonne */
func osrSetBonne(sr OGRSpatialReference, dfStandardParallel, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetBonne(sr.cValue, C.double(dfStandardParallel), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetBonne(dfStandardParallel, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetBonne(sr, dfStandardParallel, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Cylindrical Equal Area */
func osrSetCEA(sr OGRSpatialReference, dfStdP1, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetCEA(sr.cValue, C.double(dfStdP1), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetCEA(dfStdP1, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetCEA(sr, dfStdP1, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Cassini-Soldner */
func osrSetCS(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetCS(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetCS(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetCS(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Equidistant Conic */
func osrSetEC(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetEC(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetEC(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetEC(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Eckert I-VI */
func osrSetEckert(sr OGRSpatialReference, nVariation int, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetEckert(sr.cValue, C.int(nVariation), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetEckert(nVariation int, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetEckert(sr, nVariation, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Eckert IV */
func osrSetEckertIV(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetEckertIV(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetEckertIV(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetEckertIV(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Eckert VI */
func osrSetEckertVI(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetEckertVI(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetEckertVI(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetEckertVI(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Equirectangular */
func osrSetEquirectangular(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetEquirectangular(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetEquirectangular(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetEquirectangular(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Equirectangular generalized form */
func osrSetEquirectangular2(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfPseudoStdParallel1, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetEquirectangular2(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfPseudoStdParallel1), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetEquirectangular2(dfCenterLat, dfCenterLong, dfPseudoStdParallel1, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetEquirectangular2(sr, dfCenterLat, dfCenterLong, dfPseudoStdParallel1, dfFalseEasting, dfFalseNorthing)
	return
}

/** Gall Stereograpic */
func osrSetGS(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetGS(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetGS(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetGS(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Goode Homolosine */
func osrSetGH(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetGH(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetGH(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetGH(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Interrupted Goode Homolosine */
func osrSetIGH(sr OGRSpatialReference) (err error) {
	err = ogrError(OGRErr(C.OSRSetIGH(sr.cValue)))
	return
}

func (sr OGRSpatialReference) SetIGH() (err error) {
	err = osrSetIGH(sr)
	return
}

/** GEOS - Geostationary Satellite View */
func osrSetGEOS(sr OGRSpatialReference, dfCentralMeridian, dfSatelliteHeight, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetGEOS(sr.cValue, C.double(dfCentralMeridian), C.double(dfSatelliteHeight), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetGEOS(dfCentralMeridian, dfSatelliteHeight, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetGEOS(sr, dfCentralMeridian, dfSatelliteHeight, dfFalseEasting, dfFalseNorthing)
	return
}

/** Gauss Schreiber Transverse Mercator */
func osrSetGaussSchreiberTMercator(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetGaussSchreiberTMercator(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetGaussSchreiberTMercator(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetGaussSchreiberTMercator(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Gnomonic */
func osrSetGnomonic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetGnomonic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetGnomonic(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetGnomonic(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

// #ifdef undef
// /** Oblique Mercator (aka HOM (variant B) */
// OGRErr CPL_DLL OSRSetOM(OGRSpatialReferenceH hSRS, double dfCenterLat,
//                         double dfCenterLong, double dfAzimuth,
//                         double dfRectToSkew, double dfScale,
//                         double dfFalseEasting, double dfFalseNorthing);
// #endif

/** Hotine Oblique Mercator using azimuth angle */
func osrSetHOM(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetHOM(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfAzimuth), C.double(dfRectToSkew), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetHOM(dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetHOM(sr, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

func osrSetHOMAC(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetHOMAC(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfAzimuth), C.double(dfRectToSkew), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetHOMAC(dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetHOMAC(sr, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Hotine Oblique Mercator using two points on centerline */
func osrSetHOM2PNO(sr OGRSpatialReference, dfCenterLat, dfLat1, dfLong1, dfLat2, dfLong2, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetHOM2PNO(sr.cValue, C.double(dfCenterLat), C.double(dfLat1), C.double(dfLong1), C.double(dfLat2), C.double(dfLong2), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetHOM2PNO(dfCenterLat, dfLat1, dfLong1, dfLat2, dfLong2, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetHOM2PNO(sr, dfCenterLat, dfLat1, dfLong1, dfLat2, dfLong2, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** International Map of the World Polyconic */
func osrSetIWMPolyconic(sr OGRSpatialReference, dfLat1, dfLat2, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetIWMPolyconic(sr.cValue, C.double(dfLat1), C.double(dfLat2), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetIWMPolyconic(dfLat1, dfLat2, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetIWMPolyconic(sr, dfLat1, dfLat2, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Krovak Oblique Conic Conformal */
func osrSetKrovak(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfAzimuth, dfPseudoStdParallelLat, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetKrovak(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfAzimuth), C.double(dfPseudoStdParallelLat), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetKrovak(dfCenterLat, dfCenterLong, dfAzimuth, dfPseudoStdParallelLat, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetKrovak(sr, dfCenterLat, dfCenterLong, dfAzimuth, dfPseudoStdParallelLat, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Lambert Azimuthal Equal-Area */
func osrSetLAEA(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetLAEA(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetLAEA(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetLAEA(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Lambert Conformal Conic */
func osrSetLCC(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetLCC(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetLCC(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetLCC(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Lambert Conformal Conic 1SP */
func osrSetLCC1SP(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetLCC1SP(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetLCC1SP(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetLCC1SP(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Lambert Conformal Conic (Belgium) */
func osrSetLCCB(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetLCCB(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetLCCB(dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetLCCB(sr, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Miller Cylindrical */
func osrSetMC(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetMC(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetMC(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetMC(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Mercator */
func osrSetMercator(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetMercator(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetMercator(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetMercator(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Mercator 2SP */
func osrSetMercator2SP(sr OGRSpatialReference, dfStdP1, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetMercator2SP(sr.cValue, C.double(dfStdP1), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetMercator2SP(dfStdP1, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetMercator2SP(sr, dfStdP1, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Mollweide */
func osrSetMollweide(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetMollweide(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetMollweide(dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetMollweide(sr, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** New Zealand Map Grid */
func osrSetNZMG(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetNZMG(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetNZMG(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetNZMG(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Oblique Stereographic */
func osrSetOS(sr OGRSpatialReference, dfOriginLat, dfCMeridian, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetOS(sr.cValue, C.double(dfOriginLat), C.double(dfCMeridian), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetOS(dfOriginLat, dfCMeridian, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetOS(sr, dfOriginLat, dfCMeridian, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Orthographic */
func osrSetOrthographic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetOrthographic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetOrthographic(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetOrthographic(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Polyconic */
func osrSetPolyconic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetPolyconic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetPolyconic(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetPolyconic(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Polar Stereographic */
func osrSetPS(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetPS(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetPS(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetPS(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Robinson */
func osrSetRobinson(sr OGRSpatialReference, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetRobinson(sr.cValue, C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetRobinson(dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetRobinson(sr, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Sinusoidal */
func osrSetSinusoidal(sr OGRSpatialReference, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetSinusoidal(sr.cValue, C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetSinusoidal(dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetSinusoidal(sr, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Stereographic */
func osrSetStereographic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetStereographic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetStereographic(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetStereographic(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Swiss Oblique Cylindrical */
func osrSetSOC(sr OGRSpatialReference, dfLatitudeOfOrigin, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetSOC(sr.cValue, C.double(dfLatitudeOfOrigin), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetSOC(dfLatitudeOfOrigin, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetSOC(sr, dfLatitudeOfOrigin, dfCentralMeridian, dfFalseEasting, dfFalseNorthing)
	return
}

/** Transverse Mercator
 *
 * Special processing available for Transverse Mercator with GDAL &gt;= 1.10 and
 * PROJ &gt;= 4.8 : see OGRSpatialReference::exportToProj4().
 */
func osrSetTM(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetTM(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetTM(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetTM(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Transverse Mercator variant */
func osrSetTMVariant(sr OGRSpatialReference, pszVariantName string, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	cs := C.CString(pszVariantName)
	defer C.free(unsafe.Pointer(cs))
	err = ogrError(OGRErr(C.OSRSetTMVariant(sr.cValue, cs, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetTMVariant(pszVariantName string, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetTMVariant(sr, pszVariantName, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** Tunesia Mining Grid  */
func osrSetTMG(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetTMG(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetTMG(dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetTMG(sr, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Transverse Mercator (South Oriented) */
func osrSetTMSO(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetTMSO(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetTMSO(dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetTMSO(sr, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing)
	return
}

/** TPED (Two Point Equi Distant) */
func osrSetTPED(sr OGRSpatialReference, dfLat1, dfLong1, dfLat2, dfLong2, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetTPED(sr.cValue, C.double(dfLat1), C.double(dfLong1), C.double(dfLat2), C.double(dfLong2), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetTPED(dfLat1, dfLong1, dfLat2, dfLong2, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetTPED(sr, dfLat1, dfLong1, dfLat2, dfLong2, dfFalseEasting, dfFalseNorthing)
	return
}

/** VanDerGrinten */
func osrSetVDG(sr OGRSpatialReference, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetVDG(sr.cValue, C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetVDG(dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetVDG(sr, dfCenterLong, dfFalseEasting, dfFalseNorthing)
	return
}

/** Wagner I \-- VII */
func osrSetWagner(sr OGRSpatialReference, nVariation int, dfCenterLat, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetWagner(sr.cValue, C.int(nVariation), C.double(dfCenterLat), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetWagner(nVariation int, dfCenterLat, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetWagner(sr, nVariation, dfCenterLat, dfFalseEasting, dfFalseNorthing)
	return
}

/** Quadrilateralized Spherical Cube */
func osrSetQSC(sr OGRSpatialReference, dfCenterLat, dfCenterLong float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetQSC(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong))))
	return
}

func (sr OGRSpatialReference) SetQSC(dfCenterLat, dfCenterLong float64) (err error) {
	err = osrSetQSC(sr, dfCenterLat, dfCenterLong)
	return
}

/** Spherical, Cross-track, Height */
func osrSetSCH(sr OGRSpatialReference, dfPegLat, dfPegLong, dfPegHeading, dfPegHgt float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetSCH(sr.cValue, C.double(dfPegLat), C.double(dfPegLong), C.double(dfPegHeading), C.double(dfPegHgt))))
	return
}

func (sr OGRSpatialReference) SetSCH(dfPegLat, dfPegLong, dfPegHeading, dfPegHgt float64) (err error) {
	err = osrSetSCH(sr, dfPegLat, dfPegLong, dfPegHeading, dfPegHgt)
	return
}

/** Vertical Perspective / Near-sided Perspective */
func osrSetVerticalPerspective(sr OGRSpatialReference, dfTopoOriginLat, dfTopoOriginLon, dfTopoOriginHeight, dfViewPointHeight, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = ogrError(OGRErr(C.OSRSetVerticalPerspective(sr.cValue, C.double(dfTopoOriginLat), C.double(dfTopoOriginLon), C.double(dfTopoOriginHeight), C.double(dfViewPointHeight), C.double(dfFalseEasting), C.double(dfFalseNorthing))))
	return
}

func (sr OGRSpatialReference) SetVerticalPerspective(dfTopoOriginLat, dfTopoOriginLon, dfTopoOriginHeight, dfViewPointHeight, dfFalseEasting, dfFalseNorthing float64) (err error) {
	err = osrSetVerticalPerspective(sr, dfTopoOriginLat, dfTopoOriginLon, dfTopoOriginHeight, dfViewPointHeight, dfFalseEasting, dfFalseNorthing)
	return
}

func osrCalcInvFlattening(semiMajor, semiMinor float64) (result float64) {
	result = float64(C.OSRCalcInvFlattening(C.double(semiMajor), C.double(semiMinor)))
	return
}

func OSRCalcInvFlattening(semiMajor, semiMinor float64) (result float64) {
	result = osrCalcInvFlattening(semiMajor, semiMinor)
	return
}

func osrCalcSemiMinorFromInvFlattening(semiMajor, invFlattening float64) (result float64) {
	result = float64(C.OSRCalcSemiMinorFromInvFlattening(C.double(semiMajor), C.double(invFlattening)))
	return
}

func OSRCalcSemiMinorFromInvFlattening(semiMajor, invFlattening float64) (result float64) {
	result = osrCalcSemiMinorFromInvFlattening(semiMajor, invFlattening)
	return
}

func OSRCleanup() {
	C.OSRCleanup()
}

type OSRCRSType C.OSRCRSType

const (
	OSRCrsTypeGeographic2D OSRCRSType = C.OSR_CRS_TYPE_GEOGRAPHIC_2D
	OSRCrsTypeGeographic3D OSRCRSType = C.OSR_CRS_TYPE_GEOGRAPHIC_3D
	OSRCrsTypeGeocentric   OSRCRSType = C.OSR_CRS_TYPE_GEOCENTRIC
	OSRCrsTypeProjected    OSRCRSType = C.OSR_CRS_TYPE_PROJECTED
	OSRCrsTypeVertical     OSRCRSType = C.OSR_CRS_TYPE_VERTICAL
	OSRCrsTypeCompound     OSRCRSType = C.OSR_CRS_TYPE_COMPOUND
	OSRCrsTypeOther        OSRCRSType = C.OSR_CRS_TYPE_OTHER
)

type OSRCRSInfo struct {
	cValue *C.OSRCRSInfo
}

type OSRCRSInfos struct {
	cValue **OSRCRSInfo
}

type OSRCRSListParameters struct {
	cValue *C.OSRCRSListParameters
}

func osrGetCRSInfoListFromDatabase(authName string, params OSRCRSListParameters) (result []OSRCRSInfo, err error) {
	cs := C.CString(authName)
	defer C.free(unsafe.Pointer(cs))
	var count C.int
	raw := C.OSRGetCRSInfoListFromDatabase(cs, params.cValue, &count)
	if raw == nil {
		err = lastError()
		return
	}
	n := int(count)
	result = make([]OSRCRSInfo, n)
	for i := 0; i < n; i++ {
		result[i] = OSRCRSInfo{cValue: *(**C.OSRCRSInfo)(unsafe.Add(unsafe.Pointer(raw), uintptr(i)*unsafe.Sizeof(*raw)))}
	}
	return
}

func OSRGetCRSInfoListFromDatabase(authName string, params OSRCRSListParameters) (result []OSRCRSInfo, err error) {
	result, err = osrGetCRSInfoListFromDatabase(authName, params)
	return
}

func osrDestroyCRSInfoList(list **C.OSRCRSInfo) {
	C.OSRDestroyCRSInfoList(list)
}

func OSRDestroyCRSInfoList(list **C.OSRCRSInfo) {
	osrDestroyCRSInfoList(list)
}

func osrGetAuthorityListFromDatabase() (result []string) {
	raw := C.OSRGetAuthorityListFromDatabase()
	if raw == nil {
		return
	}
	defer C.CSLDestroy(raw)
	for i := 0; ; i++ {
		p := C.CSLGetField(raw, C.int(i))
		if p == nil {
			break
		}
		result = append(result, C.GoString(p))
	}
	return
}

func OSRGetAuthorityListFromDatabase() (result []string) {
	result = osrGetAuthorityListFromDatabase()
	return
}

/* -------------------------------------------------------------------- */
/*      OGRCoordinateTransform C API.                                   */
/* -------------------------------------------------------------------- */

func octNewCoordinateTransformation(source, target OGRSpatialReference) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTNewCoordinateTransformation(source.cValue, target.cValue)}
	return
}

func OCTNewCoordinateTransformation(source, target OGRSpatialReference) (result OGRCoordinateTransformation, err error) {
	result = octNewCoordinateTransformation(source, target)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func octNewCoordinateTransformationOptions() (result OGRCoordinateTransformationOptions) {
	result = OGRCoordinateTransformationOptions{cValue: C.OCTNewCoordinateTransformationOptions()}
	return
}

func OCTNewCoordinateTransformationOptions() (result OGRCoordinateTransformationOptions, err error) {
	result = octNewCoordinateTransformationOptions()
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func octCoordinateTransformationOptionsSetOperation(opts OGRCoordinateTransformationOptions, pszCO string, bReverseCO int) (result bool) {
	cs := C.CString(pszCO)
	defer C.free(unsafe.Pointer(cs))
	result = C.OCTCoordinateTransformationOptionsSetOperation(opts.cValue, cs, C.int(bReverseCO)) != 0
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

func octCoordinateTransformationOptionsSetAreaOfInterest(opts OGRCoordinateTransformationOptions, dfWestLongitudeDeg, dfSouthLatitudeDeg, dfEastLongitudeDeg, dfNorthLatitudeDeg float64) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetAreaOfInterest(opts.cValue, C.double(dfWestLongitudeDeg), C.double(dfSouthLatitudeDeg), C.double(dfEastLongitudeDeg), C.double(dfNorthLatitudeDeg)) != 0
	return
}

func (opts OGRCoordinateTransformationOptions) SetAreaOfInterest(dfWestLongitudeDeg, dfSouthLatitudeDeg, dfEastLongitudeDeg, dfNorthLatitudeDeg float64) (result bool) {
	result = octCoordinateTransformationOptionsSetAreaOfInterest(opts, dfWestLongitudeDeg, dfSouthLatitudeDeg, dfEastLongitudeDeg, dfNorthLatitudeDeg)
	return
}

func octCoordinateTransformationOptionsSetDesiredAccuracy(opts OGRCoordinateTransformationOptions, dfAccuracy float64) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetDesiredAccuracy(opts.cValue, C.double(dfAccuracy)) != 0
	return
}

func (opts OGRCoordinateTransformationOptions) SetDesiredAccuracy(dfAccuracy float64) (result bool) {
	result = octCoordinateTransformationOptionsSetDesiredAccuracy(opts, dfAccuracy)
	return
}

func octCoordinateTransformationOptionsSetBallparkAllowed(opts OGRCoordinateTransformationOptions, bAllowBallpark int) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetBallparkAllowed(opts.cValue, C.int(bAllowBallpark)) != 0
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

func octCoordinateTransformationOptionsSetOnlyBest(opts OGRCoordinateTransformationOptions, bOnlyBest bool) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetOnlyBest(opts.cValue, C._Bool(bOnlyBest)) != 0
	return
}

func (opts OGRCoordinateTransformationOptions) SetOnlyBest(bOnlyBest bool) (result bool) {
	result = octCoordinateTransformationOptionsSetOnlyBest(opts, bOnlyBest)
	return
}

func (opts OGRCoordinateTransformationOptions) Destroy() {
	C.OCTDestroyCoordinateTransformationOptions(opts.cValue)
}

func octNewCoordinateTransformationEx(source, target OGRSpatialReference, options OGRCoordinateTransformationOptions) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTNewCoordinateTransformationEx(source.cValue, target.cValue, options.cValue)}
	return
}

func OCTNewCoordinateTransformationEx(source, target OGRSpatialReference, options OGRCoordinateTransformationOptions) (result OGRCoordinateTransformation, err error) {
	result = octNewCoordinateTransformationEx(source, target, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func octClone(ct OGRCoordinateTransformation) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTClone(ct.cValue)}
	return
}

func (ct OGRCoordinateTransformation) Clone() (result OGRCoordinateTransformation, err error) {
	result = octClone(ct)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func octGetSourceCS(ct OGRCoordinateTransformation) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OCTGetSourceCS(ct.cValue)}
	return
}

func (ct OGRCoordinateTransformation) GetSourceCS() (result OGRSpatialReference) {
	result = octGetSourceCS(ct)
	return
}

func octGetTargetCS(ct OGRCoordinateTransformation) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OCTGetTargetCS(ct.cValue)}
	return
}

func (ct OGRCoordinateTransformation) GetTargetCS() (result OGRSpatialReference) {
	result = octGetTargetCS(ct)
	return
}

func octGetInverse(ct OGRCoordinateTransformation) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTGetInverse(ct.cValue)}
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
	C.OCTDestroyCoordinateTransformation(ct.cValue)
}

func octTransform(ct OGRCoordinateTransformation, x, y, z []float64) (result bool) {
	nCount := len(x)
	if nCount == 0 {
		return
	}
	cx := make([]C.double, nCount)
	cy := make([]C.double, nCount)
	for i := 0; i < nCount; i++ {
		cx[i] = C.double(x[i])
		cy[i] = C.double(y[i])
	}
	var pz *C.double
	var cz []C.double
	if len(z) > 0 {
		cz = make([]C.double, nCount)
		for i := 0; i < nCount; i++ {
			cz[i] = C.double(z[i])
		}
		pz = &cz[0]
	}
	result = C.OCTTransform(ct.cValue, C.int(nCount), &cx[0], &cy[0], pz) != 0
	for i := 0; i < nCount; i++ {
		x[i] = float64(cx[i])
		y[i] = float64(cy[i])
	}
	for i := 0; i < len(cz); i++ {
		z[i] = float64(cz[i])
	}
	return
}

func (ct OGRCoordinateTransformation) Transform(x, y, z []float64) (result bool) {
	result = octTransform(ct, x, y, z)
	return
}

func octTransformEx(ct OGRCoordinateTransformation, x, y, z []float64) (success []bool, result bool) {
	nCount := len(x)
	if nCount == 0 {
		return
	}
	cx := make([]C.double, nCount)
	cy := make([]C.double, nCount)
	for i := 0; i < nCount; i++ {
		cx[i] = C.double(x[i])
		cy[i] = C.double(y[i])
	}
	var pz *C.double
	var cz []C.double
	if len(z) > 0 {
		cz = make([]C.double, nCount)
		for i := 0; i < nCount; i++ {
			cz[i] = C.double(z[i])
		}
		pz = &cz[0]
	}
	cSuccess := make([]C.int, nCount)
	result = C.OCTTransformEx(ct.cValue, C.int(nCount), &cx[0], &cy[0], pz, &cSuccess[0]) != 0
	for i := 0; i < nCount; i++ {
		x[i] = float64(cx[i])
		y[i] = float64(cy[i])
	}
	for i := 0; i < len(cz); i++ {
		z[i] = float64(cz[i])
	}
	success = make([]bool, nCount)
	for i := 0; i < nCount; i++ {
		success[i] = cSuccess[i] != 0
	}
	return
}

func (ct OGRCoordinateTransformation) TransformEx(x, y, z []float64) (success []bool, result bool) {
	success, result = octTransformEx(ct, x, y, z)
	return
}

func octTransform4D(ct OGRCoordinateTransformation, x, y, z, t []float64) (success []bool, result bool) {
	nCount := len(x)
	if nCount == 0 {
		return
	}
	cx := make([]C.double, nCount)
	cy := make([]C.double, nCount)
	for i := 0; i < nCount; i++ {
		cx[i] = C.double(x[i])
		cy[i] = C.double(y[i])
	}
	var pz *C.double
	var cz []C.double
	if len(z) > 0 {
		cz = make([]C.double, nCount)
		for i := 0; i < nCount; i++ {
			cz[i] = C.double(z[i])
		}
		pz = &cz[0]
	}
	var pt *C.double
	var ct4 []C.double
	if len(t) > 0 {
		ct4 = make([]C.double, nCount)
		for i := 0; i < nCount; i++ {
			ct4[i] = C.double(t[i])
		}
		pt = &ct4[0]
	}
	cSuccess := make([]C.int, nCount)
	result = C.OCTTransform4D(ct.cValue, C.int(nCount), &cx[0], &cy[0], pz, pt, &cSuccess[0]) != 0
	for i := 0; i < nCount; i++ {
		x[i] = float64(cx[i])
		y[i] = float64(cy[i])
	}
	for i := 0; i < len(cz); i++ {
		z[i] = float64(cz[i])
	}
	for i := 0; i < len(ct4); i++ {
		t[i] = float64(ct4[i])
	}
	success = make([]bool, nCount)
	for i := 0; i < nCount; i++ {
		success[i] = cSuccess[i] != 0
	}
	return
}

func (ct OGRCoordinateTransformation) Transform4D(x, y, z, t []float64) (success []bool, result bool) {
	success, result = octTransform4D(ct, x, y, z, t)
	return
}

func octTransform4DWithErrorCodes(ct OGRCoordinateTransformation, x, y, z, t []float64) (errorCodes []int, result bool) {
	nCount := len(x)
	if nCount == 0 {
		return
	}
	cx := make([]C.double, nCount)
	cy := make([]C.double, nCount)
	for i := 0; i < nCount; i++ {
		cx[i] = C.double(x[i])
		cy[i] = C.double(y[i])
	}
	var pz *C.double
	var cz []C.double
	if len(z) > 0 {
		cz = make([]C.double, nCount)
		for i := 0; i < nCount; i++ {
			cz[i] = C.double(z[i])
		}
		pz = &cz[0]
	}
	var pt *C.double
	var ct4 []C.double
	if len(t) > 0 {
		ct4 = make([]C.double, nCount)
		for i := 0; i < nCount; i++ {
			ct4[i] = C.double(t[i])
		}
		pt = &ct4[0]
	}
	cErrorCodes := make([]C.int, nCount)
	result = C.OCTTransform4DWithErrorCodes(ct.cValue, C.int(nCount), &cx[0], &cy[0], pz, pt, &cErrorCodes[0]) != 0
	for i := 0; i < nCount; i++ {
		x[i] = float64(cx[i])
		y[i] = float64(cy[i])
	}
	for i := 0; i < len(cz); i++ {
		z[i] = float64(cz[i])
	}
	for i := 0; i < len(ct4); i++ {
		t[i] = float64(ct4[i])
	}
	errorCodes = make([]int, nCount)
	for i := 0; i < nCount; i++ {
		errorCodes[i] = int(cErrorCodes[i])
	}
	return
}

func (ct OGRCoordinateTransformation) Transform4DWithErrorCodes(x, y, z, t []float64) (errorCodes []int, result bool) {
	errorCodes, result = octTransform4DWithErrorCodes(ct, x, y, z, t)
	return
}

func octTransformBounds(ct OGRCoordinateTransformation, xmin, ymin, xmax, ymax float64, densifyPts int) (outXmin, outYmin, outXmax, outYmax float64, result bool) {
	var coutXmin, coutYmin, coutXmax, coutYmax C.double
	result = C.OCTTransformBounds(ct.cValue, C.double(xmin), C.double(ymin), C.double(xmax), C.double(ymax), &coutXmin, &coutYmin, &coutXmax, &coutYmax, C.int(densifyPts)) != 0
	outXmin = float64(coutXmin)
	outYmin = float64(coutYmin)
	outXmax = float64(coutXmax)
	outYmax = float64(coutYmax)
	return
}

func (ct OGRCoordinateTransformation) TransformBounds(xmin, ymin, xmax, ymax float64, densifyPts int) (outXmin, outYmin, outXmax, outYmax float64, result bool) {
	outXmin, outYmin, outXmax, outYmax, result = octTransformBounds(ct, xmin, ymin, xmax, ymax, densifyPts)
	return
}
