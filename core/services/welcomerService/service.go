package welcomerService

import (
	"multiverse/core/models"
	"multiverse/core/services/welcomerService/client"
	"multiverse/core/services/welcomerService/welcomepb"
	"os"
)

type WelcomerService interface {
	GetWelcomeMessage(user *models.User) (*welcomepb.WelcomeResponse, error)
}

func NewWelcomerService() WelcomerService {
	welcomerServer := os.Getenv("WELCOMER_SERVER")
	welcomerPort := os.Getenv("WELCOMER_PORT")
	WelcomerUseSSl := os.Getenv("WELLCOMER_USE_SSL")
	useSsl := false
	if welcomerServer == "" {
		welcomerServer = "welcomer"
	}
	if welcomerPort == "" {
		welcomerPort = "8080"
	}
	if WelcomerUseSSl == "" || WelcomerUseSSl == "false" {
		useSsl = false
	} else {
		useSsl = true
	}
	return &welcomerService{
		Host:   welcomerServer,
		Port:   welcomerPort,
		UseSSl: useSsl,
	}
}

type welcomerService struct {
	Host   string
	Port   string
	UseSSl bool
}

func (s *welcomerService) GetWelcomeMessage(user *models.User) (*welcomepb.WelcomeResponse, error) {
	conn, err := client.NewWelcomerConnection(s.UseSSl, s.Host, s.Port)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cli, err := client.NewWelcomerClient(conn)
	if err != nil {
		return nil, err
	}
	res, err := cli.Welcome(welcomepb.UserInfo{
		Name:    user.Name,
		Country: user.Address,
		Age:     21,
	})
	return res, err
}
