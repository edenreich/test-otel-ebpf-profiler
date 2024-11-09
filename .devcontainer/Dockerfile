FROM mcr.microsoft.com/vscode/devcontainers/base:0-ubuntu-22.04

ENV DEBIAN_FRONTEND=noninteractive

# Install essential packages
RUN apt-get update && apt-get -y install --no-install-recommends \
    apt-transport-https \
    zsh \
    git \
    curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Add kubernetes repository
RUN curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.31/deb/Release.key | gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg \
    && chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg \
    && echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.31/deb/ /' | tee /etc/apt/sources.list.d/kubernetes.list \
    && chmod 644 /etc/apt/sources.list.d/kubernetes.list \
    && curl https://baltocdn.com/helm/signing.asc | gpg --dearmor -o /usr/share/keyrings/helm.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list

# Install NodeJS, NPM, Go, Kubectl, Helm, Task, K3D and CLTPTL
RUN apt-get update && apt-get -y install --no-install-recommends \
    nodejs \
    npm \
    golang \
    kubectl \
    helm \
    && apt-get clean && rm -rf /var/lib/apt/lists/* \
    && npm install -g @go-task/cli \
    && curl -sSL https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash \
    && curl -sSL https://github.com/tilt-dev/ctlptl/releases/download/v0.8.35/ctlptl.0.8.35.linux.$(dpkg --print-architecture).tar.gz | tar -xz -C /usr/local/bin ctlptl

# Install Oh My Zsh with Powerlevel10k
COPY .devcontainer/p10k.zsh /home/vscode/.p10k.zsh
RUN git clone --depth=1 https://github.com/romkatv/powerlevel10k.git /home/vscode/.oh-my-zsh/custom/themes/powerlevel10k \
    && ln -sf /home/vscode/.oh-my-zsh/custom/themes/powerlevel10k/powerlevel10k.zsh-theme /home/vscode/.oh-my-zsh/custom/themes/codespaces.zsh-theme \
    && sed -i 's/ZSH_THEME="devcontainers"/ZSH_THEME="powerlevel10k\/powerlevel10k"/g' /home/vscode/.zshrc \
    && echo '[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh' >> /home/vscode/.zshrc


RUN chsh -s $(which zsh)

ENTRYPOINT [ "zsh" ]
