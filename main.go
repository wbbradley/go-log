package main

import (
	"encoding/json"
	"log"
)

type LogEntry struct {
	Details map[string]interface{}
}

func (le *LogEntry) Add(k string, val interface{}) {
	le.Details[k] = val
}

func (le *LogEntry) Report() {
	b, err := json.Marshal(le.Details)
	if err == nil {
		log.Println(string(b[:]))
	}
}

type L map[string]interface{}
type LogEntryFactory func(details L) *LogEntry

func newLogType(level string) LogEntryFactory {
	return func(details L) *LogEntry {
		logEntry := LogEntry{Details: details}
		logEntry.Details["level"] = level
		return &logEntry
	}
}

var (
	Info  = newLogType("INFO")
	Error = newLogType("ERROR")
	Warn  = newLogType("WARN")
)

func foo() {
	logEntry := Info(L{"foo": 1, "bar": 2})
	defer logEntry.Report()

	if 1 < 3 {
		logEntry.Add("baz", 3)
	}
}

func main() {
	foo()
}
