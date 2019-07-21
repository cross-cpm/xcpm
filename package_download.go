package main

import "log"

type packageDownloader struct {
}

var gDownloader = &packageDownloader{}

func GetPackageDownloader() *packageDownloader {
	return gDownloader
}

func (d *packageDownloader) Download(source *PackageSourceInfo) error {
	log.Println("source repo:", source.Git)
	return nil
}
