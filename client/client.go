package main

import (
	"fmt"
	"net"
)


//HACK: assume that below is all I need
type Client struct (
	ServerAddr string
	ServerPort int
	Conn       net.Conn
)

// trans args,usage are similar to NewAccount() in  account_manage.go
func NewClient() *Client{}

func (c *Client) Connect() error {}

func (c *Client) Login() error{}

// HACK: I don't know is it suitable here, may be it should at main.go?
func (c *Client) MainLoop() {}
