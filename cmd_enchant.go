package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// enchant <targets> <enchantment> [<level>]
func ( ws *WebServer )enchant( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "enchant %s %s %s", vars["target"], vars["enchantment"], vars["level"]  )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
