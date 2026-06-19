package profiles

import (
	"os"
	"testing"
)

func TestHomeFromDSCL(t *testing.T) {
	home := homeFromDSCL()
	if home == "" {
		t.Fatal("homeFromDSCL() returned empty string")
	}
	if _, err := os.Stat(home); err != nil {
		t.Fatalf("homeFromDSCL() = %q is not accessible: %v", home, err)
	}
}
