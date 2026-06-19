// Package granted opens the AWS console via the Granted assume CLI.
package granted

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/adamkaplun/alfred-aws-granted/internal/regions"
)

// OpenConsole runs `assume -c` with optional service and region flags.
func OpenConsole(profile, service, region string) error {
	if profile == "" {
		return fmt.Errorf("profile is required")
	}
	if region == "" {
		region = regions.Default
	}

	args := []string{"-c", profile}
	if service != "" {
		args = append(args, "-s", service)
	}
	args = append(args, "-r", region)

	cmd := exec.Command(assumeBin(), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(alfredPathEnv(), "GRANTED_ALIAS_CONFIGURED=true")

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("assume: %w", err)
	}
	return nil
}

func assumeBin() string {
	for _, candidate := range []string{
		"/opt/homebrew/bin/assume",
		"/usr/local/bin/assume",
	} {
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	if path, err := exec.LookPath("assume"); err == nil {
		return path
	}
	return "assume"
}

func alfredPathEnv() []string {
	env := append([]string(nil), os.Environ()...)
	path := strings.TrimSpace(os.Getenv("PATH"))
	if !strings.Contains(path, "/opt/homebrew/bin") {
		path = "/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:" + path
	}
	return setEnv(env, "PATH", path)
}

func setEnv(env []string, key, value string) []string {
	prefix := key + "="
	out := make([]string, 0, len(env)+1)
	replaced := false
	for _, e := range env {
		if strings.HasPrefix(e, key+"=") {
			if !replaced {
				out = append(out, prefix+value)
				replaced = true
			}
			continue
		}
		out = append(out, e)
	}
	if !replaced {
		out = append(out, prefix+value)
	}
	return out
}
