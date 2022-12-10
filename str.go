package goutils

func StringNullable(str interface{}) string {
	if str == nil {
		return ""
	}
	return str.(string)
}
