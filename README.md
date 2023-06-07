# mytelebot
This is code for mytelebot. It is writen in Go. It uses 'TELE_TOKEN' env var for telegram bot API token.
Also uses api http://api.openweathermap.org for displaying temparature in K in the city

You can access it by link

https://t.me/gotelebot_bot

Currently it supports next comands:

- /start hello  - shows version
- /get info - shows info about the bot
- /get <city_name> - show weather in the city now

How to build bot:

Don`t forget to have GO installed on your system and get the token API from bot_father from telegram
```git clone https://github.com/pontarr/mytelebot.git ```

```cd mytelebot ```

``` read -s TELE_TOKEN ``` (Ctrl+V the token from telegram)

``` export TELE_TOKEN```

``` echo $TELE_TOKEN``` (Check the API token)

Also repeat steps for Weather API variable WEATHER_API

``` go build -ldflags "-X 'github.com/pontarr/mytelebot/cmd.appVersion=v1.0.3'" ```

Start the bot

``` ./mytelebot start ```

Access in TG https://t.me/gotelebot_bot

``` /get Tokio ```

will show the temparature in Tokio now

## The Workflow
The workflow for app CI\CD is like
Git push -> Git Hub -> Git Actions on push -> Push to repository ghcr.io 
On Changes the argocd see the push and redeploys the application on k3d cluster.

![Image](/cicd.drawio.png) 