FROM cblmariner.azurecr.io/base/core:1.0.20211027

# Install build tools.
RUN tdnf install -y \
        ca-certificates-microsoft \
        build-essential \
    && tdnf clean all

# Available OpenSSL versions at https://github.com/openssl/openssl/tags.
ARG OPENSSL_VER="1.0.2u"

# Download and decompress OpenSSL tarball.
RUN [[ ${OPENSSL_VER} = 1.* ]] && openssl_name="OpenSSL_${OPENSSL_VER//./_}" || openssl_name="openssl-${OPENSSL_VER}" \
    && wget https://github.com/openssl/openssl/archive/refs/tags/${openssl_name}.tar.gz -O - | tar -xz -C /usr/local/src \
    && mv /usr/local/src/openssl-${openssl_name} /usr/local/src/openssl

# Build OpenSSL.
# Instructions taken from https://wiki.openssl.org/index.php/Compilation_and_Installation.
RUN cd /usr/local/src/openssl \
    && ./config shared \
    && make build_libs

# Copy OpenSSL shared libraries and include files into its finall location.
RUN cp -r /usr/local/src/openssl/include/openssl /usr/include \
    && cp -H /usr/local/src/openssl/libcrypto.so /usr/lib/libcrypto.so.${OPENSSL_VER}
