package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type PackageCacheBuildInfo struct {
	Cache *CacheInfo         `yaml:"cache"`
	Build []PackageBuildInfo `yaml:"build"`
}

type buildInfoManager struct {
	path     string
	filename string
	info     *PackageCacheBuildInfo
}

func NewBuildInfoManager(pkgName string) *buildInfoManager {
	return &buildInfoManager{
		path:     ".packages/build/",
		filename: fmt.Sprintf("%s.yaml", pkgName),
	}
}

func (b *buildInfoManager) Load() error {
	info := &PackageCacheBuildInfo{}
	err := LoadYaml(filepath.Join(b.path, b.filename), info)
	if err != nil {
		return err
	}

	b.info = info
	return nil
}

func (b *buildInfoManager) Save(bi []PackageBuildInfo, cache *CacheInfo) error {
	err := os.MkdirAll(b.path, os.ModePerm)
	if err != nil {
		return err
	}

	// .packages/build/<package_name>.yaml
	return SaveYaml(filepath.Join(b.path, b.filename), &PackageCacheBuildInfo{
		Cache: cache,
		Build: bi,
	})
}

func (b *buildInfoManager) GetCache() (*CacheInfo, error) {
	if b.info == nil {
		return nil, fmt.Errorf("build info not load!")
	}

	return b.info.Cache, nil
}

func (b *buildInfoManager) GetBuidInfo() ([]PackageBuildInfo, error) {
	if b.info == nil {
		return nil, fmt.Errorf("build info not load!")
	}

	return b.info.Build, nil
}
