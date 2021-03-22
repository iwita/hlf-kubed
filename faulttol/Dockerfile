FROM golang:1.15-buster AS build

COPY ./ /go/src/github.com/faulttol
WORKDIR /go/src/github.com/faulttol

# Build application
RUN go build -o chaincode -v .

# Production ready image
# Pass the binary to the prod image
FROM debian:buster as prod

COPY --from=build /go/src/github.com/faulttol/chaincode /app/chaincode

USER 1000

WORKDIR /app
CMD ./chaincode