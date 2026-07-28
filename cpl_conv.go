package gdal

/*
#include "cpl_conv_preamble.h"
*/
import "C"
import "unsafe"

/**
 * \file cpl_conv.h
 *
 * Various convenience functions for CPL.
 */

/* -------------------------------------------------------------------- */
/*      Runtime check of various configuration items.                   */
/* -------------------------------------------------------------------- */

func cplVerifyConfiguration() {
	C.CPLVerifyConfiguration()
}

func cplIsDebugEnabled() (result bool) {
	result = bool(C.CPLIsDebugEnabled())
	return
}

func cplGetConfigOption(key, dflt string) (result string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CPLGetConfigOption(cKey, cDflt))
	return
}

func cplGetThreadLocalConfigOption(key, dflt string) (result string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CPLGetThreadLocalConfigOption(cKey, cDflt))
	return
}

func cplGetGlobalConfigOption(key, dflt string) (result string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CPLGetGlobalConfigOption(cKey, cDflt))
	return
}

func cplSetConfigOption(key, value string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.CPLSetConfigOption(cKey, cValue)
}

func cplSetThreadLocalConfigOption(key, value string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.CPLSetThreadLocalConfigOption(cKey, cValue)
}

func cplDeclareKnownConfigOption(key, definition string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDefinition := C.CString(definition)
	defer C.free(unsafe.Pointer(cDefinition))
	C.CPLDeclareKnownConfigOption(cKey, cDefinition)
}

func cplGetKnownConfigOptions() (result CSLConstList) {
	raw := C.CPLGetKnownConfigOptions()
	result = cslConstList(raw)
	return
}

// CPLSetConfigOptionSubscriber (a C function-pointer typedef) and
// CPLSubscribeToSetConfigOption require a Go->C callback; both are deferred.

func cplUnsubscribeToSetConfigOption(subscriberId int) {
	C.CPLUnsubscribeToSetConfigOption(C.int(subscriberId))
}

func cplFreeConfig() {
	C.CPLFreeConfig()
}

func cplGetConfigOptions() (result CSLConstList) {
	raw := C.CPLGetConfigOptions()
	result = cslConstList(raw)
	return
}

func cplSetConfigOptions(options CSLConstList) {
	C.CPLSetConfigOptions(options.cValue)
}

func cplGetThreadLocalConfigOptions() (result CSLConstList) {
	raw := C.CPLGetThreadLocalConfigOptions()
	result = cslConstList(raw)
	return
}

func cplSetThreadLocalConfigOptions(options CSLConstList) {
	C.CPLSetThreadLocalConfigOptions(options.cValue)
}

func cplLoadConfigOptionsFromFile(filename string, overrideEnvVars int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	C.CPLLoadConfigOptionsFromFile(cFilename, C.int(overrideEnvVars))
}

func cplLoadConfigOptionsFromPredefinedFiles() {
	C.CPLLoadConfigOptionsFromPredefinedFiles()
}

/* -------------------------------------------------------------------- */
/*      Safe malloc() API.                                              */
/* -------------------------------------------------------------------- */

func cplMalloc(size uint64) (result unsafe.Pointer) {
	result = C.CPLMalloc(C.size_t(size))
	return
}

func cplCalloc(count, size uint64) (result unsafe.Pointer) {
	result = C.CPLCalloc(C.size_t(count), C.size_t(size))
	return
}

func cplRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	result = C.CPLRealloc(ptr, C.size_t(size))
	return
}

func cplStrdup(s string) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CPLStrdup(cs)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplStrlwr(s string) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.CPLStrlwr(cs)
	result = C.GoString(cs)
	return
}

/* -------------------------------------------------------------------- */
/*      Read a line from a text file, and strip of CR/LF.               */
/* -------------------------------------------------------------------- */

// The stdio FILE*-based readers CPLFGets and CPLReadLine are part of the legacy
// (non-virtualized) group and are skipped; use the VSILFILE-based CPLReadLineL/
// CPLReadLine2L/CPLReadLine3L covers below.

func cplReadLineL(file VSILFile) (result string) {
	result = C.GoString(C.CPLReadLineL(file.cValue))
	return
}

func cplReadLine2L(file VSILFile, maxCars int, options CSLConstList) (result string) {
	result = C.GoString(C.CPLReadLine2L(file.cValue, C.int(maxCars), options.cValue))
	return
}

func cplReadLine3L(file VSILFile, maxCars int, options CSLConstList) (result string, bufLength int) {
	var cBufLength C.int
	result = C.GoString(C.CPLReadLine3L(file.cValue, C.int(maxCars), &cBufLength, options.cValue))
	bufLength = int(cBufLength)
	return
}

/* -------------------------------------------------------------------- */
/*      Convert ASCII string to floating point number.                  */
/* -------------------------------------------------------------------- */

func cplAtof(s string) (result float64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = float64(C.CPLAtof(cs))
	return
}

func cplAtofDelim(s string, point byte) (result float64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = float64(C.CPLAtofDelim(cs, C.char(point)))
	return
}

func cplStrtod(s string) (result float64, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float64(C.CPLStrtod(cs, &endptr))
	rest = C.GoString(endptr)
	return
}

func cplStrtodM(s string) (result float64, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float64(C.CPLStrtodM(cs, &endptr))
	rest = C.GoString(endptr)
	return
}

func cplStrtodDelim(s string, point byte) (result float64, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float64(C.CPLStrtodDelim(cs, &endptr, C.char(point)))
	rest = C.GoString(endptr)
	return
}

func cplStrtof(s string) (result float32, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float32(C.CPLStrtof(cs, &endptr))
	rest = C.GoString(endptr)
	return
}

func cplStrtofDelim(s string, point byte) (result float32, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float32(C.CPLStrtofDelim(cs, &endptr, C.char(point)))
	rest = C.GoString(endptr)
	return
}

/* -------------------------------------------------------------------- */
/*      Convert number to string.                                       */
/* -------------------------------------------------------------------- */

func cplAtofM(s string) (result float64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = float64(C.CPLAtofM(cs))
	return
}

/* -------------------------------------------------------------------- */
/*      Read a numeric value from an ASCII character string.            */
/* -------------------------------------------------------------------- */

func cplScanString(s string, maxLength, skipLeadingSpaces, stripQuotes int) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CPLScanString(cs, C.int(maxLength), C.int(skipLeadingSpaces), C.int(stripQuotes))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplScanDouble(s string, length int) (result float64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = float64(C.CPLScanDouble(cs, C.int(length)))
	return
}

func cplScanLong(s string, length int) (result int64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = int64(C.CPLScanLong(cs, C.int(length)))
	return
}

func cplScanULong(s string, length int) (result uint64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = uint64(C.CPLScanULong(cs, C.int(length)))
	return
}

func cplScanUIntBig(s string, length int) (result GUIntBig) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = GUIntBig(C.CPLScanUIntBig(cs, C.int(length)))
	return
}

func cplAtoGIntBig(s string) (result GIntBig) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = GIntBig(C.CPLAtoGIntBig(cs))
	return
}

func cplAtoGIntBigEx(s string, warn int) (result GIntBig, overflow int) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var cOverflow C.int
	result = GIntBig(C.CPLAtoGIntBigEx(cs, C.int(warn), &cOverflow))
	overflow = int(cOverflow)
	return
}

func cplScanPointer(s string, length int) (result unsafe.Pointer) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = C.CPLScanPointer(cs, C.int(length))
	return
}

/* -------------------------------------------------------------------- */
/*      Print a value to an ASCII character string.                     */
/* -------------------------------------------------------------------- */

func cplPrintString(s string, maxLen int) (result string) {
	if maxLen <= 0 {
		return
	}
	buf := make([]byte, maxLen)
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	n := int(C.CPLPrintString((*C.char)(unsafe.Pointer(&buf[0])), cs, C.int(maxLen)))
	result = string(buf[:n])
	return
}

func cplPrintStringFill(s string, maxLen int) (result string) {
	if maxLen <= 0 {
		return
	}
	buf := make([]byte, maxLen)
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	n := int(C.CPLPrintStringFill((*C.char)(unsafe.Pointer(&buf[0])), cs, C.int(maxLen)))
	result = string(buf[:n])
	return
}

func cplPrintInt32(value GInt32, maxLen int) (result string) {
	if maxLen <= 0 {
		return
	}
	buf := make([]byte, maxLen)
	n := int(C.CPLPrintInt32((*C.char)(unsafe.Pointer(&buf[0])), C.GInt32(value), C.int(maxLen)))
	result = string(buf[:n])
	return
}

func cplPrintUIntBig(value GUIntBig, maxLen int) (result string) {
	if maxLen <= 0 {
		return
	}
	buf := make([]byte, maxLen)
	n := int(C.CPLPrintUIntBig((*C.char)(unsafe.Pointer(&buf[0])), C.GUIntBig(value), C.int(maxLen)))
	result = string(buf[:n])
	return
}

func cplPrintDouble(format string, value float64, locale string) (result string) {
	const bufSize = 512
	buf := make([]byte, bufSize)
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	cLocale := C.CString(locale)
	defer C.free(unsafe.Pointer(cLocale))
	n := int(C.CPLPrintDouble((*C.char)(unsafe.Pointer(&buf[0])), cFormat, C.double(value), cLocale))
	result = string(buf[:n])
	return
}

// CPLPrintTime operates on a "const struct tm *"; deferred pending a struct tm wrapper.

func cplPrintPointer(ptr unsafe.Pointer, maxLen int) (result string) {
	if maxLen <= 0 {
		return
	}
	buf := make([]byte, maxLen)
	n := int(C.CPLPrintPointer((*C.char)(unsafe.Pointer(&buf[0])), ptr, C.int(maxLen)))
	result = string(buf[:n])
	return
}

// CPLFormatReadableFileSize returns a std::string and is a C++-only cover; it is skipped.

/* -------------------------------------------------------------------- */
/*      Fetch a function from DLL / so.                                 */
/* -------------------------------------------------------------------- */

func cplGetSymbol(library, symbolName string) (result unsafe.Pointer) {
	cLibrary := C.CString(library)
	defer C.free(unsafe.Pointer(cLibrary))
	cSymbolName := C.CString(symbolName)
	defer C.free(unsafe.Pointer(cSymbolName))
	result = C.CPLGetSymbol(cLibrary, cSymbolName)
	return
}

/* -------------------------------------------------------------------- */
/*      Fetch executable path.                                          */
/* -------------------------------------------------------------------- */

func cplGetExecPath(maxLength int) (result string, ok bool) {
	if maxLength <= 0 {
		return
	}
	buf := make([]byte, maxLength)
	ret := int(C.CPLGetExecPath((*C.char)(unsafe.Pointer(&buf[0])), C.int(maxLength)))
	if ret == 0 {
		return
	}
	result = C.GoString((*C.char)(unsafe.Pointer(&buf[0])))
	ok = true
	return
}

/* -------------------------------------------------------------------- */
/*      Filename handling functions.                                    */
/* -------------------------------------------------------------------- */

func cplGetPath(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetPath(cPath))
	return
}

func cplGetDirname(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetDirname(cPath))
	return
}

func cplGetBasename(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetBasename(cPath))
	return
}

func cplGetExtension(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetExtension(cPath))
	return
}

func cplFormFilename(path, basename, extension string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cBasename := C.CString(basename)
	defer C.free(unsafe.Pointer(cBasename))
	cExtension := C.CString(extension)
	defer C.free(unsafe.Pointer(cExtension))
	result = C.GoString(C.CPLFormFilename(cPath, cBasename, cExtension))
	return
}

func cplFormCIFilename(path, basename, extension string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cBasename := C.CString(basename)
	defer C.free(unsafe.Pointer(cBasename))
	cExtension := C.CString(extension)
	defer C.free(unsafe.Pointer(cExtension))
	result = C.GoString(C.CPLFormCIFilename(cPath, cBasename, cExtension))
	return
}

func cplResetExtension(path, extension string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cExtension := C.CString(extension)
	defer C.free(unsafe.Pointer(cExtension))
	result = C.GoString(C.CPLResetExtension(cPath, cExtension))
	return
}

func cplProjectRelativeFilename(projectDir, secondaryFilename string) (result string) {
	cProjectDir := C.CString(projectDir)
	defer C.free(unsafe.Pointer(cProjectDir))
	cSecondaryFilename := C.CString(secondaryFilename)
	defer C.free(unsafe.Pointer(cSecondaryFilename))
	result = C.GoString(C.CPLProjectRelativeFilename(cProjectDir, cSecondaryFilename))
	return
}

func cplCleanTrailingSlash(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLCleanTrailingSlash(cPath))
	return
}

func cplGenerateTempFilename(stem string) (result string) {
	cStem := C.CString(stem)
	defer C.free(unsafe.Pointer(cStem))
	result = C.GoString(C.CPLGenerateTempFilename(cStem))
	return
}

func cplExpandTilde(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.CPLExpandTilde(cFilename))
	return
}

func cplLaunderForFilename(name, outputPath string) (result string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cOutputPath := C.CString(outputPath)
	defer C.free(unsafe.Pointer(cOutputPath))
	result = C.GoString(C.CPLLaunderForFilename(cName, cOutputPath))
	return
}

func cplGetCurrentDir() (result string) {
	raw := C.CPLGetCurrentDir()
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplGetFilename(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetFilename(cPath))
	return
}

func cplIsFilenameRelative(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.CPLIsFilenameRelative(cFilename))
	return
}

func cplExtractRelativePath(baseDir, target string) (result string, gotRelative int) {
	cBaseDir := C.CString(baseDir)
	defer C.free(unsafe.Pointer(cBaseDir))
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	var cGotRelative C.int
	result = C.GoString(C.CPLExtractRelativePath(cBaseDir, cTarget, &cGotRelative))
	gotRelative = int(cGotRelative)
	return
}

func cplCorrespondingPaths(oldFilename, newFilename string, fileList CSLConstList) (result CSLConstList) {
	cOld := C.CString(oldFilename)
	defer C.free(unsafe.Pointer(cOld))
	cNew := C.CString(newFilename)
	defer C.free(unsafe.Pointer(cNew))
	raw := C.CPLCorrespondingPaths(cOld, cNew, fileList.cValue)
	result = cslConstList(raw)
	return
}

func cplCheckForFile(filename string, siblings CSLConstList) (result int, corrected string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.CPLCheckForFile(cFilename, siblings.cValue))
	corrected = C.GoString(cFilename)
	return
}

func cplGetHomeDir() (result string) {
	result = C.GoString(C.CPLGetHomeDir())
	return
}

// The extern "C++" ...Safe variants (CPLGetPathSafe, CPLGetDirnameSafe, etc.)
// return std::string and are C++-only covers; they are skipped.

func cplHasPathTraversal(filename string) (result bool) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = bool(C.CPLHasPathTraversal(cFilename))
	return
}

func cplHasUnbalancedPathTraversal(filename string) (result bool) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = bool(C.CPLHasUnbalancedPathTraversal(cFilename))
	return
}

/* -------------------------------------------------------------------- */
/*      Find File Function                                              */
/* -------------------------------------------------------------------- */

// CPLFileFinder (a C function-pointer typedef), CPLPushFileFinder and
// CPLPopFileFinder require a Go->C callback; they are deferred.

func cplFindFile(class, basename string) (result string) {
	cClass := C.CString(class)
	defer C.free(unsafe.Pointer(cClass))
	cBasename := C.CString(basename)
	defer C.free(unsafe.Pointer(cBasename))
	result = C.GoString(C.CPLFindFile(cClass, cBasename))
	return
}

func cplDefaultFindFile(class, basename string) (result string) {
	cClass := C.CString(class)
	defer C.free(unsafe.Pointer(cClass))
	cBasename := C.CString(basename)
	defer C.free(unsafe.Pointer(cBasename))
	result = C.GoString(C.CPLDefaultFindFile(cClass, cBasename))
	return
}

func cplPushFinderLocation(location string) {
	cLocation := C.CString(location)
	defer C.free(unsafe.Pointer(cLocation))
	C.CPLPushFinderLocation(cLocation)
}

func cplPopFinderLocation() {
	C.CPLPopFinderLocation()
}

func cplFinderClean() {
	C.CPLFinderClean()
}

/* -------------------------------------------------------------------- */
/*      Safe version of stat().                                        */
/* -------------------------------------------------------------------- */

// CPLStat uses the legacy (non-virtualized) VSIStatBuf, which is skipped in
// cpl_vsi.go; CPLStat is therefore skipped too. Use VSIStatL()/VSIStatExL().

/* -------------------------------------------------------------------- */
/*      Reference counted file handle manager.                          */
/* -------------------------------------------------------------------- */

// The shared file handle manager (CPLSharedFileInfo, CPLOpenShared,
// CPLCloseShared, CPLGetSharedList, CPLDumpSharedList, CPLCleanupSharedFileMutex)
// is built on raw C FILE*, part of the legacy (non-virtualized) group skipped in
// cpl_vsi.go; it is skipped here for the same reason.

/* -------------------------------------------------------------------- */
/*      DMS to Dec to DMS conversion.                                   */
/* -------------------------------------------------------------------- */

func cplDMSToDec(is string) (result float64) {
	cIs := C.CString(is)
	defer C.free(unsafe.Pointer(cIs))
	result = float64(C.CPLDMSToDec(cIs))
	return
}

func cplDecToDMS(angle float64, axis string, precision int) (result string) {
	cAxis := C.CString(axis)
	defer C.free(unsafe.Pointer(cAxis))
	result = C.GoString(C.CPLDecToDMS(C.double(angle), cAxis, C.int(precision)))
	return
}

func cplPackedDMSToDec(packed float64) (result float64) {
	result = float64(C.CPLPackedDMSToDec(C.double(packed)))
	return
}

func cplDecToPackedDMS(dec float64) (result float64) {
	result = float64(C.CPLDecToPackedDMS(C.double(dec)))
	return
}

func cplStringToComplex(s string) (real, imag float64, ret CPLErr) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var cReal, cImag C.double
	ret = CPLErr(C.CPLStringToComplex(cs, &cReal, &cImag))
	real = float64(cReal)
	imag = float64(cImag)
	return
}

/* -------------------------------------------------------------------- */
/*      Misc other functions.                                           */
/* -------------------------------------------------------------------- */

func cplUnlinkTree(path string) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.CPLUnlinkTree(cPath))
	return
}

func cplCopyFile(newPath, oldPath string) (result int) {
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	result = int(C.CPLCopyFile(cNew, cOld))
	return
}

func cplCopyTree(newPath, oldPath string) (result int) {
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	result = int(C.CPLCopyTree(cNew, cOld))
	return
}

func cplMoveFile(newPath, oldPath string) (result int) {
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	result = int(C.CPLMoveFile(cNew, cOld))
	return
}

func cplSymlink(oldPath, newPath string, options CSLConstList) (result int) {
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	result = int(C.CPLSymlink(cOld, cNew, options.cValue))
	return
}

func cplGetRemainingFileDescriptorCount() (result int) {
	result = int(C.CPLGetRemainingFileDescriptorCount())
	return
}

/* -------------------------------------------------------------------- */
/*      Lock related functions.                                         */
/* -------------------------------------------------------------------- */

// Return code of CPLLockFileEx().
type CPLLockFileStatus C.CPLLockFileStatus

const (
	CLFSOk                   CPLLockFileStatus = C.CLFS_OK
	CLFSCannotCreateLock     CPLLockFileStatus = C.CLFS_CANNOT_CREATE_LOCK
	CLFSLockBusy             CPLLockFileStatus = C.CLFS_LOCK_BUSY
	CLFSAPIMisuse            CPLLockFileStatus = C.CLFS_API_MISUSE
	CLFSThreadCreationFailed CPLLockFileStatus = C.CLFS_THREAD_CREATION_FAILED
)

// Handle type returned by CPLLockFileEx().
type CPLLockFileHandle struct {
	cValue C.CPLLockFileHandle
}

func cplLockFileEx(lockFileName string, options CSLConstList) (handle CPLLockFileHandle, status CPLLockFileStatus) {
	cLockFileName := C.CString(lockFileName)
	defer C.free(unsafe.Pointer(cLockFileName))
	status = CPLLockFileStatus(C.CPLLockFileEx(cLockFileName, &handle.cValue, options.cValue))
	return
}

func cplUnlockFileEx(handle CPLLockFileHandle) {
	C.CPLUnlockFileEx(handle.cValue)
}

/* -------------------------------------------------------------------- */
/*      ZIP Creation.                                                   */
/* -------------------------------------------------------------------- */

// CPL_ZIP_API_OFFERED is a presence macro (no value) and is skipped.

func cplCreateZip(zipFilename string, options CSLConstList) (result unsafe.Pointer) {
	cZipFilename := C.CString(zipFilename)
	defer C.free(unsafe.Pointer(cZipFilename))
	result = C.CPLCreateZip(cZipFilename, options.cValue)
	return
}

func cplCreateFileInZip(zip unsafe.Pointer, filename string, options CSLConstList) (result CPLErr) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = CPLErr(C.CPLCreateFileInZip(zip, cFilename, options.cValue))
	return
}

func cplWriteFileInZip(zip unsafe.Pointer, buffer []byte) (result CPLErr) {
	result = CPLErr(C.CPLWriteFileInZip(zip, cBytes(buffer), C.int(len(buffer))))
	return
}

func cplCloseFileInZip(zip unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.CPLCloseFileInZip(zip))
	return
}

func cplAddFileInZip(zip unsafe.Pointer, archiveFilename, inputFilename string, fpInput VSILFile, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cArchiveFilename := C.CString(archiveFilename)
	defer C.free(unsafe.Pointer(cArchiveFilename))
	cInputFilename := C.CString(inputFilename)
	defer C.free(unsafe.Pointer(cInputFilename))
	result = CPLErr(C.CPLAddFileInZip(zip, cArchiveFilename, cInputFilename, fpInput.cValue, options.cValue, progress.cValue, progressData))
	return
}

func cplCloseZip(zip unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.CPLCloseZip(zip))
	return
}

/* -------------------------------------------------------------------- */
/*      ZLib compression                                                */
/* -------------------------------------------------------------------- */

func cplZLibDeflate(input []byte, level int) (result []byte) {
	var outBytes C.size_t
	raw := C.CPLZLibDeflate(cBytes(input), C.size_t(len(input)), C.int(level), nil, 0, &outBytes)
	if raw == nil {
		return
	}
	defer C.VSIFree(raw)
	result = C.GoBytes(raw, C.int(outBytes))
	return
}

func cplZLibInflate(input []byte) (result []byte) {
	var outBytes C.size_t
	raw := C.CPLZLibInflate(cBytes(input), C.size_t(len(input)), nil, 0, &outBytes)
	if raw == nil {
		return
	}
	defer C.VSIFree(raw)
	result = C.GoBytes(raw, C.int(outBytes))
	return
}

func cplZLibInflateEx(input []byte, allowResizeOutptr bool) (result []byte) {
	var outBytes C.size_t
	raw := C.CPLZLibInflateEx(cBytes(input), C.size_t(len(input)), nil, 0, C.bool(allowResizeOutptr), &outBytes)
	if raw == nil {
		return
	}
	defer C.VSIFree(raw)
	result = C.GoBytes(raw, C.int(outBytes))
	return
}

/* -------------------------------------------------------------------- */
/*      XML validation.                                                 */
/* -------------------------------------------------------------------- */

func cplValidateXML(xmlFilename, xsdFilename string, options CSLConstList) (result int) {
	cXML := C.CString(xmlFilename)
	defer C.free(unsafe.Pointer(cXML))
	cXSD := C.CString(xsdFilename)
	defer C.free(unsafe.Pointer(cXSD))
	result = int(C.CPLValidateXML(cXML, cXSD, options.cValue))
	return
}

/* -------------------------------------------------------------------- */
/*      Locale handling.                                                */
/* -------------------------------------------------------------------- */

// CPLsetlocale and CPLCleanupSetlocaleMutex lack CPL_DLL and are not part of the
// public exported ABI; they are skipped (matching the cpl_vsi.go precedent).

func cplIsPowerOfTwo(i uint) (result int) {
	result = int(C.CPLIsPowerOfTwo(C.uint(i)))
	return
}

/* -------------------------------------------------------------------- */
/*      Terminal related                                                */
/* -------------------------------------------------------------------- */

// CPLIsInteractive operates on a raw C FILE*, part of the legacy group skipped
// in cpl_vsi.go; it is skipped here for the same reason.

// The C++ objects after CPL_C_END (CPLLocaleC, CPLThreadLocaleC,
// CPLConfigOptionSetter, cpl::down_cast, cpl::div_round_up) are extern "C++" and
// are skipped.
