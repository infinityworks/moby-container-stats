FROM golang:1.9.0-alpine as build
LABEL maintainer "Infinity Works"

RUN apk --update add ca-certificates \
     && apk --update add --virtual build-deps git

COPY ./ /go/src/github.com/infinityworks/moby-container-stats
WORKDIR /go/src/github.com/infinityworks/moby-container-stats

RUN go get \
 && go test ./... \
 && go build -o /bin/main

FROM alpine:3.6

RUN apk --update add ca-certificates 

COPY --from=build /bin/main /bin/main


ENV LISTEN_PORT=9244
EXPOSE 9244
ENTRYPOINT [ "/bin/main" ]
