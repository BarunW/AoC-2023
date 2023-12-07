package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DaySix struct {
	filePath string
}

func CallDaySix(inputFilePath string) *DaySix{
	return &DaySix{
		filePath: inputFilePath,
	}
}

func (d *DaySix) openFile() *os.File {
	f, err := os.Open(d.filePath)
	if err != nil {
		panic(err)
	}

	return f
}

func (d *DaySix) readFile() int {
	f := d.openFile()
	defer f.Close()
    newScanner := bufio.NewScanner(f)

	ans := d.doPart(newScanner)
	
    return ans
}

func (d *DaySix) doPart(scanner *bufio.Scanner) int {
        
        scanner.Scan()
        time := strings.Fields(strings.TrimLeft(scanner.Text(), "Time: "))
        time2 := strings.Join(time, "")
        scanner.Scan()
        distance := strings.Fields(strings.TrimLeft(scanner.Text(), "Distance: ")) 
        distance2 := strings.Join(distance, "")
        
        fmt.Println(time, distance)
        
        marginErr := 1
        for i, v := range distance{
            var counter int 
            timeAllocated := d.sTOi(time[i])
            for j := 1; j < timeAllocated; j++{
                dist := (timeAllocated - j) * j 
                if dist > d.sTOi(v){
                    counter++
                }
            } 
            marginErr *= counter
        }

        count2:= 0
        fmt.Println(time2, distance2)
        for i:=1; i < d.sTOi(time2); i++{
            dist := (d.sTOi(time2) - i) * i
            if dist > d.sTOi(distance2){
                count2++
            }
        }
        return count2
        //return marginErr
}

func (d *DaySix) sTOi(s string) int{
    num, err := strconv.ParseInt(s, 10, 64)

    if err != nil{
        panic(err)
    }

    return int(num)
}





