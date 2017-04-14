package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TESTDBPATH = "/tmp/go-there-test.sqlite"

func init() {
	os.Remove(TESTDBPATH)
}

func TestDatabaseRuleRewriterCreation(t *testing.T) {
	r := newDatabaseRuleRewriter(TESTDBPATH)
	assertStringOutput(t, 0, r)
}

func assertStringOutput(t *testing.T, numRules int, r *databaseRuleRewriter) {
	assert.Equal(t, fmt.Sprintf("[db] %s. %d rules", TESTDBPATH, numRules), r.String())
}

func TestAddRules(t *testing.T) {
	r := newDatabaseRuleRewriter(TESTDBPATH)
	r.Add("/test", "http://www.google.se")
	assertStringOutput(t, 1, r)

	assert.Nil(t, r.rewrite("/jajdjdjd"))
	assert.Equal(t, "http://www.google.se", r.rewrite("/test").url)
}
