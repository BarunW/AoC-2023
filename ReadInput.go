package main

import "os"

type Do interface {
    openFile() *os.File
    readFile() int   
}







