package main

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

type rewriteServer struct {
	rewriters []rewriter
}

func (s *rewriteServer) rewrite(path string) *rewriteResult {
	for _, rule := range s.rewriters {
		if result := rule.rewrite(path); result != nil {
			return result
		}
	}
	return nil
}

func (s *rewriteServer) Rewriters() *[]rewriter {
	return &(s.rewriters)
}

func (s *rewriteServer) appendRewriter(r rewriter) {
	s.rewriters = append(s.rewriters, r)
}

func (s *rewriteServer) handleRedirect(c *gin.Context) {
	url := c.Request.URL.Path
	target := s.rewrite(url)
	if target != nil {
		if target.showLandingPage {
			c.HTML(http.StatusNotFound, "landingpage.html", gin.H{"url": target.url})
		} else {
			debug("Redirecting '%s' to '%s'", url, target.url)
			c.Redirect(http.StatusFound, target.url)
		}
	}
}

func (s *rewriteServer) listRewriters() {
	for _, rewriter := range s.rewriters {
		fmt.Printf("%s\n", rewriter)
	}
}
