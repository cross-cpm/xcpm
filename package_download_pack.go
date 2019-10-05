package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/codeskyblue/go-sh"
)

type packageTarDownloader struct {
	url string
}

func NewPackageTarDownloader(source *PackageSourceInfo) *packageTarDownloader {
	log.Println("source repo:", source.Pack)
	return &packageTarDownloader{
		url: source.Pack,
	}
}

func (d *packageTarDownloader) Download(version string) error {
	path, err := d.getPackagePath(version)
	if err != nil {
		return err
	}

	if FileExist(path) {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	log.Println("command:", "wget", "-O", path, d.url)
	err = sh.Command("wget", "-O", path, d.url).Run()
	if err != nil {
		return err
	}

	return nil
}

func (d *packageTarDownloader) getPackagePath(version string) (string, error) {
	// $HOME/.xcpm/cache/github.com/open-source-parsers/jsoncpp/archive/1.9.1.tar.gz
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

	cache_path := fmt.Sprintf("%s/%s/%s%s", home, GLOBAL_CACHE_PATH, u.Host, u.Path)
	log.Println("cache path:", cache_path)
	return cache_path, nil
}

func (d *packageTarDownloader) GetCache(version string) (*CacheInfo, error) {
	path, err := d.getPackagePath(version)
	if err != nil {
		return nil, err
	}

	return &CacheInfo{Type: "pack", Path: path}, nil
}
