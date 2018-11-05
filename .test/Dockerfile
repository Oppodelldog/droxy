FROM docker:18.06-git

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    apk add --no-cache libc6-compat
        
RUN cd /home/ && git clone https://github.com/bats-core/bats-core

ARG bashver=latest

RUN ln -s /home/bats-core/bin/bats /usr/sbin/bats

ENTRYPOINT ["bash", "/usr/sbin/bats"]