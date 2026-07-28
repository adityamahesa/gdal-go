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
var VSILOffsetMax = VSILOffset(C._VSI_L_OFFSET_MAX())

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

func vsiFOpenExL(filename, access string, setError int) (result VSILFile) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cAccess := C.CString(access)
	defer C.free(unsafe.Pointer(cAccess))
	result = VSILFile{cValue: C.VSIFOpenExL(cFilename, cAccess, C.int(setError))}
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

func vsiFCloseL(file VSILFile) (result int) {
	result = int(C.VSIFCloseL(file.cValue))
	return
}

func vsiFSeekL(file VSILFile, offset VSILOffset, whence int) (result int) {
	result = int(C.VSIFSeekL(file.cValue, C.vsi_l_offset(offset), C.int(whence)))
	return
}

func vsiFTellL(file VSILFile) (result VSILOffset) {
	result = VSILOffset(C.VSIFTellL(file.cValue))
	return
}

func vsiRewindL(file VSILFile) {
	C.VSIRewindL(file.cValue)
}

func vsiFReadL(buffer []byte, size, count uint64, file VSILFile) (result uint64) {
	result = uint64(C.VSIFReadL(cBytes(buffer), C.size_t(size), C.size_t(count), file.cValue))
	return
}

func vsiFReadMultiRangeL(nRanges int, ppData []unsafe.Pointer, panOffsets []VSILOffset, panSizes []uint64, file VSILFile) (result int) {
	if nRanges <= 0 {
		return
	}
	cOffsets := make([]C.vsi_l_offset, nRanges)
	cSizes := make([]C.size_t, nRanges)
	for i := 0; i < nRanges; i++ {
		cOffsets[i] = C.vsi_l_offset(panOffsets[i])
		cSizes[i] = C.size_t(panSizes[i])
	}
	result = int(C.VSIFReadMultiRangeL(C.int(nRanges), &ppData[0], &cOffsets[0], &cSizes[0], file.cValue))
	return
}

func vsiFWriteL(buffer []byte, size, count uint64, file VSILFile) (result uint64) {
	result = uint64(C.VSIFWriteL(cBytes(buffer), C.size_t(size), C.size_t(count), file.cValue))
	return
}

func vsiFClearErrL(file VSILFile) {
	C.VSIFClearErrL(file.cValue)
}

func vsiFErrorL(file VSILFile) (result int) {
	result = int(C.VSIFErrorL(file.cValue))
	return
}

func vsiFEofL(file VSILFile) (result int) {
	result = int(C.VSIFEofL(file.cValue))
	return
}

func vsiFTruncateL(file VSILFile, newSize VSILOffset) (result int) {
	result = int(C.VSIFTruncateL(file.cValue, C.vsi_l_offset(newSize)))
	return
}

func vsiFFlushL(file VSILFile) (result int) {
	result = int(C.VSIFFlushL(file.cValue))
	return
}

// VSIFPrintfL is variadic: deferred — format in Go and use (VSILFile).WriteL.

func vsiFPutcL(c int, file VSILFile) (result int) {
	result = int(C.VSIFPutcL(C.int(c), file.cValue))
	return
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

func vsiOverwriteFile(target VSILFile, source string) (result int) {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))
	result = int(C.VSIOverwriteFile(target.cValue, cSource))
	return
}

// Type for VSIStatL(). Wraps a "struct stat"; use the accessor methods.
type VSIStatBufL struct {
	cValue C.VSIStatBufL
}

func vsiStatL(filename string) (result VSIStatBufL, ret int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	ret = int(C.VSIStatL(cFilename, &result.cValue))
	return
}

func vsiStatbufSize(s *VSIStatBufL) (result int64) {
	result = int64(C._vsi_statbuf_size(&s.cValue))
	return
}

func vsiStatbufMode(s *VSIStatBufL) (result int) {
	result = int(C._vsi_statbuf_mode(&s.cValue))
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

func vsiIsCaseSensitiveFS(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.VSIIsCaseSensitiveFS(cFilename))
	return
}

func vsiSupportsSparseFiles(path string) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSISupportsSparseFiles(cPath))
	return
}

func vsiIsLocal(path string) (result bool) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = bool(C.VSIIsLocal(cPath))
	return
}

func vsiGetCanonicalFilename(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSIGetCanonicalFilename(cPath)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func vsiSupportsSequentialWrite(path string, allowLocalTempFile bool) (result bool) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = bool(C.VSISupportsSequentialWrite(cPath, C.bool(allowLocalTempFile)))
	return
}

func vsiSupportsRandomWrite(path string, allowLocalTempFile bool) (result bool) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = bool(C.VSISupportsRandomWrite(cPath, C.bool(allowLocalTempFile)))
	return
}

func vsiHasOptimizedReadMultiRange(path string) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSIHasOptimizedReadMultiRange(cPath))
	return
}

func vsiGetActualURL(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.VSIGetActualURL(cFilename))
	return
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

func vsiGetFileSystemOptions(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.VSIGetFileSystemOptions(cFilename))
	return
}

func vsiGetFileSystemsPrefixes() (result CSLConstList) {
	raw := C.VSIGetFileSystemsPrefixes()
	result = cslConstList(raw)
	return
}

func vsiFGetNativeFileDescriptorL(file VSILFile) (result unsafe.Pointer) {
	result = C.VSIFGetNativeFileDescriptorL(file.cValue)
	return
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

func vsiSetPathSpecificOption(prefix, key, value string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.VSISetPathSpecificOption(cPrefix, cKey, cValue)
}

func vsiClearPathSpecificOptions(prefix string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	C.VSIClearPathSpecificOptions(cPrefix)
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

func vsiSetCredential(prefix, key, value string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.VSISetCredential(cPrefix, cKey, cValue)
}

func vsiClearCredentials(prefix string) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	C.VSIClearCredentials(cPrefix)
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

/* ==================================================================== */
/*      Memory allocation                                               */
/* ==================================================================== */

func vsiCalloc(count, size uint64) (result unsafe.Pointer) {
	result = C.VSICalloc(C.size_t(count), C.size_t(size))
	return
}

func vsiMalloc(size uint64) (result unsafe.Pointer) {
	result = C.VSIMalloc(C.size_t(size))
	return
}

func vsiFree(ptr unsafe.Pointer) {
	C.VSIFree(ptr)
}

func vsiRealloc(ptr unsafe.Pointer, size uint64) (result unsafe.Pointer) {
	result = C.VSIRealloc(ptr, C.size_t(size))
	return
}

func vsiStrdup(s string) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.VSIStrdup(cs)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

// VSIFreeReleaser is a C++-only helper and is skipped.

func vsiMallocAligned(alignment, size uint64) (result unsafe.Pointer) {
	result = C.VSIMallocAligned(C.size_t(alignment), C.size_t(size))
	return
}

func vsiMallocAlignedAuto(size uint64) (result unsafe.Pointer) {
	result = C.VSIMallocAlignedAuto(C.size_t(size))
	return
}

func vsiFreeAligned(ptr unsafe.Pointer) {
	C.VSIFreeAligned(ptr)
}

func vsiMallocAlignedAutoVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMallocAlignedAutoVerbose(C.size_t(size), cFile, C.int(line))
	return
}

// VSI_MALLOC_ALIGNED_AUTO_VERBOSE(size): deferred — value-producing macro to be wrapped in a later pass.

func vsiMalloc2(size1, size2 uint64) (result unsafe.Pointer) {
	result = C.VSIMalloc2(C.size_t(size1), C.size_t(size2))
	return
}

func vsiMalloc3(size1, size2, size3 uint64) (result unsafe.Pointer) {
	result = C.VSIMalloc3(C.size_t(size1), C.size_t(size2), C.size_t(size3))
	return
}

func vsiMallocVerbose(size uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMallocVerbose(C.size_t(size), cFile, C.int(line))
	return
}

// VSI_MALLOC_VERBOSE(size): deferred — value-producing macro to be wrapped in a later pass.

func vsiMalloc2Verbose(size1, size2 uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMalloc2Verbose(C.size_t(size1), C.size_t(size2), cFile, C.int(line))
	return
}

// VSI_MALLOC2_VERBOSE(nSize1, nSize2): deferred — value-producing macro to be wrapped in a later pass.

func vsiMalloc3Verbose(size1, size2, size3 uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIMalloc3Verbose(C.size_t(size1), C.size_t(size2), C.size_t(size3), cFile, C.int(line))
	return
}

// VSI_MALLOC3_VERBOSE(nSize1, nSize2, nSize3): deferred — value-producing macro to be wrapped in a later pass.

func vsiCallocVerbose(count, size uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSICallocVerbose(C.size_t(count), C.size_t(size), cFile, C.int(line))
	return
}

// VSI_CALLOC_VERBOSE(nCount, nSize): deferred — value-producing macro to be wrapped in a later pass.

func vsiReallocVerbose(oldPtr unsafe.Pointer, newSize uint64, file string, line int) (result unsafe.Pointer) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	result = C.VSIReallocVerbose(oldPtr, C.size_t(newSize), cFile, C.int(line))
	return
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

// VSI_STRDUP_VERBOSE(pszStr): deferred — value-producing macro to be wrapped in a later pass.

func cplGetPhysicalRAM() (result GIntBig) {
	result = GIntBig(C.CPLGetPhysicalRAM())
	return
}

func cplGetUsablePhysicalRAM() (result GIntBig) {
	result = GIntBig(C.CPLGetUsablePhysicalRAM())
	return
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

func vsiReadDirRecursive(path string) (result CSLConstList) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSIReadDirRecursive(cPath)
	result = cslConstList(raw)
	return
}

func vsiReadDirEx(path string, maxFiles int) (result CSLConstList) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSIReadDirEx(cPath, C.int(maxFiles))
	result = cslConstList(raw)
	return
}

func vsiSiblingFiles(path string) (result CSLConstList) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	raw := C.VSISiblingFiles(cPath)
	result = cslConstList(raw)
	return
}

func vsiGlob(pattern string, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result CSLConstList) {
	cPattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cPattern))
	opts := options.cValue
	raw := C.VSIGlob(cPattern, opts, progress.cValue, progressData)
	result = cslConstList(raw)
	return
}

func vsiGetDirectorySeparator(path string) (result string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = C.GoString(C.VSIGetDirectorySeparator(cPath))
	return
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

// Directory entry. The C++ ctor/dtor/copy members of VSIDIREntry are skipped;
// fields are exposed through accessor methods.
type VSIDirEntry struct {
	cValue *C.VSIDIREntry
}

func vsiGetNextDirEntry(dir VSIDir) (result VSIDirEntry) {
	result = VSIDirEntry{cValue: C.VSIGetNextDirEntry(dir.cValue)}
	return
}

func vsiDirEntryName(e VSIDirEntry) (result string) {
	result = C.GoString(e.cValue.pszName)
	return
}

func vsiCloseDir(dir VSIDir) {
	C.VSICloseDir(dir.cValue)
}

func vsiMkdir(path string, mode int) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSIMkdir(cPath, C.long(mode)))
	return
}

func vsiMkdirRecursive(path string, mode int) (result int) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result = int(C.VSIMkdirRecursive(cPath, C.long(mode)))
	return
}

func vsiRmdir(dirname string) (result int) {
	cDirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(cDirname))
	result = int(C.VSIRmdir(cDirname))
	return
}

func vsiRmdirRecursive(dirname string) (result int) {
	cDirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(cDirname))
	result = int(C.VSIRmdirRecursive(cDirname))
	return
}

func vsiUnlink(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.VSIUnlink(cFilename))
	return
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

func vsiRename(oldpath, newpath string) (result int) {
	cOld := C.CString(oldpath)
	defer C.free(unsafe.Pointer(cOld))
	cNew := C.CString(newpath)
	defer C.free(unsafe.Pointer(cNew))
	result = int(C.VSIRename(cOld, cNew))
	return
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

func vsiCopyFile(source, target string, fpSource VSILFile, sourceSize VSILOffset, options CSLConstList, progress GDALProgressFunc, progressData unsafe.Pointer) (result int) {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	opts := options.cValue
	result = int(C.VSICopyFile(cSource, cTarget, fpSource.cValue, C.vsi_l_offset(sourceSize), opts, progress.cValue, progressData))
	return
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

func vsiMultipartUploadStart(filename string, options CSLConstList) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	opts := options.cValue
	raw := C.VSIMultipartUploadStart(cFilename, opts)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
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

func vsiMultipartUploadEnd(filename, uploadId string, partIds CSLConstList, totalSize VSILOffset, options CSLConstList) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cUploadId := C.CString(uploadId)
	defer C.free(unsafe.Pointer(cUploadId))
	result = int(C.VSIMultipartUploadEnd(cFilename, cUploadId, C.size_t(partIds.Count()), partIds.cValue, C.vsi_l_offset(totalSize), options.cValue))
	return
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

func vsiAbortPendingUploads(filename string) (result int) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = int(C.VSIAbortPendingUploads(cFilename))
	return
}

func vsiStrerror(errnum int) (result string) {
	result = C.GoString(C.VSIStrerror(C.int(errnum)))
	return
}

func vsiGetDiskFreeSpace(dirname string) (result GIntBig) {
	cDirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(cDirname))
	result = GIntBig(C.VSIGetDiskFreeSpace(cDirname))
	return
}

func vsiNetworkStatsReset() {
	C.VSINetworkStatsReset()
}

func vsiNetworkStatsGetAsSerializedJSON(options CSLConstList) (result string) {
	opts := options.cValue
	raw := C.VSINetworkStatsGetAsSerializedJSON(opts)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

// VSIURIToVSIPath returns a std::string and is a C++-only cover; it is skipped.

/* ==================================================================== */
/*      Install special file access handlers.                           */
/* ==================================================================== */

func vsiInstallMemFileHandler() {
	C.VSIInstallMemFileHandler()
}

func vsiInstallLargeFileHandler() {
	C.VSIInstallLargeFileHandler()
}

func vsiInstallSubFileHandler() {
	C.VSIInstallSubFileHandler()
}

// The non-CPL_DLL install handlers (Curl/Curl-streaming/S3/GS/Azure/ADLS/OSS/
// Swift/7z/Rar/GZip/Zip/Stdin/Hdfs/WebHdfs/Stdout/Tar/Cached) are not part of the
// public exported ABI and are skipped (matching the cpl_port.go precedent).

func vsiCurlClearCache() {
	C.VSICurlClearCache()
}

func vsiCurlPartialClearCache(filenamePrefix string) {
	cPrefix := C.CString(filenamePrefix)
	defer C.free(unsafe.Pointer(cPrefix))
	C.VSICurlPartialClearCache(cPrefix)
}

func vsiInstallSparseFileHandler() {
	C.VSIInstallSparseFileHandler()
}

func vsiInstallCryptFileHandler() {
	C.VSIInstallCryptFileHandler()
}

func vsiSetCryptKey(key []byte) {
	C.VSISetCryptKey((*C.GByte)(cBytes(key)), C.int(len(key)))
}

func vsiCleanupFileManager() {
	C.VSICleanupFileManager()
}

func vsiDuplicateFileSystemHandler(sourceFSName, newFSName string) (result bool) {
	cSource := C.CString(sourceFSName)
	defer C.free(unsafe.Pointer(cSource))
	cNew := C.CString(newFSName)
	defer C.free(unsafe.Pointer(cNew))
	result = bool(C.VSIDuplicateFileSystemHandler(cSource, cNew))
	return
}

func vsiFileFromMemBuffer(filename string, data []byte, takeOwnership int) (result VSILFile) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = VSILFile{cValue: C.VSIFileFromMemBuffer(cFilename, (*C.GByte)(cBytes(data)), C.vsi_l_offset(len(data)), C.int(takeOwnership))}
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

func vsiMemGenerateHiddenFilename(filename string) (result string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	result = C.GoString(C.VSIMemGenerateHiddenFilename(cFilename))
	return
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

func vsiFreeFilesystemPluginCallbacksStruct(cb VSIFilesystemPluginCallbacksStruct) {
	C.VSIFreeFilesystemPluginCallbacksStruct(cb.cValue)
}

func vsiInstallPluginHandler(prefix string, cb VSIFilesystemPluginCallbacksStruct) (result int) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	result = int(C.VSIInstallPluginHandler(cPrefix, cb.cValue))
	return
}

func vsiRemovePluginHandler(prefix string) (result int) {
	cPrefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(cPrefix))
	result = int(C.VSIRemovePluginHandler(cPrefix))
	return
}

/* ==================================================================== */
/*      Time querying.                                                  */
/* ==================================================================== */

func vsiTime() (result uint64) {
	result = uint64(C.VSITime(nil))
	return
}

func vsiCTime(t uint64) (result string) {
	result = C.GoString(C.VSICTime(C.ulong(t)))
	return
}

// VSIGMTime and VSILocalTime operate on "struct tm"; deferred pending a struct tm wrapper.

// VSIDebug1..4 are logging macros that expand to nothing outside VSI_DEBUG builds; deferred.
