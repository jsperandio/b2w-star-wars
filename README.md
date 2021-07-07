# b2w-star-wars

<p align="center">
  <img src="https://logos-download.com/wp-content/uploads/2016/09/Star_Wars_logo-1-700x300.png" width="800" height="220">
</p>

API REST para listar os planetas da franquia Star Wars.


### Feito usando

* [Go](https://golang.org/)
* [MongoDB](https://www.mongodb.com/)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)

#### Com uso de libs:

* [Go Fiber (Web Framework)](https://github.com/gofiber/fiber)
* [Mongo-driver (driver for MongoDB)](https://github.com/mongodb/mongo-go-driver)
* [Bigcache (in-memory cache)](https://github.com/allegro/bigcache)
* [Resty (Resty client lib)](https://github.com/go-resty/resty)
* [Httpmock (Mocking lib)](https://github.com/jarcoal/httpmock)

### Uso

#### Pré-requisitos

* Ambiente com Docker e Docker Compose :).

#### Usando Aplicação

```bash
$ docker-compose up --build
```

#### Endpoints


| Nome | Path | Method | Content-Type | Descrição |
| ------ | ------ | ------ | ------ | ------ |
| Listar planetas| api/v1/planets | GET | application/json | Retornar todos os planetas cadastrados no banco de dados. |
| Buscar planeta| api/v1/planet/:id | GET | application/json | Retorna o planeta pelo codigo. |
| Buscar planeta pelo nome| api/v1/planet/?name={Nome} | GET | application/json | Retorna o planeta pelo seu nome inputado. |
| Inserir planeta | api/v1/planet | POST | application/json | Insere um planeta na base de dados de acordo com o layout proposto. |
| Deletar planeta | api/v1/planet/:id | DELETE | application/json | Deleta o planeta do banco de dados de acordo com o id passado. |

## Testes

### Executar test
```bash
$ go test ./...
```

Neste projeto foi usado gerador de mocks, a ferramenta: 
* [mockery (Mocking lib)](https://github.com/vektra/mockery)