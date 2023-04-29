# mytelebot
This is code for mytelebot. It is writen in Go. It uses 'TELE_TOKEN' env var for telegram bot API token.

You can access it by link

https://t.me/gotelebot_bot

Currently it supports next comands:

- /start hello  - shows version
- /get info - shows version

How to build bot:

Don`t forget to have GO installed on your system and get the token API from bot_father from telegram
```git clone https://github.com/pontarr/mytelebot.git ```

```cd mytelebot ```

``` read -s TELE_TOKEN ``` (Ctrl+V the token from telegram)

``` export TELE_TOKEN```

``` echo $TELE_TOKEN``` (Check the API token)

``` go build -ldflags "-X 'github.com/pontarr/mytelebot/cmd.appVersion=v1.0.2'" ```

Start the bot

``` ./mytelebot start ```