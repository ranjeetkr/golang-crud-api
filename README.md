# GoLang CRUD operation with mysql DataBase


## General information
It's just a Book managagemnt system where we are doing basic CRUD operation

## Install the necessary packages
Clone this project, and install the below dependency
Run this command:

```
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
```
This command will install the basic goLang dependency.


## How to run  
```
go run main.go
```
## Description 
This goLang CRUD API having with the below API endpoint

- **POST : http://localhost:<port>/books** - API to save the book detail
- **GET : http://localhost:<port>/books/<id>** - API to get the book detail by book ID
- **GET : http://localhost:<port>/books** - API to list all the books
- **PUT : http://localhost:<port>/books/<id>** - API to update book details
- **DELETE : http://localhost:<port>/books/id** - API to delete book by id

