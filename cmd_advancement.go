package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// advancement (grant|revoke) <targets> everything
func ( ws *WebServer )advancementEverything( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	actions := []string{ "grant", "revoke" }
	if !ws.inArray( vars["action"], actions ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "advancement %s %s everything", vars["action"], vars["target"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// advancement (grant|revoke) <targets> (only|from|through|until) <advancement>
func ( ws *WebServer )advancementFilter( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	actions := []string{ "grant", "revoke" }
	if !ws.inArray( vars["action"], actions ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	filters := []string{ "only", "from", "through", "until" }
	if !ws.inArray( vars["filter"], filters ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "advancement %s %s %s %s", vars["action"], vars["target"], vars["filter"], vars["advancement"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}