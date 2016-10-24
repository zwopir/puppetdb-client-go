package puppetdb

import "net/http"

/*
Representation of a PuppetDB server instance.

Use NewServer to create a new instance.
*/
type Server struct {
	BaseUrl       string
	HTTPTransport http.RoundTripper
}

/*
Create a new instance of a Server for usage later.

This is usually the main entry point of this SDK, where you would create
this initial object and use it to perform activities on the instance in
question.
*/
func NewServer(baseUrl string) Server {
	return newServer(baseUrl, nil)
}

/*
Create a new instance of a Server for usage later.

Comparable to NewServer, but with an additional parameter to specify the http transport
(i.e. SSL options)
*/
func NewServerWithTransport(baseUrl string, httpTransport http.RoundTripper) Server {
	return newServer(baseUrl, httpTransport)
}

func newServer(baseUrl string, httpTransport http.RoundTripper) Server {
	return Server{
		BaseUrl:       baseUrl,
		HTTPTransport: httpTransport,
	}
}
