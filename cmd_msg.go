package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// tell <targets> <message>
// msg <targets> <message>
// w <targets> <message>
func ( ws *WebServer )msg( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	var post MsgPost
	if err := json.NewDecoder( r.Body ).Decode( &post ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	cmd := fmt.Sprintf( "msg %s %s", post.Target, post.Message )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}