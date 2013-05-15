// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package main

import (
	"net/http"
	"log"
	"github.com/darkhelmet/env"
)


func main() {
	port := env.StringDefault("PORT", "9000")
	pwd := env.StringDefault("PWD", "/app")
	//
	// File Server
	//
	http.Handle("/", http.FileServer(http.Dir(pwd+"/public")))
	//
	// Start Webserver
	//
	log.Println("Starting webserver on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
