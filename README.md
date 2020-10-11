# amongusmumble

A [Mumble](https://www.mumble.info/) bot for [Amoung US](http://www.innersloth.com/gameAmongUs.php) to manage muting/unmuting of players during the game.

Works in conjunction with [amonguscapture](https://github.com/denverquane/amonguscapture)


## Requirements

1. A mumble server.
2. Four mumble channels, AmongUs, Alive, Dead and Lobby. They need to be structured as follows:
![](images/MumbleChannels.jpg?raw=true)
3. A mumble registered user account for the bot. This account needs to have permission to mute/unmute users and move users between the Lobby/Dead/Alive channels.
4. The certificate for the mumble user your bot will use. You will need to convert this certificate to be in PEM format. See [Certificate Help](#Certificate Help)


## Install

1. Create folder for you bot.
2. Copy your cert and key for bot to your bot folder. See below for help on generating certicate help.
3. Download latest bot executable from [here](https://github.com/OperationalDev/amongusmumble/releases) and place it in the bot folder
4. Download v2.0.7 AmongUs capture executable from [here](https://github.com/denverquane/amonguscapture/releases) and place it in the bot folder.
5. Copy config.example to config and edit it. See [Config Example](#Config Example)
5. Run amoungusmumble.
6. Start Among Us.
7. Start Capture. Type in code 123456 and click connect.


## Build from source

1. Clonde repo
2. cd repo
3. go build .


## Certificate Help

Export your bot's certificate from mumble. This should be in a p12/pkcs format. Now convert this to a pem format with the following commands:
```
openssl pkcs12 -in botuser.p12 -nocerts -out botuser.key
openssl pkcs12 -in botuser.p12 -nocerts -nokeys -out botuser.crt
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
- [AmongUsDiscord](htps://github.com/denverquane/amongusdiscord)without their original project and capture tool, this would not be possible.


## License

[MIT](https://choosealicense.com/licenses/mit/)
