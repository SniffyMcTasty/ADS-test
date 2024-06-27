Backend - Golang
==========================

## Tasks

The list of the tasks are listed in the [Readme](../../Readme.md) in the root directory of the bitbucket project.

## Requirement

* You need to have go v1.15.x at least to be able to use the go modules [Download Golang](https://golang.org/dl/)


## Structure of the project

```
golang
├── Readme.md                           # Readme
├── api                                 # Contain all the source file necessary for the api
│    ├── database.go                    # Contain the query to get the data from mysql
│    ├── models.go                      # Contain the model
│    ├── router.go                      # Contain all the endpoint of the api
│    └── routes.go                      # Contain the implementation of the handler related to the different endpoint of the api
│
├── docker-compose.yml                  # Docker Compose for the database MySQL
├── go.mod                              # List of go modules used by the project
└── main.go                             # Main file to run the project
```

Notes: 

* You are free to organize the content of the folder `api` as you want. 
* You are free to install other go module 
* You have to use the **gin** framework pre-installed

## How to run the project

### Run API

1) Create a file `.env` based on `.env.default` at the same level

2) Start the go project

```bash 
go build && ./backendapi
```

4) Check that everything is ok

``` 
curl http://localhost:8080/api/vehicle-makes
```

Sample response

```
{
    "vehicle_makes": [
        {
            "id": 1,
            "name": "Acura",
            "url": "https://www.acura.ca"
        }
    ]
}
```

### Documentation

* This app use [Gin](https://github.com/gin-gonic/gin). This is a web framework written in Go (Golang).
* Documentation can be found here: https://github.com/gin-gonic/gin
