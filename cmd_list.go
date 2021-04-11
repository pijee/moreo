package main

import (
	"net/http"
)

// list
func ( ws *WebServer )list( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	msg := ws.rconExecute( arm.RconAccess, "list" )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// list uuids
func ( ws *WebServer )listUUIDS( w http.ResponseWriter, r *http.Request ) {
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	msg := ws.rconExecute( arm.RconAccess, "list uuids" )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
