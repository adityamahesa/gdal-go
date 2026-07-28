package gdal

import "unsafe"

func CPLVerifyConfiguration() {
	cplVerifyConfiguration()
}

func CPLIsDebugEnabled() (result bool) {
	return cplIsDebugEnabled()
}

func CPLGetConfigOption(key, dflt string) (result string) {
	return cplGetConfigOption(key, dflt)
}

func CPLGetThreadLocalConfigOption(key, dflt string) (result string) {
	return cplGetThreadLocalConfigOption(key, dflt)
}

func CPLGetGlobalConfigOption(key, dflt string) (result string) {
	return cplGetGlobalConfigOption(key, dflt)
}

func CPLSetConfigOption(key, value string) {
	cplSetConfigOption(key, value)
}

func CPLSetThreadLocalConfigOption(key, value string) {
	cplSetThreadLocalConfigOption(key, value)
}

func CPLDeclareKnownConfigOption(key, definition string) {
	cplDeclareKnownConfigOption(key, definition)
}

// CPLGetKnownConfigOptions returns an owned list of known config options; the
// caller must Destroy it.
func CPLGetKnownConfigOptions() (result CSLConstList) {
	return cplGetKnownConfigOptions()
}

func CPLUnsubscribeToSetConfigOption(subscriberId int) {
	cplUnsubscribeToSetConfigOption(subscriberId)
}

func CPLFreeConfig() {
	cplFreeConfig()
}

// CPLGetConfigOptions returns an owned list of the config options; the caller
// must Destroy it.
func CPLGetConfigOptions() (result CSLConstList) {
	return cplGetConfigOptions()
}

func CPLSetConfigOptions(options CSLConstList) {
	cplSetConfigOptions(options)
}

// CPLGetThreadLocalConfigOptions returns an owned list of the thread-local
// config options; the caller must Destroy it.
func CPLGetThreadLocalConfigOptions() (result CSLConstList) {
	return cplGetThreadLocalConfigOptions()
}

func CPLSetThreadLocalConfigOptions(options CSLConstList) {
	cplSetThreadLocalConfigOptions(options)
}

func CPLLoadConfigOptionsFromFile(filename string, overrideEnvVars int) {
	cplLoadConfigOptionsFromFile(filename, overrideEnvVars)
}

func CPLLoadConfigOptionsFromPredefinedFiles() {
	cplLoadConfigOptionsFromPredefinedFiles()
}

func CPLMalloc(size uint64) (result unsafe.Pointer) {
	return cplMalloc(size)
}

func CPLCalloc(count, size uint64) (result unsafe.Pointer) {
	return cplCalloc(count, size)
}

func CPLRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	return cplRealloc(ptr, size)
}

func CPLStrdup(s string) (result string) {
	return cplStrdup(s)
}

func CPLStrlwr(s string) (result string) {
	return cplStrlwr(s)
}

// Alias of VSIFree()
func CPLFree(ptr unsafe.Pointer) {
	vsiFree(ptr)
}

func (f VSILFile) ReadLineL() (result string) {
	return cplReadLineL(f)
}

func (f VSILFile) ReadLine2L(maxCars int, options CSLConstList) (result string) {
	return cplReadLine2L(f, maxCars, options)
}

func (f VSILFile) ReadLine3L(maxCars int, options CSLConstList) (result string, bufLength int) {
	return cplReadLine3L(f, maxCars, options)
}

func CPLAtof(s string) (result float64) {
	return cplAtof(s)
}

func CPLAtofDelim(s string, point byte) (result float64) {
	return cplAtofDelim(s, point)
}

func CPLStrtod(s string) (result float64, rest string) {
	return cplStrtod(s)
}

func CPLStrtodM(s string) (result float64, rest string) {
	return cplStrtodM(s)
}

func CPLStrtodDelim(s string, point byte) (result float64, rest string) {
	return cplStrtodDelim(s, point)
}

func CPLStrtof(s string) (result float32, rest string) {
	return cplStrtof(s)
}

func CPLStrtofDelim(s string, point byte) (result float32, rest string) {
	return cplStrtofDelim(s, point)
}

func CPLAtofM(s string) (result float64) {
	return cplAtofM(s)
}

func CPLScanString(s string, maxLength, skipLeadingSpaces, stripQuotes int) (result string) {
	return cplScanString(s, maxLength, skipLeadingSpaces, stripQuotes)
}

func CPLScanDouble(s string, length int) (result float64) {
	return cplScanDouble(s, length)
}

func CPLScanLong(s string, length int) (result int64) {
	return cplScanLong(s, length)
}

func CPLScanULong(s string, length int) (result uint64) {
	return cplScanULong(s, length)
}

func CPLScanUIntBig(s string, length int) (result GUIntBig) {
	return cplScanUIntBig(s, length)
}

func CPLAtoGIntBig(s string) (result GIntBig) {
	return cplAtoGIntBig(s)
}

func CPLAtoGIntBigEx(s string, warn int) (result GIntBig, overflow int) {
	return cplAtoGIntBigEx(s, warn)
}

func CPLScanPointer(s string, length int) (result unsafe.Pointer) {
	return cplScanPointer(s, length)
}

func CPLPrintString(s string, maxLen int) (result string) {
	return cplPrintString(s, maxLen)
}

func CPLPrintStringFill(s string, maxLen int) (result string) {
	return cplPrintStringFill(s, maxLen)
}

func CPLPrintInt32(value GInt32, maxLen int) (result string) {
	return cplPrintInt32(value, maxLen)
}

func CPLPrintUIntBig(value GUIntBig, maxLen int) (result string) {
	return cplPrintUIntBig(value, maxLen)
}

func CPLPrintDouble(format string, value float64, locale string) (result string) {
	return cplPrintDouble(format, value, locale)
}

func CPLPrintPointer(ptr unsafe.Pointer, maxLen int) (result string) {
	return cplPrintPointer(ptr, maxLen)
}

func CPLGetSymbol(library, symbolName string) (result unsafe.Pointer) {
	return cplGetSymbol(library, symbolName)
}

func CPLGetExecPath(maxLength int) (result string, ok bool) {
	return cplGetExecPath(maxLength)
}

func CPLGetPath(path string) (result string) {
	return cplGetPath(path)
}

func CPLGetDirname(path string) (result string) {
	return cplGetDirname(path)
}

func CPLGetBasename(path string) (result string) {
	return cplGetBasename(path)
}

func CPLGetExtension(path string) (result string) {
	return cplGetExtension(path)
}

func CPLFormFilename(path, basename, extension string) (result string) {
	return cplFormFilename(path, basename, extension)
}

func CPLFormCIFilename(path, basename, extension string) (result string) {
	return cplFormCIFilename(path, basename, extension)
}

func CPLResetExtension(path, extension string) (result string) {
	return cplResetExtension(path, extension)
}

func CPLProjectRelativeFilename(projectDir, secondaryFilename string) (result string) {
	return cplProjectRelativeFilename(projectDir, secondaryFilename)
}

func CPLCleanTrailingSlash(path string) (result string) {
	return cplCleanTrailingSlash(path)
}

func CPLGenerateTempFilename(stem string) (result string) {
	return cplGenerateTempFilename(stem)
}

func CPLExpandTilde(filename string) (result string) {
	return cplExpandTilde(filename)
}

func CPLLaunderForFilename(name, outputPath string) (result string) {
	return cplLaunderForFilename(name, outputPath)
}

func CPLGetCurrentDir() (result string) {
	return cplGetCurrentDir()
}

func CPLGetFilename(path string) (result string) {
	return cplGetFilename(path)
}

func CPLIsFilenameRelative(filename string) (result int) {
	return cplIsFilenameRelative(filename)
}

func CPLExtractRelativePath(baseDir, target string) (result string, gotRelative int) {
	return cplExtractRelativePath(baseDir, target)
}

// CPLCorrespondingPaths returns an owned list of renamed paths; the caller must
// Destroy it.
func CPLCorrespondingPaths(oldFilename, newFilename string, fileList CSLConstList) (result CSLConstList) {
	return cplCorrespondingPaths(oldFilename, newFilename, fileList)
}

func CPLCheckForFile(filename string, siblings CSLConstList) (result int, corrected string) {
	return cplCheckForFile(filename, siblings)
}

func CPLGetHomeDir() (result string) {
	return cplGetHomeDir()
}

func CPLHasPathTraversal(filename string) (result bool) {
	return cplHasPathTraversal(filename)
}

func CPLHasUnbalancedPathTraversal(filename string) (result bool) {
	return cplHasUnbalancedPathTraversal(filename)
}

func CPLFindFile(class, basename string) (result string) {
	return cplFindFile(class, basename)
}

func CPLDefaultFindFile(class, basename string) (result string) {
	return cplDefaultFindFile(class, basename)
}

func CPLPushFinderLocation(location string) {
	cplPushFinderLocation(location)
}

func CPLPopFinderLocation() {
	cplPopFinderLocation()
}

func CPLFinderClean() {
	cplFinderClean()
}

func CPLDMSToDec(is string) (result float64) {
	return cplDMSToDec(is)
}

func CPLDecToDMS(angle float64, axis string, precision int) (result string) {
	return cplDecToDMS(angle, axis, precision)
}

func CPLPackedDMSToDec(packed float64) (result float64) {
	return cplPackedDMSToDec(packed)
}

func CPLDecToPackedDMS(dec float64) (result float64) {
	return cplDecToPackedDMS(dec)
}

func CPLStringToComplex(s string) (real, imag float64, err error) {
	scope := errScope()
	defer scope()
	var ret CPLErr
	real, imag, ret = cplStringToComplex(s)
	err = cplErr(ret)
	return
}

func CPLUnlinkTree(path string) (result int) {
	return cplUnlinkTree(path)
}

func CPLCopyFile(newPath, oldPath string) (result int) {
	return cplCopyFile(newPath, oldPath)
}

func CPLCopyTree(newPath, oldPath string) (result int) {
	return cplCopyTree(newPath, oldPath)
}

func CPLMoveFile(newPath, oldPath string) (result int) {
	return cplMoveFile(newPath, oldPath)
}

func CPLSymlink(oldPath, newPath string, options CSLConstList) (result int) {
	return cplSymlink(oldPath, newPath, options)
}

func CPLGetRemainingFileDescriptorCount() (result int) {
	return cplGetRemainingFileDescriptorCount()
}

func CPLLockFileEx(lockFileName string, options CSLConstList) (handle CPLLockFileHandle, status CPLLockFileStatus) {
	return cplLockFileEx(lockFileName, options)
}

func (h CPLLockFileHandle) UnlockFileEx() {
	cplUnlockFileEx(h)
}

func CPLCreateZip(zipFilename string, options CSLConstList) (result unsafe.Pointer, err error) {
	scope := errScope()
	defer scope()
	result = cplCreateZip(zipFilename, options)
	if result == nil {
		err = lastError()
	}
	return
}

func CPLCreateFileInZip(zip unsafe.Pointer, filename string, options CSLConstList) (err error) {
	scope := errScope()
	defer scope()
	return cplErr(cplCreateFileInZip(zip, filename, options))
}

func CPLWriteFileInZip(zip unsafe.Pointer, buffer []byte) (err error) {
	scope := errScope()
	defer scope()
	return cplErr(cplWriteFileInZip(zip, buffer))
}

func CPLCloseFileInZip(zip unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	return cplErr(cplCloseFileInZip(zip))
}

func CPLAddFileInZip(zip unsafe.Pointer, archiveFilename, inputFilename string, fpInput VSILFile, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	return cplErr(cplAddFileInZip(zip, archiveFilename, inputFilename, fpInput, options, progress, progressData))
}

func CPLCloseZip(zip unsafe.Pointer) (err error) {
	scope := errScope()
	defer scope()
	return cplErr(cplCloseZip(zip))
}

func CPLZLibDeflate(input []byte, level int) (result []byte) {
	return cplZLibDeflate(input, level)
}

func CPLZLibInflate(input []byte) (result []byte) {
	return cplZLibInflate(input)
}

func CPLZLibInflateEx(input []byte, allowResizeOutptr bool) (result []byte) {
	return cplZLibInflateEx(input, allowResizeOutptr)
}

func CPLValidateXML(xmlFilename, xsdFilename string, options CSLConstList) (result int) {
	return cplValidateXML(xmlFilename, xsdFilename, options)
}

func CPLIsPowerOfTwo(i uint) (result int) {
	return cplIsPowerOfTwo(i)
}
