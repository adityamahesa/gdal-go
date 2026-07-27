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

// AddString appends a copy of s and returns the (possibly reallocated) owned
// list. The receiver must be an owned list (built from NullCSLConstList or a
// prior mutator); the returned list must be released with Destroy.
func (l CSLConstList) AddString(s string) CSLConstList {
	return cslAddString(l, s)
}

func cslAddStringMayFail(l CSLConstList, s string) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CSLAddStringMayFail(l.cValue, cs)
	result = cslConstList(raw)
	return
}

// AddStringMayFail is like AddString but returns a nil list on allocation
// failure instead of aborting.
func (l CSLConstList) AddStringMayFail(s string) CSLConstList {
	return cslAddStringMayFail(l, s)
}

func cslCount(l CSLConstList) (result int) {
	result = int(C.CSLCount(l.cValue))
	return
}

// Count returns the number of strings in the list.
func (l CSLConstList) Count() int {
	return cslCount(l)
}

func cslGetField(l CSLConstList, i int) (result string) {
	result = C.GoString(C.CSLGetField(l.cValue, C.int(i)))
	return
}

// GetField returns the string at index i, or "" if i is out of range.
func (l CSLConstList) GetField(i int) string {
	return cslGetField(l, i)
}

func cslDestroy(l CSLConstList) {
	C.CSLDestroy(l.cValue)
}

// Destroy frees an owned string list. Do not call it on a borrowed list.
func (l CSLConstList) Destroy() {
	cslDestroy(l)
}

func cslDuplicate(l CSLConstList) (result CSLConstList) {
	raw := C.CSLDuplicate(l.cValue)
	result = cslConstList(raw)
	return
}

// Duplicate returns a newly allocated deep copy of the list. The returned list
// must be released with Destroy.
func (l CSLConstList) Duplicate() CSLConstList {
	return cslDuplicate(l)
}

func cslMerge(l CSLConstList, override CSLConstList) (result CSLConstList) {
	raw := C.CSLMerge(l.cValue, override.cValue)
	result = cslConstList(raw)
	return
}

// Merge folds the name/value pairs from override into the receiver (an owned
// list) and returns it. The returned list must be released with Destroy.
func (l CSLConstList) Merge(override CSLConstList) CSLConstList {
	return cslMerge(l, override)
}

func cslTokenizeString(s string) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CSLTokenizeString(cs)
	result = cslConstList(raw)
	return
}

// CSLTokenizeString splits s on whitespace into an owned list that must be
// released with Destroy.
func CSLTokenizeString(s string) CSLConstList {
	return cslTokenizeString(s)
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

// CSLTokenizeStringComplex splits s on any character in delimiter. The returned
// owned list must be released with Destroy.
func CSLTokenizeStringComplex(s, delimiter string, honourStrings, allowEmptyTokens int) CSLConstList {
	return cslTokenizeStringComplex(s, delimiter, honourStrings, allowEmptyTokens)
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

// CSLTokenizeString2 splits s on any character in delimiter, controlled by the
// CSLT_* flags. The returned owned list must be released with Destroy.
func CSLTokenizeString2(s, delimiter string, flags int) CSLConstList {
	return cslTokenizeString2(s, delimiter, flags)
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

// Print writes the list, one string per line, to filename (created/truncated).
// The C function's raw FILE* is not exposed; it prints to the opened file.
func (l CSLConstList) Print(filename string) (result int, err error) {
	fp, closeFn, err := cFOpen(filename, "w")
	if err != nil {
		return 0, err
	}
	defer closeFn()
	result = cslPrint(l, fp)
	return
}

func cslLoad(fname string) (result CSLConstList) {
	cFname := C.CString(fname)
	defer C.free(unsafe.Pointer(cFname))
	raw := C.CSLLoad(cFname)
	result = cslConstList(raw)
	return
}

// CSLLoad reads fname into an owned list, one line per string. The returned list
// must be released with Destroy.
func CSLLoad(fname string) CSLConstList {
	return cslLoad(fname)
}

func cslLoad2(fname string, maxLines, maxCols int, options CSLConstList) (result CSLConstList) {
	cFname := C.CString(fname)
	defer C.free(unsafe.Pointer(cFname))
	raw := C.CSLLoad2(cFname, C.int(maxLines), C.int(maxCols), options.cValue)
	result = cslConstList(raw)
	return
}

// CSLLoad2 reads fname with limits and options into an owned list. The returned
// list must be released with Destroy.
func CSLLoad2(fname string, maxLines, maxCols int, options CSLConstList) CSLConstList {
	return cslLoad2(fname, maxLines, maxCols, options)
}

func cslSave(l CSLConstList, fname string) (result int) {
	cFname := C.CString(fname)
	defer C.free(unsafe.Pointer(cFname))
	result = int(C.CSLSave(l.cValue, cFname))
	return
}

// Save writes the list to fname, one string per line, and returns the number of
// lines written (0 on failure).
func (l CSLConstList) Save(fname string) int {
	return cslSave(l, fname)
}

func cslInsertStrings(l CSLConstList, at int, newLines CSLConstList) (result CSLConstList) {
	raw := C.CSLInsertStrings(l.cValue, C.int(at), newLines.cValue)
	result = cslConstList(raw)
	return
}

// InsertStrings inserts newLines at index at (a negative at appends) and returns
// the (possibly reallocated) owned list, which must be released with Destroy.
func (l CSLConstList) InsertStrings(at int, newLines CSLConstList) CSLConstList {
	return cslInsertStrings(l, at, newLines)
}

func cslInsertString(l CSLConstList, at int, s string) (result CSLConstList) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CSLInsertString(l.cValue, C.int(at), cs)
	result = cslConstList(raw)
	return
}

// InsertString inserts a copy of s at index at (a negative at appends) and
// returns the (possibly reallocated) owned list, released with Destroy.
func (l CSLConstList) InsertString(at int, s string) CSLConstList {
	return cslInsertString(l, at, s)
}

func cslRemoveStrings(l CSLConstList, first, num int, retStrings ***C.char) (result CSLConstList) {
	raw := C.CSLRemoveStrings(l.cValue, C.int(first), C.int(num), retStrings)
	result = cslConstList(raw)
	return
}

// RemoveStrings removes num strings starting at index first. It returns the
// (possibly reallocated) owned list and, as the second value, an owned list of
// the removed strings. Both returned lists must be released with Destroy.
func (l CSLConstList) RemoveStrings(first, num int) (result CSLConstList, removed CSLConstList) {
	var ret **C.char
	result = cslRemoveStrings(l, first, num, &ret)
	removed = cslConstList(ret)
	return
}

func cslFindString(l CSLConstList, target string) (result int) {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	result = int(C.CSLFindString(l.cValue, cTarget))
	return
}

// FindString returns the index of target (case-insensitive), or -1.
func (l CSLConstList) FindString(target string) int {
	return cslFindString(l, target)
}

func cslFindStringCaseSensitive(l CSLConstList, target string) (result int) {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	result = int(C.CSLFindStringCaseSensitive(l.cValue, cTarget))
	return
}

// FindStringCaseSensitive returns the index of target (case-sensitive), or -1.
func (l CSLConstList) FindStringCaseSensitive(target string) int {
	return cslFindStringCaseSensitive(l, target)
}

func cslPartialFindString(l CSLConstList, needle string) (result int) {
	cNeedle := C.CString(needle)
	defer C.free(unsafe.Pointer(cNeedle))
	result = int(C.CSLPartialFindString(l.cValue, cNeedle))
	return
}

// PartialFindString returns the index of the first string containing needle as a
// substring, or -1.
func (l CSLConstList) PartialFindString(needle string) int {
	return cslPartialFindString(l, needle)
}

func cslFindName(l CSLConstList, name string) (result int) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = int(C.CSLFindName(l.cValue, cName))
	return
}

// FindName returns the index of the "name=value" entry with the given name, or
// -1.
func (l CSLConstList) FindName(name string) int {
	return cslFindName(l, name)
}

func cslFetchBoolean(l CSLConstList, key string, dflt int) (result int) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	result = int(C.CSLFetchBoolean(l.cValue, cKey, C.int(dflt)))
	return
}

// FetchBoolean returns the boolean value associated with key, or dflt if absent.
func (l CSLConstList) FetchBoolean(key string, dflt int) int {
	return cslFetchBoolean(l, key, dflt)
}

func cslTestBoolean(value string) (result int) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = int(C.CSLTestBoolean(cValue))
	return
}

// CSLTestBoolean interprets value ("YES"/"NO"/"1"/"0"/...) as a boolean.
//
// Deprecated: use CPLTestBool.
func CSLTestBoolean(value string) int {
	return cslTestBoolean(value)
}

func cplTestBoolean(value string) (result int) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = int(C.CPLTestBoolean(cValue))
	return
}

// CPLTestBoolean interprets value as a boolean, returning 0 or 1.
func CPLTestBoolean(value string) int {
	return cplTestBoolean(value)
}

func cplTestBool(value string) (result bool) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	result = bool(C.CPLTestBool(cValue))
	return
}

// CPLTestBool interprets value as a boolean.
func CPLTestBool(value string) bool {
	return cplTestBool(value)
}

func cplFetchBool(l CSLConstList, key string, dflt bool) (result bool) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	result = bool(C.CPLFetchBool(l.cValue, cKey, C.bool(dflt)))
	return
}

// FetchBool returns the boolean value associated with key, or dflt if absent.
func (l CSLConstList) FetchBool(key string, dflt bool) bool {
	return cplFetchBool(l, key, dflt)
}

func cplParseMemorySize(value string, nValue *GIntBig, unitSpecified *bool) (ret CPLErr) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	ret = CPLErr(C.CPLParseMemorySize(cValue, (*C.GIntBig)(nValue), (*C.bool)(unsafe.Pointer(unitSpecified))))
	return
}

// CPLParseMemorySize parses a memory size such as "10MB" into a byte count.
// unitSpecified reports whether value carried an explicit unit.
func CPLParseMemorySize(value string) (nValue GIntBig, unitSpecified bool, err error) {
	err = cplErr(cplParseMemorySize(value, &nValue, &unitSpecified))
	return
}

func cplParseNameValue(nameValue string, key **C.char) (result string) {
	cNameValue := C.CString(nameValue)
	defer C.free(unsafe.Pointer(cNameValue))
	result = C.GoString(C.CPLParseNameValue(cNameValue, key))
	return
}

// CPLParseNameValue splits a "name=value" (or "name:value") string, returning
// the value and the name (key).
func CPLParseNameValue(nameValue string) (value string, key string) {
	var cKey *C.char
	value = cplParseNameValue(nameValue, &cKey)
	if cKey != nil {
		key = C.GoString(cKey)
		C.VSIFree(unsafe.Pointer(cKey))
	}
	return
}

func cplParseNameValueSep(nameValue string, key **C.char, sep byte) (result string) {
	cNameValue := C.CString(nameValue)
	defer C.free(unsafe.Pointer(cNameValue))
	result = C.GoString(C.CPLParseNameValueSep(cNameValue, key, C.char(sep)))
	return
}

// CPLParseNameValueSep splits a name/value string on the given separator
// character, returning the value and the name (key).
func CPLParseNameValueSep(nameValue string, sep byte) (value string, key string) {
	var cKey *C.char
	value = cplParseNameValueSep(nameValue, &cKey, sep)
	if cKey != nil {
		key = C.GoString(cKey)
		C.VSIFree(unsafe.Pointer(cKey))
	}
	return
}

func cslFetchNameValue(l CSLConstList, name string) (result string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	result = C.GoString(C.CSLFetchNameValue(l.cValue, cName))
	return
}

// FetchNameValue returns the value of the "name=value" entry with the given
// name, or "" if absent.
func (l CSLConstList) FetchNameValue(name string) string {
	return cslFetchNameValue(l, name)
}

func cslFetchNameValueDef(l CSLConstList, name, dflt string) (result string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(cDflt))
	result = C.GoString(C.CSLFetchNameValueDef(l.cValue, cName, cDflt))
	return
}

// FetchNameValueDef returns the value associated with name, or dflt if absent.
func (l CSLConstList) FetchNameValueDef(name, dflt string) string {
	return cslFetchNameValueDef(l, name, dflt)
}

func cslFetchNameValueMultiple(l CSLConstList, name string) (result CSLConstList) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	raw := C.CSLFetchNameValueMultiple(l.cValue, cName)
	result = cslConstList(raw)
	return
}

// FetchNameValueMultiple returns an owned list of all values associated with
// name. The returned list must be released with Destroy.
func (l CSLConstList) FetchNameValueMultiple(name string) CSLConstList {
	return cslFetchNameValueMultiple(l, name)
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

// AddNameValue appends a "name=value" entry (allowing duplicate names) and
// returns the owned list, which must be released with Destroy.
func (l CSLConstList) AddNameValue(name, value string) CSLConstList {
	return cslAddNameValue(l, name, value)
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

// SetNameValue sets (replacing any existing) the "name=value" entry and returns
// the owned list, which must be released with Destroy.
func (l CSLConstList) SetNameValue(name, value string) CSLConstList {
	return cslSetNameValue(l, name, value)
}

func cslSetNameValueSeparator(l CSLConstList, separator string) {
	cSeparator := C.CString(separator)
	defer C.free(unsafe.Pointer(cSeparator))
	C.CSLSetNameValueSeparator(l.cValue, cSeparator)
}

// SetNameValueSeparator rewrites each entry to use separator between name and
// value, in place.
func (l CSLConstList) SetNameValueSeparator(separator string) {
	cslSetNameValueSeparator(l, separator)
}

func cslParseCommandLine(commandLine string) (result CSLConstList) {
	cCommandLine := C.CString(commandLine)
	defer C.free(unsafe.Pointer(cCommandLine))
	raw := C.CSLParseCommandLine(cCommandLine)
	result = cslConstList(raw)
	return
}

// CSLParseCommandLine tokenizes a command line into an owned argument list,
// honouring quoting. The returned list must be released with Destroy.
func CSLParseCommandLine(commandLine string) CSLConstList {
	return cslParseCommandLine(commandLine)
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

// CPLEscapeString escapes s according to scheme (a CPLES_* constant). A length
// of -1 uses the NUL-terminated length of s.
func CPLEscapeString(s string, length, scheme int) string {
	return cplEscapeString(s, length, scheme)
}

func cplUnescapeString(s string, length *C.int, scheme int) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CPLUnescapeString(cs, length, C.int(scheme))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoStringN(raw, *length)
	return
}

// CPLUnescapeString reverses CPLEscapeString for scheme, returning the decoded
// bytes (which may contain NULs) and their length.
func CPLUnescapeString(s string, scheme int) (result string, length int) {
	var cLength C.int
	result = cplUnescapeString(s, &cLength, scheme)
	length = int(cLength)
	return
}

func cplBinaryToHex(data []byte) (result string) {
	raw := C.CPLBinaryToHex(C.int(len(data)), (*C.GByte)(cBytes(data)))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

// CPLBinaryToHex returns the hexadecimal encoding of data.
func CPLBinaryToHex(data []byte) string {
	return cplBinaryToHex(data)
}

func cplHexToBinary(hex string, bytes *C.int) (result []byte) {
	cHex := C.CString(hex)
	defer C.free(unsafe.Pointer(cHex))
	raw := C.CPLHexToBinary(cHex, bytes)
	if raw == nil {
		return
	}
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoBytes(unsafe.Pointer(raw), *bytes)
	return
}

// CPLHexToBinary decodes a hexadecimal string into bytes.
func CPLHexToBinary(hex string) (result []byte) {
	var n C.int
	result = cplHexToBinary(hex, &n)
	return
}

func cplBase64Encode(data []byte) (result string) {
	raw := C.CPLBase64Encode(C.int(len(data)), (*C.GByte)(cBytes(data)))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

// CPLBase64Encode returns the base64 encoding of data.
func CPLBase64Encode(data []byte) string {
	return cplBase64Encode(data)
}

func cplBase64DecodeInPlace(base64 string) (result []byte) {
	cs := C.CString(base64)
	defer C.free(unsafe.Pointer(cs))
	n := C.CPLBase64DecodeInPlace((*C.GByte)(unsafe.Pointer(cs)))
	result = C.GoBytes(unsafe.Pointer(cs), n)
	return
}

// CPLBase64DecodeInPlace decodes a base64 string and returns the decoded bytes.
func CPLBase64DecodeInPlace(base64 string) []byte {
	return cplBase64DecodeInPlace(base64)
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

// CPLGetValueType classifies value as a string, real, or integer.
func CPLGetValueType(value string) CPLValueType {
	return cplGetValueType(value)
}

func cplToupper(c int) (result int) {
	result = int(C.CPLToupper(C.int(c)))
	return
}

// CPLToupper is a locale-independent toupper().
func CPLToupper(c int) int {
	return cplToupper(c)
}

func cplTolower(c int) (result int) {
	result = int(C.CPLTolower(C.int(c)))
	return
}

// CPLTolower is a locale-independent tolower().
func CPLTolower(c int) int {
	return cplTolower(c)
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

// CPLStrlcpy copies src into a buffer of destSize bytes and returns the copied
// (possibly truncated) string together with the length of src.
func CPLStrlcpy(src string, destSize int) (result string, srcLen int) {
	return cplStrlcpy(src, destSize)
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

// CPLStrlcat appends src to dst within a buffer of destSize bytes and returns
// the resulting (possibly truncated) string and the length it tried to build.
func CPLStrlcat(dst, src string, destSize int) (result string, n int) {
	return cplStrlcat(dst, src, destSize)
}

func cplStrnlen(s string, maxLen int) (result uint64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = uint64(C.CPLStrnlen(cs, C.size_t(maxLen)))
	return
}

// CPLStrnlen returns the length of s, capped at maxLen.
func CPLStrnlen(s string, maxLen int) uint64 {
	return cplStrnlen(s, maxLen)
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

// CPLEncodingCharSize returns the character size (in bytes) of the encoding, or
// -1 if unknown.
func CPLEncodingCharSize(encoding string) int {
	return cplEncodingCharSize(encoding)
}

func cplClearRecodeWarningFlags() {
	C.CPLClearRecodeWarningFlags()
}

func CPLClearRecodeWarningFlags() {
	cplClearRecodeWarningFlags()
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

// CPLRecode converts source from srcEncoding to dstEncoding (CPL_ENC_* names).
func CPLRecode(source, srcEncoding, dstEncoding string) string {
	return cplRecode(source, srcEncoding, dstEncoding)
}

// CPLRecodeFromWChar and CPLRecodeToWChar operate on wchar_t*, whose width is
// platform-dependent and not cleanly representable in cgo; they are deferred.

func cplIsUTF8(data string, length int) (result int) {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))
	result = int(C.CPLIsUTF8(cData, C.int(length)))
	return
}

// CPLIsUTF8 reports (as 0/1) whether the first length bytes of data are valid
// UTF-8. A length of -1 uses the NUL-terminated length.
func CPLIsUTF8(data string, length int) int {
	return cplIsUTF8(data, length)
}

func cplIsASCII(data string, length int) (result bool) {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))
	result = bool(C.CPLIsASCII(cData, C.size_t(length)))
	return
}

// CPLIsASCII reports whether the first length bytes of data are pure ASCII.
func CPLIsASCII(data string, length int) bool {
	return cplIsASCII(data, length)
}

func cplForceToASCII(data string, length int, replacement byte) (result string) {
	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))
	raw := C.CPLForceToASCII(cData, C.int(length), C.char(replacement))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

// CPLForceToASCII replaces every non-ASCII byte in the first length bytes of
// data with replacement. A length of -1 uses the NUL-terminated length.
func CPLForceToASCII(data string, length int, replacement byte) string {
	return cplForceToASCII(data, length, replacement)
}

func cplUTF8ForceToASCII(s string, replacement byte) (result string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	raw := C.CPLUTF8ForceToASCII(cs, C.char(replacement))
	defer C.VSIFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

// CPLUTF8ForceToASCII transliterates s to ASCII, replacing untranslatable
// characters with replacement.
func CPLUTF8ForceToASCII(s string, replacement byte) string {
	return cplUTF8ForceToASCII(s, replacement)
}

func cplStrlenUTF8(s string) (result int) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.CPLStrlenUTF8(cs))
	return
}

// CPLStrlenUTF8 returns the number of UTF-8 characters in s.
//
// Deprecated: use CPLStrlenUTF8Ex.
func CPLStrlenUTF8(s string) int {
	return cplStrlenUTF8(s)
}

func cplStrlenUTF8Ex(s string) (result uint64) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	result = uint64(C.CPLStrlenUTF8Ex(cs))
	return
}

// CPLStrlenUTF8Ex returns the number of UTF-8 characters in s.
func CPLStrlenUTF8Ex(s string) uint64 {
	return cplStrlenUTF8Ex(s)
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

// CPLCanRecode reports (as 0/1) whether testStr can be recoded from srcEncoding
// to dstEncoding.
func CPLCanRecode(testStr, srcEncoding, dstEncoding string) int {
	return cplCanRecode(testStr, srcEncoding, dstEncoding)
}

// Everything after CPL_C_END is extern "C++" and is skipped: CPLRemoveSQLComments,
// the CPLString class and CPLOPrintf/CPLOvPrintf/CPLQuotedSQLIdentifier,
// CPLURLGetValue/CPLURLAddKVP, the CPLStringList class, and the cpl:: namespace
// helpers (CSLUniquePtr, iterators, Iterate/IterateNameValue, ToVector, ...).
