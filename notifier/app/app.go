package app

import (
	"log"
	"multiverse/notifier/config"
	"multiverse/notifier/services/brokerService"
	"multiverse/notifier/services/emailService"
)

type App struct {
	BrokerService brokerService.BrokerService
	MailService   emailService.MailService
}

func NewApp() *App {
	return &App{
		BrokerService: brokerService.NewBrokerService(),
		MailService:   emailService.NewMailService(config.Configs.App.SenderEmail),
	}
}

func (a *App) Start() error {
	taskChann, errChann, err := a.BrokerService.StartConsuming()
	if err != nil {
		return err
	}
	defer a.BrokerService.CloseConnection()
	defer a.BrokerService.CloseChannel()
	go func() {
		for task := range taskChann {
			log.Println("Received task:", task)
			err := a.MailService.SendEmail(task.Target, task.Text)
			if err != nil {
				log.Println("Error sending email:", err)
				errChann <- err
			}
		}
	}()
	for err := range errChann {
		log.Printf("Received error: %s", err.Error())
		return err
	}
	return nil
}
