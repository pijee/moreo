package main

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/pijee/pjwt"
)

// POST : perform user's authentification and send token
func ( ws *WebServer )login( w http.ResponseWriter, r *http.Request ) {
	// Fill struct with json post
	var post AuthentificationPost
	if err := json.NewDecoder( r.Body ).Decode( &post ); err != nil {
		w.WriteHeader( http.StatusBadRequest )	// 400
		return
	}

	// Verify identity
	password, ok := ws.usersDB[ post.User.Name ]
	if !ok || post.User.Password != password {
		w.WriteHeader( http.StatusForbidden )	// 403
		return
	}

	// UUID Generation
	id, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader( http.StatusInternalServerError )	// 500
		return
	}

	// Token creation
	t := time.Now().Add( TOKEN_LIFE ).Unix()
	accessTk := AccessToken{ UUID: id.String(), Expire: t }
	dbTk := RconAccess{ Host: post.Rcon.Host, Port: post.Rcon.Port, Password: post.Rcon.Password, Expire: t }
	pt := new( pjwt.PJWToken )
	str, err := pt.CreateToken( accessTk, ws.key )
	if err != nil {
		w.WriteHeader( http.StatusInternalServerError )	// 500
		return
	}

	// Send token and store it into internal db
	w.Write( []byte(str) )
	ws.tokensDB[ id.String() ] = dbTk
}
