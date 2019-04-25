module github.com/Oppodelldog/droxy

go 1.12

require (
	github.com/BurntSushi/toml v0.3.0
	github.com/Masterminds/semver v1.4.2
	github.com/Microsoft/go-winio v0.4.11
	github.com/Oppodelldog/filediscovery v0.0.0-20180705182803-4da0ea434b80
	github.com/davecgh/go-spew v1.1.0
	github.com/docker/distribution v2.6.2+incompatible
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.3.3
	github.com/drone/envsubst v0.0.0-20180915092829-73c8b7a36b56
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/magiconair/properties v1.7.6
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/sirupsen/logrus v1.0.4
	github.com/spf13/cobra v0.0.1
	github.com/spf13/pflag v1.0.0
	github.com/stretchr/objx v0.1.0
	github.com/stretchr/testify v1.2.1
	golang.org/x/crypto v0.0.0-20180127211104-1875d0a70c90
	golang.org/x/net v0.0.0-20181005035420-146acd28ed58
	golang.org/x/sys v0.0.0-20180202135801-37707fdb30a5
)

replace github.com/drone/envsubst => github.com/Oppodelldog/envsubst v0.0.0-20180915092829-73c8b7a36b56
