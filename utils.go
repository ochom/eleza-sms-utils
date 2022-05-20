package smsutils

import (
	"fmt"
	"os"
)

func getEnv(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf("env `%s` not defined", name)
	}

	return val, nil
}
