#!/bin/sh
set -e

if ! getent group errors >/dev/null 2>&1; then
    groupadd --system errors
fi

if ! getent passwd errors >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/errors --shell /bin/bash -g errors errors
fi
