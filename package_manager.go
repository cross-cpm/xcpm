package main

import (
	"fmt"
	"log"
)

type PackageInfo struct {
	Source []PackageSourceInfo `yaml:"source"`
	Build  []PackageBuildInfo  `yaml:"build"`
}

type PackageBuildInfo struct {
	Type string `yaml:"type"`
	// Command string `yaml:"command"`
}

type PackageSourceInfo struct {
	Version string `yaml:"version"`
	Git     string `yaml:"git"`
	Pack    string `yaml:"pack"`
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

func (p *packageManager) GetBuildInfo(version string) ([]PackageBuildInfo, error) {
	// TODO: match by version

	return p.info.Build, nil
}
