package todomodal

func isCloseModalKey(key string) bool {
	return key == "esc"
}

func isAppendTodoKey(key string) bool {
	return key == "enter"
}
