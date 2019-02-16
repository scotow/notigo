# Notigo

💬 Send iOS/Android notifications using IFTTT's Webhook 💬

```
go get -u github.com/Scotow/notigo
```


## IFTTT

From Wikipedia:

*[IFTTT](https://ifttt.com/) is a free web-based service to create chains of simple conditional statements, called applets. An applet is triggered by changes that occur within other web services such as Gmail, Facebook, Telegram, Instagram, or Pinterest.*

IFTTT proposes hundreds of triggers, but the one that Notigo uses is the [Webhook](https://ifttt.com/maker_webhooks) trigger (also known as Maker Event).

By creating an IFTTT applet that send a rich notification to your device when a Webhook is triggered, we can create a simple wrapper that call the specified URL to trigger it from a HTTP call.


## IFTTT account and mobile app

In order to receive a notification from IFTTT, you have to create an IFTTT [account](https://ifttt.com/join) and download the [iOS](https://itunes.apple.com/us/app/ifttt/id660944635?mt=8) app or the [Android](https://play.google.com/store/apps/details?id=com.ifttt.ifttt&hl=en) app. 


## Creating the IFTTT applet

Next you need to create the corresponding applet in your IFTTT account. Applets that use Webhook as a trigger can't be share like other applets, so you need to create it manually:

* Go to the applet [creation](https://ifttt.com/create) page;
* Search for `webhook` and select the `Receive a web request` trigger;
* Specify the name of the event (`notigo` is the default one used in the [command example](https://github.com/Scotow/notigo/tree/master/cmd/notigo));
* Click on `Create trigger`;
* For the `that` action, search for `notification` and select the `Send a rich notification from the IFTTT app` action;
* Use the `Add ingredient` button to add `value1` as a title and `value2` as a message. Leave the others blank.

The final configuration of the applet looks like this:

![Applet](applet.png?raw=true)


## Getting the Webhook key

The last step before using the applet is to get your Webhook key. Head to the [Webhook settings page](https://ifttt.com/services/maker_webhooks/settings) and get the key from the URL.

For example, if the URL is `https://maker.ifttt.com/use/eHolJ7y7b8KVk4wUgZS6mY`, your key is `eHolJ7y7b8KVk4wUgZS6mY`.

NB: Clicking the `Edit connection` button reset the event key, but the old key seems to be usable for a while.

Now that you have created the applet and got your Webhook key, you can use the library or the example command.


## Using the library

Here is a simple example that send a notification with "Test" as a title and "Hello from GitHub" as a message:

```go
package main

import (
	"fmt"
	"log"

	"github.com/scotow/notigo"
)

func main() {
	notification := notigo.NewNotification("Test", "Hello from GitHub")
	key := notigo.Key("eHolJ7y7b8KVk4wUgZS6mY")

	err := key.Send(notification)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Notification sent.")
}
```

You can use the `func (k *Key) SendEvent(n Notification, event string) (err error)` and specify and custom event name if you registered a different one while creating the applet.

Using an empty string as a title or using the `func NewMessage(message string) Notification` function will try to use the hostname of the machine.


## Using the command

This repo has a simple [command](https://github.com/Scotow/notigo/tree/master/cmd/notigo) that allows you to send a notification from your favorite shell.

```
Usage of notigo:
  [-e EVENT]
    	event name (default "notigo")
  [-k KEY]...
    	IFTTT authentication key(s), ~/.config/notigo if not set
  [-t TITLE]
    	notification title
```

The command will use the Webhook key(s) stored in the `~/.config/notigo` if it exists. If it doesn't exist you have to specify your key(s) by using the `-k KEY` option.

The default event name is `notigo` and can be changed with the `-e EVENT` option.

You can specify a title for the notification by using the `-t TITLE` option.

The body of the notification is the concatenation of the remaining arguments or read from `STDIN` if there is none.

***Enjoy simple notifications!***
