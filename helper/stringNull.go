package helper

func Get(dataString []byte) string {
	if len(dataString) == 0 {
		return ""
	}
	return string(dataString)
}
