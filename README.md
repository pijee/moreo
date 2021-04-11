# RestFull API server for Minecraft (JAVA Edition) Rcon Server writen in Go

Moreo provide a web proxy server than interact with your minecraft server's rcon service for you by an RestFull API.

# How it works ?
Moreo authentificate users by small database file (explication bellow).
Once you're connected, Moreo send you token to interact with API (like a normal RestFUll API).
Each token spread **one rcon server**. That's mean you can interact with many differents servers with the same account in same times, with differents tokens.
Moreo execute command on rcon server for you and send you result.

For managing connection, Moreo use [Will Robert's Minecraft-Client](http://github.com/willroberts/minecraft-client)

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

# List of commands
I used [minecraft wiki commands](https://minecraft.fandom.com/wiki/Commands) to implement commands supported by the API.

# Work in progress

***Some functions may be truncated due to usage of rest api.***

### Server basics
- [x] Authentification
- [x] Token creation and verification
- [x] User login
- [x] Debug Rcon Connection
- [ ] Swagger
- [ ] Dockerfile

### Minecraft Commands
- [x] [/advancement](https://minecraft.fandom.com/wiki/Commands/advancement)
- [x] [/attribute](https://minecraft.fandom.com/wiki/Commands/attribute)
- [x] [/ban](https://minecraft.fandom.com/wiki/Commands/ban)
- [x] [/ban-ip](https://minecraft.fandom.com/wiki/Commands/ban#ban-ip)
- [x] [/banlist](https://minecraft.fandom.com/wiki/Commands/ban#banlist)
- [x] [/bossbar](https://minecraft.fandom.com/wiki/Commands/bossbar)
- [x] [/clear](https://minecraft.fandom.com/wiki/Commands/clear)
- [x] [/clone](https://minecraft.fandom.com/wiki/Commands/clone)
- [ ] [/data](https://minecraft.fandom.com/wiki/Commands/data)
- [x] [/debug](https://minecraft.fandom.com/wiki/Commands/debug)
- [x] [/defaultgamemode](https://minecraft.fandom.com/wiki/Commands/defaultgamemode)
- [x] [/difficulty](https://minecraft.fandom.com/wiki/Commands/difficulty)
- [x] [/effect](https://minecraft.fandom.com/wiki/Commands/effect)
- [x] [/enchant](https://minecraft.fandom.com/wiki/Commands/enchant)
- [ ] [/execute](https://minecraft.fandom.com/wiki/Commands/execute)
- [x] [/experience](https://minecraft.fandom.com/wiki/Commands/experience)
- [ ] [/fill](https://minecraft.fandom.com/wiki/Commands/fill)
- [ ] [/forceload](https://minecraft.fandom.com/wiki/Commands/forceload)
- [ ] [/function](https://minecraft.fandom.com/wiki/Commands/function)
- [x] [/gamemode](https://minecraft.fandom.com/wiki/Commands/gamemode)
- [x] [/gamerule](https://minecraft.fandom.com/wiki/Commands/gamerule)
- [x] [/give](https://minecraft.fandom.com/wiki/Commands/give)
- [x] [/help](https://minecraft.fandom.com/wiki/Commands/help)
- [x] [/kick](https://minecraft.fandom.com/wiki/Commands/kick)
- [x] [/kill](https://minecraft.fandom.com/wiki/Commands/kill)
- [x] [/list](https://minecraft.fandom.com/wiki/Commands/list)
- [x] [/locate](https://minecraft.fandom.com/wiki/Commands/locate)
- [x] [/locatebiome](https://minecraft.fandom.com/wiki/Commands/locatebiome)
- [ ] [/loot](https://minecraft.fandom.com/wiki/Commands/loot)
- [x] [/me](https://minecraft.fandom.com/wiki/Commands/me)
- [x] [/msg](https://minecraft.fandom.com/wiki/Commands/msg)
- [x] [/op](https://minecraft.fandom.com/wiki/Commands/op)
- [x] [/pardon](https://minecraft.fandom.com/wiki/Commands/pardon)
- [x] [/pardon-ip](https://minecraft.fandom.com/wiki/Commands/pardon#pardon-ip)
- [ ] [/particle](https://minecraft.fandom.com/wiki/Commands/particle)
- [ ] [/playsound](https://minecraft.fandom.com/wiki/Commands/playsound)
- [ ] [/publish](https://minecraft.fandom.com/wiki/Commands/publish)
- [ ] [/recipe](https://minecraft.fandom.com/wiki/Commands/recipe)
- [ ] [/reload](https://minecraft.fandom.com/wiki/Commands/reload)
- [ ] [/save-all](https://minecraft.fandom.com/wiki/Commands/save#save-all)
- [ ] [/save-off](https://minecraft.fandom.com/wiki/Commands/save#save-off)
- [ ] [/save-on](https://minecraft.fandom.com/wiki/Commands/save#save-on)
- [ ] [/say](https://minecraft.fandom.com/wiki/Commands/say)
- [ ] [/schedule](https://minecraft.fandom.com/wiki/Commands/schedule)
- [ ] [/scoreboard](https://minecraft.fandom.com/wiki/Commands/scoreboard)
- [ ] [/seed](https://minecraft.fandom.com/wiki/Commands/seed)
- [ ] [/setblock](https://minecraft.fandom.com/wiki/Commands/setblock)
- [ ] [/setidletimeout](https://minecraft.fandom.com/wiki/Commands/setidletimeout)
- [ ] [/setworldspawn](https://minecraft.fandom.com/wiki/Commands/setworldspawn)
- [ ] [/spawnpoint](https://minecraft.fandom.com/wiki/Commands/spawnpoint)
- [ ] [/spectate](https://minecraft.fandom.com/wiki/Commands/spectate)
- [ ] [/spreadplayers](https://minecraft.fandom.com/wiki/Commands/spreadplayers)
- [ ] [/stop](https://minecraft.fandom.com/wiki/Commands/stop)
- [ ] [/stopsound](https://minecraft.fandom.com/wiki/Commands/stopsound)
- [ ] [/summon](https://minecraft.fandom.com/wiki/Commands/summon)
- [ ] [/tag](https://minecraft.fandom.com/wiki/Commands/tag)
- [ ] [/team](https://minecraft.fandom.com/wiki/Commands/team)
- [ ] [/teammsg](https://minecraft.fandom.com/wiki/Commands/teammsg)
- [ ] [/teleport](https://minecraft.fandom.com/wiki/Commands/teleport)
- [ ] [/tell](https://minecraft.fandom.com/wiki/Commands/tell)
- [ ] [/tellraw](https://minecraft.fandom.com/wiki/Commands/tellraw)
- [ ] [/time](https://minecraft.fandom.com/wiki/Commands/time)
- [ ] [/title](https://minecraft.fandom.com/wiki/Commands/title)
- [ ] [/tm](https://minecraft.fandom.com/wiki/Commands/tm)
- [ ] [/tp](https://minecraft.fandom.com/wiki/Commands/tp)
- [ ] [/trigger](https://minecraft.fandom.com/wiki/Commands/trigger)
- [ ] [/w](https://minecraft.fandom.com/wiki/Commands/w)
- [ ] [/weather](https://minecraft.fandom.com/wiki/Commands/weather)
- [ ] [/whitelist](https://minecraft.fandom.com/wiki/Commands/whitelist)
- [ ] [/worldborder](https://minecraft.fandom.com/wiki/Commands/worldborder)
- [ ] [/xp](https://minecraft.fandom.com/wiki/Commands/xp)
- [ ] [/replaceitem](https://minecraft.fandom.com/wiki/Commands/replaceitem)
