package profiles

import (
	"os"
	"path/filepath"
	"testing"
)

func TestList(t *testing.T) {
	dir := t.TempDir()
	config := `[default]
region = us-east-1

[profile dev]
region = eu-west-1

[profile prod]
role_arn = arn:aws:iam::123456789012:role/Admin
`
	credentials := `[dev]
aws_access_key_id = AKIATEST
aws_secret_access_key = secret

[staging]
aws_access_key_id = AKIATEST2
aws_secret_access_key = secret2
`
	if err := os.WriteFile(filepath.Join(dir, "config"), []byte(config), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "credentials"), []byte(credentials), 0o600); err != nil {
		t.Fatal(err)
	}

	t.Setenv("AWS_CONFIG_FILE", filepath.Join(dir, "config"))
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", filepath.Join(dir, "credentials"))
	t.Setenv("HOME", dir)

	got, err := List()
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"default", "dev", "prod", "staging"}
	if len(got) != len(want) {
		t.Fatalf("List() = %v, want %v", got, want)
	}
	for i, name := range want {
		if got[i] != name {
			t.Fatalf("List()[%d] = %q, want %q (full: %v)", i, got[i], name, got)
		}
	}
}
