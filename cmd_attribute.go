package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// attribute <target> <attribute> get [<scale>]
func ( ws *WebServer )attributeGet( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "attribute %s %s get %s", vars["target"], vars["attribute"], vars["scale"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// attribute <target> <attribute> base get [<scale>]
func ( ws *WebServer )attributeBaseGet( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "attribute %s %s base get %s", vars["target"], vars["attribute"], vars["scale"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// attribute <target> <attribute> base set <value>
func ( ws *WebServer )attributeSet( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "attribute %s %s set %s", vars["target"], vars["attribute"], vars["value"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// attribute <target> <attribute> modifier add <uuid> <name> <value> (add|multiply|multiply_base)
func ( ws *WebServer )attributeModifierAdd( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	operations := []string{ "add", "multiply", "multiply_base" }
	if !ws.inArray( vars["operation"], operations ) {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	cmd := fmt.Sprintf( "attribute %s %s modifier add %s %s %s %s", vars["target"], vars["attribute"], vars["uuid"], vars["name"], vars["value"], vars["operation"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// attribute <target> <attribute> modifier remove <uuid>
func ( ws *WebServer )attributeModifierRemove( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "attribute %s %s modifier remove %s", vars["target"], vars["attribute"], vars["uuid"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}

// attribute <target> <attribute> modifier value get <uuid> [<scale>]
func ( ws *WebServer )attributeModifierGet( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	arm := ws.checkAccess( r )
	if arm.HTTPCode != 200 {
		w.WriteHeader( arm.HTTPCode )
		return
	}

	cmd := fmt.Sprintf( "attribute %s %s modifier value get %s %s", vars["target"], vars["attribute"], vars["uuid"], vars["scale"] )
	msg := ws.rconExecute( arm.RconAccess, cmd )
	if msg.Code != 200 {
		w.WriteHeader( msg.Code )
		return
	}

	w.Write( []byte(msg.Body) )
}