package main

import "log"

func doCliDownload(pkgName string) error {
	// package.yaml 文件解析
	lpm := GetLocalPackageManager()
	di, err := lpm.GetDependencyInfo(pkgName)
	//log.Println("dependency info:", di, err)
	if err != nil {
		return err
	}

	// 包描述文件加载器
	// 当前工程: packages/<package_name>.yaml
	// 全局: ~/.xcpm/packages/<package_name>.yaml
	pm := NewPackageManager(pkgName)
	si, err := pm.GetSourceInfo(di.Version)
	//log.Println("package source info:", si)
	if err != nil {
		return err
	}

	// 源码下载
	// 下载目录: ~/.xcpm/cache/github.com/libjpeg-turbo/libjpeg-turbo.git/<tag>/
	pd := NewPackageDownloader(si)
	err = pd.Download(di.Version)
	if err != nil {
		return err
	}

	// TODO: 编译构建信息写入文件
	//    .packages/build/<package_name>.yaml
	// err = pm.WriteBuildInfo()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func doCliBuild(pkgName, toolchain string) error {
	pb := NewPackageBuilder(pkgName)
	err := pb.Build(toolchain)
	if err != nil {
		return err
	}

	return nil
}

func doCliUpdate() error {
	pkgName := "libjpeg-turbo"
	pkgLib := NewPackageManager(pkgName)
	log.Println(pkgLib)
	return nil
}
