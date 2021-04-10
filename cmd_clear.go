package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// clear [<target>] [<item>] [<maxCount>]
func ( ws *WebServer )clear( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "clear %s %s %s", vars["target"], vars["item"], vars["maxcount"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
