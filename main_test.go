package main

import "testing"

func TestState(t *testing.T) {
	var enString string = "Valhall"
	var feilString string = "Noe annet"

	if result := GjørIngenTing(enString); result != feilString {
		t.Errorf("GjørIngenTing(enString) = %q, want %q", result, venstre)
	}
}
