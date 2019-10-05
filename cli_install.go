package main

func doCliInstallAll(toolchain string) error {
	// package.yaml 文件解析
	lpm := GetLocalPackageManager()
	dependencies, err := lpm.GetDependencies()
	if err != nil {
		return err
	}

	// 遍历依赖包信息，下载并构建依赖包
	for pkg, info := range dependencies {
		err = doCliDownload(pkg, info.Version)
		if err != nil {
			continue
		}

		err = doCliBuild(pkg, toolchain)
		if err != nil {
			continue
		}
	}

	return nil
}

func doCliInstall(pkg, toolchain string) error {
	// package.yaml 文件解析
	lpm := GetLocalPackageManager()
	info, err := lpm.GetDependencyInfo(pkg)
	if err != nil {
		return err
	}

	// 下载并构建依赖包
	err = doCliDownload(pkg, info.Version)
	if err != nil {
		return err
	}

	err = doCliBuild(pkg, toolchain)
	if err != nil {
		return err
	}

	return nil
}
