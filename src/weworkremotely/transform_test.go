// +build test

package weworkremotely

import (
	"testing"
)

func TestExtractTextFromHTML(t *testing.T) {
	const input string = "<div>Sample <span>text</span></div>"

	expectedOutput := "Sample text"

	actualOutput := extractTextFromHTML(input)

	if expectedOutput != actualOutput {
		t.Errorf("%s failed: Expected output %s, got %s",
			t.Name(), expectedOutput, actualOutput)
	}
}
