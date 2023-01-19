package route

import (
	"net/http"
	"testing"

	"github.com/starudream/go-lib/router"
)

func TestMirror(t *testing.T) {
	router.T(t,
		router.TCase{
			Method: http.MethodGet,
			Path:   "/latest",
		},
	)
}
