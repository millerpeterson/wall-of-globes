package main

import "testing"

func TestVideoFilterVlcArg(t *testing.T) {
	expected := "--video-filter=somefilter{four=def,one=1,three=abc,two=110}"

	cropAddArgs := make(map[string]string)
	cropAddArgs["one"] = "1"
	cropAddArgs["two"] = "110"
	cropAddArgs["three"] = "abc"
	cropAddArgs["four"] = "def"
	actual := VideoFilterArg("somefilter", cropAddArgs)

	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

func TestCropFilterArg(t *testing.T) {
	expected := "--video-filter=croppadd{cropbottom=140,cropleft=1800,cropright=36000,croptop=10}"

	cropArgs := map[string]int{
		"top":    10,
		"bottom": 140,
		"left":   1800,
		"right":  36000,
	}
	actual := CropFilterArg(cropArgs)

	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}
