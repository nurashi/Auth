package httpAuth

import "fmt"

func SendVerificationEmail(email string, token string) {
	verificationLink := fmt.Sprintf("http://example.com/verify?token=%s", token)
	fmt.Printf("Sent verification email to %s with link: %s\n", email, verificationLink)
}
