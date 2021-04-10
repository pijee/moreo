package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// clone <begin> <end> <destination> [replace|masked] [force|move|normal]
func ( ws *WebServer )clone( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	actions := []string{ "replace", "masked" }
	if !ws.inArray( vars["action"], actions ) {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	modes := []string{ "force", "move", "normal" }
	if !ws.inArray( vars["mode"], modes ) {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	var post ClonePost
	if err := json.NewDecoder( r.Body ).Decode( &post ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "clone %d %d %d %d %d %d %d %d %d %s %s",
		post.Begin.X, post.Begin.Y, post.Begin.Z,
		post.Destination.X, post.Destination.Y, post.Destination.Z,
		post.End.X, post.End.Y, post.End.Z,
		vars["action"], vars["mode"],
	)
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// clone <begin> <end> <destination> filtered <filter> [force|move|normal]
func ( ws *WebServer )cloneFitlered( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	modes := []string{ "force", "move", "normal" }
	if !ws.inArray( vars["mode"], modes ) {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	var post ClonePost
	if err := json.NewDecoder( r.Body ).Decode( &post ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "clone %d %d %d %d %d %d %d %d %d filtered %s %s",
		post.Begin.X, post.Begin.Y, post.Begin.Z,
		post.Destination.X, post.Destination.Y, post.Destination.Z,
		post.End.X, post.End.Y, post.End.Z,
		vars["filter"], vars["mode"],
	)
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}
