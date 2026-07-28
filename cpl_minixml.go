package gdal

/*
#include "cpl_minixml_preamble.h"
*/
import "C"
import "unsafe"

// /** XML node type */
type CPLXMLNodeType C.CPLXMLNodeType

const (
	CXTElement   CPLXMLNodeType = C.CXT_Element
	CXTText      CPLXMLNodeType = C.CXT_Text
	CXTAttribute CPLXMLNodeType = C.CXT_Attribute
	CXTComment   CPLXMLNodeType = C.CXT_Comment
	CXTLiteral   CPLXMLNodeType = C.CXT_Literal
)

// Document node structure for a single parsed XML fragment.
// Allocated via CPL functions and freed with CPLDestroyXMLNode.
// Nodes form a tree via psChild and psNext links.
type CPLXMLNode struct {
	cValue *C.CPLXMLNode
}

func cplParseXMLString(xml string) (result CPLXMLNode) {
	cs := C.CString(xml)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLParseXMLString(cs)}
	return
}

func cplDestroyXMLNode(node CPLXMLNode) {
	C.CPLDestroyXMLNode(node.cValue)
}

func cplGetXMLNode(root CPLXMLNode, path string) (result CPLXMLNode) {
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLGetXMLNode(root.cValue, cs)}
	return
}

func cplSearchXMLNode(root CPLXMLNode, target string) (result CPLXMLNode) {
	cs := C.CString(target)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLSearchXMLNode(root.cValue, cs)}
	return
}

func cplGetXMLValue(root CPLXMLNode, path string, dflt string) (result string) {
	csPath := C.CString(path)
	defer C.free(unsafe.Pointer(csPath))
	csDflt := C.CString(dflt)
	defer C.free(unsafe.Pointer(csDflt))
	result = C.GoString(C.CPLGetXMLValue(root.cValue, csPath, csDflt))
	return
}

func cplCreateXMLNode(parent CPLXMLNode, eType CPLXMLNodeType, text string) (result CPLXMLNode) {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLCreateXMLNode(parent.cValue, C.CPLXMLNodeType(eType), cs)}
	return
}

func cplSerializeXMLTree(node CPLXMLNode) (result string) {
	raw := C.CPLSerializeXMLTree(node.cValue)
	defer vsiFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func cplAddXMLChild(parent CPLXMLNode, child CPLXMLNode) {
	C.CPLAddXMLChild(parent.cValue, child.cValue)
}

func cplRemoveXMLChild(parent CPLXMLNode, child CPLXMLNode) (result int) {
	result = int(C.CPLRemoveXMLChild(parent.cValue, child.cValue))
	return
}

func cplAddXMLSibling(olderSibling CPLXMLNode, newSibling CPLXMLNode) {
	C.CPLAddXMLSibling(olderSibling.cValue, newSibling.cValue)
}

func cplCreateXMLElementAndValue(parent CPLXMLNode, name string, value string) (result CPLXMLNode) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))
	result = CPLXMLNode{cValue: C.CPLCreateXMLElementAndValue(parent.cValue, csName, csValue)}
	return
}

func cplAddXMLAttributeAndValue(parent CPLXMLNode, name string, value string) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))
	C.CPLAddXMLAttributeAndValue(parent.cValue, csName, csValue)
}

func cplCloneXMLTree(tree CPLXMLNode) (result CPLXMLNode) {
	result = CPLXMLNode{cValue: C.CPLCloneXMLTree(tree.cValue)}
	return
}

func cplSetXMLValue(root CPLXMLNode, path string, value string) (result int) {
	csPath := C.CString(path)
	defer C.free(unsafe.Pointer(csPath))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))
	result = int(C.CPLSetXMLValue(root.cValue, csPath, csValue))
	return
}

func cplStripXMLNamespace(root CPLXMLNode, namespace string, recurse int) {
	cs := C.CString(namespace)
	defer C.free(unsafe.Pointer(cs))
	C.CPLStripXMLNamespace(root.cValue, cs, C.int(recurse))
}

func cplCleanXMLElementName(name string) (result string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.CPLCleanXMLElementName(cs)
	result = C.GoString(cs)
	return
}

func cplParseXMLFile(filename string) (result CPLXMLNode) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLParseXMLFile(cs)}
	return
}

func cplSerializeXMLTreeToFile(tree CPLXMLNode, filename string) (result int) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.CPLSerializeXMLTreeToFile(tree.cValue, cs))
	return
}

func cplXMLNodeGetRAMUsageEstimate(node CPLXMLNode) (result uint64) {
	result = uint64(C.CPLXMLNodeGetRAMUsageEstimate(node.cValue))
	return
}
