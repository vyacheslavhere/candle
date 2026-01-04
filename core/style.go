package core

/*
 * Represents style of error reporting
 */
type Style struct {
	ErrorPrefix        string
	ErrorPrefixColor   string
	TitleColor         string
	FileNameColor      string
	FileNameColon      string
	FileNameColonColor string
	LineNumberColor    string
	LineColor          string
	SpaceChar          string
	ArrowChar          string
	ArrowCharColor     string
	CaptionColor       string
}

// Default style
func DefaultStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   RedColor,
		TitleColor:         ResetColor,
		FileNameColor:      CyanColor,
		FileNameColon:      ":",
		FileNameColonColor: ResetColor,
		LineNumberColor:    GrayColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     RedColor,
		CaptionColor:       RedColor,
	}
}
