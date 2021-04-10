package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// effect clear [<target>] [<effect>]
func ( ws *WebServer )effectClear( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "effect clear %s %s", vars["target"], vars["effect"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// effect give <targets> <effect> [<seconds:int>] [<amplifier:int>] [<hideParticles:boolean>]
func ( ws *WebServer )effectGive( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	booleans := []string{ "true", "false" }
	if !ws.inArray( vars["hideparticles"], booleans ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "effect give %s %s %s %s %s", vars["target"], vars["effect"], vars["seconds"], vars["amplifier"] , vars["hideparticles"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
