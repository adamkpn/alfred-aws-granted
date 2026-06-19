// Package regions lists AWS regions supported by Granted console access.
package regions

// Default is the region used when none is explicitly selected.
const Default = "us-east-2"

// All returns AWS regions that Granted accepts via the -r flag.
// The default region is listed first.
func All() []string {
	others := []string{
		"af-south-1",
		"ap-east-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"ap-south-1",
		"ap-south-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"ap-southeast-4",
		"ap-southeast-5",
		"ca-central-1",
		"cn-north-1",
		"cn-northwest-1",
		"eu-central-1",
		"eu-central-2",
		"eu-north-1",
		"eu-south-1",
		"eu-south-2",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-west-4",
		"eu-west-6",
		"il-central-1",
		"me-central-1",
		"me-central-2",
		"me-south-1",
		"sa-east-1",
		"us-east-1",
		"us-gov-east-1",
		"us-gov-west-1",
		"us-west-1",
		"us-west-2",
		"us-west-3",
	}

	out := make([]string, 0, len(others)+1)
	out = append(out, Default)
	for _, region := range others {
		if region == Default {
			continue
		}
		out = append(out, region)
	}
	return out
}
