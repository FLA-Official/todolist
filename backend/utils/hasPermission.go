package utils

func hasPermission(role string, required ...string) bool {
	for _, r := range required {
		if role == r {
			return true
		}
	}
	return false
}
