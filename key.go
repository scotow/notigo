package notigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	endPointFormat = "https://maker.ifttt.com/trigger/%s/with/key/%s"

	// The default event name used to communicated with the IFTTT's Webhook service.
	DefaultEvent = "notigo"
)

// Key is the key given by IFTTT while configuring the Webhooks service.
// Check the README for more information on how to get a API key.
type Key string

// Send the specified Notification using the default event key ("notigo").
// n is the Notification that will be send.
// This method returns any error that may have occurred while sending the Notification.
// The errors are mainly due to a wrong API key or a network problem.
func (k *Key) Send(n Notification) error {
	return k.SendEvent(n, DefaultEvent)
}

// Send the specified Notification using a custom event key.
// n is the Notification that will be send. event is the event name passed to IFTTT.
// This method returns any error that may have occurred while sending the Notification.
// The errors are mainly due to a wrong API key or a network problem.
func (k *Key) SendEvent(n Notification, event string) error {
	data, err := json.Marshal(&n)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf(endPointFormat, event, string(*k)), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = resp.Body.Close()
		if err != nil {
			return err
		}

		err = errors.New(string(content))
		return err
	}

	return nil
}
