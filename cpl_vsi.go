package gdal

/*
#include "cpl_vsi_preamble.h"
*/
import "C"
import "unsafe"

/**
 * \file cpl_vsi.h
 *
 * Virtual System Interface (VSI): hookable covers over standard C I/O, memory
 * allocation and other system services, allowing virtualization of disk I/O.
 */

// The legacy stdio file-access covers (VSIFOpen/VSIFClose/VSIFSeek/VSIFTell/
// VSIRewind/VSIFFlush/VSIFRead/VSIFWrite/VSIFGets/VSIFPuts/VSIFPrintf/VSIFGetc/
// VSIFPutc/VSIUngetc/VSIFEof) operate on a raw C FILE* and predate the
// virtualization API. They are skipped in favor of the VSILFile large-file API.

// VSIStat()/VSIStatBuf and the VSI_IS* mode-test macros are part of the same
// legacy (non-virtualized) group and are skipped; use VSIStatL()/VSIStatExL().

/* ==================================================================== */
/*      64bit stdio file access functions.                              */
/* ==================================================================== */

// Type for a file offset
type VSILOffset C.vsi_l_offset

// Maximum value for a file offset
var VSILOffsetMax = VSILOffset(C._VSI_L_OFFSET_MAX)

// Opaque type for a FILE that implements the VSIVirtualHandle API
type VSILFile struct {
	cValue *C.VSILFILE
}

func vsiFOpenL(filename, access string) (result VSILFile) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cAccess := C.CString(access)
	defer C.free(unsafe.Pointer(cAccess))
	result = VSILFile{cValue: C.VSIFOpenL(cFilename, cAccess)}
	return
}

func VSIFOpenL(filename, access string) (result VSILFile, err error) {
	result = vsiFOpenL(filename, access)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func vsiFOpenExL(filename, access string, setError int) (result VSILFile) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cAccess := C.CString(access)
	defer C.free(unsafe.Pointer(cAccess))
	result = VSILFile{cValue: C.VSIFOpenExL(cFilename, cAccess, C.int(setError))}
	return
}

func VSIFOpenExL(filename, access string, setError int) (result VSILFile, err error) {
	result = vsiFOpenExL(filename, access, setError)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func vsiFOpenEx2L(filename, access string, setError int, options CSLConstList) (result VSILFile) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cAccess := C.CString(access)
	defer C.free(unsafe.Pointer(cAccess))
	result = VSILFile{cValue: C.VSIFOpenEx2L(cFilename, cAccess, C.int(setError), options.cValue)}
	return
}

func VSIFOpenEx2L(filename, access string, setError int, options CSLConstList) (result VSILFile, err error) {
	result = vsiFOpenEx2L(filename, access, setError, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func vsiFCloseL(file VSILFile) (result int) {
	result = int(C.VSIFCloseL(file.cValue))
	return
}

func (f VSILFile) CloseL() (result int) {
	return vsiFCloseL(f)
}

func vsiFSeekL(file VSILFile, offset VSILOffset, whence int) (result int) {
	result = int(C.VSIFSeekL(file.cValue, C.vsi_l_offset(offset), C.int(whence)))
	return
}

func (f VSILFile) SeekL(offset VSILOffset, whence int) (result int) {
	return vsiFSeekL(f, offset, whence)
}

func vsiFTellL(file VSILFile) (result VSILOffset) {
	result = VSILOffset(C.VSIFTellL(file.cValue))
	return
}

func (f VSILFile) TellL() (result VSILOffset) {
	return vsiFTellL(f)
}

func vsiRewindL(file VSILFile) {
	C.VSIRewindL(file.cValue)
}

func (f VSILFile) RewindL() {
	vsiRewindL(f)
}

func vsiFReadL(buffer []byte, size, count uint64, file VSILFile) (result uint64) {
	result = uint64(C.VSIFReadL(cBytes(buffer), C.size_t(size), C.size_t(count), file.cValue))
	return
}

// ReadL reads up to len(buffer) bytes into buffer and returns the number of
// bytes read (VSIFReadL called with a block size of 1).
func (f VSILFile) ReadL(buffer []byte) (result uint64) {
	return vsiFReadL(buffer, 1, uint64(len(buffer)), f)
}

func vsiFReadMultiRangeL(nRanges int, ppData *unsafe.Pointer, panOffsets *C.vsi_l_offset, panSizes *C.size_t, file VSILFile) (result int) {
	result = int(C.VSIFReadMultiRangeL(C.int(nRanges), ppData, panOffsets, panSizes, file.cValue))
	return
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
	cOffsets := make([]C.vsi_l_offset, n)
	cSizes := make([]C.size_t, n)
	for i := 0; i < n; i++ {
		ppData[i] = C.VSIMalloc(C.size_t(sizes[i]))
		cOffsets[i] = C.vsi_l_offset(offsets[i])
		cSizes[i] = C.size_t(sizes[i])
	}
	ret = vsiFReadMultiRangeL(n, &ppData[0], &cOffsets[0], &cSizes[0], f)
	result = make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = C.GoBytes(ppData[i], C.int(sizes[i]))
		C.VSIFree(ppData[i])
	}
	return
}

func vsiFWriteL(buffer []byte, size, count uint64, file VSILFile) (result uint64) {
	result = uint64(C.VSIFWriteL(cBytes(buffer), C.size_t(size), C.size_t(count), file.cValue))
	return
}

// WriteL writes len(buffer) bytes from buffer and returns the number of bytes
// written (VSIFWriteL called with a block size of 1).
func (f VSILFile) WriteL(buffer []byte) (result uint64) {
	return vsiFWriteL(buffer, 1, uint64(len(buffer)), f)
}

func vsiFClearErrL(file VSILFile) {
	C.VSIFClearErrL(file.cValue)
}

func (f VSILFile) ClearErrL() {
	vsiFClearErrL(f)
}

func vsiFErrorL(file VSILFile) (result int) {
	result = int(C.VSIFErrorL(file.cValue))
	return
}

func (f VSILFile) ErrorL() (result int) {
	return vsiFErrorL(f)
}

func vsiFEofL(file VSILFile) (result int) {
	result = int(C.VSIFEofL(file.cValue))
	return
}

func (f VSILFile) EofL() (result int) {
	return vsiFEofL(f)
}

func vsiFTruncateL(file VSILFile, newSize VSILOffset) (result int) {
	result = int(C.VSIFTruncateL(file.cValue, C.vsi_l_offset(newSize)))
	return
}

func (f VSILFile) TruncateL(newSize VSILOffset) (result int) {
	return vsiFTruncateL(f, newSize)
}

func vsiFFlushL(file VSILFile) (result int) {
	result = int(C.VSIFFlushL(file.cValue))
	return
}

func (f VSILFile) FlushL() (result int) {
	return vsiFFlushL(f)
}

// VSIFPrintfL is variadic: deferred — format in Go and use (VSILFile).WriteL.

func vsiFPutcL(c int, file VSILFile) (result int) {
	result = int(C.VSIFPutcL(C.int(c), file.cValue))
	return
}

func (f VSILFile) PutcL(c int) (result int) {
	return vsiFPutcL(c, f)
}

// Range status
type VSIRangeStatus C.VSIRangeStatus

const (
	VSIRangeStatusUnknown VSIRangeStatus = C.VSI_RANGE_STATUS_UNKNOWN
	VSIRangeStatusData    VSIRangeStatus = C.VSI_RANGE_STATUS_DATA
	VSIRangeStatusHole    VSIRangeStatus = C.VSI_RANGE_STATUS_HOLE
)

func vsiFGetRangeStatusL(file VSILFile, start, length VSILOffset) (result VSIRangeStatus) {
	result = VSIRangeStatus(C.VSIFGetRangeStatusL(file.cValue, C.vsi_l_offset(start), C.vsi_l_offset(length)))
	return
}

func (f VSILFile) GetRangeStatusL(start, length VSILOffset) (result VSIRangeStatus) {
	return vsiFGetRangeStatusL(f, start, length)
}

func vsiIngestFile(file VSILFile, filename string, maxSize int64) (result []byte, ret int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	var pabyRet *C.GByte
	var size C.vsi_l_offset
	ret = int(C.VSIIngestFile(file.cValue, cFilename, &pabyRet, &size, C.GIntBig(maxSize)))
	if pabyRet != nil {
		result = C.GoBytes(unsafe.Pointer(pabyRet), C.int(size))
		C.VSIFree(unsafe.Pointer(pabyRet))
	}
	return
}

func VSIIngestFile(file VSILFile, filename string, maxSize int64) (result []byte, err error) {
	var ret int
	result, ret = vsiIngestFile(file, filename, maxSize)
	if ret == 0 {
		err = lastError()
	}
	return
}

func vsiOverwriteFile(target VSILFile, source string) (result int) {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))
	result = int(C.VSIOverwriteFile(target.cValue, cSource))
	return
}

func VSIOverwriteFile(target VSILFile, source string) (result int) {
	return vsiOverwriteFile(target, source)
}

// Type for VSIStatL(). Wraps a "struct stat"; use the accessor methods.
type VSIStatBufL struct {
	cValue C.VSIStatBufL
}

// Size returns the st_size field (file size in bytes).
func (s VSIStatBufL) Size() int64 {
	return int64(C._vsi_statbuf_size(&s.cValue))
}

// Mode returns the st_mode field (see the VSI_IS* macros in the C header).
func (s VSIStatBufL) Mode() int {
	return int(C._vsi_statbuf_mode(&s.cValue))
}

func vsiStatL(filename string) (result VSIStatBufL, ret int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	ret = int(C.VSIStatL(cFilename, &result.cValue))
	return
}

func VSIStatL(filename string) (result VSIStatBufL, err error) {
	var ret int
	result, ret = vsiStatL(filename)
	if ret != 0 {
		err = lastError()
	}
	return
}

const (
	VSIStatExistsFlag   = C.VSI_STAT_EXISTS_FLAG
	VSIStatNatureFlag   = C.VSI_STAT_NATURE_FLAG
	VSIStatSizeFlag     = C.VSI_STAT_SIZE_FLAG
	VSIStatSetErrorFlag = C.VSI_STAT_SET_ERROR_FLAG
	VSIStatCacheOnly    = C.VSI_STAT_CACHE_ONLY
)

func vsiStatExL(filename string, flags int) (result VSIStatBufL, ret int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	ret = int(C.VSIStatExL(cFilename, &result.cValue, C.int(flags)))
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

func vsiIsCaseSensitiveFS(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.VSIIsCaseSensitiveFS(cFilename))
	return
}

func VSIIsCaseSensitiveFS(filename string) (result int) {
	return vsiIsCaseSensitiveFS(filename)
}

func vsiSupportsSparseFiles(path string) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSISupportsSparseFiles(cPath))
	return
}

func VSISupportsSparseFiles(path string) (result int) {
	return vsiSupportsSparseFiles(path)
}

func vsiIsLocal(path string) (result bool) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = bool(C.VSIIsLocal(cPath))
	return
}

func VSIIsLocal(path string) (result bool) {
	return vsiIsLocal(path)
}

func vsiGetCanonicalFilename(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSIGetCanonicalFilename(cPath)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func VSIGetCanonicalFilename(path string) (result string) {
	return vsiGetCanonicalFilename(path)
}

func vsiSupportsSequentialWrite(path string, allowLocalTempFile bool) (result bool) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = bool(C.VSISupportsSequentialWrite(cPath, C.bool(allowLocalTempFile)))
	return
}

func VSISupportsSequentialWrite(path string, allowLocalTempFile bool) (result bool) {
	return vsiSupportsSequentialWrite(path, allowLocalTempFile)
}

func vsiSupportsRandomWrite(path string, allowLocalTempFile bool) (result bool) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = bool(C.VSISupportsRandomWrite(cPath, C.bool(allowLocalTempFile)))
	return
}

func VSISupportsRandomWrite(path string, allowLocalTempFile bool) (result bool) {
	return vsiSupportsRandomWrite(path, allowLocalTempFile)
}

func vsiHasOptimizedReadMultiRange(path string) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSIHasOptimizedReadMultiRange(cPath))
	return
}

func VSIHasOptimizedReadMultiRange(path string) (result int) {
	return vsiHasOptimizedReadMultiRange(path)
}

func vsiGetActualURL(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.VSIGetActualURL(cFilename))
	return
}

func VSIGetActualURL(filename string) (result string) {
	return vsiGetActualURL(filename)
}

func vsiGetSignedURL(filename string, options CSLConstList) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	opts := options.cValue
	raw := C.VSIGetSignedURL(cFilename, opts)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func VSIGetSignedURL(filename string, options CSLConstList) (result string) {
	return vsiGetSignedURL(filename, options)
}

func vsiGetFileSystemOptions(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.VSIGetFileSystemOptions(cFilename))
	return
}

func VSIGetFileSystemOptions(filename string) (result string) {
	return vsiGetFileSystemOptions(filename)
}

func vsiGetFileSystemsPrefixes() (result CSLConstList) {
	raw := C.VSIGetFileSystemsPrefixes()
	result = cslConstList(raw)
	return
}

func VSIGetFileSystemsPrefixes() (result CSLConstList) {
	return vsiGetFileSystemsPrefixes()
}

func vsiFGetNativeFileDescriptorL(file VSILFile) (result unsafe.Pointer) {
	result = C.VSIFGetNativeFileDescriptorL(file.cValue)
	return
}

func (f VSILFile) GetNativeFileDescriptorL() (result unsafe.Pointer) {
	return vsiFGetNativeFileDescriptorL(f)
}

func vsiGetFileMetadata(filename, domain string, options CSLConstList) (result CSLConstList) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	opts := options.cValue
	raw := C.VSIGetFileMetadata(cFilename, cDomain, opts)
	result = cslConstList(raw)
	return
}

func VSIGetFileMetadata(filename, domain string, options CSLConstList) (result CSLConstList) {
	return vsiGetFileMetadata(filename, domain, options)
}

func vsiSetFileMetadata(filename string, metadata CSLConstList, domain string, options CSLConstList) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	md := metadata.cValue
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))
	opts := options.cValue
	result = int(C.VSISetFileMetadata(cFilename, md, cDomain, opts))
	return
}

func VSISetFileMetadata(filename string, metadata CSLConstList, domain string, options CSLConstList) (result int) {
	return vsiSetFileMetadata(filename, metadata, domain, options)
}

func vsiSetPathSpecificOption(prefix, key, value string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.VSISetPathSpecificOption(cPrefix, cKey, cValue)
}

func VSISetPathSpecificOption(prefix, key, value string) {
	vsiSetPathSpecificOption(prefix, key, value)
}

func vsiClearPathSpecificOptions(prefix string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	C.VSIClearPathSpecificOptions(cPrefix)
}

func VSIClearPathSpecificOptions(prefix string) {
	vsiClearPathSpecificOptions(prefix)
}

func vsiGetPathSpecificOption(path, key, dflt string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.VSIGetPathSpecificOption(cPath, cKey, cDflt))
	return
}

func VSIGetPathSpecificOption(path, key, dflt string) (result string) {
	return vsiGetPathSpecificOption(path, key, dflt)
}

func vsiSetCredential(prefix, key, value string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.VSISetCredential(cPrefix, cKey, cValue)
}

// Deprecated: use VSISetPathSpecificOption.
func VSISetCredential(prefix, key, value string) {
	vsiSetCredential(prefix, key, value)
}

func vsiClearCredentials(prefix string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	C.VSIClearCredentials(cPrefix)
}

// Deprecated: use VSIClearPathSpecificOptions.
func VSIClearCredentials(prefix string) {
	vsiClearCredentials(prefix)
}

func vsiGetCredential(path, key, dflt string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.VSIGetCredential(cPath, cKey, cDflt))
	return
}

// Deprecated: use VSIGetPathSpecificOption.
func VSIGetCredential(path, key, dflt string) (result string) {
	return vsiGetCredential(path, key, dflt)
}

/* ==================================================================== */
/*      Memory allocation                                               */
/* ==================================================================== */

func vsiCalloc(count, size uint64) (result unsafe.Pointer) {
	result = C.VSICalloc(C.size_t(count), C.size_t(size))
	return
}

func VSICalloc(count, size uint64) (result unsafe.Pointer) {
	return vsiCalloc(count, size)
}

func vsiMalloc(size uint64) (result unsafe.Pointer) {
	result = C.VSIMalloc(C.size_t(size))
	return
}

func VSIMalloc(size uint64) (result unsafe.Pointer) {
	return vsiMalloc(size)
}

func vsiFree(ptr unsafe.Pointer) {
	C.VSIFree(ptr)
}

func VSIFree(ptr unsafe.Pointer) {
	vsiFree(ptr)
}

func vsiRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	result = C.VSIRealloc(ptr, C.size_t(size))
	return
}

func VSIRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	return vsiRealloc(ptr, size)
}

func vsiStrdup(s string) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.VSIStrdup(cs)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func VSIStrdup(s string) (result string) {
	return vsiStrdup(s)
}

// VSIFreeReleaser is a C++-only helper and is skipped.

func vsiMallocAligned(alignment, size uint64) (result unsafe.Pointer) {
	result = C.VSIMallocAligned(C.size_t(alignment), C.size_t(size))
	return
}

func VSIMallocAligned(alignment, size uint64) (result unsafe.Pointer) {
	return vsiMallocAligned(alignment, size)
}

func vsiMallocAlignedAuto(size uint64) (result unsafe.Pointer) {
	result = C.VSIMallocAlignedAuto(C.size_t(size))
	return
}

func VSIMallocAlignedAuto(size uint64) (result unsafe.Pointer) {
	return vsiMallocAlignedAuto(size)
}

func vsiFreeAligned(ptr unsafe.Pointer) {
	C.VSIFreeAligned(ptr)
}

func VSIFreeAligned(ptr unsafe.Pointer) {
	vsiFreeAligned(ptr)
}

func vsiMallocAlignedAutoVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMallocAlignedAutoVerbose(C.size_t(size), cFile, C.int(line))
	return
}

func VSIMallocAlignedAutoVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMallocAlignedAutoVerbose(size, file, line)
}

// VSI_MALLOC_ALIGNED_AUTO_VERBOSE(size): deferred — value-producing macro to be wrapped in a later pass.

func vsiMalloc2(size1, size2 uint64) (result unsafe.Pointer) {
	result = C.VSIMalloc2(C.size_t(size1), C.size_t(size2))
	return
}

func VSIMalloc2(size1, size2 uint64) (result unsafe.Pointer) {
	return vsiMalloc2(size1, size2)
}

func vsiMalloc3(size1, size2, size3 uint64) (result unsafe.Pointer) {
	result = C.VSIMalloc3(C.size_t(size1), C.size_t(size2), C.size_t(size3))
	return
}

func VSIMalloc3(size1, size2, size3 uint64) (result unsafe.Pointer) {
	return vsiMalloc3(size1, size2, size3)
}

func vsiMallocVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMallocVerbose(C.size_t(size), cFile, C.int(line))
	return
}

func VSIMallocVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMallocVerbose(size, file, line)
}

// VSI_MALLOC_VERBOSE(size): deferred — value-producing macro to be wrapped in a later pass.

func vsiMalloc2Verbose(size1, size2 uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMalloc2Verbose(C.size_t(size1), C.size_t(size2), cFile, C.int(line))
	return
}

func VSIMalloc2Verbose(size1, size2 uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMalloc2Verbose(size1, size2, file, line)
}

// VSI_MALLOC2_VERBOSE(nSize1, nSize2): deferred — value-producing macro to be wrapped in a later pass.

func vsiMalloc3Verbose(size1, size2, size3 uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMalloc3Verbose(C.size_t(size1), C.size_t(size2), C.size_t(size3), cFile, C.int(line))
	return
}

func VSIMalloc3Verbose(size1, size2, size3 uint64, file string, line int) (result unsafe.Pointer) {
	return vsiMalloc3Verbose(size1, size2, size3, file, line)
}

// VSI_MALLOC3_VERBOSE(nSize1, nSize2, nSize3): deferred — value-producing macro to be wrapped in a later pass.

func vsiCallocVerbose(count, size uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSICallocVerbose(C.size_t(count), C.size_t(size), cFile, C.int(line))
	return
}

func VSICallocVerbose(count, size uint64, file string, line int) (result unsafe.Pointer) {
	return vsiCallocVerbose(count, size, file, line)
}

// VSI_CALLOC_VERBOSE(nCount, nSize): deferred — value-producing macro to be wrapped in a later pass.

func vsiReallocVerbose(oldPtr unsafe.Pointer, newSize uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIReallocVerbose(oldPtr, C.size_t(newSize), cFile, C.int(line))
	return
}

func VSIReallocVerbose(oldPtr unsafe.Pointer, newSize uint64, file string, line int) (result unsafe.Pointer) {
	return vsiReallocVerbose(oldPtr, newSize, file, line)
}

// VSI_REALLOC_VERBOSE(pOldPtr, nNewSize): deferred — value-producing macro to be wrapped in a later pass.

func vsiStrdupVerbose(s string, file string, line int) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	raw := C.VSIStrdupVerbose(cs, cFile, C.int(line))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func VSIStrdupVerbose(s string, file string, line int) (result string) {
	return vsiStrdupVerbose(s, file, line)
}

// VSI_STRDUP_VERBOSE(pszStr): deferred — value-producing macro to be wrapped in a later pass.

func cplGetPhysicalRAM() (result GIntBig) {
	result = GIntBig(C.CPLGetPhysicalRAM())
	return
}

func CPLGetPhysicalRAM() (result GIntBig) {
	return cplGetPhysicalRAM()
}

func cplGetUsablePhysicalRAM() (result GIntBig) {
	result = GIntBig(C.CPLGetUsablePhysicalRAM())
	return
}

func CPLGetUsablePhysicalRAM() (result GIntBig) {
	return cplGetUsablePhysicalRAM()
}

/* ==================================================================== */
/*      Other...                                                        */
/* ==================================================================== */

func vsiReadDir(path string) (result CSLConstList) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSIReadDir(cPath)
	result = cslConstList(raw)
	return
}

func VSIReadDir(path string) (result CSLConstList) {
	return vsiReadDir(path)
}

// Alias of VSIReadDir()
func CPLReadDir(path string) (result CSLConstList) {
	return vsiReadDir(path)
}

func vsiReadDirRecursive(path string) (result CSLConstList) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSIReadDirRecursive(cPath)
	result = cslConstList(raw)
	return
}

func VSIReadDirRecursive(path string) (result CSLConstList) {
	return vsiReadDirRecursive(path)
}

func vsiReadDirEx(path string, maxFiles int) (result CSLConstList) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSIReadDirEx(cPath, C.int(maxFiles))
	result = cslConstList(raw)
	return
}

func VSIReadDirEx(path string, maxFiles int) (result CSLConstList) {
	return vsiReadDirEx(path, maxFiles)
}

func vsiSiblingFiles(path string) (result CSLConstList) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSISiblingFiles(cPath)
	result = cslConstList(raw)
	return
}

func VSISiblingFiles(path string) (result CSLConstList) {
	return vsiSiblingFiles(path)
}

func vsiGlob(pattern string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CSLConstList) {
	cPattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cPattern))
	opts := options.cValue
	raw := C.VSIGlob(cPattern, opts, progress.cValue, progressData)
	result = cslConstList(raw)
	return
}

func VSIGlob(pattern string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CSLConstList) {
	return vsiGlob(pattern, options, progress, progressData)
}

func vsiGetDirectorySeparator(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.VSIGetDirectorySeparator(cPath))
	return
}

func VSIGetDirectorySeparator(path string) (result string) {
	return vsiGetDirectorySeparator(path)
}

// Opaque type for a directory iterator
type VSIDir struct {
	cValue *C.VSIDIR
}

func vsiOpenDir(path string, recurseDepth int, options CSLConstList) (result VSIDir) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	opts := options.cValue
	result = VSIDir{cValue: C.VSIOpenDir(cPath, C.int(recurseDepth), opts)}
	return
}

func VSIOpenDir(path string, recurseDepth int, options CSLConstList) (result VSIDir, err error) {
	result = vsiOpenDir(path, recurseDepth, options)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

// Directory entry. The C++ ctor/dtor/copy members of VSIDIREntry are skipped;
// fields are exposed through accessor methods.
type VSIDirEntry struct {
	cValue *C.VSIDIREntry
}

// Name returns the entry filename.
func (e VSIDirEntry) Name() string {
	return C.GoString(e.cValue.pszName)
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

func vsiGetNextDirEntry(dir VSIDir) (result VSIDirEntry) {
	result = VSIDirEntry{cValue: C.VSIGetNextDirEntry(dir.cValue)}
	return
}

func (d VSIDir) GetNextDirEntry() (result VSIDirEntry) {
	return vsiGetNextDirEntry(d)
}

func vsiCloseDir(dir VSIDir) {
	C.VSICloseDir(dir.cValue)
}

func (d VSIDir) Close() {
	vsiCloseDir(d)
}

func vsiMkdir(path string, mode int) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSIMkdir(cPath, C.long(mode)))
	return
}

func VSIMkdir(path string, mode int) (result int) {
	return vsiMkdir(path, mode)
}

func vsiMkdirRecursive(path string, mode int) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSIMkdirRecursive(cPath, C.long(mode)))
	return
}

func VSIMkdirRecursive(path string, mode int) (result int) {
	return vsiMkdirRecursive(path, mode)
}

func vsiRmdir(dirname string) (result int) {
	cDirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(cDirname))
	result = int(C.VSIRmdir(cDirname))
	return
}

func VSIRmdir(dirname string) (result int) {
	return vsiRmdir(dirname)
}

func vsiRmdirRecursive(dirname string) (result int) {
	cDirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(cDirname))
	result = int(C.VSIRmdirRecursive(cDirname))
	return
}

func VSIRmdirRecursive(dirname string) (result int) {
	return vsiRmdirRecursive(dirname)
}

func vsiUnlink(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.VSIUnlink(cFilename))
	return
}

func VSIUnlink(filename string) (result int) {
	return vsiUnlink(filename)
}

func vsiUnlinkBatch(files CSLConstList) (result []int) {
	arr := files.cValue
	raw := C.VSIUnlinkBatch(arr)
	if raw == nil {
		return
	}
	defer C.VSIFree(unsafe.Pointer(raw))
	n := files.Count()
	slice := unsafe.Slice(raw, n)
	result = make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = int(slice[i])
	}
	return
}

func VSIUnlinkBatch(files CSLConstList) (result []int) {
	return vsiUnlinkBatch(files)
}

func vsiRename(oldpath, newpath string) (result int) {
	cOld := C.CString(oldpath)
	defer C.free(unsafe.Pointer(cOld))
	cNew := C.CString(newpath)
	defer C.free(unsafe.Pointer(cNew))
	result = int(C.VSIRename(cOld, cNew))
	return
}

func VSIRename(oldpath, newpath string) (result int) {
	return vsiRename(oldpath, newpath)
}

func vsiMove(oldpath, newpath string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int) {
	cOld := C.CString(oldpath)
	defer C.free(unsafe.Pointer(cOld))
	cNew := C.CString(newpath)
	defer C.free(unsafe.Pointer(cNew))
	opts := options.cValue
	result = int(C.VSIMove(cOld, cNew, opts, progress.cValue, progressData))
	return
}

func VSIMove(oldpath, newpath string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int) {
	return vsiMove(oldpath, newpath, options, progress, progressData)
}

func vsiCopyFile(source, target string, fpSource VSILFile, sourceSize VSILOffset, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int) {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	opts := options.cValue
	result = int(C.VSICopyFile(cSource, cTarget, fpSource.cValue, C.vsi_l_offset(sourceSize), opts, progress.cValue, progressData))
	return
}

func VSICopyFile(source, target string, fpSource VSILFile, sourceSize VSILOffset, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int) {
	return vsiCopyFile(source, target, fpSource, sourceSize, options, progress, progressData)
}

func vsiCopyFileRestartable(source, target, inputPayload string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int, outputPayload string) {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	cInput := C.CString(inputPayload)
	defer C.free(unsafe.Pointer(cInput))
	opts := options.cValue
	var cOutput *C.char
	result = int(C.VSICopyFileRestartable(cSource, cTarget, cInput, &cOutput, opts, progress.cValue, progressData))
	if cOutput != nil {
		outputPayload = C.GoString(cOutput)
		C.VSIFree(unsafe.Pointer(cOutput))
	}
	return
}

func VSICopyFileRestartable(source, target, inputPayload string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int, outputPayload string) {
	return vsiCopyFileRestartable(source, target, inputPayload, options, progress, progressData)
}

func vsiSync(source, target string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int, outputs CSLConstList) {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	var cOutputs **C.char
	result = int(C.VSISync(cSource, cTarget, options.cValue, progress.cValue, progressData, &cOutputs))
	outputs = cslConstList(cOutputs)
	return
}

// VSISync mirrors files/directories from source to target. outputs is an owned
// list of the files produced; the caller must Destroy it.
func VSISync(source, target string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int, outputs CSLConstList) {
	return vsiSync(source, target, options, progress, progressData)
}

func vsiMultipartUploadGetCapabilities(filename string) (result, nonSequentialUploadSupported, parallelUploadSupported, abortSupported int, minPartSize, maxPartSize uint64, maxPartCount int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	var nonSeq, parallel, abort, maxCount C.int
	var minPart, maxPart C.size_t
	result = int(C.VSIMultipartUploadGetCapabilities(cFilename, &nonSeq, &parallel, &abort, &minPart, &maxPart, &maxCount))
	nonSequentialUploadSupported = int(nonSeq)
	parallelUploadSupported = int(parallel)
	abortSupported = int(abort)
	minPartSize = uint64(minPart)
	maxPartSize = uint64(maxPart)
	maxPartCount = int(maxCount)
	return
}

func VSIMultipartUploadGetCapabilities(filename string) (result, nonSequentialUploadSupported, parallelUploadSupported, abortSupported int, minPartSize, maxPartSize uint64, maxPartCount int) {
	return vsiMultipartUploadGetCapabilities(filename)
}

func vsiMultipartUploadStart(filename string, options CSLConstList) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	opts := options.cValue
	raw := C.VSIMultipartUploadStart(cFilename, opts)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func VSIMultipartUploadStart(filename string, options CSLConstList) (result string) {
	return vsiMultipartUploadStart(filename, options)
}

func vsiMultipartUploadAddPart(filename, uploadId string, partNumber int, fileOffset VSILOffset, data []byte, options CSLConstList) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cUploadId := C.CString(uploadId)
	defer C.free(unsafe.Pointer(cUploadId))
	opts := options.cValue
	raw := C.VSIMultipartUploadAddPart(cFilename, cUploadId, C.int(partNumber), C.vsi_l_offset(fileOffset), cBytes(data), C.size_t(len(data)), opts)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func VSIMultipartUploadAddPart(filename, uploadId string, partNumber int, fileOffset VSILOffset, data []byte, options CSLConstList) (result string) {
	return vsiMultipartUploadAddPart(filename, uploadId, partNumber, fileOffset, data, options)
}

func vsiMultipartUploadEnd(filename, uploadId string, partIds CSLConstList, totalSize VSILOffset, options CSLConstList) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cUploadId := C.CString(uploadId)
	defer C.free(unsafe.Pointer(cUploadId))
	result = int(C.VSIMultipartUploadEnd(cFilename, cUploadId, C.size_t(partIds.Count()), partIds.cValue, C.vsi_l_offset(totalSize), options.cValue))
	return
}

func VSIMultipartUploadEnd(filename, uploadId string, partIds CSLConstList, totalSize VSILOffset, options CSLConstList) (result int) {
	return vsiMultipartUploadEnd(filename, uploadId, partIds, totalSize, options)
}

func vsiMultipartUploadAbort(filename, uploadId string, options CSLConstList) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cUploadId := C.CString(uploadId)
	defer C.free(unsafe.Pointer(cUploadId))
	opts := options.cValue
	result = int(C.VSIMultipartUploadAbort(cFilename, cUploadId, opts))
	return
}

func VSIMultipartUploadAbort(filename, uploadId string, options CSLConstList) (result int) {
	return vsiMultipartUploadAbort(filename, uploadId, options)
}

func vsiAbortPendingUploads(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.VSIAbortPendingUploads(cFilename))
	return
}

func VSIAbortPendingUploads(filename string) (result int) {
	return vsiAbortPendingUploads(filename)
}

func vsiStrerror(errnum int) (result string) {
	result = C.GoString(C.VSIStrerror(C.int(errnum)))
	return
}

func VSIStrerror(errnum int) (result string) {
	return vsiStrerror(errnum)
}

func vsiGetDiskFreeSpace(dirname string) (result GIntBig) {
	cDirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(cDirname))
	result = GIntBig(C.VSIGetDiskFreeSpace(cDirname))
	return
}

func VSIGetDiskFreeSpace(dirname string) (result GIntBig) {
	return vsiGetDiskFreeSpace(dirname)
}

func vsiNetworkStatsReset() {
	C.VSINetworkStatsReset()
}

func VSINetworkStatsReset() {
	vsiNetworkStatsReset()
}

func vsiNetworkStatsGetAsSerializedJSON(options CSLConstList) (result string) {
	opts := options.cValue
	raw := C.VSINetworkStatsGetAsSerializedJSON(opts)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func VSINetworkStatsGetAsSerializedJSON(options CSLConstList) (result string) {
	return vsiNetworkStatsGetAsSerializedJSON(options)
}

// VSIURIToVSIPath returns a std::string and is a C++-only cover; it is skipped.

/* ==================================================================== */
/*      Install special file access handlers.                           */
/* ==================================================================== */

func vsiInstallMemFileHandler() {
	C.VSIInstallMemFileHandler()
}

func VSIInstallMemFileHandler() {
	vsiInstallMemFileHandler()
}

func vsiInstallLargeFileHandler() {
	C.VSIInstallLargeFileHandler()
}

func VSIInstallLargeFileHandler() {
	vsiInstallLargeFileHandler()
}

func vsiInstallSubFileHandler() {
	C.VSIInstallSubFileHandler()
}

func VSIInstallSubFileHandler() {
	vsiInstallSubFileHandler()
}

// The non-CPL_DLL install handlers (Curl/Curl-streaming/S3/GS/Azure/ADLS/OSS/
// Swift/7z/Rar/GZip/Zip/Stdin/Hdfs/WebHdfs/Stdout/Tar/Cached) are not part of the
// public exported ABI and are skipped (matching the cpl_port.go precedent).

func vsiCurlClearCache() {
	C.VSICurlClearCache()
}

func VSICurlClearCache() {
	vsiCurlClearCache()
}

func vsiCurlPartialClearCache(filenamePrefix string) {
	cPrefix := C.CString(filenamePrefix)
	defer C.free(unsafe.Pointer(cPrefix))
	C.VSICurlPartialClearCache(cPrefix)
}

func VSICurlPartialClearCache(filenamePrefix string) {
	vsiCurlPartialClearCache(filenamePrefix)
}

func vsiInstallSparseFileHandler() {
	C.VSIInstallSparseFileHandler()
}

func VSIInstallSparseFileHandler() {
	vsiInstallSparseFileHandler()
}

func vsiInstallCryptFileHandler() {
	C.VSIInstallCryptFileHandler()
}

func VSIInstallCryptFileHandler() {
	vsiInstallCryptFileHandler()
}

func vsiSetCryptKey(key []byte) {
	C.VSISetCryptKey((*C.GByte)(cBytes(key)), C.int(len(key)))
}

func VSISetCryptKey(key []byte) {
	vsiSetCryptKey(key)
}

func vsiCleanupFileManager() {
	C.VSICleanupFileManager()
}

func VSICleanupFileManager() {
	vsiCleanupFileManager()
}

func vsiDuplicateFileSystemHandler(sourceFSName, newFSName string) (result bool) {
	cSource := C.CString(sourceFSName)
	defer C.free(unsafe.Pointer(cSource))
	cNew := C.CString(newFSName)
	defer C.free(unsafe.Pointer(cNew))
	result = bool(C.VSIDuplicateFileSystemHandler(cSource, cNew))
	return
}

func VSIDuplicateFileSystemHandler(sourceFSName, newFSName string) (result bool) {
	return vsiDuplicateFileSystemHandler(sourceFSName, newFSName)
}

func vsiFileFromMemBuffer(filename string, data []byte, takeOwnership int) (result VSILFile) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = VSILFile{cValue: C.VSIFileFromMemBuffer(cFilename, (*C.GByte)(cBytes(data)), C.vsi_l_offset(len(data)), C.int(takeOwnership))}
	return
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

func vsiGetMemFileBuffer(filename string, unlinkAndSeize int) (result []byte) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	var length C.vsi_l_offset
	raw := C.VSIGetMemFileBuffer(cFilename, &length, C.int(unlinkAndSeize))
	if raw == nil {
		return
	}
	result = C.GoBytes(unsafe.Pointer(raw), C.int(length))
	if unlinkAndSeize != 0 {
		C.VSIFree(unsafe.Pointer(raw))
	}
	return
}

func VSIGetMemFileBuffer(filename string, unlinkAndSeize int) (result []byte) {
	return vsiGetMemFileBuffer(filename, unlinkAndSeize)
}

func vsiMemGenerateHiddenFilename(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.VSIMemGenerateHiddenFilename(cFilename))
	return
}

func VSIMemGenerateHiddenFilename(filename string) (result string) {
	return vsiMemGenerateHiddenFilename(filename)
}

// VSIWriteFunction (a C function-pointer typedef) and VSIStdoutSetRedirection
// require a C-created callback and a raw FILE*; both are deferred.

// The filesystem plugin callback typedefs (VSIFilesystemPluginStatCallback,
// UnlinkCallback, RenameCallback, MkdirCallback, RmdirCallback, ReadDirCallback,
// SiblingFilesCallback, OpenCallback, TellCallback, SeekCallback, ReadCallback,
// ReadMultiRangeCallback, GetRangeStatusCallback, EofCallback, WriteCallback,
// FlushCallback, TruncateCallback, CloseCallback, AdviseReadCallback,
// ErrorCallback, ClearErrCallback) bridge C into user code and are deferred; the
// callbacks struct below is wrapped as an opaque handle only.

// Opaque handle to a VSIFilesystemPluginCallbacksStruct. Its individual callback
// fields are not yet bridged to Go functions.
type VSIFilesystemPluginCallbacksStruct struct {
	cValue *C.VSIFilesystemPluginCallbacksStruct
}

func vsiAllocFilesystemPluginCallbacksStruct() (result VSIFilesystemPluginCallbacksStruct) {
	result = VSIFilesystemPluginCallbacksStruct{cValue: C.VSIAllocFilesystemPluginCallbacksStruct()}
	return
}

func VSIAllocFilesystemPluginCallbacksStruct() (result VSIFilesystemPluginCallbacksStruct) {
	return vsiAllocFilesystemPluginCallbacksStruct()
}

func vsiFreeFilesystemPluginCallbacksStruct(cb VSIFilesystemPluginCallbacksStruct) {
	C.VSIFreeFilesystemPluginCallbacksStruct(cb.cValue)
}

func (cb VSIFilesystemPluginCallbacksStruct) Free() {
	vsiFreeFilesystemPluginCallbacksStruct(cb)
}

func vsiInstallPluginHandler(prefix string, cb VSIFilesystemPluginCallbacksStruct) (result int) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	result = int(C.VSIInstallPluginHandler(cPrefix, cb.cValue))
	return
}

func VSIInstallPluginHandler(prefix string, cb VSIFilesystemPluginCallbacksStruct) (result int) {
	return vsiInstallPluginHandler(prefix, cb)
}

func vsiRemovePluginHandler(prefix string) (result int) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	result = int(C.VSIRemovePluginHandler(cPrefix))
	return
}

func VSIRemovePluginHandler(prefix string) (result int) {
	return vsiRemovePluginHandler(prefix)
}

/* ==================================================================== */
/*      Time querying.                                                  */
/* ==================================================================== */

func vsiTime() (result uint64) {
	result = uint64(C.VSITime(nil))
	return
}

func VSITime() (result uint64) {
	return vsiTime()
}

func vsiCTime(t uint64) (result string) {
	result = C.GoString(C.VSICTime(C.ulong(t)))
	return
}

func VSICTime(t uint64) (result string) {
	return vsiCTime(t)
}

// VSIGMTime and VSILocalTime operate on "struct tm"; deferred pending a struct tm wrapper.

// VSIDebug1..4 are logging macros that expand to nothing outside VSI_DEBUG builds; deferred.
