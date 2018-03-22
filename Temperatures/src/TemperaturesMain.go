package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // n: the number of temperatures to analyse
    var n int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&n)
    
    scanner.Scan()
    inputs := strings.Split(scanner.Text()," ")
    var temperatures []int64
    for i := 0; i < n; i++ {
        // t: a temperature expressed as an integer ranging from -273 to 5526
        t,_ := strconv.ParseInt(inputs[i],10,32)
        temperatures = append(temperatures, t)
    }
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println(ExtractTemperature(&temperatures)) // Write answer to stdout
}

func ExtractTemperature(v *[]int64) int {
	if len(*v) > 0 {
		var minDifference float64 = 5526
        for _, temperature := range *v {
			// t: a temperature expressed as an integer ranging from -273 to 5526
			tFloat := float64(temperature)
            if  math.Abs(tFloat) < math.Abs(minDifference) {
				minDifference = tFloat
			} else if math.Abs(tFloat) == math.Abs(minDifference) && int64(minDifference) + temperature == 0 { 
				minDifference = math.Abs(tFloat)
			} 
        }
        return int(minDifference)
	} else {
		return 0
	}
}