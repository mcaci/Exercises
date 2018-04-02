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
	
	var golfCourse []string 
    for i := 0; i < height; i++ {
        var row string
		fmt.Scan(&row)
		golfCourse = append(golfCourse, row)
	}
	
	golfCourseWithPaths := FindPath(golfCourse)
    for _, row := range golfCourseWithPaths { 
        fmt.Println(row);
    }
}

func FindPath(golfCourseMap []string) []string {
	emptyGolfMap := CopyToEmptyGolfCourseMap(golfCourseMap)
	starts := FindBalls(golfCourseMap)
	ends := FindHoles(golfCourseMap)
	paths := pathsFromBallToHole(starts, ends)
	return BuildGolfCourseWithPaths(paths, emptyGolfMap)
}

type Path struct {
	Start *Ball
	StepSequence *[]string
}

func BuildGolfCourseWithPaths(paths [](*Path), golfMap []string) []string {
	for _, path := range paths {
		for _, direction := range *(path.StepSequence) {
			line := golfMap[path.Start.X()]
			golfMap[path.Start.X()] = line[:path.Start.Y()] + direction + line[path.Start.Y()+1:]
		}
	}
	return golfMap
}

func pathsFromBallToHole(balls, holes []Coordinate) [](*Path) {
	var paths [](*Path)
	for i, ball := range balls {
		ballAtCurrentPosition := ball.(Ball)
		var sequence []string	
		for !AreInSamePosition(ballAtCurrentPosition, holes[i]) {
			direction := getDirection(ballAtCurrentPosition, holes[i])
			sequence = append(sequence, direction)
			ballAtCurrentPosition.Move(direction)
		}
		ballAtStartPosition := ball.(Ball)
		paths = append(paths, &Path{&ballAtStartPosition, &sequence})
	}
	return paths
}

func getDirection(ball, hole Coordinate) string {
	var direction string
	if ball.X() < hole.X() && ball.Y() == hole.Y() {
		direction = "v"
	} else if ball.X() == hole.X() && ball.Y() < hole.Y() {
		direction = ">"
	} else if ball.X() > hole.X() && ball.Y() == hole.Y() {
		direction = "^"
	} else if ball.X() == hole.X() && ball.Y() > hole.Y() {
		direction = "<"
	}	
	return direction
}


type Coordinate interface {
	X() int
	Y() int
}

func AreInSamePosition(a, b Coordinate) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}

func FindElements(golfCourseMap []string, elementPredicate func (byte) bool, elementSupplier func (int, int) Coordinate) []Coordinate {
	var elements []Coordinate
	H := len(golfCourseMap)
	for i := 0; i < H; i++ {
		for j := 0; j < len(golfCourseMap[i]); j++ {
			if elementPredicate(golfCourseMap[i][j]) {
				elements = append(elements, elementSupplier(i, j))
			}
		}
	}
	return elements
}

func FindBalls(golfCourseMap []string) []Coordinate {
	return FindElements(golfCourseMap, isPointABall, ballSupplier)
}

func FindHoles(golfCourseMap []string) []Coordinate {
	return FindElements(golfCourseMap, isPointAHole, holeSupplier)
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

type Ball struct {
	x, y int 
}

func isPointABall(point byte) bool {
	return point <= 57 && point >= 48 // between 0 and 9
}

func ballSupplier(x, y int) Coordinate {
	return Ball{x, y}
}

func (ball *Ball) Move(direction string) {
	switch direction {
		case "v": ball.IncrX()
		case "^": ball.DecrX()
		case ">": ball.IncrY()
		case "<": ball.DecrY()
	}
}

func (ball *Ball) IncrX() {
	ball.x++
}

func (ball *Ball) DecrX() {
	ball.x--
}

func (ball *Ball) IncrY() {
	ball.y++
}

func (ball *Ball) DecrY() {
	ball.y--
}

func (ball Ball) X() int {
	return ball.x
}

func (ball Ball) Y() int {
	return ball.y
}

type Hole struct {
	x, y int 
}

func isPointAHole(point byte) bool {
	return point == 72 // letter "H"
}

func holeSupplier(x, y int) Coordinate {
	return &Hole{x, y}
}

func (hole Hole) X() int {
	return hole.x
}

func (hole Hole) Y() int {
	return hole.y
}