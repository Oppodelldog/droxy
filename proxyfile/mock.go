package proxyfile

type fileCreationFunctionMock struct {
	returnValue error
	parmSrc     string
	parmDst     string
	calls       int
}

func (m *fileCreationFunctionMock) FileCreationFunc(src, dst string) error {
	m.parmSrc = src
	m.parmDst = dst
	m.calls++

	return m.returnValue
}
