# Changelog for unreleased

The following sections list the changes for unreleased.

## Summary

 * Chg #14: Use Alpine base image and define healthcheck

## Details

 * Change #14: Use Alpine base image and define healthcheck

   We just switched the base image to Alpine Linux instead of a scratch image as this could lead to
   issues related to healthcheck commands defined by the Docker CLI. Beside that we have added the
   healthcheck command to the Dockerfile for having a proper handling by default.

   https://github.com/webhippie/errors/pull/14


# Changelog for 1.0.1

The following sections list the changes for 1.0.1.

## Summary

 * Fix #10: Add healthcheck endpoint to http

## Details

 * Bugfix #10: Add healthcheck endpoint to http

   We had the healthchecks only on our metrics port available, but this behavior broke the helm
   installation for ingress-nginx as this part is not configurable. Just add it to the http port to
   get this working correctly.

   https://github.com/webhippie/errors/issues/10


# Changelog for 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #4: Initial release of basic version

## Details

 * Change #4: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/webhippie/errors/issues/4


