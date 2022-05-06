package services

import (
	"github.com/klasrak/go-meli-test-dojo/clients/swapi"
	"github.com/klasrak/go-meli-test-dojo/models"
)

func GetStarshipService(id int) (models.Starship, error) {
	return swapi.Instance.GetStarship(id)
}

func GetStarshipsService() (models.Starships, error) {
	return swapi.Instance.GetStarships()
}

func GetPeopleService(id int) (models.People, error) {
	return swapi.Instance.GetPeople(id)
}

func GetPeopleListService() (models.PeopleList, error) {
	return swapi.Instance.GetPeopleList()
}
