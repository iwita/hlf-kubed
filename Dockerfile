FROM golang:1.14.6-buster AS build

COPY ./ /go/src/github.com/fabcar
WORKDIR /go/src/github.com/fabcar

# Build application
RUN go build -o chaincode -v .

# Production ready image
# Pass the binary to the prod image
FROM buster as prod

COPY --from=build /go/src/github.com/fabcar/chaincode /app/chaincode

USER 1000

WORKDIR /app
CMD ./chaincode