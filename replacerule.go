package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

type replaceRule struct {
	pattern string
	regexp  *regexp.Regexp
	target  string
	parts   []string
}

type replaceRuleConfig struct {
	Pattern string
	Target  string
}

func (r *replaceRule) rewrite(path string) *string {
	found := r.regexp.FindAllStringSubmatch(path, -1)
	if found == nil {
		return nil
	}
	newPath := r.target
	for i, replacePattern := range r.parts {
		newPath = strings.Replace(newPath, "{"+replacePattern+"}", found[0][i+1], -1)
	}

	return &newPath
}

func (r *replaceRule) String() string {
	return fmt.Sprintf("[replace] %v -> %v", r.pattern, r.target)
}

func newReplaceRule(pattern string, target string) *replaceRule {
	var parts []string
	if strings.HasPrefix(pattern, "/") {
		parts = strings.Split(pattern[1:], "/")
	} else {
		parts = strings.Split(pattern, "/")
	}

	var regexpString string

	regexpString = `^`
	replaceParts := make([]string, 0)
	for _, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			regexpString += `\/(\w+)`
			replaceParts = append(replaceParts, part[1:len(part)-1])
		} else {
			regexpString += `\/` + part
		}
	}
	regexpString += `$`

	return &replaceRule{
		pattern: pattern,
		regexp:  regexp.MustCompile(regexpString),
		parts:   replaceParts,
		target:  target,
	}
}

func loadReplaceRulesFromConfig(r *rewriteServer) {
	var rules []replaceRuleConfig
	err := viper.UnmarshalKey("replace-rules", &rules)

	if err != nil {
		fmt.Printf("Could not parse replace-rules from config file: %s \n", err)
		return
	}
	for _, rule := range rules {
		r.appendRewriter(newReplaceRule(rule.Pattern, rule.Target))
	}
}
