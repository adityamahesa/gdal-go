#ifndef GDAL_GO_OGR_CORE_PREAMBLE_H
#define GDAL_GO_OGR_CORE_PREAMBLE_H

#include "ogr_core.h"

static const char* _OLCRandomRead(void)          { return OLCRandomRead; }
static const char* _OLCSequentialWrite(void)     { return OLCSequentialWrite; }
static const char* _OLCRandomWrite(void)         { return OLCRandomWrite; }
static const char* _OLCFastSpatialFilter(void)   { return OLCFastSpatialFilter; }
static const char* _OLCFastFeatureCount(void)    { return OLCFastFeatureCount; }
static const char* _OLCFastGetExtent(void)       { return OLCFastGetExtent; }
static const char* _OLCFastGetExtent3D(void)     { return OLCFastGetExtent3D; }
static const char* _OLCCreateField(void)         { return OLCCreateField; }
static const char* _OLCDeleteField(void)         { return OLCDeleteField; }
static const char* _OLCReorderFields(void)       { return OLCReorderFields; }
static const char* _OLCAlterFieldDefn(void)      { return OLCAlterFieldDefn; }
static const char* _OLCAlterGeomFieldDefn(void)  { return OLCAlterGeomFieldDefn; }
static const char* _OLCTransactions(void)        { return OLCTransactions; }
static const char* _OLCDeleteFeature(void)       { return OLCDeleteFeature; }
static const char* _OLCUpsertFeature(void)       { return OLCUpsertFeature; }
static const char* _OLCUpdateFeature(void)       { return OLCUpdateFeature; }
static const char* _OLCFastSetNextByIndex(void)  { return OLCFastSetNextByIndex; }
static const char* _OLCStringsAsUTF8(void)       { return OLCStringsAsUTF8; }
static const char* _OLCIgnoreFields(void)        { return OLCIgnoreFields; }
static const char* _OLCCreateGeomField(void)     { return OLCCreateGeomField; }
static const char* _OLCCurveGeometries(void)     { return OLCCurveGeometries; }
static const char* _OLCMeasuredGeometries(void)  { return OLCMeasuredGeometries; }
static const char* _OLCZGeometries(void)         { return OLCZGeometries; }
static const char* _OLCRename(void)              { return OLCRename; }
static const char* _OLCFastGetArrowStream(void)  { return OLCFastGetArrowStream; }
static const char* _OLCFastWriteArrowBatch(void) { return OLCFastWriteArrowBatch; }

static const char* _ODsCCreateLayer(void)                     { return ODsCCreateLayer; }
static const char* _ODsCDeleteLayer(void)                     { return ODsCDeleteLayer; }
static const char* _ODsCCreateGeomFieldAfterCreateLayer(void) { return ODsCCreateGeomFieldAfterCreateLayer; }
static const char* _ODsCCurveGeometries(void)                 { return ODsCCurveGeometries; }
static const char* _ODsCTransactions(void)                    { return ODsCTransactions; }
static const char* _ODsCEmulatedTransactions(void)            { return ODsCEmulatedTransactions; }
static const char* _ODsCMeasuredGeometries(void)              { return ODsCMeasuredGeometries; }
static const char* _ODsCZGeometries(void)                     { return ODsCZGeometries; }
static const char* _ODsCRandomLayerRead(void)                 { return ODsCRandomLayerRead; }
static const char* _ODsCRandomLayerWrite(void)                { return ODsCRandomLayerWrite; }
static const char* _ODsCAddFieldDomain(void)                  { return ODsCAddFieldDomain; }
static const char* _ODsCDeleteFieldDomain(void)               { return ODsCDeleteFieldDomain; }
static const char* _ODsCUpdateFieldDomain(void)               { return ODsCUpdateFieldDomain; }

static const char* _ODrCCreateDataSource(void) { return ODrCCreateDataSource; }
static const char* _ODrCDeleteDataSource(void) { return ODrCDeleteDataSource; }

static const char* _OLMD_FID64(void) { return OLMD_FID64; }

static int _ALTER_ALL_FLAG(void)                 { return ALTER_ALL_FLAG; }
static int _ALTER_GEOM_FIELD_DEFN_ALL_FLAG(void) { return ALTER_GEOM_FIELD_DEFN_ALL_FLAG; }
static int _OGR_F_VAL_ALL(void)                  { return OGR_F_VAL_ALL; }

#endif /* GDAL_GO_OGR_CORE_PREAMBLE_H */
