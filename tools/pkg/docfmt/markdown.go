// Package docfmt normalizes generated provider documentation.
package docfmt

import (
	"regexp"
	"strings"
)

const calloutSignature = "-> **Syntax Rule:**"

const syntaxRulesCallout = "-> **Syntax Rule:** This provider uses OneOf groups for mutually " +
	"exclusive options. Fields documented as \"Optional Block\" use empty block " +
	"syntax `field_name {}`, **never** `field_name = true`. Boolean attributes " +
	"(like `add_hsts`, `http_redirect`) use `= true/false` as normal."

var (
	headingPattern = regexp.MustCompile(`^#{1,6}\s+\S`)
	listPattern    = regexp.MustCompile(`^\s*(?:[-+*]|[0-9]+\.)\s+\S`)
	fencePattern   = regexp.MustCompile(`^\s*(?:` + "```" + `|~~~)`)
	boldPattern    = regexp.MustCompile(`^\*\*([^*]+)\*\*$`)
	blockContext   = regexp.MustCompile(`^An? .+ block(?: \(.+\))? supports the following:$`)
)

// InjectSyntaxRulesCallout adds the generated schema-syntax note exactly once.
func InjectSyntaxRulesCallout(content string) string {
	if strings.Contains(content, calloutSignature) {
		return content
	}

	for _, marker := range []string{"## Argument Reference", "## Schema"} {
		markerStart := strings.Index(content, marker)
		if markerStart < 0 {
			continue
		}
		lineEndOffset := strings.Index(content[markerStart:], "\n")
		if lineEndOffset < 0 {
			lineEndOffset = len(content) - markerStart
		}
		lineEnd := markerStart + lineEndOffset
		prefix := strings.TrimRight(content[:markerStart], "\n")
		heading := content[markerStart:lineEnd]
		remainder := strings.TrimLeft(content[lineEnd:], "\n")

		var result strings.Builder
		if prefix != "" {
			result.WriteString(prefix)
			result.WriteString("\n\n")
		}
		result.WriteString(heading)
		result.WriteString("\n\n")
		result.WriteString(syntaxRulesCallout)
		if remainder != "" {
			result.WriteString("\n\n")
			result.WriteString(remainder)
		}
		return strings.TrimRight(result.String(), "\n") + "\n"
	}

	return content
}

// NormalizeMarkdownSpacing applies the blank-line rules generated docs need.
func NormalizeMarkdownSpacing(content string) string {
	lines := strings.Split(content, "\n")
	output := make([]string, 0, len(lines))
	inFence := false
	previousWasList := false

	appendBlank := func() {
		if len(output) > 0 && output[len(output)-1] != "" {
			output = append(output, "")
		}
	}

	for _, rawLine := range lines {
		line := strings.TrimRight(rawLine, " \t")
		if line == "" {
			appendBlank()
			continue
		}

		if inFence {
			output = append(output, line)
			if fencePattern.MatchString(line) {
				inFence = false
				appendBlank()
			}
			previousWasList = false
			continue
		}

		if fencePattern.MatchString(line) {
			appendBlank()
			output = append(output, line)
			inFence = true
			previousWasList = false
			continue
		}

		if headingPattern.MatchString(line) {
			appendBlank()
			output = append(output, line)
			appendBlank()
			previousWasList = false
			continue
		}

		isList := listPattern.MatchString(line)
		if isList && !previousWasList {
			appendBlank()
		} else if !isList && previousWasList {
			appendBlank()
		}
		output = append(output, line)
		previousWasList = isList
	}

	return strings.TrimRight(strings.Join(output, "\n"), "\n") + "\n"
}

// InsertSectionBeforeOnce inserts a generated section before a heading once.
func InsertSectionBeforeOnce(content, sectionHeading, section, beforeHeading string) string {
	normalizedContent := "\n" + content
	if strings.Contains(normalizedContent, "\n"+sectionHeading+"\n") {
		return content
	}

	markerStart := strings.Index(normalizedContent, "\n"+beforeHeading+"\n")
	if markerStart < 0 {
		return strings.TrimRight(content, "\n") + "\n\n" +
			strings.Trim(section, "\n") + "\n"
	}
	// normalizedContent has one synthetic leading newline, so its marker index is
	// also the heading start index in the original content.
	prefix := strings.TrimRight(content[:markerStart], "\n")
	suffix := strings.TrimLeft(content[markerStart:], "\n")
	return prefix + "\n\n" + strings.Trim(section, "\n") + "\n\n" + suffix
}

// PromoteNestedBlockHeadings turns only block labels into navigable headings.
func PromoteNestedBlockHeadings(content string) string {
	lines := strings.Split(content, "\n")
	for index, line := range lines {
		match := boldPattern.FindStringSubmatch(strings.TrimSpace(line))
		if match == nil {
			continue
		}
		next := index + 1
		for next < len(lines) && strings.TrimSpace(lines[next]) == "" {
			next++
		}
		if next < len(lines) && blockContext.MatchString(strings.TrimSpace(lines[next])) {
			lines[index] = "#### " + match[1]
		}
	}
	return strings.Join(lines, "\n")
}
