package main

import (
	"log"
	"os"
	"strconv"
)

// Getenv retrieves an environment variable, or default if not found
func Getenv(name string, defaultVal string) string {
	s := os.Getenv(name)
	if s == "" {
		return defaultVal
	}
	return s
}

// GetenvBool retrieves an environment variable as a bool, or default if not found
func GetenvBool(name string, defaultVal bool) bool {
	s := os.Getenv(name)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatalf("Failed to parse environment var %s, err: %s", name, err)
	}
	return v
}
