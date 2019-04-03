//use templ and custom function
package main

import (
	"./xkcd"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"
)

//%.16s means max length of string is 16, it will cut exceed part
//{{}} - {{}} this is the loop tag
const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------
Number: {{.Num}}
Title:  {{.Title | printf "%.16s"}}
Age:    {{.Year | yearsAgo}}
{{end}}`

func yearsAgo(year string) int {
	curyear := time.Now().Year()
	intyear, _ := strconv.Atoi(year)
	return curyear - intyear
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"yearsAgo": yearsAgo}).
	Parse(templ))

/*
3 issues:
--------------------------
Number: 571
Title:  Can&#39;t Sleep
Age:    10
--------------------------
Number: 572
Title:  Together
Age:    10
--------------------------
Number: 573
Title:  Parental Trollin
Age:    10

*/
func main() {
	result, err := xkcd.SearchSingle()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Items[0].Title)
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
