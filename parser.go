package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParsePortRange(input string) (int, int, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return -1, -1, fmt.Errorf("invalid format")
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil || start < 0 || end < start {
		return -1, -1, fmt.Errorf("invalid range")
	}

	return start, end, nil
}
