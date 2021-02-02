# ACCOUNT SERVICE

Simple account service. This service is covers user-registration, user-login and some example end points which need an authorization (With OAuth 2.0). The database only support for PostgreSQL.

## Contents
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Quick Start](#quick-start)

## Tech stack
* Go, with modules:
    * [gin](github.com/gin-gonic/gin)
    * [godotenv](github.com/joho/godotenv)
    * [gorm](github.com/jinzhu/gorm)
    * [validator](github.com/go-playground/validator/v10)
    * [redigo](github.com/gomodule/redigo)
    * [jwt-go](github.com/dgrijalva/jwt-go)
* PostgreSQL
* Redis
* JWT

## Prerequisites
* Go v1.14
* PostgreSQL
* Redis

## Setup
1. Clone this project
```
git clone https://github.com/yaumulisnain/account-service
```
2. Setup your database.
    * You can copy and run DDL script from ```src/migration/00_ddl.sql``` for create database, or create your database manually.
    * You also need to copy and run SQL script from ```src/migration/01_all_tables.sql``` for create all tables.
    * You can copy and run SQL script from ```src/migration/02_dummy.sql``` for create dummy user data, or insert manually.

3. You must create ```.env``` file, you can copy from ```.env.dev``` and then customize the value.
```
APP_ENV=dev
TZ=UTC
PORT=:3000

DB_HOST=localhost
DB_DATABASE=db_account
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_PORT=5432

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=11

JWT_SECRET_KEY=97a89sdasdkada7s
JWT_OWNER=account-service
```

## Quick Start
This project using Makefile for simplify your commands.

1. Run project
```
$ make run
```

2. Build project
```
$ make build
```

3. Run binary file
```
$ make start
```