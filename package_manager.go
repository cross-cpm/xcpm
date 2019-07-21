package main

import (
	"fmt"
	"log"
)

type PackageInfo struct {
	Source []PackageSourceInfo `yaml:"source"`
}

type PackageSourceInfo struct {
	Git string `yaml:"git"`
}

type packageManager struct {
	info PackageInfo
}

// 包描述文件加载器
func NewPackageManager(name string) *packageManager {
	filename, err := findPackageLibFile(name)
	log.Println("package lib filename:", name, filename)
	if err != nil {
		log.Fatal("find package lib failed!", name, err)
	}

	lib := &packageManager{}
	LoadYaml(filename, &lib.info)
	log.Println("lib info", lib.info)

	return lib
}

func findPackageLibFile(name string) (string, error) {
	filename := fmt.Sprintf("packages/%s.yaml", name)
	log.Println("filename:", filename)
	if FileExist(filename) {
		return filename, nil
	}

	// TODO: find in global packages lib

	return "", fmt.Errorf("package(%s) lib file not found!", name)
}

func (p *packageManager) GetSourceInfo(version string) (*PackageSourceInfo, error) {

	// TODO: match by version
	if len(p.info.Source) > 0 {
		return &p.info.Source[0], nil
	}

	return nil, fmt.Errorf("source not found")
}
