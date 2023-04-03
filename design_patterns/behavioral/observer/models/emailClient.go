package models

import (
	"fmt"
	"math/rand"
)

type EmailClient struct {
	id string
}

func NewEmailClient() *EmailClient {
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	length := 4
	b := make([]rune, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return &EmailClient{
		id: string(b),
	}
}

func (e *EmailClient) GetId() string {
	return e.id
}

func (e *EmailClient) UpdateValues(value string) {
	fmt.Printf("Sending Email - %s is avalaible from client %s\n", value, e.id)
}
