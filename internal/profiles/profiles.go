// Package profiles loads AWS CLI profile names from the standard config files.
package profiles

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/ini.v1"
)

// List returns sorted unique profile names from AWS config and credentials files.
func List() ([]string, error) {
	profiles := make(map[string]struct{})

	for _, name := range fromAWSCLI() {
		profiles[name] = struct{}{}
	}
	for _, name := range fromFile(configPath(), true) {
		profiles[name] = struct{}{}
	}
	for _, name := range fromFile(credentialsPath(), false) {
		profiles[name] = struct{}{}
	}

	out := make([]string, 0, len(profiles))
	for name := range profiles {
		out = append(out, name)
	}
	sort.Strings(out)
	return out, nil
}

// DebugInfo returns paths used for profile discovery. Useful when List is empty.
func DebugInfo() string {
	return fmt.Sprintf(
		"home=%q aws=%q config=%q credentials=%q",
		homeDir(),
		awsBin(),
		configPath(),
		credentialsPath(),
	)
}

func configPath() string {
	if path := os.Getenv("AWS_CONFIG_FILE"); path != "" {
		return path
	}
	return filepath.Join(awsHome(), "config")
}

func credentialsPath() string {
	if path := os.Getenv("AWS_SHARED_CREDENTIALS_FILE"); path != "" {
		return path
	}
	return filepath.Join(awsHome(), "credentials")
}

func awsHome() string {
	if home := homeDir(); home != "" {
		return filepath.Join(home, ".aws")
	}
	return ".aws"
}

// homeDir resolves the user home directory. Alfred often runs scripts without HOME set,
// and os.UserHomeDir() fails in that case on macOS.
func homeDir() string {
	if dir, err := os.UserHomeDir(); err == nil && dir != "" {
		return dir
	}
	if dir := strings.TrimSpace(os.Getenv("HOME")); dir != "" {
		return dir
	}
	if u, err := user.Current(); err == nil && u.HomeDir != "" {
		return u.HomeDir
	}
	if home := homeFromDSCL(); home != "" {
		return home
	}
	return ""
}

func homeFromDSCL() string {
	username := strings.TrimSpace(os.Getenv("USER"))
	if username == "" {
		if u, err := user.Current(); err == nil {
			username = u.Username
		}
	}
	if username == "" {
		return ""
	}

	out, err := exec.Command("/usr/bin/dscl", ".", "-read", "/Users/"+username, "NFSHomeDirectory").Output()
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "NFSHomeDirectory:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "NFSHomeDirectory:"))
		}
	}
	return ""
}

func fromFile(path string, configFile bool) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	parser, err := ini.Load(data)
	if err != nil {
		return nil
	}

	var profiles []string
	for _, section := range parser.Sections() {
		name := section.Name()
		if name == ini.DefaultSection {
			continue
		}
		if configFile {
			switch {
			case name == "default":
				profiles = append(profiles, "default")
			case strings.HasPrefix(name, "profile "):
				profiles = append(profiles, strings.TrimPrefix(name, "profile "))
			}
			continue
		}
		profiles = append(profiles, name)
	}
	return profiles
}

func fromAWSCLI() []string {
	cmd := exec.Command(awsBin(), "configure", "list-profiles")
	cmd.Env = alfredEnv()

	out, err := cmd.Output()
	if err != nil {
		return nil
	}

	var profiles []string
	for _, line := range strings.Split(string(out), "\n") {
		name := strings.TrimSpace(line)
		if name != "" {
			profiles = append(profiles, name)
		}
	}
	return profiles
}

func awsBin() string {
	for _, candidate := range []string{
		"/opt/homebrew/bin/aws",
		"/usr/local/bin/aws",
	} {
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	if path, err := exec.LookPath("aws"); err == nil {
		return path
	}
	return "aws"
}

func alfredEnv() []string {
	env := append([]string(nil), os.Environ()...)
	env = setEnv(env, "PATH", "/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:"+os.Getenv("PATH"))
	if strings.TrimSpace(getEnv(env, "HOME")) == "" {
		if home := homeDir(); home != "" {
			env = setEnv(env, "HOME", home)
		}
	}
	return env
}

func getEnv(env []string, key string) string {
	prefix := key + "="
	for _, e := range env {
		if strings.HasPrefix(e, prefix) {
			return strings.TrimPrefix(e, prefix)
		}
	}
	return ""
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
