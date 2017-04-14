package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStaticRule(t *testing.T) {
	assert := assert.New(t)
	var rule StaticRule
	rule = StaticRule{}
	assert.Equal(rule.rewrite("").url, "")

	rule = StaticRule{From: "/non/matching"}
	assert.Nil(rule.rewrite(""))
}
