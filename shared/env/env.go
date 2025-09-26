package env

import (
	"os"
)

func GetEnv(key string, defultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defultValue
	}
	return value
}
