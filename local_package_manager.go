package main

import (
	"log"
)

type DependencyInfo struct {
	Version string `yaml:"version"`
}

type DependenciesMap map[string]*DependencyInfo

type LocalPackageInfo struct {
	Dependencies DependenciesMap `yaml:"dependencies"`
}

type localPackageManager struct {
	info LocalPackageInfo
}

var gLocalPackageManager = &localPackageManager{}

func GetLocalPackageManager() *localPackageManager {
	return gLocalPackageManager
}

func init() {
	LoadYaml("package.yaml", &gLocalPackageManager.info)
	log.Println(gLocalPackageManager.info)
}

func (p *localPackageManager) GetDependencies() (DependenciesMap, error) {
	return p.info.Dependencies, nil
}

func (p *localPackageManager) GetDependencyInfo(pkgName string) (*DependencyInfo, error) {
	return p.info.Dependencies[pkgName], nil
}
