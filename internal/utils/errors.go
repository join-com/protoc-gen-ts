package utils

import (
	"log"
	"strings"
)

func LogError(messages ...string) {
	paddedMessages := make([]string, len(messages))
	for i, msg := range messages {
		paddedMessages[i] = "    " + msg
	}
	paddedMsg := strings.Join(paddedMessages, "\n")

	log.Fatal("protoc-gen-ts:\n" + paddedMsg + "\n")
}
