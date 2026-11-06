package gdal

func VSIErrorReset() {
	vsiErrorReset()
}

func VSIGetLastErrorNo() (result VSIErrorNum) {
	result = vsiGetLastErrorNo()
	return
}

func VSIGetLastErrorMsg() (result string) {
	result = vsiGetLastErrorMsg()
	return
}

// VSIToCPLError re-emits the pending VSI error, if any, as a CPLError of class
// eErrClass. VSIE_FileError maps to eDefaultErrorNo; the other VSIE codes map
// to their CPLE counterpart. Reports whether a VSI error was pending.
func VSIToCPLError(eErrClass CPLErr, eDefaultErrorNo CPLErrorNum) (result bool) {
	result = vsiToCPLError(eErrClass, eDefaultErrorNo) != 0
	return
}

func VSIToCPLErrorWithMsg(eErrClass CPLErr, eDefaultErrorNo CPLErrorNum, msg string) {
	vsiToCPLErrorWithMsg(eErrClass, eDefaultErrorNo, msg)
}
