package marky

import (
	"bufio"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	hTagRegex        = regexp.MustCompile(`^(#+).*`)
	pTagRegex        = regexp.MustCompile(`^[^#]+`)
	linkTagRegex     = regexp.MustCompile(`(\[(.*?)\]\((.*?)\))`)
	strongEmTagRegex = regexp.MustCompile(`\*\*\*(.*?)\*\*\*`)
	strongTagRegex   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	emTagRegex       = regexp.MustCompile(`\*(.*?)\*`)

	ErrMarkdownTemplateFound = errors.New("No markup found")
)

func NewMarkdown(text string) *Markdown {

	if text == "" {
		return &Markdown{}
	}

	return &Markdown{
		MarkdownTemplate: text,
	}
}

type Markdown struct {
	MarkdownTemplate string
}

func CreateHeaderTag(text string, size int) string {
	return "<h" + strconv.Itoa(size) + ">" + strings.Trim(text, " ") + "</h" + strconv.Itoa(size) + ">\n"
}

func CreatePTag(text string) string {
	return "<p>" + text + "</p>\n"
}

func CreateLinkTag(text, href string) string {
	return "<a href='" + href + "'>" + text + "</a>"
}

func CreateEmTag(text string) string {
	return "<em>" + text + "</em>"
}

func CreateStrongTag(text string) string {
	return "<strong>" + text + "</strong>"
}

func RenderLines(text string) string {

	renderedText := ""

	if matched := hTagRegex.FindStringSubmatch(text); len(matched) == 2 {
		headerSize := len(matched[1])
		renderedText += CreateHeaderTag(text[headerSize:], headerSize)
	} else if matched := pTagRegex.FindStringSubmatch(text); len(matched) == 1 {
		renderedText += CreatePTag(matched[0])
	}

	return renderedText
}

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

func RenderHighlightTags(text string) string {

	for {
		if matched := strongEmTagRegex.FindStringSubmatch(text); len(matched) == 2 {
			emTag := CreateEmTag(matched[1])
			strongTag := CreateStrongTag(emTag)
			text = strings.Replace(text, matched[0], strongTag, -1)
		} else if matched := strongTagRegex.FindStringSubmatch(text); len(matched) == 2 {
			strongTag := CreateStrongTag(matched[1])
			text = strings.Replace(text, matched[0], strongTag, -1)
		} else if matched := emTagRegex.FindStringSubmatch(text); len(matched) == 2 {
			emTag := CreateEmTag(matched[1])
			text = strings.Replace(text, matched[0], emTag, -1)
		} else {
			break
		}
	}

	return text
}

func (md *Markdown) Compile() string {

	if md.MarkdownTemplate == "" {
		return ""
	}

	renderedHtml := ""
	scanner := bufio.NewScanner(strings.NewReader(md.MarkdownTemplate))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		// get text line
		text := scanner.Text()

		// render markdown template to html
		renderedHtml += RenderLines(text)
		renderedHtml = RenderLinks(renderedHtml)
		renderedHtml = RenderHighlightTags(renderedHtml)
	}

	return renderedHtml
}
