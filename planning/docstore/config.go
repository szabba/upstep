package docstore

import (
	"os"
)

func envOrElse(name, whenEmpty string) string {
	v := os.Getenv(name)
	if v == "" {
		return whenEmpty
	}
	return v
}