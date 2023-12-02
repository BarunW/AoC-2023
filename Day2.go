package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DayTwo struct {
	filePath string
}

func CallDayTwo(inputFilePath string) *DayTwo{
	return &DayTwo{
		filePath: inputFilePath,
	}
}

func (d *DayTwo) openFile() *os.File {
	f, err := os.Open(d.filePath)
	if err != nil {
		panic(err)
	}

	return f
}

func (d *DayTwo) readFile() int {
	f := d.openFile()
	defer f.Close()
    newScanner := bufio.NewScanner(f)

	ans := d.doPart(newScanner)
	
    return ans
}

func (d *DayTwo) doPart(scanner *bufio.Scanner) int{
    // 12 red cubes, 13 green cubes, and 14 blue cubes
    gameId := 1
//  totalGameId  := 0
    var power int

    for scanner.Scan(){ 
        balls := make(map[string]int ,3)
        balls["blue"] = 0
        balls["green"] = 0
        balls["red"] = 0
        
        token := strings.TrimSpace(scanner.Text())

        var color string
 //     var isValidToken bool = true

 //     inside:
        for i := len(token)-1; i >= 1; i--{
            if token[i] >= 'a' && token[i] <= 'z'{  
               // fmt.Println(string(token[i]))
                color = string(token[i]) + color 
            }
            
            if token[i] == ' ' && token[i-1] == ',' || token[i-1] == ';' || token[i-1] == ':' {
                if color != ""{
                    // to catch bugs cuz i am soydev
                    if _, ok := balls[color]; !ok {
                        fmt.Println(gameId)
                        panic("KEY NOT FOUND")
                    }
                    
                    s := string(token[i+1])+string(token[i+2])
                    num ,err := strconv.Atoi(strings.TrimSpace(s))
                    if err != nil {
                        panic(err)
                    }

                    balls[color] = d.compare(num, balls[color])
 //                    if stat := d.underThreshold(color, num); !stat{
 //                        fmt.Println(stat,balls)
 //                        isValidToken = stat
 //                        break inside 
 //                    }
                }
            }

            if token[i] == ',' || token[i] == ';' {
                color = ""
                
            }    
            
        }

        fmt.Println(balls, gameId)
 //        if isValidToken{
 //            totalGameId += gameId
 //        }
        power += balls["green"] * balls["red"] * balls["blue"]
        
        gameId++
    }
 //    return totalGameId
    return power
}

func(d *DayTwo) compare(a, b int) int{
    if a > b {
        return a 
    }
    return b
}

func(d *DayTwo) underThreshold(color string, num int) bool {
    switch color {
    case "green":
        if num > 13 {
            return false
        }
    
    case "red":
        if num > 12 {
            return false
        }

    case "blue":
        if num > 14 {
            return false
        }
    }
    return true
}


