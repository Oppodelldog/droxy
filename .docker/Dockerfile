FROM golang:1.11.1

WORKDIR /go/src/github.com/Oppodelldog/droxy

RUN go get -u gopkg.in/alecthomas/gometalinter.v2 && \
	go get -u github.com/golang/dep/cmd/dep && \
	go get -u golang.org/x/tools/cmd/goimports && \
    wget -O codecov https://codecov.io/bash && chmod ugo+x codecov && cp codecov /usr/local/bin/.

