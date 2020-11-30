package main

import "github.com/pkg/errors"

var errNegative = errors.New("inputs cannot be negative")

func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errNegative
	}
	return a + b, nil
}
