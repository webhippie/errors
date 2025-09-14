#!/bin/sh
set -e

chown -R errors:errors /etc/errors
chown -R errors:errors /var/lib/errors
chmod 750 /var/lib/errors

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet errors.service; then
        systemctl restart errors.service
    fi
fi
