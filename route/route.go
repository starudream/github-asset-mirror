package route

import (
	"net/http"

	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/router"

	. "github.com/starudream/github-asset-mirror/config"
)

func Register() {
	router.Handle(http.MethodGet, "/", index)
	router.Handle(http.MethodPost, "/", index)

	router.Handle(http.MethodGet, "/:ver", mirror)
	router.Handle(http.MethodGet, "/:ver/:os/:arch", mirror)

	_, _ = FormatURL(nil)
}

func index(c *router.Context) {
	c.JSONOK(router.M{"version": constant.VERSION, "bidtime": constant.BIDTIME})
}
