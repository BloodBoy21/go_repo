package main

import "fmt"

//SMS EMAIl
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotificationSender struct {

}
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type SMSNotification struct {

}
func (SMSNotification) SendNotification() {
	fmt.Println("Sending SMS notification")
}

func (SMSNotification) GetSender()ISender {
	return SMSNotificationSender{}
}
//Email
type EmailNotification struct{

}
func (EmailNotification) SendNotification(){
	fmt.Println("Sending Email")
}
type EmailNotificationSender struct{

}
func(EmailNotification)GetSender()ISender {
	return EmailNotificationSender{}
}
func (EmailNotificationSender)GetSenderMethod()string{
return "Email"
}
func (EmailNotificationSender)GetSenderChannel () string {
	return "SES"
}

//Factory
func getNotificationFactory(notificationType string)(INotificationFactory,error){
	if notificationType == "SMS"{
		return &SMSNotification{},nil
	}
	if notificationType == "Email"{
		return &EmailNotification{},nil
	}
	return nil,fmt.Errorf("no notification type")
}

func SendNotification(f INotificationFactory){
	f.SendNotification()
}

func getMethod(f INotificationFactory){
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main(){
	smsFactory,_ := getNotificationFactory("SMS")
	emailFactory,_ := getNotificationFactory("Email")
	SendNotification(smsFactory)
	SendNotification(emailFactory)
	getMethod(smsFactory)
	getMethod(emailFactory)

}