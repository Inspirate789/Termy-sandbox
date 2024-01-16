FROM docker:dind

ARG GOLANG_VERSION=1.21.0
ARG VAULT_VERSION=1.15.4

RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz &&  \
    tar -C /usr/local -zxvf go${GOLANG_VERSION}.linux-amd64.tar.gz &&  \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV GOBIN /go/bin

ADD https://releases.hashicorp.com/vault/${VAULT_VERSION}/vault_${VAULT_VERSION}_linux_amd64.zip /tmp/
ADD https://releases.hashicorp.com/vault/${VAULT_VERSION}/vault_${VAULT_VERSION}_SHA256SUMS      /tmp/
ADD https://releases.hashicorp.com/vault/${VAULT_VERSION}/vault_${VAULT_VERSION}_SHA256SUMS.sig  /tmp/

WORKDIR /tmp/

RUN apk --update add --virtual verify gpgme \
 && gpg --keyserver keyserver.ubuntu.com --recv-key 0x72D7468F \
 && gpg --verify /tmp/vault_${VAULT_VERSION}_SHA256SUMS.sig \
 && apk del verify \
 && cat vault_${VAULT_VERSION}_SHA256SUMS | grep linux_amd64 | sha256sum -c \
 && unzip vault_${VAULT_VERSION}_linux_amd64.zip \
 && mv vault /usr/local/bin/ \
 && rm -rf /tmp/* \
 && rm -rf /var/cache/apk/*
