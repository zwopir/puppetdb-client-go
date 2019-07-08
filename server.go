package puppetdb

import (
	"net/http"
	"time"
)

/*
Server Representation of a PuppetDB server instance.

Use NewServer to create a new instance.
*/
type Server struct {
	client  *http.Client
	BaseUrl string
}

// Opt is an option func
type Opt func(*Server)

/*
NewServer Create a new instance of a Server for usage later.

This is usually the main entry point of this SDK, where you would create
this initial object and use it to perform activities on the instance in
question.
*/
func NewServer(baseUrl string, opts ...Opt) *Server {
	s := &Server{
		client:  http.DefaultClient,
		BaseUrl: baseUrl,
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

// WithRoundTripper sets the roundtrippter for the api http client
func WithRoundTripper(rt http.RoundTripper) Opt {
	return func(s *Server) {
		s.client.Transport = rt
	}
}

// WithTimeout sets the http client timeout
func WithTimeout(d time.Duration) Opt {
	return func(s *Server) {
		s.client.Timeout = d
	}
}
