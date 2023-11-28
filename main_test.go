package main_test

import (
	gf "github.com/lizongying/go-github-faster"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIps(t *testing.T) {
	github := gf.NewGithub(22, "api")
	ips := github.GetIps()
	t.Log(ips)
	assert.Greater(t, len(ips), 0)
}
