package main

import (
	"net/http"
	"time"
)

type Client struct {
	client http.Client
	opts   ClientOpts
}

type ClientOpts struct {
	BaseUrl string
	Auth    *ClientAuth
}

type ClientAuth struct {
	Token string
}

func New(opts *ClientOpts) *Client {
	c := http.Client{Timeout: time.Duration(1) * time.Minute}

	return &Client{
		client: c,
		opts:   *opts,
	}
}
