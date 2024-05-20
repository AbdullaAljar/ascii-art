package main

import(
	"testing"
	"os/exec"
)

func testMain(t *testing.T){
	cmd := exec.Command("./main 'hello'")
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("expected ____" + string(output))
	} else {
		t.Errorf("expected ____" + string(output))
	}
}