package core

/*
 * Location structure.
 * Represents source code location.
 */
type Location struct {
	Line  int
	Start int
	End   int
}

// Creates new location
func NewLocation(line, start, end int) Location {
	return Location{
		Line:  line,
		Start: start,
		End:   end,
	}
}
