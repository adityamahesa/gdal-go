package gdal

import "fmt"

func ogrError(e OGRErr) error {
	defer CPLErrorReset()
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
	defer CPLErrorReset()
	msg := CPLGetLastErrorMsg()
	if msg != "" {
		msg = " " + msg
	}
	return fmt.Errorf("[%s][%s]%s", CPLGetLastErrorType().String(),
		CPLGetLastErrorNo().String(), msg)
}

func cplErr(e CPLErr) error {
	defer CPLErrorReset()
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
