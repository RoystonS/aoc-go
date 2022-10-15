# shellcheck shell=bash
set -e

export DEBIAN_FRONTEND=noninteractive
apt-get -y install --no-install-recommends \
    bc \
    pup \
    lynx
