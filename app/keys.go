package app

func isQuitKey(key string) bool {
	return key == "q"
}

func isTodoListFocusKey(key string) bool {
	return key == "1"
}

func isCalendarFocusKey(key string) bool {
	return key == "2"
}
