package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var width, height int
    fmt.Scan(&width, &height)
    
    for i := 0; i < height; i++ {
        var row string
        fmt.Scan(&row)
    }
    for i := 0; i < height; i++ {
        fmt.Println(">.");
    }
}