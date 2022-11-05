package services

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	propertyDao "microservicio/daos/property"
	"microservicio/dtos"
	model "microservicio/models"
	e "microservicio/utils/errors"
	"net/http"
	"os"
	"sync"
)

type propertyService struct{}

type propertyServiceInterface interface {
	GetProperty(id string) (dtos.PropertyDto, e.ApiError)
	GetProperties() (dtos.PropertiesDto, e.ApiError)
	InsertMany(propertiesDto dtos.PropertiesDto) (dtos.PropertiesDto, e.ApiError)
	InsertProperty(propertyDto dtos.PropertyDto) (dtos.PropertyDto, e.ApiError)
}

var (
	PropertyService propertyServiceInterface
)

func init() {
	PropertyService = &propertyService{}
}

func (s *propertyService) GetProperties() (dtos.PropertiesDto, e.ApiError) {
	var properties = propertyDao.GetAll()
	var propertiesDtoArray dtos.PropertiesDto

	var wg sync.WaitGroup
	wg.Add(len(properties))

	for _, property := range properties {
		var propertyDto dtos.PropertyDto

		if property.Id.Hex() == "000000000000000000000000" {
			return propertiesDtoArray, e.NewBadRequestApiError("error in insert")
		}
		go func(url string) {
			defer wg.Done()
			fileName := RandStringBytes() + ".jpg"
			fmt.Println("Downloading", url, "to", fileName)

			output, err := os.Create(fileName)
			if err != nil {
				log.Fatal("Error while creating", fileName, "- ", err)
			}
			defer output.Close()

			res, err := http.Get(url)
			if err != nil {
				log.Fatal("http get error: ", err)
			} else {
				defer res.Body.Close()
				_, err = io.Copy(output, res.Body)
				if err != nil {
					log.Fatal("Error while downloading", url, "-", err)

				} else {
					fmt.Println("Downloaded", fileName)
				}
			}
		}(property.Image)
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
		propertyDto.Image = property.Image
		propertyDto.Id = property.Id.Hex()

		propertiesDtoArray = append(propertiesDtoArray, propertyDto)
	}
	wg.Wait()
	return propertiesDtoArray, nil

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (s *propertyService) GetProperty(id string) (dtos.PropertyDto, e.ApiError) {

	var property model.Property = propertyDao.GetById(id)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(url string) {
		defer wg.Done()
		fileName := RandStringBytes() + ".jpg"
		fmt.Println("Downloading", url, "to", fileName)

		output, err := os.Create(fileName)
		if err != nil {
			log.Fatal("Error while creating", fileName, "- ", err)
		}
		defer output.Close()

		res, err := http.Get(url)
		if err != nil {
			log.Fatal("http get error: ", err)
		} else {
			defer res.Body.Close()
			_, err = io.Copy(output, res.Body)
			if err != nil {
				log.Fatal("Error while downloading", url, "-", err)

			} else {
				fmt.Println("Downloaded", fileName)
			}
		}
	}(property.Image)
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
	propertyDto.Image = property.Image
	propertyDto.Id = property.Id.Hex()
	wg.Wait()
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
	property.Image = propertyDto.Image
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
	propertyDto.Image = property.Image
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

func (s *propertyService) InsertMany(propertiesDto dtos.PropertiesDto) (dtos.PropertiesDto, e.ApiError) {
	var propertiesDtoArray dtos.PropertiesDto

	for _, propertyDto := range propertiesDto {
		var property model.Property

		property.Tittle = propertyDto.Tittle
		property.Size = propertyDto.Size
		property.Image = propertyDto.Image
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
			return propertiesDto, e.NewBadRequestApiError("error in insert")
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
		propertyDto.Image = property.Image
		propertyDto.Rooms = property.Rooms
		propertyDto.Id = property.Id.Hex()

		propertiesDtoArray = append(propertiesDtoArray, propertyDto)
	}

	return propertiesDtoArray, nil
}
