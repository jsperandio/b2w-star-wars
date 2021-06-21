# b2w-star-wars

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

