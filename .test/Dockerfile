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
        
RUN cd /home/ && git clone https://github.com/bats-core/bats-core

ARG bashver=latest

RUN ln -s /home/bats-core/bin/bats /usr/sbin/bats

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ENTRYPOINT ["bash", "/usr/sbin/bats"]