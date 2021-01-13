#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
ARG GO_VER
ARG ALPINE_VER

FROM arm64v8/golang:1.14-buster as builder
ARG GO_LDFLAGS
ARG GO_TAGS

RUN apt install -y \
	gcc \
	git;

#RUN apk add --no-cache \
#	gcc \
#	git \
#	libtool \
#	musl-dev \
#	binutils-gold;

ADD . $GOPATH/src/github.com/hyperledger/fabric-ca
RUN ls
RUN GOOS=linux GOARCH=arm64 go install -tags "${GO_TAGS}" -ldflags "${GO_LDFLAGS}" \
	github.com/hyperledger/fabric-ca/cmd/fabric-ca-server \
	&& GOOS=linux GOARCH=arm64 go install -tags "${GO_TAGS}" -ldflags "${GO_LDFLAGS}" \
	github.com/hyperledger/fabric-ca/cmd/fabric-ca-client

FROM arm64v8/ubuntu
#RUN apk add --no-cache \
#	tzdata \
#	libltdl;
ENV FABRIC_CA_HOME /etc/hyperledger/fabric-ca-server
COPY --from=builder /go/bin /usr/local/bin
EXPOSE 7054
CMD fabric-ca-server start -b admin:adminpw
