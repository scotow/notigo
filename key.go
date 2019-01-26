package notigo

import (
    "bytes"
    "errors"
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

const (
    endPointFormat  = "https://maker.ifttt.com/trigger/%s/with/key/%s"
    DefaultEvent    = "notigo"
)

type Key string

func (k *Key) Send(n Notification) (err error) {
    return k.SendEvent(n, DefaultEvent)
}

func (k *Key) SendEvent(n Notification, event string) (err error) {
    data, err := json.Marshal(&n)
    if err != nil {
        return
    }

    resp, err := http.Post(fmt.Sprintf(endPointFormat, event, string(*k)), "application/json", bytes.NewBuffer(data))
    if err != nil {
        return
    }

    if resp.StatusCode != http.StatusOK {
        content, err2 := ioutil.ReadAll(resp.Body)
        if err2 != nil {
            err = err2
            return
        }

        err = resp.Body.Close()
        if err != nil {
            err = err2
            return
        }

        err = errors.New(string(content))
        return
    }

    return
}
