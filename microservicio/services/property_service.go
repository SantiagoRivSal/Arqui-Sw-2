package services

import (
	propertyDao "microservicio/daos/property"
	"microservicio/dtos"
	model "microservicio/models"
	e "microservicio/utils/errors"
)

type propertyService struct{}

type propertyServiceInterface interface {
	GetProperty(id string) (dtos.PropertyDto, e.ApiError)

	InsertProperty(propertyDto dtos.PropertyDto) (dtos.PropertyDto, e.ApiError)
}

var (
	PropertyService propertyServiceInterface
)

func init() {
	PropertyService = &propertyService{}
}

func (s *propertyService) GetProperty(id string) (dtos.PropertyDto, e.ApiError) {

	var property model.Property = propertyDao.GetById(id)
	var propertyDto dtos.PropertyDto

	if property.Id.Hex() == "000000000000000000000000" {
		return propertyDto, e.NewBadRequestApiError("property not found")
	}
	propertyDto.Tittle = property.Tittle
	propertyDto.Size = property.Size
	propertyDto.Bathrooms = property.Bathrooms
	propertyDto.Service = property.Service
	propertyDto.Address.City = property.Address.City
	propertyDto.Address.State = property.Address.State
	propertyDto.Address.Country = property.Address.Country
	propertyDto.Address.Street = property.Address.Street
	propertyDto.Price = property.Price
	propertyDto.Rooms = property.Rooms
	propertyDto.Id = property.Id.Hex()
	return propertyDto, nil
}

func (s *propertyService) InsertProperty(propertyDto dtos.PropertyDto) (dtos.PropertyDto, e.ApiError) {

	var property model.Property

	property.Tittle = propertyDto.Tittle
	property.Size = propertyDto.Size
	property.Price = propertyDto.Price
	property.Rooms = propertyDto.Rooms
	property.Service = propertyDto.Service
	property.Bathrooms = propertyDto.Bathrooms
	property.Description = propertyDto.Description

	property.Address.City = propertyDto.Address.City
	property.Address.Country = propertyDto.Address.Country
	property.Address.State = propertyDto.Address.State
	property.Address.Street = propertyDto.Address.Street

	property = propertyDao.Insert(property)

	if property.Id.Hex() == "000000000000000000000000" {
		return propertyDto, e.NewBadRequestApiError("error in insert")
	}
	propertyDto.Tittle = property.Tittle
	propertyDto.Size = property.Size
	propertyDto.Bathrooms = property.Bathrooms
	propertyDto.Service = property.Service
	propertyDto.Address.City = property.Address.City
	propertyDto.Address.State = property.Address.State
	propertyDto.Address.Country = property.Address.Country
	propertyDto.Address.Street = property.Address.Street
	propertyDto.Price = property.Price
	propertyDto.Rooms = property.Rooms
	propertyDto.Id = property.Id.Hex()

	return propertyDto, nil
}
