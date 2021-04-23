package main

import (
	"encoding/json"
	"net/http"
)

// Use by /post route to send user's command to rcon server
func ( ws *WebServer )postCommand( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	var post CommandPost
	if err := json.NewDecoder( r.Body ).Decode( &post ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	msg := ws.rconExecute( arm.RconAccess, post.Cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
