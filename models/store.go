package models

type Store struct {
	Id                int    `json:"id"`
	Nome              string `json:"nome" validate:"nonzero"`
	RazaoSocial       string `json:"razao social" validate:"min=5"`
	Endereco          string `json:"endereco" validate:"nonzero"`
	Estado            string `json:"estado" validate:"nonzero"`
	Cidade            string `json:"cidade" validate:"nonzero"`
	Cep               string `json:"cep" validate:"len=8, regexp=^[0-9]*$"`
	NumLoja           int    `json:"numero da loja" validate:"nonzero"`
	IdEstabelecimento int    `json:"id do estabelecimento" validate:"nonzero"`
}
