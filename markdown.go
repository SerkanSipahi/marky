package marky

import (
	"bufio"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

func NewMarkdown(text string) *Markdown {

	if text == "" {
		return &Markdown{}
	}

	return &Markdown{
		Text: text,
	}
}

var (
	hTagRegex        = regexp.MustCompile(`^(#+).*`)
	pTagRegex        = regexp.MustCompile(`^[^#]+`)
	linkTagRegex     = regexp.MustCompile(`(\[(.*?)\]\((.*?)\))`)
	strongEmTagRegex = regexp.MustCompile(`\*\*\*(.*?)\*\*\*`)
	strongTagRegex   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	emTagRegex       = regexp.MustCompile(`\*(.*?)\*`)

	ErrMarkdownTemplateFound = errors.New("No markup found")
)

type Markdown struct {
	Text string
	Html string
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

func (md *Markdown) Compile() string {

	if md.Text == "" {
		return ""
	}

	// init line by line scanner
	scanner := bufio.NewScanner(strings.NewReader(md.Text))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		text := scanner.Text()

		// render lines (header and paragraph)
		if matched := hTagRegex.FindStringSubmatch(text); len(matched) == 2 {
			headerSize := len(matched[1])
			md.Html += CreateHeaderTag(text[headerSize:], headerSize)
		} else if matched := pTagRegex.FindStringSubmatch(text); len(matched) == 1 {
			md.Html += CreatePTag(matched[0])
		}

		// render links
		for {
			if matched := linkTagRegex.FindStringSubmatch(md.Html); len(matched) == 4 {
				md.Html = strings.Replace(md.Html, matched[1], CreateLinkTag(matched[2], matched[3]), -1)
			} else {
				break
			}
		}

		// render strong or em
		for {
			if matched := strongEmTagRegex.FindStringSubmatch(md.Html); len(matched) == 2 {
				emTag := CreateEmTag(matched[1])
				strongTag := CreateStrongTag(emTag)
				md.Html = strings.Replace(md.Html, matched[0], strongTag, -1)
			} else if matched := strongTagRegex.FindStringSubmatch(md.Html); len(matched) == 2 {
				strongTag := CreateStrongTag(matched[1])
				md.Html = strings.Replace(md.Html, matched[0], strongTag, -1)
			} else if matched := emTagRegex.FindStringSubmatch(md.Html); len(matched) == 2 {
				emTag := CreateEmTag(matched[1])
				md.Html = strings.Replace(md.Html, matched[0], emTag, -1)
			} else {
				break
			}
		}

	}

	return md.Html
}
