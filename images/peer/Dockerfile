# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER
ARG ALPINE_VER

FROM arm64v8/ubuntu as peer-base
#RUN apk add --no-cache tzdata
# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN echo 'hosts: files dns' > /etc/nsswitch.conf

FROM arm64v8/golang:1.14-buster as golang
RUN apt install -y \
	bash \
	gcc \
	git \
	make \
        binutils;

ADD . $GOPATH/src/github.com/hyperledger-v2.2.0/fabric
WORKDIR $GOPATH/src/github.com/hyperledger-v2.2.0/fabric

FROM golang as peer
ARG GO_TAGS
RUN make peer GO_TAGS=${GO_TAGS} GOOS=linux GOARCH=arm64


FROM peer-base
ENV FABRIC_CFG_PATH /etc/hyperledger/fabric
VOLUME /etc/hyperledger/fabric
VOLUME /var/hyperledger
COPY --from=peer /go/src/github.com/hyperledger-v2.2.0/fabric/build/bin /usr/local/bin
COPY --from=peer /go/src/github.com/hyperledger-v2.2.0/fabric/sampleconfig/msp ${FABRIC_CFG_PATH}/msp
COPY --from=peer /go/src/github.com/hyperledger-v2.2.0/fabric/sampleconfig/core.yaml ${FABRIC_CFG_PATH}
EXPOSE 7051
CMD ["peer","node","start"]
