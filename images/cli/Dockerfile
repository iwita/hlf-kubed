# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER
ARG ALPINE_VER
FROM arm64v8/golang:1.14-buster as golang

RUN apt install -y \
	bash \
	gcc \
	git \
	make \
	binutils;

ADD . $GOPATH/src/github.com/hyperledger-v2.2.0/fabric
WORKDIR $GOPATH/src/github.com/hyperledger-v2.2.0/fabric

FROM golang as tools
RUN make configtxgen configtxlator cryptogen peer discover idemixgen

FROM arm64v8/golang:1.14-buster
# git is required to support `go list -m`
RUN apt install -y \
	bash \
        jq \
	git;
ENV FABRIC_CFG_PATH /etc/hyperledger/fabric
VOLUME /etc/hyperledger/fabric
COPY --from=tools /go/src/github.com/hyperledger-v2.2.0/fabric/build/bin /usr/local/bin
COPY --from=tools /go/src/github.com/hyperledger-v2.2.0/fabric/sampleconfig ${FABRIC_CFG_PATH}
