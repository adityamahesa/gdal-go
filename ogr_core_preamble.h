#ifndef GDAL_GO_OGR_CORE_PREAMBLE_H
#define GDAL_GO_OGR_CORE_PREAMBLE_H

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

#endif /* GDAL_GO_OGR_CORE_PREAMBLE_H */
