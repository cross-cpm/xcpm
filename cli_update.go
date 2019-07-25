package main

import "log"

func doCliDownload(pkgName string) error {
	lpm := GetLocalPackageManager()
	di, err := lpm.GetDependencyInfo(pkgName)
	//log.Println("dependency info:", di, err)
	if err != nil {
		return err
	}

	pm := NewPackageManager(pkgName)
	si, err := pm.GetSourceInfo(di.Version)
	//log.Println("package source info:", si)
	if err != nil {
		return err
	}

	pd := NewPackageDownloader(si)
	err = pd.Download(di.Version)
	if err != nil {
		return err
	}

	return nil
}

func doCliBuild(pkgName string) error {
	return nil
}

func doCliUpdate() error {
	pkgName := "libjpeg-turbo"
	pkgLib := NewPackageManager(pkgName)
	log.Println(pkgLib)
	return nil
}
