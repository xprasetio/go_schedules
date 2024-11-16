package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xprasetio/go_schedules.git/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
}
type service struct {
	client *http.Client // untuk get data dari api
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	// if err != nil {
	// 	return nil, err
	// }

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil { // Menangani kesalahan unmarshalling
		return nil, err
	}
	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}
	return
}
