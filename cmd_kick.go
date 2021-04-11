package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// kick <target> [<reason>]
func ( ws *WebServer )kick( w http.ResponseWriter, r *http.Request ) {
	var post KickPost
	if err := json.NewDecoder( r.Body ).Decode( &post ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "kick %s %s", post.Name, post.Reason )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}