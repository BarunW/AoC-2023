package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type ans int

const (
    BasePath string = "./inputs" 
)

func DOne() {
    filePath := filepath.Join(BasePath,"day1_input2.txt")
    if len(os.Args) != 0{
        filePath , _= filepath.Abs(os.Args[1])
    }

    fmt.Println(filePath)
    c := CallDayOne(filePath)

    result := AddDay(c)
    fmt.Println(result)
}

func AddDay( d Do) ans { 
    return ans(d.readFile())
}

func main(){    
    DOne()

}





