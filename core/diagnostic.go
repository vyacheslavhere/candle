package core

/*
 * Diagnostic is a part of error
 * with the source code location and label provided
 */
type Diagnostic struct {
	Source   string
	Location Location
	Label    string
}

// Creates new diagnostic
func NewDiagnostic(source string, location Location, label string) Diagnostic {
	return Diagnostic{
		Source:   source,
		Location: location,
		Label:    label,
	}
}
