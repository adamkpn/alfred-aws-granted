// Package icons provides Alfred icon paths for workflow results.
package icons

const (
	// Workflow is the Alfred workflow bundle icon.
	Workflow = "icon.png"
	// Profile is the AWS profile icon used for every profile.
	Profile = "icons/profile.png"
	// Region is the AWS region icon.
	Region = "region.png"
)

// ProfileIcon returns the same rounded AWS icon for every profile.
func ProfileIcon(string) string {
	return Profile
}
