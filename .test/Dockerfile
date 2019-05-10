FROM docker:18.06-git

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    apk add --no-cache libc6-compat && \
    apk add --update make && \
    apk add gcc \
            musl-dev \
            openssl \
            go
        

RUN wget https://github.com/ovh/venom/releases/download/v0.25.0/venom.linux-amd64 && mv venom.linux-amd64 venom && chmod u+x venom && cp venom /usr/local/bin/venom   

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ENTRYPOINT ["/usr/local/bin/venom"]