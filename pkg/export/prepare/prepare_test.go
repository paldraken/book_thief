package prepare

import (
	"reflect"
	"testing"
)

func TestRemoveAttributes(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "<p class=\"test\">Hello, world!</p>",
			expected: "<p>Hello, world!</p>",
		},
		{
			input:    "<p style=\"color:red;\">Welcome!</p>",
			expected: "<p>Welcome!</p>",
		},
		{
			input:    "<div>Hello, <span style=\"font-weight:bold;\">world</span>!</div>",
			expected: "<div>Hello, <span>world</span>!</div>",
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		got := removeAttributes(tc.input)

		// Compare actual vs expected
		if got != tc.expected {
			t.Errorf("removeAttributes(%q) = %q; expected %q", tc.input, got, tc.expected)
		}
	}
}

func TestReplaceTags(t *testing.T) {
	var tagReplace = map[string]string{
		"div": "p",
		"b":   "strong",
	}

	// Define test cases
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "<div>Hello, world!</div>",
			expected: "<p>Hello, world!</p>",
		},
		{
			input:    "<b>Bold text</b>",
			expected: "<strong>Bold text</strong>",
		},
		{
			input:    "<p>Lorem <div>ipsum dolor</div> sit amet</p>",
			expected: "<p>Lorem <p>ipsum dolor</p> sit amet</p>",
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		got := replaceTags(tc.input, tagReplace)

		// Compare actual vs expected
		if got != tc.expected {
			t.Errorf("replaceTags(%q) = %q; expected %q", tc.input, got, tc.expected)
		}
	}
}

func TestCollectAllTags(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "<html><head><title>Test Page</title></head><body><h1>Test Page</h1><p>Sample text</p></body></html>",
			expected: []string{"<html>", "<head>", "<title>", "</title>", "</head>", "<body>", "<h1>", "</h1>", "<p>", "</p>", "</body>", "</html>"},
		},
		{
			input:    "<p>Sample text</p>",
			expected: []string{"<p>", "</p>"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		got := collectAllTags(tc.input)

		// Compare actual vs expected
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("collectAllTags(%q) = %v; expected %v", tc.input, got, tc.expected)
		}
	}
}

func TestFindUnclosedTags(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"<html>", "<head>", "<title>", "</title>", "</head>", "<body>", "<h1>", "</h1>", "<p>", "</p>", "</body>", "</html>"},
			expected: []string{},
		},
		{
			input:    []string{"<html>", "<title>", "<head>", "</head>", "<body>", "<h1>", "</h1>", "<p>", "</p>", "</body>", "</html>"},
			expected: []string{"</title>"},
		},
		{
			input:    []string{"<html>", "<head>", "<title>", "</title>", "<body>", "<h1>", "</h1>", "<p>", "</p>", "</html>"},
			expected: []string{"</head>", "</body>"},
		},
		{
			input:    []string{"<html>", "<head>", "<title>", "</p>", "</title>", "<body>", "<h1>", "</h1>", "</body>", "</html>"},
			expected: []string{"</head>", "<p>"},
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		got := findUnclosedTags(tc.input)

		// Compare actual vs expected
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("findUnclosedTags(%v) = %v; expected %v", tc.input, got, tc.expected)
		}
	}
}
