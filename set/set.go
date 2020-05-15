package set

// Set represents a set of strings
type Set map[string]bool

// Add adds entries into a set
func (s Set) Add(keys ...string) {
	for _, key := range keys {
		s[key] = true
	}
}

// Contains check existence of an entry in a set
func (s Set) Contains(key string) bool {
	return s[key]
}
