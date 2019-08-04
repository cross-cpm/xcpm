package main

import (
	"fmt"
	"log"
)

type PackageBuiler interface {
	Build(toolchain string) error
}

type PackageBuildInfo struct {
}

func NewPackageBuilder(pkgName string) PackageBuiler {
	bi, err := loadBuildYaml(pkgName)
	if err != nil {
		log.Println("load build info error!", err)
		return nil
	}

	return NewPackageCMakeBuiler(pkgName, bi)
}

func loadBuildYaml(pkgName string) (*PackageBuildInfo, error) {
	// load: .packages/build/<package_name>.yaml
	filename := fmt.Sprintf(".packages/build/%s.yaml", pkgName)
	bi := &PackageBuildInfo{}
	err := LoadYaml(filename, bi)
	if err != nil {
		return nil, err
	}

	return bi, nil
}
