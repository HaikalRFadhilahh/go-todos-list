package helper

import "os"

func GetEnv(k string, df string) string {
	if os.Getenv(k) == "" {
		return df
	}

	return os.Getenv(k)
}
