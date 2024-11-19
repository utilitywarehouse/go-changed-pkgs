package sql

import "C"

import (
	"embed"
)

//go:embed *.sql
var fs embed.FS
