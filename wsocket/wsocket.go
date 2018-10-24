// Package wsocket ...
// Copyright (c) 2018 pirakansa
package wsocket

import (
	"container/list"
	"net/http"

	"golang.org/x/net/websocket"
)

const bufsize = 65536

var conns = list.List{}

func wsHandler(ws *websocket.Conn) {
	con := conns.PushBack(ws)
	defer func() {
		ws.Close()
		conns.Remove(con)
	}()
	for {
		msg := make([]byte, bufsize)
		n, err := ws.Read(msg)
		if err != nil {
			break
		}
		for e := conns.Front(); e != nil; e = e.Next() {
			e.Value.(*websocket.Conn).Write(msg[:n])
		}
	}
}

// Serve websocket
func Serve() {
	http.Handle("/ws", websocket.Handler(wsHandler))
}
