package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var isHue = flag.Bool("elvish", false, "second case")
	flag.Parse()

	lines, err := LoadFromFile("test_data.txt")
	grid := makeGrid()
	var count int
	if err != nil {
		log.Println(err)
	} else {
		for _, line := range lines {
			cmd := Parse(line, *isHue)
			count += LocalCount(grid, cmd)
		}
		fmt.Println(count)
	}
}

func makeGrid() [][]int {
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	return grid
}
