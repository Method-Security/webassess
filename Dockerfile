FROM ubuntu:24.04

ARG CLI_NAME="webassess"
ARG TARGETARCH

RUN apt update && apt install -y jq ca-certificates git curl

# Install ollama
RUN curl -fsSL https://ollama.com/install.sh | sh

# Setup Method Directory Structure
RUN \
  mkdir -p /opt/method/${CLI_NAME}/ && \
  mkdir -p /opt/method/${CLI_NAME}/var/data && \
  mkdir -p /opt/method/${CLI_NAME}/var/data/tmp && \
  mkdir -p /opt/method/${CLI_NAME}/var/conf && \
  mkdir -p /opt/method/${CLI_NAME}/var/conf/models && \
  mkdir -p /opt/method/${CLI_NAME}/var/log && \
  mkdir -p /opt/method/${CLI_NAME}/service/bin && \
  mkdir -p /mnt/output

# Copy the CLI binary and make it executable
COPY ${CLI_NAME} /opt/method/${CLI_NAME}/service/bin/${CLI_NAME}

# Install models, ollama needs to be running both for install and entrypoint
# A single dedicated script is used to have ollama serve be run in the same terminal
ENV OLLAMA_MODELS=/opt/method/${CLI_NAME}/var/conf/models
RUN echo '#!/usr/bin/env bash\n\nollama serve &\nsleep 5\nollama list\nollama pull qwen2.5:0.5b\nollama pull qwen2.5:3b\nollama pull gemma2:2b' > /opt/method/${CLI_NAME}/var/conf/setup-ollama.sh
RUN chmod +x /opt/method/${CLI_NAME}/var/conf/setup-ollama.sh
RUN /opt/method/${CLI_NAME}/var/conf/setup-ollama.sh

RUN \
  adduser --disabled-password --gecos '' method && \
  chown -R method:method /opt/method/${CLI_NAME}/ && \
  chown -R method:method /mnt/output

USER method

WORKDIR /opt/method/${CLI_NAME}/

ENV PATH="/opt/method/${CLI_NAME}/service/bin:${PATH}"
ENTRYPOINT [ "webassess" ]