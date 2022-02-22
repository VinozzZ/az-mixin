package az

import (
	"get.porter.sh/mixin/az/pkg"
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/pkgmgmt"
	"get.porter.sh/porter/pkg/porter/version"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	return version.PrintVersion(m.Context, opts, m.Version(opts))
}

func (m *Mixin) Version(opts version.Options) mixin.Metadata {
	return mixin.Metadata{
		Name: "az",
		VersionInfo: pkgmgmt.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Porter Authors",
		},
	}
}
