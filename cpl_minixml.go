package gdal

/*
#include "cpl_minixml.h"
#include "cpl_conv.h" // TODO: implement cpl_conv.go
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

func CPLParseXMLString(xml string) (result CPLXMLNode, err error) {
	result = cplParseXMLString(xml)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func cplDestroyXMLNode(node CPLXMLNode) {
	C.CPLDestroyXMLNode(node.cValue)
}

func (n CPLXMLNode) Destroy() {
	cplDestroyXMLNode(n)
}

func cplGetXMLNode(root CPLXMLNode, path string) (result CPLXMLNode) {
	cs := C.CString(path)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLGetXMLNode(root.cValue, cs)}
	return
}

func (n CPLXMLNode) GetNode(path string) (result CPLXMLNode, err error) {
	result = cplGetXMLNode(n, path)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func cplSearchXMLNode(root CPLXMLNode, target string) (result CPLXMLNode) {
	cs := C.CString(target)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLSearchXMLNode(root.cValue, cs)}
	return
}

func (n CPLXMLNode) SearchNode(target string) (result CPLXMLNode, err error) {
	result = cplSearchXMLNode(n, target)
	if result.cValue == nil {
		err = lastError()
	}
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

func (n CPLXMLNode) GetValue(path, dflt string) string {
	return cplGetXMLValue(n, path, dflt)
}

func cplCreateXMLNode(parent CPLXMLNode, eType CPLXMLNodeType, text string) (result CPLXMLNode) {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLCreateXMLNode(parent.cValue, C.CPLXMLNodeType(eType), cs)}
	return
}

func (n CPLXMLNode) CreateNode(eType CPLXMLNodeType, text string) (result CPLXMLNode, err error) {
	result = cplCreateXMLNode(n, eType, text)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func cplSerializeXMLTree(node CPLXMLNode) (result string) {
	raw := C.CPLSerializeXMLTree(node.cValue)
	defer C.CPLFree(unsafe.Pointer(raw))
	result = C.GoString(raw)
	return
}

func (n CPLXMLNode) SerializeTree() string {
	return cplSerializeXMLTree(n)
}

func cplAddXMLChild(parent CPLXMLNode, child CPLXMLNode) {
	C.CPLAddXMLChild(parent.cValue, child.cValue)
}

func (n CPLXMLNode) AddChild(child CPLXMLNode) {
	cplAddXMLChild(n, child)
}

func cplRemoveXMLChild(parent CPLXMLNode, child CPLXMLNode) (result int) {
	result = int(C.CPLRemoveXMLChild(parent.cValue, child.cValue))
	return
}

func (n CPLXMLNode) RemoveChild(child CPLXMLNode) int {
	return cplRemoveXMLChild(n, child)
}

func cplAddXMLSibling(olderSibling CPLXMLNode, newSibling CPLXMLNode) {
	C.CPLAddXMLSibling(olderSibling.cValue, newSibling.cValue)
}

func (n CPLXMLNode) AddSibling(newSibling CPLXMLNode) {
	cplAddXMLSibling(n, newSibling)
}

func cplCreateXMLElementAndValue(parent CPLXMLNode, name string, value string) (result CPLXMLNode) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))
	result = CPLXMLNode{cValue: C.CPLCreateXMLElementAndValue(parent.cValue, csName, csValue)}
	return
}

func (n CPLXMLNode) CreateElementAndValue(name, value string) (result CPLXMLNode, err error) {
	result = cplCreateXMLElementAndValue(n, name, value)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func cplAddXMLAttributeAndValue(parent CPLXMLNode, name string, value string) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))
	C.CPLAddXMLAttributeAndValue(parent.cValue, csName, csValue)
}

func (n CPLXMLNode) AddAttributeAndValue(name, value string) {
	cplAddXMLAttributeAndValue(n, name, value)
}

func cplCloneXMLTree(tree CPLXMLNode) (result CPLXMLNode) {
	result = CPLXMLNode{cValue: C.CPLCloneXMLTree(tree.cValue)}
	return
}

func (n CPLXMLNode) CloneTree() (result CPLXMLNode, err error) {
	result = cplCloneXMLTree(n)
	if result.cValue == nil {
		err = lastError()
	}
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

func (n CPLXMLNode) SetValue(path, value string) int {
	return cplSetXMLValue(n, path, value)
}

func cplStripXMLNamespace(root CPLXMLNode, namespace string, recurse int) {
	cs := C.CString(namespace)
	defer C.free(unsafe.Pointer(cs))
	C.CPLStripXMLNamespace(root.cValue, cs, C.int(recurse))
}

func (n CPLXMLNode) StripNamespace(namespace string, recurse int) {
	cplStripXMLNamespace(n, namespace, recurse)
}

func cplCleanXMLElementName(name string) (result string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.CPLCleanXMLElementName(cs)
	result = C.GoString(cs)
	return
}

func CPLCleanXMLElementName(name string) (result string) {
	result = cplCleanXMLElementName(name)
	return
}

func cplParseXMLFile(filename string) (result CPLXMLNode) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = CPLXMLNode{cValue: C.CPLParseXMLFile(cs)}
	return
}

func CPLParseXMLFile(filename string) (result CPLXMLNode, err error) {
	result = cplParseXMLFile(filename)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func cplSerializeXMLTreeToFile(tree CPLXMLNode, filename string) (result int) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	result = int(C.CPLSerializeXMLTreeToFile(tree.cValue, cs))
	return
}

func (n CPLXMLNode) SerializeTreeToFile(filename string) int {
	return cplSerializeXMLTreeToFile(n, filename)
}

func cplXMLNodeGetRAMUsageEstimate(node CPLXMLNode) (result uint64) {
	result = uint64(C.CPLXMLNodeGetRAMUsageEstimate(node.cValue))
	return
}

func (n CPLXMLNode) GetRAMUsageEstimate() uint64 {
	return cplXMLNodeGetRAMUsageEstimate(n)
}
