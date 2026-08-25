package app

func isQuitKey(key string) bool {
	if key == "q" {
		return true
	}
	return false
}

func isOpenTodoModalKey(key string) bool {
	if key == "a" {
		return true
	}
	return false
}
