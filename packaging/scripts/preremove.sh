#!/bin/sh
set -e

systemctl stop errors.service || true
systemctl disable errors.service || true
