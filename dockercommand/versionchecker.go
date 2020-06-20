package dockercommand

import (
	"github.com/Masterminds/semver"
	"github.com/Oppodelldog/droxy/logger"
)

type versionChecker struct {
	dockerVersion string
}

func (vc versionChecker) isVersionSupported(versionConstraint string) bool {
	constraints, err := semver.NewConstraint(versionConstraint)
	if err != nil {
		logger.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	dockerSemVer, err := semver.NewVersion(vc.dockerVersion)
	if err != nil {
		logger.Errorf("unable to check version constraint '%s': %v", versionConstraint, err)

		return false
	}

	return constraints.Check(dockerSemVer)
}
