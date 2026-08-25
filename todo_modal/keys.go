package todomodal

func isCloseModalKey(key string) bool {
	if key == "esc" {
		return true
	}
	return false
}

func isAppendTodoKey(key string) bool {
	if key == "enter" {
		return true
	}
	return false
}
