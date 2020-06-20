package dockercommand

import "testing"

func TestVersionChecker_isVersionSupported(t *testing.T) {
	testCases := map[string]struct {
		dockerVersion string
		inputVersion  string
		want          bool
	}{
		"empty":                  {want: false},
		"docker version empty":   {want: false, inputVersion: "1.0"},
		"input version empty":    {want: false, dockerVersion: "1.0"},
		"docker version too low": {want: false, dockerVersion: "1.0", inputVersion: "2.0"},
		"versions same":          {want: true, dockerVersion: "2.0", inputVersion: "2.0"},
		"docker version higher":  {want: true, dockerVersion: "1.9", inputVersion: ">1.0"},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			got := versionChecker{dockerVersion: testCase.dockerVersion}.isVersionSupported(testCase.inputVersion)
			want := testCase.want

			if want != got {
				t.Fatalf("want: %v, got: %v", want, got)
			}
		})
	}
}
