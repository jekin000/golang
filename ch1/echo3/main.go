//Echo3 prints its comand-line arguments by join
package main

import (
	"fmt"
	"os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:]," "))
    exe1_cmdname()
    exe2_args_eachline()
}

func exe1_cmdname(){
    fmt.Println("exe1_cmdname:")
    fmt.Println("\t"+os.Args[0])
}

func exe2_args_eachline(){
    fmt.Println("exe2_args_eachline:")
    for _,arg := range os.Args[1:] {
        fmt.Println("\t"+arg)
    }
}

func exe3_add_and_join(){
}
