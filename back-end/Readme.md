Automotive Data Solution - Backend skill test
==========================

### API
This Api allows a user to get information about vehicles and to manage the vehicle in general (make, model,...)

### Requirements

* You need to install docker desktop and Sign Up [Download Docker](https://www.docker.com/products/docker-desktop)

### Instructions

* You will have to implement the API in PHP, in Golang or in Python.
* A framework is imposed in all the languages

You can complete the tasks in any order.

We recommend to set aside 2 hours for the test but it is not necessary/expected that you complete all the tasks.


### Database Structure

![Database structure](./data/img/gotest_mpd.png "Database Structure")


### Tasks

* TASK 1: As a user, I want to list all the vehicle model
    - Implement a new endpoint to do this action
* TASK 2: As a user, I want to add a vehicle make
    - Update the database migration script `data/sql/update_database.sql` add a column `created` in all the tables existing. (This column contain the date and the time when the data has been added)
    - Implement a new endpoint to do this action
* TASK 3: As a user, I want to update a vehicle make
    - Update the database migration script `data/sql/update_database.sql` add a column `updated` in all the tables existing. (This column contain the date and the time when the data has been changed)
    - Implement a new endpoint to do this action
* TASK 4: As a user, I want to be able to remove a vehicle make logically
    - Update the database migration script `data/sql/update_database.sql` add a column `state` in all the tables existing. This column can take 2 values 0/1. (0 means the entry has been deleted)
    - Implement a new endpoint to do this action

* TASK 5: Secure the call to insert / update / delete data with a JWT token

* TASK 6: As a user, I want to know all the vehicles available (means state=1) by make/year/model
    - Implement a new endpoint to do this action
* TASK 7: As a user, I want to be able to add a new vehicle
    - Update the database migration script `data/sql/update_database.sql` to avoid to have duplicate data for this couple (make,year,model)  
    - Implement a new endpoint to do this action

* TASK 8: As a user, I want to be able to paginate the list of vehicle result
    - Implement a new endpoint to do this action

* TASK 9: Add some Unit Test (two at least)
    - Implement a new endpoint to do this action


Extra tasks:

* As a user, I want to add a vehicle model
* As a user, I want to update a vehicle model
* As a user, I want to delete a vehicle model logically
* As a user, I want the detail of a vehicle model


Note: At least provide us a list of curl or a postman collection (much appreciated) of all the endpoints you have created.


### Documentation

We use docker compose to have a database MySQL in local accessible by this project.
The documentation of how to use *docker-compose* can be found here: https://docs.docker.com/compose/

The documentation of JWT token can be found there https://jwt.io/


### How to submit your work?

Create a zip file and send it by email.
