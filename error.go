package gdal

import (
	"fmt"
	"runtime"
)

// errScope pins the goroutine to its OS thread and installs a quiet error
// handler for the duration of one GDAL call. GDAL's error context and its
// handler stack are both thread-local, so without the pin the goroutine could
// be rescheduled between the C call and the CPLGetLastError* reads below and
// observe a different thread's context — reporting no error for a real
// failure, or popping a handler off an empty stack.
//
// The returned func must be deferred immediately, as the first statement of
// the function, so that it unwinds after every other deferred call.
func errScope() func() {
	runtime.LockOSThread()
	CPLPushErrorHandler(CPLQuietErrorHandler)
	CPLErrorReset()
	return func() {
		CPLPopErrorHandler()
		runtime.UnlockOSThread()
	}
}

func ogrError(e OGRErr) error {
	if e == OGRErrNone {
		return nil
	}
	msg := CPLGetLastErrorMsg()
	if msg != "" {
		msg = " " + msg
	}
	return fmt.Errorf("[%s][%s][%s]%s", CPLGetLastErrorType().String(),
		CPLGetLastErrorNo().String(), e.String(), msg)
}

func lastError() error {
	errType := CPLGetLastErrorType()
	if !(errType == CEFailure || errType == CEFatal) {
		return nil
	}
	errMsg := CPLGetLastErrorMsg()
	if errMsg != "" {
		errMsg = " " + errMsg
	}
	return fmt.Errorf("[%s][%s]%s", errType.String(),
		CPLGetLastErrorNo().String(), errMsg)
}

func cplErr(e CPLErr) error {
	if !(e == CEFailure || e == CEFatal) {
		return nil
	}
	msg := CPLGetLastErrorMsg()
	if msg != "" {
		msg = " " + msg
	}
	return fmt.Errorf("[%s][%s][%s]%s", CPLGetLastErrorType().String(),
		CPLGetLastErrorNo().String(), e.String(), msg)
}
