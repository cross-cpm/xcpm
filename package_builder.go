package main

type PackageBuiler interface {
	Build() error
	GetPath() (string, error)
}

func NewPackageBuilder(pkgName string, toolchain string, bi []PackageBuildInfo) PackageBuiler {

	// TODO: choose one from build info array
	info := &bi[0]

	return NewPackageCMakeBuiler(pkgName, toolchain, info)
}
