# RestFull API server for Minecraft (JAVA Edition) Rcon Server writen in Go

Moreo provide a web proxy server than interact with your minecraft server's rcon service for you by an RestFull API.

# How it works ?
Moreo authentificate users by small database file (explication bellow).
Once you're connected, Moreo send you token to interact with API (like a normal RestFUll API).
Each token spread **one rcon server**. That's mean you can interact with many differents servers with the same account in same times, with differents tokens.
Moreo execute command on rcon server for you and send you result.

For managing connection, Moreo use [Will Robert's Minecraft-Client](github.com/willroberts/minecraft-client)

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

# List of commands
I used [minecraft wiki commands](https://minecraft.fandom.com/wiki/Commands) to implement commands supported by the API.

# Work in progress
- [x] Authentification
- [x] Token creation and verification
- [x] User login
- [x] Debug Rcon Connection
- [ ] Swagger
- [ ] Dockerfile
- [x] [/advancement](https://minecraft.fandom.com/wiki/Commands/advancement)
- [x] [/attribute](https://minecraft.fandom.com/wiki/Commands/attribute)
- [x] [/ban](https://minecraft.fandom.com/wiki/Commands/ban)
- [x] [/ban-ip](https://minecraft.fandom.com/wiki/Commands/ban#ban-ip)
- [x] [/banlist](https://minecraft.fandom.com/wiki/Commands/ban#banlist)
- [ ] [/bossbar](https://minecraft.fandom.com/wiki/Commands/bossbar)
- [ ] [/clear](https://minecraft.fandom.com/wiki/Commands/clear)
- [ ] [/clone](https://minecraft.fandom.com/wiki/Commands/clone)
- [ ] [/data](https://minecraft.fandom.com/wiki/Commands/data)
- [ ] [/debug](https://minecraft.fandom.com/wiki/Commands/debug)
- [ ] [/defaultgamemode](https://minecraft.fandom.com/wiki/Commands/defaultgamemode)
- [ ] [/difficulty](https://minecraft.fandom.com/wiki/Commands/difficulty)
- [ ] [/effect](https://minecraft.fandom.com/wiki/Commands/effect)
- [ ] [/enchant](https://minecraft.fandom.com/wiki/Commands/enchant)
- [ ] [/execute](https://minecraft.fandom.com/wiki/Commands/execute)
- [ ] [/experience](https://minecraft.fandom.com/wiki/Commands/experience)
- [ ] [/fill](https://minecraft.fandom.com/wiki/Commands/fill)
- [ ] [/forceload](https://minecraft.fandom.com/wiki/Commands/forceload)
- [ ] [/function](https://minecraft.fandom.com/wiki/Commands/function)
- [ ] [/gamemode](https://minecraft.fandom.com/wiki/Commands/gamemode)
- [ ] [/gamerule](https://minecraft.fandom.com/wiki/Commands/gamerule)
- [ ] [/give](https://minecraft.fandom.com/wiki/Commands/give)
- [ ] [/help](https://minecraft.fandom.com/wiki/Commands/help)
- [ ] [/kick](https://minecraft.fandom.com/wiki/Commands/kick)
- [ ] [/kill](https://minecraft.fandom.com/wiki/Commands/kick/kill)
- [ ] [/list](https://minecraft.fandom.com/wiki/Commands/kick/list)
- [ ] [/locate](https://minecraft.fandom.com/wiki/Commands/kick/locate)
- [ ] [/locatebiome](https://minecraft.fandom.com/wiki/Commands/kick/locatebiome)
- [ ] [/loot](https://minecraft.fandom.com/wiki/Commands/kick/loot)
- [ ] [/me](https://minecraft.fandom.com/wiki/Commands/kick/me)
- [ ] [/msg](https://minecraft.fandom.com/wiki/Commands/kick/msg)
- [ ] [/op](https://minecraft.fandom.com/wiki/Commands/kick/op)
- [ ] [/pardon](https://minecraft.fandom.com/wiki/Commands/kick/pardon)
- [ ] [/pardon-ip](https://minecraft.fandom.com/wiki/Commands/kick/pardon#pardon-ip)
- [ ] [/particle](https://minecraft.fandom.com/wiki/Commands/kick/particle)
- [ ] [/playsound](https://minecraft.fandom.com/wiki/Commands/kick/playsound)
- [ ] [/publish](https://minecraft.fandom.com/wiki/Commands/kick/publish)
- [ ] [/recipe](https://minecraft.fandom.com/wiki/Commands/kick/recipe)
- [ ] [/reload](https://minecraft.fandom.com/wiki/Commands/kick/reload)
- [ ] [/save-all](https://minecraft.fandom.com/wiki/Commands/kick/save#save-all)
- [ ] [/save-off](https://minecraft.fandom.com/wiki/Commands/kick/save#save-off)
- [ ] [/save-on](https://minecraft.fandom.com/wiki/Commands/kick/save#save-on)
- [ ] [/say](https://minecraft.fandom.com/wiki/Commands/kick/say)
- [ ] [/schedule](https://minecraft.fandom.com/wiki/Commands/kick/schedule)
- [ ] [/scoreboard](https://minecraft.fandom.com/wiki/Commands/kick/scoreboard)
- [ ] [/seed](https://minecraft.fandom.com/wiki/Commands/kick/seed)
- [ ] [/setblock](https://minecraft.fandom.com/wiki/Commands/kick/setblock)
- [ ] [/setidletimeout](https://minecraft.fandom.com/wiki/Commands/kick/setidletimeout)
- [ ] [/setworldspawn](https://minecraft.fandom.com/wiki/Commands/kick/setworldspawn)
- [ ] [/spawnpoint](https://minecraft.fandom.com/wiki/Commands/kick/spawnpoint)
- [ ] [/spectate](https://minecraft.fandom.com/wiki/Commands/kick/spectate)
- [ ] [/spreadplayers](https://minecraft.fandom.com/wiki/Commands/kick/spreadplayers)
- [ ] [/stop](https://minecraft.fandom.com/wiki/Commands/kick/stop)
- [ ] [/stopsound](https://minecraft.fandom.com/wiki/Commands/kick/stopsound)
- [ ] [/summon](https://minecraft.fandom.com/wiki/Commands/kick/summon)
- [ ] [/tag](https://minecraft.fandom.com/wiki/Commands/kick/tag)
- [ ] [/team](https://minecraft.fandom.com/wiki/Commands/kick/team)
- [ ] [/teammsg](https://minecraft.fandom.com/wiki/Commands/kick/teammsg)
- [ ] [/teleport](https://minecraft.fandom.com/wiki/Commands/kick/teleport)
- [ ] [/tell](https://minecraft.fandom.com/wiki/Commands/kick/tell)
- [ ] [/tellraw](https://minecraft.fandom.com/wiki/Commands/kick/tellraw)
- [ ] [/time](https://minecraft.fandom.com/wiki/Commands/kick/time)
- [ ] [/title](https://minecraft.fandom.com/wiki/Commands/kick/title)
- [ ] [/tm](https://minecraft.fandom.com/wiki/Commands/kick/tm)
- [ ] [/tp](https://minecraft.fandom.com/wiki/Commands/kick/tp)
- [ ] [/trigger](https://minecraft.fandom.com/wiki/Commands/kick/trigger)
- [ ] [/w](https://minecraft.fandom.com/wiki/Commands/kick/w)
- [ ] [/weather](https://minecraft.fandom.com/wiki/Commands/kick/weather)
- [ ] [/whitelist](https://minecraft.fandom.com/wiki/Commands/kick/whitelist)
- [ ] [/worldborder](https://minecraft.fandom.com/wiki/Commands/kick/worldborder)
- [ ] [/xp](https://minecraft.fandom.com/wiki/Commands/kick/xp)
- [ ] [/replaceitem](https://minecraft.fandom.com/wiki/Commands/kick/replaceitem)
