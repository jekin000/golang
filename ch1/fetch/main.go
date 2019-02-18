//fetch can dump the url's webpage information
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	//"reflect"
	//fmt.Println(reflect.TypeOf(resp),reflect.ValueOf(resp).Kind())
	//exe 1.8
	"strings"
)

func ioutilReadAll(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

//exe 1.7
func ioCopy(resp *http.Response) {
	_, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io.Copy %v\n", err)
	}
	//if io.Copy read b,err := ,
	//fmt.Printf("%s",b) -> %!s(int64=4154)
	//confirmed from book, the b is copy bytes, book call it nbytes
}

//exe 1.7-1.9
func ioCopy_1_7_1_9(resp *http.Response) {
	_, err := io.Copy(os.Stdout, resp.Body)

	fmt.Fprintf(os.Stderr, "http statuscode = %v\n", resp.Status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io.Copy %v\n", err)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		//exe 1.8
		if !strings.HasPrefix(url, "http://") {
			//url = "http://"+url
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//ioutilReadAll(resp)
		//exe 1.7
		ioCopy_1_7_1_9(resp)
	}
}
