package test_helpers

func AreAnnotationsEqual(annotations1 []string, annotations2 []string) bool {
	if len(annotations1) != len(annotations2) {
		return false
	}

	for i := range annotations1 {
		if annotations1[i] != annotations2[i] {
			return false
		}
	}

	return true
}
