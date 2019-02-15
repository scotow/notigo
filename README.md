# Notigo

💬 Send iOS/Android notifications using IFTTT's webhook 💬

### IFTTT

From Wikipedia:

*[IFTTT](https://ifttt.com/) is a free web-based service to create chains of simple conditional statements, called applets. An applet is triggered by changes that occur within other web services such as Gmail, Facebook, Telegram, Instagram, or Pinterest.*

IFTTT proposes hundreads of triggers, but the one that Notigo uses is the [webhook](https://ifttt.com/maker_webhooks) trigger (also known as Maker Event).

By creating an IFTTT applet that send a rich notification to your device when a webhook is triggered, we can create a simple wrapper that call the specified URL to trigger it from a HTTP call.


### IFTTT account and mobile app

In order to receive a notification from IFTTT, you have to create an IFTTT [account](https://ifttt.com/join) and download the [iOS](https://itunes.apple.com/us/app/ifttt/id660944635?mt=8) app or the [Android](https://play.google.com/store/apps/details?id=com.ifttt.ifttt&hl=en) app. 


### Creating the IFTTT applet

Next you need to create the corresponding applet in your IFTTT account. Applets that use Webhook as a trigger can't be share like other applets, so you need to create it manually:

* Go to the applet [creation](https://ifttt.com/create) page;
* Search for `webhook` and select the `Receive a web request` trigger;
* Specify the name of the event (`notigo` is the default one used in the [example command](https://github.com/Scotow/notigo/tree/master/cmd/notigo));
* Click on `Create trigger`;
* For the `that` action, search for `notification` and select the `Send a rich notification from the IFTTT app` action;
* Use the `Add ingredient` button to add `value1` as a title and `value2` as a message. Leave the others blank.

The final configuration of the applet looks like this:

![Applet](applet.png?raw=true)

Once the applet created you can now use the library or the command example.

