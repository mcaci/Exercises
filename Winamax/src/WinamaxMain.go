package main

import "fmt"
import "strconv"
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

type Coordinate struct {
	X, Y int 
}

func FindStart(golfCourseMap []string) (*Coordinate, bool) {
	return findPoint(golfCourseMap, isPointANumber)
}

func FindEnd(golfCourseMap []string) (*Coordinate, bool) {
	return findPoint(golfCourseMap, isPointAHole)
}

func findPoint(golfCourseMap []string, pointPredicate func(string) bool) (*Coordinate, bool) {
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if pointPredicate(string(golfCourseMap[i][j])) {
				return &Coordinate{i, j}, true
			}
		}
	}
	return &Coordinate{-1, -1}, false
}

func isPointAHole(point string) bool {
	return point == "H"
}

func isPointANumber(point string) bool {
	_, err := strconv.Atoi(point)
	return err == nil
}

func PathFromBallToHole(ballPosition, holePosition *Coordinate) []string {
	var direction string
	var sequence []string
	pathCoordinate := *ballPosition
	for pathCoordinate != *holePosition {
		direction = getDirection(&pathCoordinate, holePosition)
		sequence = append(sequence, direction)
	}
	return sequence
}

func getDirection(ballPosition, holePosition *Coordinate) string {
	direction := holePositionComparedToBall(ballPosition, holePosition)
	switch direction {
		case "v": (ballPosition.X)++
		case "^": (ballPosition.X)--
		case ">": (ballPosition.Y)++
		case "<": (ballPosition.Y)--
	}
	return direction
}

func holePositionComparedToBall(ballPosition, holePosition *Coordinate) string {
	var direction string
	if ballPosition.X == holePosition.X && ballPosition.Y > holePosition.Y {
		direction = "<"
	} else if ballPosition.X == holePosition.X && ballPosition.Y < holePosition.Y {
		direction = ">"
	} else if ballPosition.X > holePosition.X && ballPosition.Y == holePosition.Y {
		direction = "^"
	} else if ballPosition.X < holePosition.X && ballPosition.Y == holePosition.Y {
		direction = "v"
	}
	return direction
}

func CopyToEmptyGolfCourseMap(golfCourseMap []string) []string {
	var emptyGolfMap []string
	for i := 0; i < len(golfCourseMap); i++ {
		emptyGolfMap = append(emptyGolfMap, createMapLine(golfCourseMap[i]))
	}
	return emptyGolfMap
}

func createMapLine(golfCourseMapLine string) string {
	var golfLine string
	mapElement := "."
	for i := 0; i < len(golfCourseMapLine); i++ {
		golfLine += mapElement
	}
	return golfLine
}