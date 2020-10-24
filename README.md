# AmongUs Mumble

A [Mumble](https://www.mumble.info/) bot for [Amoung US](http://www.innersloth.com/gameAmongUs.php) to manage muting/unmuting of players during the game.

Works in conjunction with [amonguscapture](https://github.com/denverquane/amonguscapture)

## How it works
Players join the lobby channel. When the game starts, players are moved to the alive channel and are muted/deafened. When a body is reported or the emergency button is pressed, players that are not dead are unmuted/undeafened. Players that are dead, stay muted, but are undeafened so they can hear what's going on. When the game resumes, alive players are once again muted/deafened while dead players are moved to the dead channel and can communicate freely.


## Requirements

1. A mumble server.
2. Four mumble channels, AmongUs, Alive, Dead and Lobby. They need to be structured as follows:
![](images/MumbleChannels.jpg?raw=true)
3. A mumble registered user account for the bot. This account needs to have permission to mute/unmute users and move users between the Lobby/Dead/Alive channels.
4. The certificate for the mumble user your bot will use. You will need to convert this certificate to be in PEM format. See [Certificate Help](#certificate-help)


## Install

1. Create folder for you bot.
2. Copy your cert and key for bot to your bot folder. See [below](#certificate-help) for help on generating certicate help.
3. Download latest bot executable from [here](https://github.com/OperationalDev/amongusmumble/releases) and place it in the bot folder
4. Download v2.0.7 AmongUs capture executable from [here](https://github.com/denverquane/amonguscapture/releases) and place it in the bot folder.
5. Copy config.example to config and edit it. See [Config Example](#config-example)
5. Run amoungusmumble.
6. Start Among Us.
7. Start Capture. Type in code 123456 and click connect.

## Usage

Once the bot and capture are running, players need to join the lobby channel and join the lobby of the game. In mumble, players need to set a comment that is the same as their in game name. E.g. If my mumble user is Bob and my in game name is "Not the Imposter", then my mumble user should have the comment set, Not the Imposter.

![](images/MumbleComment.jpg?raw=true)


## Build from source

1. Clone repo
2. cd repo
3. go build .


## Certificate Help

Export your bot's certificate from mumble (See this [video](https://www.typefrag.com/mumble/tutorials/backup-or-import-certificate/) for help). This should be in a p12/pkcs format. Now convert this to a pem format with the following commands:
```
openssl pkcs12 -info -in botuser.p12 -nodes -out 6ix9ine.key -nocerts
openssl pkcs12 -info -in botuser.p12 -nodes -out botuser.crt -nokeys
```

## Config Example

```
cert: "certname.crt" # Certificate for your mumble user.
key: "cername.key" # Contains private key for your mumble user certificate.
listenaddress: "0.0.0.0" # The address amongusmumble will listen on. You probably don't need to change this.
listenport: "8123" # The port amongusmumble will listen on. You probably don't need to change this.
mumbleserver: "mymumbleserver.com:64738" # Your mumble server and port.
username: "botname" # Your bot's username
```

## Similar Projects
- [AmongUsDiscord](https://github.com/denverquane/amongusdiscord) without their original project and capture tool, this would not be possible.


## License

[MIT](https://choosealicense.com/licenses/mit/)
