package main

import (
	"regexp"
	"strings"
)

type replaceRule struct {
	regexp *regexp.Regexp
	target string
	parts  []string
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

func newReplaceRule(pattern string, target string) *replaceRule {
	if strings.HasPrefix(pattern, "/") {
		pattern = pattern[1:]
	}

	parts := strings.Split(pattern, "/")

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
		regexp: regexp.MustCompile(regexpString),
		parts:  replaceParts,
		target: target,
	}
}
