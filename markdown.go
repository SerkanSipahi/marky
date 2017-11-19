package marky

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

// markdown regex´s for h, p, a, strong and em tags
var (
	hTagRegex        = regexp.MustCompile(`^(#+).*`)
	pTagRegex        = regexp.MustCompile(`^[^#]+`)
	linkTagRegex     = regexp.MustCompile(`(\[(.*?)\]\((.*?)\))`)
	strongEmTagRegex = regexp.MustCompile(`\*\*\*(.*?)\*\*\*`)
	strongTagRegex   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	emTagRegex       = regexp.MustCompile(`\*(.*?)\*`)
)

// marky.NewMarkdown creates an instance of Markdown module
// which can receive a markdown text that can be compiled with
// Markdown.Compile() to html
func NewMarkdown(text string) *Markdown {

	if text == "" {
		return &Markdown{}
	}
	return &Markdown{
		MarkdownTemplate: text,
	}
}

// Markdown struct which holds the markdown template
type Markdown struct {
	MarkdownTemplate string
}

// CreateHeaderTag creates a <h{n}> tag by given text and size (h1, h2, ..., ...)
// There size is limited to 6 according the w3c specification
// see https://www.w3schools.com/tags/tag_hn.asp
func CreateHeaderTag(text string, size int, newLine bool) string {
	if size > 6 {
		size = 6
	}

	text = "<h" + strconv.Itoa(size) + ">" + strings.Trim(text, " ") + "</h" + strconv.Itoa(size) + ">"
	if newLine {
		text = text + "\n"
	}

	return text
}

// CreatePTag creates an <p> paragraph tag
func CreatePTag(text string, newLine bool) string {

	text = "<p>" + text + "</p>"
	if newLine {
		text = text + "\n"
	}

	return text

}

// CreateLinkTag creates an <a> link by given text and href
func CreateLinkTag(text, href string) string {
	return "<a href='" + href + "'>" + text + "</a>"
}

// CreateEmTag creates an <em> tag by given text
func CreateEmTag(text string) string {
	return "<em>" + text + "</em>"
}

// CreateStrongTag creates a <strong> tag by given text
func CreateStrongTag(text string) string {
	return "<strong>" + text + "</strong>"
}

// RenderLines renders the typical markdown header tags to html.
// # Hello   -> becomes -> <h1>Hello</h1>
// ## Hello  -> becomes -> <h2>Hello</h2>
// ### Hello -> becomes -> <h3>Hello</h3>
// and so on ...
func RenderLines(text string, newLine bool) string {

	renderedText := ""

	if matched := hTagRegex.FindStringSubmatch(text); len(matched) == 2 {
		headerSize := len(matched[1])
		renderedText += CreateHeaderTag(text[headerSize:], headerSize, newLine)
	} else if matched := pTagRegex.FindStringSubmatch(text); len(matched) == 1 {
		renderedText += CreatePTag(matched[0], newLine)
	}

	return renderedText
}

// RenderLinks renders the typical markdown a link to html.
// [Linkname](http://example.com) -> becomes -> <a href="http://example.com">Link</a>
func RenderLinks(text string) string {

	for {
		if matched := linkTagRegex.FindStringSubmatch(text); len(matched) == 4 {
			text = strings.Replace(text, matched[1], CreateLinkTag(matched[2], matched[3]), -1)
		} else {
			break
		}
	}

	return text
}

// RenderHighlightTags render string and em tags.
// string and em (***some text*** —> <strong><em>some text</em></strong> )
// strong (**some text** —> <strong>some text<strong> )
// emphasized (*some text* —> <em>some text</em> )
func RenderHighlightTags(text string) string {

	for {
		// render string and em
		if matched := strongEmTagRegex.FindStringSubmatch(text); len(matched) == 2 {
			emTag := CreateEmTag(matched[1])
			strongTag := CreateStrongTag(emTag)
			text = strings.Replace(text, matched[0], strongTag, -1)
			continue
		}

		// render strong
		if matched := strongTagRegex.FindStringSubmatch(text); len(matched) == 2 {
			strongTag := CreateStrongTag(matched[1])
			text = strings.Replace(text, matched[0], strongTag, -1)
			continue
		}

		// render emphasized
		if matched := emTagRegex.FindStringSubmatch(text); len(matched) == 2 {
			emTag := CreateEmTag(matched[1])
			text = strings.Replace(text, matched[0], emTag, -1)
			continue
		}

		// nothing more to render
		return text
	}

}

// Compile renders and returns markdown template to html
func (md *Markdown) Compile() string {

	if md.MarkdownTemplate == "" {
		return md.MarkdownTemplate
	}

	renderedHtml := ""
	scanner := bufio.NewScanner(strings.NewReader(md.MarkdownTemplate))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		// get text line
		text := scanner.Text()

		// render markdown template to html
		renderedHtml += RenderLines(text, true)
		renderedHtml = RenderLinks(renderedHtml)
		renderedHtml = RenderHighlightTags(renderedHtml)
	}

	return renderedHtml
}
