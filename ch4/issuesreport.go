package main

import (
	"time"
	"log"
	"os"
	"text/template"
	"gopl.io/ch4/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}