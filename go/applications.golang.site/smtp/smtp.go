package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// func PlainAuth(identity, username, password, host string) Auth
	from := "flavono123@gmail.com"
	to := []string{"flavono123@gmail.com"} // Initialize a string array with an array literal
	// Use [App password](https://security.google.com/settings/security/apppasswords) since I use 2FA for Google account
	// ref. https://gist.github.com/jpillora/cb46d183eca0710d909a?permalink_comment_id=3239532#gistcomment-3239532
	// Eliding the password, this codes are would not working
	pass := "..."
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	headerSubject := "Subject: Test subj\r\n"
	headerBlank := "\r\n"
	body := "Test\r\n"

	msg := []byte(headerSubject + headerBlank + body) // Type Convert: strings to byte slice

	// Also works for 25
	port := 587 // ref. https://sendgrid.com/blog/whats-the-difference-between-ports-465-and-587/
	addr := fmt.Sprintf("smtp.gmail.com:%d", port)
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("vim-go")
}
