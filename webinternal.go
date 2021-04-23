package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/pijee/pjwt"
	"github.com/willroberts/minecraft-client"
)

// Clean service routine
func ( ws *WebServer )cleanService() {
	// Watch end of program
	end := make( chan os.Signal, 1 )
	signal.Notify( end, syscall.SIGINT, syscall.SIGTERM )

	// Schedulde clean tokens database and new key
	tokensdbTicker := time.NewTicker( DB_CLEAN_INTERVAL )
	keyTicker := time.NewTicker( KEY_GENERATION_INTERVAL )
	for {
		select {
			// Stop server
			case <-end:
				tokensdbTicker.Stop()
				keyTicker.Stop()
				ws.Shutdown( context.TODO() )
				return

			// Range over all tokens and check validity. If token is expired, drop it
			case <-tokensdbTicker.C:
				for k,v := range ws.tokensDB {
					if v.Expire < time.Now().Unix() {
						delete( ws.tokensDB, k )
					}
				}

			// Create new key and truncate tokens db
			case <-keyTicker.C:
				// Gen new sign key
				rand.Read( ws.key )

				// Truncate tokens db
				ws.tokensDB = make( map[string]RconAccess )

		}
	}
}

// Function who parse token from header and return AccessReturnMessage struct. Http Code error will be in
// HTTP Code returns :
// - Token invalid, empty or not present in request : 407
// - Token expired or not found in internal db : 403
// - Token parse error : 500
// - Grant Access : 200
func ( ws *WebServer )checkAccess( r *http.Request ) AccessReturnMessage {
	var rcon RconAccess

	// Verify header
	authorizationHeader := r.Header.Get( "Authorization" )
	bearer := strings.Split( authorizationHeader, " " )
	if (len(bearer) != 2) || (bearer[0] != "bearer") {
		return AccessReturnMessage{ http.StatusProxyAuthRequired, rcon }	// HTTP 407
	}

	// Verify token
	pt := new( pjwt.PJWToken )
	if !pt.ValidToken( bearer[1], ws.key ) {
		return AccessReturnMessage{ http.StatusForbidden, rcon } 	// HTTP 403
	}

	// Extract data
	var at AccessToken
	if err := pt.ExtractPlayloadFromToken( bearer[1], &at ); err != nil {
		return AccessReturnMessage{ http.StatusInternalServerError, rcon } 	// HTTP 500
	}

	// Token expiration
	if at.Expire < time.Now().Unix() {
		return AccessReturnMessage{ http.StatusForbidden, rcon }	// HTTP 403
	}

	// Get RconAccess in internal DB
	rcon, ok := ws.tokensDB[ at.UUID ]
	if !ok {
		return AccessReturnMessage{ http.StatusForbidden, rcon }	// HTTP 403
	}

	return AccessReturnMessage{ http.StatusOK, rcon }	// HTTP 200
}

// Connect, execute and get response from rcon server
// HTTP Codes:
// - connection to rcon server failed : 502
// - authentification to rcon server failed : 401
// - command execution error : 400
// - command ok : 200
func ( ws *WebServer )rconExecute( rcon RconAccess, cmd string ) RconReturnMessage {
	// Rcon server connection
	uri := fmt.Sprintf( "%s:%d", rcon.Host, rcon.Port )
	client, err := minecraft.NewClient( uri )
	if err != nil { return RconReturnMessage{ http.StatusBadGateway, "" } }	// HTTP 502
	defer client.Close()

	// Rcon server authentification
	if  err := client.Authenticate( rcon.Password ); err != nil {
		fmt.Println( err )
		return RconReturnMessage{ http.StatusUnauthorized, "" } 	// HTTP 401
	}

	// Send command to rcon server and get response
	resp, err := client.SendCommand( cmd )
	if err != nil { return RconReturnMessage{ http.StatusBadRequest, "" } }		// HTTP 400

	return RconReturnMessage{ http.StatusOK, resp.Body }
}
