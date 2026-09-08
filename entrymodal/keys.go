package entrymodal

func isCloseModalKey(key string) bool {
	return key == "esc"
}

func isAppendTodoKey(key string) bool {
	return key == "enter"
}

func isToggleTimeKey(key string) bool {
	return key == "tab"
}
