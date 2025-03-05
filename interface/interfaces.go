package main

import (
	"flag"
	"fmt"
)

type Target interface {
	Send(mesage string) error
}

type EmailTarget struct {
	Email string
}

func (e *EmailTarget) Send(mesage string) error {
	fmt.Printf("Sending Email to %s message : %s \n", e.Email, mesage)
	return nil

}

type SMSTarget struct {
	Ph_No string
}

func (s *SMSTarget) Send(mesage string) error {
	fmt.Printf("Sending SMS to %s message : %s \n", s.Ph_No, mesage)
	return nil
}

type PushNotificationTarget struct {
	Id string
}

func (p *PushNotificationTarget) Send(mesage string) error {
	fmt.Printf("Sending PushNotificationTarget to %s message : %s \n", p.Id, mesage)
	return nil
}

type NotificationManager struct {
	Targets []Target
}

func (nM *NotificationManager) AddTarget(target Target) {
	nM.Targets = append(nM.Targets, target)
}
func (nM *NotificationManager) SendToAllTargets(mesage string) {

	for _, t := range nM.Targets {
		err := t.Send(mesage)
		if err != nil {
			fmt.Println("Error in sending message ", err)
		}
	}
}
func main() {
	emailFlag := flag.String("email", "", "")
	smsFlag := flag.String("sms", "", "")
	pushNotiFlag := flag.String("ph_no", "", "")
	flag.Parse()
	notificationManager := &NotificationManager{}
	if *emailFlag != "" {
		emailTarget := &EmailTarget{Email: *emailFlag}
		notificationManager.AddTarget(emailTarget)
	}
	if *smsFlag != "" {
		smsTarget := &SMSTarget{Ph_No: *smsFlag}
		notificationManager.AddTarget(smsTarget)
	}
	if *pushNotiFlag != "" {
		pushNoTarget := &PushNotificationTarget{Id: *pushNotiFlag}
		notificationManager.AddTarget(pushNoTarget)
	}
	notificationManager.SendToAllTargets("Test Notification")

}
