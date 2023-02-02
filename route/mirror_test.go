package route

import (
	"net/http"
	"testing"

	"github.com/starudream/go-lib/router"
)

func TestMirror(t *testing.T) {
	router.TE(t,
		router.TC{
			Method:  http.MethodGet,
			Pattern: "/latest",
		},
		router.TC{
			Method:  http.MethodGet,
			Pattern: "/latest/windows/amd64",
		},
	)
}
