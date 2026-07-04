package gdal

/*
#include "cpl_error.h"

static CPLErrorHandler cplLoggingErrorHandlerFn()         { return CPLLoggingErrorHandler; }
static CPLErrorHandler cplDefaultErrorHandlerFn()         { return CPLDefaultErrorHandler; }
static CPLErrorHandler cplQuietErrorHandlerFn()           { return CPLQuietErrorHandler; }
static CPLErrorHandler cplQuietWarningsErrorHandlerFn()   { return CPLQuietWarningsErrorHandler; }

const CPLErr _VALIDATE_POINTER_ERR = VALIDATE_POINTER_ERR;
*/
import "C"
import "unsafe"

type CPLErr C.CPLErr

const (
	CENone    CPLErr = C.CE_None
	CEDebug   CPLErr = C.CE_Debug
	CEWarning CPLErr = C.CE_Warning
	CEFailure CPLErr = C.CE_Failure
	CEFatal   CPLErr = C.CE_Fatal
)

func (e CPLErr) String() string {
	switch e {
	case CENone:
		return "CE_NONE"
	case CEDebug:
		return "CE_DEBUG"
	case CEWarning:
		return "CE_WARNING"
	case CEFailure:
		return "CE_FAILURE"
	case CEFatal:
		return "CE_FATAL"
	default:
		return "CE_UNKNOWN"
	}
}

type CPLErrorNum C.CPLErrorNum

const (
	CPLENone                      CPLErrorNum = C.CPLE_None
	CPLEAppDefined                CPLErrorNum = C.CPLE_AppDefined
	CPLEOutOfMemory               CPLErrorNum = C.CPLE_OutOfMemory
	CPLEFileIO                    CPLErrorNum = C.CPLE_FileIO
	CPLEOpenFailed                CPLErrorNum = C.CPLE_OpenFailed
	CPLEIllegalArg                CPLErrorNum = C.CPLE_IllegalArg
	CPLENotSupported              CPLErrorNum = C.CPLE_NotSupported
	CPLEAssertionFailed           CPLErrorNum = C.CPLE_AssertionFailed
	CPLENoWriteAccess             CPLErrorNum = C.CPLE_NoWriteAccess
	CPLEUserInterrupt             CPLErrorNum = C.CPLE_UserInterrupt
	CPLEObjectNull                CPLErrorNum = C.CPLE_ObjectNull
	CPLEHttpResponse              CPLErrorNum = C.CPLE_HttpResponse
	CPLEBucketNotFound            CPLErrorNum = C.CPLE_BucketNotFound
	CPLEObjectNotFound            CPLErrorNum = C.CPLE_ObjectNotFound
	CPLEAccessDenied              CPLErrorNum = C.CPLE_AccessDenied
	CPLEInvalidCredentials        CPLErrorNum = C.CPLE_InvalidCredentials
	CPLESignatureDoesNotMatch     CPLErrorNum = C.CPLE_SignatureDoesNotMatch
	CPLEObjectStorageGenericError CPLErrorNum = C.CPLE_ObjectStorageGenericError
)

func (e CPLErrorNum) String() string {
	switch e {
	case CPLENone:
		return "CPLE_NONE"
	case CPLEAppDefined:
		return "CPLE_APP_DEFINED"
	case CPLEOutOfMemory:
		return "CPLE_OUT_OF_MEMORY"
	case CPLEFileIO:
		return "CPLE_FILE_IO"
	case CPLEOpenFailed:
		return "CPLE_OPEN_FAILED"
	case CPLEIllegalArg:
		return "CPLE_ILLEGAL_ARG"
	case CPLENotSupported:
		return "CPLE_NOT_SUPPORTED"
	case CPLEAssertionFailed:
		return "CPLE_ASSERTION_FAILED"
	case CPLENoWriteAccess:
		return "CPLE_NO_WRITE_ACCESS"
	case CPLEUserInterrupt:
		return "CPLE_USER_INTERRUPT"
	case CPLEObjectNull:
		return "CPLE_OBJECT_NULL"
	case CPLEHttpResponse:
		return "CPLE_HTTP_RESPONSE"
	case CPLEBucketNotFound:
		return "CPLE_BUCKET_NOT_FOUND"
	case CPLEObjectNotFound:
		return "CPLE_OBJECT_NOT_FOUND"
	case CPLEAccessDenied:
		return "CPLE_ACCESS_DENIED"
	case CPLEInvalidCredentials:
		return "CPLE_INVALID_CREDENTIALS"
	case CPLESignatureDoesNotMatch:
		return "CPLE_SIGNATURE_DOES_NOT_MATCH"
	case CPLEObjectStorageGenericError:
		return "CPLE_OBJECT_STORAGE_GENERIC_ERROR"
	default:
		return "CPLE_UNKNOWN"
	}
}

// void CPL_DLL CPLError(CPLErr eErrClass, CPLErrorNum err_no,
//                       CPL_FORMAT_STRING(const char *fmt), ...)
//     CPL_PRINT_FUNC_FORMAT(3, 4);

// void CPL_DLL CPLErrorV(CPLErr, CPLErrorNum, const char *, va_list);

func cplEmergencyError(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.CPLEmergencyError(cs)
}

func CPLEmergencyError(msg string) {
	cplEmergencyError(msg)
}

func cplErrorReset() {
	C.CPLErrorReset()
}

func CPLErrorReset() {
	cplErrorReset()
}

func cplGetLastErrorNo() (result CPLErrorNum) {
	result = CPLErrorNum(C.CPLGetLastErrorNo())
	return
}

func CPLGetLastErrorNo() (result CPLErrorNum) {
	result = cplGetLastErrorNo()
	return
}

func cplGetLastErrorType() (result CPLErr) {
	result = CPLErr(C.CPLGetLastErrorType())
	return
}

func CPLGetLastErrorType() (result CPLErr) {
	result = cplGetLastErrorType()
	return
}

func cplGetLastErrorMsg() (result string) {
	result = C.GoString(C.CPLGetLastErrorMsg())
	return
}

func CPLGetLastErrorMsg() (result string) {
	result = cplGetLastErrorMsg()
	return
}

func cplGetErrorCounter() (result uint32) {
	result = uint32(C.CPLGetErrorCounter())
	return
}

func CPLGetErrorCounter() (result uint32) {
	result = cplGetErrorCounter()
	return
}

func cplGetErrorHandlerUserData() (result unsafe.Pointer) {
	result = C.CPLGetErrorHandlerUserData()
	return
}

func CPLGetErrorHandlerUserData() (result unsafe.Pointer) {
	result = cplGetErrorHandlerUserData()
	return
}

func cplErrorSetState(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.CPLErrorSetState(C.CPLErr(eErrClass), C.CPLErrorNum(errNo), cs)
}

func CPLErrorSetState(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cplErrorSetState(eErrClass, errNo, msg)
}

func cplCallPreviousHandler(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.CPLCallPreviousHandler(C.CPLErr(eErrClass), C.CPLErrorNum(errNo), cs)
}

func CPLCallPreviousHandler(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cplCallPreviousHandler(eErrClass, errNo, msg)
}

func cplCleanupErrorMutex() {
	C.CPLCleanupErrorMutex()
}

func CPLCleanupErrorMutex() {
	cplCleanupErrorMutex()
}

type CPLErrorHandler struct {
	cValue C.CPLErrorHandler
}

var CPLLoggingErrorHandler = CPLErrorHandler{
	cValue: C.cplLoggingErrorHandlerFn(),
}

var CPLDefaultErrorHandler = CPLErrorHandler{
	cValue: C.cplDefaultErrorHandlerFn(),
}

var CPLQuietErrorHandler = CPLErrorHandler{
	cValue: C.cplQuietErrorHandlerFn(),
}

var CPLQuietWarningsErrorHandler = CPLErrorHandler{
	cValue: C.cplQuietWarningsErrorHandlerFn(),
}

func cplTurnFailureIntoWarning(bOn int) {
	C.CPLTurnFailureIntoWarning(C.int(bOn))
}

func CPLTurnFailureIntoWarning(bOn int) {
	cplTurnFailureIntoWarning(bOn)
}

func cplGetErrorHandler(ppUserData *unsafe.Pointer) (result CPLErrorHandler) {
	result = CPLErrorHandler{
		cValue: C.CPLGetErrorHandler(ppUserData),
	}
	return
}

func CPLGetErrorHandler(ppUserData *unsafe.Pointer) (result CPLErrorHandler) {
	result = cplGetErrorHandler(ppUserData)
	return
}

func cplSetErrorHandler(handler CPLErrorHandler) (result CPLErrorHandler) {
	result = CPLErrorHandler{
		cValue: C.CPLSetErrorHandler(handler.cValue),
	}
	return
}

func CPLSetErrorHandler(handler CPLErrorHandler) (result CPLErrorHandler) {
	result = cplSetErrorHandler(handler)
	return
}

func cplSetErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) (result CPLErrorHandler) {
	result = CPLErrorHandler{
		cValue: C.CPLSetErrorHandlerEx(handler.cValue, userdata),
	}
	return
}

func CPLSetErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) (result CPLErrorHandler) {
	result = cplSetErrorHandlerEx(handler, userdata)
	return
}

func cplPushErrorHandler(handler CPLErrorHandler) {
	C.CPLPushErrorHandler(handler.cValue)
}

func CPLPushErrorHandler(handler CPLErrorHandler) {
	cplPushErrorHandler(handler)
}

func cplPushErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) {
	C.CPLPushErrorHandlerEx(handler.cValue, userdata)
}

func CPLPushErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) {
	cplPushErrorHandlerEx(handler, userdata)
}

func cplSetCurrentErrorHandlerCatchDebug(bCatchDebug int) {
	C.CPLSetCurrentErrorHandlerCatchDebug(C.int(bCatchDebug))
}

func CPLSetCurrentErrorHandlerCatchDebug(bCatchDebug int) {
	cplSetCurrentErrorHandlerCatchDebug(bCatchDebug)
}

func cplPopErrorHandler() {
	C.CPLPopErrorHandler()
}

func CPLPopErrorHandler() {
	cplPopErrorHandler()
}

// void CPL_DLL CPLDebug(const char *, CPL_FORMAT_STRING(const char *), ...)
//     CPL_PRINT_FUNC_FORMAT(2, 3);
// void CPL_DLL CPLDebugProgress(const char *, CPL_FORMAT_STRING(const char *),
//                               ...) CPL_PRINT_FUNC_FORMAT(2, 3);

/** Same as CPLDebug(), but expands to nothing for non-DEBUG builds.
 * @since GDAL 3.1
 */
// #define CPLDebugOnly(...) CPLDebug(__VA_ARGS__)

/*
	This line should be implementing CPLAssert behaviors
*/

const ValidatePointerErr = C._VALIDATE_POINTER_ERR
