package module01

func RevString(str string) string {
	var revStr string

	for _, r := range str {
		revStr = string(r) + revStr
	}
	return revStr

}
