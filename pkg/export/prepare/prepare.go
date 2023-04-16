package prepare

import (
	"math"
	"regexp"
	"strings"
)

func SanitizeForFB2(text string) string {

	text = strings.ReplaceAll(text, "<br>", "")
	text = strings.ReplaceAll(text, "</br>", "")
	// text = strings.ReplaceAll(text, "</p>", "</p>\n")

	tags := collectAllTags(text)

	unclosedTags := findUnclosedTags(tags)

	text = closeUncloseTags(text, unclosedTags)

	return text
}

func collectAllTags(s string) []string {
	found := []string{}
	var inTag bool
	var currTag string

	for i := 0; i < len(s); i++ {
		if s[i] == '<' && !inTag {
			inTag = true
			currTag = ""
		}

		if inTag {
			if s[i] == '>' {
				found = append(found, strings.Split(currTag, " ")[0]+">")
				s = s[i+1:]
				inTag = false
				i = -1
			} else {
				currTag += string(s[i])
			}
		}
	}

	return found
}

func findUnclosedTags(tags []string) []string {
	m := make(map[string]int)
	result := []string{}

	for _, tag := range tags {

		// skip self closed tags
		if len(tag) > 3 && tag[len(tag)-2:] == "/>" {
			continue
		}

		if tag[0] == '<' && tag[1] != '/' {
			v := m[tag]
			m[tag] = v + 1
		} else {
			closeFor := tag[0:1] + tag[2:]
			v := m[closeFor]
			m[closeFor] = v - 1
		}
	}

	for tag, v := range m {
		if v == 0 {
			continue
		}
		if v > 0 {
			tag = "</" + tag[1:]
		}

		v = int(math.Abs(float64(v)))
		for i := 0; i < v; i++ {
			result = append(result, tag)
		}
	}
	return result
}

func closeUncloseTags(text string, tags []string) string {
	append := ""
	for _, tag := range tags {
		append = append + tag
	}
	return text + append
}

func replaceTags(text string, tagReplace map[string]string) string {
	for tag1, tag2 := range tagReplace {
		text = strings.ReplaceAll(text, "<"+tag1+">", "<"+tag2+">")
		text = strings.ReplaceAll(text, "</"+tag1+">", "</"+tag2+">")
	}
	return text
}

func removeAttributes(text string) string {
	re := regexp.MustCompile(`(?i)<([a-zA-Z0-9]+\b)([^>]+)>`)
	return re.ReplaceAllString(text, "<$1>")
}
