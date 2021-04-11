package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// help
func ( ws *WebServer )help( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	msg := ws.rconExecute( arm.RconAccess, "help" )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// help <command>
func ( ws *WebServer )helpCmd( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "help %s", vars["cmd"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}