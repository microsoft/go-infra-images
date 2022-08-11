#!/usr/bin/env pwsh
# From https://github.com/dotnet/dotnet-buildtools-prereqs-docker/blob/c54d34b99929259dc16252692e7bbd0da781559c/build.ps1
[cmdletbinding()]
param(
    [string]$DockerfilePath = "*",
    [string]$ImageBuilderCustomArgs
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
pushd $PSScriptRoot
$ImageBuilderCustomArgs = "$ImageBuilderCustomArgs --var UniqueId=$(Get-Date -Format yyyyMMddHHmmss)"
try {
    ./eng/common/Invoke-ImageBuilder.ps1 "build --path '$DockerfilePath' $ImageBuilderCustomArgs"
}
finally {
    popd
}
