package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/scotow/notigo"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	keySubPath = ".config/notigo/config"
)

var (
	keysFlag keys
	//keyFlag     = flag.String("k", "", "IFTTT authentication key, ~/.config/notigo if not set")
	eventFlag = flag.String("e", notigo.DefaultEvent, "event name")
	titleFlag = flag.String("t", "", "notification title")
)

func findKeys() (keys, error) {
	if len(keysFlag) != 0 {
		return keysFlag, nil
	}

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	keyPath := filepath.Join(usr.HomeDir, keySubPath)

	file, err := os.Open(keyPath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var k keys
	for scanner.Scan() {
		k = append(k, notigo.Key(strings.TrimSpace(scanner.Text())))
	}

	err = scanner.Err()
	if err != nil {
		_ = file.Close()
		return nil, err
	}

	err = file.Close()
	return k, err
}

func getMessage() (message string, err error) {
	args := flag.Args()
	if len(args) > 0 {
		message = strings.Join(args, " ")
		return
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	message = string(bytes)
	return
}

func main() {
	flag.Var(&keysFlag, "k", "IFTTT authentication key(s), ~/.config/notigo if not set")
	flag.Parse()

	keys, err := findKeys()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "cannot get API key:", err)
		os.Exit(1)
	}

	message, err := getMessage()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "cannot parse the message:", err)
		os.Exit(1)
	}

	if message == "" {
		message = "[No content]"
	}

	for _, key := range keys {
		err = key.SendEvent(notigo.NewNotification(*titleFlag, message), *eventFlag)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "cannot send notification:", err)
			os.Exit(1)
		}
	}
}
