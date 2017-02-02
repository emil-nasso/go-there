package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStaticRule(t *testing.T) {
	assert := assert.New(t)
	var rule staticRule
	rule = staticRule{}
	assert.Equal(*rule.rewrite(""), "")

	rule = staticRule{from: "/non/matching"}
	assert.Nil(rule.rewrite(""))
}
