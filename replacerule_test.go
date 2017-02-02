package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglePart(t *testing.T) {
	assert := assert.New(t)
	var rule *replaceRule
	rule = newReplaceRule("/hello/{name}", "/zup/{name}")
	assert.Equal("/zup/world", *(rule.rewrite("/hello/world")))
	assert.Equal("/zup/world250", *(rule.rewrite("/hello/world250")))
	assert.Nil(rule.rewrite("/test"))
	assert.Nil(rule.rewrite("/hello/you/safs"))
}

func TestMultiPart(t *testing.T) {
	assert := assert.New(t)
	var rule *replaceRule

	rule = newReplaceRule("/hello/{name}/{something}", "/zup/{name}/{something}")
	assert.Equal("/zup/world/yo", *(rule.rewrite("/hello/world/yo")))
	assert.Nil(rule.rewrite("/test"))
}

func TestMultiPartFlipped(t *testing.T) {
	assert := assert.New(t)
	var rule *replaceRule

	rule = newReplaceRule("/hello/{name}/{something}", "/zup/{something}/{name}")
	assert.Equal("/zup/yo/world", *(rule.rewrite("/hello/world/yo")))
	assert.Nil(rule.rewrite("/test"))
}
