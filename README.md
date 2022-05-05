# Errors

[![Current Tag](https://img.shields.io/github/v/tag/webhippie/errors?sort=semver)](https://github.com/webhippie/errors) [![Build Status](https://github.com/webhippie/errors/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/errors/actions) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Docker Size](https://img.shields.io/docker/image-size/webhippie/errors/latest)](https://hub.docker.com/r/webhippie/errors) [![Docker Pulls](https://img.shields.io/docker/pulls/webhippie/errors)](https://hub.docker.com/r/webhippie/errors) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/errors.svg)](https://pkg.go.dev/github.com/webhippie/errors) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/errors)](https://goreportcard.com/report/github.com/webhippie/errors) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/8dbcb22838214efd940e75d2cffc31bc)](https://www.codacy.com/gh/webhippie/errors/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/errors&amp;utm_campaign=Badge_Grade)

This project simply provides customizeable default and error pages for your
Nginx Ingress Controller on a Kubernetes cluster. By default it already ships
with simple but nice pages for standard errors.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/webhippie/errors.git
cd errors

make generate build

./bin/errors -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/webhippie/errors/releases
[dockerhub]: https://hub.docker.com/r/webhippie/errors/tags/
[quay]: https://quay.io/repository/webhippie/errors?tab=tags
[docs]: https://webhippie.github.io/errors/#getting-started
[golang]: http://golang.org/doc/install.html
