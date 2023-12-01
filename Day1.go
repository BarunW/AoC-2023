package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var byteInt = map[byte]string{
    byte('1') :"1",
    byte('2') :"2",
    byte('3') :"3",
    byte('4') :"4",
    byte('5') :"5",
    byte('6') :"6",
    byte('7') :"7",
    byte('8') :"8",
    byte('9') :"9",
}

var subStrToInt = map[string]string{
    "one" : "1",
    "two" : "2",
    "three" : "3",
    "four" : "4",
    "five" : "5",
    "six" : "6",
    "seven" : "7",
    "eight" : "8",
    "nine" : "9",
}

type DayOne struct {
	filePath string
}

func CallDayOne(inputFilePath string) *DayOne {
	return &DayOne{
		filePath: inputFilePath,
	}
}

func (d *DayOne) openFile() *os.File {
	f, err := os.Open(d.filePath)
	if err != nil {
		panic(err)
	}

	return f
}

func (d *DayOne) readFile() int {
	f := d.openFile()
	defer f.Close()
    newScanner := bufio.NewScanner(f)

	ans := d.doPart2(newScanner)
	
    return ans
}

func (d *DayOne) doPart(scanner *bufio.Scanner) int {
    fmt.Printf("%+v\n", byteInt)

    var totalCalibration int 

    for scanner.Scan() {

        calPerToken := make([]string,2)
        calPerToken[0] = ""
        calPerToken[1] = ""
        token := scanner.Bytes()

        for i:=0; i < len(token); i++{
            v, ok := byteInt[token[i]]
            if ok {
       //         fmt.Println(v,token[i])
                if calPerToken[0] == ""{
                    calPerToken[0] = v
                    continue
                }
                
                calPerToken[1] = v
            }
        }
        
        if calPerToken[1] == "" {
            calPerToken[1] = calPerToken[0]
        }
        num , err := strconv.Atoi(calPerToken[0]+calPerToken[1])
        if err != nil{
            fmt.Println("Unable to convert the String at token", string(token))
            panic(err)
        }
        totalCalibration =  totalCalibration+num
        fmt.Println(totalCalibration, calPerToken)
	}
	return totalCalibration
}


func (d *DayOne) doPart2(scanner *bufio.Scanner) int {
    var totalCalibration int 
    for scanner.Scan() {

        calPerToken := make([]string,2)
        calPerToken[0] = ""
        calPerToken[1] = ""
        token := scanner.Bytes()
        others := ""

        for i:=0; i < len(token); i++{
            v, ok := byteInt[token[i]]
            if ok {
       //         fmt.Println(v,token[i])
                if calPerToken[0] == ""{
                    calPerToken[0] = v
                    continue
                }
                calPerToken[1] = v
            }
        // second part
            others += string(token[i])
            for k, v := range subStrToInt{
                if strings.Contains(others, k){
                    if calPerToken[0] == ""{
                        calPerToken[0] = v
                    } else {
                        calPerToken[1] = v
                    }
                    others = string(others[len(others)-1])
                }

            }
            
            
        }

        if calPerToken[1] == "" {
            calPerToken[1] = calPerToken[0]
        }
        num , err := strconv.Atoi(calPerToken[0]+calPerToken[1])
        if err != nil{
            fmt.Println("Unable to convert the String at token", string(token))
            panic(err)
        }
        totalCalibration =  totalCalibration+num
	}

	return totalCalibration
}







