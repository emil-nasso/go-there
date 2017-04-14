package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// StaticRule - TODO
type StaticRule struct {
	gorm.Model
	From        string
	To          string
	LandingPage bool
}

func (r StaticRule) rewrite(path string) *rewriteResult {
	if path == r.From {
		return &rewriteResult{r.To, r.LandingPage}
	}
	return nil
}

func (r StaticRule) String() string {
	return fmt.Sprintf("[static](%t) %v -> %v", r.LandingPage, r.From, r.To)
}

func loadStaticRulesFromConfig(r *rewriteServer) {
	var rules []StaticRule
	err := viper.UnmarshalKey("static-rules", &rules)

	if err != nil {
		fmt.Printf("Could not parse static-rules from config file: %s \n", err)
		return
	}
	for _, rule := range rules {
		r.appendRewriter(rule)
	}

}
