package gdal

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

func WkbFlatten(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = wkbFlatten(x)
	return
}

func WkbHasZ(x OGRwkbGeometryType) bool {
	return wkbHasZ(x)
}

func WkbSetZ(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = wkbSetZ(x)
	return
}

func WkbHasM(x OGRwkbGeometryType) bool {
	return wkbHasM(x)
}

func WkbSetM(x OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = wkbSetM(x)
	return
}

func (gt OGRwkbGeometryType) ToName() (result string) {
	result = ogrGeometryTypeToName(gt)
	return
}

func (gt OGRwkbGeometryType) Merge(extra OGRwkbGeometryType) (result OGRwkbGeometryType) {
	result = ogrMergeGeometryTypes(gt, extra)
	return
}

func (gt OGRwkbGeometryType) MergeEx(extra OGRwkbGeometryType, allowPromotingToCurves int) (result OGRwkbGeometryType) {
	result = ogrMergeGeometryTypesEx(gt, extra, allowPromotingToCurves)
	return
}

func (gt OGRwkbGeometryType) Flatten() (result OGRwkbGeometryType) {
	result = ogrGTFlatten(gt)
	return
}

func (gt *OGRwkbGeometryType) SetZ() {
	*gt = ogrGTSetZ(*gt)
}

func (gt *OGRwkbGeometryType) SetM() {
	*gt = ogrGTSetM(*gt)
}

func (gt *OGRwkbGeometryType) SetModifier(bSetZ, bSetM int) {
	*gt = ogrGTSetModifier(*gt, bSetZ, bSetM)
}

func (gt OGRwkbGeometryType) HasZ() (result bool) {
	result = ogrGTHasZ(gt) != 0
	return
}

func (gt OGRwkbGeometryType) HasM() (result bool) {
	result = ogrGTHasM(gt) != 0
	return
}

func (gt OGRwkbGeometryType) IsSubClassOf(eSuperType OGRwkbGeometryType) (result bool) {
	result = ogrGTIsSubClassOf(gt, eSuperType) != 0
	return
}

func (gt OGRwkbGeometryType) IsCurve() (result bool) {
	result = ogrGTIsCurve(gt) != 0
	return
}

func (gt OGRwkbGeometryType) IsSurface() (result bool) {
	result = ogrGTIsSurface(gt) != 0
	return
}

func (gt OGRwkbGeometryType) IsNonLinear() (result bool) {
	result = ogrGTIsNonLinear(gt) != 0
	return
}

func (gt OGRwkbGeometryType) GetCollection() (result OGRwkbGeometryType) {
	result = ogrGTGetCollection(gt)
	return
}

func (gt OGRwkbGeometryType) GetSingle() (result OGRwkbGeometryType) {
	result = ogrGTGetSingle(gt)
	return
}

func (gt OGRwkbGeometryType) GetCurve() (result OGRwkbGeometryType) {
	result = ogrGTGetCurve(gt)
	return
}

func (gt OGRwkbGeometryType) GetLinear() (result OGRwkbGeometryType) {
	result = ogrGTGetLinear(gt)
	return
}

func OGRGetMS(fSec float32) (result int) {
	result = ogrGetMS(fSec)
	return
}

func OGRParseDate(input string, options int) (result OGRField, ok bool) {
	result = InitOGRField()
	ok = ogrParseDate(input, result, options) != 0
	return
}
