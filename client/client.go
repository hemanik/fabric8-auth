package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"net/http"
)

// Client is the alm service client.
type Client struct {
	*goaclient.Client
	JWTSigner goaclient.Signer
	Encoder   *goa.HTTPEncoder
	Decoder   *goa.HTTPDecoder
}

// New instantiates the client.
func New(c *http.Client) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetJWTSigner sets the request signer for the jwt security scheme.
func (c *Client) SetJWTSigner(signer goaclient.Signer) {
	c.JWTSigner = signer
}