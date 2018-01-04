package service

import (
	"encoding/base64"
	"strconv"
	"strings"
)

func decodeCursor(after *string) (*int, error) {
	decodedValue := defaultDecodedIndex
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(string(*after))
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return nil, err
		}
		decodedValue = i
	}
	return &decodedValue, nil
}