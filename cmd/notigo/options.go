package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	. "github.com/scotow/notigo"
)

const (
	keySubPath = ".config/notigo/keys"
)

type options struct {
	Keys       []string      `short:"k" long:"key" description:"List of key(s) to use" value-name:"KEY"`
	KeysPaths  []string      `short:"K" long:"keys-path" description:"List of file path(s) that contains key(s)" value-name:"PATH"`
	Event      string        `short:"e" long:"event" description:"Event key passed to IFTTT" default:"notigo" value-name:"EVENT"`
	Title      string        `short:"t" long:"title" description:"Title of the notification(s)" value-name:"TITLE"`
	Files      []string      `short:"f" long:"file" description:"List of file(s) used for content" value-name:"PATH"`
	Merge      bool          `short:"m" long:"merge" description:"Content should be merged"`
	MergeSep   string        `short:"s" long:"merge-separator" description:"Separator used while merging content" default:"\n" value-name:"SEPARATOR"`
	Delay      time.Duration `short:"d" long:"delay" description:"Delay between two notification" default:"3s" value-name:"DELAY"`
	Concurrent bool          `short:"c" long:"concurrent" description:"Concurrently send notifications to the keys"`
}

func findKeys(opt options) ([]Key, error) {
	var keys []Key

	for _, key := range opt.Keys {
		keys = append(keys, Key(key))
	}

	// Parse files in specified files.
	for _, keyPath := range opt.KeysPaths {
		found, err := keysFromFile(keyPath)
		if err != nil {
			return nil, err
		}
		keys = append(keys, found...)
	}

	// Fallback to home directory config file.
	if len(opt.Keys)+len(opt.KeysPaths) == 0 {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}

		keysPath := filepath.Join(usr.HomeDir, keySubPath)
		if exists(keysPath) {
			found, err := keysFromFile(keysPath)
			if err != nil {
				return nil, err
			}
			keys = append(keys, found...)
		}
	}

	return keys, nil
}

func keysFromFile(path string) ([]Key, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var keys []Key
	for scanner.Scan() {
		keys = append(keys, Key(strings.TrimSpace(scanner.Text())))
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func findMessages(opt options, args []string) ([]string, error) {
	var messages []string

	// Create a message from remaining arguments.
	data := strings.TrimSpace(strings.Join(args, " "))
	if len(data) > 0 {
		messages = append(messages, data)
	}

	// Get messages from files.
	for _, path := range opt.Files {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		data := strings.TrimSpace(string(bytes))
		if len(data) > 0 {
			messages = append(messages, string(bytes))
		}
	}

	// Fallback to stdin.
	if len(args)+len(opt.Files) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}

		data := strings.TrimSpace(string(bytes))
		if len(data) > 0 {
			messages = append(messages, string(bytes))
		}
	}

	if opt.Merge && len(messages) > 0 {
		messages = []string{strings.Join(messages, opt.MergeSep)}
	}

	return messages, nil
}
