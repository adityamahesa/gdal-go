package gdal

// AddString appends a copy of s and returns the (possibly reallocated) owned
// list. The receiver must be an owned list (built from NullCSLConstList or a
// prior mutator); the returned list must be released with Destroy.
func (l CSLConstList) AddString(s string) CSLConstList {
	return cslAddString(l, s)
}

// AddStringMayFail is like AddString but returns a nil list on allocation
// failure instead of aborting.
func (l CSLConstList) AddStringMayFail(s string) CSLConstList {
	return cslAddStringMayFail(l, s)
}

// Count returns the number of strings in the list.
func (l CSLConstList) Count() int {
	return cslCount(l)
}

// GetField returns the string at index i, or "" if i is out of range.
func (l CSLConstList) GetField(i int) string {
	return cslGetField(l, i)
}

// Destroy frees an owned string list. Do not call it on a borrowed list.
func (l CSLConstList) Destroy() {
	cslDestroy(l)
}

// Duplicate returns a newly allocated deep copy of the list. The returned list
// must be released with Destroy.
func (l CSLConstList) Duplicate() CSLConstList {
	return cslDuplicate(l)
}

// Merge folds the name/value pairs from override into the receiver (an owned
// list) and returns it. The returned list must be released with Destroy.
func (l CSLConstList) Merge(override CSLConstList) CSLConstList {
	return cslMerge(l, override)
}

// CSLTokenizeString splits s on whitespace into an owned list that must be
// released with Destroy.
func CSLTokenizeString(s string) CSLConstList {
	return cslTokenizeString(s)
}

// CSLTokenizeStringComplex splits s on any character in delimiter. The returned
// owned list must be released with Destroy.
func CSLTokenizeStringComplex(s, delimiter string, honourStrings, allowEmptyTokens int) CSLConstList {
	return cslTokenizeStringComplex(s, delimiter, honourStrings, allowEmptyTokens)
}

// CSLTokenizeString2 splits s on any character in delimiter, controlled by the
// CSLT_* flags. The returned owned list must be released with Destroy.
func CSLTokenizeString2(s, delimiter string, flags int) CSLConstList {
	return cslTokenizeString2(s, delimiter, flags)
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

// CSLLoad reads fname into an owned list, one line per string. The returned list
// must be released with Destroy.
func CSLLoad(fname string) CSLConstList {
	return cslLoad(fname)
}

// CSLLoad2 reads fname with limits and options into an owned list. The returned
// list must be released with Destroy.
func CSLLoad2(fname string, maxLines, maxCols int, options CSLConstList) CSLConstList {
	return cslLoad2(fname, maxLines, maxCols, options)
}

// Save writes the list to fname, one string per line, and returns the number of
// lines written (0 on failure).
func (l CSLConstList) Save(fname string) int {
	return cslSave(l, fname)
}

// InsertStrings inserts newLines at index at (a negative at appends) and returns
// the (possibly reallocated) owned list, which must be released with Destroy.
func (l CSLConstList) InsertStrings(at int, newLines CSLConstList) CSLConstList {
	return cslInsertStrings(l, at, newLines)
}

// InsertString inserts a copy of s at index at (a negative at appends) and
// returns the (possibly reallocated) owned list, released with Destroy.
func (l CSLConstList) InsertString(at int, s string) CSLConstList {
	return cslInsertString(l, at, s)
}

// RemoveStrings removes num strings starting at index first. It returns the
// (possibly reallocated) owned list and, as the second value, an owned list of
// the removed strings. Both returned lists must be released with Destroy.
func (l CSLConstList) RemoveStrings(first, num int) (result CSLConstList, removed CSLConstList) {
	result = cslRemoveStrings(l, first, num, &removed)
	return
}

// FindString returns the index of target (case-insensitive), or -1.
func (l CSLConstList) FindString(target string) int {
	return cslFindString(l, target)
}

// FindStringCaseSensitive returns the index of target (case-sensitive), or -1.
func (l CSLConstList) FindStringCaseSensitive(target string) int {
	return cslFindStringCaseSensitive(l, target)
}

// PartialFindString returns the index of the first string containing needle as a
// substring, or -1.
func (l CSLConstList) PartialFindString(needle string) int {
	return cslPartialFindString(l, needle)
}

// FindName returns the index of the "name=value" entry with the given name, or
// -1.
func (l CSLConstList) FindName(name string) int {
	return cslFindName(l, name)
}

// FetchBoolean returns the boolean value associated with key, or dflt if absent.
func (l CSLConstList) FetchBoolean(key string, dflt int) int {
	return cslFetchBoolean(l, key, dflt)
}

// CSLTestBoolean interprets value ("YES"/"NO"/"1"/"0"/...) as a boolean.
//
// Deprecated: use CPLTestBool.
func CSLTestBoolean(value string) int {
	return cslTestBoolean(value)
}

// CPLTestBoolean interprets value as a boolean, returning 0 or 1.
func CPLTestBoolean(value string) int {
	return cplTestBoolean(value)
}

// CPLTestBool interprets value as a boolean.
func CPLTestBool(value string) bool {
	return cplTestBool(value)
}

// FetchBool returns the boolean value associated with key, or dflt if absent.
func (l CSLConstList) FetchBool(key string, dflt bool) bool {
	return cplFetchBool(l, key, dflt)
}

// CPLParseMemorySize parses a memory size such as "10MB" into a byte count.
// unitSpecified reports whether value carried an explicit unit.
func CPLParseMemorySize(value string) (nValue GIntBig, unitSpecified bool, err error) {
	err = cplErr(cplParseMemorySize(value, &nValue, &unitSpecified))
	return
}

// CPLParseNameValue splits a "name=value" (or "name:value") string, returning
// the value and the name (key).
func CPLParseNameValue(nameValue string) (value string, key string) {
	value = cplParseNameValue(nameValue, &key)
	return
}

// CPLParseNameValueSep splits a name/value string on the given separator
// character, returning the value and the name (key).
func CPLParseNameValueSep(nameValue string, sep byte) (value string, key string) {
	value = cplParseNameValueSep(nameValue, &key, sep)
	return
}

// FetchNameValue returns the value of the "name=value" entry with the given
// name, or "" if absent.
func (l CSLConstList) FetchNameValue(name string) string {
	return cslFetchNameValue(l, name)
}

// FetchNameValueDef returns the value associated with name, or dflt if absent.
func (l CSLConstList) FetchNameValueDef(name, dflt string) string {
	return cslFetchNameValueDef(l, name, dflt)
}

// FetchNameValueMultiple returns an owned list of all values associated with
// name. The returned list must be released with Destroy.
func (l CSLConstList) FetchNameValueMultiple(name string) CSLConstList {
	return cslFetchNameValueMultiple(l, name)
}

// AddNameValue appends a "name=value" entry (allowing duplicate names) and
// returns the owned list, which must be released with Destroy.
func (l CSLConstList) AddNameValue(name, value string) CSLConstList {
	return cslAddNameValue(l, name, value)
}

// SetNameValue sets (replacing any existing) the "name=value" entry and returns
// the owned list, which must be released with Destroy.
func (l CSLConstList) SetNameValue(name, value string) CSLConstList {
	return cslSetNameValue(l, name, value)
}

// SetNameValueSeparator rewrites each entry to use separator between name and
// value, in place.
func (l CSLConstList) SetNameValueSeparator(separator string) {
	cslSetNameValueSeparator(l, separator)
}

// CSLParseCommandLine tokenizes a command line into an owned argument list,
// honouring quoting. The returned list must be released with Destroy.
func CSLParseCommandLine(commandLine string) CSLConstList {
	return cslParseCommandLine(commandLine)
}

// CPLEscapeString escapes s according to scheme (a CPLES_* constant). A length
// of -1 uses the NUL-terminated length of s.
func CPLEscapeString(s string, length, scheme int) string {
	return cplEscapeString(s, length, scheme)
}

// CPLUnescapeString reverses CPLEscapeString for scheme, returning the decoded
// bytes (which may contain NULs) and their length.
func CPLUnescapeString(s string, scheme int) (result string, length int) {
	result = cplUnescapeString(s, &length, scheme)
	return
}

// CPLBinaryToHex returns the hexadecimal encoding of data.
func CPLBinaryToHex(data []byte) string {
	return cplBinaryToHex(data)
}

// CPLHexToBinary decodes a hexadecimal string into bytes.
func CPLHexToBinary(hex string) (result []byte) {
	result = cplHexToBinary(hex, nil)
	return
}

// CPLBase64Encode returns the base64 encoding of data.
func CPLBase64Encode(data []byte) string {
	return cplBase64Encode(data)
}

// CPLBase64DecodeInPlace decodes a base64 string and returns the decoded bytes.
func CPLBase64DecodeInPlace(base64 string) []byte {
	return cplBase64DecodeInPlace(base64)
}

// CPLGetValueType classifies value as a string, real, or integer.
func CPLGetValueType(value string) CPLValueType {
	return cplGetValueType(value)
}

// CPLToupper is a locale-independent toupper().
func CPLToupper(c int) int {
	return cplToupper(c)
}

// CPLTolower is a locale-independent tolower().
func CPLTolower(c int) int {
	return cplTolower(c)
}

// CPLStrlcpy copies src into a buffer of destSize bytes and returns the copied
// (possibly truncated) string together with the length of src.
func CPLStrlcpy(src string, destSize int) (result string, srcLen int) {
	return cplStrlcpy(src, destSize)
}

// CPLStrlcat appends src to dst within a buffer of destSize bytes and returns
// the resulting (possibly truncated) string and the length it tried to build.
func CPLStrlcat(dst, src string, destSize int) (result string, n int) {
	return cplStrlcat(dst, src, destSize)
}

// CPLStrnlen returns the length of s, capped at maxLen.
func CPLStrnlen(s string, maxLen int) uint64 {
	return cplStrnlen(s, maxLen)
}

// CPLEncodingCharSize returns the character size (in bytes) of the encoding, or
// -1 if unknown.
func CPLEncodingCharSize(encoding string) int {
	return cplEncodingCharSize(encoding)
}

func CPLClearRecodeWarningFlags() {
	cplClearRecodeWarningFlags()
}

// CPLRecode converts source from srcEncoding to dstEncoding (CPL_ENC_* names).
func CPLRecode(source, srcEncoding, dstEncoding string) string {
	return cplRecode(source, srcEncoding, dstEncoding)
}

// CPLIsUTF8 reports (as 0/1) whether the first length bytes of data are valid
// UTF-8. A length of -1 uses the NUL-terminated length.
func CPLIsUTF8(data string, length int) int {
	return cplIsUTF8(data, length)
}

// CPLIsASCII reports whether the first length bytes of data are pure ASCII.
func CPLIsASCII(data string, length int) bool {
	return cplIsASCII(data, length)
}

// CPLForceToASCII replaces every non-ASCII byte in the first length bytes of
// data with replacement. A length of -1 uses the NUL-terminated length.
func CPLForceToASCII(data string, length int, replacement byte) string {
	return cplForceToASCII(data, length, replacement)
}

// CPLUTF8ForceToASCII transliterates s to ASCII, replacing untranslatable
// characters with replacement.
func CPLUTF8ForceToASCII(s string, replacement byte) string {
	return cplUTF8ForceToASCII(s, replacement)
}

// CPLStrlenUTF8 returns the number of UTF-8 characters in s.
//
// Deprecated: use CPLStrlenUTF8Ex.
func CPLStrlenUTF8(s string) int {
	return cplStrlenUTF8(s)
}

// CPLStrlenUTF8Ex returns the number of UTF-8 characters in s.
func CPLStrlenUTF8Ex(s string) uint64 {
	return cplStrlenUTF8Ex(s)
}

// CPLCanRecode reports (as 0/1) whether testStr can be recoded from srcEncoding
// to dstEncoding.
func CPLCanRecode(testStr, srcEncoding, dstEncoding string) int {
	return cplCanRecode(testStr, srcEncoding, dstEncoding)
}
