package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// difficulty [easy|hard|normal|peaceful]
func ( ws *WebServer )difficulty( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	difficulties := []string{ "easy", "hard", "normal", "peaceful" }
	if !ws.inArray( vars["difficulty"], difficulties ) {
		w.WriteHeader( http.StatusBadRequest )	// HTTP 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "difficulty %s", vars["difficulty"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
