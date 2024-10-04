# Microsoft Go Infrastructure Docker Images

The Dockerfiles in this repository are used by Microsoft to produce Docker images that support building and testing the Go compiler inside Microsoft infrastructure.

## Prerequisites

* [PowerShell 6+](https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell)

## Building the Docker images

To build a specific docker image, run `docker build .` in a directory that contains a Dockerfile.

To build all Docker tags in this repository, run `pwsh build.ps1` in the root of the repository.

## Using the images in CI

A list of images is maintained in the [images.md](./images.md) file.

## Updating a Dockerfile

After adding or removing any Dockerfile, run `go run ./cmd/geninfra` in the root of the repository.
This regenerates files such as the `manifest.json` file according to directory naming conventions.
The `manifest.json` file is used by the `build.ps1` build process and CI builds.

As long as your Dockerfile matches the conventions we expect in this repository, the command will succeed.
If it doesn't succeed, you may need to move the Dockerfile or edit the command.

Once the PR is merged, wait for the next rolling build, or check that your commit has been mirrored to the [internal repository](https://dev.azure.com/dnceng/internal/_git/microsoft-go-infra-images) and run [the pipeline](https://dev.azure.com/dnceng/internal/_build?definitionId=1170) manually.

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
