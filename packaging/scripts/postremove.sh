#!/bin/sh
set -e

if [ ! -d /var/lib/errors ] && [ ! -d /etc/errors ]; then
    userdel errors 2>/dev/null || true
    groupdel errors 2>/dev/null || true
fi
