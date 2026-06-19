// Alfred workflow for opening AWS Console via Granted with fuzzy search.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"

	"github.com/adamkaplun/alfred-aws-granted/internal/granted"
	"github.com/adamkaplun/alfred-aws-granted/internal/icons"
	"github.com/adamkaplun/alfred-aws-granted/internal/profiles"
	"github.com/adamkaplun/alfred-aws-granted/internal/regions"
	"github.com/adamkaplun/alfred-aws-granted/internal/services"
)

var wf *aw.Workflow

func main() {
	mode := commandMode()
	if mode == "open" {
		if err := runOpen(openArgs()); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	wf = aw.New()
	switch mode {
	case "profiles":
		wf.Run(runProfiles)
	case "services":
		wf.Run(runServices)
	case "regions":
		wf.Run(runRegions)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", mode)
		os.Exit(1)
	}
}

func commandMode() string {
	base := strings.ToLower(filepath.Base(os.Args[0]))
	switch {
	case strings.Contains(base, "profiles"):
		return "profiles"
	case strings.Contains(base, "services"):
		return "services"
	case strings.Contains(base, "regions"):
		return "regions"
	case strings.Contains(base, "open"):
		return "open"
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "profiles", "services", "regions", "open":
			return os.Args[1]
		}
	}
	return "profiles"
}

func usesSubcommandArg() bool {
	if len(os.Args) <= 1 {
		return false
	}
	switch os.Args[1] {
	case "profiles", "services", "regions", "open":
		return true
	default:
		return false
	}
}

func queryArg() string {
	if usesSubcommandArg() {
		if len(os.Args) > 2 {
			return os.Args[2]
		}
		return ""
	}
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return ""
}

func openArgs() []string {
	if usesSubcommandArg() {
		return os.Args[2:]
	}
	return os.Args[1:]
}

func runProfiles() {
	profileNames, err := profiles.List()
	if err != nil {
		wf.Fatalf("list profiles: %v", err)
	}

	for _, name := range profileNames {
		wf.NewItem(name).
			Arg(name).
			UID(name).
			Match(name).
			Icon(&aw.Icon{Value: icons.ProfileIcon(name)}).
			Valid(true)
	}

	query := strings.TrimSpace(queryArg())
	if query != "" {
		wf.Filter(query)
	}

	if len(wf.Feedback.Items) == 0 {
		wf.WarnEmpty(
			"No AWS profiles found",
			profiles.DebugInfo(),
		)
	}
	wf.SendFeedback()
}

func runServices() {
	for _, svc := range services.All() {
		item := wf.NewItem(svc.Title).
			Arg(svc.Arg).
			UID(svc.UID).
			Subtitle(svc.Keywords).
			Valid(true)
		if svc.Icon != "" {
			item.Icon(&aw.Icon{Value: svc.Icon})
		}
		item.Match(svc.Keywords)
	}

	query := strings.TrimSpace(queryArg())
	if query != "" {
		wf.Filter(query)
	}
	wf.WarnEmpty("No services matched", "Try a different query")
	wf.SendFeedback()
}

func runRegions() {
	for _, region := range regions.All() {
		item := wf.NewItem(region).
			Arg(region).
			UID(region).
			Icon(&aw.Icon{Value: icons.Region}).
			Valid(true)
		if region == regions.Default {
			item.Subtitle("Default region")
		}
	}

	query := strings.TrimSpace(queryArg())
	if query != "" {
		wf.Filter(query)
	}
	wf.WarnEmpty("No regions matched", "Try a different query")
	wf.SendFeedback()
}

func runOpen(args []string) error {
	profile, service, region := parseOpenArgs(args)
	return granted.OpenConsole(profile, service, region)
}

func parseOpenArgs(args []string) (profile, service, region string) {
	if len(args) == 0 {
		return "", "", ""
	}

	joined := strings.Join(args, " ")
	if strings.Contains(joined, ",") {
		parts := strings.Split(joined, ",")
		if len(parts) > 0 {
			profile = strings.TrimSpace(parts[0])
		}
		if len(parts) > 1 {
			service = strings.TrimSpace(parts[1])
		}
		if len(parts) > 2 {
			region = strings.TrimSpace(parts[2])
		}
		return profile, service, region
	}

	switch len(args) {
	case 1:
		profile = args[0]
	case 2:
		profile, service = args[0], args[1]
	default:
		profile, service, region = args[0], args[1], args[2]
	}
	return profile, service, region
}
