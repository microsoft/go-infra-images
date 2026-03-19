FROM mcr.microsoft.com/cbl-mariner/base/core:2.0

# Install dependencies needed to build, test, and longtest Go.
RUN set -eux; \
    # Upgrade all packages per https://eng.ms/docs/more/containers-secure-supply-chain/updating
	tdnf update -y; \
    tdnf install -y \
        binutils \
        util-linux \
        ca-certificates-microsoft \
        gcc \
        git \
        glibc \
        glibc-devel \
        iana-etc \
        kernel-headers \
        # Pin powershell: 7.6.0 requires dotnet-runtime-10.0, which isn't available for CBL-Mariner 2.0.
        powershell-7.5.5 \
        sudo \
        tar \
    ; \
    tdnf clean all
