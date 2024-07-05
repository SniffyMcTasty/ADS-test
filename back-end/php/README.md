Backend - PHP
===============

## Tasks

The list of the tasks listed in the [Readme](../Readme.md) in the root directory of the bitbucket project.

## Requirement

* You need to have composer installed [Download composer](https://getcomposer.org/download/)

* You need to install docker desktop and Sign Up [Download Docker](https://www.docker.com/products/docker-desktop)

## Notes

* You are free to organize the content of the folder `src` as you want.
* You are free to install any packages

## How to run the project

### Install the application

```bash
composer install
```

### Run API

#### Without docker

1) update the database host in the config `config/application.ini` with the following value

``` 
host = "127.0.0.1"
```

2) Run the PHP server

```bash 
composer start
```

#### With docker

1) update the database host in the config `config/application.ini` with the following value

``` 
host = "host.docker.internal"
```

Note: be sure you have in your host file the following line (MacOS or Linux: `/etc/hosts` Windows `C:\Windows\System32\drivers\etc/hosts`)

```
127.0.0.1 host.docker.internal
```

2) Spawn containers

```bash 
docker-compose up
```

### Check tha api

Check that everything is ok

```bash
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

### Unit Tests

```bash
composer test
```

### Documentation

* This app use [PHP Slim Framework](https://www.slimframework.com/). This is a web framework written in PHP.
* Complementary documentation can be found here: https://odan.github.io/2019/11/05/slim4-tutorial.html
