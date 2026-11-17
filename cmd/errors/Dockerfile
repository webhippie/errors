FROM --platform=$BUILDPLATFORM golang:1.27.4-alpine@sha256:70bd7ddea8b6cbe5d6943cde279d6b67feb10cc924fb9100d1d8757c1575b280 AS builder

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

FROM alpine:3.24@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf

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
