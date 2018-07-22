FROM golang:1.10.3

WORKDIR /go/src/github.com/Oppodelldog/droxy

RUN go get -u gopkg.in/alecthomas/gometalinter.v2 && \
	go get -u github.com/golang/dep/cmd/dep && \
	go get -u golang.org/x/tools/cmd/goimports && \
    go get -u github.com/mattn/goveralls

