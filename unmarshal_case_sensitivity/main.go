package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Actions  string `json:"aktions"`
}

func main() {
	payload := []byte(`{
	"aktions": "first",
	"aKtIoNs": "second",
	"aKtionſ": "third"
}`)

	u := User{}
	err := json.Unmarshal(payload, &u)
	if err != nil {
		slog.Error("error using json.Unmarshal to retrieve action", "error", err)
		os.Exit(1)
	}
	slog.Info("json.Unmarshal", "action", u.Actions)
	if u.Actions != "first" {
		slog.Error("per the documentation the case-sensitive match should be preferred")
		os.Exit(1)
	}
}
