# Go Fiber with MongoDB

## Introduction:

Fiber is a quickly growing super-fast backend framework, mainly written in golang with performance in mind. This repository intends to use go fiber along with MongoDB a NoSQL database to emulate a generic CRUD operations project and hopefully introduce proper structuring in order to have code that is easily maintanable and expandable.

This is a simple todo app, with basic CRUD functionality such as insertion, reading one/many, updation, deletion of one/many emulating basic operations needed by majority of DB based applications for easier adaptation.

## Organization:

### 1. config:

This has the module responsible for loading data from creds.env and returning it as environment variable values.

### 2. database:

Module is responsible for connection to database and maintaining the reference to database and particular collection operations are performed on.

### 3. handler:

Main db and route logic required for handling the endpoints.

### 4. middleware:

This project does not use any middleware but if used can go over here.

### 5. model:

Since all CRUD operations are done based with models, this helps maintain structure of data being sent to and received from clients. This project has a todo model which is used for all the CRUD operations.

### 6. router:

Sets up the routes for the endpoints, with logic provided by handler modules.

### 7. creds.env:

File that needs to be created from which values such as port, db credentials are loaded, can be directly set as environmental variables as well.

## Setup:

### If building from source:

1. Once this repository is cloned and golang is installed in the system, navigate to this directory and run

```
go mod download
```

2. Once the dependencies are downloaded, using sample.env as reference either create a file called creds.env with the same keys or directly configure same keys as environment variables.
3. After the configuration and ensuring that the db is operational, run either

```
go build # to get the executable to run

or

go run main.go
```

### Alternatively:

Just run the already built executable if you are just trying it out with a database. Executable is compiled for Linux only at the moment.

## Announcement:

If you find this repository useful to you in any way, do help me by starring this repository and let your friends know about this repo. I'll try my best to keep it updated.
