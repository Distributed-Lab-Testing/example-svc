FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/Distributed-Lab-Testing/example-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/example-svc /go/src/github.com/Distributed-Lab-Testing/example-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/example-svc /usr/local/bin/example-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["example-svc"]
