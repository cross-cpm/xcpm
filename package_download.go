package main

import "log"

type CacheInfo struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

type PackageDownloader interface {
	Download(version string) error
	GetCache(version string) (*CacheInfo, error)
}

func NewPackageDownloader(source *PackageSourceInfo) PackageDownloader {
	log.Println("package source", source)
	if source.Git != "" {
		return NewPackageGitDownloader(source)
	} else if source.Url != "" {
		return NewPackageTarDownloader(source)
	}

	return nil
}
