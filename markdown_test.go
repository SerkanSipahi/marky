package marky_test

import (
	"github.com/serkansipahi/marky"
	"io/ioutil"
	"testing"
)

func TestNewMarkdown(t *testing.T) {

	markdownTemplate, _ := ioutil.ReadFile("./markdown_test.md")
	expectedHeaders, _ := ioutil.ReadFile("./markdown_test_expected.txt")

	markdown := marky.NewMarkdown(string(markdownTemplate))
	code := markdown.Compile()

	if code != string(expectedHeaders) {
		t.Error(
			"expected", string(expectedHeaders),
			"got", code,
		)
	}

}

func TestCreateHeaderTag(t *testing.T) {

	// without newLine option and limited h tag size according specification

	expectedTag := "<h1>Hello World</h1>"
	createdTags := marky.CreateHeaderTag("Hello World", 1, false)
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

	expectedTag = "<h6>Hello World</h6>"
	createdTags = marky.CreateHeaderTag("Hello World", 6, false)
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

	expectedTag = "<h6>Hello World</h6>"
	createdTags = marky.CreateHeaderTag("Hello World", 12, false)
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

	// with newLine option
	expectedTag = "<h3>Hello World</h3>\n"
	createdTags = marky.CreateHeaderTag("Hello World", 3, true)
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}
}

func TestCreatePTag(t *testing.T) {

	// without newLine option
	expectedTag := "<p>Hello World</p>"
	createdTags := marky.CreatePTag("Hello World", false)
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

	// with newLine option
	expectedTag = "<p>Hello World</p>\n"
	createdTags = marky.CreatePTag("Hello World", true)
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

}

func TestCreateLinkTag(t *testing.T) {

	expectedTag := "<a href='http://example.com'>Hello World</a>"
	createdTags := marky.CreateLinkTag("Hello World", "http://example.com")
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

}

func TestCreateEmTag(t *testing.T) {

	expectedTag := "<em>Hello World</em>"
	createdTags := marky.CreateEmTag("Hello World")
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

}

func TestCreateStrongTag(t *testing.T) {

	expectedTag := "<strong>Hello World</strong>"
	createdTags := marky.CreateStrongTag("Hello World")
	if createdTags != expectedTag {
		t.Error(
			"expected", expectedTag,
			"got", createdTags,
		)
	}

}
