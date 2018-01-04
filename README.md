## Go Graphql Starter

This project aims to use [neelance/graphql-go](https://github.com/neelance/graphql-go) to build a starter web application. This project would be continuously under development. Pull request is welcome. 

#### RoadMap:
- [x] Integrated with sqlx
- [x] Integrated with graphql-go
- [x] Use go-bindata to generate Go code from .graphql file
- [ ] Use psql
- [ ] Integrated with dataloader
- [ ] Add authentication & authorization
- [ ] Add unit test cases
- [ ] Support subscription

#### Usage:

1. Install go-bindata
    ```
    go get -u github.com/jteeuwen/go-bindata/...
    ```

2. Setup GOPATH

    For example: MacOS
    ```
    export GOPATH=/Users/${username}/go
    export PATH=$PATH:$GOPATH/bin
    ```

3. Run the following command at root directory to generate Go code from .graphql file
    ```
    go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
    ```
    
    There would be bindata.go generated under `schema` folder
    
4. Start the server
    ```
    go build server.go
    ```
    
#### Test:

- Run Unit Tests
    ```
    go test
    ```
    
#### Reference

-[neelance/graphql-go](https://github.com/neelance/graphql-go)

-[tonyghita/graphql-go-example](https://github.com/tonyghita/graphql-go-example) 
