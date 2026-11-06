package gdal

/*
#include "cpl_vsi_error_preamble.h"
*/
import "C"
import "unsafe"

type VSIErrorNum C.VSIErrorNum

const (
	VSIENone                      VSIErrorNum = C.VSIE_None
	VSIEFileError                 VSIErrorNum = C.VSIE_FileError
	VSIEHttpError                 VSIErrorNum = C.VSIE_HttpError
	VSIEObjectStorageGenericError VSIErrorNum = C.VSIE_ObjectStorageGenericError
	VSIEAccessDenied              VSIErrorNum = C.VSIE_AccessDenied
	VSIEBucketNotFound            VSIErrorNum = C.VSIE_BucketNotFound
	VSIEObjectNotFound            VSIErrorNum = C.VSIE_ObjectNotFound
	VSIEInvalidCredentials        VSIErrorNum = C.VSIE_InvalidCredentials
	VSIESignatureDoesNotMatch     VSIErrorNum = C.VSIE_SignatureDoesNotMatch
)

// Deprecated aliases kept for 1:1 parity with cpl_vsi_error.h.
const (
	VSIEAWSError                 VSIErrorNum = C.VSIE_AWSError
	VSIEAWSAccessDenied          VSIErrorNum = C.VSIE_AWSAccessDenied
	VSIEAWSBucketNotFound        VSIErrorNum = C.VSIE_AWSBucketNotFound
	VSIEAWSObjectNotFound        VSIErrorNum = C.VSIE_AWSObjectNotFound
	VSIEAWSInvalidCredentials    VSIErrorNum = C.VSIE_AWSInvalidCredentials
	VSIEAWSSignatureDoesNotMatch VSIErrorNum = C.VSIE_AWSSignatureDoesNotMatch
)

// void CPL_DLL VSIError(VSIErrorNum err_no, CPL_FORMAT_STRING(const char *fmt),
//                       ...) CPL_PRINT_FUNC_FORMAT(2, 3);

func vsiErrorReset() {
	C.VSIErrorReset()
}

func vsiGetLastErrorNo() (result VSIErrorNum) {
	result = VSIErrorNum(C.VSIGetLastErrorNo())
	return
}

func vsiGetLastErrorMsg() (result string) {
	result = C.GoString(C.VSIGetLastErrorMsg())
	return
}

// const char *VSIErrorNumToString(int eErr);
//
// Not wrapped: declared without CPL_DLL, so libgdal does not export it and the
// binding would fail to link.

func vsiToCPLError(eErrClass CPLErr, eDefaultErrorNo CPLErrorNum) (result int) {
	result = int(C.VSIToCPLError(C.CPLErr(eErrClass), C.CPLErrorNum(eDefaultErrorNo)))
	return
}

func vsiToCPLErrorWithMsg(eErrClass CPLErr, eDefaultErrorNo CPLErrorNum, msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.VSIToCPLErrorWithMsg(C.CPLErr(eErrClass), C.CPLErrorNum(eDefaultErrorNo), cs)
}
