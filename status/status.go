package status

type BuildStatus struct {
	DisplayName  string   `json:displayName`
	Url          string   `json:url`
	Color        string   `json:color`
	HealthReport []Health `json:healthReport`
}

type Health struct {
	Description string `json:description`
	Score       int    `json:score`
}
