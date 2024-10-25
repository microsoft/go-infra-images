FROM mcr.microsoft.com/cbl-mariner/base/core:2.0

# Install dependencies needed to build, test, and longtest Go.
RUN set -eux; \
    # Upgrade all packages per https://eng.ms/docs/more/containers-secure-supply-chain/updating
	tdnf update -y; \
    tdnf install -y \
        binutils \
        ca-certificates-microsoft \
        gcc \
        git \
        glibc \
        glibc-devel \
        iana-etc \
        kernel-headers \
        mercurial \
        powershell \
        sudo \
    ; \
    tdnf clean all
