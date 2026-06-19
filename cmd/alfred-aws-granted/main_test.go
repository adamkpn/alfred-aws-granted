package main

import "testing"

func TestParseOpenArgs(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		wantProfile     string
		wantService     string
		wantRegion      string
	}{
		{
			name:        "comma separated",
			args:        []string{"prod-admin,ec2,us-east-1"},
			wantProfile: "prod-admin",
			wantService: "ec2",
			wantRegion:  "us-east-1",
		},
		{
			name:        "profile only",
			args:        []string{"prod-admin"},
			wantProfile: "prod-admin",
		},
		{
			name:        "profile and service",
			args:        []string{"prod-admin", "s3"},
			wantProfile: "prod-admin",
			wantService: "s3",
		},
		{
			name:        "three positional args",
			args:        []string{"prod-admin", "lambda", "eu-west-1"},
			wantProfile: "prod-admin",
			wantService: "lambda",
			wantRegion:  "eu-west-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			profile, service, region := parseOpenArgs(tt.args)
			if profile != tt.wantProfile || service != tt.wantService || region != tt.wantRegion {
				t.Fatalf("parseOpenArgs(%v) = (%q, %q, %q), want (%q, %q, %q)",
					tt.args, profile, service, region, tt.wantProfile, tt.wantService, tt.wantRegion)
			}
		})
	}
}
