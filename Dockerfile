FROM ubuntu:24.04

ARG CLI_NAME="aiassess"
ARG TARGETARCH

RUN apt update && apt install -y ca-certificates git curl

# Install ollama
#RUN curl -fsSL https://ollama.com/install.sh | sh

# Setup Method Directory Structure
RUN \
  mkdir -p /opt/method/${CLI_NAME}/ && \
  mkdir -p /opt/method/${CLI_NAME}/var/data && \
  mkdir -p /opt/method/${CLI_NAME}/var/data/tmp && \
  mkdir -p /opt/method/${CLI_NAME}/var/conf && \
  mkdir -p /opt/method/${CLI_NAME}/var/log && \
  mkdir -p /opt/method/${CLI_NAME}/service/bin && \
  mkdir -p /mnt/output

# Copy the CLI binary and make it executable
COPY ${CLI_NAME} /opt/method/${CLI_NAME}/service/bin/${CLI_NAME}
RUN chmod +x /opt/method/${CLI_NAME}/service/bin/${CLI_NAME}

# Install models, ollama needs to be running both for install and entrypoint
COPY configs/setup-ollama.sh /opt/method/${CLI_NAME}/var/conf/setup-ollama.sh
RUN chmod +x /opt/method/${CLI_NAME}/var/conf/setup-ollama.sh
#RUN /opt/method/${CLI_NAME}/var/conf/setup-ollama.sh

# Startup script to start ollama serve and then do aiassess
COPY configs/startup.sh /opt/method/${CLI_NAME}/service/bin/startup.sh
RUN chmod +x /opt/method/${CLI_NAME}/service/bin/startup.sh

RUN \
  adduser --disabled-password --gecos '' method && \
  chown -R method:method /opt/method/${CLI_NAME}/ && \
  chown -R method:method /mnt/output

USER method

WORKDIR /opt/method/${CLI_NAME}/

ENV PATH="/opt/method/${CLI_NAME}/service/bin:${PATH}"
ENTRYPOINT [ "/opt/method/aiassess/service/bin/startup.sh" ]