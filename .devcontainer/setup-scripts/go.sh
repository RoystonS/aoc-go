# shellcheck shell=bash
export GOTARGET=/usr/local/go

toolstemp=/tmp/gotools

# See https://github.com/golang/vscode-go/blob/v0.35.2/src/goToolsInformation.ts

export PATH=${GOTARGET}/bin:${PATH}

mkdir -p "${toolstemp}"
cd "${toolstemp}" || exit
export GOPATH=${toolstemp}
export GOCACHE=${toolstemp}/cache

go_pkgs="golang.org/x/tools/gopls@latest \
    honnef.co/go/tools/cmd/staticcheck@latest \
    golang.org/x/lint/golint@latest \
    github.com/mgechev/revive@latest \
    github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
    github.com/ramya-rao-a/go-outline@latest \
    github.com/go-delve/delve/cmd/dlv@latest \
    github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

echo "${go_pkgs}" | xargs -n 1 go install -v
mv bin/* "${GOTARGET}/bin/"
cd /
rm -rf "${toolstemp}"
