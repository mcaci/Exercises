package main

import "testing"
import "temperature"

func TestNoTemperature(t *testing.T) { // Test6
	v := []int64{}
	expectedTemperature := 0
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}

func TestOneTemperature(t *testing.T) {
	v := []int64{1}
	expectedTemperature := 1
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}

func TestTwoTemperatures(t *testing.T) {
	v := []int64{10, 1}
	expectedTemperature := 1
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}

func TestMixedSample(t *testing.T) {
	v := []int64{1, -2, -8, 4, 5, 1}
	expectedTemperature := 1
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}

func TestSampleWitNegativesOnly(t *testing.T) {
	v := []int64{-12, -5, -137, -5}
	expectedTemperature := -5
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}

func TestWhenElementEquallyDistantTo0TakePositiveValue(t *testing.T) {
	v := []int64{42, 5, 12, 21, -5, 24, 5}
	expectedTemperature := 5
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}

func TestAgainWhenElementEquallyDistantTo0TakePositiveValue(t *testing.T) {
	v := []int64{42, -5, 12, 21, 5, 24, -5}
	expectedTemperature := 5
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}


func TestAnotherMixedSet(t *testing.T) {
	v := []int64{-5, -4, -2, 12, -40, 4, 2, 18, 11, 5, -2}
	expectedTemperature := 2
	actualTemperature := temperature.ExtractTemperature(&v)
	if expectedTemperature !=actualTemperature {
		t.Fatalf("Expected %d but got %d", expectedTemperature, actualTemperature)
	}
}
