package core

import (
	"fmt"
	"strings"
)

/*
 * Reports diagnostics, if they given
 */
func ReportDiagnostics(builder *strings.Builder, error *Error, style *Style) {
	if len(error.Diagnostics) > 0 {
		for _, diagnostic := range error.Diagnostics {
			file, line := ReadSource(diagnostic.Source, diagnostic.Location, style)

			fmt.Fprintln(builder, "")
			fmt.Fprintf(builder, "%s%s%s%s\n", style.FileNameColor, file, style.FileNameColonColor, style.FileNameColon)
			fmt.Fprintf(builder, "%s%d %s%s\n", style.LineNumberColor, diagnostic.Location.Line, style.LineColor, line)

			spacesAmount := diagnostic.Location.Start + len(fmt.Sprintf("%d", diagnostic.Location.Line))
			arrowsAmount := diagnostic.Location.End - diagnostic.Location.Start

			spaces := strings.Repeat(style.SpaceChar, spacesAmount)
			arrows := strings.Repeat(style.ArrowChar, arrowsAmount)

			fmt.Fprintf(builder, "%s%s%s\n", spaces, style.ArrowCharColor, arrows)
			fmt.Fprintf(builder, "%s%s%s%s\n", spaces, style.CaptionColor, diagnostic.Label, ResetColor)
		}
		fmt.Fprintf(builder, "\n")
	}
}

/*
 * Reports an error
 */
func Report(error *Error, style *Style) string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "%s%s%s %s\n", style.ErrorPrefixColor, style.ErrorPrefix, style.TitleColor, error.Title)
	ReportDiagnostics(&builder, error, style)

	return builder.String()
}
