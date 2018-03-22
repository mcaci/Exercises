package main

import "fmt"
import "os"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // nbFloors: number of floors
    // width: width of the area
    // nbRounds: maximum number of rounds
    // exitFloor: floor on which the exit is found
    // exitPos: position of the exit on its floor
    // nbTotalClones: number of generated clones
    // nbAdditionalElevators: ignore (always zero)
    // nbElevators: number of elevators
    var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
    fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)
    
    turn := 0
    var elevators Elevators
        fmt.Fprintln(os.Stderr, elevators)
    for i := 0; i < nbElevators; i++ {
        // elevatorFloor: floor on which this elevator is found
        // elevatorPos: position of the elevator on its floor
        var elevatorFloor, elevatorPos int
        fmt.Scan(&elevatorFloor, &elevatorPos)
        elevators = append(elevators, Elevator{elevatorFloor, elevatorPos})
        fmt.Fprintln(os.Stderr, elevatorFloor)
        fmt.Fprintln(os.Stderr, elevatorPos)
        fmt.Fprintln(os.Stderr, elevators)
    }
    elevators = append(elevators, Elevator{exitFloor, exitPos})
    sort.Sort(elevators)

    for {
        // cloneFloor: floor of the leading clone
        // clonePos: position of the leading clone on its floor
        // direction: direction of the leading clone: LEFT or RIGHT
        var cloneFloor, clonePos int
        var direction string
        fmt.Scan(&cloneFloor, &clonePos, &direction)
        marvin := Clone{cloneFloor, clonePos, direction}
        marvinAction := marvin.EvaluateAction(elevators)
        
        // marvinAction := "BLOCK"
        fmt.Println(marvinAction)
        // lastDirection = marvinAction
        fmt.Fprintln(os.Stderr, marvin)
        fmt.Fprintln(os.Stderr, elevators)
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        // if turn == 1 {
        //     fmt.Println(marvinAction)
        // } else {
        //     fmt.Println("WAIT") // action: WAIT or BLOCK
        // }
        turn++
    }
}

type Elevator struct {
    ElevatorFloor, ElevatorPos int
}

type Elevators []Elevator

func (elevator Elevators) Len() int {
    return len(elevator)
}
func (elevator Elevators) Swap(i, j int) {
    elevator[i], elevator[j] = elevator[j], elevator[i]
}
func (elevator Elevators) Less(i, j int) bool {
    return elevator[i].ElevatorFloor < elevator[j].ElevatorFloor
}

type Clone struct {
	CloneFloor, ClonePos int
	Direction string
}

func (marvin *Clone) EvaluateAction(elevators []Elevator) string { 
	if marvin.hasCloneArrived() || marvin.isCloneAtAnElevator(elevators) {
		return "WAIT"
	} else {
		return marvin.action(elevators)
	}
}

func (marvin *Clone) action(elevators []Elevator) string {
	marvinExpectedDir := InWhichDirectionTheCloneShouldGo(marvin.ClonePos, elevators[marvin.CloneFloor].ElevatorPos)
	if marvinExpectedDir == marvin.Direction {
		return "WAIT"	
	} else {
		return "BLOCK"
	}
}

func (marvin *Clone) hasCloneArrived() bool {
	return marvin.CloneFloor < 0
}

func (marvin *Clone) isCloneAtAnElevator(elevators []Elevator) bool {
	return marvin.ClonePos == elevators[marvin.CloneFloor].ElevatorPos
}

func InWhichDirectionTheCloneShouldGo(clonePos, elevatorPos int) string {
	if clonePos < elevatorPos {
		return "RIGHT"
	} else if clonePos > elevatorPos{
		return "LEFT"
	} else {
	    return "STAY"
	}
}