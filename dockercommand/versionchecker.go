package dockercommand

import (
	"github.com/Masterminds/semver"
	"github.com/sirupsen/logrus"
)

type versionChecker struct {
	versionProvider dockerVersionProvider
}

func (vc versionChecker) isVersionSupported(versionConstraint string) bool {
	constraints, err := semver.NewConstraint(versionConstraint)
	if err != nil {
		logrus.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	dockerVersion, err := vc.versionProvider.getAPIVersion()
	if err != nil {
		logrus.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	dockerSemVer, err := semver.NewVersion(dockerVersion)
	if err != nil {
		logrus.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	return constraints.Check(dockerSemVer)
}
