# Twitch Chat Bot
 customizable twitch chat bot built using Go.
 Utilizes the [twitch EventSub](https://dev.twitch.tv/docs/eventsub/) using websockets to retrieve events occuring in twitch chat channels.

 This is an attempted remake of [DaBolinkBot](https://github.com/dabolink/DaBolinkBot). while lacking a lot of the intial version this version attempts to build the bots in a more scalable way as well as removing a lot of the fallbacks of the intial version because of its implementation in Python (mostly around the GIL and loose typing)

## Entry point
### auth
Add ability to generate a access token from a code to be added to your `.env` file.

## Packages
### chat
merging of command and twitch packages
creates the bots and manages them.

runs commands per message recieved from `twitch`
### command
holds details around setting up commands and running commands against inputs.
#### counter
Command Specific counter that keeps track of a user and a number they typed, plays a "game" where the numbers must be counted sequentially and each user can only count a single number.
### env
Parse Environment variables from a .env file
### maps
Implementations for specific reusable map types
#### default map
Implements a map that if the value doesn't exist at a specific key uses a builder to initialize the value.

### twitch
twitch specific eventsub / websocket negotiations

handles twitch specific requests while passing forward text message events and other channel specific events.
### variables
Adds ability to store / retrieve variables by key while still attempting to keep some sembalance of typing.
### ws
Holds concerns around setting up and keeping a websocket connection alive as well as reading values from the websocket.

## TODOs
see [/docs/TODO.md](docs/TODO.md)