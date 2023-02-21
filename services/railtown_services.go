package services

import (
	"encoding/json"
	"net/http"
	"railtown/domain/railtown"
	"railtown/domain/requests"
	"railtown/logger"
	"railtown/utills/date_utills"
	"railtown/utills/responses"
)

var (
	RailTownsService railTownServiceInterface = &railTownsService{}
)

type railTownsService struct{}

type railTownServiceInterface interface {
	GETTemperatureService(*railtown.RailTownREQ) (*railtown.RailTownRES, *responses.Response)
}

func (s *railTownsService) GETTemperatureService(req *railtown.RailTownREQ) (*railtown.RailTownRES, *responses.Response) {
	logger.Info("Enter to GETTemperatureService service successfully")

	request := requests.Request{}

	request.Init("https://api.weatherapi.com/v1/current.json", "GET", nil)

	request.AddToURL(req.Query)
	body, err := request.GET()
	if err != nil {
		return nil, err
	}

	ret := railtown.RailTownRES{}
	if marshalErr := json.Unmarshal(body, &ret); marshalErr != nil {
		logger.Error("error when trying to marshal response to struct", marshalErr)
		err := responses.NewBadRequestError("", "", http.StatusBadRequest)
		return nil, err
	}

	ret.CreatedAt = date_utills.DateTimeService.GetNowDBFormat()

	//saveErr := ret.Save()
	//if saveErr != nil {
	//	return nil, err
	//}

	logger.Info("Close from GETTemperatureService service successfully")
	return &ret, nil
}
