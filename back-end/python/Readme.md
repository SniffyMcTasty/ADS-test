Backend - Python
=================

## Tasks

The list of the tasks are listed in the [Readme](../../Readme.md) in the root directory of the bitbucket project.

## Requirement

* You need to have python v3.11.x or higher [Download Python](https://www.python.org/downloads/)

## Structure of the project

```
python
├── Readme.md                           # Readme
├── app                                 # Contain all the source file necessary for the api
│   ├── __init__.py
│   ├── api                             # Contain all the endpoint of the api
│   │   ├── __init__.py
│   │   └── vehicle.py                  
│   ├── models                          # Contain the models
│   │   ├── __init__.py
│   │   └── entities.py
│   ├── repositories                    # Contain the repository to access the data in the database
│   │   ├── __init__.py
│   │   └── vehicle.py
│   └── services                        # Contain the services
│       ├── __init__.py
│       └── vehicle.py
├── config.py
│
├── docker-compose.yml                  # Docker Compose for the database MySQL
├── .flaskenv                           # Env variable use by the framework flask
├── requirements.txt                    # Python dependencies required
└── run.py                              # Main file to run the project
```

Notes:

* You are free to organize the content of the folder `app` as you want.
* You are free to install other library
* You have to use the **flask** framework pre-installed

## How to run the project

### Setup the project

1) Create a file `.env` based on `.env.default` at the same level

2) Install the dependencies
``` 
pip install -r requirements.txt
```

### Run API

1) Start the python project

``` 
flask run
```

2) Check that everything is ok

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

3) Run the tests

``` 
flask test
```

### Documentation

This app use [Flask](https://flask.palletsprojects.com/en/2.2.x/) . This is a web framework for python.

We use [SqlAlchemy](https://www.sqlalchemy.org/) to access to the mysql database.
