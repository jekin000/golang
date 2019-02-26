//Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import(
    "fmt"
    "os"
    "strconv"
    //$GOPATH/src/myexe/ch2/tempconv/tempconv.go
    //$GOPATH/src/myexe/ch2/tempconv/conv.go
    //"myexe/ch2/tempconv"
    "./tempconv"
)


//#go run cf.go 37
//37ᵒF = 2.7777777777777777ᵒC, 37ᵒC = 98.6ᵒF

func main(){
    for _,arg := range os.Args[1:]{
        t,err := strconv.ParseFloat(arg,64)
        if err != nil{
            fmt.Fprintf(os.Stderr,"cf:%v\n",err)
            os.Exit(1)
        }

        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)
        fmt.Printf("%s = %s, %s = %s\n",
            f,tempconv.FToC(f),c,tempconv.CToF(c))
    }
}
