package main

import (
	"fmt"
	"log"
)

type packageCMakeBuiler struct {
	pkgName   string
	buildInfo *PackageBuildInfo
}

func NewPackageCMakeBuiler(pkgName string, bi *PackageBuildInfo) *packageCMakeBuiler {
	return &packageCMakeBuiler{
		pkgName:   pkgName,
		buildInfo: bi,
	}
}

func (b *packageCMakeBuiler) Build(toolchain string) error {
	log.Println("cmake build ...")
	prefixRootPath, err := getPrefixRootPath(toolchain)
	if err != nil {
		return err
	}

	codePath := fmt.Sprintf("%s/usr/local/%s", prefixRootPath, b.pkgName)
	log.Println("codePath:", codePath)
	return nil
}

func getPrefixRootPath(toolchain string) (string, error) {
	if toolchain == "" {
		return ".packages/prefix_root", nil
	}

	return "", fmt.Errorf("unknown prefix root path")
}

func (b *packageCMakeBuiler) prepare() {

}
