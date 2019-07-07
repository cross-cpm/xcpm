package main

type packDownloader struct {
}

var gDownloader = &packDownloader{}

func GetDownloader() *packDownloader {
	return gDownloader
}
