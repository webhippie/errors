FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Errors" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

EXPOSE 8080
ENTRYPOINT ["/usr/bin/errors"]

ENV ERRORS_ASSETS /usr/share/errors

RUN apk add --no-cache ca-certificates mailcap bash

COPY assets/ /usr/share/errors/
COPY dist/binaries/errors-*-linux-amd64 /usr/bin/errors
