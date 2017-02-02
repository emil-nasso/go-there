package main

type staticRule struct {
	from string
	to   string
}

func (r staticRule) rewrite(path string) *string {
	if path == r.from {
		return &r.to
	}
	return nil
}
