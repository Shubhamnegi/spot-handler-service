package notification

import "os"

var (
	DEFAIULT_SLACK_URL      string = "#default"
	DEFAIULT_SLACK_USERNAME string = "#default"
	DEFAIULT_SLACK_CHANNEL  string = "#default"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
