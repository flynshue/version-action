FROM golang:alpine3.16 as build

WORKDIR $GOPATH/src/

ADD . .

RUN go get -d -v && go build -o /go/bin/version-action

FROM alpine:latest

COPY --from=build /go/bin/version-action /go/bin/version-action

CMD ["/go/bin/version-action"]