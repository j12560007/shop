# Shop API

## Introduction

```
This is an API project that contains shop features.

```

## Build APP

```
go build main.go

```

## Run APP

```
./main.exe

```

## Framework

```
gin-gonic

```

## Project Structure

```
|   main.go                 // App Entrance
|   go.mod
|   go.sum
|   README.md
|
+---config                  // App Configurations
|     server_conf.go
|     server.toml
|
+---dao                     // Connection And Functions Of Database
|     order_dao.go
|     product_dao.go
|     user_dao.go
|
+---entity                  // Structures For Request And Response
|     order_entity.go
|     product_entity.go
|     user_info_entity.go
|
+---route                   // Routers And Controllers
|     router.go
|
+---services                // Service Of API
|     manager_service.go
|     order_service.go
|     product_service.go
|     user_service.go
|
+---tool                    // Tool Usage
|     error_tool.go
|     tool.go
|

```

## Database Diagram

```
![database](https://github.com/j12560007/shop/blob/main/shop_db.jpg)

```
