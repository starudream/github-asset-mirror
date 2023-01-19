package github

import (
	"os"
	"testing"

	. "github.com/starudream/github-asset-mirror/config"
)

func TestMain(m *testing.M) {
	Default()

	os.Exit(m.Run())
}
