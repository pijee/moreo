package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Create rcon connection without internal authentification by token
func ( ws *WebServer )debugRCON(  w http.ResponseWriter, r *http.Request) {
	var cmd DebugPost
	if err := json.NewDecoder( r.Body ).Decode( &cmd ); err != nil {
		fmt.Println( err )
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	result := ws.rconExecute( RconAccess{ Host: cmd.Host, Port: cmd.Port, Password: cmd.Password }, cmd.Cmd )
	fmt.Println( result )
}