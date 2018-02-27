package filediscovery

type fileLocationProviderMock struct {
	wasCalled         bool
	fileNameParameter string
}

func (m *fileLocationProviderMock) WasCalled() bool {
	return m.wasCalled
}

func (m *fileLocationProviderMock) GetCalledFilenameParameter() string {
	return m.fileNameParameter
}

func (m *fileLocationProviderMock) GetFunc(returnValue string, returnError error) FileLocationProvider {
	return func(fileName string) (string, error) {
		m.wasCalled = true
		m.fileNameParameter = fileName

		return returnValue, returnError
	}
}

func newFileLocationProviderMock() (*fileLocationProviderMock, FileLocationProvider) {
	mock := &fileLocationProviderMock{}
	return mock, mock.GetFunc("", nil)
}
