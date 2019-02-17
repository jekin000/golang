//fetch can dump the url's webpage information
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"io"
	//"reflect"
	//fmt.Println(reflect.TypeOf(resp),reflect.ValueOf(resp).Kind())
)

func ioutilReadAll(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

//exe 1.7
func ioCopy(resp *http.Response){
	_,err := io.Copy(os.Stdout,resp.Body)
	if err != nil{
		fmt.Fprintf(os.Stderr,"io.Copy %v\n",err)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//ioutilReadAll(resp)
		ioCopy(resp)
	}
}
