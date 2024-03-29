# ~~RestFull API server for Minecraft (JAVA Edition) Rcon Server writen in Go~~

~~Moreo provide a web proxy server than interact with your minecraft server rcon service for you by an RestFull API.~~

#
# Web Proxy server for Mineraft (JAVA Edition) Rcon Server written in Go
After some works on this project, i realized that build rest full api for each functions is boring and limited.
- Boring to write exactly some code for each functions (test token, parse data, send to rcon server)
- Limited because somes Minecraft functions are very complex to implement and obiously, i had to take some decisions to make it works

So i decided to write only one function who receives any minecraft rcon commands in POST and send it directly to rcon server.
This new way allows to use unofficial Minecraft Server like Spigot because we are not limited to some functions now.
And if a new rcon function appears, it will be automaticly supported by Moreo.

#
# How it works ?
Moreo authentificate users by checking user/password in small database file (explication bellow) and by sending token after.
Once you're authentified, you send both token (authentification bearer) and POST query with your command, Moreo do the rest.
It will send your command to your Minecraft rcon server and return you result in body message.
Each token spreads **one rcon server**. That's mean you can interact with many differents servers with the same account in same times, with differents tokens (Moreo keep rcon authentification in memory and reuse it each time you need).

For managing rcon connection, Moreo use [Will Robert's Minecraft-Client](http://github.com/willroberts/minecraft-client)

#
# Users Database
Database for identify users is a simple file in the directory of your choice. Format use is :
```
username:password<newline>
```

For example :
```
toto:123some
tata:456magic
titi:678password
```
#
# ENV Variables
For a docker usage (*i'm still work on it*), Moreo's configuration is from environment variable.

### Must be initialized
```
MOREO_WEB_PORT  = local port where Moreo listen
MOREO_DB_PATH   = local db file (description above)
```

### Extra options (can be empty or uninitialized)
```
MOREO_WEB_IP        = IP to make Moreo listening on (by default, all availables)
MOREO_TOKEN_LIFE    = token life (by default 15 minutes)
MOREO_DEBUG         = to active debugs routes and functionnalities
```

#
# Authentification

To authenticated user and get access token, you need to post data in json format

```json
{
    "user": {
        "name": "",     // user name in local db file
        "password": ""  // user's password
    },
    "rcon": {
        "host": "",     // rcon name or ip
        "port": 25575,  // rcon port
        "password": ""  // password set in server.properties
    }
}
```
#
# Send command

To send command, post data in json format with authentification bearer token

```json
{
    "cmd": "any minecraft rcon commands available"
}
```

#
# Debug connection

To debug rcon connection, you can enable with env **MOREO_DEBUG=1**.
By using this, you disable Moreo's authentification process and send POST commands to rcon directly with the url **http(s)://[DOMAIN:PORT]/debug/rcon/post** and this structure (json format)

```json
{
    "host": "",         // rcon domain or ip
    "port": 25575,      // rcon port
    "password": "",     // rcon password (server.properties)
    "cmd": ""           // command to be executed
}
```

#
# Tokens
Moreo use my own [jwt package implementation](https://github.com/pijee/pjwt).
When starting, Moreo generate random 512 bits key to sign tokens. Every 2 hours, this key is regenerated and the tokens db is truncated.

#
# Example (with curl)

A small example with curl on terminal. I have a spigot server running on docker with rcon opened on my local computer (localhost 25575 "password")

>Starting Moreo
```console
$ export MOREO_WEB_PORT=8080
$ export MOREO_DB_FILE=./users.db
$ ./moreo
```
>Authentification using curl
```console
$ curl -v -X POST -H "Content-Type: application/json" -d '{ "user":{ "name": "toto", "password": "123"}, "rcon":{ "host": "127.0.0.1", "port": 25575, "password": "password"}}' http://127.0.0.1:8080/login
```

>Curl result
```console
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> POST /login HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.76.1
> Accept: */*
> Content-Type: application/json
> Content-Length: 116
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 23 Apr 2021 11:19:42 GMT
< Content-Length: 210
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 127.0.0.1 left intact
eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiZjUxMGM0N2ItNGE4OC00NWM2LTUyNzQtMDU1MmNhM2E2OTNlIiwiZXhwIjoxNjE5MTc3NjgyfQ.dSBerBJKKRgGvUxexP70jKJL1vgbgWgIawtzoWZ_LV-7EkE0ZWju0dvDBZJg0PqI-YdehfbIT6jaOmg0xVhYQg
```

>It's ok, i 've got my token, continue ...

```console
$ curl -v -X POST -H "Content-Type: application/json" -H "Authorization: bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiZjUxMGM0N2ItNGE4OC00NWM2LTUyNzQtMDU1MmNhM2E2OTNlIiwiZXhwIjoxNjE5MTc3NjgyfQ.dSBerBJKKRgGvUxexP70jKJL1vgbgWgIawtzoWZ_LV-7EkE0ZWju0dvDBZJg0PqI-YdehfbIT6jaOmg0xVhYQg" -d '{ "cmd": "help" }' http://127.0.0.1:8080/post
```

>Rcon response
```console
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 23 Apr 2021 11:23:39 GMT
< Content-Length: 434
< Content-Type: text/plain; charset=utf-8
<
§e--------- §fHelp: Index (1/10) §e--------------------
§7Use /help [n] to get page n of help.
§6Aliases: §fLists command aliases
§6Bukkit: §fAll commands for Bukkit
§6Minecraft: §fAll commands for Minecraft
§6/advancement: §fA Mojang provided command.
§6/attribute: §fA Mojang provided command.
§6/ban: §fA Mojang provided command.
§6/ban-ip: §fA Mojang provided command.
§6/banlist: §fA Mojang provided command.
* Connection #0 to host 127.0.0.1 left intact
```

>Let's try some other command
```console
$ curl -v -X POST -H "Content-Type: application/json" -H "Authorization: bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiZjUxMGM0N2ItNGE4OC00NWM2LTUyNzQtMDU1MmNhM2E2OTNlIiwiZXhwIjoxNjE5MTc3NjgyfQ.dSBerBJKKRgGvUxexP70jKJL1vgbgWgIawtzoWZ_LV-7EkE0ZWju0dvDBZJg0PqI-YdehfbIT6jaOmg0xVhYQg" -d '{ "cmd": "say Hello World !" }' http://127.0.0.1:8080/post
```

>Curl result
```console
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 23 Apr 2021 11:24:44 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact
```

>Minecraft server console (docker logs -f [container_name])
```console
[11:24:44] [RCON Listener #1/INFO]: Thread RCON Client /172.19.0.1 started
[11:24:44] [Server thread/INFO]: [Rcon] Hello World !
[11:24:44] [RCON Client /172.19.0.1 #10/INFO]: Thread RCON Client /172.19.0.1 shutting down
```

>In game

![screen1](./images/screen1.png)

#
# Timeout

All tokens uuid are referenced in internal memory database with timestamp.
An internal go routine is running each X minutes (by default : 5) and deletes all tokens outdated. You can make token'life much longer by using **MOREO_TOKEN_LIFE** env variable.

## Example
If i'm waiting few minutes and send same request

```console
$ curl -v -X POST -H "Content-Type: application/json" -H "Authorization: bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiZjUxMGM0N2ItNGE4OC00NWM2LTUyNzQtMDU1MmNhM2E2OTNlIiwiZXhwIjoxNjE5MTc3NjgyfQ.dSBerBJKKRgGvUxexP70jKJL1vgbgWgIawtzoWZ_LV-7EkE0ZWju0dvDBZJg0PqI-YdehfbIT6jaOmg0xVhYQg" -d '{ "cmd": "say Hello World !" }' http://127.0.0.1:8080/post
```

>Curl result
```console
* Mark bundle as not supporting multiuse
< HTTP/1.1 403 Forbidden
< Date: Fri, 23 Apr 2021 11:34:53 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact
```

>403 sent because my token is outdated


#
# Proxy Errors
>Send command directly without authentification
```console
$ curl -v -X POST -H "Content-Type: application/json" -d '{"cmd": "say Hello World !"}' http://127.0.0.1:8080/post
```

>Curl result
```console
$ * Mark bundle as not supporting multiuse
< HTTP/1.1 407 Proxy Authentication Required
< Date: Fri, 23 Apr 2021 12:23:23 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact
```

> 407 : proxy authentification required
> Try now  with wrong user password
```console
$ curl -v -X POST -H "Content-Type: application/json" -d '{ "user":{ "name": "tata", "password": "456"}, "rcon":{ "host": "127.0.0.1", "port": 25575, "password": "password"}}' http://127.0.0.1:8080/login
```

>Curl result
```console
* Mark bundle as not supporting multiuse
< HTTP/1.1 403 Forbidden
< Date: Fri, 23 Apr 2021 12:25:39 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact
```
> error 403

#
# Rcon errors
If you send wrong rcon commands, this is the Minecraft rcon server who will responds to you in body part, not Moreo. You have to parse responses so.

```console
$ curl -v -X POST -H "Content-Type: application/json" -H "Authorization: bearer XXX" -d '{ "cmd": "Hello there !"}' http://127.0.0.1:8080/post
```

>Curl response
```console
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Fri, 23 Apr 2021 12:34:12 GMT
< Content-Length: 40
< Content-Type: text/plain; charset=utf-8
<
Unknown command. Type "/help" for help.
* Connection #0 to host 127.0.0.1 left intact
```

#
# Work in progress

### Server basics
- [x] Authentification
- [x] Token creation and verification
- [x] User login
- [x] Debug Rcon Connection
- [x] Send commands to rcon server
- [ ] Dockerfile
- [ ] Write small simple html client