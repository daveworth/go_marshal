package main

import (
	"encoding/json"
	"encoding/xml"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Query struct {
	Actions string `json:"actions" yaml:"actions" xml:"actions"`
}

func main() {
	q := Query{}
	polyglot := []byte(`{
		"actions": "Action_1",
		"aCtionſ": "Action_2",
		"llkajkljskljskld": "<xml><actions>Action_3</actions></xml>"
}`)

	err := json.Unmarshal(polyglot, &q)
	if err != nil {
		slog.Error("error using json.Unmarshal to retrieve action", "error", err)
		os.Exit(1)
	}
	slog.Info("json.Unmarshal", "action", q.Actions)

	err = yaml.Unmarshal(polyglot, &q)
	if err != nil {
		slog.Error("error using json.Unmarshal to retrieve action", "error", err)
		os.Exit(1)
	}
	slog.Info("yaml.Unmarshal", "action", q.Actions)

	err = xml.Unmarshal(polyglot, &q)
	if err != nil {
		slog.Error("error using json.Unmarshal to retrieve action", "error", err)
		os.Exit(1)
	}
	slog.Info("xml.Unmarshal", "action", q.Actions)
}
