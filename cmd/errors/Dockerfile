FROM --platform=$BUILDPLATFORM golang:1.27.3-alpine@sha256:f2c872bcee949b7be9c490efd02b48a8a2fbf947cdf533f1c58525ec9d8170fd AS builder

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

FROM alpine:3.24@sha256:c6b073f8e9e58d881f6ffba6ebf3b9ebd60b9e7d6e8337018c0e68bb5c356873

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
