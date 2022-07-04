FROM ubuntu:22.04

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Ensure apt is in non-interactive to avoid prompts
ENV DEBIAN_FRONTEND=noninteractive

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Install tools.
RUN apt update && apt -y install --no-install-recommends \
    ca-certificates curl wget git openssh-client locales build-essential \
    python3 python3-pip unzip \
    && locale-gen en_US.UTF-8

ENV DEBIAN_FRONTEND=dialog

# Install GO
RUN wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz \
    && rm -rf go1.18.3.linux-amd64.tar.gz

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

# Install tfswitch
RUN wget https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh \
    && bash install.sh -b $HOME/bin

# Install pre-commit dependency terrascan
RUN curl -L "$(curl -s https://api.github.com/repos/tenable/terrascan/releases/latest | grep -o -E "https://.+?_Linux_x86_64.tar.gz")" > terrascan.tar.gz \
    && tar -xf terrascan.tar.gz terrascan && rm terrascan.tar.gz \
    && install terrascan $HOME/bin && rm terrascan

# Install pre-commit dependency tflint
RUN curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh -O \
    && TFLINT_INSTALL_PATH=$HOME/bin bash install_linux.sh && rm install_linux.sh

ENV PATH $HOME/.local/bin:$HOME/go/bin:$HOME/bin:$PATH
