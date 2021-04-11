package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ban <targets> [<reason>] (reason in body with BanUserPost struct)
func ( ws *WebServer )ban( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	var post ReasonPost
	if err := json.NewDecoder( r.Body ).Decode( &post ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	cmd := fmt.Sprintf( "ban %s %s", vars["target"], post.Message )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}