// Issues prints a table of GitHub issues matching the search terms.
// Get https://api.gitbub.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder: dial tcp 199.59.242.151:443: i/o timeout
// curl https://api.github.com/search/issues?q=repo%3Agolang%2Fgo

package main

import (
	"./github"
	"fmt"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	for i := 0; i < 1000; i++ {
		fmt.Printf("Connecting... [%d]\n", i)
		if err != nil {
			fmt.Println(err)
			result, err = github.SearchIssues(os.Args[1:])
		} else {
			break
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	/*
		result, err := github.SearchIssues(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}
	*/

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, items := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			items.Number, items.User.Login, items.Title)
	}
}
