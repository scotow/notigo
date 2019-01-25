package notigo

import (
    "bytes"
    "errors"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

const (
    endPoint = "https://maker.ifttt.com/trigger/notigo/with/key/"
)

type Key string

func (k *Key) Send(n Notification) (err error) {
    data, err := json.Marshal(&n)
    if err != nil {
        return
    }

    resp, err := http.Post(endPoint + string(*k), "application/json", bytes.NewBuffer(data))
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
