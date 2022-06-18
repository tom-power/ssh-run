package shared

func DefaultString(v string, defaultV string) string {
	if v == "" {
		return defaultV
	}
	return v
}
