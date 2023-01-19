package route

import (
	"net/http"
	"path"
	fp "path/filepath"
	"strings"

	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/router"

	. "github.com/starudream/github-asset-mirror/config"
	"github.com/starudream/github-asset-mirror/github"

	"github.com/starudream/github-asset-mirror/internal/osx"
)

func mirror(c *router.Context) {
	var (
		ver  = strings.ToLower(c.Param("ver"))
		os   = strings.ToLower(c.Param("os"))
		arch = strings.ToLower(c.Param("arch"))
	)

	if ver == "" || ver == "latest" {
		release, err := github.GetLatestRelease()
		if err != nil {
			log.Ctx(c).Error().Msgf("get latest release failed: %v", err)
			c.JSON(http.StatusBadRequest, router.ErrBadRequest)
			return
		}
		ver = *release.Name
	}

	if os == "" {
		os = "linux"
	}
	if arch == "" {
		arch = "amd64"
	}

	data := map[string]any{
		"ver":      ver,
		"os":       os,
		"arch":     arch,
		"platform": "client",
	}

	for k, vs := range c.AllQuery() {
		if len(vs) > 0 {
			data[k] = vs[0]
		}
	}

	url, err := FormatURL(data)
	if err != nil {
		log.Ctx(c).Error().Msgf("format url failed: %v", err)
		c.JSON(http.StatusInternalServerError, router.ErrInternal)
		return
	}

	filename := path.Base(url)
	filepath := fp.Join(C.Storage, filename)

	if !osx.ExistFile(filepath) {
		_, err = osx.SaveFile(osx.ProxyURL(C.Proxy, url), filepath)
		if err != nil {
			log.Ctx(c).Error().Msgf("download file failed: %v", err)
			c.JSON(http.StatusInternalServerError, router.ErrInternal)
			return
		}
	}

	c.ATTACHMENT(filepath, filename)
}
