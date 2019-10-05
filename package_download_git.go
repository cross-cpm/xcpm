package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/codeskyblue/go-sh"
)

type packageGitDownloader struct {
	url string
}

func NewPackageGitDownloader(source *PackageSourceInfo) *packageGitDownloader {
	log.Println("source repo:", source.Git)
	return &packageGitDownloader{
		url: source.Git,
	}
}

func (d *packageGitDownloader) Download(version string) error {
	path, err := d.getPackagePath(version)
	if err != nil {
		return err
	}

	if FileExist(path) {
		return nil
	}

	err = sh.Command("git", "clone", "--depth", "1", "-b", version, d.url, path).Run()
	if err != nil {
		return err
	}

	return nil
}

func (d *packageGitDownloader) getPackagePath(version string) (string, error) {
	// $HOME/.xcpm/cache/github.com/libjpeg-turbo/libjpeg-turbo.git/<tag>/
	u, err := url.Parse(d.url)
	if err != nil {
		return "", err
	}

	// log.Println("url:", u, err)
	// log.Println(u.Host)
	// log.Println(u.Path)
	home, err := Home()
	if err != nil {
		return "", err
	}

	cache_path := fmt.Sprintf("%s/%s/%s%s/%s/", home, GLOBAL_CACHE_PATH, u.Host, u.Path, version)
	log.Println("cache path:", cache_path)
	return cache_path, nil
}

func (d *packageGitDownloader) GetCache(version string) (*CacheInfo, error) {
	path, err := d.getPackagePath(version)
	if err != nil {
		return nil, err
	}

	return &CacheInfo{Type: "repo", Path: path}, nil
}
