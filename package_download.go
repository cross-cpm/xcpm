package main

type PackageDownloader interface {
	Download(version string) error
}

func NewPackageDownloader(source *PackageSourceInfo) PackageDownloader {
	return NewPackageGitDownloader(source)
}
