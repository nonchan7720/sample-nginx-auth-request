FROM golang:1.21.1-bullseye

ARG GIT_VERSION=2.39.1

RUN apt-get update \
    && apt-get install --no-install-recommends -y \
    make \
    sudo \
    protobuf-compiler \
    vim \
    procps \
    curl \
    openssh-client \
    gettext \
    libexpat1-dev \
    libz-dev \
    libssl-dev \
    libcurl4-openssl-dev \
    && cd /usr/local/src \
    && wget https://github.com/git/git/archive/refs/tags/v${GIT_VERSION}.tar.gz \
    && tar -xzf v${GIT_VERSION}.tar.gz \
    && cd git-${GIT_VERSION} \
    && make prefix=/usr/local all -j "$(nproc)" \
    && make prefix=/usr/local install \
    && go install github.com/cweill/gotests/gotests@v1.6.0 \
    && go install github.com/fatih/gomodifytags@v1.16.0 \
    && go install github.com/josharian/impl@v1.2.0 \
    && go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2 \
    && go install golang.org/x/tools/gopls@v0.11.0 \
    && go install github.com/stormcat24/protodep@v0.1.7 \
    && go install github.com/go-delve/delve/cmd/dlv@v1.21.0

ARG USER_NAME=app \
    GROUP_NAME=app \
    USER_ID=1000 \
    GROUP_ID=1000 \
    PASSWORD=app

RUN groupadd -r -g $GROUP_ID ${GROUP_NAME} \
    && useradd -m -s /bin/bash -u $USER_ID -g $GROUP_ID -G sudo $USER_NAME \
    && echo $USER_NAME:$PASSWORD | chpasswd \
    && echo "$USER_NAME   ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && mkdir -p /src \
    && chown $USER_ID:$GROUP_ID -R /go /src \
    && rm -rf /var/lib/apt/lists/*

USER $USER_NAME

WORKDIR /src
