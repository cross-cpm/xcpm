package main

import "fmt"

type PackageCacheBuildInfo struct {
	Cache *CacheInfo         `yaml:"cache"`
	Build []PackageBuildInfo `yaml:"build"`
}

type buildInfoManager struct {
	filename string
	info     *PackageCacheBuildInfo
}

func NewBuildInfoManager(pkgName string) *buildInfoManager {
	return &buildInfoManager{
		filename: fmt.Sprintf(".packages/build/%s.yaml", pkgName),
	}
}

func (b *buildInfoManager) Load() error {
	info := &PackageCacheBuildInfo{}
	err := LoadYaml(b.filename, info)
	if err != nil {
		return err
	}

	b.info = info
	return nil
}

func (b *buildInfoManager) Save(bi []PackageBuildInfo, cache *CacheInfo) error {
	// .packages/build/<package_name>.yaml
	return SaveYaml(b.filename, &PackageCacheBuildInfo{
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
