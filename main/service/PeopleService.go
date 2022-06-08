package service

import (
	"GoEmployeeMgmt/main/model/people"
	"fmt"
)

type (
	PeopleService interface {
		GetAllPeople() ([]people.People, error)
		GetOnePerson(personId int) (*people.People, error)
	}
)

type defaultPeopleRepo struct {
	repo people.Repo
}

func (p defaultPeopleRepo) GetAllPeople() ([]people.People, error) {
	fmt.Println("entering service.GetAllPeople")

	allPeople, err := p.repo.FindAllPeople()
	if err != nil {
		return nil, err
	}
	fmt.Println("entering service.GetAllPeople")

	return allPeople, nil
}

func (p defaultPeopleRepo) GetOnePerson(personId int) (*people.People, error) {

	person, err := p.repo.FindOnePerson(personId)
	if err != nil {
		return nil, err
	}

	return person, nil

}

func NewPeopleService(repo people.Repo) defaultPeopleRepo {
	fmt.Println("entering NewPeopleService")
	return defaultPeopleRepo{repo: repo}
}
