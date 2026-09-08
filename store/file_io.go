package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const dataFileName = "data.json"

// A missing data file is treated as an empty store.
func load() (Data, error) {
	dir, err := dataPath()
	if err != nil {
		return Data{}, err
	}

	bytes, err := os.ReadFile(filepath.Join(dir, dataFileName))
	if errors.Is(err, os.ErrNotExist) {
		return Data{}, nil
	}
	if err != nil {
		return Data{}, fmt.Errorf("데이터 파일 읽어오는 중 에러: %w", err)
	}

	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return Data{}, fmt.Errorf("json 데이터로 변환 중 에러: %w", err)
	}

	return data, nil
}

func save(data Data) error {
	dir, err := dataPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("데이터 저장 경로 생성 중 에러: %w", err)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json 데이터를 구조체로 변환 중 에러: %w", err)
	}

	return os.WriteFile(filepath.Join(dir, dataFileName), bytes, 0600)
}
