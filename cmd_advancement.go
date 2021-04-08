package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// advancement (grant|revoke) <targets> everything
func ( ws *WebServer )advancementEverything( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	if (vars["action"] != "grant") && (vars["action"] != "revoke") {
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

// advancement (grant|revoke) <targets> only <advancement>
// advancement (grant|revoke) <targets> from <advancement>
// advancement (grant|revoke) <targets> through <advancement>
// advancement (grant|revoke) <targets> until <advancement>
func ( ws *WebServer )advancementFilter( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	if (vars["action"] != "grant") && (vars["action"] != "revoke") {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	if (vars["filter"] != "only") && (vars["filter"] != "from") && (vars["filter"] != "through") && (vars["filter"] != "until") {
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