package main

import "net/http"

type rewriteServer struct {
	rules []rewriter
}

func (s *rewriteServer) rewrite(path string) *string {
	for _, rule := range s.rules {
		if result := rule.rewrite(path); result != nil {
			return result
		}
	}
	return nil
}

func (s *rewriteServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := s.rewrite(r.URL.Path)
	if target == nil {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, *target, http.StatusPermanentRedirect)
	}
}
