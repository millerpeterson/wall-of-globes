package main

import "testing"

func TestVideoFilterVlcArgs(t *testing.T) {
	expected := "--video-filter=cropadd{croptop=10,cropbottom=140,cropleft=10,cropright=160}"

	cropAddArgs := make(map[string]string)
	cropAddArgs["croptop"] = "10"
	cropAddArgs["cropbottom"] = "140"
	cropAddArgs["cropleft"] = "10"
	cropAddArgs["cropright"] = "160"
	actual := VideoFilterVlcArgs("cropadd", cropAddArgs)

	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}
