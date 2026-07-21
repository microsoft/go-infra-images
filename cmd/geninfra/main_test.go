package main

import (
	"os/exec"
	"testing"
)

func TestCurrentGeninfraOutputRepeatable(t *testing.T) {
	cmd := exec.Command("go", "run", "./cmd/geninfra", "-exit-code")
	cmd.Dir = "../../"
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("geninfra failed: %v\n%s", err, out)
	}
	if len(out) > 0 {
		t.Fatalf("geninfra output:\n%s", out)
	}
}
