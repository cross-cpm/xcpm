package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/codeskyblue/go-sh"
)

type tarPackagePrepare struct {
	cachePath string
	buildPath string
}

func NewTarPackagePrepare(cachePath string, buildPath string) *tarPackagePrepare {
	return &tarPackagePrepare{
		cachePath: cachePath,
		buildPath: buildPath,
	}
}

func (p *tarPackagePrepare) Prepare() error {
	log.Printf("extract code from cache(%s) to build path(%s)!\n", p.cachePath, p.buildPath)

	buildParentPath := filepath.Dir(p.buildPath)
	tmpBuildPath := filepath.Join(buildParentPath, "tmp", filepath.Base(p.buildPath))

	err := os.RemoveAll(tmpBuildPath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(tmpBuildPath, os.ModePerm)
	if err != nil {
		return err
	}

	// TODO: 改为使用 shutil 工具包，以实现跨平台
	s := sh.NewSession()
	s.ShowCMD = true
	s.SetDir(tmpBuildPath)
	s.Call("tar", "xf", p.cachePath)

	tarTopDir, err := findTarTopDir(tmpBuildPath)
	if err != nil {
		return err
	}

	err = os.RemoveAll(p.buildPath)
	if err != nil {
		return err
	}

	err = os.Rename(tarTopDir, p.buildPath)
	if err != nil {
		return err
	}

	return nil
}

// 找到TAR包的顶层目录
func findTarTopDir(dir string) (string, error) {
	subdirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	count := 0
	lastdir := ""
	for _, subdir := range subdirs {
		// 如果含有文件，则 TopDir 就是传入的顶层目录
		if !subdir.IsDir() {
			break
		}
		count = count + 1
		lastdir = subdir.Name()
	}

	if count == 1 {
		// 如果只含有一个子目录，则认为 TopDir 是这个子目录
		return filepath.Join(dir, lastdir), nil
	}

	return dir, nil
}
