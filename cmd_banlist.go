package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// banlist
func ( ws *WebServer )banlist( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	msg := ws.rconExecute( arm.RconAccess, "banlist" )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// banlist [ips|players]
func ( ws *WebServer )banlistFiltered( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	if (vars["cat"] != "ips") && (vars["cat"] != "players") {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	cmd := fmt.Sprintf( "banlist %s", vars["cat"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}