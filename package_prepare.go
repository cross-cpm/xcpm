package main

type PackagePreparer interface {
	Prepare() error
}

func NewPackagePrepare(cache *CacheInfo, buildPath string) PackagePreparer {
	if cache.Type == "repo" {
		return NewRepoPackagePrepare(cache.Path, buildPath)
	}

	return nil
}
