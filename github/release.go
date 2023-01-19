package github

import (
	"context"
	"time"

	"github.com/google/go-github/v49/github"

	"github.com/starudream/go-lib/log"

	. "github.com/starudream/github-asset-mirror/config"
)

var (
	xLatestReleaseKey    = "__LATEST__"
	xLatestReleaseExpire = time.Minute
)

func GetLatestRelease() (*github.RepositoryRelease, error) {
	if mc.Has(xLatestReleaseKey) {
		v, err := mc.Get(xLatestReleaseKey)
		if err == nil && v != nil {
			x := v.(*github.RepositoryRelease)
			log.Debug().Msgf("github: get latest release from cache, %s", *x.Name)
			return x, nil
		}
	}
	release, _, err := cli.Repositories.GetLatestRelease(context.Background(), C.Owner, C.Repo)
	if err == nil {
		err = mc.SetWithExpire(xLatestReleaseKey, release, xLatestReleaseExpire)
	}
	return release, err
}
