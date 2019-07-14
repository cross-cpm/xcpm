package main

type packDownloader struct {
}

var gDownloader = &packDownloader{}

func GetDownloader() *packDownloader {
	return gDownloader
}

func (d *packDownloader) Download(source *PackageLibSourceInfo) error {
	return nil
}
