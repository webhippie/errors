FROM --platform=$BUILDPLATFORM golang:1.25.5-alpine3.21@sha256:b4dbd292a0852331c89dfd64e84d16811f3e3aae4c73c13d026c4d200715aff6 AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/errors
COPY . /go/src/errors/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.23@sha256:865b95f46d98cf867a156fe4a135ad3fe50d2056aa3f25ed31662dff6da4eb62

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 errors && \
    adduser -D -u 1337 -h /var/lib/errors -G errors errors

EXPOSE 8080 8081
VOLUME ["/var/lib/errors"]
ENTRYPOINT ["/usr/bin/errors"]
CMD ["server"]
HEALTHCHECK CMD ["/usr/bin/errors", "health"]

COPY --from=builder /go/src/errors/bin/errors /usr/bin/errors
WORKDIR /var/lib/errors
USER errors
