package strain

// Ints is a slice of ints
type Ints []int

// Lists is a slice of slice of ints
type Lists [][]int

// Strings is a slice of strings
type Strings []string

// Discard filters a list, based on failure of the given function
func (collection Ints) Discard(f func(int) bool) (filtered Ints) {
	for _, i := range collection {
		if !f(i) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}

// Keep filters a list, based on success of the given function
func (collection Ints) Keep(f func(int) bool) (filtered Ints) {
	for _, i := range collection {
		if f(i) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}

// Keep filters a list, based on success of the given function
func (collection Strings) Keep(f func(string) bool) (filtered Strings) {
	for _, i := range collection {
		if f(i) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}

// Keep filters a list, based on success of the given function
func (collection Lists) Keep(f func([]int) bool) (filtered Lists) {
	for _, i := range collection {
		if f(i) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}
