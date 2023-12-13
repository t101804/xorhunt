package runner

import (
	"errors"
	"strings"
)

func sanitize(data string) (string, error) {
	data = strings.Trim(data, "\n\t\"' ")
	if data == "" {
		return "", errors.New("empty data")
	}
	return data, nil
}
