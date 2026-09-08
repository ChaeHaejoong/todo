package appointment

func isDeleteKey(key string) bool {
	return key == "d"
}

func isCursorDownKey(key string) bool {
	return key == "j" || key == "down"
}

func isCursorUpKey(key string) bool {
	return key == "k" || key == "up"
}
