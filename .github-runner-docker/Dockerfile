# Self-Hosted GitHub Actions Runner with All Dependencies
#
# This Dockerfile creates a fully self-contained runner with:
# - GitHub Actions runner
# - Go 1.23
# - Terraform CLI
# - All required build tools
#
# Build: docker build -t f5xc-runner .
# Run:   docker-compose up -d

FROM ubuntu:22.04

# Versions - update these as needed
ARG RUNNER_VERSION=2.329.0
ARG GO_VERSION=1.24.2
ARG TERRAFORM_VERSION=1.9.8
ARG GH_VERSION=2.97.0
ARG RUNNER_UID=1001
# TARGETARCH is automatically set by Docker buildx (amd64, arm64)
ARG TARGETARCH

# Prevent interactive prompts
ENV DEBIAN_FRONTEND=noninteractive
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install base dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl=7.81.0-1ubuntu1.25 \
    wget=1.21.2-2ubuntu1.4 \
    git=1:2.34.1-1ubuntu1.17 \
    jq=1.6-2.1ubuntu3.2 \
    unzip=6.0-26ubuntu3.2 \
    zip=3.0-12build2 \
    ca-certificates=20260601~22.04.1 \
    gnupg=2.2.27-3ubuntu2.5 \
    lsb-release=11.1.0ubuntu4 \
    sudo=1.9.9-1ubuntu2.6 \
    libicu70=70.1-2 \
    libssl3=3.0.2-0ubuntu1.26 \
    && rm -rf /var/lib/apt/lists/*

# Install Go
RUN curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${TARGETARCH}.tar.gz" | tar -C /usr/local -xzf -
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/home/runner/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Install Terraform
RUN curl -fsSL "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_${TARGETARCH}.zip" -o terraform.zip \
    && unzip terraform.zip -d /usr/local/bin \
    && rm terraform.zip \
    && chmod +x /usr/local/bin/terraform

# Install GitHub CLI
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg \
    && chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
    && apt-get update \
    && apt-get install -y --no-install-recommends "gh=${GH_VERSION}" \
    && rm -rf /var/lib/apt/lists/*

# Create runner user
RUN useradd --uid "${RUNNER_UID}" -m -s /bin/bash runner \
    && echo "runner ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

# Setup runner directory
WORKDIR /home/runner/actions-runner

# Download and extract GitHub Actions runner
RUN curl -fsSL "https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-${TARGETARCH}-${RUNNER_VERSION}.tar.gz" | tar -xzf -

# Create tool cache directory for actions/setup-go and other actions
# This prevents permission errors when GitHub Actions try to cache tools
RUN mkdir -p /home/runner/actions-runner/_work/_tool \
    && chown -R runner:runner /home/runner

# Switch to runner user
USER ${RUNNER_UID}

# Verify installations
RUN go version && terraform version && gh --version

# Entrypoint script
COPY --chown=runner:runner entrypoint.sh /home/runner/entrypoint.sh
RUN chmod +x /home/runner/entrypoint.sh

ENTRYPOINT ["/home/runner/entrypoint.sh"]
