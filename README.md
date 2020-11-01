# iseng

CRUD API Boilerplate

[Installation](#installation)

## Installation

### Prerequisites
- Golang(>=1.11) - Download and Install [Golang](https://golang.org/)
- Swaggo - Download and Install [Swaggo](https://github.com/swaggo/swag)
- Mockery - Download and Install [Mockery](https://github.com/vektra/mockery)

### How To Use
- Download Dependency Golang

This command is useful for download dependency golang
```
    $ make download
```

- Running Unit Test

This command is useful for running unit testing golang
```
    $ make unittest
```

- Running Test

This command is useful for running testing golang
```
    $ make test
```

- Running Prepare Linter

This command is useful for prepare linter golang
``` 
    $ make lint-prepare
```

- Running Linter

This command is useful for check quality of golang code
```
    $ make lint
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