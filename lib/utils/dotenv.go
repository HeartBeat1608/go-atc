package utils

import (
	"os"
	"strings"
)

func LoadDotEnv(filename string) error {
	if filename == "" {
		filename = ".env"
	}

	byteBuffer, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(byteBuffer), "\n") {
		parts := strings.Split(line, "=")
		key, value := parts[0], strings.Join(parts[1:], "")
		if key != "" {
			if err = os.Setenv(key, value); err != nil {
				return err
			}
		}
	}

	return nil
}
