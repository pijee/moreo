package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// debug (start|stop|report)
func ( ws *WebServer )debug( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	actions := []string{ "start", "stop", "report" }
	if !ws.inArray( vars["action"], actions ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "debug %s", vars["action"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
