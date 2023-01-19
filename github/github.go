package github

import (
	"net/http"
	"time"

	"github.com/google/go-github/v49/github"

	"github.com/starudream/go-lib/cache"
)

var (
	cli = github.NewClient(&http.Client{Timeout: 10 * time.Second})

	mc = cache.SIMPLE()
)
