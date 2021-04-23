package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Create New WebServer
func NewWebServer() *WebServer {
	ws := &WebServer {
		Server: http.Server {
			Addr: fmt.Sprintf( "%s:%s", os.Getenv( "MOREO_WEB_IP" ), os.Getenv( "MOREO_WEB_PORT" ) ),
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		key: make( []byte, 32 ),
		router: mux.NewRouter(),
		tokensDB: make( map[string]RconAccess ),
	}

	// Init key for signing tokens
	// rand.Read( ws.key )
	ws.key = []byte{ 1, 2, 3, 4, 5, 6, 7, 8, 9 }

	// Routine for cleaning tokens database and watching sigterm signals
	go ws.cleanService()

	// Init users database
	var err error
	ws.usersDB, err = LoadUsersDatabase( os.Getenv( "MOREO_DB_PATH" ) )
	if err != nil {
		log.Fatal( "Unable to load users database" )
	}

	//
	// All routes
	//

	// User proxy login
	ws.router.HandleFunc( "/login", ws.login ).Methods( "POST" )

	// For debugging
	if os.Getenv( "MOREO_DEBUG" ) == "1" {
		ws.router.HandleFunc( "/debug/rcon/post", ws.debugRCON ).Methods( "POST" )	// create & execute rcon command
	}

	// Send command to rcon server after being authenticated
	ws.router.HandleFunc( "/post", ws.postCommand ).Methods( "POST" )

	ws.Handler = ws.router
	return ws
}
