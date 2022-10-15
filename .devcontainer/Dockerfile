ARG VARIANT="1.19"
FROM golang:${VARIANT}

COPY setup-scripts/* /tmp/setup-scripts/

ARG USERNAME=vscode
ARG USER_UID=1000

RUN bash /tmp/setup-scripts/common-debian.sh "${USERNAME}" "${USER_UID}"
RUN bash /tmp/setup-scripts/go.sh
RUN bash /tmp/setup-scripts/aoc-debian.sh

# USER vscode
# RUN go get -x <your-dependency-or-tool>

RUN rm -rf /tmp/setup-scripts/
