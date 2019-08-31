package main

import (
	"fmt"
	"log"
)

type packageCMakeBuiler struct {
	pkgName   string
	toolchain string
	buildInfo *PackageBuildInfo
}

func NewPackageCMakeBuiler(pkgName string, toolchain string, bi *PackageBuildInfo) *packageCMakeBuiler {
	return &packageCMakeBuiler{
		pkgName:   pkgName,
		toolchain: toolchain,
		buildInfo: bi,
	}
}

func (b *packageCMakeBuiler) GetPath() (string, error) {
	prefixRootPath, err := getPrefixRootPath(b.toolchain)
	if err != nil {
		return "", err
	}

	codePath := fmt.Sprintf("%s/usr/local/%s", prefixRootPath, b.pkgName)
	log.Println("codePath:", codePath)
	return codePath, nil
}

func getPrefixRootPath(toolchain string) (string, error) {
	if toolchain == "" {
		return ".packages/prefix_root", nil
	}

	return "", fmt.Errorf("unknown prefix root path")
}

func (b *packageCMakeBuiler) Build() error {
	log.Println("cmake build ...")
	return nil
}
