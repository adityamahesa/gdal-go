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

func CPLVerifyConfiguration() {
	cplVerifyConfiguration()
}

func cplIsDebugEnabled() (result bool) {
	result = bool(C.CPLIsDebugEnabled())
	return
}

func CPLIsDebugEnabled() (result bool) {
	return cplIsDebugEnabled()
}

func cplGetConfigOption(key, dflt string) (result string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CPLGetConfigOption(cKey, cDflt))
	return
}

func CPLGetConfigOption(key, dflt string) (result string) {
	return cplGetConfigOption(key, dflt)
}

func cplGetThreadLocalConfigOption(key, dflt string) (result string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CPLGetThreadLocalConfigOption(cKey, cDflt))
	return
}

func CPLGetThreadLocalConfigOption(key, dflt string) (result string) {
	return cplGetThreadLocalConfigOption(key, dflt)
}

func cplGetGlobalConfigOption(key, dflt string) (result string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CPLGetGlobalConfigOption(cKey, cDflt))
	return
}

func CPLGetGlobalConfigOption(key, dflt string) (result string) {
	return cplGetGlobalConfigOption(key, dflt)
}

func cplSetConfigOption(key, value string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.CPLSetConfigOption(cKey, cValue)
}

func CPLSetConfigOption(key, value string) {
	cplSetConfigOption(key, value)
}

func cplSetThreadLocalConfigOption(key, value string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.CPLSetThreadLocalConfigOption(cKey, cValue)
}

func CPLSetThreadLocalConfigOption(key, value string) {
	cplSetThreadLocalConfigOption(key, value)
}

func cplDeclareKnownConfigOption(key, definition string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDefinition := C.CString(definition)
	defer C.free(unsafe.Pointer(cDefinition))
	C.CPLDeclareKnownConfigOption(cKey, cDefinition)
}

func CPLDeclareKnownConfigOption(key, definition string) {
	cplDeclareKnownConfigOption(key, definition)
}

func cplGetKnownConfigOptions() (result CSLConstList) {
	raw := C.CPLGetKnownConfigOptions()
	result = cslConstList(raw)
	return
}

// CPLGetKnownConfigOptions returns an owned list of known config options; the
// caller must Destroy it.
func CPLGetKnownConfigOptions() (result CSLConstList) {
	return cplGetKnownConfigOptions()
}

// CPLSetConfigOptionSubscriber (a C function-pointer typedef) and
// CPLSubscribeToSetConfigOption require a Go->C callback; both are deferred.

func cplUnsubscribeToSetConfigOption(subscriberId int) {
	C.CPLUnsubscribeToSetConfigOption(C.int(subscriberId))
}

func CPLUnsubscribeToSetConfigOption(subscriberId int) {
	cplUnsubscribeToSetConfigOption(subscriberId)
}

func cplFreeConfig() {
	C.CPLFreeConfig()
}

func CPLFreeConfig() {
	cplFreeConfig()
}

func cplGetConfigOptions() (result CSLConstList) {
	raw := C.CPLGetConfigOptions()
	result = cslConstList(raw)
	return
}

// CPLGetConfigOptions returns an owned list of the config options; the caller
// must Destroy it.
func CPLGetConfigOptions() (result CSLConstList) {
	return cplGetConfigOptions()
}

func cplSetConfigOptions(options CSLConstList) {
	C.CPLSetConfigOptions(options.cValue)
}

func CPLSetConfigOptions(options CSLConstList) {
	cplSetConfigOptions(options)
}

func cplGetThreadLocalConfigOptions() (result CSLConstList) {
	raw := C.CPLGetThreadLocalConfigOptions()
	result = cslConstList(raw)
	return
}

// CPLGetThreadLocalConfigOptions returns an owned list of the thread-local
// config options; the caller must Destroy it.
func CPLGetThreadLocalConfigOptions() (result CSLConstList) {
	return cplGetThreadLocalConfigOptions()
}

func cplSetThreadLocalConfigOptions(options CSLConstList) {
	C.CPLSetThreadLocalConfigOptions(options.cValue)
}

func CPLSetThreadLocalConfigOptions(options CSLConstList) {
	cplSetThreadLocalConfigOptions(options)
}

func cplLoadConfigOptionsFromFile(filename string, overrideEnvVars int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	C.CPLLoadConfigOptionsFromFile(cFilename, C.int(overrideEnvVars))
}

func CPLLoadConfigOptionsFromFile(filename string, overrideEnvVars int) {
	cplLoadConfigOptionsFromFile(filename, overrideEnvVars)
}

func cplLoadConfigOptionsFromPredefinedFiles() {
	C.CPLLoadConfigOptionsFromPredefinedFiles()
}

func CPLLoadConfigOptionsFromPredefinedFiles() {
	cplLoadConfigOptionsFromPredefinedFiles()
}

/* -------------------------------------------------------------------- */
/*      Safe malloc() API.                                              */
/* -------------------------------------------------------------------- */

func cplMalloc(size uint64) (result unsafe.Pointer) {
	result = C.CPLMalloc(C.size_t(size))
	return
}

func CPLMalloc(size uint64) (result unsafe.Pointer) {
	return cplMalloc(size)
}

func cplCalloc(count, size uint64) (result unsafe.Pointer) {
	result = C.CPLCalloc(C.size_t(count), C.size_t(size))
	return
}

func CPLCalloc(count, size uint64) (result unsafe.Pointer) {
	return cplCalloc(count, size)
}

func cplRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	result = C.CPLRealloc(ptr, C.size_t(size))
	return
}

func CPLRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	return cplRealloc(ptr, size)
}

func cplStrdup(s string) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CPLStrdup(cs)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func CPLStrdup(s string) (result string) {
	return cplStrdup(s)
}

func cplStrlwr(s string) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.CPLStrlwr(cs)
	result = C.GoString(cs)
	return
}

func CPLStrlwr(s string) (result string) {
	return cplStrlwr(s)
}

// Alias of VSIFree()
func CPLFree(ptr unsafe.Pointer) {
	vsiFree(ptr)
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

func (f VSILFile) ReadLineL() (result string) {
	return cplReadLineL(f)
}

func cplReadLine2L(file VSILFile, maxCars int, options CSLConstList) (result string) {
	result = C.GoString(C.CPLReadLine2L(file.cValue, C.int(maxCars), options.cValue))
	return
}

func (f VSILFile) ReadLine2L(maxCars int, options CSLConstList) (result string) {
	return cplReadLine2L(f, maxCars, options)
}

func cplReadLine3L(file VSILFile, maxCars int, options CSLConstList) (result string, bufLength int) {
	var cBufLength C.int
	result = C.GoString(C.CPLReadLine3L(file.cValue, C.int(maxCars), &cBufLength, options.cValue))
	bufLength = int(cBufLength)
	return
}

func (f VSILFile) ReadLine3L(maxCars int, options CSLConstList) (result string, bufLength int) {
	return cplReadLine3L(f, maxCars, options)
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

func CPLAtof(s string) (result float64) {
	return cplAtof(s)
}

func cplAtofDelim(s string, point byte) (result float64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = float64(C.CPLAtofDelim(cs, C.char(point)))
	return
}

func CPLAtofDelim(s string, point byte) (result float64) {
	return cplAtofDelim(s, point)
}

func cplStrtod(s string) (result float64, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float64(C.CPLStrtod(cs, &endptr))
	rest = C.GoString(endptr)
	return
}

func CPLStrtod(s string) (result float64, rest string) {
	return cplStrtod(s)
}

func cplStrtodM(s string) (result float64, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float64(C.CPLStrtodM(cs, &endptr))
	rest = C.GoString(endptr)
	return
}

func CPLStrtodM(s string) (result float64, rest string) {
	return cplStrtodM(s)
}

func cplStrtodDelim(s string, point byte) (result float64, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float64(C.CPLStrtodDelim(cs, &endptr, C.char(point)))
	rest = C.GoString(endptr)
	return
}

func CPLStrtodDelim(s string, point byte) (result float64, rest string) {
	return cplStrtodDelim(s, point)
}

func cplStrtof(s string) (result float32, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float32(C.CPLStrtof(cs, &endptr))
	rest = C.GoString(endptr)
	return
}

func CPLStrtof(s string) (result float32, rest string) {
	return cplStrtof(s)
}

func cplStrtofDelim(s string, point byte) (result float32, rest string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var endptr *C.char
	result = float32(C.CPLStrtofDelim(cs, &endptr, C.char(point)))
	rest = C.GoString(endptr)
	return
}

func CPLStrtofDelim(s string, point byte) (result float32, rest string) {
	return cplStrtofDelim(s, point)
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

func CPLAtofM(s string) (result float64) {
	return cplAtofM(s)
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

func CPLScanString(s string, maxLength, skipLeadingSpaces, stripQuotes int) (result string) {
	return cplScanString(s, maxLength, skipLeadingSpaces, stripQuotes)
}

func cplScanDouble(s string, length int) (result float64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = float64(C.CPLScanDouble(cs, C.int(length)))
	return
}

func CPLScanDouble(s string, length int) (result float64) {
	return cplScanDouble(s, length)
}

func cplScanLong(s string, length int) (result int64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = int64(C.CPLScanLong(cs, C.int(length)))
	return
}

func CPLScanLong(s string, length int) (result int64) {
	return cplScanLong(s, length)
}

func cplScanULong(s string, length int) (result uint64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = uint64(C.CPLScanULong(cs, C.int(length)))
	return
}

func CPLScanULong(s string, length int) (result uint64) {
	return cplScanULong(s, length)
}

func cplScanUIntBig(s string, length int) (result GUIntBig) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = GUIntBig(C.CPLScanUIntBig(cs, C.int(length)))
	return
}

func CPLScanUIntBig(s string, length int) (result GUIntBig) {
	return cplScanUIntBig(s, length)
}

func cplAtoGIntBig(s string) (result GIntBig) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = GIntBig(C.CPLAtoGIntBig(cs))
	return
}

func CPLAtoGIntBig(s string) (result GIntBig) {
	return cplAtoGIntBig(s)
}

func cplAtoGIntBigEx(s string, warn int) (result GIntBig, overflow int) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var cOverflow C.int
	result = GIntBig(C.CPLAtoGIntBigEx(cs, C.int(warn), &cOverflow))
	overflow = int(cOverflow)
	return
}

func CPLAtoGIntBigEx(s string, warn int) (result GIntBig, overflow int) {
	return cplAtoGIntBigEx(s, warn)
}

func cplScanPointer(s string, length int) (result unsafe.Pointer) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = C.CPLScanPointer(cs, C.int(length))
	return
}

func CPLScanPointer(s string, length int) (result unsafe.Pointer) {
	return cplScanPointer(s, length)
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

func CPLPrintString(s string, maxLen int) (result string) {
	return cplPrintString(s, maxLen)
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

func CPLPrintStringFill(s string, maxLen int) (result string) {
	return cplPrintStringFill(s, maxLen)
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

func CPLPrintInt32(value GInt32, maxLen int) (result string) {
	return cplPrintInt32(value, maxLen)
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

func CPLPrintUIntBig(value GUIntBig, maxLen int) (result string) {
	return cplPrintUIntBig(value, maxLen)
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

func CPLPrintDouble(format string, value float64, locale string) (result string) {
	return cplPrintDouble(format, value, locale)
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

func CPLPrintPointer(ptr unsafe.Pointer, maxLen int) (result string) {
	return cplPrintPointer(ptr, maxLen)
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

func CPLGetSymbol(library, symbolName string) (result unsafe.Pointer) {
	return cplGetSymbol(library, symbolName)
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

func CPLGetExecPath(maxLength int) (result string, ok bool) {
	return cplGetExecPath(maxLength)
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

func CPLGetPath(path string) (result string) {
	return cplGetPath(path)
}

func cplGetDirname(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetDirname(cPath))
	return
}

func CPLGetDirname(path string) (result string) {
	return cplGetDirname(path)
}

func cplGetBasename(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetBasename(cPath))
	return
}

func CPLGetBasename(path string) (result string) {
	return cplGetBasename(path)
}

func cplGetExtension(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetExtension(cPath))
	return
}

func CPLGetExtension(path string) (result string) {
	return cplGetExtension(path)
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

func CPLFormFilename(path, basename, extension string) (result string) {
	return cplFormFilename(path, basename, extension)
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

func CPLFormCIFilename(path, basename, extension string) (result string) {
	return cplFormCIFilename(path, basename, extension)
}

func cplResetExtension(path, extension string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cExtension := C.CString(extension)
	defer C.free(unsafe.Pointer(cExtension))
	result = C.GoString(C.CPLResetExtension(cPath, cExtension))
	return
}

func CPLResetExtension(path, extension string) (result string) {
	return cplResetExtension(path, extension)
}

func cplProjectRelativeFilename(projectDir, secondaryFilename string) (result string) {
	cProjectDir := C.CString(projectDir)
	defer C.free(unsafe.Pointer(cProjectDir))
	cSecondaryFilename := C.CString(secondaryFilename)
	defer C.free(unsafe.Pointer(cSecondaryFilename))
	result = C.GoString(C.CPLProjectRelativeFilename(cProjectDir, cSecondaryFilename))
	return
}

func CPLProjectRelativeFilename(projectDir, secondaryFilename string) (result string) {
	return cplProjectRelativeFilename(projectDir, secondaryFilename)
}

func cplCleanTrailingSlash(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLCleanTrailingSlash(cPath))
	return
}

func CPLCleanTrailingSlash(path string) (result string) {
	return cplCleanTrailingSlash(path)
}

func cplGenerateTempFilename(stem string) (result string) {
	cStem := C.CString(stem)
	defer C.free(unsafe.Pointer(cStem))
	result = C.GoString(C.CPLGenerateTempFilename(cStem))
	return
}

func CPLGenerateTempFilename(stem string) (result string) {
	return cplGenerateTempFilename(stem)
}

func cplExpandTilde(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.CPLExpandTilde(cFilename))
	return
}

func CPLExpandTilde(filename string) (result string) {
	return cplExpandTilde(filename)
}

func cplLaunderForFilename(name, outputPath string) (result string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cOutputPath := C.CString(outputPath)
	defer C.free(unsafe.Pointer(cOutputPath))
	result = C.GoString(C.CPLLaunderForFilename(cName, cOutputPath))
	return
}

func CPLLaunderForFilename(name, outputPath string) (result string) {
	return cplLaunderForFilename(name, outputPath)
}

func cplGetCurrentDir() (result string) {
	raw := C.CPLGetCurrentDir()
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func CPLGetCurrentDir() (result string) {
	return cplGetCurrentDir()
}

func cplGetFilename(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.CPLGetFilename(cPath))
	return
}

func CPLGetFilename(path string) (result string) {
	return cplGetFilename(path)
}

func cplIsFilenameRelative(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.CPLIsFilenameRelative(cFilename))
	return
}

func CPLIsFilenameRelative(filename string) (result int) {
	return cplIsFilenameRelative(filename)
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

func CPLExtractRelativePath(baseDir, target string) (result string, gotRelative int) {
	return cplExtractRelativePath(baseDir, target)
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

// CPLCorrespondingPaths returns an owned list of renamed paths; the caller must
// Destroy it.
func CPLCorrespondingPaths(oldFilename, newFilename string, fileList CSLConstList) (result CSLConstList) {
	return cplCorrespondingPaths(oldFilename, newFilename, fileList)
}

func cplCheckForFile(filename string, siblings CSLConstList) (result int, corrected string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.CPLCheckForFile(cFilename, siblings.cValue))
	corrected = C.GoString(cFilename)
	return
}

func CPLCheckForFile(filename string, siblings CSLConstList) (result int, corrected string) {
	return cplCheckForFile(filename, siblings)
}

func cplGetHomeDir() (result string) {
	result = C.GoString(C.CPLGetHomeDir())
	return
}

func CPLGetHomeDir() (result string) {
	return cplGetHomeDir()
}

// The extern "C++" ...Safe variants (CPLGetPathSafe, CPLGetDirnameSafe, etc.)
// return std::string and are C++-only covers; they are skipped.

func cplHasPathTraversal(filename string) (result bool) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = bool(C.CPLHasPathTraversal(cFilename))
	return
}

func CPLHasPathTraversal(filename string) (result bool) {
	return cplHasPathTraversal(filename)
}

func cplHasUnbalancedPathTraversal(filename string) (result bool) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = bool(C.CPLHasUnbalancedPathTraversal(cFilename))
	return
}

func CPLHasUnbalancedPathTraversal(filename string) (result bool) {
	return cplHasUnbalancedPathTraversal(filename)
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

func CPLFindFile(class, basename string) (result string) {
	return cplFindFile(class, basename)
}

func cplDefaultFindFile(class, basename string) (result string) {
	cClass := C.CString(class)
	defer C.free(unsafe.Pointer(cClass))
	cBasename := C.CString(basename)
	defer C.free(unsafe.Pointer(cBasename))
	result = C.GoString(C.CPLDefaultFindFile(cClass, cBasename))
	return
}

func CPLDefaultFindFile(class, basename string) (result string) {
	return cplDefaultFindFile(class, basename)
}

func cplPushFinderLocation(location string) {
	cLocation := C.CString(location)
	defer C.free(unsafe.Pointer(cLocation))
	C.CPLPushFinderLocation(cLocation)
}

func CPLPushFinderLocation(location string) {
	cplPushFinderLocation(location)
}

func cplPopFinderLocation() {
	C.CPLPopFinderLocation()
}

func CPLPopFinderLocation() {
	cplPopFinderLocation()
}

func cplFinderClean() {
	C.CPLFinderClean()
}

func CPLFinderClean() {
	cplFinderClean()
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

func CPLDMSToDec(is string) (result float64) {
	return cplDMSToDec(is)
}

func cplDecToDMS(angle float64, axis string, precision int) (result string) {
	cAxis := C.CString(axis)
	defer C.free(unsafe.Pointer(cAxis))
	result = C.GoString(C.CPLDecToDMS(C.double(angle), cAxis, C.int(precision)))
	return
}

func CPLDecToDMS(angle float64, axis string, precision int) (result string) {
	return cplDecToDMS(angle, axis, precision)
}

func cplPackedDMSToDec(packed float64) (result float64) {
	result = float64(C.CPLPackedDMSToDec(C.double(packed)))
	return
}

func CPLPackedDMSToDec(packed float64) (result float64) {
	return cplPackedDMSToDec(packed)
}

func cplDecToPackedDMS(dec float64) (result float64) {
	result = float64(C.CPLDecToPackedDMS(C.double(dec)))
	return
}

func CPLDecToPackedDMS(dec float64) (result float64) {
	return cplDecToPackedDMS(dec)
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

func CPLStringToComplex(s string) (real, imag float64, err error) {
	var ret CPLErr
	real, imag, ret = cplStringToComplex(s)
	err = cplErr(ret)
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

func CPLUnlinkTree(path string) (result int) {
	return cplUnlinkTree(path)
}

func cplCopyFile(newPath, oldPath string) (result int) {
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	result = int(C.CPLCopyFile(cNew, cOld))
	return
}

func CPLCopyFile(newPath, oldPath string) (result int) {
	return cplCopyFile(newPath, oldPath)
}

func cplCopyTree(newPath, oldPath string) (result int) {
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	result = int(C.CPLCopyTree(cNew, cOld))
	return
}

func CPLCopyTree(newPath, oldPath string) (result int) {
	return cplCopyTree(newPath, oldPath)
}

func cplMoveFile(newPath, oldPath string) (result int) {
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	result = int(C.CPLMoveFile(cNew, cOld))
	return
}

func CPLMoveFile(newPath, oldPath string) (result int) {
	return cplMoveFile(newPath, oldPath)
}

func cplSymlink(oldPath, newPath string, options CSLConstList) (result int) {
	cOld := C.CString(oldPath)
	defer C.free(unsafe.Pointer(cOld))
	cNew := C.CString(newPath)
	defer C.free(unsafe.Pointer(cNew))
	result = int(C.CPLSymlink(cOld, cNew, options.cValue))
	return
}

func CPLSymlink(oldPath, newPath string, options CSLConstList) (result int) {
	return cplSymlink(oldPath, newPath, options)
}

func cplGetRemainingFileDescriptorCount() (result int) {
	result = int(C.CPLGetRemainingFileDescriptorCount())
	return
}

func CPLGetRemainingFileDescriptorCount() (result int) {
	return cplGetRemainingFileDescriptorCount()
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

func CPLLockFileEx(lockFileName string, options CSLConstList) (handle CPLLockFileHandle, status CPLLockFileStatus) {
	return cplLockFileEx(lockFileName, options)
}

func cplUnlockFileEx(handle CPLLockFileHandle) {
	C.CPLUnlockFileEx(handle.cValue)
}

func (h CPLLockFileHandle) UnlockFileEx() {
	cplUnlockFileEx(h)
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

func CPLCreateZip(zipFilename string, options CSLConstList) (result unsafe.Pointer, err error) {
	result = cplCreateZip(zipFilename, options)
	if result == nil {
		err = lastError()
	}
	return
}

func cplCreateFileInZip(zip unsafe.Pointer, filename string, options CSLConstList) (result CPLErr) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = CPLErr(C.CPLCreateFileInZip(zip, cFilename, options.cValue))
	return
}

func CPLCreateFileInZip(zip unsafe.Pointer, filename string, options CSLConstList) (err error) {
	return cplErr(cplCreateFileInZip(zip, filename, options))
}

func cplWriteFileInZip(zip unsafe.Pointer, buffer []byte) (result CPLErr) {
	result = CPLErr(C.CPLWriteFileInZip(zip, cBytes(buffer), C.int(len(buffer))))
	return
}

func CPLWriteFileInZip(zip unsafe.Pointer, buffer []byte) (err error) {
	return cplErr(cplWriteFileInZip(zip, buffer))
}

func cplCloseFileInZip(zip unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.CPLCloseFileInZip(zip))
	return
}

func CPLCloseFileInZip(zip unsafe.Pointer) (err error) {
	return cplErr(cplCloseFileInZip(zip))
}

func cplAddFileInZip(zip unsafe.Pointer, archiveFilename, inputFilename string, fpInput VSILFile, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CPLErr) {
	cArchiveFilename := C.CString(archiveFilename)
	defer C.free(unsafe.Pointer(cArchiveFilename))
	cInputFilename := C.CString(inputFilename)
	defer C.free(unsafe.Pointer(cInputFilename))
	result = CPLErr(C.CPLAddFileInZip(zip, cArchiveFilename, cInputFilename, fpInput.cValue, options.cValue, progress.cValue, progressData))
	return
}

func CPLAddFileInZip(zip unsafe.Pointer, archiveFilename, inputFilename string, fpInput VSILFile, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	return cplErr(cplAddFileInZip(zip, archiveFilename, inputFilename, fpInput, options, progress, progressData))
}

func cplCloseZip(zip unsafe.Pointer) (result CPLErr) {
	result = CPLErr(C.CPLCloseZip(zip))
	return
}

func CPLCloseZip(zip unsafe.Pointer) (err error) {
	return cplErr(cplCloseZip(zip))
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

func CPLZLibDeflate(input []byte, level int) (result []byte) {
	return cplZLibDeflate(input, level)
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

func CPLZLibInflate(input []byte) (result []byte) {
	return cplZLibInflate(input)
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

func CPLZLibInflateEx(input []byte, allowResizeOutptr bool) (result []byte) {
	return cplZLibInflateEx(input, allowResizeOutptr)
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

func CPLValidateXML(xmlFilename, xsdFilename string, options CSLConstList) (result int) {
	return cplValidateXML(xmlFilename, xsdFilename, options)
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

func CPLIsPowerOfTwo(i uint) (result int) {
	return cplIsPowerOfTwo(i)
}

/* -------------------------------------------------------------------- */
/*      Terminal related                                                */
/* -------------------------------------------------------------------- */

// CPLIsInteractive operates on a raw C FILE*, part of the legacy group skipped
// in cpl_vsi.go; it is skipped here for the same reason.

// The C++ objects after CPL_C_END (CPLLocaleC, CPLThreadLocaleC,
// CPLConfigOptionSetter, cpl::down_cast, cpl::div_round_up) are extern "C++" and
// are skipped.
