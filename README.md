
# Order Service

Sample Golang With Echo Framework with system order 


## Run Locally

Clone the project

```bash
  git clone https://github.com/anggaramp/order-service.git
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies and run docker 

```bash
  make setup
```
Create Database

```bash
  import file setup.sql for create database
```

Setup DB Connection

```bash
  Update file .env.local on host DB base on your system
```
Start the server

```bash
  make run
```
Migration Your Table

```bash
 Hit Endpoint /v1/admin/migration on postman
```
