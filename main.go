package main

import (
	"school/nroute"
	_ "github.com/lib/pq"
)


func main() {
	r := nroute.Nroute()
   	r.Run(":1234")
}