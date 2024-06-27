Database - Mysql
==========================

### Requirements

* You need to install docker desktop and Sign Up [Download Docker](https://www.docker.com/products/docker-desktop)

### Database Structure

![Database structure](../data/mpd.png "Database Structure")

### Credentials Database

``` 
USER: usrtest
PASSWORD: pwdtest
ROOT_PASSWORD: pwdtest
DATABASE: local_dbtest
```

### Start/Stop MySQL

1) Start the database MySQL
```
./runmysql.sh  #Choose start in the contextual menu
```

2) Stop the database MySQL
```
./runmysql.sh  #Choose stop in the contextual menu
```

3) Stop and Purge the database MySQL
```
./runmysql.sh  #Choose purge in the contextual menu
```

**IMPORTANT**: Everytime you modify the file `sql/update_database.sql` you have to run the script `runmysql.sh` with the command *purge*

Note: You can apply directly your SQL queries to test them in your favorite Mysql UI by using the credentials.
