package app

import (
	"os/exec"

	tea "charm.land/bubbletea/v2"
	"github.com/chaehaejoong/todo/internal/focus"
)

func fcitxCmd(option string) tea.Cmd {
	return func() tea.Msg {
		_ = exec.Command("fcitx5-remote", option).Run()
		return nil
	}
}

func isComponentFocus(target focus.Target) bool {
	switch target {
	case focus.TodoList, focus.Calendar, focus.Clock, focus.Appointment:
		return true
	default:
		return false
	}
}
