package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codeskyblue/go-sh"
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

	codePath := fmt.Sprintf("%s/usr/src/%s", prefixRootPath, b.pkgName)
	log.Println("codePath:", codePath)
	return codePath, nil
}

func getPrefixRootPath(toolchain string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	prefix := "prefix_root"
	if toolchain != "" {
		prefix = toolchain + "_prefix_root"
	}

	return fmt.Sprintf("%s/.packages/%s", pwd, prefix), nil
}

func (b *packageCMakeBuiler) Build() error {
	log.Println("cmake build ...")

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
	s.Call("bash", "-c", "cmake -DCMAKE_INSTALL_PREFIX=$PREFIX_ROOT .")
	s.Call("make")
	s.Call("make", "install")

	return nil
}
