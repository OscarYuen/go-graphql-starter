## Go Graphql Starter
[![GitHub license](https://img.shields.io/github/license/OscarYuen/go-graphql-starter.svg)](https://github.com/OscarYuen/go-graphql-starter/blob/master/LICENSE)


This project aims to use [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go) to build a starter web application. This project has already been used as backend application in production. 

In case you need to get called from another frontend side, CORS may needed to be enabled in this application as this project mainly focuses on backend logic at this stage. 


This project would be continuously under development for enhancement. Pull request is welcome. 


#### RoadMap:
- [x] Integrated with sqlx
- [x] Integrated with graphql-go
- [x] Use go-bindata to generate Go code from .graphql file
- [x] Use psql
- [x] Integrated with dataloader
- [x] Add authentication & authorization
- [ ] Add unit test cases
- [ ] Support subscription
- [ ] Support web-socket notification and messaging

#### Structure
    go-graphql-starer
    │   README.md
    │   server.go           --- the entry file
    │   Config.toml         --- the configuration file for setting server parameter
    │   Dockerfile
    │   Gopkg.lock          --- generated file from dependency management tool, dep
    │   Gopkg.toml          --- generated file from dependency management tool, dep
    │   graphiql.html       --- the html for graphiql which is used for testing query & mutation
    └───context             --- application context like db configuration
    └───data                --- storing the sql data patch for different version
    │   └───1.0             --- storing sql data patch for version 1.0
    │      └───...          --- sql files
    └───handler             --- the handler used for chaining http request like authentication, logging etc.
    └───loader              --- implementation of dataloader for caching and batching the graphql query
    └───model               --- the folder putting struct file
    └───resolver            --- implementation of graphql resolvers
    └───schema              --- implementation of graphql resolvers
    │   │   schema.go       --- used for generate go code from static graphql files inside 'type' folder
    │   │   schema.graphql  --- graphql root schema
    │   └───type            --- folder for storing graphql schema in *.graphql format
    │       └───...         --- graphql schema files in *.graphql format
    └───service             --- services for users, authorization etc.
    └───util                --- utilities
    
#### Requirement:

1. Postgres database
2. Golang

Remark: If you want to use other databases, please feel free to change the driver in `config/db.go`

#### Usage(Without docker):

1. Run the sql scripts under `data/1.0` folder inside Postgres database console

2. Install go-bindata
    ```
    go get -u github.com/jteeuwen/go-bindata/...
    ```

3. Setup GOPATH (Optional if already set)

    For example: MacOS
    ```
    export GOPATH=/Users/${username}/go
    export PATH=$PATH:$GOPATH/bin
    ```

4. Run the following command at root directory to generate Go code from .graphql file
    ```
    go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
    ```

    OR

    ```
    go generate ./schema
    ```
    There would be bindata.go generated under `schema` folder

5. Start the server (Ensure your postgres database is live and its setting in Config.toml is correct)
    ```
    go build server.go
    ```
    
#### Usage(With docker):

1. Run the sql scripts under `data/1.0` folder inside Postgres database console

2. Build docker image
    ```
    docker build -t go-graphql-starter .
    ```

3. Run docker image (Ensure your database setting in Config.toml is correct)
    ```
    docker run go-graphql-starter
    ```

#### Usage(With docker-compose):

1. Create a folder `/psqldata` on your OS system and set it for file sharing in docker

2. Create and starter services by docker-compose
    ```
    docker-compose up
    ```
    
        
#### Graphql Example:

Test in graphiql by the following endpoint

```
localhost:3000
```

Basically there are two graphql queries and one mutation

##### Query:
1. Get a user by email
2. Get user list by cursor pagination

##### Mutation:

To query a list of users, you need to be authenticated.
Authentication is not required for other operations.
In order to authenticate, here are the steps to follow:

1. Create a user

```graphql
mutation {
  createUser (email: "tester@tester.com", password: "123456") {
    id
  }
}
```

2. Log in by submitting your email and password through a Basic Authorization Header.

Here's an example on how to achieve this:

       a. Download [Insomnia](https://insomnia.rest/) OR Other RESTful endpoint testing tools e.g. Postman

       b. Create a new POST request, paste `localhost:3000/login` in the URL bar and go to the Header tab

       c. Generate your basic Authorization on [blitter.se](https://www.blitter.se/utils/basic-authentication-header-generator/)

       d. In Insomnia, first column, type Authorization, second column enter the Basic you just copied `Basic dGVzdGVyQHRlc3Rlci5jb206a3Rta3Rt`

       e. Click send, you should get a jwt token back.

You can change the Authorization of request header in `graphiql.html` and restart the server to see the effect of authentication using token

#### Test:

- Run Unit Tests
    ```
    go test
    ```
    
#### Reference

-[graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)

-[tonyghita/graphql-go-example](https://github.com/tonyghita/graphql-go-example)
