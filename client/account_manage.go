package main

import (
	"fmt"
	"net/smtp"
	"crypto/rand"
	"encoding/hex"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Manager struct {
	db *sql.DB // pointer, avoid copying
}

func NewAccount() (*Manager, error){
	// DONE
	// create table
	// poor error handle
	// return value in the form of data pairs
	// TODO
	// store information of user, email account, account
	// add  more functions here to handle the exception and additional function( like a part of secure verify )
	// main point is to create a space to stash user-info
	db, err := sql.Open("sqlite3", "./chatroom.db") // already uni-test
	if err != nil {
		return nil, err
	}

	_, err = db.Exec('

		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE,
			password TEXT,
			token TEXT
		)
	')

	if err ! = nil {
		return nil, err
	}

	return &Manager{db:db}, nil
} 

// generate
func (am *Manager) Register(email, password string) error {
	// generate token to verify wether it is user 
	// 16 byte token
	// 
	token := make([byte], 16)
	_, err := rand.Read(token)
	if err != nil {
		return err
	}

	tokenStr := hex.EncodeToString(token)
	
	_. err = am.db.Exec("INSERT INTO users (email, password, token) VALUES (?, ?, ?), email, password, tokenStr")
	if err != nil{
		return err
	}

	err = sendVerificationEmail(email, tokenStr)
	if err != nil {
		return err
	}

	return nil
} 
// verify part 
func (am *Manager) Login(email, password string) (string, error) {
	var storedPassword, token string
	err := am.db.QueryRow("SELECT password, token FROM users WHERE email = ?", email).Scan(&storedPassword, &token)
	if err != nil {
		return "", err
	}

	if storedPassword != password {
		return "", fmt.Errorf("invalid password")
	}

	return token, nil
}
// send
func sendVerificationEmail(email, token string) error {
	// TEST
	// HACK: Maybe verify it by using local server without a domain
	from := "email@email.org"
	password := "nopasswd"
	to := []string{email}
	smtpHost := "smtp.email.org"
	smtpPort := "587"

	message := []byte("Subject: ChatRoom Verification\r\b\r\n"+ "Please click the following link to verify your email: https://email.com/verify?token" + token)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth,from,to,message)
	if err != nil {
		return err
	}
	return nil

}
