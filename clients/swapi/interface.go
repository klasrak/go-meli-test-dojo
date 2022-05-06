package swapi

import "github.com/klasrak/go-meli-test-dojo/models"

type Client interface {
	GetStarship(id int) (models.Starship, error)
	GetStarships() (models.Starships, error)
	GetPeople(id int) (models.People, error)
	GetPeopleList() (models.PeopleList, error)
}

var (
	defaultInstance Client = NewSWAPIClient()
	Instance               = defaultInstance
)
