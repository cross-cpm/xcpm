package main

type CacheInfo struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

type PackageDownloader interface {
	Download(version string) error
	GetCache(version string) (*CacheInfo, error)
}

func NewPackageDownloader(source *PackageSourceInfo) PackageDownloader {
	return NewPackageGitDownloader(source)
}
