package icons

import "testing"

func TestProfileIconIsConsistent(t *testing.T) {
	t.Parallel()

	profiles := []string{
		"default",
		"cn",
		"plainid-tech-prod.622209932229.AdministratorAccess",
		"plainid-product.705653228102.PowerUserAccessWithRoles",
	}

	for _, profile := range profiles {
		if got := ProfileIcon(profile); got != Profile {
			t.Fatalf("ProfileIcon(%q) = %q, want %q", profile, got, Profile)
		}
	}
}
