module github.com/Oppodelldog/droxy

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/semver v1.4.2
	github.com/Microsoft/go-winio v0.4.11 // indirect
	github.com/Oppodelldog/filediscovery v0.3.0
	github.com/docker/distribution v2.6.2+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.3.3 // indirect
	github.com/drone/envsubst v0.0.0-20180915092829-73c8b7a36b56
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/stevvooe/resumable v0.0.0-20180830230917-22b14a53ba50 // indirect
	github.com/stretchr/testify v1.3.0
	golang.org/x/net v0.0.0-20181005035420-146acd28ed58 // indirect
)

replace github.com/drone/envsubst => github.com/Oppodelldog/envsubst v0.0.0-20180915092829-73c8b7a36b56
