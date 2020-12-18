package arguments

import (
	"errors"
	"github.com/Oppodelldog/droxy/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testCase struct {
	input       string
	overwrites  *[]string
	requireVars bool

	want      string
	wantErr   error
	wantPanic bool

	prepare func()
	cleanup func()
}

type testCases map[string]testCase

func (c testCases) run(t *testing.T) {
	for name, testData := range c {
		t.Run(name, func(t *testing.T) {
			runSubstitution(t, testData)
		})
	}
}

func Test_newEnvVarResolver_invalidEnvVarMappings(t *testing.T) {
	def := config.CommandDefinition{
		EnvVarOverwrites: &[]string{"VAR=TEST=1"},
	}

	assert.Panics(t, func() { newEnvVarResolver(def) })
}

func Test_resolveEnvVar(t *testing.T) {
	testCases{
		"no var provided, no change": {
			input: "VAR=VALUE",
			want:  "VAR=VALUE",
		},
		"var provided, but not set, substitute empty": {
			input: "VAR=${VALUE}",
			want:  "VAR=",
		},
		"var provided, env is set, substitute from env value": {
			input: "VAR=${VALUE}",
			want:  "VAR=123",
			prepare: func() {
				must(t, os.Setenv("VALUE", "123"))
			},
		},
		"var provided, env is set, overwrites is set, substitute from overwrites": {
			input: "VAR=${VALUE}",
			want:  "VAR=456",
			prepare: func() {
				must(t, os.Setenv("VALUE", "123"))
			},
			cleanup: func() {
				os.Unsetenv("VALUE")
			},
			overwrites: &[]string{"VALUE=456"},
		},
	}.run(t)
}

func Test_resolveEnvVar_strict(t *testing.T) {
	testCases{
		"no var provided, no change": {
			input:       "VAR=VALUE",
			want:        "VAR=VALUE",
			requireVars: true,
		},
		"var provided, but not set, substitute empty": {
			input:       "VAR=${VALUE}",
			want:        "VAR=",
			requireVars: true,
			wantPanic:   true,
		},
		"var provided, and not set, substitute from envVar": {
			input:       "VAR=${VALUE}",
			want:        "VAR=123",
			requireVars: true,
			prepare: func() {
				must(t, os.Setenv("VALUE", "123"))
			},
		},
	}.run(t)
}

func runSubstitution(t *testing.T, testData testCase) {
	if testData.prepare != nil {
		testData.prepare()
	}

	defer func() {
		if testData.cleanup != nil {
			testData.cleanup()
		}
	}()

	def := config.CommandDefinition{
		RequireEnvVars:   &testData.requireVars,
		EnvVarOverwrites: testData.overwrites,
	}

	if testData.wantPanic {
		assert.Panics(t, func() {
			_, _ = newEnvVarResolver(def).substitute(testData.input)
		})
	} else {
		res, err := newEnvVarResolver(def).substitute(testData.input)
		if !errors.Is(testData.wantErr, err) {
			t.Fatalf("want: %v, got: %v", testData.wantErr, err)
		}
		if testData.want != res {
			t.Fatalf("want: %v, got: %v", testData.want, res)
		}
	}
}

func must(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
