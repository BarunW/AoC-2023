package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DayThree struct {
    filePath string
}

func CallDayThree(inputFilePath string) *DayThree{
    return &DayThree{
        filePath: inputFilePath,
    }
}

func (d *DayThree) openFile() *os.File {
    f, err := os.Open(d.filePath)
    if err != nil {
        panic(err)
    }

    return f
}

func (d *DayThree) readFile() int {
    f := d.openFile()
    defer f.Close()
    newScanner := bufio.NewScanner(f)

    ans := d.doPart(newScanner)

    return ans
}

func (d *DayThree ) doPart(scanner *bufio.Scanner) int {
    var prevToken string 
    var nextToken string
    var sumOfPartNum int
    var lines []string

    for scanner.Scan(){
        lines = append(lines, scanner.Text())
    }

    for i, v := range lines{
        token := v
        if i < len(lines) - 1{
           nextToken = lines[i+1] 
        }
        var num string 
        for i, v := range token{
            if d.isDigit(v){
                num += string(v) 
                if i == len(token) -1{ 
                    l := i - len(num)
                    if l > -1 && d.isSymbol(rune(token[l])){
                        fmt.Println(num)
                        sumOfPartNum += d.sTOi(num)

                    } else if result := d.handleCondition(i, num, prevToken, nextToken); result{
                        sumOfPartNum += d.sTOi(num)
                    }
                    num = ""
                }

                if i < len(token) -1 && token[i+1] == '.' {
                    l := i - len(num)
                    if l > -1 && d.isSymbol(rune(token[l])){
 //                        fmt.Println(num)
                        sumOfPartNum += d.sTOi(num)
                    }

                    if result := d.handleCondition(i, num, prevToken, nextToken); result{    
 //                        fmt.Println(num)
                        sumOfPartNum += d.sTOi(num)
                    }
                    num = ""
                }

                if i < len(token) -1 && d.isSymbol(rune(token[i+1])){
 //                    fmt.Println(num)
                    sumOfPartNum += d.sTOi(num)
                    num = ""
                }

            }
        }

        prevToken = token
    }

    return sumOfPartNum
}

func (d *DayThree) handleCondition(index int, s , prev, next string) bool {
    switch index{
    case len(next) - 1:
        for i := index-len(s); i <= index; i++{
            if prev != "" && d.isSymbol(rune(prev[i])) {
                return true
            }

            if next != "" && d.isSymbol(rune(next[i])){
                return true
            }
            
        }
        return false

    default:
        for i := index - len(s); i <= index + 1; i++{
            if i < 0 {
                i = 0
            }
            if prev != "" && d.isSymbol(rune(prev[i])){
                return true  
            } 
            if next != "" && d.isSymbol(rune(next[i])){
                return true   
            }
        }
        return false
    }

}

func (d *DayThree) sTOi(s string) int{
    num, err := strconv.Atoi(s)
    if err != nil {
        panic( err )
    }

    return num
}

func (d *DayThree) isDigit(r rune) bool {
    if r >= '0' && r <= '9'{
        return true
    }
    return false
}

func (d *DayThree) isSymbol(r rune) bool {
    if !d.isDigit(r) && r != '.' {
        return true 
    }
    return false
}

