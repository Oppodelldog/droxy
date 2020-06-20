package dockercommand

import (
	"errors"

	"github.com/Masterminds/semver"
	"github.com/Oppodelldog/droxy/logger"
)

var errCheckingVersionConstraint = errors.New("unable to check version constraint")

type versionChecker struct {
	dockerVersion string
}

func (vc versionChecker) isVersionSupported(versionConstraint string) bool {
	constraints, err := semver.NewConstraint(versionConstraint)
	if err != nil {
		logger.Errorf("%v '%s': %v", errCheckingVersionConstraint, versionConstraint, err)

		return false
	}

	dockerSemVer, err := semver.NewVersion(vc.dockerVersion)
	if err != nil {
		logger.Errorf("%v '%s': %v", errCheckingVersionConstraint, versionConstraint, err)

		return false
	}

	return constraints.Check(dockerSemVer)
}
