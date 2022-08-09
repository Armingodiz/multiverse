package welcomerService

import (
	"multiverse/core/models"
	"multiverse/core/services/welcomerService/client"
	"multiverse/core/services/welcomerService/welcomepb"
)

type WelcomerService interface {
	GetWelcomeMessage(user *models.User) (*welcomepb.WelcomeResponse, error)
}

func NewWelcomerService() WelcomerService {
	return &welcomerService{}
}

type welcomerService struct {
}

func (s *welcomerService) GetWelcomeMessage(user *models.User) (*welcomepb.WelcomeResponse, error) {
	conn, err := client.NewWelcomerConnection(false, "welcomer", "8080")
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
