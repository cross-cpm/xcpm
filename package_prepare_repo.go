package main

import "log"

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
	return nil
}
