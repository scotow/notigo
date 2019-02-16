package main

import (
	"github.com/scotow/notigo"
	"strings"
)

type keys []notigo.Key

func (k *keys) String() string {
	if k == nil {
		return ""
	} else {
		keyStrings := make([]string, 0, len(*k))
		for _, key := range *k {
			keyStrings = append(keyStrings, string(key))
		}
		return strings.Join(keyStrings, " ")
	}
}

func (k *keys) Set(value string) error {
	*k = append(*k, notigo.Key(value))
	return nil
}
