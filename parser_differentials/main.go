package main

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/buger/jsonparser"
)

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Action   string `json:"action"`
}

func main() {
	payload := []byte(`{
	"action": "first",
	"action": "second"
}`)
	action, _, _, err := jsonparser.Get(payload, "action")
	if err != nil {
		slog.Error("error using jsonparser to retrieve action", "error", err)
		os.Exit(1)
	}
	slog.Info("jsonparser.Get", "action", action)

	u := User{}
	err = json.Unmarshal(payload, &u)
	if err != nil {
		slog.Error("error using json.Unmarshal to retrieve action", "error", err)
		os.Exit(1)
	}
	slog.Info("json.Unmarshal", "action", u.Action)

	if string(action) != u.Action {
		slog.Error("value mismatch between parsers")
		os.Exit(1)
	}
}
