FROM golang:1.15.6-alpine3.12 as build-env

RUN mkdir /meshpaas
RUN apk --update add ca-certificates
RUN apk add make git
WORKDIR /meshpaas
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install ./...

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=build-env /go/bin/ /usr/local/bin/
WORKDIR /workspace
EXPOSE 8820
EXPOSE 8821

ENTRYPOINT ["/usr/local/bin/meshpaas"]