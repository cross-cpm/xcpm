package main

import "log"

func doCliUpdate() {
	pkgName := "libjpeg-turbo"
	pkgLib := NewPackageLib(pkgName)
	log.Println(pkgLib)
}
