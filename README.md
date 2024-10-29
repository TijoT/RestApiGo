<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [REST API with Go and Gin](#rest-api-with-go-and-gin)
  - [Prerequisites](#prerequisites)
  - [Layout](#layout)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# REST API with Go and Gin

Project demonstrates REST API implementation with Go and Gin framework

## Prerequisites
Install external dependencies 
* [gin](https://github.com/gin-gonic/gin)
* [jwt](https://github.com/golang-jwt/jwt)
* [go-sqlite3](https://github.com/mattn/go-sqlite3)
* [crypto](https://github.com/golang/crypto)

## Layout

.
|-- README.md
|-- api-test
|   |-- README.md
|   |-- cancel-registration.http
|   |-- create-event.http
|   |-- create-user.http
|   |-- delete-event.http
|   |-- get-events.http
|   |-- get-single-event.http
|   |-- login.http
|   |-- register.http
|   `-- update-event.http
|-- api.db
|-- db
|   |-- README.md
|   `-- db.go
|-- go.mod
|-- go.sum
|-- main.go
|-- middlewares
|   |-- README.md
|   `-- auth.go
|-- models
|   |-- README.md
|   |-- event.go
|   `-- user.go
|-- routes
|   |-- README.md
|   |-- events.go
|   |-- register.go
|   |-- routes.go
|   `-- users.go
`-- utils
    |-- README.md
    |-- hash.go
    `-- jwt.go