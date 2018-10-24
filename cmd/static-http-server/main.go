// Copyright (c) 2018 pirakansa
package main

import (
	"fmt"
	"net/http"
	"flag"

	"github.com/pirakansa/go-static-http-server/spage"
	"github.com/pirakansa/go-static-http-server/wsocket"
)

// Args ...
type Args struct {
	Port int
}

func getArgs() Args {
	args := new(Args)

	flag.IntVar(&args.Port, "port", 8080, "port number")
	flag.IntVar(&args.Port, "p", 8080, "(short) --port")
	flag.Parse()

	return *args
}

func main() {
	args := getArgs()
	wsocket.Serve()
	spage.Serve()

	http.ListenAndServe(fmt.Sprintf(":%d", args.Port), nil)
}
