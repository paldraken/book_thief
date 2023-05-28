package chapters

import "testing"

func TestDecodeText(t *testing.T) {
	secret := "a4484ed145bd5efa5c14a37fefba744b"
	text := "^D\nЖёўЭќFЈѳџѵѳѝЌAЧњЋХќЄЏЦїѻЏЄѶЭЁА`ЈѰЎЅЄќЍЏЏѓЩУEЦЇѲ"
	userId := "628353"
	expected := "<p>Самый простой способ отвязаться от назойливых рас"

	// Test decoding with correct inputs
	decoded := decodeText(secret, text, userId)
	if decoded != expected {
		t.Errorf("Expected %s but got %s", expected, decoded)
	}

	// Test decoding with incorrect secret
	decoded = decodeText("wrong secret", text, userId)
	if decoded == expected {
		t.Errorf("Expected decoding to fail with wrong secret")
	}

	// Test decoding with incorrect userId
	decoded = decodeText(secret, text, "456")
	if decoded == expected {
		t.Errorf("Expected decoding to fail with wrong userId")
	}

	// Test decoding with empty text
	decoded = decodeText(secret, "", userId)
	if decoded != "" {
		t.Errorf("Expected empty string but got %s", decoded)
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"12345", "54321"},
		{"racecar", "racecar"},
	}

	for _, tc := range testCases {
		result := reverseString(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s but got %s for input %s", tc.expected, result, tc.input)
		}
	}
}
