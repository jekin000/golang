//Dup2 prints the text of each line that appears more than
//once in the stdin or file, preceded by its count.
package main

import(
    "os"
    "fmt"
    "bufio"
)

type NameLn struct{
    Count int
    Filenames map[string]bool
}

func main(){
    counts := make(map[string]*NameLn)
    files := os.Args[1:]
    if len(files) == 0{
        countLines(os.Stdin,counts)
    } else {
        for _,arg := range files{
            f,err := os.Open(arg)
            if err != nil{
                fmt.Fprintf(os.Stderr,"dup2: %v\n",err)
                continue
            }
            countLines(f,counts)
            f.Close()
        }
    }

    for line,nameln := range counts{
        fmt.Printf("%d\t%s%v\n",nameln.Count,line,keys(nameln.Filenames))
    }
}


func in_map(fname string,filenames map[string]bool) bool{
    _,ok := filenames[fname]
    if ok {
        return true
    } else {
        return false
    }
}

func keys(fmap map[string]bool) []string{
    //WARNING: use make will cause keys extend 2 times
    //in the end, if we append a.txt,b.txt to keys, 
    //keys will be ['','','a.txt','b.txt']
    //if we print keys, it is  [  a.txt b.txt]
    //keys := make([]string,len(fmap))

    var keys []string
    for k,_ := range fmap{
        keys = append(keys,k)
    }

    return keys
}

func countLines(f *os.File, counts map[string]*NameLn){
    input := bufio.NewScanner(f)
    for input.Scan(){
        key := input.Text()
        _,ok := counts[key]
        if ok {
            counts[key].Count++
        } else {
            counts[key] = new(NameLn)
            counts[key].Count = 1
            counts[key].Filenames = make(map[string]bool)
        }

        if in_map(f.Name(),counts[key].Filenames){
            continue
        } else {
            counts[key].Filenames[f.Name()] = true
        }
    }
}
