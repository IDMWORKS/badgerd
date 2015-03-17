package status

type BuildStatus struct {
	DisplayName string `json:displayName`
	Url         string `json:url`
	Color       string `json:color`
}
