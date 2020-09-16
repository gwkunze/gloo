package common

import (
	"sort"

	"github.com/hashicorp/go-version"
	"github.com/rotisserie/eris"
)

type PluginVersion string

type PluginVersionList []PluginVersion

func (p PluginVersionList) Len() int {
	return len(p)
}

func (p PluginVersionList) Less(i, j int) bool {
	vi, err := version.NewVersion(string(p[i]))
	if err != nil {
		return true
	}
	vj, err := version.NewVersion(string(p[j]))
	if err != nil {
		return true
	}
	return vi.LessThan(vj)
}

func (p PluginVersionList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type PlatformDownload map[string]string

// VersionedPlugins is a map from version to map from binary name to download url.
type VersionedPlugins map[PluginVersion]PlatformDownload

func (p VersionedPlugins) ListVersions() PluginVersionList {
	return p.listDescending()
}

func (p VersionedPlugins) Latest() (PluginVersion, error) {
	versions := p.listDescending()
	if len(versions) > 0 {
		return versions[0], nil
	}
	return "", eris.New("There are no releases for this plugin.")
}

func (p VersionedPlugins) listDescending() PluginVersionList {
	var versions PluginVersionList
	for version := range p {
		versions = append(versions, version)
	}
	sort.Reverse(versions)
	return versions
}
