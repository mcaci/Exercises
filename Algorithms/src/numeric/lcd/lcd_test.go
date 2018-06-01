package lcd

import "testing"

func TestLcdOf1And1(t *testing.T) {
	lcd := LCD(1, 1)
	expectedLcd := 1
	if expectedLcd != lcd {
		t.Fatalf("Expected result to be %d but %d was computed", expectedLcd, lcd)
	}
}

func TestLcdOf6And10(t *testing.T) {
	lcd := LCD(6, 10)
	expectedLcd := 30
	if expectedLcd != lcd {
		t.Fatalf("Expected result to be %d but %d was computed", expectedLcd, lcd)
	}
}

func TestLcdOf10And6(t *testing.T) {
	lcd := LCD(10, 6)
	expectedLcd := 30
	if expectedLcd != lcd {
		t.Fatalf("Expected result to be %d but %d was computed", expectedLcd, lcd)
	}
}
