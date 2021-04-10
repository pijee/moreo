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

}
