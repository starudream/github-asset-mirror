package route

import (
	"net/http"

	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/router"

	. "github.com/starudream/github-asset-mirror/config"
)

func Register() {
	router.Handle(http.MethodGet, "/_health", health)

	router.Handle(http.MethodGet, "/", mirror)
	router.Handle(http.MethodGet, "/:ver", mirror)
	router.Handle(http.MethodGet, "/:ver/:os", mirror)
	router.Handle(http.MethodGet, "/:ver/:os/:arch", mirror)

	_, _ = FormatURL(nil)
}

func health(c *router.Context) {
	c.OK(map[string]any{"version": constant.VERSION, "bidtime": constant.BIDTIME})
}
