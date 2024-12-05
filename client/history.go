package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type History struct {
	db *sql.DB
}

func NewHistoryFile(*History, error){}

func (hm *History) Save() error{}

func (hm *History) Get() ([]Message, error) {}

func (hm *History) Cls() ([]Message, error) {}


