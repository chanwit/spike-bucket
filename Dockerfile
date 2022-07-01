FROM alpine:latest
USER root
RUN apk --no-cache add ca-certificates \
  && update-ca-certificates \

COPY main /main

ENTRYPOINT /main
