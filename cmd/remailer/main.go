package main

import "github.com/Didstopia/remailer/pkg/smtpserver"

func main() {
	if err := smtpserver.Start(); err != nil {
		panic(err)
	}
}
