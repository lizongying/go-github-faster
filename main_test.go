package main_test

import (
	gf "github.com/lizongying/go-github-faster"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIps(t *testing.T) {
	res := gf.GetIps()
	t.Log(res)
	assert.Greater(t, len(res), 0)
}
