package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	. "github.com/scotow/notigo"
)

func main() {
	var opt options
	args, err := flags.Parse(&opt)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "cannot parse options:", err)
		os.Exit(1)
	}

	keys, err := findKeys(opt)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "cannot find IFTTT key(s):", err)
		os.Exit(1)
	}

	if len(keys) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "no IFTTT key(s) found")
		os.Exit(1)
	}

	messages, err := findMessages(opt, args)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "cannot find messages:", err)
		os.Exit(1)
	}

	if len(messages) == 0 {
		messages = []string{" "}
	}

	if opt.Concurrent {
		errorChan := make(chan error)

		for _, key := range keys {
			go func(key Key) {
				err := sendNotifications(key, opt.Event, opt.Title, messages, opt.Delay)
				errorChan <- err
			}(key)
		}

		var err error
		for range keys {
			err = <-errorChan
		}

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "cannot send notification:", err)
			os.Exit(1)
		}
	} else {
		for _, key := range keys {
			err := sendNotifications(key, opt.Event, opt.Title, messages, opt.Delay)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "cannot send notification:", err)
				os.Exit(1)
			}
		}
	}
}

func sendNotifications(key Key, event, title string, messages []string, delay time.Duration) error {
	if len(messages) == 0 {
		return nil
	}

	for _, message := range messages[:len(messages)-1] {
		err := key.SendEvent(NewNotification(title, message), event)
		if err != nil {
			return err
		}

		time.Sleep(delay)
	}

	// Send the last notification.
	err := key.SendEvent(NewNotification(title, messages[len(messages)-1]), event)
	if err != nil {
		return err
	}

	return nil
}
