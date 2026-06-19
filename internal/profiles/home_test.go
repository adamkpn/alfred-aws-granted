package profiles

import (
	"os"
	"runtime"
	"testing"
)

func TestHomeFromDSCL(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("homeFromDSCL uses macOS dscl")
	}

	home := homeFromDSCL()
	if home == "" {
		t.Fatal("homeFromDSCL() returned empty string")
	}
	if _, err := os.Stat(home); err != nil {
		t.Fatalf("homeFromDSCL() = %q is not accessible: %v", home, err)
	}
}
