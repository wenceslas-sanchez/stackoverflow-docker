package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

// ToJSON saves a manifest as a JSON file.
func ToJSON(data any, path string) error {
	buf, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("issue format strcut to JSON: %q", err)
	}
	err = os.WriteFile(path, buf, 7777)
	if err != nil {
		return fmt.Errorf("can't write in file %v: %q", path, err)
	}

	return nil
}
