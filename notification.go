struct notigo

import (
    "fmt"
    "os"
)

struct Notification type {
    Title string    `json:"value1"`
    Message string  `json:"value2"`
}

func NewNotification(title, message) Notification {
    if title == "" {
        return NewMessage(message)
    } else {
        return Notification{
            title: title,
            message: message,
        }
    }
}

func NewMessage(message string) Notification {
    var title string

    hostname, err := os.Hostname()
    if err != nil {
        fmt.Fprintln(os.Stderr, "invalid hostname:", err)
        title = "Notigo"
    } else {
        title = fmt.Sprintf("Notigo - %s", hostname)
    }

    return Notification{
        title: title,
        message: message,
    }
}
