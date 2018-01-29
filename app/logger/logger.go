package logger

import (
	"log"
)

// LogDebug log some debug information
func LogDebug(msg string) {
	log.Printf("[DEBUG] %s\n", msg)
}

// LogDebug log supporting information
func LogInfo(msg string) {
	log.Printf("[INFO] %s\n", msg)
}

// LogWarn logs warning information
func LogWarn(msg string) {
	log.Printf("[WARN] %s\n", msg)
}

// LogError log errors, do not throw panic
func LogError(err error) {
	log.Printf("[ERROR] %s\n", err.Error())
}
