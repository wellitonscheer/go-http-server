package whats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Template struct {
	Name     string `json:"name"`
	Language struct {
		Code string `json:"code"`
	} `json:"language"`
}

type Message struct {
	MessagingProduct string   `json:"messaging_product"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Template         Template `json:"template"`
}

func SendMessage(to string) error {
	if to == "" {
		return fmt.Errorf("Invalid parameter")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	phoneNumberId := os.Getenv("PHONE_NUMBER_ID")
	authToken := os.Getenv("AUTH_TOKEN")

	url := fmt.Sprint("https://graph.facebook.com/v20.0/", phoneNumberId, "/messages")

	msg := Message{
		MessagingProduct: "whatsapp",
		To:               to,
		Type:             "template",
		Template: Template{
			Name: "hello_world",
			Language: struct {
				Code string `json:"code"`
			}{
				Code: "en_US",
			},
		},
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response: %v", resp.Status)
	}

	fmt.Println("Message sent successfully!")
	return nil
}
