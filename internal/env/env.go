package env

import (
	"encoding/json"
	"os"
)

func Load[T any](config *T) error {
	data, err := os.ReadFile(".env")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &config)

}
