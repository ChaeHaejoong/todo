package store

import (
	"fmt"
	"os"
	"path/filepath"
)

// XDG_DATA_HOME takes precedence 
// When it's not set, the path falls back to $HOME/.local/share
func dataPath() (string, error) {
	if dataHome := os.Getenv("XDG_DATA_HOME"); dataHome != "" {
		return filepath.Join(dataHome, "todo"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("홈 디렉토리가 환경변수에 등록되어있지 않습니다.: %w", err)
	}

	return filepath.Join(homeDir, ".local", "share", "todo"), nil
}
