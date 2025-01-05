package takeout

import (
	"encoding/json"
)

type service struct {
}

type Service interface {
	Get(fileBytes []byte) Takeout
}

func (s *service) Get(fileBytes []byte) Takeout {

	var takeout Takeout
	err := json.Unmarshal(fileBytes, &takeout.Transactions)
	if err != nil {
		panic(err)
	}
	return takeout

}

func NewService() Service {
	return &service{}
}
