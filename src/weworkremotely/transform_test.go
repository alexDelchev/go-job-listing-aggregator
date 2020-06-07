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

func TestDeconstructTitle(t *testing.T) {
	input := "value1:value2:value3"

	expectedFirstValue := "value1"
	expectedSecondValue := "value2:value3"

	actualFistValue, actualSecondValue := deconstructTitle(input)

	if expectedFirstValue != actualFistValue {
		t.Errorf("%s failed: Expected first value %s, got %s",
			t.Name(), expectedFirstValue, actualFistValue)
	}

	if expectedSecondValue != actualSecondValue {
		t.Errorf("%s failed: Expected sconed value %s, got %s",
			t.Name(), expectedSecondValue, actualSecondValue)
	}
}
