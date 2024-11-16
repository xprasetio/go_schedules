package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/xprasetio/go_schedules.git/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	GetStationById(id string) (response []ScheduleResponse, err error)
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
	if err != nil {
		return nil, err
	}

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

func (s *service) GetStationById(id string) (response []ScheduleResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}
	var schedule []Schedule
	err = json.Unmarshal(byteResponse, &schedule)
	if err != nil {
		return
	}
	//schedule selected by id
	var scheduleSelected Schedule
	for _, item := range schedule {
		if item.StationId == id {
			scheduleSelected = item
			break
		}
	}
	if scheduleSelected.StationId == "" {
		return nil, errors.New("station not found")
	}
	response, err = ConvertDataToResponse(scheduleSelected)
	if err != nil {
		return
	}
	return
}

func ConvertDataToResponse(schedule Schedule) (response []ScheduleResponse, err error) {
	var (
		LebakBulusTripName = "Lebak Bulus"
		BundaranHITripName = "Bundaran Hilir"
	)
	scheduleBundaranHI := schedule.ScheduleBundaranHI
	scheduleLebakBulus := schedule.ScheduleLebakBulus

	scheduleBundaranHIparsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)
	if err != nil {
		return nil, err
	}
	scheduleLebakBulusparsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)
	if err != nil {
		return nil, err
	}

	// convert to response
	for _, item := range scheduleBundaranHIparsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: BundaranHITripName,
				Time:        item.Format("15:04"),
			})
		}
	}
	for _, item := range scheduleLebakBulusparsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: LebakBulusTripName,
				Time:        item.Format("15:04"),
			})
		}
	}
	return
}

func ConvertScheduleToTimeFormat(schedule string) (response []time.Time, err error) {
	var (
		parsedTime time.Time
		schedules  = strings.Split(schedule, ",")
	)
	for _, item := range schedules {
		trimmedItem := strings.TrimSpace(item)
		if trimmedItem == "" {
			continue
		}
		parsedTime, err = time.Parse("15:04", trimmedItem)
		if err != nil {
			return nil, errors.New("format waktu tidak valid")
		}
		response = append(response, parsedTime)
	}
	return
}
