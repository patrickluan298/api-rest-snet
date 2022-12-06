package models

type Establishment struct {
	Id                 int    `json:"id"`
	Nome               string `json:"nome"`
	RazaoSocial        string `json:"razao social"`
	Endereco           string `json:"endereco"`
	Estado             string `json:"estado"`
	Cidade             string `json:"cidade"`
	Cep                string `json:"cep"`
	NumEstabelecimento int    `json:"numero do estabelecimento"`
}

type Store struct {
	Id                int    `json:"id"`
	Nome              string `json:"nome"`
	RazaoSocial       string `json:"razao social"`
	Endereco          string `json:"endereco"`
	Estado            string `json:"estado"`
	Cidade            string `json:"cidade"`
	Cep               string `json:"cep"`
	NumLoja           int    `json:"numero da loja"`
	IdEstabelecimento int    `json:"id do estabelecimento"`
}

type Message struct {
	Message string `json:"message"`
}
