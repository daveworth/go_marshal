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
	IsAdmin  bool   `json:"-,"`
}

func main() {
	u := User{}
	err := json.Unmarshal([]byte(`{
    "username": "jofra",
    "password": "qwerty123!",
    "-": true
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
