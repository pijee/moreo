package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	TOKEN_LIFE 				= 15 * time.Minute
	DB_CLEAN_INTERVAL		= 5 * time.Minute
	KEY_GENERATION_INTERVAL = 2 * time.Hour
)

// Check environment variables for connections
func init() {
	end := false
	if len( os.Getenv( "MOREO_WEB_PORT" ) ) == 0 { end = true }
	if len( os.Getenv( "MOREO_DB_PATH" ) ) == 0 { end = true }
	if end {
		fmt.Println( "Please set environment variables before !" )
		fmt.Println( "- MOREO_WEB_PORT : \tport for web server to listen on" )
		fmt.Println( "- MOREO_DB_PATH : \tpath to user's db file for authentification service" )
		fmt.Println( "" )
		fmt.Println( "Extra Options : (can be empty or uninitialized)" )
		fmt.Println( "- MOREO_WEB_IP : \tip for webserver to listen on (by default, all ip availables)" )
		fmt.Println( "- MOREO_TOKEN_LIFE : \tnumber of minutes for user's authentification without sending login/password (by default : 15 minutes)" )
		fmt.Println( "- MOREO_DEBUG : \tactive debugs functionnalities (value need to be 1, default 0)" )
		os.Exit( 0 )
	}

	// Support admin's token life setting
	tl, err := strconv.Atoi( os.Getenv( "MOREO_TOKEN_LIFE") )
	if err == nil { TOKEN_LIFE = time.Duration(tl) * time.Minute }
}

// Start of program
func main() {
	ws := NewWebServer()
	if err := ws.ListenAndServe(); err != nil { fmt.Println( err ) }
}
