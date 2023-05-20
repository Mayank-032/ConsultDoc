package routes

import (
	"errors"
	"healthcare-service/domain/entity"
)

type FetchDoctorListRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (req FetchDoctorListRequest) Validate() error {
	if len(req.Latitude) == 0 || len(req.Longitude) == 0 {
		return errors.New("address details missing")
	}
	return nil
}

func (req FetchDoctorListRequest) toAddressDto() entity.Address {
	return entity.Address{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}
