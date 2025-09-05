package main

import (
	"fmt"
	"os"
	"time"
)

type stderrActionsLoggerStruct struct{}

func newStderrActionsLogger() *stderrActionsLoggerStruct {
	logger := &stderrActionsLoggerStruct{}
	return logger
}

func (*stderrActionsLoggerStruct) LogActionError(timestamp time.Time, message string, actionInvocationId string, action string) {
	fmt.Fprintf(os.Stderr, "%s [ACTION_ERROR] %s (action invocation id: %s, action: %s)\n", timestamp.Local().Format(time.Kitchen), message, actionInvocationId, action)
}
