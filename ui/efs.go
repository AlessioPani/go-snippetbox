package ui

import "embed"

// Special comment directive. When our application is compiled (as part of either go build or go run), 
// the comment //go:embed "static" instructs Go to store the files from our ui/static folder in 
// an embedded filesystem referenced by the global variable Files.

//go:embed "static"
var Files embed.FS
