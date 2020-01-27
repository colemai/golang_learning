package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    // "os"
    "strconv"
    // "strings"

)


func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main () {
    
    dat, err := ioutil.ReadFile("data/hrank_sock.txt")
    check(err)

    fmt.Print(string(nTemp))

}