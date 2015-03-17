package badge

import (
	"fmt"
	"github.com/IDMWORKS/badgerd/status"
	"strings"
)

var (
	ErrorBadge = "build-error.svg"
)

func ForBuildColor(color string) (string, error) {
	switch {
	case color == "blue":
		return "build-passing.svg", nil
	case color == "red":
		return "build-failing.svg", nil
	case strings.Index(color, "_anime") > 0:
		return "build-building.svg", nil
	}

	return ErrorBadge, fmt.Errorf("Unknown build color: '%s'", color)
}
