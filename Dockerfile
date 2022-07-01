FROM alpine:3.15
USER root

RUN apk --no-cache add ca-certificates \
  && update-ca-certificates

COPY main /main

ENTRYPOINT /main
