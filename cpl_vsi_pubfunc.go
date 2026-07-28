package gdal

import "unsafe"

func VSIFOpenL(filename, access string) (result VSILFile, err error) {
	result = vsiFOpenL(filename, access)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func VSIFOpenExL(filename, access string, setError int) (result VSILFile, err error) {
	result = vsiFOpenExL(filename, access, setError)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func VSIFOpenEx2L(filename, access string, setError int, options CSLConstList) (result VSILFile, err error) {
	result = vsiFOpenEx2L(filename, access, setError, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (f VSILFile) CloseL() (result int) {
	return vsiFCloseL(f)
}

func (f VSILFile) SeekL(offset VSILOffset, whence int) (result int) {
	return vsiFSeekL(f, offset, whence)
}

func (f VSILFile) TellL() (result VSILOffset) {
	return vsiFTellL(f)
}

func (f VSILFile) RewindL() {
	vsiRewindL(f)
}

// ReadL reads up to len(buffer) bytes into buffer and returns the number of
// bytes read (VSIFReadL called with a block size of 1).
func (f VSILFile) ReadL(buffer []byte) (result uint64) {
	return vsiFReadL(buffer, 1, uint64(len(buffer)), f)
}

// ReadMultiRangeL reads len(offsets) ranges, the i-th of sizes[i] bytes starting
// at offsets[i]. Data is staged through C-allocated buffers, so the returned
// slices are independent Go copies.
func (f VSILFile) ReadMultiRangeL(offsets []VSILOffset, sizes []int) (result [][]byte, ret int) {
	n := len(offsets)
	if n == 0 {
		return
	}
	ppData := make([]unsafe.Pointer, n)
	goSizes := make([]uint64, n)
	for i := 0; i < n; i++ {
		ppData[i] = vsiMalloc(uint64(sizes[i]))
		goSizes[i] = uint64(sizes[i])
	}
	ret = vsiFReadMultiRangeL(n, ppData, offsets, goSizes, f)
	result = make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = goBytes(ppData[i], sizes[i])
		vsiFree(ppData[i])
	}
	return
}

// WriteL writes len(buffer) bytes from buffer and returns the number of bytes
// written (VSIFWriteL called with a block size of 1).
func (f VSILFile) WriteL(buffer []byte) (result uint64) {
	return vsiFWriteL(buffer, 1, uint64(len(buffer)), f)
}

func (f VSILFile) ClearErrL() {
	vsiFClearErrL(f)
}

func (f VSILFile) ErrorL() (result int) {
	return vsiFErrorL(f)
}

func (f VSILFile) EofL() (result int) {
	return vsiFEofL(f)
}

func (f VSILFile) TruncateL(newSize VSILOffset) (result int) {
	return vsiFTruncateL(f, newSize)
}

func (f VSILFile) FlushL() (result int) {
	return vsiFFlushL(f)
}

func (f VSILFile) PutcL(c int) (result int) {
	return vsiFPutcL(c, f)
}

func (f VSILFile) GetRangeStatusL(start, length VSILOffset) (result VSIRangeStatus) {
	return vsiFGetRangeStatusL(f, start, length)
}

func VSIIngestFile(file VSILFile, filename string, maxSize int64) (result []byte, err error) {
	var ret int
	result, ret = vsiIngestFile(file, filename, maxSize)
	if ret == 0 {
		err = lastError()
	}
	return
}

func VSIOverwriteFile(target VSILFile, source string) (result int) {
	return vsiOverwriteFile(target, source)
}

// Size returns the st_size field (file size in bytes).
func (s VSIStatBufL) Size() int64 {
	return vsiStatbufSize(&s)
}

// Mode returns the st_mode field (see the VSI_IS* macros in the C header).
func (s VSIStatBufL) Mode() int {
	return vsiStatbufMode(&s)
}

func VSIStatL(filename string) (result VSIStatBufL, err error) {
	var ret int
	result, ret = vsiStatL(filename)
	if ret != 0 {
		err = lastError()
	}
	return
}

func VSIStatExL(filename string, flags int) (result VSIStatBufL, err error) {
	var ret int
	result, ret = vsiStatExL(filename, flags)
	if ret != 0 {
		err = lastError()
	}
	return
}

func VSIIsCaseSensitiveFS(filename string) (result int) {
	return vsiIsCaseSensitiveFS(filename)
}

func VSISupportsSparseFiles(path string) (result int) {
	return vsiSupportsSparseFiles(path)
}

func VSIIsLocal(path string) (result bool) {
	return vsiIsLocal(path)
}

func VSIGetCanonicalFilename(path string) (result string) {
	return vsiGetCanonicalFilename(path)
}

func VSISupportsSequentialWrite(path string, allowLocalTempFile bool) (result bool) {
	return vsiSupportsSequentialWrite(path, allowLocalTempFile)
}

func VSISupportsRandomWrite(path string, allowLocalTempFile bool) (result bool) {
	return vsiSupportsRandomWrite(path, allowLocalTempFile)
}

func VSIHasOptimizedReadMultiRange(path string) (result int) {
	return vsiHasOptimizedReadMultiRange(path)
}

func VSIGetActualURL(filename string) (result string) {
	return vsiGetActualURL(filename)
}

func VSIGetSignedURL(filename string, options CSLConstList) (result string) {
	return vsiGetSignedURL(filename, options)
}

func VSIGetFileSystemOptions(filename string) (result string) {
	return vsiGetFileSystemOptions(filename)
}

func VSIGetFileSystemsPrefixes() (result CSLConstList) {
	return vsiGetFileSystemsPrefixes()
}

func (f VSILFile) GetNativeFileDescriptorL() (result unsafe.Pointer) {
	return vsiFGetNativeFileDescriptorL(f)
}

func VSIGetFileMetadata(filename, domain string, options CSLConstList) (result CSLConstList) {
	return vsiGetFileMetadata(filename, domain, options)
}

func VSISetFileMetadata(filename string, metadata CSLConstList, domain string, options CSLConstList) (result int) {
	return vsiSetFileMetadata(filename, metadata, domain, options)
}

func VSISetPathSpecificOption(prefix, key, value string) {
	vsiSetPathSpecificOption(prefix, key, value)
}

func VSIClearPathSpecificOptions(prefix string) {
	vsiClearPathSpecificOptions(prefix)
}

func VSIGetPathSpecificOption(path, key, dflt string) (result string) {
	return vsiGetPathSpecificOption(path, key, dflt)
}

// Deprecated: use VSISetPathSpecificOption.
func VSISetCredential(prefix, key, value string) {
	vsiSetCredential(prefix, key, value)
}

// Deprecated: use VSIClearPathSpecificOptions.
func VSIClearCredentials(prefix string) {
	vsiClearCredentials(prefix)
}

// Deprecated: use VSIGetPathSpecificOption.
func VSIGetCredential(path, key, dflt string) (result string) {
	return vsiGetCredential(path, key, dflt)
}

func VSICalloc(count, size uint64) (result unsafe.Pointer) {
	return vsiCalloc(count, size)
}

func VSIMalloc(size uint64) (result unsafe.Pointer) {
	return vsiMalloc(size)
}

func VSIFree(ptr unsafe.Pointer) {
	vsiFree(ptr)
}

func VSIRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	return vsiRealloc(ptr, size)
}

func VSIStrdup(s string) (result string) {
	return vsiStrdup(s)
}

func VSIMallocAligned(alignment, size uint64) (result unsafe.Pointer) {
	return vsiMallocAligned(alignment, size)
}

func VSIMallocAlignedAuto(size uint64) (result unsafe.Pointer) {
	return vsiMallocAlignedAuto(size)
}

func VSIFreeAligned(ptr unsafe.Pointer) {
	vsiFreeAligned(ptr)
}

func VSIMallocAlignedAutoVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMallocAlignedAutoVerbose(size, file, line)
}

func VSIMalloc2(size1, size2 uint64) (result unsafe.Pointer) {
	return vsiMalloc2(size1, size2)
}

func VSIMalloc3(size1, size2, size3 uint64) (result unsafe.Pointer) {
	return vsiMalloc3(size1, size2, size3)
}

func VSIMallocVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMallocVerbose(size, file, line)
}

func VSIMalloc2Verbose(size1, size2 uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMalloc2Verbose(size1, size2, file, line)
}

func VSIMalloc3Verbose(size1, size2, size3 uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMalloc3Verbose(size1, size2, size3, file, line)
}

func VSICallocVerbose(count, size uint64, file string, line int) (result unsafe.Pointer) {
	return vsiCallocVerbose(count, size, file, line)
}

func VSIReallocVerbose(oldPtr unsafe.Pointer, newSize uint64, file string, line int) (result unsafe.Pointer) {
	return vsiReallocVerbose(oldPtr, newSize, file, line)
}

func VSIStrdupVerbose(s string, file string, line int) (result string) {
	return vsiStrdupVerbose(s, file, line)
}

func CPLGetPhysicalRAM() (result GIntBig) {
	return cplGetPhysicalRAM()
}

func CPLGetUsablePhysicalRAM() (result GIntBig) {
	return cplGetUsablePhysicalRAM()
}

func VSIReadDir(path string) (result CSLConstList) {
	return vsiReadDir(path)
}

// Alias of VSIReadDir()
func CPLReadDir(path string) (result CSLConstList) {
	return vsiReadDir(path)
}

func VSIReadDirRecursive(path string) (result CSLConstList) {
	return vsiReadDirRecursive(path)
}

func VSIReadDirEx(path string, maxFiles int) (result CSLConstList) {
	return vsiReadDirEx(path, maxFiles)
}

func VSISiblingFiles(path string) (result CSLConstList) {
	return vsiSiblingFiles(path)
}

func VSIGlob(pattern string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CSLConstList) {
	return vsiGlob(pattern, options, progress, progressData)
}

func VSIGetDirectorySeparator(path string) (result string) {
	return vsiGetDirectorySeparator(path)
}

func VSIOpenDir(path string, recurseDepth int, options CSLConstList) (result VSIDir, err error) {
	result = vsiOpenDir(path, recurseDepth, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// Name returns the entry filename.
func (e VSIDirEntry) Name() string {
	return vsiDirEntryName(e)
}

// Mode returns the file mode. See the VSI_IS* macros in the C header.
func (e VSIDirEntry) Mode() int {
	return int(e.cValue.nMode)
}

// Size returns the file size.
func (e VSIDirEntry) Size() VSILOffset {
	return VSILOffset(e.cValue.nSize)
}

// MTime returns the last modification time (seconds since 1970/01/01).
func (e VSIDirEntry) MTime() int64 {
	return int64(e.cValue.nMTime)
}

// ModeKnown reports whether Mode() is meaningful.
func (e VSIDirEntry) ModeKnown() bool {
	return e.cValue.bModeKnown != 0
}

// SizeKnown reports whether Size() is meaningful.
func (e VSIDirEntry) SizeKnown() bool {
	return e.cValue.bSizeKnown != 0
}

// MTimeKnown reports whether MTime() is meaningful.
func (e VSIDirEntry) MTimeKnown() bool {
	return e.cValue.bMTimeKnown != 0
}

// Extra returns the NULL-terminated list of extra properties. The list is
// borrowed (owned by the directory entry); do not Destroy it.
func (e VSIDirEntry) Extra() CSLConstList {
	return cslConstList(e.cValue.papszExtra)
}

func (d VSIDir) GetNextDirEntry() (result VSIDirEntry) {
	return vsiGetNextDirEntry(d)
}

func (d VSIDir) Close() {
	vsiCloseDir(d)
}

func VSIMkdir(path string, mode int) (result int) {
	return vsiMkdir(path, mode)
}

func VSIMkdirRecursive(path string, mode int) (result int) {
	return vsiMkdirRecursive(path, mode)
}

func VSIRmdir(dirname string) (result int) {
	return vsiRmdir(dirname)
}

func VSIRmdirRecursive(dirname string) (result int) {
	return vsiRmdirRecursive(dirname)
}

func VSIUnlink(filename string) (result int) {
	return vsiUnlink(filename)
}

func VSIUnlinkBatch(files CSLConstList) (result []int) {
	return vsiUnlinkBatch(files)
}

func VSIRename(oldpath, newpath string) (result int) {
	return vsiRename(oldpath, newpath)
}

func VSIMove(oldpath, newpath string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int) {
	return vsiMove(oldpath, newpath, options, progress, progressData)
}

func VSICopyFile(source, target string, fpSource VSILFile, sourceSize VSILOffset, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int) {
	return vsiCopyFile(source, target, fpSource, sourceSize, options, progress, progressData)
}

func VSICopyFileRestartable(source, target, inputPayload string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int, outputPayload string) {
	return vsiCopyFileRestartable(source, target, inputPayload, options, progress, progressData)
}

// VSISync mirrors files/directories from source to target. outputs is an owned
// list of the files produced; the caller must Destroy it.
func VSISync(source, target string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int, outputs CSLConstList) {
	return vsiSync(source, target, options, progress, progressData)
}

func VSIMultipartUploadGetCapabilities(filename string) (result, nonSequentialUploadSupported, parallelUploadSupported, abortSupported int, minPartSize, maxPartSize uint64, maxPartCount int) {
	return vsiMultipartUploadGetCapabilities(filename)
}

func VSIMultipartUploadStart(filename string, options CSLConstList) (result string) {
	return vsiMultipartUploadStart(filename, options)
}

func VSIMultipartUploadAddPart(filename, uploadId string, partNumber int, fileOffset VSILOffset, data []byte, options CSLConstList) (result string) {
	return vsiMultipartUploadAddPart(filename, uploadId, partNumber, fileOffset, data, options)
}

func VSIMultipartUploadEnd(filename, uploadId string, partIds CSLConstList, totalSize VSILOffset, options CSLConstList) (result int) {
	return vsiMultipartUploadEnd(filename, uploadId, partIds, totalSize, options)
}

func VSIMultipartUploadAbort(filename, uploadId string, options CSLConstList) (result int) {
	return vsiMultipartUploadAbort(filename, uploadId, options)
}

func VSIAbortPendingUploads(filename string) (result int) {
	return vsiAbortPendingUploads(filename)
}

func VSIStrerror(errnum int) (result string) {
	return vsiStrerror(errnum)
}

func VSIGetDiskFreeSpace(dirname string) (result GIntBig) {
	return vsiGetDiskFreeSpace(dirname)
}

func VSINetworkStatsReset() {
	vsiNetworkStatsReset()
}

func VSINetworkStatsGetAsSerializedJSON(options CSLConstList) (result string) {
	return vsiNetworkStatsGetAsSerializedJSON(options)
}

func VSIInstallMemFileHandler() {
	vsiInstallMemFileHandler()
}

func VSIInstallLargeFileHandler() {
	vsiInstallLargeFileHandler()
}

func VSIInstallSubFileHandler() {
	vsiInstallSubFileHandler()
}

func VSICurlClearCache() {
	vsiCurlClearCache()
}

func VSICurlPartialClearCache(filenamePrefix string) {
	vsiCurlPartialClearCache(filenamePrefix)
}

func VSIInstallSparseFileHandler() {
	vsiInstallSparseFileHandler()
}

func VSIInstallCryptFileHandler() {
	vsiInstallCryptFileHandler()
}

func VSISetCryptKey(key []byte) {
	vsiSetCryptKey(key)
}

func VSICleanupFileManager() {
	vsiCleanupFileManager()
}

func VSIDuplicateFileSystemHandler(sourceFSName, newFSName string) (result bool) {
	return vsiDuplicateFileSystemHandler(sourceFSName, newFSName)
}

// VSIFileFromMemBuffer creates an in-memory file backed by data. When
// takeOwnership is 0 the buffer is referenced (not copied), so data must remain
// valid for the lifetime of the file handle.
func VSIFileFromMemBuffer(filename string, data []byte, takeOwnership int) (result VSILFile, err error) {
	result = vsiFileFromMemBuffer(filename, data, takeOwnership)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func VSIGetMemFileBuffer(filename string, unlinkAndSeize int) (result []byte) {
	return vsiGetMemFileBuffer(filename, unlinkAndSeize)
}

func VSIMemGenerateHiddenFilename(filename string) (result string) {
	return vsiMemGenerateHiddenFilename(filename)
}

func VSIAllocFilesystemPluginCallbacksStruct() (result VSIFilesystemPluginCallbacksStruct) {
	return vsiAllocFilesystemPluginCallbacksStruct()
}

func (cb VSIFilesystemPluginCallbacksStruct) Free() {
	vsiFreeFilesystemPluginCallbacksStruct(cb)
}

func VSIInstallPluginHandler(prefix string, cb VSIFilesystemPluginCallbacksStruct) (result int) {
	return vsiInstallPluginHandler(prefix, cb)
}

func VSIRemovePluginHandler(prefix string) (result int) {
	return vsiRemovePluginHandler(prefix)
}

func VSITime() (result uint64) {
	return vsiTime()
}

func VSICTime(t uint64) (result string) {
	return vsiCTime(t)
}
