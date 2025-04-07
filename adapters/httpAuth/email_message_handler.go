package httpAuth

import "fmt"

func SendVerificationEmail(email string, token string) {
	// Example: send an email with a link to verify the email
	verificationLink := fmt.Sprintf("http://example.com/verify?token=%s", token)
	fmt.Printf("Sent verification email to %s with link: %s\n", email, verificationLink)
}
