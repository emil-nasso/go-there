package main

import (
	"fmt"
	"net/http"
)

type rewriteServer struct {
	rewriters []rewriter
}

func (s *rewriteServer) rewrite(path string) *string {
	for _, rule := range s.rewriters {
		if result := rule.rewrite(path); result != nil {
			return result
		}
	}
	return nil
}

func (s *rewriteServer) appendRewriter(r rewriter) {
	s.rewriters = append(s.rewriters, r)
}

func (s *rewriteServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := s.rewrite(r.URL.Path)
	if target == nil {
		http.NotFound(w, r)
	} else {
		debug("Redirecting '%s' to '%s'", r.URL.Path, *target)
		http.Redirect(w, r, *target, http.StatusPermanentRedirect)
	}
}

func (s *rewriteServer) listRewriters() {
	for _, rewriter := range s.rewriters {
		fmt.Printf("%s\n", rewriter)
	}
}
