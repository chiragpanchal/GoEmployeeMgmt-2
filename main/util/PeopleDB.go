package util

import (
	"GoEmployeeMgmt/main/model/people"
	"fmt"
	"log"
)

func (rd RepoDb) FindAllPeople() ([]people.People, error) {
	fmt.Println("entering util.FindAllPeople")
	pgClient := rd.DbClient

	findAllSql := "select person_id, first_name, last_name, gender, marital_status, date_of_birth, joining_date, COALESCE(job_id,0)job_id , COALESCE(department_id,0) department_id, COALESCE(location_id,0) location_id, employee_number from dm_person per"

	peopleList := make([]people.People, 0)

	err := pgClient.Select(&peopleList, findAllSql)
	if err != nil {
		log.Fatal("Error while querying people records" + err.Error())
		return nil, err
	}
	fmt.Println("entering util.FindAllPeople")
	return peopleList, nil

}

func (rd RepoDb) FindOnePerson(personId int) (*people.People, error) {
	pgClient := rd.DbClient

	findOneSql := "select person_id, first_name, last_name, gender, marital_status, date_of_birth, joining_date, job_id, department_id, location_id, employee_number from dm_person per where per.person_id=$1"

	var findPerson people.People

	err := pgClient.Get(&findPerson, findOneSql)
	if err != nil {
		return nil, err
	}

	return &findPerson, nil
}
