// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/microsoft/go-infra/buildmodel/dockermanifest"
	"github.com/microsoft/go-infra/stringutil"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "geninfra",
	RunE: run,
	Long: `Generates files in this repo based on the Dockerfiles in src.

- manifest.json: lets .NET Docker infrastructure build this repository.
- images.md: a list of the images in a format that a dev can more easily read and use.

Dockerfile paths are parsed using this convention:

  src/{distro}/{version}/{arch}/{name...}

"name..." may include more "/". They are treated as "-" if needed to produce a tag name.
Distros and arches have special handling, so ./cmd/geninfra/main.go may need to be adjusted when
new platforms or special cases are added.
	`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	var images []*dockermanifest.Image

	err := filepath.WalkDir("src", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		pathParts := strings.Split(path, string(filepath.Separator))
		if len(pathParts) < 6 {
			log.Printf("Not enough path parts to apply convention, skipping: %v\n", pathParts)
			return nil
		}
		distro := pathParts[1]
		version := pathParts[2]
		arch := pathParts[3]
		name := strings.Join(pathParts[4:len(pathParts)-1], "-")

		var os string
		switch distro {
		case "cbl-mariner", "debian", "ubuntu", "azurelinux":
			os = "linux"
		default:
			log.Printf("Didn't recognize distro, update %v: %v\n", "cmd/gen/main.go", distro)
			return nil
		}

		var archVariant string
		switch arch {
		case "armv7":
			arch = "arm"
			archVariant = "v7"
		}

		tag := distro + "-" + version + "-" + arch + archVariant + "-" + name

		images = append(images, &dockermanifest.Image{
			SharedTags: make(map[string]dockermanifest.Tag),
			Platforms: []*dockermanifest.Platform{
				{
					Architecture: arch,
					Variant:      archVariant,
					Dockerfile:   filepath.ToSlash(filepath.Dir(path)),
					OS:           os,
					OSVersion:    version,
					Tags: map[string]dockermanifest.Tag{
						// Pinned tag to use in e.g microsoft/go CI, for stability.
						tag + "-$(System:TimeStamp)-$(System:DockerfileGitCommitSha)": struct{}{},
						// Floating tag to use for build dependencies between Dockerfiles in this
						// repo. The .NET Docker infrastructure matches up this tag with "FROM x"
						// statements to determine the build order.
						tag: struct{}{},
					},
				},
			},
		})
		return nil
	})
	if err != nil {
		return err
	}

	m := &dockermanifest.Manifest{
		Readme:    map[string]string{"path": "README.md"},
		Registry:  "mcr.microsoft.com",
		Variables: make(map[string]interface{}),
		Includes:  make([]string, 0),
		Repos: []*dockermanifest.Repo{
			{
				ID:     "infra-images",
				Name:   "microsoft-go/infra-images",
				Images: images,
			},
		},
	}

	if err := stringutil.WriteJSONFile("manifest.json", m); err != nil {
		return err
	}

	return writeImagesMD(m)
}

func writeImagesMD(m *dockermanifest.Manifest) error {
	var b strings.Builder
	fmt.Fprint(&b, `<!-- This file was generated by 'go run ./cmd/geninfra'. Do not edit manually. -->

# Go infra images list

The following list has been generated based on the manifest.json file in this repository.
The full URIs have been calculated here to make the tags easy to use.
If a build hasn't occurred yet, the list may be out of date.

For an accurate but harder to use list of currently available tags, see [the list API](https://mcr.microsoft.com/v2/microsoft-go/infra-images/tags/list).

`)

	for _, repo := range m.Repos {
		for _, image := range repo.Images {
			for _, platform := range image.Platforms {
				for tag := range platform.Tags {
					if strings.Contains(tag, "$") {
						continue
					}
					fmt.Fprintf(&b, "`%v` ([%v](./%v/Dockerfile))\n", tag, platform.Dockerfile, platform.Dockerfile)
					fmt.Fprintf(&b, "```\n%v\n```\n\n", m.Registry+"/"+repo.Name+":"+tag)
				}
			}
		}
	}
	return os.WriteFile("images.md", []byte(b.String()), 0o666)
}
