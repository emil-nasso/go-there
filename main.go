package main

import "net/http"

type rewriter interface {
	rewrite(string) *string
}

func main() {
	rewriteServer := rewriteServer{
		rules: []rewriter{
			staticRule{from: "/google", to: "http://www.google.com"},
			newReplaceRule("/g/{query}", "https://www.google.se/#q={query}"),
		},
	}
	http.HandleFunc("/", rewriteServer.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
