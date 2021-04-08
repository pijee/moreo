package main

import (
	"bufio"
	"os"
	"strings"
)

// Parse db file and return a map with [user]=password
func LoadUsersDatabase( _path string ) ( map[string]string, error ) {
	file, err := os.Open( _path )
	if err != nil { return nil, err }
	defer file.Close()

	usersDB := make( map[string]string )
	scanner := bufio.NewScanner( file )
	for scanner.Scan() {
		inputs := strings.Split( scanner.Text(), ":" )
		if len(inputs) != 2 { continue }

		user, password := inputs[0], inputs[1]
		usersDB[ user ] = password
	}

	if err := scanner.Err(); err != nil { return nil, err }
	return usersDB, nil
}
