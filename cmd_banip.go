package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ban-ip <target> [<reason>] (target and reason on BanIPPost struct)
func ( ws *WebServer )banip( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	var ban BanIPPost
	if err := json.NewDecoder( r.Body ).Decode( &ban ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	cmd := fmt.Sprintf( "ban-ip %s %s", ban.Target, ban.Reason )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}