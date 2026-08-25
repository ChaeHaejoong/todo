package main

import (
	"fmt"
	"os"

	"github.com/chaehaejoong/todo/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(app.New(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "프로그램 실행 실패:", err)
		os.Exit(1)
	}
}
