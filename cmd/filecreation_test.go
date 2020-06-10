package cmd

import (
	"os"
	"testing"
)

type proxyFilesCreatorMock struct {
	isForced bool
	calls    int
}

func (p *proxyFilesCreatorMock) CreateProxyFiles(isForced bool) error {
	p.calls++
	p.isForced = isForced

	return nil
}

func TestNewCloneCommandWrapper(t *testing.T) {
	mock := proxyFilesCreatorMock{}
	name := "testCommandName"

	command := createCommand(name, &mock)

	gotName := command.Name()
	wantName := name

	if wantName != gotName {
		t.Fatalf("wantName: %v, gotName: %v", wantName, gotName)
	}

	err := command.Execute()
	if err != nil {
		t.Fatalf("did not expect Execute to return an error, but got: %v", err)
	}

	gotCalls := mock.calls
	wantCalls := 1

	if wantCalls != gotCalls {
		t.Fatalf("wantCalls: %v, gotCalls: %v", wantCalls, gotCalls)
	}

	gotForced := mock.isForced
	wantForced := false

	if wantForced != gotForced {
		t.Fatalf("wantForced: %v, gotForced: %v", wantForced, gotForced)
	}
}

func TestNewCloneCommandWrapper_Forced(t *testing.T) {
	mock := proxyFilesCreatorMock{}
	name := ""

	originalArgLen := len(os.Args)
	os.Args = append(os.Args, "-f")

	defer func() {
		os.Args = os.Args[:originalArgLen]
	}()

	err := createCommand(name, &mock).Execute()
	if err != nil {
		t.Fatalf("did not expect Execute to return an error, but got: %v", err)
	}

	gotForced := mock.isForced
	wantForced := true

	if wantForced != gotForced {
		t.Fatalf("wantForced: %v, gotForced: %v", wantForced, gotForced)
	}
}
