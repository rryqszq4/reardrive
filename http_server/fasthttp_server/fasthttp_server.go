package fasthttp_server

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"reardrive/core/modules"
	"net"
	"time"
)

type FastHttpServer struct {
	modules.HttpServerModule

	Server *fasthttp.Server

	Listener net.Listener

	GracefulListener net.Listener

}

func (self *FastHttpServer) Run_before_module() {
	self.initHttpServer()
}

func (self *FastHttpServer) Close_module() {

	err := self.GracefulListener.Close()
	if err != nil {
		self.Logger.Errorf("error with graceful close: %s", err)
	}

	self.Logger.Infof("Server gracefully stopped")

	self.Stop()
}

func (self *FastHttpServer) initHttpServer() {

	var err error
	self.Listener, err = reuseport.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		self.Logger.Errorf("error in reuseport listener: %s", err)
		return
	}

	duration := 5 * time.Second
	self.GracefulListener = NewGracefulListener(self.Listener, duration)

	self.Server = &fasthttp.Server{
		Handler: 				self.requestHandler,
		Name: 					"My super server",
		ReadTimeout:          	5 * time.Second,
		WriteTimeout:         	10 * time.Second,
		MaxConnsPerIP:        	500,
		MaxRequestsPerConn:   	500,
		MaxKeepaliveDuration: 	5 * time.Second,
	}

	self.Hooker.Listener = self.listener
}

func (self *FastHttpServer) listener() {


	self.Logger.Info("Http server serving...")
	if err := self.Server.Serve(self.Listener); err != nil {
		self.Logger.Errorf("Error in Serve: %s", err)
	}

}

func (self *FastHttpServer) requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}