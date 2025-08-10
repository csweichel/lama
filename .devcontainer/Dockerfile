FROM mcr.microsoft.com/devcontainers/go:1.24.6

# Install build tools required for lama development
RUN cd /tmp && \
    # Install GoReleaser v0.108.0 (version from legacy .gitpod.yml)
    curl -OL https://github.com/goreleaser/goreleaser/releases/download/v0.108.0/goreleaser_Linux_x86_64.tar.gz && \
    tar xfz goreleaser_Linux_x86_64.tar.gz && \
    mv goreleaser /usr/local/bin/ && \
    rm goreleaser_Linux_x86_64.tar.gz && \
    # Install UPX v3.95 (version from legacy .gitpod.yml)
    curl -OL https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz && \
    tar xf upx-3.95-amd64_linux.tar.xz && \
    mv upx-3.95-amd64_linux/upx /usr/local/bin/ && \
    rm -rf upx-3.95-amd64_linux* && \
    # Verify installations
    goreleaser --version && \
    upx --version
