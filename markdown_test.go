package marky_test

import (
	"github.com/serkansipahi/marky"
	"io/ioutil"
	"testing"
)

func TestNewMarkdown(t *testing.T) {

	markdownTemplate, _ := ioutil.ReadFile("./markdown_test.md")
	expectedHeaders, _ := ioutil.ReadFile("./markdown_test_expected.txt")

	_, markdown := marky.NewMarkdown(string(markdownTemplate))
	code := markdown.Compile()

	if code != string(expectedHeaders) {
		t.Error(
			"expected", string(expectedHeaders),
			"got", code,
		)
	}

}
