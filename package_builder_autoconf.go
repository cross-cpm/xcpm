package main

import (
	"fmt"
	"log"

	"github.com/codeskyblue/go-sh"
)

type packageAutoConfBuiler struct {
	pkgName   string
	toolchain string
	buildInfo *PackageBuildInfo
}

func NewPackageAutoconfBuilder(pkgName string, toolchain string, bi *PackageBuildInfo) *packageAutoConfBuiler {
	return &packageAutoConfBuiler{
		pkgName:   pkgName,
		toolchain: toolchain,
		buildInfo: bi,
	}
}

func (b *packageAutoConfBuiler) GetPath() (string, error) {
	prefixRootPath, err := getPrefixRootPath(b.toolchain)
	if err != nil {
		return "", err
	}

	codePath := fmt.Sprintf("%s/usr/src/%s", prefixRootPath, b.pkgName)
	log.Println("codePath:", codePath)
	return codePath, nil
}

func (b *packageAutoConfBuiler) Build() error {
	log.Println("autoconf build ...")

	prefix, err := getPrefixRootPath(b.toolchain)
	if err != nil {
		return err
	}

	workdir, err := b.GetPath()
	if err != nil {
		return err
	}

	if !FileExist(workdir) {
		return fmt.Errorf("workdir not found!")
	}

	s := sh.NewSession()
	s.ShowCMD = true
	s.SetDir(workdir)
	s.SetEnv("PREFIX_ROOT", prefix)

	configs := make([]interface{}, 0)
	configs = append(configs, "./configure")
	configs = append(configs, fmt.Sprintf("--prefix=%s", prefix))
	for _, v := range b.buildInfo.Configure {
		configs = append(configs, v)
	}
	s.Call("sh", configs...)
	s.Call("make")
	s.Call("make", "install")

	return nil
}
