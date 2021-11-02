# iseng

CRUD API Boilerplate

[Installation](#installation)

## Installation

### Prerequisites
- Golang(>=1.11) - Download and Install [Golang](https://golang.org/)
- Redis - Download and Install [Redis](https://redis.io/download)
- MySQL - Download and Install [MySQL](https://www.apachefriends.org/download.html)

### How To Use
- Download Dependency Golang

This command is useful for download dependency golang
```
    $ make download
```

This command is useful for migrate dummy data
```
    $ make migrate
```

- Running Server

This command is useful for running and generate swagger documentation
```
    $ make run
```

- Create Mock

This Command is useful for creating mock file from interface{}
```
    $ mockery --name=InterfaceName --output=cmd/mocks  --recursive
```
