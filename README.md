# ownCloud: Errors

[![Build Status](https://drone.owncloud.com/api/badges/owncloud-docker/errors/status.svg)](https://drone.owncloud.com/owncloud-docker/errors)
[![](https://images.microbadger.com/badges/image/owncloud/errors.svg)](https://microbadger.com/images/owncloud/errors "Get your own image badge on microbadger.com")

This is our Docker image used as a error page within our cluster to show a proper page, it is based on our [Ubuntu container](https://registry.hub.docker.com/u/owncloud/ubuntu/).


## Build locally

The available versions should be already pushed to the Docker Hub, but in case you want to try a change locally you can always execute the following command to get this image built locally:

```
docker build -t owncloud/errors:latest .
```


## Versions

To get an overview about the available versions please take a look at the [GitHub branches](https://github.com/owncloud-docker/errors/branches/all) or our [Docker Hub tags](https://hub.docker.com/r/owncloud/errors/tags/), these lists are always up to date.


## Volumes

* N/A


## Ports

* 8080


## Available environment variables

```
ERRORS_ADDRESS 0.0.0.0:8080
ERRORS_ASSETS /usr/share/errors
```


## Issues, Feedback and Ideas

Open an [Issue](https://github.com/owncloud-docker/errors/issues)


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

MIT


## Copyright

```
Copyright (c) 2018 Thomas Boerger <tboerger@owncloud.com>
```
