package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type File interface {
	ToString() (string, error)
	ToBytes() ([]byte, error)
}

func Save(toFile File, name string, perm os.FileMode) error {
	fileB, _ := json.MarshalIndent(toFile, "", " ")

	if err := os.WriteFile(name, fileB, perm); err != nil {
		return fmt.Errorf("can't save file: %q", err)
	}
	return nil
}
