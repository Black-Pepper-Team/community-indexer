FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/black-pepper-team/community-indexer
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/community-indexer /go/src/github.com/black-pepper-team/community-indexer


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/community-indexer /usr/local/bin/community-indexer
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["community-indexer"]
