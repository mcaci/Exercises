package dontpanictest

import "testing"
import "dontpanic"

func TestLeadingCloneShouldKeepGoing(t *testing.T) {
	marvin := dontpanic.Clone{0, 1, "RIGHT"}
	expectedAction := "WAIT"
	marvinAction := marvin.EvaluateAction(elevators)
	if marvinAction != expectedAction {
		t.Fatalf("Expected %s but got %s", expectedAction, marvinAction)
	}
}

func TestLeadingCloneShouldBlock(t *testing.T) {
	marvin := dontpanic.Clone{1, 5, "RIGHT"}
	expectedAction := "BLOCK"
	marvinAction := marvin.EvaluateAction(elevators)
	if marvinAction != expectedAction {
		t.Fatalf("Expected %s but got %s", expectedAction, marvinAction)
	}
}


func TestCloneNotArrivedShouldWait(t *testing.T) {
	marvin := dontpanic.Clone{-1, 1, "RIGHT"}
	expectedAction := "WAIT"
	marvinAction := marvin.EvaluateAction(elevators)
	if marvinAction != expectedAction {
		t.Fatalf("Expected %s but got %s", expectedAction, marvinAction)
	}
}

func TestCloneShouldWaitWhenUnderTheElevator(t *testing.T) {
	marvin := dontpanic.Clone{0, 5, "RIGHT"}
	expectedAction := "WAIT"
	marvinAction := marvin.EvaluateAction(elevators)
	if marvinAction != expectedAction {
		t.Fatalf("Expected %s but got %s", expectedAction, marvinAction)
	}
}