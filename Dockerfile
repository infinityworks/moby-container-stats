FROM golang:1.15.1 as build
LABEL maintainer "Infinity Works"

COPY ./ /go/src/github.com/infinityworks/moby-container-stats
WORKDIR /go/src/github.com/infinityworks/moby-container-stats

RUN go get \
 && go test ./... \
 && go build -o /bin/main

ENV LISTEN_PORT=9244
EXPOSE 9244
ENTRYPOINT [ "/bin/main" ]
