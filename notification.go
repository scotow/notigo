package notigo

import (
	"os"
)

// A Notification is a simple struct that group a title and content.
// Depending on the OS your are using, the length of the authorized content may vary.
// Sending more 1000 characters per Notification may be problematic.
type Notification struct {
	Title   string `json:"value1"`
	Message string `json:"value2"`
}

// Create a Notification.
// title and message are self-explanatory.
// If the title of the Notification is empty, NewMessage is called.
func NewNotification(title, message string) Notification {
	if title == "" {
		return NewMessage(message)
	} else {
		return Notification{
			Title:   title,
			Message: message,
		}
	}
}

// Create a new Notification with a specified content.
// message is the content of the Notification.
// This method use the hostname as title or "notigo" if the hostname could not be determined.
func NewMessage(message string) Notification {
	var title string

	hostname, err := os.Hostname()
	if err != nil {
		title = "Notigo"
	} else {
		title = hostname
	}

	return Notification{
		Title:   title,
		Message: message,
	}
}
