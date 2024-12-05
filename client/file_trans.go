package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type FileTransfer struct {
	Conn net.Conn
}

func NewFileTransfer(conn net.Conn) *FileTransfer {
	return &FileTransfer{Conn: conn}
}

func (ft *FileTransfer) SendFile(filePath string) error {}

func (ft *FileTransfer) Receive() error {}

func (ft *FileTransfer) transfer() error {}
