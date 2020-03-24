package su

import (
	"os/exec"
	"testing"
)

func TestIsElevationRequiredErr(t *testing.T) {
	cmd := exec.Command("regedit")
	err := cmd.Run()
	if !IsElevationRequired(err) {
		t.Fatal("false negative error")
	}
	cmd = exec.Command("not exist command")
	err = cmd.Run()
	if IsElevationRequired(err) {
		t.Fatal("false positive error")
	}
	if err != nil {
		// println(err.Error())
	}
}
