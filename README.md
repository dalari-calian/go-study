# Projeto CRUD em GoLang para Gestão de Clientes e Carros

Este é um projeto em GoLang que implementa um CRUD (Create, Read, Update, Delete) para gestão de clientes e carros. O projeto utiliza uma arquitetura simples e eficaz para interação com um banco de dados relacional.

## Funcionalidades

- Cadastro, leitura, atualização e exclusão de clientes
- Cadastro, leitura, atualização e exclusão de carros
- Restrição: Não é possível excluir um cliente se ele possuir carros cadastrados. Primeiro é necessário excluir os carros associados ao cliente.

## Pré-requisitos

- GoLang instalado
- Banco de dados relacional (MySQL, PostgreSQL, etc.)
- Configurar a URL do banco de dados no arquivo `db.go`

## Como executar o projeto

1. Clone o repositório do projeto:

git clone <https://github.com/dalari-calian/go-study.git>

2. Configure a URL do banco de dados no arquivo `db.go` conforme o seu ambiente:

```go
// Arquivo: db.go

const (
    host     = "localhost"
    port     = 5432 // Porta padrão do PostgreSQL
    user     = "seu_usuario"
    password = "sua_senha"
    dbname   = "nome_do_banco_de_dados"
)
```
1 - Navegue até o diretório cmd:
```cd cmd```

2 - Execute o comando para rodar o projeto:
```go run main.go```

## Teste Unitários

1 - Navegue até a pasta tests
```cd tests```

2 - Execute o comando para rodar os testes:
```go test -v tests```

## Autor
Calian Dalari - cdalari@gmail.com
