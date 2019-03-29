// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"./github"
	"fmt"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, items := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			items.Number, items.User.Login, items.Title)
	}
}
