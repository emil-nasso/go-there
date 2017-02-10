package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticRoutes(t *testing.T) {
	assert := assert.New(t)

	server := rewriteServer{
		rewriters: []rewriter{
			StaticRule{From: "/hello", To: "/world"},
			StaticRule{From: "/tjenna", To: "/mannen"},
		},
	}

	assert.Nil(server.rewrite(""))
	assert.Nil(server.rewrite("/some404"))
	assert.Equal(*(server.rewrite("/hello")), "/world")
	assert.Equal(*(server.rewrite("/tjenna")), "/mannen")
}
