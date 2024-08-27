FROM ubuntu:20.04

# Based on https://github.com/dotnet/dotnet-buildtools-prereqs-docker/blob/c4da7a7d44bb3e9c177d1a65127010f777b8987c/src/ubuntu/20.04/Dockerfile
# Then modified to remove some unnecessary tools and include dependencies for Go long tests.

# Install the base toolchain we need to build anything (clang, cmake, make and the like).
RUN apt-get update \
    && apt-get install -y \
        cmake \
        clang-12 \
        gdb \
        liblldb-12-dev \
        lldb-12 \
        llvm-12 \
        locales \
        make \
        sudo \
        wget \
    && rm -rf /var/lib/apt/lists/*

# Install tools used by the AzDO build automation.
RUN apt-get update \
    && apt-get install -y \
        git \
        nodejs \
        npm \
        tar \
        zip \
        curl \
    && rm -rf /var/lib/apt/lists/*

# Install Go long test prerequisites.
RUN apt-get update \
    && apt-get install -y \
        mercurial \
    && rm -rf /var/lib/apt/lists/*
