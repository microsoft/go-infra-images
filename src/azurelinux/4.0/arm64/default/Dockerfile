# mcr.microsoft.com/azurelinux/base/core:4.0 doesn't distribute a PowerShell package,
# see https://github.com/microsoft/azurelinux/pull/2799.
# Use the .NET SDK image as a base image instead, which includes PowerShell and
# the necessary dependencies to use it (i.e. .NET Runtime and .NET SDK).
FROM mcr.microsoft.com/dotnet/sdk:11.0-preview-azurelinux4.0

# Install dependencies needed to build, test, and longtest Go.
RUN set -eux; \
    # Upgrade all packages per https://eng.ms/docs/more/containers-secure-supply-chain/updating.
    dnf update -y; \
    dnf install -y \
        binutils \
        ca-certificates \
        gcc \
        git \
        glibc \
        glibc-devel \
        kernel-headers \
        setup \
        sudo \
    ; \
    dnf clean all
