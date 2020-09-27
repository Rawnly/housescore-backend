package util

import "os"

type Env struct{}

func (e *Env) Get(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}
	return value
}
