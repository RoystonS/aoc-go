# shellcheck shell=bash
set -e

USERNAME=${1}
USER_UID=${2:-autogen}

apt_get_update_if_needed() {
    if [ ! -d "/var/lib/apt/lists" ] || [ "$(find /var/lib/apt/lists -type f | wc -l)" = "0" ]; then
        echo "Running apt-get update..."
        apt-get update
    else
        echo "Skipping apt-get update."
    fi
}

apt_get_update_if_needed

export DEBIAN_FRONTEND=noninteractive
apt-get -y install --no-install-recommends \
    apt-utils \
    ca-certificates \
    apt-transport-https \
    jq \
    unzip \
    zip \
    curl \
    sudo \
    vim-tiny \
    locales

# Create user
if [ "${USER_UID}" = "autogen" ]; then
    useradd -s /bin/bash -m $USERNAME
else
    useradd -s /bin/bash --uid $USER_UID -m $USERNAME
fi

# Grant passwordless sudo to user
mkdir -p /etc/sudoers.d
echo $USERNAME ALL=\(root\) NOPASSWD:ALL >/etc/sudoers.d/$USERNAME
chmod 0440 /etc/sudoers.d/$USERNAME
