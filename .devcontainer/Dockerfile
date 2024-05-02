FROM ubuntu:24.04

ARG USERNAME=ubuntu
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Ensure apt is in non-interactive to avoid prompts
ENV DEBIAN_FRONTEND=noninteractive
ENV PIP_BREAK_SYSTEM_PACKAGES 1

# Create the user
RUN apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Install tools.
RUN apt update && apt -y install --no-install-recommends \
    ca-certificates curl wget git openssh-client locales build-essential \
    python3 python3-pip unzip jq \
    && locale-gen en_US.UTF-8

# Install GO
RUN wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz \
    && rm -rf go1.22.2.linux-amd64.tar.gz

ENV PATH /usr/local/go/bin:$PATH

USER $USERNAME
ENV HOME /home/$USERNAME
WORKDIR $HOME

# vscode go extension dependencies
RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest

# Install pre-commit go dependencies
RUN go install github.com/terraform-docs/terraform-docs@v0.16.0 \
    && go install github.com/aquasecurity/tfsec/cmd/tfsec@latest
# ENV PATH $HOME/go/bin:$PATH

# Install pre-commit and it's pip dependencies
RUN pip install pre-commit checkov
# ENV PATH $HOME/.local/bin:$PATH

# Install tenv
RUN LATEST_VERSION=$(curl --silent https://api.github.com/repos/tofuutils/tenv/releases/latest | jq -r .tag_name) \
    && curl -O -L "https://github.com/tofuutils/tenv/releases/latest/download/tenv_${LATEST_VERSION}_amd64.deb" \
    && sudo dpkg -i "tenv_${LATEST_VERSION}_amd64.deb"

# Install pre-commit dependency terrascan
RUN curl -L "$(curl -s https://api.github.com/repos/tenable/terrascan/releases/latest | grep -o -E "https://.+?_Linux_x86_64.tar.gz")" > terrascan.tar.gz \
    && tar -xf terrascan.tar.gz terrascan && rm terrascan.tar.gz \
    && install -D terrascan $HOME/bin && rm terrascan

# Install pre-commit dependency tflint
RUN curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh -O \
    && bash install_linux.sh && rm install_linux.sh

ENV PATH /usr/bin:$HOME/.local/bin:$HOME/go/bin:$HOME/bin:$PATH
