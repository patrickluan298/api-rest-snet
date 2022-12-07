package repository

import (
	"api-rest-snet/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "753159"
	Dbname   = "api-rest-snet"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Connection() {
	connection := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

	db, err = sql.Open("postgres", connection)
	checkErr(err)
	err = db.Ping()
	checkErr(err)
}

func InsertEstablishment(e *models.Establishment) error {
	Connection()

	query := `insert into Establishment (Nome, RazaoSocial, Endereco, Estado, Cidade, Cep, NumEstablishment) values ($1, $2, $3, $4, $5, $6, $7) returning id`
	id := 0
	err = db.QueryRow(query, e.Nome, e.RazaoSocial, e.Endereco, e.Estado, e.Cidade, e.Cep, e.NumEstabelecimento).Scan(&id)
	defer db.Close()

	if err != nil {
		fmt.Println("Error registering establishment:", err.Error())

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("Existing establishment number")
		}

		return err
	}

	SelectEstablishment(id)
	fmt.Println("Establishment successfully registered. Registration ID:", id)

	return nil
}

func SelectEstablishment(id int) error {
	Connection()

	query := `select id from Establishment where id = $1`
	r, _ := db.Query(query, id)
	defer db.Close()

	var idT int = 0
	r.Next()
	if err := r.Scan(&idT); err != nil {
		fmt.Println("Error selecting establishment:", err.Error())

		if strings.Contains(err.Error(), "Rows are closed") {
			return errors.New("The entered ID is invalid")
		}

		return err
	} else {
		if id == idT {
			query := `select * from Establishment where id = $1`
			r, _ := db.Query(query, id)

			for r.Next() {
				e := models.Establishment{}
				err = r.Scan(&e.Id, &e.Nome, &e.RazaoSocial, &e.Endereco, &e.Estado, &e.Cidade, &e.Cep, &e.NumEstabelecimento)
				checkErr(err)

				fmt.Println("Id:", e.Id)
				fmt.Println("Nome:", e.Nome)
				fmt.Println("Razão Social:", e.RazaoSocial)
				fmt.Println("Endereço:", e.Endereco)
				fmt.Println("Estado:", e.Estado)
				fmt.Println("Cidade:", e.Cidade)
				fmt.Println("Cep:", e.Cep)
				fmt.Println("Número do Estabelecimento:", e.NumEstabelecimento)
				fmt.Println("")
			}
		}
		return nil
	}
}

func UpdateEstablishment(e *models.Establishment, id int) {
	Connection()

	if e.Nome != "" {
		query := `UPDATE Establishment set Nome = $1 WHERE id = $2`
		_, err = db.Exec(query, e.Nome, id)
	}
	if e.RazaoSocial != "" {
		query := `UPDATE Establishment set RazaoSocial = $1 WHERE id = $2`
		_, err = db.Exec(query, e.RazaoSocial, id)
	}
	if e.Endereco != "" {
		query := `UPDATE Establishment set Endereco = $1 WHERE id = $2`
		_, err = db.Exec(query, e.Endereco, id)
	}
	if e.Estado != "" {
		query := `UPDATE Establishment set Estado = $1 WHERE id = $2`
		_, err = db.Exec(query, e.Estado, id)
	}
	if e.Cidade != "" {
		query := `UPDATE Establishment set Cidade = $1 WHERE id = $2`
		_, err = db.Exec(query, e.Cidade, id)
	}
	if e.Cep != "" {
		query := `UPDATE Establishment set Cep = $1 WHERE id = $2`
		_, err = db.Exec(query, e.Cep, id)
	}

	checkErr(err)
	defer db.Close()

	SelectEstablishment(id)
	fmt.Println("Establishment updated successfully")
}

func SelectAllEstablishments() error {
	Connection()

	r, err := db.Query(`select * from Establishment`)
	checkErr(err)
	defer db.Close()

	if err != nil {
		fmt.Println("Error when selecting all establishments:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return errors.New("Error when selecting all establishments")
		}

		return err
	}

	for r.Next() {
		e := models.Establishment{}
		err = r.Scan(&e.Id, &e.Nome, &e.RazaoSocial, &e.Endereco, &e.Estado, &e.Cidade, &e.Cep, &e.NumEstabelecimento)
		checkErr(err)

		fmt.Println("Id:", e.Id)
		fmt.Println("Nome:", e.Nome)
		fmt.Println("Razão Social:", e.RazaoSocial)
		fmt.Println("Endereço:", e.Endereco)
		fmt.Println("Estado:", e.Estado)
		fmt.Println("Cidade:", e.Cidade)
		fmt.Println("Cep:", e.Cep)
		fmt.Println("Número do Estabelecimento:", e.NumEstabelecimento)
		fmt.Println("")
	}

	return nil
}

func DeleteEstablishment(id int) error {
	err = SelectEstablishment(id)

	if err != nil {
		fmt.Println("Error deleting establishment:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return errors.New("The entered ID is invalid")
		}

		return err
	}

	Connection()
	query := `delete from Establishment where id = $1`
	_, err = db.Exec(query, id)

	defer db.Close()
	fmt.Println("Id deleted successfully")

	return nil
}

// func (e *models.Establishment) Validate() error {
// 	return validation.ValidateStruct(
// 		validation.Field(&e.Nome, validation.Required, validation.Length(5, 50)),
// 		validation.Field(&e.RazaoSocial, validation.Required, validation.Length(5, 50)),
// 		validation.Field(&e.Endereco, validation.Required, validation.Length(5, 50)),
// 		validation.Field(&e.Estado, validation.Required),
// 		validation.Field(&e.Cidade, validation.Required),
// 		validation.Field(&e.Cep, validation.Required, validation.Length(10, 10)),
// 		validation.Field(&e.NumEstablishment, validation.Required),
// 	)
// }

// e := Establishment{}
// err := e.Validate()