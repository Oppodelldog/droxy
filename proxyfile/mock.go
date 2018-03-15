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

type mockFileCreationStrategy struct {
	returnValue               error
	parmCommandBinaryFilePath string
	parmCommandNameFileName   string
	calls                     int
}

func (m *mockFileCreationStrategy) CreateProxyFile(commandBinaryFilePath string, commandNameFileName string) error {
	m.parmCommandBinaryFilePath = commandBinaryFilePath
	m.parmCommandNameFileName = commandNameFileName
	m.calls++

	return m.returnValue
}
