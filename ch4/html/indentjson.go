//like movie, just print indent json
package main

import (
	"./xkcd"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	pres, _ := xkcd.SearchSingle()

	datafmt, err := json.MarshalIndent(*pres, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", datafmt)
}
