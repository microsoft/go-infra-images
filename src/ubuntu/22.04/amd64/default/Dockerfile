FROM ubuntu.azurecr.io/ubuntu:22.04

# Install dependencies needed to build, test, and longtest Go.
RUN apt-get update \
    && apt-get install -y \
        ca-certificates \
        binutils \
        gcc \
        gdb \
        git \
        mercurial \
        sudo \
    && rm -rf /var/lib/apt/lists/*

# Install tools used to add the Microsoft package repository, necessary to install PowerShell.
RUN apt-get update \
    && apt-get install -y \
        curl \
        gnupg\
        apt-transport-https \
    && rm -rf /var/lib/apt/lists/*

# Add the Microsoft package repository.
RUN curl https://packages.microsoft.com/keys/microsoft.asc | apt-key add - \
    && echo "deb https://packages.microsoft.com/ubuntu/22.04/prod/ jammy main" > /etc/apt/sources.list.d/microsoft.list

# Install PowerShell.
RUN apt-get update \
    && apt-get install -y \
        powershell \
    && rm -rf /var/lib/apt/lists/*
