package main

import "log"

type PackageBuiler interface {
	Build() error
	GetPath() (string, error)
}

func NewPackageBuilder(pkgName string, toolchain string, bi []PackageBuildInfo) PackageBuiler {

	// TODO: choose one from build info array by toolchain
	info := &bi[0]
	log.Println("build type:", info.Type)

	if info.Type == "cmake" {
		return NewPackageCMakeBuiler(pkgName, toolchain, info)
	} else if info.Type == "autoconf" {
		return NewPackageAutoconfBuilder(pkgName, toolchain, info)
	}

	return nil
}
