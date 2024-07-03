package repository

import (
	"context"
	"fmt"
	"log"
	"vyatsuAPIInventory/vyatsuAPI/models"
)

func (repo *PGrepo) GetEmployees(offset int, limit int) ([]models.Employee, error) {
	rows, err := repo.pool.Query(context.Background(), "select * from employees LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var employees []models.Employee
	for rows.Next() {
		e := models.Employee{}
		err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Position, &e.PersonnelNumber)
		if err != nil {
			fmt.Println(err)
			continue
		}
		employees = append(employees, e)
	}
	return employees, nil
}
func (repo *PGrepo) GetEmployeesWithFilterByPosition(offset int, limit int, position string) ([]models.Employee, error) {
	rows, err := repo.pool.Query(context.Background(), "select * from employees where position = $3 LIMIT $1 OFFSET $2", limit, offset, position)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var employees []models.Employee
	for rows.Next() {
		e := models.Employee{}
		err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Position, &e.PersonnelNumber)
		if err != nil {
			fmt.Println(err)
			continue
		}
		employees = append(employees, e)
	}
	return employees, nil
}
