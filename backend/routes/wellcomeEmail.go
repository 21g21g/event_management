package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func init() {
	err := godotenv.Load()
	if err != nil {

		log.Fatalf("Error loading .env file")

	}

}
func GetSMTPConfig() *SMTPConfig {

	return &SMTPConfig{
		Host:     "smtp.mailtrap.io",
		Port:     "587",
		Username: os.Getenv("MAILTRAP_USERNAME"),
		Password: os.Getenv("MAILTRAP_PASSWORD"),
		From:     "assegagebeyehu21@gmail.com", // Your email from address
	}

}

func sendWelcomeEmail(toEmail, username string) error {
	smtpConfig := GetSMTPConfig()

	// Construct the email message
	subject := "Welcome to event management"
	body := fmt.Sprintf("Hello %s,\n\nWelcome to my event management website! i am  glad to have you with me.\n\nBest regards,\n gebeyehu", username)
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", smtpConfig.From, toEmail, subject, body)

	// Set up authentication information
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)

	// Send the email
	err := smtp.SendMail(
		smtpConfig.Host+":"+smtpConfig.Port,
		auth,
		smtpConfig.From,
		[]string{toEmail},
		[]byte(message),
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

type User struct {
	Event struct {
		Data struct {
			New struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"new"`
		} `json:"data"`
	} `json:"event"`
}

func WelcomeEmailHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	email := user.Event.Data.New.Email
	log.Printf("email to %s: %v", email)

	username := user.Event.Data.New.Username
	err = sendWelcomeEmail(email, username)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", user.Event.Data.New.Email, err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Email sent successfully to %s", user.Event.Data.New.Email)
}
