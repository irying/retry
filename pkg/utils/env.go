package utils

import (
	"os"
)

// GetEnvWithDefault will fetch a environment variable's value and assign
// a default value if the environment variable isn't set
func GetEnvWithDefault(key, fallback string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return fallback
}
