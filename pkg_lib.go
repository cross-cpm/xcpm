package main

import (
	"fmt"
	"log"
)

type PackageLibInfo struct {
	Source []PackageLibSourceInfo `yaml:"source"`
}

type PackageLibSourceInfo struct {
	Git string `yaml:"git"`
}

type pkgLib struct {
	info PackageLibInfo
}

func NewPackageLib(name string) *pkgLib {
	filename, err := findPackageLibFile(name)
	log.Println("package lib filename:", name, filename)
	if err != nil {
		log.Fatal("find package lib failed!", name, err)
	}

	lib := &pkgLib{}
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
