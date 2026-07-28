package gdal

/*
#include "cpl_error_preamble.h"
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

// void CPL_DLL CPLError(CPLErr eErrClass, CPLErrorNum err_no,
//                       CPL_FORMAT_STRING(const char *fmt), ...)
//     CPL_PRINT_FUNC_FORMAT(3, 4);

// void CPL_DLL CPLErrorV(CPLErr, CPLErrorNum, const char *, va_list);

func cplEmergencyError(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.CPLEmergencyError(cs)
}

func cplErrorReset() {
	C.CPLErrorReset()
}

func cplGetLastErrorNo() (result CPLErrorNum) {
	result = CPLErrorNum(C.CPLGetLastErrorNo())
	return
}

func cplGetLastErrorType() (result CPLErr) {
	result = CPLErr(C.CPLGetLastErrorType())
	return
}

func cplGetLastErrorMsg() (result string) {
	result = C.GoString(C.CPLGetLastErrorMsg())
	return
}

func cplGetErrorCounter() (result uint32) {
	result = uint32(C.CPLGetErrorCounter())
	return
}

func cplGetErrorHandlerUserData() (result unsafe.Pointer) {
	result = C.CPLGetErrorHandlerUserData()
	return
}

func cplErrorSetState(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.CPLErrorSetState(C.CPLErr(eErrClass), C.CPLErrorNum(errNo), cs)
}

func cplCallPreviousHandler(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.CPLCallPreviousHandler(C.CPLErr(eErrClass), C.CPLErrorNum(errNo), cs)
}

func cplCleanupErrorMutex() {
	C.CPLCleanupErrorMutex()
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

func cplGetErrorHandler(ppUserData *unsafe.Pointer) (result CPLErrorHandler) {
	result = CPLErrorHandler{
		cValue: C.CPLGetErrorHandler(ppUserData),
	}
	return
}

func cplSetErrorHandler(handler CPLErrorHandler) (result CPLErrorHandler) {
	result = CPLErrorHandler{
		cValue: C.CPLSetErrorHandler(handler.cValue),
	}
	return
}

func cplSetErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) (result CPLErrorHandler) {
	result = CPLErrorHandler{
		cValue: C.CPLSetErrorHandlerEx(handler.cValue, userdata),
	}
	return
}

func cplPushErrorHandler(handler CPLErrorHandler) {
	C.CPLPushErrorHandler(handler.cValue)
}

func cplPushErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) {
	C.CPLPushErrorHandlerEx(handler.cValue, userdata)
}

func cplSetCurrentErrorHandlerCatchDebug(bCatchDebug int) {
	C.CPLSetCurrentErrorHandlerCatchDebug(C.int(bCatchDebug))
}

func cplPopErrorHandler() {
	C.CPLPopErrorHandler()
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

var ValidatePointerErr = C._VALIDATE_POINTER_ERR()
