# `eng`: build infrastructure

This directory contains build infrastructure files used to build the prerequisites images that Microsoft uses to build Go.

The directory name, "eng", is short for "engineering". This name is used because it matches the engineering directory used in microsoft/go, and also because auto-updates to the "eng/common" directory only work with this absolute location.

The `common` directory is copied from https://github.com/dotnet/dotnet-docker/tree/main/eng/common when an update is necessary.
