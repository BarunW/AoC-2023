package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ans int

const (
    BasePath string = "./inputs" 
)

func DOne() {
    c := CallDayOne(AbsPath())
    
    result := AddDay(c)
    fmt.Println(result)
}

func D_Two(){
    c := CallDayTwo(AbsPath()) 
    result := AddDay(c)
    fmt.Println(result)
}


func D_Three(){
    c := CallDayThree(AbsPath()) 
    result := AddDay(c)
    fmt.Println(result)
}

func D_Four(){
    c := CallDayFour(AbsPath()) 
    result := AddDay(c)
    fmt.Println(result)
}

 //
 //func D_Five(){
 //    c := CallDayFive(AbsPath()) 
 //    result := AddDay(c)
 //    fmt.Println(result)
 //}

func D_Six(){
    c := CallDaySix(AbsPath())
    result := AddDay(c)
    fmt.Println(result)
}

func AbsPath() string{ 
    fmt.Println(os.Args[1])
    filePath , err := filepath.Abs(os.Args[1])
    
    if err != nil{
       fmt.Println(err)
       os.Exit(1)
    }
    
    return filePath
}

func AddDay( d Do) ans { 
    return ans(d.readFile())
}

func main(){    
    start := time.Now()
    // DOne()
   // D_Two()
   // D_Three()
    
 //    D_Five()
    D_Six()
    fmt.Println(time.Since(start))
}





