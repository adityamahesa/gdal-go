package gdal

/*
#include "ogr_srs_api_preamble.h"
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
var SRSWKTWgs84LatLong = C.GoString(C._SRS_WKT_WGS84_LAT_LONG())

// /* ==================================================================== */
// /*      Some "standard" strings.                                        */
// /* ==================================================================== */

// /** Albers_Conic_Equal_Area projection */
var SRSPTAlbersConicEqualArea = C.GoString(C._SRS_PT_ALBERS_CONIC_EQUAL_AREA())

// /** Azimuthal_Equidistant projection */
var SRSPTAzimuthalEquidistant = C.GoString(C._SRS_PT_AZIMUTHAL_EQUIDISTANT())

// /** Cassini_Soldner projection */
var SRSPTCassiniSoldner = C.GoString(C._SRS_PT_CASSINI_SOLDNER())

// /** Cylindrical_Equal_Area projection */
var SRSPTCylindricalEqualArea = C.GoString(C._SRS_PT_CYLINDRICAL_EQUAL_AREA())

// /** Cylindrical_Equal_Area projection */
var SRSPTBonne = C.GoString(C._SRS_PT_BONNE())

// /** Eckert_I projection */
var SRSPTEckertI = C.GoString(C._SRS_PT_ECKERT_I())

// /** Eckert_II projection */
var SRSPTEckertII = C.GoString(C._SRS_PT_ECKERT_II())

// /** Eckert_III projection */
var SRSPTEckertIII = C.GoString(C._SRS_PT_ECKERT_III())

// /** Eckert_IV projection */
var SRSPTEckertIV = C.GoString(C._SRS_PT_ECKERT_IV())

// /** Eckert_V projection */
var SRSPTEckertV = C.GoString(C._SRS_PT_ECKERT_V())

// /** Eckert_VI projection */
var SRSPTEckertVI = C.GoString(C._SRS_PT_ECKERT_VI())

// /** Equidistant_Conic projection */
var SRSPTEquidistantConic = C.GoString(C._SRS_PT_EQUIDISTANT_CONIC())

// /** Equirectangular projection */
var SRSPTEquirectangular = C.GoString(C._SRS_PT_EQUIRECTANGULAR())

// /** Gall_Stereographic projection */
var SRSPTGallStereographic = C.GoString(C._SRS_PT_GALL_STEREOGRAPHIC())

// /** Gauss_Schreiber_Transverse_Mercator projection */
var SRSPTGaussschreibertmercator = C.GoString(C._SRS_PT_GAUSSSCHREIBERTMERCATOR())

// /** Geostationary_Satellite projection */
var SRSPTGeostationarySatellite = C.GoString(C._SRS_PT_GEOSTATIONARY_SATELLITE())

// /** Goode_Homolosine projection */
var SRSPTGoodeHomolosine = C.GoString(C._SRS_PT_GOODE_HOMOLOSINE())

// /** Interrupted_Goode_Homolosine projection */
var SRSPTIgh = C.GoString(C._SRS_PT_IGH())

// /** Gnomonic projection */
var SRSPTGnomonic = C.GoString(C._SRS_PT_GNOMONIC())

// /** Hotine_Oblique_Mercator_Azimuth_Center projection */
var SRSPTHotineObliqueMercatorAzimuthCenter = C.GoString(C._SRS_PT_HOTINE_OBLIQUE_MERCATOR_AZIMUTH_CENTER())

// /** Hotine_Oblique_Mercator projection */
var SRSPTHotineObliqueMercator = C.GoString(C._SRS_PT_HOTINE_OBLIQUE_MERCATOR())

// /** Hotine_Oblique_Mercator_Two_Point_Natural_Origin projection */
var SRSPTHotineObliqueMercatorTwoPointNaturalOrigin = C.GoString(C._SRS_PT_HOTINE_OBLIQUE_MERCATOR_TWO_POINT_NATURAL_ORIGIN())

// /** Laborde_Oblique_Mercator projection */
var SRSPTLabardeObliqueMercator = C.GoString(C._SRS_PT_LABORDE_OBLIQUE_MERCATOR())

// /** Lambert_Conformal_Conic_1SP projection */
var SRSPTLambertConformalConic1SP = C.GoString(C._SRS_PT_LAMBERT_CONFORMAL_CONIC_1SP())

// /** Lambert_Conformal_Conic_2SP projection */
var SRSPTLambertConformalConic2SP = C.GoString(C._SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP())

// /** Lambert_Conformal_Conic_2SP_Belgium projection */
var SRSPTLambertConformalConic2SPBelgium = C.GoString(C._SRS_PT_LAMBERT_CONFORMAL_CONIC_2SP_BELGIUM())

// /** Lambert_Azimuthal_Equal_Area projection */
var SRSPTLambertAzimuthalEqualArea = C.GoString(C._SRS_PT_LAMBERT_AZIMUTHAL_EQUAL_AREA())

// /** Mercator_1SP projection */
var SRSPTMercator1SP = C.GoString(C._SRS_PT_MERCATOR_1SP())

// /** Mercator_2SP projection */
var SRSPTMercator2SP = C.GoString(C._SRS_PT_MERCATOR_2SP())

// /** Mercator_Auxiliary_Sphere is used used by ESRI to mean EPSG:3875 */
var SRSPTMercatorAuxiliarySphere = C.GoString(C._SRS_PT_MERCATOR_AUXILIARY_SPHERE())

// /** Miller_Cylindrical projection */
var SRSPTMillerCylindrical = C.GoString(C._SRS_PT_MILLER_CYLINDRICAL())

// /** Mollweide projection */
var SRSPTMollweide = C.GoString(C._SRS_PT_MOLLWEIDE())

// /** New_Zealand_Map_Grid projection */
var SRSPTNewZealandMapGrid = C.GoString(C._SRS_PT_NEW_ZEALAND_MAP_GRID())

// /** Oblique_Stereographic projection */
var SRSPTObliqueStereographic = C.GoString(C._SRS_PT_OBLIQUE_STEREOGRAPHIC())

// /** Orthographic projection */
var SRSPTOrthographic = C.GoString(C._SRS_PT_ORTHOGRAPHIC())

// /** Polar_Stereographic projection */
var SRSPTPolarStereographic = C.GoString(C._SRS_PT_POLAR_STEREOGRAPHIC())

// /** Polyconic projection */
var SRSPTPolyconic = C.GoString(C._SRS_PT_POLYCONIC())

// /** Robinson projection */
var SRSPTRobinson = C.GoString(C._SRS_PT_ROBINSON())

// /** Sinusoidal projection */
var SRSPTSinusoidal = C.GoString(C._SRS_PT_SINUSOIDAL())

// /** Stereographic projection */
var SRSPTStereographic = C.GoString(C._SRS_PT_STEREOGRAPHIC())

// /** Swiss_Oblique_Cylindrical projection */
var SRSPTSwissObliqueCylindrical = C.GoString(C._SRS_PT_SWISS_OBLIQUE_CYLINDRICAL())

// /** Transverse_Mercator projection */
var SRSPTTransverseMercator = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR())

// /** Transverse_Mercator_South_Orientated projection */
var SRSPTTransverseMercatorSouthOriented = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_SOUTH_ORIENTED())

// /* special mapinfo variants on Transverse Mercator */
// /** Transverse_Mercator_MapInfo_21 projection */
var SRSPTTransverseMercatorMI21 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_21())

// /** Transverse_Mercator_MapInfo_22 projection */
var SRSPTTransverseMercatorMI22 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_22())

// /** Transverse_Mercator_MapInfo_23 projection */
var SRSPTTransverseMercatorMI23 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_23())

// /** Transverse_Mercator_MapInfo_24 projection */
var SRSPTTransverseMercatorMI24 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_24())

// /** Transverse_Mercator_MapInfo_25 projection */
var SRSPTTransverseMercatorMI25 = C.GoString(C._SRS_PT_TRANSVERSE_MERCATOR_MI_25())

// /** Tunisia_Mining_Grid projection */
var SRSPTTunisiaMiningGrid = C.GoString(C._SRS_PT_TUNISIA_MINING_GRID())

// /** Two_Point_Equidistant projection */
var SRSPTTwoPointEquidistant = C.GoString(C._SRS_PT_TWO_POINT_EQUIDISTANT())

// /** VanDerGrinten projection */
var SRSPTVandergrinten = C.GoString(C._SRS_PT_VANDERGRINTEN())

// /** Krovak projection */
var SRSPTKrovak = C.GoString(C._SRS_PT_KROVAK())

// /** International_Map_of_the_World_Polyconic projection */
var SRSPTIMWPolyconic = C.GoString(C._SRS_PT_IMW_POLYCONIC())

// /** Wagner_I projection */
var SRSPTWagnerI = C.GoString(C._SRS_PT_WAGNER_I())

// /** Wagner_II projection */
var SRSPTWagnerII = C.GoString(C._SRS_PT_WAGNER_II())

// /** Wagner_III projection */
var SRSPTWagnerIII = C.GoString(C._SRS_PT_WAGNER_III())

// /** Wagner_IV projection */
var SRSPTWagnerIV = C.GoString(C._SRS_PT_WAGNER_IV())

// /** Wagner_V projection */
var SRSPTWagnerV = C.GoString(C._SRS_PT_WAGNER_V())

// /** Wagner_VI projection */
var SRSPTWagnerVI = C.GoString(C._SRS_PT_WAGNER_VI())

// /** Wagner_VII projection */
var SRSPTWagnerVII = C.GoString(C._SRS_PT_WAGNER_VII())

// /** Quadrilateralized_Spherical_Cube projection */
var SRSPTQsc = C.GoString(C._SRS_PT_QSC())

// /** Aitoff projection */
var SRSPTAitoff = C.GoString(C._SRS_PT_AITOFF())

// /** Winkel_I projection */
var SRSPTWinkelI = C.GoString(C._SRS_PT_WINKEL_I())

// /** Winkel_II projection */
var SRSPTWinkelII = C.GoString(C._SRS_PT_WINKEL_II())

// /** Winkel_Tripel projection */
var SRSPTWinkelTripel = C.GoString(C._SRS_PT_WINKEL_TRIPEL())

// /** Craster_Parabolic projection */
var SRSPTCrasterParabolic = C.GoString(C._SRS_PT_CRASTER_PARABOLIC())

// /** Loximuthal projection */
var SRSPTLoximuthal = C.GoString(C._SRS_PT_LOXIMUTHAL())

// /** Quartic_Authalic projection */
var SRSPTQuarticAuthalic = C.GoString(C._SRS_PT_QUARTIC_AUTHALIC())

// /** Spherical_Cross_Track_Height projection */
var SRSPTSch = C.GoString(C._SRS_PT_SCH())

// /** central_meridian projection parameter */
var SRSPPCentralMeridian = C.GoString(C._SRS_PP_CENTRAL_MERIDIAN())

// /** scale_factor projection parameter */
var SRSPPScaleFactor = C.GoString(C._SRS_PP_SCALE_FACTOR())

// /** standard_parallel_1 projection parameter */
var SRSPPStandardParallel1 = C.GoString(C._SRS_PP_STANDARD_PARALLEL_1())

// /** standard_parallel_2 projection parameter */
var SRSPPStandardParallel2 = C.GoString(C._SRS_PP_STANDARD_PARALLEL_2())

// /** pseudo_standard_parallel_1 projection parameter */
var SRSPPPseudoStdParallel1 = C.GoString(C._SRS_PP_PSEUDO_STD_PARALLEL_1())

// /** longitude_of_center projection parameter */
var SRSPPLongitudeOfCenter = C.GoString(C._SRS_PP_LONGITUDE_OF_CENTER())

// /** latitude_of_center projection parameter */
var SRSPPLatitudeOfCenter = C.GoString(C._SRS_PP_LATITUDE_OF_CENTER())

// /** longitude_of_origin projection parameter */
var SRSPPLongitudeOfOrigin = C.GoString(C._SRS_PP_LONGITUDE_OF_ORIGIN())

// /** latitude_of_origin projection parameter */
var SRSPPLatitudeOfOrigin = C.GoString(C._SRS_PP_LATITUDE_OF_ORIGIN())

// /** false_easting projection parameter */
var SRSPPFalseEasting = C.GoString(C._SRS_PP_FALSE_EASTING())

// /** false_northing projection parameter */
var SRSPPFalseNorthing = C.GoString(C._SRS_PP_FALSE_NORTHING())

// /** azimuth projection parameter */
var SRSPPAzimuth = C.GoString(C._SRS_PP_AZIMUTH())

// /** longitude_of_point_1 projection parameter */
var SRSPPLongitudeOfPoint1 = C.GoString(C._SRS_PP_LONGITUDE_OF_POINT_1())

// /** latitude_of_point_1 projection parameter */
var SRSPPLatitudeOfPoint1 = C.GoString(C._SRS_PP_LATITUDE_OF_POINT_1())

// /** longitude_of_point_2 projection parameter */
var SRSPPLongitudeOfPoint2 = C.GoString(C._SRS_PP_LONGITUDE_OF_POINT_2())

// /** latitude_of_point_2 projection parameter */
var SRSPPLatitudeOfPoint2 = C.GoString(C._SRS_PP_LATITUDE_OF_POINT_2())

// /** longitude_of_point_3 projection parameter */
var SRSPPLongitudeOfPoint3 = C.GoString(C._SRS_PP_LONGITUDE_OF_POINT_3())

// /** latitude_of_point_3 projection parameter */
var SRSPPLatitudeOfPoint3 = C.GoString(C._SRS_PP_LATITUDE_OF_POINT_3())

// /** rectified_grid_angle projection parameter */
var SRSPPRectifiedGridAngle = C.GoString(C._SRS_PP_RECTIFIED_GRID_ANGLE())

// /** landsat_number projection parameter */
var SRSPPLandsatNumber = C.GoString(C._SRS_PP_LANDSAT_NUMBER())

// /** path_number projection parameter */
var SRSPPPathNumber = C.GoString(C._SRS_PP_PATH_NUMBER())

// /** perspective_point_height projection parameter */
var SRSPPPerspectivePointHeight = C.GoString(C._SRS_PP_PERSPECTIVE_POINT_HEIGHT())

// /** satellite_height projection parameter */
var SRSPPSatelliteHeight = C.GoString(C._SRS_PP_SATELLITE_HEIGHT())

// /** fipszone projection parameter */
var SRSPPFipszone = C.GoString(C._SRS_PP_FIPSZONE())

// /** zone projection parameter */
var SRSPPZone = C.GoString(C._SRS_PP_ZONE())

// /** Latitude_Of_1st_Point projection parameter */
var SRSPPLatitudeOf1stPoint = C.GoString(C._SRS_PP_LATITUDE_OF_1ST_POINT())

// /** Longitude_Of_1st_Point projection parameter */
var SRSPPLongitudeOf1stPoint = C.GoString(C._SRS_PP_LONGITUDE_OF_1ST_POINT())

// /** Latitude_Of_2nd_Point projection parameter */
var SRSPPLatitudeOf2ndPoint = C.GoString(C._SRS_PP_LATITUDE_OF_2ND_POINT())

// /** Longitude_Of_2nd_Point projection parameter */
var SRSPPLongitudeOf2ndPoint = C.GoString(C._SRS_PP_LONGITUDE_OF_2ND_POINT())

// /** peg_point_latitude projection parameter */
var SRSPPPegPointLatitude = C.GoString(C._SRS_PP_PEG_POINT_LATITUDE())

// /** peg_point_longitude projection parameter */
var SRSPPPegPointLongitude = C.GoString(C._SRS_PP_PEG_POINT_LONGITUDE())

// /** peg_point_heading projection parameter */
var SRSPPPegPointHeading = C.GoString(C._SRS_PP_PEG_POINT_HEADING())

// /** peg_point_height projection parameter */
var SRSPPPegPointHeight = C.GoString(C._SRS_PP_PEG_POINT_HEIGHT())

// /** Linear unit Meter */
var SRSULMeter = C.GoString(C._SRS_UL_METER())

// /** Linear unit Foot (International) */
var SRSULFoot = C.GoString(C._SRS_UL_FOOT())

// /** Linear unit Foot (International) conversion factor to meter*/
var SRSULFootConv = C.GoString(C._SRS_UL_FOOT_CONV())

// /** Linear unit Foot */
var SRSULUSFoot = C.GoString(C._SRS_UL_US_FOOT())

// /** Linear unit Foot conversion factor to meter */
var SRSULUSFootConv = C.GoString(C._SRS_UL_US_FOOT_CONV())

// /** Linear unit Nautical Mile */
var SRSULNauticalMile = C.GoString(C._SRS_UL_NAUTICAL_MILE())

// /** Linear unit Nautical Mile conversion factor to meter */
var SRSULNauticalMileConv = C.GoString(C._SRS_UL_NAUTICAL_MILE_CONV())

// /** Linear unit Link */
var SRSULLink = C.GoString(C._SRS_UL_LINK())

// /** Linear unit Link conversion factor to meter */
var SRSULLinkConv = C.GoString(C._SRS_UL_LINK_CONV())

// /** Linear unit Chain */
var SRSULChain = C.GoString(C._SRS_UL_CHAIN())

// /** Linear unit Chain conversion factor to meter */
var SRSULChainConv = C.GoString(C._SRS_UL_CHAIN_CONV())

// /** Linear unit Rod */
var SRSULRod = C.GoString(C._SRS_UL_ROD())

// /** Linear unit Rod conversion factor to meter */
var SRSULRodConv = C.GoString(C._SRS_UL_ROD_CONV())

// /** Linear unit Link_Clarke */
var SRSULLinkClarke = C.GoString(C._SRS_UL_LINK_Clarke())

// /** Linear unit Link_Clarke conversion factor to meter */
var SRSULLinkClarkeConv = C.GoString(C._SRS_UL_LINK_Clarke_CONV())

// /** Linear unit Kilometer */
var SRSULKilometer = C.GoString(C._SRS_UL_KILOMETER())

// /** Linear unit Kilometer conversion factor to meter */
var SRSULKilometerConv = C.GoString(C._SRS_UL_KILOMETER_CONV())

// /** Linear unit Decimeter */
var SRSULDecimeter = C.GoString(C._SRS_UL_DECIMETER())

// /** Linear unit Decimeter conversion factor to meter */
var SRSULDecimeterConv = C.GoString(C._SRS_UL_DECIMETER_CONV())

// /** Linear unit Decimeter */
var SRSULCentimeter = C.GoString(C._SRS_UL_CENTIMETER())

// /** Linear unit Decimeter conversion factor to meter */
var SRSULCentimeterConv = C.GoString(C._SRS_UL_CENTIMETER_CONV())

// /** Linear unit Millimeter */
var SRSULMillimeter = C.GoString(C._SRS_UL_MILLIMETER())

// /** Linear unit Millimeter conversion factor to meter */
var SRSULMillimeterConv = C.GoString(C._SRS_UL_MILLIMETER_CONV())

// /** Linear unit Nautical_Mile_International */
var SRSULIntlNautMile = C.GoString(C._SRS_UL_INTL_NAUT_MILE())

// /** Linear unit Nautical_Mile_International conversion factor to meter */
var SRSULIntlNautMileConv = C.GoString(C._SRS_UL_INTL_NAUT_MILE_CONV())

// /** Linear unit Inch_International */
var SRSULIntlInch = C.GoString(C._SRS_UL_INTL_INCH())

// /** Linear unit Inch_International conversion factor to meter */
var SRSULIntlInchConv = C.GoString(C._SRS_UL_INTL_INCH_CONV())

// /** Linear unit Foot_International */
var SRSULIntlFoot = C.GoString(C._SRS_UL_INTL_FOOT())

// /** Linear unit Foot_International conversion factor to meter */
var SRSULIntlFootConv = C.GoString(C._SRS_UL_INTL_FOOT_CONV())

// /** Linear unit Yard_International */
var SRSULIntlYard = C.GoString(C._SRS_UL_INTL_YARD())

// /** Linear unit Yard_International conversion factor to meter */
var SRSULIntlYardConv = C.GoString(C._SRS_UL_INTL_YARD_CONV())

// /** Linear unit Statute_Mile_International */
var SRSULIntlStatMile = C.GoString(C._SRS_UL_INTL_STAT_MILE())

// /** Linear unit Statute_Mile_Internationalconversion factor to meter */
var SRSULIntlStatMileConv = C.GoString(C._SRS_UL_INTL_STAT_MILE_CONV())

// /** Linear unit Fathom_International */
var SRSULIntlFathom = C.GoString(C._SRS_UL_INTL_FATHOM())

// /** Linear unit Fathom_International conversion factor to meter */
var SRSULIntlFathomConv = C.GoString(C._SRS_UL_INTL_FATHOM_CONV())

// /** Linear unit Chain_International */
var SRSULIntlChain = C.GoString(C._SRS_UL_INTL_CHAIN())

// /** Linear unit Chain_International conversion factor to meter */
var SRSULIntlChainConv = C.GoString(C._SRS_UL_INTL_CHAIN_CONV())

// /** Linear unit Link_International */
var SRSULIntlLink = C.GoString(C._SRS_UL_INTL_LINK())

// /** Linear unit Link_International conversion factor to meter */
var SRSULIntlLinkConv = C.GoString(C._SRS_UL_INTL_LINK_CONV())

// /** Linear unit Inch_US_Surveyor */
var SRSULUSInch = C.GoString(C._SRS_UL_US_INCH())

// /** Linear unit Inch_US_Surveyor conversion factor to meter */
var SRSULUSInchConv = C.GoString(C._SRS_UL_US_INCH_CONV())

// /** Linear unit Yard_US_Surveyor */
var SRSULUSYard = C.GoString(C._SRS_UL_US_YARD())

// /** Linear unit Yard_US_Surveyor conversion factor to meter */
var SRSULUSYardConv = C.GoString(C._SRS_UL_US_YARD_CONV())

// /** Linear unit Chain_US_Surveyor */
var SRSULUSChain = C.GoString(C._SRS_UL_US_CHAIN())

// /** Linear unit Chain_US_Surveyor conversion factor to meter */
var SRSULUSChainConv = C.GoString(C._SRS_UL_US_CHAIN_CONV())

// /** Linear unit Statute_Mile_US_Surveyor */
var SRSULUSStatMile = C.GoString(C._SRS_UL_US_STAT_MILE())

// /** Linear unit Statute_Mile_US_Surveyor conversion factor to meter */
var SRSULUSStatMileConv = C.GoString(C._SRS_UL_US_STAT_MILE_CONV())

// /** Linear unit Yard_Indian */
var SRSULIndianYard = C.GoString(C._SRS_UL_INDIAN_YARD())

// /** Linear unit Yard_Indian conversion factor to meter */
var SRSULIndianYardConv = C.GoString(C._SRS_UL_INDIAN_YARD_CONV())

// /** Linear unit Foot_Indian */
var SRSULIndianFoot = C.GoString(C._SRS_UL_INDIAN_FOOT())

// /** Linear unit Foot_Indian conversion factor to meter */
var SRSULIndianFootConv = C.GoString(C._SRS_UL_INDIAN_FOOT_CONV())

// /** Linear unit Chain_Indian */
var SRSULIndianChain = C.GoString(C._SRS_UL_INDIAN_CHAIN())

// /** Linear unit Chain_Indian conversion factor to meter */
var SRSULIndianChainConv = C.GoString(C._SRS_UL_INDIAN_CHAIN_CONV())

// /** Angular unit degree */
var SRSUADegree = C.GoString(C._SRS_UA_DEGREE())

// /** Angular unit degree conversion factor to radians */
var SRSUADegreeConv = C.GoString(C._SRS_UA_DEGREE_CONV())

// /** Angular unit radian */
var SRSUARadian = C.GoString(C._SRS_UA_RADIAN())

// /** Prime meridian Greenwich */
var SRSPMGreenwich = C.GoString(C._SRS_PM_GREENWICH())

// /** North_American_Datum_1927 datum name */
var SRSDNNad27 = C.GoString(C._SRS_DN_NAD27())

// /** North_American_Datum_1983 datum name */
var SRSDNNad83 = C.GoString(C._SRS_DN_NAD83())

// /** WGS_1972 datum name */
var SRSDNWgs72 = C.GoString(C._SRS_DN_WGS72())

// /** WGS_1984 datum name */
var SRSDNWgs84 = C.GoString(C._SRS_DN_WGS84())

// /** Semi-major axis of the WGS84 ellipsoid */
const SRSWgs84SemiMajor = C.SRS_WGS84_SEMIMAJOR

// /** Inverse flattening of the WGS84 ellipsoid */
const SRSWgs84InvFlattening = C.SRS_WGS84_INVFLATTENING

// #ifndef SWIG
// /* -------------------------------------------------------------------- */
// /*      C Wrappers for C++ objects and methods.                         */
// /* -------------------------------------------------------------------- */

func osrSetPROJSearchPaths(paths CSLConstList) {
	cPaths := paths.cValue
	C.OSRSetPROJSearchPaths(cPaths)
}

func osrGetPROJSearchPaths() (result CSLConstList) {
	raw := C.OSRGetPROJSearchPaths()
	result = cslConstList(raw)
	return
}

func osrSetPROJAuxDbPaths(paths CSLConstList) {
	cPaths := paths.cValue
	C.OSRSetPROJAuxDbPaths(cPaths)
}

func osrGetPROJAuxDbPaths() (result CSLConstList) {
	raw := C.OSRGetPROJAuxDbPaths()
	result = cslConstList(raw)
	return
}

func osrSetPROJEnableNetwork(enabled int) {
	C.OSRSetPROJEnableNetwork(C.int(enabled))
}

func osrGetPROJEnableNetwork() (result int) {
	result = int(C.OSRGetPROJEnableNetwork())
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

func osrNewSpatialReference(wkt string) (result OGRSpatialReference) {
	var cs *C.char
	if wkt != "" {
		cs = C.CString(wkt)
		defer C.free(unsafe.Pointer(cs))
	}
	result = OGRSpatialReference{cValue: C.OSRNewSpatialReference(cs)}
	return
}

func osrCloneGeogCS(sr OGRSpatialReference) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OSRCloneGeogCS(sr.cValue)}
	return
}

func osrClone(sr OGRSpatialReference) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OSRClone(sr.cValue)}
	return
}

func osrDestroySpatialReference(sr OGRSpatialReference) {
	C.OSRDestroySpatialReference(sr.cValue)
}

func osrReference(sr OGRSpatialReference) (result int) {
	result = int(C.OSRReference(sr.cValue))
	return
}

func osrDereference(sr OGRSpatialReference) (result int) {
	result = int(C.OSRDereference(sr.cValue))
	return
}

func osrRelease(sr OGRSpatialReference) {
	C.OSRRelease(sr.cValue)
}

func osrValidate(sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRValidate(sr.cValue))
	return
}

func osrImportFromEPSG(sr OGRSpatialReference, nCode int) (result OGRErr) {
	result = OGRErr(C.OSRImportFromEPSG(sr.cValue, C.int(nCode)))
	return
}

func osrImportFromEPSGA(sr OGRSpatialReference, nCode int) (result OGRErr) {
	result = OGRErr(C.OSRImportFromEPSGA(sr.cValue, C.int(nCode)))
	return
}

func osrImportFromWkt(sr OGRSpatialReference, wkt string) (result OGRErr) {
	cs := C.CString(wkt)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRImportFromWkt(sr.cValue, &cs))
	return
}

func osrImportFromProj4(sr OGRSpatialReference, proj4 string) (result OGRErr) {
	cs := C.CString(proj4)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRImportFromProj4(sr.cValue, cs))
	return
}

func osrImportFromESRI(sr OGRSpatialReference, lines CSLConstList) (result OGRErr) {
	cLines := lines.cValue
	result = OGRErr(C.OSRImportFromESRI(sr.cValue, cLines))
	return
}

func osrImportFromPCI(sr OGRSpatialReference, proj, units string, arParams []float64) (result OGRErr) {
	csProj := C.CString(proj)
	defer C.free(unsafe.Pointer(csProj))
	csUnits := C.CString(units)
	defer C.free(unsafe.Pointer(csUnits))
	var pParams *C.double
	if len(arParams) > 0 {
		pParams = (*C.double)(unsafe.Pointer(&arParams[0]))
	}
	result = OGRErr(C.OSRImportFromPCI(sr.cValue, csProj, csUnits, pParams))
	return
}

func osrImportFromUSGS(sr OGRSpatialReference, projSys, zone int, arParams []float64, datum int) (result OGRErr) {
	var pParams *C.double
	if len(arParams) > 0 {
		pParams = (*C.double)(unsafe.Pointer(&arParams[0]))
	}
	result = OGRErr(C.OSRImportFromUSGS(sr.cValue, C.long(projSys), C.long(zone), pParams, C.long(datum)))
	return
}

func osrImportFromXML(sr OGRSpatialReference, xmlString string) (result OGRErr) {
	cs := C.CString(xmlString)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRImportFromXML(sr.cValue, cs))
	return
}

func osrImportFromDict(sr OGRSpatialReference, dictFile, code string) (result OGRErr) {
	csDictFile := C.CString(dictFile)
	defer C.free(unsafe.Pointer(csDictFile))
	csCode := C.CString(code)
	defer C.free(unsafe.Pointer(csCode))
	result = OGRErr(C.OSRImportFromDict(sr.cValue, csDictFile, csCode))
	return
}

func osrImportFromPanorama(sr OGRSpatialReference, projSys, datum, ellipsoid int, arParams []float64) (result OGRErr) {
	var pParams *C.double
	if len(arParams) > 0 {
		pParams = (*C.double)(unsafe.Pointer(&arParams[0]))
	}
	result = OGRErr(C.OSRImportFromPanorama(sr.cValue, C.long(projSys), C.long(datum), C.long(ellipsoid), pParams))
	return
}

func osrImportFromOzi(sr OGRSpatialReference, lines CSLConstList) (result OGRErr) {
	cLines := lines.cValue
	result = OGRErr(C.OSRImportFromOzi(sr.cValue, cLines))
	return
}

func osrImportFromMICoordSys(sr OGRSpatialReference, coordSys string) (result OGRErr) {
	cs := C.CString(coordSys)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRImportFromMICoordSys(sr.cValue, cs))
	return
}

func osrImportFromERM(sr OGRSpatialReference, proj, datum, units string) (result OGRErr) {
	csProj := C.CString(proj)
	defer C.free(unsafe.Pointer(csProj))
	csDatum := C.CString(datum)
	defer C.free(unsafe.Pointer(csDatum))
	csUnits := C.CString(units)
	defer C.free(unsafe.Pointer(csUnits))
	result = OGRErr(C.OSRImportFromERM(sr.cValue, csProj, csDatum, csUnits))
	return
}

func osrImportFromUrl(sr OGRSpatialReference, url string) (result OGRErr) {
	cs := C.CString(url)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRImportFromUrl(sr.cValue, cs))
	return
}

func osrImportFromCF1(sr OGRSpatialReference, keyValues CSLConstList, units string) (result OGRErr) {
	cKeyValues := keyValues.cValue
	csUnits := C.CString(units)
	defer C.free(unsafe.Pointer(csUnits))
	result = OGRErr(C.OSRImportFromCF1(sr.cValue, cKeyValues, csUnits))
	return
}

func osrExportToWkt(sr OGRSpatialReference) (result string, status OGRErr) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToWkt(sr.cValue, &p))
	defer CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	status = ogrErr
	return
}

func osrExportToWktEx(sr OGRSpatialReference, options CSLConstList) (result string, status OGRErr) {
	cOptions := options.cValue
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToWktEx(sr.cValue, &p, cOptions))
	defer CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	status = ogrErr
	return
}

func osrExportToPrettyWkt(sr OGRSpatialReference, simplify int) (result string, status OGRErr) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToPrettyWkt(sr.cValue, &p, C.int(simplify)))
	defer CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	status = ogrErr
	return
}

func osrExportToPROJJSON(sr OGRSpatialReference, options CSLConstList) (result string, status OGRErr) {
	cOptions := options.cValue
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToPROJJSON(sr.cValue, &p, cOptions))
	defer CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	status = ogrErr
	return
}

func osrExportToProj4(sr OGRSpatialReference) (result string, status OGRErr) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToProj4(sr.cValue, &p))
	defer CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	status = ogrErr
	return
}

func osrExportToPCI(sr OGRSpatialReference) (proj, units string, params []float64, status OGRErr) {
	var cProj, cUnits *C.char
	var cParams *C.double
	ogrErr := OGRErr(C.OSRExportToPCI(sr.cValue, &cProj, &cUnits, &cParams))
	defer CPLFree(unsafe.Pointer(cProj))
	defer CPLFree(unsafe.Pointer(cUnits))
	defer CPLFree(unsafe.Pointer(cParams))
	proj = C.GoString(cProj)
	units = C.GoString(cUnits)
	const nPCIParams = 17
	params = make([]float64, nPCIParams)
	for i := 0; i < nPCIParams; i++ {
		params[i] = float64(*(*C.double)(unsafe.Add(unsafe.Pointer(cParams), uintptr(i)*unsafe.Sizeof(*cParams))))
	}
	status = ogrErr
	return
}

func osrExportToUSGS(sr OGRSpatialReference) (projSys, zone int, params []float64, datum int, status OGRErr) {
	var cProjSys, cZone, cDatum C.long
	var cParams *C.double
	ogrErr := OGRErr(C.OSRExportToUSGS(sr.cValue, &cProjSys, &cZone, &cParams, &cDatum))
	defer CPLFree(unsafe.Pointer(cParams))
	projSys = int(cProjSys)
	zone = int(cZone)
	datum = int(cDatum)
	const nUSGSParams = 15
	params = make([]float64, nUSGSParams)
	for i := 0; i < nUSGSParams; i++ {
		params[i] = float64(*(*C.double)(unsafe.Add(unsafe.Pointer(cParams), uintptr(i)*unsafe.Sizeof(*cParams))))
	}
	status = ogrErr
	return
}

func osrExportToXML(sr OGRSpatialReference, dialect string) (result string, status OGRErr) {
	var cs *C.char
	if dialect != "" {
		cs = C.CString(dialect)
		defer C.free(unsafe.Pointer(cs))
	}
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToXML(sr.cValue, &p, cs))
	defer CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	status = ogrErr
	return
}

func osrExportToPanorama(sr OGRSpatialReference) (projSys, datum, ellipsoid, zone int, params []float64, status OGRErr) {
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
	status = ogrErr
	return
}

func osrExportToMICoordSys(sr OGRSpatialReference) (result string, status OGRErr) {
	var p *C.char
	ogrErr := OGRErr(C.OSRExportToMICoordSys(sr.cValue, &p))
	defer CPLFree(unsafe.Pointer(p))
	result = C.GoString(p)
	status = ogrErr
	return
}

func osrExportToERM(sr OGRSpatialReference) (proj, datum, units string, status OGRErr) {
	cProj := make([]C.char, 32)
	cDatum := make([]C.char, 32)
	cUnits := make([]C.char, 32)
	ogrErr := OGRErr(C.OSRExportToERM(sr.cValue, &cProj[0], &cDatum[0], &cUnits[0]))
	proj = C.GoString(&cProj[0])
	datum = C.GoString(&cDatum[0])
	units = C.GoString(&cUnits[0])
	status = ogrErr
	return
}

func osrExportToCF1(sr OGRSpatialReference, options CSLConstList) (gridMappingName string, keyValues CSLConstList, units string, status OGRErr) {
	cOptions := options.cValue
	var cGridMappingName *C.char
	var cKeyValues **C.char
	var cUnits *C.char
	ogrErr := OGRErr(C.OSRExportToCF1(sr.cValue, &cGridMappingName, &cKeyValues, &cUnits, cOptions))
	defer CPLFree(unsafe.Pointer(cGridMappingName))
	defer CPLFree(unsafe.Pointer(cUnits))
	gridMappingName = C.GoString(cGridMappingName)
	units = C.GoString(cUnits)
	keyValues = cslConstList(cKeyValues)
	status = ogrErr
	return
}

func osrMorphToESRI(sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRMorphToESRI(sr.cValue))
	return
}

func osrMorphFromESRI(sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRMorphFromESRI(sr.cValue))
	return
}

func osrStripVertical(sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRStripVertical(sr.cValue))
	return
}

func osrConvertToOtherProjection(sr OGRSpatialReference, targetProjection string, options CSLConstList) (result OGRSpatialReference) {
	cs := C.CString(targetProjection)
	defer C.free(unsafe.Pointer(cs))
	cOptions := options.cValue
	result = OGRSpatialReference{cValue: C.OSRConvertToOtherProjection(sr.cValue, cs, cOptions)}
	return
}

func osrGetName(sr OGRSpatialReference) (result string) {
	result = C.GoString(C.OSRGetName(sr.cValue))
	return
}

func osrGetCelestialBodyName(sr OGRSpatialReference) (result string) {
	result = C.GoString(C.OSRGetCelestialBodyName(sr.cValue))
	return
}

func osrSetAttrValue(sr OGRSpatialReference, nodePath, newValue string) (result OGRErr) {
	csPath := C.CString(nodePath)
	defer C.free(unsafe.Pointer(csPath))
	csValue := C.CString(newValue)
	defer C.free(unsafe.Pointer(csValue))
	result = OGRErr(C.OSRSetAttrValue(sr.cValue, csPath, csValue))
	return
}

func osrGetAttrValue(sr OGRSpatialReference, name string, iChild int) (result string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OSRGetAttrValue(sr.cValue, cs, C.int(iChild)))
	return
}

func osrSetAngularUnits(sr OGRSpatialReference, name string, inDegrees float64) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetAngularUnits(sr.cValue, cs, C.double(inDegrees)))
	return
}

func osrGetAngularUnits(sr OGRSpatialReference) (result float64, name string) {
	var pName *C.char
	result = float64(C.OSRGetAngularUnits(sr.cValue, &pName))
	name = C.GoString(pName)
	return
}

func osrSetLinearUnits(sr OGRSpatialReference, name string, inMeters float64) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetLinearUnits(sr.cValue, cs, C.double(inMeters)))
	return
}

func osrSetTargetLinearUnits(sr OGRSpatialReference, targetKey, name string, inMeters float64) (result OGRErr) {
	csKey := C.CString(targetKey)
	defer C.free(unsafe.Pointer(csKey))
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	result = OGRErr(C.OSRSetTargetLinearUnits(sr.cValue, csKey, csName, C.double(inMeters)))
	return
}

func osrSetLinearUnitsAndUpdateParameters(sr OGRSpatialReference, name string, inMeters float64) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetLinearUnitsAndUpdateParameters(sr.cValue, cs, C.double(inMeters)))
	return
}

func osrGetLinearUnits(sr OGRSpatialReference) (result float64, name string) {
	var pName *C.char
	result = float64(C.OSRGetLinearUnits(sr.cValue, &pName))
	name = C.GoString(pName)
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

func osrGetPrimeMeridian(sr OGRSpatialReference) (result float64, name string) {
	var pName *C.char
	result = float64(C.OSRGetPrimeMeridian(sr.cValue, &pName))
	name = C.GoString(pName)
	return
}

func osrIsGeographic(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsGeographic(sr.cValue) != 0
	return
}

func osrIsDerivedGeographic(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsDerivedGeographic(sr.cValue) != 0
	return
}

func osrIsLocal(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsLocal(sr.cValue) != 0
	return
}

func osrIsProjected(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsProjected(sr.cValue) != 0
	return
}

func osrIsDerivedProjected(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsDerivedProjected(sr.cValue) != 0
	return
}

func osrIsCompound(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsCompound(sr.cValue) != 0
	return
}

func osrIsGeocentric(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsGeocentric(sr.cValue) != 0
	return
}

func osrIsVertical(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsVertical(sr.cValue) != 0
	return
}

func osrIsDynamic(sr OGRSpatialReference) (result bool) {
	result = C.OSRIsDynamic(sr.cValue) != 0
	return
}

func osrHasPointMotionOperation(sr OGRSpatialReference) (result bool) {
	result = C.OSRHasPointMotionOperation(sr.cValue) != 0
	return
}

func osrIsSameGeogCS(sr, other OGRSpatialReference) (result bool) {
	result = C.OSRIsSameGeogCS(sr.cValue, other.cValue) != 0
	return
}

func osrIsSameVertCS(sr, other OGRSpatialReference) (result bool) {
	result = C.OSRIsSameVertCS(sr.cValue, other.cValue) != 0
	return
}

func osrIsSame(sr, other OGRSpatialReference) (result bool) {
	result = C.OSRIsSame(sr.cValue, other.cValue) != 0
	return
}

func osrIsSameEx(sr, other OGRSpatialReference, options CSLConstList) (result bool) {
	cOptions := options.cValue
	result = C.OSRIsSameEx(sr.cValue, other.cValue, cOptions) != 0
	return
}

func osrSetCoordinateEpoch(sr OGRSpatialReference, epoch float64) {
	C.OSRSetCoordinateEpoch(sr.cValue, C.double(epoch))
}

func osrGetCoordinateEpoch(sr OGRSpatialReference) (result float64) {
	result = float64(C.OSRGetCoordinateEpoch(sr.cValue))
	return
}

func osrSetLocalCS(sr OGRSpatialReference, name string) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetLocalCS(sr.cValue, cs))
	return
}

func osrSetProjCS(sr OGRSpatialReference, name string) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetProjCS(sr.cValue, cs))
	return
}

func osrSetGeocCS(sr OGRSpatialReference, name string) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetGeocCS(sr.cValue, cs))
	return
}

func osrSetWellKnownGeogCS(sr OGRSpatialReference, name string) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetWellKnownGeogCS(sr.cValue, cs))
	return
}

func osrSetFromUserInput(sr OGRSpatialReference, definition string) (result OGRErr) {
	cs := C.CString(definition)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetFromUserInput(sr.cValue, cs))
	return
}

func osrSetFromUserInputEx(sr OGRSpatialReference, definition string, options CSLConstList) (result OGRErr) {
	cs := C.CString(definition)
	defer C.free(unsafe.Pointer(cs))
	cOptions := options.cValue
	result = OGRErr(C.OSRSetFromUserInputEx(sr.cValue, cs, cOptions))
	return
}

func osrCopyGeogCSFrom(sr, src OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRCopyGeogCSFrom(sr.cValue, src.cValue))
	return
}

func osrSetTOWGS84(sr OGRSpatialReference, dx, dy, dz, ex, ey, ez, ppm float64) (result OGRErr) {
	result = OGRErr(C.OSRSetTOWGS84(sr.cValue, C.double(dx), C.double(dy), C.double(dz), C.double(ex), C.double(ey), C.double(ez), C.double(ppm)))
	return
}

func osrGetTOWGS84(sr OGRSpatialReference) (params []float64, status OGRErr) {
	const n = 7
	cParams := make([]C.double, n)
	ogrErr := OGRErr(C.OSRGetTOWGS84(sr.cValue, &cParams[0], C.int(n)))
	params = make([]float64, n)
	for i := 0; i < n; i++ {
		params[i] = float64(cParams[i])
	}
	status = ogrErr
	return
}

func osrAddGuessedTOWGS84(sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRAddGuessedTOWGS84(sr.cValue))
	return
}

func osrSetCompoundCS(sr OGRSpatialReference, name string, horizSRS, vertSRS OGRSpatialReference) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetCompoundCS(sr.cValue, cs, horizSRS.cValue, vertSRS.cValue))
	return
}

func osrPromoteTo3D(sr OGRSpatialReference, name string) (result OGRErr) {
	var cs *C.char
	if name != "" {
		cs = C.CString(name)
		defer C.free(unsafe.Pointer(cs))
	}
	result = OGRErr(C.OSRPromoteTo3D(sr.cValue, cs))
	return
}

func osrDemoteTo2D(sr OGRSpatialReference, name string) (result OGRErr) {
	var cs *C.char
	if name != "" {
		cs = C.CString(name)
		defer C.free(unsafe.Pointer(cs))
	}
	result = OGRErr(C.OSRDemoteTo2D(sr.cValue, cs))
	return
}

func osrSetGeogCS(sr OGRSpatialReference, geogName, datumName, ellipsoidName string, semiMajor, invFlattening float64, pmName string, pmOffset float64, units string, convertToRadians float64) (result OGRErr) {
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
	result = OGRErr(C.OSRSetGeogCS(sr.cValue, csGeog, csDatum, csEllipsoid, C.double(semiMajor), C.double(invFlattening), csPM, C.double(pmOffset), csUnits, C.double(convertToRadians)))
	return
}

func osrSetVertCS(sr OGRSpatialReference, vertCSName, vertDatumName string, vertDatumType int) (result OGRErr) {
	csName := C.CString(vertCSName)
	defer C.free(unsafe.Pointer(csName))
	csDatum := C.CString(vertDatumName)
	defer C.free(unsafe.Pointer(csDatum))
	result = OGRErr(C.OSRSetVertCS(sr.cValue, csName, csDatum, C.int(vertDatumType)))
	return
}

func osrGetSemiMajor(sr OGRSpatialReference) (result float64, status OGRErr) {
	var ogrErr C.OGRErr
	result = float64(C.OSRGetSemiMajor(sr.cValue, &ogrErr))
	status = OGRErr(ogrErr)
	return
}

func osrGetSemiMinor(sr OGRSpatialReference) (result float64, status OGRErr) {
	var ogrErr C.OGRErr
	result = float64(C.OSRGetSemiMinor(sr.cValue, &ogrErr))
	status = OGRErr(ogrErr)
	return
}

func osrGetInvFlattening(sr OGRSpatialReference) (result float64, status OGRErr) {
	var ogrErr C.OGRErr
	result = float64(C.OSRGetInvFlattening(sr.cValue, &ogrErr))
	status = OGRErr(ogrErr)
	return
}

func osrSetAuthority(sr OGRSpatialReference, targetKey, authority string, code int) (result OGRErr) {
	csKey := C.CString(targetKey)
	defer C.free(unsafe.Pointer(csKey))
	csAuth := C.CString(authority)
	defer C.free(unsafe.Pointer(csAuth))
	result = OGRErr(C.OSRSetAuthority(sr.cValue, csKey, csAuth, C.int(code)))
	return
}

func osrGetAuthorityCode(sr OGRSpatialReference, targetKey string) (result string) {
	cs := C.CString(targetKey)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OSRGetAuthorityCode(sr.cValue, cs))
	return
}

func osrGetAuthorityName(sr OGRSpatialReference, targetKey string) (result string) {
	cs := C.CString(targetKey)
	defer C.free(unsafe.Pointer(cs))
	result = C.GoString(C.OSRGetAuthorityName(sr.cValue, cs))
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

func osrSetProjection(sr OGRSpatialReference, projection string) (result OGRErr) {
	cs := C.CString(projection)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetProjection(sr.cValue, cs))
	return
}

func osrSetProjParm(sr OGRSpatialReference, name string, value float64) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetProjParm(sr.cValue, cs, C.double(value)))
	return
}

func osrGetProjParm(sr OGRSpatialReference, name string, dfDefault float64) (result float64, status OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	var ogrErr C.OGRErr
	result = float64(C.OSRGetProjParm(sr.cValue, cs, C.double(dfDefault), &ogrErr))
	status = OGRErr(ogrErr)
	return
}

func osrSetNormProjParm(sr OGRSpatialReference, name string, value float64) (result OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetNormProjParm(sr.cValue, cs, C.double(value)))
	return
}

func osrGetNormProjParm(sr OGRSpatialReference, name string, dfDefault float64) (result float64, status OGRErr) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	var ogrErr C.OGRErr
	result = float64(C.OSRGetNormProjParm(sr.cValue, cs, C.double(dfDefault), &ogrErr))
	status = OGRErr(ogrErr)
	return
}

func osrSetUTM(sr OGRSpatialReference, zone, north int) (result OGRErr) {
	result = OGRErr(C.OSRSetUTM(sr.cValue, C.int(zone), C.int(north)))
	return
}

func osrGetUTMZone(sr OGRSpatialReference) (zone int, north bool) {
	var bNorth C.int
	zone = int(C.OSRGetUTMZone(sr.cValue, &bNorth))
	north = bNorth != 0
	return
}

func osrSetStatePlane(sr OGRSpatialReference, zone, nad83 int) (result OGRErr) {
	result = OGRErr(C.OSRSetStatePlane(sr.cValue, C.int(zone), C.int(nad83)))
	return
}

func osrSetStatePlaneWithUnits(sr OGRSpatialReference, zone, nad83 int, unitName string, unit float64) (result OGRErr) {
	cs := C.CString(unitName)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetStatePlaneWithUnits(sr.cValue, C.int(zone), C.int(nad83), cs, C.double(unit)))
	return
}

func osrAutoIdentifyEPSG(sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRAutoIdentifyEPSG(sr.cValue))
	return
}

func osrFindMatches(sr OGRSpatialReference, options CSLConstList, count *int, matchConfidence *[]int) (result OGRSpatialReferences) {
	cOptions := options.cValue
	var n C.int
	var conf *C.int
	result = OGRSpatialReferences{cValue: C.OSRFindMatches(sr.cValue, cOptions, &n, &conf)}
	if count != nil {
		*count = int(n)
	}
	if conf != nil {
		if matchConfidence != nil {
			src := unsafe.Slice(conf, int(n))
			mc := make([]int, int(n))
			for i := range mc {
				mc[i] = int(src[i])
			}
			*matchConfidence = mc
		}
		CPLFree(unsafe.Pointer(conf))
	}
	return
}

func osrFreeSRSArray(pahSRS OGRSpatialReferences) {
	C.OSRFreeSRSArray(pahSRS.cValue)
}

func osrEPSGTreatsAsLatLong(sr OGRSpatialReference) (result bool) {
	result = C.OSREPSGTreatsAsLatLong(sr.cValue) != 0
	return
}

func osrEPSGTreatsAsNorthingEasting(sr OGRSpatialReference) (result bool) {
	result = C.OSREPSGTreatsAsNorthingEasting(sr.cValue) != 0
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

func osrGetAxesCount(sr OGRSpatialReference) (result int) {
	result = int(C.OSRGetAxesCount(sr.cValue))
	return
}

func osrSetAxes(sr OGRSpatialReference, targetKey, xAxisName string, xOrientation OGRAxisOrientation, yAxisName string, yOrientation OGRAxisOrientation) (result OGRErr) {
	csKey := C.CString(targetKey)
	defer C.free(unsafe.Pointer(csKey))
	csX := C.CString(xAxisName)
	defer C.free(unsafe.Pointer(csX))
	csY := C.CString(yAxisName)
	defer C.free(unsafe.Pointer(csY))
	result = OGRErr(C.OSRSetAxes(sr.cValue, csKey, csX, C.OGRAxisOrientation(xOrientation), csY, C.OGRAxisOrientation(yOrientation)))
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

func osrSetAxisMappingStrategy(sr OGRSpatialReference, strategy OSRAxisMappingStrategy) {
	C.OSRSetAxisMappingStrategy(sr.cValue, C.OSRAxisMappingStrategy(strategy))
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

func osrSetDataAxisToSRSAxisMapping(sr OGRSpatialReference, mapping []int) (result OGRErr) {
	if len(mapping) == 0 {
		return
	}
	cMapping := make([]C.int, len(mapping))
	for i, m := range mapping {
		cMapping[i] = C.int(m)
	}
	result = OGRErr(C.OSRSetDataAxisToSRSAxisMapping(sr.cValue, C.int(len(mapping)), &cMapping[0]))
	return
}

/** Albers Conic Equal Area */
func osrSetACEA(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetACEA(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Azimuthal Equidistant */
func osrSetAE(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetAE(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Bonne */
func osrSetBonne(sr OGRSpatialReference, dfStandardParallel, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetBonne(sr.cValue, C.double(dfStandardParallel), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Cylindrical Equal Area */
func osrSetCEA(sr OGRSpatialReference, dfStdP1, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetCEA(sr.cValue, C.double(dfStdP1), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Cassini-Soldner */
func osrSetCS(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetCS(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Equidistant Conic */
func osrSetEC(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetEC(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Eckert I-VI */
func osrSetEckert(sr OGRSpatialReference, nVariation int, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetEckert(sr.cValue, C.int(nVariation), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Eckert IV */
func osrSetEckertIV(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetEckertIV(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Eckert VI */
func osrSetEckertVI(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetEckertVI(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Equirectangular */
func osrSetEquirectangular(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetEquirectangular(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Equirectangular generalized form */
func osrSetEquirectangular2(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfPseudoStdParallel1, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetEquirectangular2(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfPseudoStdParallel1), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Gall Stereograpic */
func osrSetGS(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetGS(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Goode Homolosine */
func osrSetGH(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetGH(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Interrupted Goode Homolosine */
func osrSetIGH(sr OGRSpatialReference) (result OGRErr) {
	result = OGRErr(C.OSRSetIGH(sr.cValue))
	return
}

/** GEOS - Geostationary Satellite View */
func osrSetGEOS(sr OGRSpatialReference, dfCentralMeridian, dfSatelliteHeight, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetGEOS(sr.cValue, C.double(dfCentralMeridian), C.double(dfSatelliteHeight), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Gauss Schreiber Transverse Mercator */
func osrSetGaussSchreiberTMercator(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetGaussSchreiberTMercator(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Gnomonic */
func osrSetGnomonic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetGnomonic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
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
func osrSetHOM(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetHOM(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfAzimuth), C.double(dfRectToSkew), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

func osrSetHOMAC(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfAzimuth, dfRectToSkew, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetHOMAC(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfAzimuth), C.double(dfRectToSkew), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Hotine Oblique Mercator using two points on centerline */
func osrSetHOM2PNO(sr OGRSpatialReference, dfCenterLat, dfLat1, dfLong1, dfLat2, dfLong2, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetHOM2PNO(sr.cValue, C.double(dfCenterLat), C.double(dfLat1), C.double(dfLong1), C.double(dfLat2), C.double(dfLong2), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** International Map of the World Polyconic */
func osrSetIWMPolyconic(sr OGRSpatialReference, dfLat1, dfLat2, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetIWMPolyconic(sr.cValue, C.double(dfLat1), C.double(dfLat2), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Krovak Oblique Conic Conformal */
func osrSetKrovak(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfAzimuth, dfPseudoStdParallelLat, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetKrovak(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfAzimuth), C.double(dfPseudoStdParallelLat), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Lambert Azimuthal Equal-Area */
func osrSetLAEA(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetLAEA(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Lambert Conformal Conic */
func osrSetLCC(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetLCC(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Lambert Conformal Conic 1SP */
func osrSetLCC1SP(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetLCC1SP(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Lambert Conformal Conic (Belgium) */
func osrSetLCCB(sr OGRSpatialReference, dfStdP1, dfStdP2, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetLCCB(sr.cValue, C.double(dfStdP1), C.double(dfStdP2), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Miller Cylindrical */
func osrSetMC(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetMC(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Mercator */
func osrSetMercator(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetMercator(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Mercator 2SP */
func osrSetMercator2SP(sr OGRSpatialReference, dfStdP1, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetMercator2SP(sr.cValue, C.double(dfStdP1), C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Mollweide */
func osrSetMollweide(sr OGRSpatialReference, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetMollweide(sr.cValue, C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** New Zealand Map Grid */
func osrSetNZMG(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetNZMG(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Oblique Stereographic */
func osrSetOS(sr OGRSpatialReference, dfOriginLat, dfCMeridian, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetOS(sr.cValue, C.double(dfOriginLat), C.double(dfCMeridian), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Orthographic */
func osrSetOrthographic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetOrthographic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Polyconic */
func osrSetPolyconic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetPolyconic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Polar Stereographic */
func osrSetPS(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetPS(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Robinson */
func osrSetRobinson(sr OGRSpatialReference, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetRobinson(sr.cValue, C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Sinusoidal */
func osrSetSinusoidal(sr OGRSpatialReference, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetSinusoidal(sr.cValue, C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Stereographic */
func osrSetStereographic(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetStereographic(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Swiss Oblique Cylindrical */
func osrSetSOC(sr OGRSpatialReference, dfLatitudeOfOrigin, dfCentralMeridian, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetSOC(sr.cValue, C.double(dfLatitudeOfOrigin), C.double(dfCentralMeridian), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Transverse Mercator
 *
 * Special processing available for Transverse Mercator with GDAL &gt;= 1.10 and
 * PROJ &gt;= 4.8 : see OGRSpatialReference::exportToProj4().
 */
func osrSetTM(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetTM(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Transverse Mercator variant */
func osrSetTMVariant(sr OGRSpatialReference, pszVariantName string, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	cs := C.CString(pszVariantName)
	defer C.free(unsafe.Pointer(cs))
	result = OGRErr(C.OSRSetTMVariant(sr.cValue, cs, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Tunesia Mining Grid  */
func osrSetTMG(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetTMG(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Transverse Mercator (South Oriented) */
func osrSetTMSO(sr OGRSpatialReference, dfCenterLat, dfCenterLong, dfScale, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetTMSO(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong), C.double(dfScale), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** TPED (Two Point Equi Distant) */
func osrSetTPED(sr OGRSpatialReference, dfLat1, dfLong1, dfLat2, dfLong2, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetTPED(sr.cValue, C.double(dfLat1), C.double(dfLong1), C.double(dfLat2), C.double(dfLong2), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** VanDerGrinten */
func osrSetVDG(sr OGRSpatialReference, dfCenterLong, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetVDG(sr.cValue, C.double(dfCenterLong), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Wagner I \-- VII */
func osrSetWagner(sr OGRSpatialReference, nVariation int, dfCenterLat, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetWagner(sr.cValue, C.int(nVariation), C.double(dfCenterLat), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

/** Quadrilateralized Spherical Cube */
func osrSetQSC(sr OGRSpatialReference, dfCenterLat, dfCenterLong float64) (result OGRErr) {
	result = OGRErr(C.OSRSetQSC(sr.cValue, C.double(dfCenterLat), C.double(dfCenterLong)))
	return
}

/** Spherical, Cross-track, Height */
func osrSetSCH(sr OGRSpatialReference, dfPegLat, dfPegLong, dfPegHeading, dfPegHgt float64) (result OGRErr) {
	result = OGRErr(C.OSRSetSCH(sr.cValue, C.double(dfPegLat), C.double(dfPegLong), C.double(dfPegHeading), C.double(dfPegHgt)))
	return
}

/** Vertical Perspective / Near-sided Perspective */
func osrSetVerticalPerspective(sr OGRSpatialReference, dfTopoOriginLat, dfTopoOriginLon, dfTopoOriginHeight, dfViewPointHeight, dfFalseEasting, dfFalseNorthing float64) (result OGRErr) {
	result = OGRErr(C.OSRSetVerticalPerspective(sr.cValue, C.double(dfTopoOriginLat), C.double(dfTopoOriginLon), C.double(dfTopoOriginHeight), C.double(dfViewPointHeight), C.double(dfFalseEasting), C.double(dfFalseNorthing)))
	return
}

func osrCalcInvFlattening(semiMajor, semiMinor float64) (result float64) {
	result = float64(C.OSRCalcInvFlattening(C.double(semiMajor), C.double(semiMinor)))
	return
}

func osrCalcSemiMinorFromInvFlattening(semiMajor, invFlattening float64) (result float64) {
	result = float64(C.OSRCalcSemiMinorFromInvFlattening(C.double(semiMajor), C.double(invFlattening)))
	return
}

func osrCleanup() {
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

// OSRCRSInfos wraps a C OSRCRSInfo* array (its length carries the C
// result-count argument).
type OSRCRSInfos struct {
	cValue **C.OSRCRSInfo
}

type OSRCRSListParameters struct {
	cValue *C.OSRCRSListParameters
}

func osrGetCRSInfoListFromDatabase(authName string, params OSRCRSListParameters, count *int) (result OSRCRSInfos) {
	cs := C.CString(authName)
	defer C.free(unsafe.Pointer(cs))
	var n C.int
	result.cValue = C.OSRGetCRSInfoListFromDatabase(cs, params.cValue, &n)
	if count != nil {
		*count = int(n)
	}
	return
}

func osrDestroyCRSInfoList(list OSRCRSInfos) {
	C.OSRDestroyCRSInfoList(list.cValue)
}

func osrGetAuthorityListFromDatabase() (result CSLConstList) {
	raw := C.OSRGetAuthorityListFromDatabase()
	result = cslConstList(raw)
	return
}

/* -------------------------------------------------------------------- */
/*      OGRCoordinateTransform C API.                                   */
/* -------------------------------------------------------------------- */

func octNewCoordinateTransformation(source, target OGRSpatialReference) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTNewCoordinateTransformation(source.cValue, target.cValue)}
	return
}

func octNewCoordinateTransformationOptions() (result OGRCoordinateTransformationOptions) {
	result = OGRCoordinateTransformationOptions{cValue: C.OCTNewCoordinateTransformationOptions()}
	return
}

func octCoordinateTransformationOptionsSetOperation(opts OGRCoordinateTransformationOptions, pszCO string, bReverseCO int) (result bool) {
	cs := C.CString(pszCO)
	defer C.free(unsafe.Pointer(cs))
	result = C.OCTCoordinateTransformationOptionsSetOperation(opts.cValue, cs, C.int(bReverseCO)) != 0
	return
}

func octCoordinateTransformationOptionsSetAreaOfInterest(opts OGRCoordinateTransformationOptions, dfWestLongitudeDeg, dfSouthLatitudeDeg, dfEastLongitudeDeg, dfNorthLatitudeDeg float64) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetAreaOfInterest(opts.cValue, C.double(dfWestLongitudeDeg), C.double(dfSouthLatitudeDeg), C.double(dfEastLongitudeDeg), C.double(dfNorthLatitudeDeg)) != 0
	return
}

func octCoordinateTransformationOptionsSetDesiredAccuracy(opts OGRCoordinateTransformationOptions, dfAccuracy float64) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetDesiredAccuracy(opts.cValue, C.double(dfAccuracy)) != 0
	return
}

func octCoordinateTransformationOptionsSetBallparkAllowed(opts OGRCoordinateTransformationOptions, bAllowBallpark int) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetBallparkAllowed(opts.cValue, C.int(bAllowBallpark)) != 0
	return
}

func octCoordinateTransformationOptionsSetOnlyBest(opts OGRCoordinateTransformationOptions, bOnlyBest bool) (result bool) {
	result = C.OCTCoordinateTransformationOptionsSetOnlyBest(opts.cValue, C._Bool(bOnlyBest)) != 0
	return
}

func octDestroyCoordinateTransformationOptions(opts OGRCoordinateTransformationOptions) {
	C.OCTDestroyCoordinateTransformationOptions(opts.cValue)
}

func octNewCoordinateTransformationEx(source, target OGRSpatialReference, options OGRCoordinateTransformationOptions) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTNewCoordinateTransformationEx(source.cValue, target.cValue, options.cValue)}
	return
}

func octDestroyCoordinateTransformation(ct OGRCoordinateTransformation) {
	C.OCTDestroyCoordinateTransformation(ct.cValue)
}

func octClone(ct OGRCoordinateTransformation) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTClone(ct.cValue)}
	return
}

func octGetSourceCS(ct OGRCoordinateTransformation) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OCTGetSourceCS(ct.cValue)}
	return
}

func octGetTargetCS(ct OGRCoordinateTransformation) (result OGRSpatialReference) {
	result = OGRSpatialReference{cValue: C.OCTGetTargetCS(ct.cValue)}
	return
}

func octGetInverse(ct OGRCoordinateTransformation) (result OGRCoordinateTransformation) {
	result = OGRCoordinateTransformation{cValue: C.OCTGetInverse(ct.cValue)}
	return
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

func octTransformBounds(ct OGRCoordinateTransformation, xmin, ymin, xmax, ymax float64, densifyPts int) (outXmin, outYmin, outXmax, outYmax float64, result bool) {
	var coutXmin, coutYmin, coutXmax, coutYmax C.double
	result = C.OCTTransformBounds(ct.cValue, C.double(xmin), C.double(ymin), C.double(xmax), C.double(ymax), &coutXmin, &coutYmin, &coutXmax, &coutYmax, C.int(densifyPts)) != 0
	outXmin = float64(coutXmin)
	outYmin = float64(coutYmin)
	outXmax = float64(coutXmax)
	outYmax = float64(coutYmax)
	return
}
