package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// gamerule <rule name> [<value>]
// https://minecraft.fandom.com/wiki/Game_rule
func ( ws *WebServer )gamerule( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "gamerule %s %s", vars["rule"], vars["value"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
