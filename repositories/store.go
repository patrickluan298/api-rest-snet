package repositories

import (
	"api-rest-snet/models"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func InsertStore(e *models.Store) error {
	Connection()

	query := `insert into Store (Nome, RazaoSocial, Endereco, Estado, Cidade, Cep, NumLoja, IdEstabelecimento) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`
	id := 0
	err = db.QueryRow(query, e.Nome, e.RazaoSocial, e.Endereco, e.Estado, e.Cidade, e.Cep, e.NumLoja, e.IdEstabelecimento).Scan(&id)
	defer db.Close()

	if err != nil {
		fmt.Println("Error registering store:", err.Error())

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("Existing store number")
		}
		if strings.Contains(err.Error(), "insert or update on table \"store\" violates foreign key constraint") {
			return errors.New("Unregistered establishment")
		}

		return err
	}

	SelectStore(id)
	fmt.Println("Store successfully registered. Registration ID:", id)

	return nil
}

func SelectStore(NumLoja int) error {
	Connection()

	query := `select NumLoja from Store where NumLoja = $1`
	r, _ := db.Query(query, NumLoja)
	defer db.Close()

	var NumLojaDB int = 0
	r.Next()
	if err := r.Scan(&NumLojaDB); err != nil {
		fmt.Println("Error selecting store:", err.Error())

		if strings.Contains(err.Error(), "Rows are closed") {
			return errors.New("The entered ID is invalid")
		}

		return err
	} else {
		if NumLoja == NumLojaDB {
			query := `select * from Store where NumLoja = $1`
			r, _ := db.Query(query, NumLoja)

			for r.Next() {
				e := models.Store{}
				err = r.Scan(&e.Id, &e.Nome, &e.RazaoSocial, &e.Endereco, &e.Estado, &e.Cidade, &e.Cep, &e.NumLoja, &e.IdEstabelecimento)
				checkErr(err)

				fmt.Println("Id:", e.Id)
				fmt.Println("Nome:", e.Nome)
				fmt.Println("Razão Social:", e.RazaoSocial)
				fmt.Println("Endereço:", e.Endereco)
				fmt.Println("Estado:", e.Estado)
				fmt.Println("Cidade:", e.Cidade)
				fmt.Println("Cep:", e.Cep)
				fmt.Println("Número da loja:", e.NumLoja)
				fmt.Println("Id do estabelecimento:", e.IdEstabelecimento)
				fmt.Println("")
			}
		}
		return nil
	}
}

func UpdateStore(e *models.Store, NumLoja int) {
	Connection()

	if e.Nome != "" {
		query := `UPDATE Store set Nome = $1 WHERE NumLoja = $2`
		_, err = db.Exec(query, e.Nome, NumLoja)
	}
	if e.RazaoSocial != "" {
		query := `UPDATE Store set RazaoSocial = $1 WHERE NumLoja = $2`
		_, err = db.Exec(query, e.RazaoSocial, NumLoja)
	}
	if e.Endereco != "" {
		query := `UPDATE Store set Endereco = $1 WHERE NumLoja = $2`
		_, err = db.Exec(query, e.Endereco, NumLoja)
	}
	if e.Estado != "" {
		query := `UPDATE Store set Estado = $1 WHERE NumLoja = $2`
		_, err = db.Exec(query, e.Estado, NumLoja)
	}
	if e.Cidade != "" {
		query := `UPDATE Store set Cidade = $1 WHERE NumLoja = $2`
		_, err = db.Exec(query, e.Cidade, NumLoja)
	}
	if e.Cep != "" {
		query := `UPDATE Store set Cep = $1 WHERE NumLoja = $2`
		_, err = db.Exec(query, e.Cep, NumLoja)
	}
	if e.NumLoja != 0 {
		query := `UPDATE Store set NumLoja = $1 WHERE NumLoja = $2`
		_, err = db.Exec(query, e.NumLoja, NumLoja)
	}
	checkErr(err)
	SelectStore(NumLoja)

	fmt.Println("Store updated successfully")
	defer db.Close()
}

func SelectAllStores() error {
	Connection()

	r, err := db.Query(`select * from Store`)
	checkErr(err)
	defer db.Close()

	if err != nil {
		fmt.Println("Error when selecting all stores:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return errors.New("Error when selecting all stores")
		}

		return err
	}

	for r.Next() {
		e := models.Store{}
		err = r.Scan(&e.Id, &e.Nome, &e.RazaoSocial, &e.Endereco, &e.Estado, &e.Cidade, &e.Cep, &e.NumLoja, &e.IdEstabelecimento)
		checkErr(err)

		fmt.Println("Id:", e.Id)
		fmt.Println("Nome:", e.Nome)
		fmt.Println("Razão Social:", e.RazaoSocial)
		fmt.Println("Endereço:", e.Endereco)
		fmt.Println("Estado:", e.Estado)
		fmt.Println("Cidade:", e.Cidade)
		fmt.Println("Cep:", e.Cep)
		fmt.Println("Número da loja:", e.NumLoja)
		fmt.Println("Id do estabelecimento:", e.IdEstabelecimento)
		fmt.Println("")
	}
	return nil
}

func DeleteStore(NumLoja int) error {
	err := SelectStore(NumLoja)

	if err != nil {
		fmt.Println("Error deleting store:", err.Error())

		if strings.Contains(err.Error(), "invalid memory address or nil pointer dereference goroutine") {
			return errors.New("The entered ID is invalid")
		}

		return err
	}

	Connection()
	query := `delete from Store where NumLoja = $1`
	_, err = db.Exec(query, NumLoja)

	fmt.Println("Store deleted successfully")

	defer db.Close()
	return nil
}
