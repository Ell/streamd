package streamd

func SliceIncludes[T comparable](data []T, entry T) bool {
	for _, x := range data {
		if x == entry {
			return true
		}
	}

	return false
}
