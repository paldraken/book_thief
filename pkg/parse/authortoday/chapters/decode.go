package chapters

// Алгоритм дешифровки текста главы
func decodeText(secret, text, userId string) string {
	magic := reverseString(secret) + "@_@" + userId
	counter := 0
	result := ""

	for _, c := range text {
		mIdx := int(float64(counter % len(magic)))
		newCh := int(c) ^ int(magic[mIdx])
		counter++
		result = result + string(rune(newCh))
	}
	return result
}

func reverseString(str string) string {
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res = res + string(str[i])
	}
	return res
}
