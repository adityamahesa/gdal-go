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

// lastErrorOr is lastError for a call whose failure sentinel (NULL, -1, ...)
// was already observed: it never returns nil. Many GDAL functions fail without
// posting a CPLError (GDALOpenEx without GDAL_OF_VERBOSE_ERROR, VSIStatL,
// lookups that miss, ...), so lastError alone would report success. fn is the
// C function name, so the message can be looked up in the GDAL docs.
//
// Use plain lastError only where GDAL documents NULL as a legitimate value
// (end of iteration, absent color table, SQL without a result set, ...).
func lastErrorOr(fn string) error {
	if err := lastError(); err != nil {
		return err
	}
	return fmt.Errorf("[%s][%s] %s failed without a GDAL error message",
		CEFailure.String(), CPLENone.String(), fn)
}

// vsiLastErrorOr is lastErrorOr for VSI calls that can record a VSIError
// (VSIFOpenExL/VSIFOpenEx2L with setError, VSIStatExL with
// VSIStatSetErrorFlag, VSIFileFromMemBuffer). The VSI error context is
// separate from the CPL one, so the pending VSIError is first promoted with
// VSIToCPLError, which keeps its message and maps VSIE codes to CPLE codes.
// Callers must VSIErrorReset at entry so a stale VSIError cannot leak in.
func vsiLastErrorOr(fn string) error {
	VSIToCPLError(CEFailure, CPLEFileIO)
	return lastErrorOr(fn)
}

// cplErrNoValues is cplErr for functions whose CE_Warning means "no values
// returned" (statistics/histogram queries with force=0): the outputs are then
// unset, so the warning must surface as an error.
func cplErrNoValues(e CPLErr, fn string) error {
	if e == CEWarning {
		return fmt.Errorf("[%s][%s] %s: no values available (pass force=1 to compute them)",
			CEWarning.String(), CPLENone.String(), fn)
	}
	return cplErr(e)
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
