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
	exitWithTextIfError("cannot parse options:", err)

	keys, err := findKeys(opt)
	exitWithTextIfError("cannot find IFTTT key(s):", err)

	if len(keys) == 0 {
		exitWithText("no IFTTT key(s) found")
	}

	messages, err := findMessages(opt, args)
	exitWithTextIfError("cannot find messages:", err)

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

		hasError := false
		for range keys {
			err := <-errorChan
			if err != nil {
				printError("cannot send notification:", err)
				hasError = true
			}
		}

		if hasError {
			os.Exit(1)
		}
	} else {
		for _, key := range keys {
			err := sendNotifications(key, opt.Event, opt.Title, messages, opt.Delay)
			exitWithTextIfError("cannot send notification:", err)
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

func printError(args ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, args...)
}

func exitWithText(args ...interface{}) {
	printError(args...)
	os.Exit(1)
}

func exitWithTextIfError(text string, err error) {
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		exitWithText(text, err.Error())
	}
}
