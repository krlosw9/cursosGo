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

// -----------------------------------------------------
type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	fmt.Println("Send notification Email")
}

type EmailNotificationSender struct{}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

// -----------------------------------------------------
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("no notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
