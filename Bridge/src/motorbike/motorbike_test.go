package motorbike

import "testing"

func TestOnPlatformSlowdownIfMoving(t *testing.T) {
	testMoto := motorbike{11, 5}
	testBridge := bridge{6, 4, 8}

	motorbikeCommand := testMoto.nextCommand(&testBridge)
	if "SLOW" != motorbikeCommand {
		t.Fatalf("Motorbike should slow down on the platform but it executed: %s", motorbikeCommand)
	}
}

func TestOnPlatformStopIfNotMoving(t *testing.T) {
	testMoto := motorbike{11, 0}
	testBridge := bridge{6, 4, 8}

	motorbikeCommand := testMoto.nextCommand(&testBridge)
	if "WAIT" != motorbikeCommand {
		t.Fatalf("Motorbike should stop when not moving on the platform but it executed: %s", motorbikeCommand)
	}
}


func TestJumpCondition(t *testing.T) {
	testMoto := motorbike{5, 5}
	testBridge := bridge{6, 4, 8}

	motorbikeCommand := testMoto.nextCommand(&testBridge)
	if "JUMP" != motorbikeCommand {
		t.Fatalf("Motorbike should stop when not moving on the platform but it executed: %s", motorbikeCommand)
	}
}