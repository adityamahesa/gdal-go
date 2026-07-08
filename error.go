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
