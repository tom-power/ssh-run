package shared

func DefaultString(v string, defaultV string) string {
	if v == "" {
		return defaultV
	}
	return v
}

func IsNotEmpty(value string) bool {
	return len(value) > 0
}
