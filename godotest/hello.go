package main

import "fmt"

func hello(user string) string {
	if user == "" {
		return "Hello unknown user!"
	}
	return fmt.Sprintf("Hello %v!", user)
}
