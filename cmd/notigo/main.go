package main

import (
    "fmt"
    "os"
    "github.com/scotow/notigo"
)

func main() {
    key := notigo.Key("b-1dH-8C5dW8clrbLHoOWa")

    err := key.Send(notigo.NewMessage("hello"))
    if err != nil {
        fmt.Fprintln(os.Stderr, "cannot send notification:", err)
        os.Exit(1)
    }

    fmt.Println(err)
}
