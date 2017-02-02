package main

type databaseRule struct {
	dbHandler string
	table     string
}

func (r databaseRule) rewrite(path string) *string {
	return nil
}
