package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ValidateThaiID("1234567890122"))
	fmt.Println(ValidateThaiID("1234567890121"))
}

func ValidateThaiID(id string) error {
	if len(id) != 13 {
		return errors.New("id digits incorrect")
	}

	splited := strings.Split(id, "")
	sum := 0
	for i, j := 0, 13; j > 1; i, j = i+1, j-1 {
		val, _ := strconv.Atoi(splited[i])
		sum = sum + val*j
	}

	moded := sum % 11
	result := 11 - moded

	last := result % 10
	lastID, _ := strconv.Atoi(splited[len(splited)-1])

	if last != lastID {
		return errors.New("id incorrect")
	}
	return errors.New("nil")
}
