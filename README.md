# Errors

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/errors/status.svg)](http://github.dronehippie.de/webhippie/errors)
[![Stories in Ready](https://badge.waffle.io/webhippie/errors.svg?label=ready&title=Ready)](http://waffle.io/webhippie/errors)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/8dbcb22838214efd940e75d2cffc31bc)](https://www.codacy.com/app/webhippie/errors?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/errors&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/webhippie/errors?status.svg)](http://godoc.org/github.com/webhippie/errors)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/errors)](http://goreportcard.com/report/github.com/webhippie/errors)
[![](https://images.microbadger.com/badges/image/tboerger/errors.svg)](http://microbadger.com/images/tboerger/errors "Get your own image badge on microbadger.com")


This project simply provides customizeable default and error pages for your Nginx Ingress controller on a Kubernetes cluster.


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.webhippie.de/misc/errors). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring of the dependencies you have to use a Go version `>= 1.6`. It is also possible to just simply execute the `go get github.com/webhippie/errors/cmd/errors` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/errors
cd $GOPATH/src/github.com/webhippie/errors
make clean retool sync build

./errors -h
```


## Security

If you find a security issue please contact thomas@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <http://www.webhippie.de>
```
