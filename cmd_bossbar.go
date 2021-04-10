package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// bossbar add <id> <name>
func ( ws *WebServer )bossbarAdd( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar add %s %s", vars["id"], vars["name"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar get <id> (max|players|value|visible)
func ( ws *WebServer )bossbarGet( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	values := []string{ "max", "players", "value", "visible" }
	if !ws.inArray( vars["value"], values ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar get %s %s", vars["id"], vars["value"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar list
func ( ws *WebServer )bossbarList( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	msg := ws.rconExecute( arm.RconAccess, "bossbar list" )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar remove <id>
func ( ws *WebServer )bossbarRemove( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar remove %s", vars["id"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar set <id> color (blue|green|pink|purple|red|white|yellow)
func ( ws *WebServer )bossbarSetColor( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	colors := []string{ "blue", "green", "pink", "purple", "red", "white", "yellow" }
	if !ws.inArray( vars["color"], colors ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar set %s color %s", vars["id"], vars["color"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar set <id> max <max>
func ( ws *WebServer )bossbarSetMax( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar set %s max %s", vars["id"], vars["max"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar set <id> name <name>
func ( ws *WebServer )bossbarSetName( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar set %s name %s", vars["id"], vars["name"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar set <id> players [<targets>]
func ( ws *WebServer )bossbarSetPlayers( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar set %s players %s", vars["id"], vars["player"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar set <id> style (notched_6|notched_10|notched_12|notched_20|progress)
func ( ws *WebServer )bossbarSetStyle( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	styles := []string{ "notched_6", "notched_10", "notched_12", "notched_20", "progress" }
	if !ws.inArray( vars["style"], styles ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar set %s style %s", vars["id"], vars["style"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar set <id> players [<targets>]
func ( ws *WebServer )bossbarSetValue( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar set %s value %s", vars["id"], vars["value"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// bossbar set <id> style (notched_6|notched_10|notched_12|notched_20|progress)
func ( ws *WebServer )bossbarSetVisible( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	visibles := []string{ "true", "false" }
	if !ws.inArray( vars["visible"], visibles ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "bossbar set %s visible %s", vars["id"], vars["visible"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
