// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Geninfraimagesmanifest generates the manifest.json file to let .NET Docker infrastructure build
// this repository. It searches "src" for Dockerfiles and parses the directory path to determine how
// each Dockerfile should be built and tagged.
//
// A Dockerfile path looks like "src/{distro}/{version}/{arch}/{name...}", where "name..." includes
// the remaining path with "/" changed to "-". Distros and arches have special handling included
// in this file, and the special handling is not intended to proactively handle future cases.
// This command should be edited when new requirements and conventions are necessary.
package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/microsoft/go-infra/buildmodel/dockermanifest"
	"github.com/microsoft/go-infra/stringutil"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
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
		case "cbl-mariner", "debian", "ubuntu":
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

	return stringutil.WriteJSONFile("manifest.json", m)
}
