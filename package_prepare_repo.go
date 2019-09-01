package main

import (
	"log"
	"os/exec"

	"github.com/cross-cpm/go-shutil"
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
	//FIXME: use: shutil.RmTree(p.buildPath, nil)
	exec.Command("rm", "-fr", p.buildPath).Run()
	_, err := shutil.CopyTree(p.cachePath, p.buildPath, nil)
	if err != nil {
		return err
	}

	return nil
}
