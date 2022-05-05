---
title: "Getting Started"
date: 2022-05-04T00:00:00+00:00
anchor: "getting-started"
weight: 20
---

## Installation

So far we are offering only a few different variants for the installation. You
can choose between [Docker][docker] or pre-built binaries which are stored on
our download mirror and GitHub releases. Maybe we will also provide system
packages for the major distributions later if we see the need for it.

### Docker

Generally we are offering the images through
[quay.io/webhippie/errors][quay] and [webhippie/errors][dockerhub], so
feel free to choose one of the providers. Maybe we will come up with Kustomize
manifests or some Helm chart.

### Binaries

Simply download a binary matching your operating system and your architecture
from our [downloads][downloads] or the GitHub releases and place it within your
path like `/usr/local/bin` if you are using macOS or Linux.

## Configuration

We provide overall three different variants of configuration. The variant based
on environment variables and commandline flags are split up into global values
and command-specific values.

### Envrionment variables

If you prefer to configure the service with environment variables you can see
the available variables below.

#### Global

ERRORS_CONFIG_FILE
: Path to optional config file

ERRORS_LOG_LEVEL
: Set logging level, defaults to `info`

ERRORS_LOG_COLOR
: Enable colored logging, defaults to `true`

ERRORS_LOG_PRETTY
: Enable pretty logging, defaults to `true`

#### Server

ERRORS_METRICS_ADDR
: Address to bind the metrics, defaults to `0.0.0.0:8081`

ERRORS_METRICS_TOKEN
: Token to make metrics secure

ERRORS_SERVER_ADDR
: Address to bind the server, defaults to `0.0.0.0:8080`

ERRORS_SERVER_PPROF
: Enable pprof debugging, defaults to `false`

ERRORS_SERVER_ROOT
: Root path of the server, defaults to `/`

ERRORS_SERVER_HOST
: External access to server, defaults to `http://localhost:8080`

ERRORS_SERVER_CERT
: Path to cert for SSL encryption

ERRORS_SERVER_KEY
: Path to key for SSL encryption

ERRORS_SERVER_STRICT_CURVES
: Use strict SSL curves, defaults to `false`

ERRORS_SERVER_STRICT_CIPHERS
: Use strict SSL ciphers, defaults to `false`

ERRORS_SERVER_TEMPLATES
: Folder for custom templates

ERRORS_SERVER_ERRORS
: Path for overriding errors

#### Health

ERRORS_METRICS_ADDR
: Address to bind the metrics, defaults to `0.0.0.0:8081`

### Commandline flags

If you prefer to configure the service with commandline flags you can see the
available variables below.

#### Global

--config-file
: Path to optional config file

--log-level
: Set logging level, defaults to `info`

--log-color
: Enable colored logging, defaults to `true`

--log-pretty
: Enable pretty logging, defaults to `true`

#### Server

--metrics-addr
: Address to bind the metrics, defaults to `0.0.0.0:8081`

--metrics-token
: Token to make metrics secure

--server-addr
: Address to bind the server, defaults to `0.0.0.0:8080`

--server-pprof
: Enable pprof debugging, defaults to `false`

--server-root
: Root path of the server, defaults to `/`

--server-host
: External access to server, defaults to `http://localhost:8080`

--server-cert
: Path to cert for SSL encryption

--server-key
: Path to key for SSL encryption

--strict-curves
: Use strict SSL curves, defaults to `false`

--strict-ciphers
: Use strict SSL ciphers, defaults to `false`

--templates-path
: Folder for custom templates

--errors-path
: Path for overriding errors

#### Health

--metrics-addr
: Address to bind the metrics, defaults to `0.0.0.0:8081`

### Configuration file

So far we support multiple file formats like `json` or `yaml`, if you want to
get a full example configuration just take a look at [our repository][repo],
there you can always see the latest configuration format. These example configs
include all available options and the default values. The configuration file
will be automatically loaded if it's placed at
`/etc/errors/config.yml`, `${HOME}/.errors/config.yml` or
`$(pwd)/errors/config.yml`.

## Usage

The program provides a few sub-commands on execution. The available config
methods have already been mentioned above. Generally you can always see a
formated help output if you execute the binary similar to something like
 `errors --help`.

[docker]: https://www.docker.com/
[quay]: https://quay.io/repository/webhippie/errors
[dockerhub]: https://hub.docker.com/r/webhippie/errors
[downloads]: https://dl.webhippie.de/#errors/
[repo]: https://github.com/webhippie/errors/tree/master/config
