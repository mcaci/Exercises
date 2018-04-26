package motorbike

import "testing"

const executed = " but it executed: %s"

func TestOnPlatformSlowdownIfMoving(t *testing.T) {
	testMoto := newMotorbike(11, 5)
	testBridge := bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "SLOW", t)
}

func TestOnPlatformStopIfNotMoving(t *testing.T) {
	testMoto := newMotorbike(11, 0)
	testBridge := bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "WAIT", t)
}

func TestJumpCondition(t *testing.T) {
	testMoto := newMotorbike(5, 5)
	testBridge := bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "JUMP", t)
}

func TestSpeedCondition(t *testing.T) {
	testMoto := newMotorbike(0, 0)
	testBridge := bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "SPEED", t)
}

func TestMaxSpeedCondition(t *testing.T) {
	testMoto := newMotorbike(3, 5)
	testBridge := bridge{6, 4, 8}
	testMotorbikeCommand(testMoto, &testBridge, "WAIT", t)
}

func testMotorbikeCommand(m *motorbike, b* bridge, expectedCommand string, t *testing.T) {
	motorbikeCommand := m.nextCommand(b)
	if expectedCommand != motorbikeCommand {
		t.Fatalf("Motorbike should accelerate at start" + executed, motorbikeCommand)
	}
}