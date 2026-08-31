package todolist

func isOpenModalKey(key string) bool {
	return key == "a"
}

func isCursorDownKey(key string) bool {
	return key == "j" || key == "down"
}

func isCursorUpKey(key string) bool {
	return key == "k" || key == "up"
}
