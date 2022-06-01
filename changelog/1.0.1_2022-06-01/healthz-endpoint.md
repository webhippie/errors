Bugfix: Add healthcheck endpoint to http

We had the healthchecks only on our metrics port available, but this behavior
broke the helm installation for ingress-nginx as this part is not configurable.
Just add it to the http port to get this working correctly.

https://github.com/webhippie/errors/issues/10
