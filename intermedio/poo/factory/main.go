package main

import "fmt"

// Abstract Factory: Familias de objetos relacionados sin especificar sus clases concretas
// https://refactoring.guru/es/design-patterns/abstract-factory
// SMS | Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// -----------------------------------------------------
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Send notificationvia SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct{}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
