FROM --platform=$BUILDPLATFORM golang:1.26.5-alpine@sha256:99e12cfb19b753915f9b9fdc5a99f1869a24a69d3a0955832d5702e7fa68f1be AS builder

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

FROM alpine:3.24@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

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
