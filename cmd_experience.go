package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// experience (add|set) <targets> <amount> [levels|points]
// xp (add|set) <targets> <amount> [levels|points]
func ( ws *WebServer )experienceModify( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	actions := []string{ "add", "set" }
	if !ws.inArray( vars["action"], actions ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	levels := []string{ "levels", "points" }
	if !ws.inArray( vars["level"], levels ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "experience %s %s %s %s", vars["action"], vars["target"], vars["amount"], vars["level"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// experience query <targets> (levels|points)
// xp query <targets> (levels|points)
func ( ws *WebServer )experienceQuery( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	levels := []string{ "levels", "points" }
	if !ws.inArray( vars["level"], levels ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "experience query %s %s", vars["target"], vars["level"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}