package main

import (
    "fmt"
    "os"
    "os/user"
    "io/ioutil"
    "bufio"
    "flag"
    "strings"
    "path/filepath"
    "github.com/scotow/notigo"
)

const (
    keySubPath = ".config/notigo/config"
)

var (
    keyFlag     = flag.String("k", "", "IFTTT authentification key, ~/.config/notigo if not set")
    titleFlag   = flag.String("t", "", "notification title")
)

func findKey() (key notigo.Key, err error) {
    if *keyFlag != "" {
        key = notigo.Key(*keyFlag)
        return
    }

    user, err := user.Current()
    if err != nil {
        return
    }

    keyPath := filepath.Join(user.HomeDir, keySubPath)

    file, err := os.Open(keyPath)
    if err != nil {
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()

    key = notigo.Key(strings.TrimSpace(scanner.Text()))
    err = scanner.Err()

    return
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
    flag.Parse()

    key, err := findKey()
    if err != nil {
        fmt.Fprintln(os.Stderr, "cannot get API key:", err)
        os.Exit(1)
    }

    message, err := getMessage()
    if err != nil {
        fmt.Fprintln(os.Stderr, "cannot parse the message:", err)
        os.Exit(1)
    }

    if message == "" {
        message = "[No content]"
    }

    err = key.Send(notigo.NewNotification(*titleFlag, message))
    if err != nil {
        fmt.Fprintln(os.Stderr, "cannot send notification:", err)
        os.Exit(1)
    }
}
