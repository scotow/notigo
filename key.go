package notigo

import (
    "encoding/json"
)

const (
    endPoint- = "https://maker.ifttt.com/trigger/notigo/with/key/"
)

type Key string

func (k *Key) Send(n Notification) err error {
    data, err := json.Marshal(&n)
    if err != nil {
        return
    }

    resp, err := http.Post(fmt.Sprintf(endPoint, string(*k)), "application/json", bytes.NewBuffer(data))
    if err != nil {
        return
    }

    return
}
