package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "CHATROOM: ", log.LstdFlags|log.Lshortfile)
}
