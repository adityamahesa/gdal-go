package gdal

import "unsafe"

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

func CPLEmergencyError(msg string) {
	cplEmergencyError(msg)
}

func CPLErrorReset() {
	cplErrorReset()
}

func CPLGetLastErrorNo() (result CPLErrorNum) {
	result = cplGetLastErrorNo()
	return
}

func CPLGetLastErrorType() (result CPLErr) {
	result = cplGetLastErrorType()
	return
}

func CPLGetLastErrorMsg() (result string) {
	result = cplGetLastErrorMsg()
	return
}

func CPLGetErrorCounter() (result uint32) {
	result = cplGetErrorCounter()
	return
}

func CPLGetErrorHandlerUserData() (result unsafe.Pointer) {
	result = cplGetErrorHandlerUserData()
	return
}

func CPLErrorSetState(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cplErrorSetState(eErrClass, errNo, msg)
}

func CPLCallPreviousHandler(eErrClass CPLErr, errNo CPLErrorNum, msg string) {
	cplCallPreviousHandler(eErrClass, errNo, msg)
}

func CPLCleanupErrorMutex() {
	cplCleanupErrorMutex()
}

func CPLTurnFailureIntoWarning(bOn int) {
	cplTurnFailureIntoWarning(bOn)
}

func CPLGetErrorHandler(ppUserData *unsafe.Pointer) (result CPLErrorHandler) {
	result = cplGetErrorHandler(ppUserData)
	return
}

func CPLSetErrorHandler(handler CPLErrorHandler) (result CPLErrorHandler) {
	result = cplSetErrorHandler(handler)
	return
}

func CPLSetErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) (result CPLErrorHandler) {
	result = cplSetErrorHandlerEx(handler, userdata)
	return
}

func CPLPushErrorHandler(handler CPLErrorHandler) {
	cplPushErrorHandler(handler)
}

func CPLPushErrorHandlerEx(handler CPLErrorHandler, userdata unsafe.Pointer) {
	cplPushErrorHandlerEx(handler, userdata)
}

func CPLSetCurrentErrorHandlerCatchDebug(bCatchDebug int) {
	cplSetCurrentErrorHandlerCatchDebug(bCatchDebug)
}

func CPLPopErrorHandler() {
	cplPopErrorHandler()
}
