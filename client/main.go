// sign up and log in

package main

import (
	"buffo"
	"fmt"
	"os"
	"strings"

)

func main() {
	fmt.Println("👋")

	Manager.err := NewAccount()

	if err != nil {
		fmt.Println("fail to initialize account manager:", err)
		os.Exit(1)
		var email, password string
	}
}


