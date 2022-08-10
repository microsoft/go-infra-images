FROM ubuntu:18.04

# Based on https://github.com/dotnet/dotnet-buildtools-prereqs-docker/blob/5c627c7ceeb94a58033cdc927f349b6ae8ffe333/src/ubuntu/18.04/amd64/Dockerfile
# Then modified to remove some unnecessary tools and include dependencies for Go long tests.

# Install the base toolchain we need to build anything (clang, cmake, make and the like).
RUN apt-get update \
    && apt-get install -y \
        clang-3.9 \
        gdb \
        liblldb-3.9-dev \
        lldb-3.9 \
        llvm-3.9 \
        locales \
        make \
        python-lldb-3.9 \
        sudo \
        wget \
    && rm -rf /var/lib/apt/lists/* \
    && wget https://cmake.org/files/v3.10/cmake-3.10.2-Linux-x86_64.tar.gz \
    && tar -xf cmake-3.10.2-Linux-x86_64.tar.gz --strip 1 -C /usr/local \
    && rm cmake-3.10.2-Linux-x86_64.tar.gz

# Install tools used by the AzDO build automation.
RUN apt-get update \
    && apt-get install -y \
        git \
        nodejs \
        npm \
        tar \
        zip \
    && rm -rf /var/lib/apt/lists/*

# Install Go long test prerequisites.
RUN apt-get update \
    && apt-get install -y \
        mercurial \
    && rm -rf /var/lib/apt/lists/*
