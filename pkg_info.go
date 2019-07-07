package main

import (
	"log"
)

type PackageInfo struct {
	Dependencies map[string]struct {
		Version string `yaml:"version"`
	} `yaml:"dependencies"`
}

type pkgInfo struct {
	info PackageInfo
}

var gPackageInfo = &pkgInfo{}

func GetPackInfo() *pkgInfo {
	return gPackageInfo
}

func init() {
	LoadYaml("package.yaml", &gPackageInfo.info)
	log.Println(gPackageInfo.info)
}
