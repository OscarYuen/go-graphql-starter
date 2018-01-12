package service

import (
	"encoding/base64"
	"fmt"
	graphql "github.com/neelance/graphql-go"
	"strconv"
	"strings"
)

func EncodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1))))
}

func DecodeCursor(after *string) (*int, error) {
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
