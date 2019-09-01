package main

import (
	"log"

	"github.com/termie/go-shutil"
)

type repoPackagePrepare struct {
	cachePath string
	buildPath string
}

func NewRepoPackagePrepare(cachePath string, buildPath string) *repoPackagePrepare {
	return &repoPackagePrepare{
		cachePath: cachePath,
		buildPath: buildPath,
	}
}

func (p *repoPackagePrepare) Prepare() error {
	log.Printf("copy code from cache(%s) to build path(%s)!\n", p.cachePath, p.buildPath)
	//shutil.RmTree(p.buildPath, nil)
	err := shutil.CopyTree(p.cachePath, p.buildPath, nil)
	if err != nil {
		return err
	}

	return nil
}
