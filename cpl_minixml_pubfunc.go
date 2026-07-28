package gdal

func CPLParseXMLString(xml string) (result CPLXMLNode, err error) {
	result = cplParseXMLString(xml)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (n CPLXMLNode) Destroy() {
	cplDestroyXMLNode(n)
}

func (n CPLXMLNode) GetNode(path string) (result CPLXMLNode, err error) {
	result = cplGetXMLNode(n, path)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (n CPLXMLNode) SearchNode(target string) (result CPLXMLNode, err error) {
	result = cplSearchXMLNode(n, target)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (n CPLXMLNode) GetValue(path, dflt string) string {
	return cplGetXMLValue(n, path, dflt)
}

func (n CPLXMLNode) CreateNode(eType CPLXMLNodeType, text string) (result CPLXMLNode, err error) {
	result = cplCreateXMLNode(n, eType, text)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (n CPLXMLNode) SerializeTree() string {
	return cplSerializeXMLTree(n)
}

func (n CPLXMLNode) AddChild(child CPLXMLNode) {
	cplAddXMLChild(n, child)
}

func (n CPLXMLNode) RemoveChild(child CPLXMLNode) int {
	return cplRemoveXMLChild(n, child)
}

func (n CPLXMLNode) AddSibling(newSibling CPLXMLNode) {
	cplAddXMLSibling(n, newSibling)
}

func (n CPLXMLNode) CreateElementAndValue(name, value string) (result CPLXMLNode, err error) {
	result = cplCreateXMLElementAndValue(n, name, value)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (n CPLXMLNode) AddAttributeAndValue(name, value string) {
	cplAddXMLAttributeAndValue(n, name, value)
}

func (n CPLXMLNode) CloneTree() (result CPLXMLNode, err error) {
	result = cplCloneXMLTree(n)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (n CPLXMLNode) SetValue(path, value string) int {
	return cplSetXMLValue(n, path, value)
}

func (n CPLXMLNode) StripNamespace(namespace string, recurse int) {
	cplStripXMLNamespace(n, namespace, recurse)
}

func CPLCleanXMLElementName(name string) (result string) {
	result = cplCleanXMLElementName(name)
	return
}

func CPLParseXMLFile(filename string) (result CPLXMLNode, err error) {
	result = cplParseXMLFile(filename)
	if result.cValue == nil {
		err = lastError()
	}
	return
}

func (n CPLXMLNode) SerializeTreeToFile(filename string) int {
	return cplSerializeXMLTreeToFile(n, filename)
}

func (n CPLXMLNode) GetRAMUsageEstimate() uint64 {
	return cplXMLNodeGetRAMUsageEstimate(n)
}
