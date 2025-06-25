package helpers

import (
	"fmt"
	"math/big"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(userEmail string, code *big.Int) error {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("MAIL_ADDRESS"))
	m.SetHeader("To", userEmail);
	m.SetHeader("Subject", "Account Verification Code");
	html := fmt.Sprintf(`
        <h2>Your Verification Code</h2>
        <p>Please use the following 6-digit code to verify your email address:</p>
        <h1 style="color:blue;">%s</h1>
        <p>If you did not request this code, please ignore this email.</p>
    `, code)

	m.SetBody("text/html", html)
	d := gomail.NewDialer("smtp.gmail.com", 587 , os.Getenv("MAIL_ADDRESS"), os.Getenv("MAIL_PASS"))

	err := d.DialAndSend(m)

	return err


}