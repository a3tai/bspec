package chat

import (
	"regexp"
	"strings"

	"github.com/alperdrsnn/clime"
)

// MarkdownFormatter handles console formatting of markdown text
type MarkdownFormatter struct {
	// Configuration for output formatting
	enableColors bool
}

// NewMarkdownFormatter creates a new markdown formatter
func NewMarkdownFormatter(enableColors bool) *MarkdownFormatter {
	return &MarkdownFormatter{
		enableColors: enableColors,
	}
}

// Format converts markdown text to formatted console output
func (f *MarkdownFormatter) Format(markdown string) string {
	if !f.enableColors {
		return f.stripMarkdown(markdown)
	}

	text := markdown

	// Headers
	text = f.formatHeaders(text)

	// Code blocks
	text = f.formatCodeBlocks(text)

	// Inline code
	text = f.formatInlineCode(text)

	// Bold text
	text = f.formatBold(text)

	// Italic text
	text = f.formatItalic(text)

	// Lists
	text = f.formatLists(text)

	// Links
	text = f.formatLinks(text)

	return text
}

// formatHeaders formats markdown headers with colors
func (f *MarkdownFormatter) formatHeaders(text string) string {
	// H1 headers (# Header)
	h1Regex := regexp.MustCompile(`^# (.+)$`)
	text = h1Regex.ReplaceAllStringFunc(text, func(match string) string {
		content := h1Regex.FindStringSubmatch(match)[1]
		return clime.Success.Sprint("■ " + content) + "\n" + strings.Repeat("=", len(content)+2)
	})

	// H2 headers (## Header)
	h2Regex := regexp.MustCompile(`^## (.+)$`)
	text = h2Regex.ReplaceAllStringFunc(text, func(match string) string {
		content := h2Regex.FindStringSubmatch(match)[1]
		return clime.Info.Sprint("▶ " + content) + "\n" + strings.Repeat("-", len(content)+2)
	})

	// H3 headers (### Header)
	h3Regex := regexp.MustCompile(`^### (.+)$`)
	text = h3Regex.ReplaceAllStringFunc(text, func(match string) string {
		content := h3Regex.FindStringSubmatch(match)[1]
		return clime.Warning.Sprint("• " + content)
	})

	return text
}

// formatCodeBlocks formats markdown code blocks
func (f *MarkdownFormatter) formatCodeBlocks(text string) string {
	// Multi-line code blocks with language
	codeBlockRegex := regexp.MustCompile("(?s)```(\\w+)?\\n?(.*?)```")
	text = codeBlockRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := codeBlockRegex.FindStringSubmatch(match)
		language := parts[1]
		code := strings.TrimSpace(parts[2])

		var header string
		if language != "" {
			header = clime.MagentaColor.Sprint("┌── " + language + " ──")
		} else {
			header = clime.MagentaColor.Sprint("┌── code ──")
		}

		// Add borders to each line
		lines := strings.Split(code, "\n")
		formattedLines := make([]string, len(lines))
		for i, line := range lines {
			formattedLines[i] = clime.MagentaColor.Sprint("│ ") + clime.GreenColor.Sprint(line)
		}

		footer := clime.MagentaColor.Sprint("└" + strings.Repeat("─", 10))

		return header + "\n" + strings.Join(formattedLines, "\n") + "\n" + footer
	})

	return text
}

// formatInlineCode formats inline code with backticks
func (f *MarkdownFormatter) formatInlineCode(text string) string {
	inlineCodeRegex := regexp.MustCompile("`([^`]+)`")
	return inlineCodeRegex.ReplaceAllStringFunc(text, func(match string) string {
		code := inlineCodeRegex.FindStringSubmatch(match)[1]
		return clime.GreenColor.Sprint(code)
	})
}

// formatBold formats bold text
func (f *MarkdownFormatter) formatBold(text string) string {
	boldRegex := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	return boldRegex.ReplaceAllStringFunc(text, func(match string) string {
		content := boldRegex.FindStringSubmatch(match)[1]
		return clime.Success.Sprint(content)
	})
}

// formatItalic formats italic text
func (f *MarkdownFormatter) formatItalic(text string) string {
	italicRegex := regexp.MustCompile(`\*([^*]+)\*`)
	return italicRegex.ReplaceAllStringFunc(text, func(match string) string {
		content := italicRegex.FindStringSubmatch(match)[1]
		return clime.BlueColor.Sprint(content)
	})
}

// formatLists formats markdown lists
func (f *MarkdownFormatter) formatLists(text string) string {
	lines := strings.Split(text, "\n")
	formattedLines := make([]string, len(lines))

	for i, line := range lines {
		// Bullet lists
		if strings.HasPrefix(strings.TrimSpace(line), "- ") {
			content := strings.TrimSpace(line)[2:]
			indent := len(line) - len(strings.TrimSpace(line))
			formattedLines[i] = strings.Repeat(" ", indent) + clime.BlueColor.Sprint("• ") + content
		} else if strings.HasPrefix(strings.TrimSpace(line), "* ") {
			content := strings.TrimSpace(line)[2:]
			indent := len(line) - len(strings.TrimSpace(line))
			formattedLines[i] = strings.Repeat(" ", indent) + clime.BlueColor.Sprint("• ") + content
		} else {
			// Numbered lists
			numberedRegex := regexp.MustCompile(`^(\s*)(\d+)\.\s(.+)$`)
			if match := numberedRegex.FindStringSubmatch(line); len(match) > 0 {
				indent := match[1]
				number := match[2]
				content := match[3]
				formattedLines[i] = indent + clime.Warning.Sprint(number+".") + " " + content
			} else {
				formattedLines[i] = line
			}
		}
	}

	return strings.Join(formattedLines, "\n")
}

// formatLinks formats markdown links
func (f *MarkdownFormatter) formatLinks(text string) string {
	linkRegex := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	return linkRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := linkRegex.FindStringSubmatch(match)
		linkText := parts[1]
		url := parts[2]
		return clime.BlueColor.Sprint(linkText) + clime.MagentaColor.Sprint(" ("+url+")")
	})
}

// stripMarkdown removes markdown formatting for plain text output
func (f *MarkdownFormatter) stripMarkdown(text string) string {
	// Remove headers
	headerRegex := regexp.MustCompile(`^#{1,6}\s+(.+)$`)
	text = headerRegex.ReplaceAllString(text, "$1")

	// Remove code blocks
	codeBlockRegex := regexp.MustCompile("(?s)```\\w*\\n?(.*?)```")
	text = codeBlockRegex.ReplaceAllString(text, "$1")

	// Remove inline code
	inlineCodeRegex := regexp.MustCompile("`([^`]+)`")
	text = inlineCodeRegex.ReplaceAllString(text, "$1")

	// Remove bold and italic
	boldRegex := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	text = boldRegex.ReplaceAllString(text, "$1")

	italicRegex := regexp.MustCompile(`\*([^*]+)\*`)
	text = italicRegex.ReplaceAllString(text, "$1")

	// Remove links
	linkRegex := regexp.MustCompile(`\[([^\]]+)\]\([^)]+\)`)
	text = linkRegex.ReplaceAllString(text, "$1")

	return text
}