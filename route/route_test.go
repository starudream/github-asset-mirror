package route

import (
	"os"
	"testing"

	. "github.com/starudream/github-asset-mirror/config"
)

func TestMain(m *testing.M) {
	Default()
	Register()

	os.Exit(m.Run())
}
