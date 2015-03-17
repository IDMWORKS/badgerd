package badge

import (
	"fmt"
	"github.com/IDMWORKS/badgerd/status"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	BuildErrorBadge    = "build-error.svg"
	CoverageErrorBadge = "coverage-error.svg"
)

func ForBuildStatus(status *status.BuildStatus) (string, error) {
	color := status.Color
	switch {
	case color == "blue":
		return "build-passing.svg", nil
	case color == "red":
		return "build-failing.svg", nil
	case strings.Index(color, "_anime") > 0:
		return "build-building.svg", nil
	}

	return BuildErrorBadge, fmt.Errorf("Unknown build color: '%s'", color)
}

func ForRCov(status *status.BuildStatus) (string, error) {
	for i := range status.HealthReport {
		h := status.HealthReport[i]
		matched, err := regexp.MatchString("Rcov coverage", h.Description)
		if err != nil {
			return CoverageErrorBadge, fmt.Errorf("Parse error: '%s'", err)
		}

		if matched {
			digitRe := regexp.MustCompile("Code coverage \\d{1,3}\\.\\d{2}%\\(([^\\)]+)\\)")
			coverage, err := strconv.ParseFloat(digitRe.FindStringSubmatch(h.Description)[1], 32)
			if err != nil {
				return CoverageErrorBadge, fmt.Errorf("Parse error: '%s'", err)
			}

			coverage = math.Ceil(coverage)
			return fmt.Sprintf("coverage-%d.svg", int(coverage)), nil
		}
	}
	return CoverageErrorBadge, nil
}
