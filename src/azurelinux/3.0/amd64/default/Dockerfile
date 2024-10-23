# mcr.microsoft.com/azurelinux/base/core:3.0 doesn't distribute a PowerShell package,
# see https://github.com/microsoft/azurelinux/pull/2799.
# Use the .NET SDK image as a base image instead, which includes PowerShell and
# the necessary dependencies to use it (i.e. .NET Runtime and .NET SDK).
FROM mcr.microsoft.com/dotnet/sdk:9.0-azurelinux3.0-amd64

# Install dependencies needed to build, test, and longtest Go.
RUN set -eux; \
    # Upgrade all packages per https://eng.ms/docs/more/containers-secure-supply-chain/updating.
	tdnf update -y; \
    tdnf install -y \
        binutils \
        ca-certificates \
        gcc \
        git \
        glibc \
        glibc-devel \
        iana-etc \
        kernel-headers \
        mercurial \
    ; \
    tdnf clean all
