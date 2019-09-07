FROM golang:1.13.0

WORKDIR /go/src/github.com/Oppodelldog/droxy

RUN go get -u golang.org/x/tools/cmd/goimports
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.17.1
RUN wget -O codecov https://codecov.io/bash && chmod ugo+x codecov && cp codecov $GOPATH/bin/.

RUN chmod -R a+rw $GOPATH/src
    