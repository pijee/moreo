package main

import "os"

func ( ws *WebServer )setAllRoutes() {
	// User proxy login
	ws.router.HandleFunc( "/login", ws.login ).Methods( "POST" )

	// For debugging
	if os.Getenv( "MOREO_DEBUG" ) == "1" {
		ws.router.HandleFunc( "/debug/rcon/post", ws.debugRCON )	// create & execute rcon command
	}

	//
	// https://minecraft.fandom.com/wiki/Commands
	//

	// https://minecraft.fandom.com/wiki/Commands/advancement
	ws.router.HandleFunc( "/advancement/{action}/{target}/everything", ws.advancementEverything ).Methods( "PATCH" )
	ws.router.HandleFunc( "/advancement/{action}/{target}/{filter}/{advancement}", ws.advancementFilter ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/attribute
	ws.router.HandleFunc( "/attribute/{target}/{attribute}/get/{scale}", ws.attributeGet ).Methods( "GET" )
	ws.router.HandleFunc( "/attribute/{target}/{attribute}/base/get/{scale}", ws.attributeBaseGet ).Methods( "GET" )
	ws.router.HandleFunc( "/attribute/{target}/{attribute}/set/{value}", ws.attributeSet ).Methods( "PATCH" )
	ws.router.HandleFunc( "/attribute/{target}/{attribute}/modifier/add/{uuid}/{name}/{value}/{operation}", ws.attributeModifierAdd ).Methods( "POST" )
	ws.router.HandleFunc( "/attribute/{target}/{attribute}/modifier/remove/{uuid}", ws.attributeModifierRemove ).Methods( "DELETE" )
	ws.router.HandleFunc( "/attribute/{target}/{attribute}/modifier/value/get/{uuid}/{scale}", ws.attributeModifierGet ).Methods( "GET" )

	// https://minecraft.fandom.com/wiki/Commands/ban (reason in BanUserPost struct)
	ws.router.HandleFunc( "/ban/{target}", ws.ban ).Methods( "POST" )

	// https://minecraft.fandom.com/wiki/Commands/ban#ban-ip (target and reason in BanIPPost struct)
	ws.router.HandleFunc( "/banip", ws.banip ).Methods( "POST" )

	// https://minecraft.fandom.com/wiki/Commands/ban#banlist
	ws.router.HandleFunc( "/banlist/", ws.banlist ).Methods( "GET" )
	ws.router.HandleFunc( "/banlist/{cat}", ws.banlistFiltered ).Methods( "GET" )

	// https://minecraft.fandom.com/wiki/Commands/bossbar
	ws.router.HandleFunc( "/bossbar/add/{id}/{name}", ws.bossbarAdd ).Methods( "POST" )
	ws.router.HandleFunc( "/bossbar/get/{id}/{value}", ws.bossbarGet ).Methods( "GET" )
	ws.router.HandleFunc( "/bossbar/list", ws.bossbarList ).Methods( "GET" )
	ws.router.HandleFunc( "/bossbar/remove/{id}", ws.bossbarRemove ).Methods( "DELETE" )
	ws.router.HandleFunc( "/bossbar/set/{id}/color/{color}", ws.bossbarSetColor ).Methods( "PATCH" )
	ws.router.HandleFunc( "/bossbar/set/{id}/max/{max}", ws.bossbarSetMax ).Methods( "PATCH" )
	ws.router.HandleFunc( "/bossbar/set/{id}/name/{name}", ws.bossbarSetName ).Methods( "PATCH" )
	ws.router.HandleFunc( "/bossbar/set/{id}/players/{player}", ws.bossbarSetPlayers ).Methods( "PATCH" )
	ws.router.HandleFunc( "/bossbar/set/{id}/style/{style}", ws.bossbarSetStyle ).Methods( "PATCH" )
	ws.router.HandleFunc( "/bossbar/set/{id}/value/{value}", ws.bossbarSetValue ).Methods( "PATCH" )
	ws.router.HandleFunc( "/bossbar/set/{id}/visible/{visible}", ws.bossbarSetVisible ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/clear
	ws.router.HandleFunc( "/bossbar/clear/{target}/{item}/{maxcount}", ws.clear ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/clone
	ws.router.HandleFunc( "/clone/{action}/{mode}", ws.clone ).Methods( "POST" )
	ws.router.HandleFunc( "/clone/filtered/{filter}/{mode}", ws.cloneFitlered ).Methods( "POST" )

	// https://minecraft.fandom.com/wiki/Commands/debug
	ws.router.HandleFunc( "/debug/{action}", ws.debug ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/defaultgamemode
	ws.router.HandleFunc( "/defaultgamemode/{mode}", ws.defaultgamemode ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/difficulty
	ws.router.HandleFunc( "/difficulty/{difficulty}", ws.difficulty ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/effect
	ws.router.HandleFunc( "/effect/clear/{effect}", ws.effectClear ).Methods( "PATCH" )
	ws.router.HandleFunc( "/effect/give/{effect}/{seconds}/{amplifier}/{hideparticles}", ws.effectGive ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/enchant
	ws.router.HandleFunc( "/enchant/{target}/{enchantment}/{level}", ws.enchant ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/experience
	ws.router.HandleFunc( "/experience/{action}/{target}/{enchantment}/{level}", ws.experienceModify ).Methods( "PATCH" )
	ws.router.HandleFunc( "/experience/query/{target}/{level}", ws.experienceQuery ).Methods( "GET" )

	// https://minecraft.fandom.com/wiki/Commands/gamemode
	ws.router.HandleFunc( "/gamemode/{mode}/{target}", ws.gamemode ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/gamerule
	ws.router.HandleFunc( "/gamerule/{rule}/{value}", ws.gamerule ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/give
	ws.router.HandleFunc( "/give/{target}/{item}/{count}", ws.give ).Methods( "PATCH" )

	// https://minecraft.fandom.com/wiki/Commands/help
	ws.router.HandleFunc( "/help", ws.help ).Methods( "GET" )
	ws.router.HandleFunc( "/help/{cmd}", ws.helpCmd ).Methods( "GET" )

	// https://minecraft.fandom.com/wiki/Commands/kick
	ws.router.HandleFunc( "/kick", ws.kick ).Methods( "POST" )

	// https://minecraft.fandom.com/wiki/Commands/kill
	ws.router.HandleFunc( "/kill/{target}", ws.kill ).Methods( "DELETE" )

	// https://minecraft.fandom.com/wiki/Commands/list
	ws.router.HandleFunc( "/list", ws.list ).Methods( "GET" )
	ws.router.HandleFunc( "/list/uuids", ws.listUUIDS ).Methods( "GET" )

	// https://minecraft.fandom.com/wiki/Commands/locate
	ws.router.HandleFunc( "/locate/{type}", ws.locate ).Methods( "GET" )

	// https://minecraft.fandom.com/wiki/Commands/locatebiome
	ws.router.HandleFunc( "/locatebiome/{id}", ws.locatebiome ).Methods( "GET" )

	// https://minecraft.fandom.com/wiki/Commands/me
	ws.router.HandleFunc( "/me", ws.me ).Methods( "POST" )

	// https://minecraft.fandom.com/wiki/Commands/msg
	ws.router.HandleFunc( "/msg", ws.msg ).Methods( "POST" )
}
