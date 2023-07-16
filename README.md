
# Chistes-API / Jokes-API

Esta documentación describe los endpoints disponibles en la API REST de GitHub para interactuar con diferentes aspectos del sistema. La API de GitHub se basa en la arquitectura REST y utiliza el formato JSON para la representación de datos.

---


This documentation describes the available endpoints in the GitHub REST API for interacting with different aspects of the system. The GitHub API is based on the REST architecture and uses the JSON format for data representation.


 
## Tech Stack

* Golang
* [Fiber](https://docs.gofiber.io/)
* [database/sql](https://pkg.go.dev/database/sql)
* [dotenv](https://pkg.go.dev/poseur.com/dotenv)


## API Reference

### Get frases


Solicita una frase aleatoria de una base de datos que contiene 85 frases.

Request a random phrase from a database that contains 85 phrases.

```http
  GET /api/frases
```

#### Response
```json
[
    {
        "text": "* En los aviones el tiempo se pasa volando. "
    }
]
```

### Get chistes


Solicita un chiste aleatorio de una base de datos que contiene 256 chistes.

Request a random joke from a database that contains 256 jokes.

```http
  GET /api/chistes
```

#### Response
```json
[
    {
        "text": "- Mamá, mamá, (Una cerdita a su madre): ¿porqué tenemos una raja abajo? Por que si la tuvieras arriba serías una alcancía. "
    }
]
```

### Get refranes

Solicita un refran aleatorio de una base de datos que contiene 2780 refranes.

Request a random proverb from a database that contains 2780 proverbs.

```http
  GET /api/refranes
```

#### Response
```json
[
    {
        "text": "El que la sigue, la consigue.,"
    }
]
```



## Usado por / Used By:

Este proyecto es usado en :

This project is used by :

- [payaso-bot](https://github.com/liensanchez/bot-golang)


