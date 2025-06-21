package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"omitempty"`
	// IsAdmin  bool   `json:"omitempty,"` // This also fires the semgrep rule
}

func main() {
	u := User{}
	err := json.Unmarshal([]byte(`{
    "username": "jofra",
    "password": "qwerty123!",
    "omitempty": true
}`), &u)
	if err != nil {
		slog.Error("error unmarshalling JSON string", "error", err)
		return
	}
	fmt.Printf("Result: %#v\n", u)
	if u.IsAdmin {
		slog.Error("IsAdmin overwritten by malicious user input")
		os.Exit(1)
	}
}
