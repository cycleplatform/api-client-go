FROM golang:1.26-bookworm

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

RUN apt-get update \
    && apt-get install -y --no-install-recommends curl ca-certificates git sudo \
    && rm -rf /var/lib/apt/lists/*

# Node.js (matches the version used in CI) for building the api-spec submodule
COPY --from=node:22-bookworm-slim /usr/local/bin/node /usr/local/bin/node
COPY --from=node:22-bookworm-slim /usr/local/lib/node_modules /usr/local/lib/node_modules
RUN ln -s ../lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm \
    && ln -s ../lib/node_modules/npm/bin/npx-cli.js /usr/local/bin/npx \
    && node --version && npm --version

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID --create-home --shell /bin/bash $USERNAME \
    && echo "$USERNAME ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && chown -R $USER_UID:$USER_GID /go /usr/local/bin

USER $USERNAME
ENV HOME=/home/$USERNAME

RUN curl -fsSL https://claude.ai/install.sh | bash
ENV PATH="$HOME/.local/bin:$HOME/.claude/bin:${PATH}"
RUN claude --version

WORKDIR /workspace
