package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Struct to parse Post request on /login
type AuthentificationPost struct {
	User struct {
		Name 		string 	`json:"name"`
		Password 	string 	`json:"password"`
	} `json:"user"`

	Rcon struct {
		Host 		string 	`json:"host"`
		Port 		uint16 	`json:"port"`
		Password 	string 	`json:"password"`
	} `json:"rcon"`
}

// Message from authentification procedure
type AccessReturnMessage struct {
	HTTPCode 	int
	RconAccess 	RconAccess
}

// Struct for rcon server access in WebServer.tokensDB
type RconAccess struct {
	Host 		string 	`json:"host"`
	Port 		uint16 	`json:"port"`
	Password 	string 	`json:"password"`
	Expire		int64 	`json:"exp"`		// will be trigger by cleanservice
}

// Struct of token from client to api access
type AccessToken struct {
	UUID 	string 	`json:"uuid"`
	Expire	int64 	`json:"exp"`
}

// WebServer Struct
type WebServer struct {
	http.Server
	key 		[]byte
    router 		*mux.Router
	usersDB 	map[string]string
	tokensDB 	map[string]RconAccess
}

// Message returning from rcon execute commande
type RconReturnMessage struct {
	Code int
	Body string
}

// Struct for /ban command
type BanUserPost struct {
	Reason 	string  `json:"reason"`
}

// Struct for /banip command
type BanIPPost struct {
	Target 	string 	`json:"target"`
	Reason 	string  `json:"reason"`
}

// For receiving cmd in debug url
type DebugPost struct {
	Host 		string 	`json:"host"`
	Port 		uint16 	`json:"port"`
	Password 	string 	`json:"password"`
	Cmd 		string 	`json:"cmd"`
}