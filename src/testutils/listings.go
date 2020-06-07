// +build test

package testutils

func CompareStringSlices(a []string, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for index, value := range a {
		if value != b[index] {
			return false
		}
	}

	return true
}
