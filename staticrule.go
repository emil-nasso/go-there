package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type staticRule struct {
	From string
	To   string
}

func (r staticRule) rewrite(path string) *string {
	if path == r.From {
		return &r.To
	}
	return nil
}

func (r staticRule) String() string {
	return fmt.Sprintf("[static] %v -> %v", r.From, r.To)
}

func loadStaticRulesFromConfig(r *rewriteServer) {
	var rules []staticRule
	err := viper.UnmarshalKey("static-rules", &rules)

	if err != nil {
		fmt.Printf("Could not parse static-rules from config file: %s \n", err)
		return
	}
	for _, rule := range rules {
		r.appendRewriter(rule)
	}

}
