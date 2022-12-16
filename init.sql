CREATE TABLE Establishment(
	Id serial primary key,
	Nome varchar(50),
	RazaoSocial varchar(50),
	Endereco varchar(100),
	Estado varchar(20),
	Cidade varchar(20),
	Cep varchar(15),
	NumEstablishment integer unique
);

CREATE TABLE Store(
	Id serial primary key,
	Nome varchar(50),
	RazaoSocial varchar(50),
	Endereco varchar(100),
	Estado varchar(20),
	Cidade varchar(20),
	Cep varchar(15),
	NumLoja integer unique,
	IdEstabelecimento integer
);

alter table Store
add constraint fk_store
foreign key (IdEstabelecimento)
references Establishment (NumEstablishment);