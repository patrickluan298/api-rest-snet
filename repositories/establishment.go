package repositories

import (
	"api-rest-snet/models"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

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
	SelectEstablishment(id)

	fmt.Println("Establishment updated successfully")
	defer db.Close()
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

	fmt.Println("Id deleted successfully")

	defer db.Close()
	return nil
}
