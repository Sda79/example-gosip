package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgo/sip"
)

func main() {
	ua, _ := sipgo.NewUA()        // Build user agent
	srv, _ := sipgo.NewServer(ua) // Creating server handle for ua

	var testHandler = func(req *sip.Request, tx sip.ServerTransaction) {
		res := sip.NewResponseFromRequest(req, 200, "OK", nil)
		srv.WriteResponse(res)
	}

	srv.OnInvite(testHandler)

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	go srv.ListenAndServe(ctx, "udp", "127.0.0.1:5060")
	go srv.ListenAndServe(ctx, "tcp", "127.0.0.1:5061")
	go srv.ListenAndServe(ctx, "ws", "127.0.0.1:5080")
	<-ctx.Done()
}
