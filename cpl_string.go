package gdal

/*
#include "cpl_string_preamble.h"
*/
import "C"
import "unsafe"

/**
 * \file cpl_string.h
 *
 * Various convenience functions for working with strings and string lists.
 *
 * A StringList is an array of strings with a NULL last pointer, wrapped here by
 * the CSLConstList handle (see cpl_port.go). Build one by chaining AddString
 * from NullCSLConstList and read it back with Count/GetField. A common
 * convention stores name/value pairs formatted as "<name>=<value>" (":" is also
 * accepted).
 */

func cslAddString(l CSLConstList, s string) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CSLAddString(l.cValue, cs)
	result = cslConstList(raw)
	return
}

func cslAddStringMayFail(l CSLConstList, s string) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CSLAddStringMayFail(l.cValue, cs)
	result = cslConstList(raw)
	return
}

func cslCount(l CSLConstList) (result int) {
	result = int(C.CSLCount(l.cValue))
	return
}

func cslGetField(l CSLConstList, i int) (result string) {
	result = C.GoString(C.CSLGetField(l.cValue, C.int(i)))
	return
}

func cslDestroy(l CSLConstList) {
	C.CSLDestroy(l.cValue)
}

func cslDuplicate(l CSLConstList) (result CSLConstList) {
	raw := C.CSLDuplicate(l.cValue)
	result = cslConstList(raw)
	return
}

func cslMerge(l CSLConstList, override CSLConstList) (result CSLConstList) {
	raw := C.CSLMerge(l.cValue, override.cValue)
	result = cslConstList(raw)
	return
}

func cslTokenizeString(s string) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CSLTokenizeString(cs)
	result = cslConstList(raw)
	return
}

func cslTokenizeStringComplex(s, delimiter string, honourStrings, allowEmptyTokens int) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	cDelimiter := C.CString(delimiter)
	defer C.free(unsafe.Pointer(cDelimiter))
	raw := C.CSLTokenizeStringComplex(cs, cDelimiter, C.int(honourStrings), C.int(allowEmptyTokens))
	result = cslConstList(raw)
	return
}

func cslTokenizeString2(s, delimiter string, flags int) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	cDelimiter := C.CString(delimiter)
	defer C.free(unsafe.Pointer(cDelimiter))
	raw := C.CSLTokenizeString2(cs, cDelimiter, C.int(flags))
	result = cslConstList(raw)
	return
}

const (
	// Flag for CSLTokenizeString2() to honour strings
	CSLTHonourStrings = C.CSLT_HONOURSTRINGS
	// Flag for CSLTokenizeString2() to allow empty tokens
	CSLTAllowEmptyTokens = C.CSLT_ALLOWEMPTYTOKENS
	// Flag for CSLTokenizeString2() to preserve quotes
	CSLTPreserveQuotes = C.CSLT_PRESERVEQUOTES
	// Flag for CSLTokenizeString2() to preserve escape characters
	CSLTPreserveEscapes = C.CSLT_PRESERVEESCAPES
	// Flag for CSLTokenizeString2() to strip leading spaces
	CSLTStripLeadSpaces = C.CSLT_STRIPLEADSPACES
	// Flag for CSLTokenizeString2() to strip trailing spaces
	CSLTStripEndSpaces = C.CSLT_STRIPENDSPACES
)

func cslPrint(l CSLConstList, fp cFile) (result int) {
	result = int(C.CSLPrint(l.cValue, fp))
	return
}

func cslLoad(fname string) (result CSLConstList) {
	cFname := C.CString(fname)
	defer C.free(unsafe.Pointer(cFname))
	raw := C.CSLLoad(cFname)
	result = cslConstList(raw)
	return
}

func cslLoad2(fname string, maxLines, maxCols int, options CSLConstList) (result CSLConstList) {
	cFname := C.CString(fname)
	defer C.free(unsafe.Pointer(cFname))
	raw := C.CSLLoad2(cFname, C.int(maxLines), C.int(maxCols), options.cValue)
	result = cslConstList(raw)
	return
}

func cslSave(l CSLConstList, fname string) (result int) {
	cFname := C.CString(fname)
	defer C.free(unsafe.Pointer(cFname))
	result = int(C.CSLSave(l.cValue, cFname))
	return
}

func cslInsertStrings(l CSLConstList, at int, newLines CSLConstList) (result CSLConstList) {
	raw := C.CSLInsertStrings(l.cValue, C.int(at), newLines.cValue)
	result = cslConstList(raw)
	return
}

func cslInsertString(l CSLConstList, at int, s string) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CSLInsertString(l.cValue, C.int(at), cs)
	result = cslConstList(raw)
	return
}

func cslRemoveStrings(l CSLConstList, first, num int, retStrings *CSLConstList) (result CSLConstList) {
	// A nil papszRetStrings tells C to free the removed strings itself, so only
	// ask for them back when the caller supplied somewhere to put them.
	if retStrings == nil {
		result = cslConstList(C.CSLRemoveStrings(l.cValue, C.int(first), C.int(num), nil))
		return
	}
	var ret **C.char
	result = cslConstList(C.CSLRemoveStrings(l.cValue, C.int(first), C.int(num), &ret))
	*retStrings = cslConstList(ret)
	return
}

func cslFindString(l CSLConstList, target string) (result int) {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	result = int(C.CSLFindString(l.cValue, cTarget))
	return
}

func cslFindStringCaseSensitive(l CSLConstList, target string) (result int) {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	result = int(C.CSLFindStringCaseSensitive(l.cValue, cTarget))
	return
}

func cslPartialFindString(l CSLConstList, needle string) (result int) {
	cNeedle := C.CString(needle)
	defer C.free(unsafe.Pointer(cNeedle))
	result = int(C.CSLPartialFindString(l.cValue, cNeedle))
	return
}

func cslFindName(l CSLConstList, name string) (result int) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = int(C.CSLFindName(l.cValue, cName))
	return
}

func cslFetchBoolean(l CSLConstList, key string, dflt int) (result int) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	result = int(C.CSLFetchBoolean(l.cValue, cKey, C.int(dflt)))
	return
}

func cslTestBoolean(value string) (result int) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = int(C.CSLTestBoolean(cValue))
	return
}

func cplTestBoolean(value string) (result int) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = int(C.CPLTestBoolean(cValue))
	return
}

func cplTestBool(value string) (result bool) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = bool(C.CPLTestBool(cValue))
	return
}

func cplFetchBool(l CSLConstList, key string, dflt bool) (result bool) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	result = bool(C.CPLFetchBool(l.cValue, cKey, C.bool(dflt)))
	return
}

func cplParseMemorySize(value string, nValue *GIntBig, unitSpecified *bool) (ret CPLErr) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	ret = CPLErr(C.CPLParseMemorySize(cValue, (*C.GIntBig)(nValue), (*C.bool)(unsafe.Pointer(unitSpecified))))
	return
}

func cplParseNameValue(nameValue string, key *string) (result string) {
	cNameValue := C.CString(nameValue)
	defer C.free(unsafe.Pointer(cNameValue))
	var cKey *C.char
	result = C.GoString(C.CPLParseNameValue(cNameValue, &cKey))
	if cKey != nil {
		defer C.VSIFree(unsafe.Pointer(cKey))
		if key != nil {
			*key = C.GoString(cKey)
		}
	}
	return
}

func cplParseNameValueSep(nameValue string, key *string, sep byte) (result string) {
	cNameValue := C.CString(nameValue)
	defer C.free(unsafe.Pointer(cNameValue))
	var cKey *C.char
	result = C.GoString(C.CPLParseNameValueSep(cNameValue, &cKey, C.char(sep)))
	if cKey != nil {
		defer C.VSIFree(unsafe.Pointer(cKey))
		if key != nil {
			*key = C.GoString(cKey)
		}
	}
	return
}

func cslFetchNameValue(l CSLConstList, name string) (result string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = C.GoString(C.CSLFetchNameValue(l.cValue, cName))
	return
}

func cslFetchNameValueDef(l CSLConstList, name, dflt string) (result string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CSLFetchNameValueDef(l.cValue, cName, cDflt))
	return
}

func cslFetchNameValueMultiple(l CSLConstList, name string) (result CSLConstList) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	raw := C.CSLFetchNameValueMultiple(l.cValue, cName)
	result = cslConstList(raw)
	return
}

func cslAddNameValue(l CSLConstList, name, value string) (result CSLConstList) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	raw := C.CSLAddNameValue(l.cValue, cName, cValue)
	result = cslConstList(raw)
	return
}

func cslSetNameValue(l CSLConstList, name, value string) (result CSLConstList) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	raw := C.CSLSetNameValue(l.cValue, cName, cValue)
	result = cslConstList(raw)
	return
}

func cslSetNameValueSeparator(l CSLConstList, separator string) {
	cSeparator := C.CString(separator)
	defer C.free(unsafe.Pointer(cSeparator))
	C.CSLSetNameValueSeparator(l.cValue, cSeparator)
}

func cslParseCommandLine(commandLine string) (result CSLConstList) {
	cCommandLine := C.CString(commandLine)
	defer C.free(unsafe.Pointer(cCommandLine))
	raw := C.CSLParseCommandLine(cCommandLine)
	result = cslConstList(raw)
	return
}

const (
	// Scheme for CPLEscapeString()/CPLUnescapeString() for backslash quoting
	CPLESBackslashQuotable = C.CPLES_BackslashQuotable
	// Scheme for CPLEscapeString()/CPLUnescapeString() for XML
	CPLESXML = C.CPLES_XML
	// Scheme for CPLEscapeString()/CPLUnescapeString() for URL
	CPLESURL = C.CPLES_URL
	// Scheme for CPLEscapeString()/CPLUnescapeString() for SQL
	CPLESSQL = C.CPLES_SQL
	// Scheme for CPLEscapeString()/CPLUnescapeString() for CSV
	CPLESCSV = C.CPLES_CSV
	// Scheme for CPLEscapeString()/CPLUnescapeString() for XML (preserves quotes)
	CPLESXMLButQuotes = C.CPLES_XML_BUT_QUOTES
	// Scheme for CPLEscapeString()/CPLUnescapeString() for CSV (forced quoting)
	CPLESCSVForceQuoting = C.CPLES_CSV_FORCE_QUOTING
	// Scheme for CPLEscapeString()/CPLUnescapeString() for SQL identifiers
	CPLESSQLI = C.CPLES_SQLI
)

func cplEscapeString(s string, length, scheme int) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CPLEscapeString(cs, C.int(length), C.int(scheme))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplUnescapeString(s string, length *int, scheme int) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	var cLength C.int
	raw := C.CPLUnescapeString(cs, &cLength, C.int(scheme))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoStringN(raw, cLength)
	if length != nil {
		*length = int(cLength)
	}
	return
}

func cplBinaryToHex(data []byte) (result string) {
	raw := C.CPLBinaryToHex(C.int(len(data)), (*C.GByte)(cBytes(data)))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplHexToBinary(hex string, bytes *int) (result []byte) {
	cHex := C.CString(hex)
	defer C.free(unsafe.Pointer(cHex))
	var n C.int
	raw := C.CPLHexToBinary(cHex, &n)
	if raw == nil {
		return
	}
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoBytes(unsafe.Pointer(raw), n)
	if bytes != nil {
		*bytes = int(n)
	}
	return
}

func cplBase64Encode(data []byte) (result string) {
	raw := C.CPLBase64Encode(C.int(len(data)), (*C.GByte)(cBytes(data)))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplBase64DecodeInPlace(base64 string) (result []byte) {
	cs := C.CString(base64)
	defer C.free(unsafe.Pointer(cs))
	n := C.CPLBase64DecodeInPlace((*C.GByte)(unsafe.Pointer(cs)))
	result = C.GoBytes(unsafe.Pointer(cs), n)
	return
}

// Type of value
type CPLValueType C.CPLValueType

const (
	// String
	CPLValueString CPLValueType = C.CPL_VALUE_STRING
	// Real number
	CPLValueReal CPLValueType = C.CPL_VALUE_REAL
	// Integer
	CPLValueInteger CPLValueType = C.CPL_VALUE_INTEGER
)

func cplGetValueType(value string) (result CPLValueType) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = CPLValueType(C.CPLGetValueType(cValue))
	return
}

func cplToupper(c int) (result int) {
	result = int(C.CPLToupper(C.int(c)))
	return
}

func cplTolower(c int) (result int) {
	result = int(C.CPLTolower(C.int(c)))
	return
}

func cplStrlcpy(src string, destSize int) (result string, srcLen int) {
	cs := C.CString(src)
	defer C.free(unsafe.Pointer(cs))
	if destSize <= 0 {
		srcLen = int(C.CPLStrlcpy(nil, cs, 0))
		return
	}
	buf := make([]byte, destSize)
	srcLen = int(C.CPLStrlcpy((*C.char)(unsafe.Pointer(&buf[0])), cs, C.size_t(destSize)))
	result = C.GoString((*C.char)(unsafe.Pointer(&buf[0])))
	return
}

func cplStrlcat(dst, src string, destSize int) (result string, n int) {
	cSrc := C.CString(src)
	defer C.free(unsafe.Pointer(cSrc))
	if destSize <= 0 {
		n = int(C.CPLStrlcat(nil, cSrc, 0))
		return
	}
	buf := make([]byte, destSize+1)
	copy(buf[:destSize], dst)
	n = int(C.CPLStrlcat((*C.char)(unsafe.Pointer(&buf[0])), cSrc, C.size_t(destSize)))
	result = C.GoString((*C.char)(unsafe.Pointer(&buf[0])))
	return
}

func cplStrnlen(s string, maxLen int) (result uint64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = uint64(C.CPLStrnlen(cs, C.size_t(maxLen)))
	return
}

// The locale-independent formatting functions CPLvsnprintf, CPLsnprintf,
// CPLsprintf, CPLprintf, CPLsscanf, CPLSPrintf, CSLAppendPrintf and CPLVASPrintf
// are variadic / va_list based; they are deferred — format in Go instead.

const (
	// Encoding of the current locale
	CPLEncLocale = C.CPL_ENC_LOCALE
	// UTF-8 encoding
	CPLEncUTF8 = C.CPL_ENC_UTF8
	// UTF-16 encoding
	CPLEncUTF16 = C.CPL_ENC_UTF16
	// UCS-2 encoding
	CPLEncUCS2 = C.CPL_ENC_UCS2
	// UCS-4 encoding
	CPLEncUCS4 = C.CPL_ENC_UCS4
	// ASCII encoding
	CPLEncASCII = C.CPL_ENC_ASCII
	// ISO-8859-1 (LATIN1) encoding
	CPLEncISO8859_1 = C.CPL_ENC_ISO8859_1
)

func cplEncodingCharSize(encoding string) (result int) {
	cEncoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(cEncoding))
	result = int(C.CPLEncodingCharSize(cEncoding))
	return
}

func cplClearRecodeWarningFlags() {
	C.CPLClearRecodeWarningFlags()
}

func cplRecode(source, srcEncoding, dstEncoding string) (result string) {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))
	cSrcEncoding := C.CString(srcEncoding)
	defer C.free(unsafe.Pointer(cSrcEncoding))
	cDstEncoding := C.CString(dstEncoding)
	defer C.free(unsafe.Pointer(cDstEncoding))
	raw := C.CPLRecode(cSource, cSrcEncoding, cDstEncoding)
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

// CPLRecodeFromWChar and CPLRecodeToWChar operate on wchar_t*, whose width is
// platform-dependent and not cleanly representable in cgo; they are deferred.

func cplIsUTF8(data string, length int) (result int) {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))
	result = int(C.CPLIsUTF8(cData, C.int(length)))
	return
}

func cplIsASCII(data string, length int) (result bool) {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))
	result = bool(C.CPLIsASCII(cData, C.size_t(length)))
	return
}

func cplForceToASCII(data string, length int, replacement byte) (result string) {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))
	raw := C.CPLForceToASCII(cData, C.int(length), C.char(replacement))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplUTF8ForceToASCII(s string, replacement byte) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CPLUTF8ForceToASCII(cs, C.char(replacement))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplStrlenUTF8(s string) (result int) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.CPLStrlenUTF8(cs))
	return
}

func cplStrlenUTF8Ex(s string) (result uint64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = uint64(C.CPLStrlenUTF8Ex(cs))
	return
}

func cplCanRecode(testStr, srcEncoding, dstEncoding string) (result int) {
	cTestStr := C.CString(testStr)
	defer C.free(unsafe.Pointer(cTestStr))
	cSrcEncoding := C.CString(srcEncoding)
	defer C.free(unsafe.Pointer(cSrcEncoding))
	cDstEncoding := C.CString(dstEncoding)
	defer C.free(unsafe.Pointer(cDstEncoding))
	result = int(C.CPLCanRecode(cTestStr, cSrcEncoding, cDstEncoding))
	return
}

// Everything after CPL_C_END is extern "C++" and is skipped: CPLRemoveSQLComments,
// the CPLString class and CPLOPrintf/CPLOvPrintf/CPLQuotedSQLIdentifier,
// CPLURLGetValue/CPLURLAddKVP, the CPLStringList class, and the cpl:: namespace
// helpers (CSLUniquePtr, iterators, Iterate/IterateNameValue, ToVector, ...).
