FROM cblmariner.azurecr.io/base/core:1.0

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
        openssl-devel \
        powershell \
    ; \
    tdnf clean all
