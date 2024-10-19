package email

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/antremis/server/internal/MailClient"
	"github.com/google/uuid"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	var jsonBody emailSchema
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		log.Fatalf("Unable to decode json body: %v", err)
	}
	to := jsonBody.To
	subject := jsonBody.Subject
	body := jsonBody.Body
	tags := jsonBody.Tags

	message := fmt.Sprintf(`From: <akash.annonymail@gmail.com>
To: <%s>
Subject: %s

%s



ID: %s
%s`, to, subject, body, uuid.New(), parseTags(tags))
	
	if err := MailClient.SendMail(message); err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, fmt.Sprintf("Failed to send email: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Email sent successfully!")
}
