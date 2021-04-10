package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// defaultgamemode (survival|creative|adventure|spectator)
func ( ws *WebServer )defaultgamemode( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	modes := []string{ "survival", "creative", "adventure", "spectator" }
	if !ws.inArray( vars["mode"], modes ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "defaultgamemode %s", vars["mode"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
