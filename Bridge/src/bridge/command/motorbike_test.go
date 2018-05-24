package command

import (
	"bridge/env"
	"testing"
)

const executed = " but it executed: %s"

func TestOnPlatformSlowdownIfMoving(t *testing.T) {
	testMoto := env.NewMotorbike(11, 5)
	testBridge := env.Bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "SLOW", t)
}

func TestOnPlatformStopIfNotMoving(t *testing.T) {
	testMoto := env.NewMotorbike(11, 0)
	testBridge := env.Bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "WAIT", t)
}

func TestJumpCondition(t *testing.T) {
	testMoto := env.NewMotorbike(5, 5)
	testBridge := env.Bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "JUMP", t)
}

func TestSpeedCondition(t *testing.T) {
	testMoto := env.NewMotorbike(0, 0)
	testBridge := env.Bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "SPEED", t)
}

func TestMaxSpeedCondition(t *testing.T) {
	testMoto := env.NewMotorbike(3, 5)
	testBridge := env.Bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "WAIT", t)
}

func testMotorbikeCommand(m *(env.Motorbike), b *(env.Bridge), expectedCommand string, t *testing.T) {
	motorbikeCommand := nextCommand(m, b)
	if expectedCommand != motorbikeCommand {
		t.Fatalf("Motorbike should accelerate at start"+executed, motorbikeCommand)
	}
}
